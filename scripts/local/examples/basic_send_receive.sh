#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # stop on first error

##
## Variables provided by run_setup.sh
##

my_pk=$user_private_key
my_0x=$user_address
c_bid=$c_chain_blockchain_id
a_bid=$subnet_a_blockchain_id
c_sid=$c_chain_subnet_id
a_sid=$subnet_a_subnet_id
c_rpc=$c_chain_rpc_url
a_rpc=$subnet_a_rpc_url
c_bix=$c_chain_blockchain_id_hex
a_bix=$subnet_a_blockchain_id_hex
tp_0x=$teleporter_contract_address
wm_0x=$warp_messenger_precompile_addr

##
## Test covers:
## - Sending bidirectional cross chain messages between C-chain and a subnet chain,
##   by calling the teleporter contract sendCrossChainMessage function directly.
## - Checking message delivery for message that was sent on destination chain.
## - Calling to the warp precompile address directly for blockchain ID.
##

cd contracts

##
## Deploy an ERC20 to be used in the E2E test
##

erc=$(forge create \
    --rpc-url $c_rpc --private-key $my_pk src/Mocks/ExampleERC20.sol:ExampleERC20)
erc_c0x=$(parseContractAddress "$erc")
echo "ERC20 deployed to $erc_c0x on C-chain"

erc=$(forge create \
    --rpc-url $a_rpc --private-key $my_pk src/Mocks/ExampleERC20.sol:ExampleERC20)
erc_a0x=$(parseContractAddress "$erc")
echo "ERC20 deployed to $erc_a0x on subnet A"

##
## Send from C-chain to subnet A
##

echo "Sending from C-chain to subnet A"
bid=$(cast call $wm_0x "getBlockchainID()(bytes32)" --rpc-url $c_rpc)
echo "Got blockchain ID $bid"

echo "Sending call to teleporter $tp_0x $a_bix $c_rpc"
mid=$(cast call $tp_0x "getNextMessageID(bytes32)(bytes32)" $a_bix --rpc-url $c_rpc)
echo "Next message ID for subnet $a_bix is $mid"

##
## Directly send a TX to teleporter's sendCrossChainMessage
##

sccm_tgt_id=$a_bix
sccm_tgt_0x=abcedf1234abcedf1234abcedf1234abcedf1234
send_bytes32=000000000000000000000000$sccm_tgt_0x
sccm_fee=255
sccm_gas=10000
sccm_data=cafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafe

##
## Approve teleporter to some ERC20 from user for TXs
##

approve_u256=100000000000000000000000000000
cast send $erc_c0x \
    "approve(address,uint256)(bool)" $tp_0x $approve_u256 \
    --rpc-url $c_rpc --private-key $my_pk
u256=$(cast call $erc_c0x \
    "allowance(address,address)(uint256)" $my_0x $tp_0x \
    --rpc-url $c_rpc)
u256=$(echo $u256 | cut -d' ' -f1)
if [[ $u256 -ne $approve_u256 ]]; then
    echo $u256
    echo $approve_u256
    echo "Approving teleporter to spend ERC20 from user failed"
    exit 1
else
    echo "Approved teleporter to spend the test ERC20 token from user"
fi

mid=$(cast call $tp_0x \
    "sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes))(bytes32)" \
    "($sccm_tgt_id,$sccm_tgt_0x,($erc_c0x,$sccm_fee),$sccm_gas,[],$sccm_data)" \
    --rpc-url $c_rpc --from $my_0x)
echo "Got starting ID $mid to teleport address $tp_0x"
cast send $tp_0x \
    "sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes))(bytes32)" \
    "($sccm_tgt_id,$sccm_tgt_0x,($erc_c0x,$sccm_fee),$sccm_gas,[],$sccm_data)" \
    --rpc-url $c_rpc --private-key $my_pk

idx=0
msg=$(cast call $tp_0x "messageReceived(bytes32)(bool)" $mid --rpc-url $a_rpc)
until [[ $msg == "true" ]]; do
    if [[ idx -ge 10 ]]; then
        echo "Destination chain on subnet A did not receive message before timeout"
        exit 1
    fi
    echo "Waiting for destination chain on subnet A to receive message ID $mid; retry count: $idx"
    sleep 3
    msg=$(cast call $tp_0x "messageReceived(bytes32)(bool)" $mid --rpc-url $a_rpc)
    idx=$((idx + 1))
done

echo "Received on subnet A is $msg"

##
## Send from subnet A to C-Chain
##

echo "Sending from subnet A to C-Chain"
bid=$(cast call $wm_0x "getBlockchainID()(bytes32)" --rpc-url $a_rpc)
echo "Got blockchain ID $bid"

echo "Sending call to teleporter $tp_0x $c_bix $a_sid"
mid=$(cast call $tp_0x "getNextMessageID(bytes32)(bytes32)" $c_bix --rpc-url $a_rpc)
echo "Next message ID for subnet $c_bix is $mid"

##
## Directly send a few TXs to teleporter's sendCrossChainMessage
##

sccm_tgt_id=$c_bix
sccm_tgt_0x=abcedf1234abcedf1234abcedf1234abcedf1234
send_bytes32=000000000000000000000000$sccm_tgt_0x
sccm_data=cafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafebabecafe

##
## Approve teleporter to some ERC20 from user for TXs
##

cast send $erc_a0x \
    "approve(address,uint256)(bool)" $tp_0x $approve_u256 \
    --rpc-url $a_rpc --private-key $my_pk
mid=$(cast call $erc_a0x \
    "allowance(address,address)(uint256)" $my_0x $tp_0x \
    --rpc-url $a_rpc)
mid=$(echo $mid | cut -d' ' -f1)
if [[ $mid -ne $approve_u256 ]]; then
    echo $mid
    echo "Approving teleporter to spend ERC20 from user failed"
    exit 1
else
    echo "Approved teleporter to spend the test ERC20 from user"
fi

mid=$(cast call $tp_0x \
    "sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes))(bytes32)" \
    "($sccm_tgt_id,$sccm_tgt_0x,($erc_a0x,$sccm_fee),$sccm_gas,[],$sccm_data)" \
    --rpc-url $a_rpc --from $my_0x)
echo "Got starting ID $mid to teleport address $tp_0x"
echo "Got IDs $a_bix $c_bix $a_sid $c_sid"
cast send $tp_0x \
    "sendCrossChainMessage((bytes32,address,(address,uint256),uint256,address[],bytes))(bytes32)" \
    "($sccm_tgt_id,$sccm_tgt_0x,($erc_a0x,$sccm_fee),$sccm_gas,[],$sccm_data)" \
    --rpc-url $a_rpc --private-key $my_pk

idx=0
msg=$(cast call $tp_0x "messageReceived(bytes32)(bool)" $mid --rpc-url $c_rpc)
until [[ $msg == "true" ]]; do
    if [[ idx -ge 10 ]]; then
        echo "C-Chain did not receive message before timeout"
        exit 1
    fi
    echo "Waiting for C-chain to receive message ID $mid; retry count: $idx"
    sleep 3
    msg=$(cast call $tp_0x "messageReceived(bytes32)(bool)" $mid --rpc-url $c_rpc)
    idx=$((idx + 1))
done

echo "Received on C-chain is $msg"
