#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# TODO We want to run the native token bridge with our Fuji scripts.
# This file isn't compeletely set up properly yet.

set -e # Stop on first error

# Variables provided by run_test.sh:
#   c_chain_url
#   user_private_key
#   user_address_bytes
#   user_address
#   relayer_address
#   subnet_a_chain_id
#   subnet_a_subnet_id
#   subnet_a_url
#   subnet_a_chain_id_hex
#   subnet_a_subnet_id_hex
#   teleporter_contract_address

# Test covers:
# - Creating bridged tokens on subnets different from the original subnet with the native token asset.
# - Bridging directly from token native subnet to a subnet with its bridged token
# - Bridging a wrapped token of one subnet to another subnet's bridged token of the same native token asset, through unwrapping then wrapping.
# - Unwrap bridged token back to origin subnet with native token asset.

# Deploy the ERC20 bridge contract to both chains.
bridge_a_deploy_result=$(forge create --private-key $user_private_key --constructor-args $teleporter_contract_address \
    --rpc-url $subnet_a_url src/CrossChainApplications/ERC20Bridge/ERC20Bridge.sol:ERC20Bridge)
ntm_a_address=$(parseContractAddress "$bridge_a_deploy_result")
echo "Native token minting contract deployed to subnet A at $ntm_a_address."

c_chain_deploy_result=$(forge create --private-key $user_private_key --constructor-args $teleporter_contract_address \
    --rpc-url $c_chain_url src/CrossChainApplications/ERC20Bridge/ERC20Bridge.sol:ERC20Bridge)
c_chain_ntr_address=$(parseContractAddress "$c_chain_deploy_result")
echo ""Native token bridge contract deployed to C Chain at $c_chain_ntr_address."

cd ..

# Send a transaction that initiates the ERC20 token being added to be bridged to subnet B
# Function signature for submitCreateBridgeToken includes an ERC20 type parameter, but we can pass in address
# since the contract solidity type gets mapped to address abi type https://docs.soliditylang.org/en/latest/abi-spec.html#mapping-solidity-to-abi-types.
create_bridge_token_message_fee=1000000000000000000
cast send $bridge_a_address "submitCreateBridgeToken(bytes32,address,address,address,uint256)" \
    $subnet_b_chain_id_hex \
    $ntm_b_address \
    $native_erc20_contract_address \
    $native_erc20_contract_address \
    $create_bridge_token_message_fee \
    --private-key $user_private_key --rpc-url $subnet_a_url
echo "Sent a transaction on Subnet A to add support for the the ERC20 token to the bridge on Subnet B."


# Wait for the cross chain message to be delivered by a relayer.
sleep 10

# Check that the bridge token was added on Subnet B
bridge_token_subnet_b_contract_address=$(cast call $ntm_b_address "nativeToWrappedTokens(bytes32,address,address)(address)" \
    $subnet_a_chain_id_hex \
    $bridge_a_address \
    $native_erc20_contract_address \
    --rpc-url $subnet_b_url)
echo "The bridge token contract address on Subnet B is $bridge_token_subnet_b_contract_address"
if [[ "$bridge_token_subnet_b_contract_address" == "0x0000000000000000000000000000000000000000" ]]; then
    echo "Bridge token contract was not created on Subnet B."
    exit 1
fi

# Now that the bridge token has been added, send a bridge transfer for the newly added token from subnet A to subnet B.
total_amount=13000000000000000000
primary_fee_amount=1000000000000000000
secondary_fee_amount=0
cast send $bridge_a_address "bridgeTokens(bytes32,address,address,address,uint256,uint256,uint256)" \
    $subnet_b_chain_id_hex \
    $bridge_b_address \
    $native_erc20_contract_address \
    $user_address \
    $total_amount \
    $primary_fee_amount \
    $secondary_fee_amount \
    --private-key $user_private_key --rpc-url $subnet_a_url
echo "Sent a transaction to bridgeTokens from Subnet A to Subnet B."

# Wait for the cross chain message to be processed by a relayer.
sleep 10

# Check that all the settings of the new bridge token are correct.
actual_native_chain_id=$(cast call $bridge_token_subnet_b_contract_address "nativeChainID()(bytes32)" --rpc-url $subnet_b_url)
echo "Bridge token native chain ID: $actual_native_chain_id"
actual_native_bridge=$(cast call $bridge_token_subnet_b_contract_address "nativeBridge()(address)" --rpc-url $subnet_b_url)
echo "Bridge token native bridge address: $actual_native_bridge"
actual_native_asset=$(cast call $bridge_token_subnet_b_contract_address "nativeAsset()(address)" --rpc-url $subnet_b_url)
echo "Bridge token native asset: $actual_native_asset"
actual_name=$(cast call $bridge_token_subnet_b_contract_address "name()(string)" --rpc-url $subnet_b_url)
echo "Bridge token name: $actual_name"
actual_symbol=$(cast call $bridge_token_subnet_b_contract_address "symbol()(string)" --rpc-url $subnet_b_url)
echo "Bridge token symbol: $actual_symbol"
actual_decimals=$(cast call $bridge_token_subnet_b_contract_address "decimals()(uint8)" --rpc-url $subnet_b_url)
echo "Bridge token decimals: $actual_decimals"
if [[ "$actual_native_chain_id" != "0x$subnet_a_chain_id_hex" ||
      "$actual_native_bridge" != "$bridge_a_address" ||
      "$actual_native_asset" != "$native_erc20_contract_address" ||
      "$actual_name" != "Mock Token" ||
      "$actual_symbol" != "EXMP" ||
      "$actual_decimals" != "18" ]]; then
    echo "Bridge token contract was not created or did not have expected values."
    echo "Actual native chain ID: $actual_native_chain_id, Expected: $subnet_a_chain_id_hex"
    echo "Actual native bridge: $actual_native_bridge, Expected: $bridge_a_address"
    echo "Actual native asset: $actual_native_asset, Expected: $native_erc20_contract_address"
    echo "Actual name: $actual_name, Expected: Mock Token"
    echo "Actual symbol: $actual_symbol, Expected: EXMP"
    echo "Actual decimals: $actual_decimals, Expected: 18"
    exit 1
fi

# Check the recipient balance of the new bridge token.
actual_bridge_token_balance=$(cast call $bridge_token_subnet_b_contract_address "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_b_url)
echo "Bridge token balance is $actual_bridge_token_balance"

if [[ "$actual_bridge_token_balance" != "12000000000000000000" ]]; then
    echo "Bridge token balance for $user_address did not match expected."
    echo "Actual balance: $actual_bridge_token_balance, Expected: 12000000000000000000"
    exit 1
fi
echo "Bridge token balance matches expected."

# Approve the bridge contract on subnet B to spent the wrapped tokens in the user account.
cast send $bridge_token_subnet_b_contract_address "approve(address,uint256)(bool)" $bridge_b_address 000000000000000000000000000000000000000000FFFFFFFFFFFFFFFFFFFFFF --private-key $user_private_key --rpc-url $subnet_b_url
result=$(cast call $bridge_token_subnet_b_contract_address "allowance(address,address)(uint256)" $user_address $bridge_b_address --rpc-url $subnet_b_url)
if [[ $result != 309485009821345068724781055 ]]; then # FFFFFFFFFFFFFFFFFFFFFF in decimal form is 309485009821345068724781055
    echo $result
    echo "Error approving bridge contract on subnet B to spend bridged ERC20 from user account."
    exit 1
fi
echo "Approved the subnet B bridge contract to spend the bridged ERC20 token from the user account."

# Unwrap bridged tokens back to subnet A.
total_amount=11000000000000000000
primary_fee_amount=1000000000000000000
secondary_fee_amount=1000000000000000000
cast send $bridge_b_address "bridgeTokens(bytes32,address,address,address,uint256,uint256,uint256)" \
    $subnet_a_chain_id_hex \
    $bridge_a_address \
    $bridge_token_subnet_b_contract_address \
    $user_address \
    $total_amount \
    $primary_fee_amount \
    $secondary_fee_amount \
    --private-key $user_private_key --rpc-url $subnet_b_url
echo "Sent a transaction to bridge tokens from Subnet B to Subnet A."

# Wait for the cross chain message to be delivered by a relayer.
sleep 10

# Check the redeemable reward balance of the relayer if the relayer address was set.
if [ ! -z "$relayer_address" ]; then
    actual_relayer_redeemable_balance=$(cast call $teleporter_contract_address "checkRelayerRewardAmount(address,address)(uint256)" $relayer_address $native_erc20_contract_address --rpc-url $subnet_a_url)
    echo "Native ERC20 token balance for relayer account is $actual_relayer_redeemable_balance"
    if [[ "$actual_relayer_redeemable_balance" != "2000000000000000000" ]]; then
        echo "Redeemable rewards for relayer account ($relayer_address) did not match expected."
        echo "Actual balance: $actual_relayer_redeemable_balance, Expected: 2000000000000000000"
        exit 1
    fi
    echo "Redeemable rewards balance matches expected for relayer on Subnet A."
fi

# Check the balance of the native token after the unwrap
actual_native_token_default_account_balance=$(cast call $native_erc20_contract_address "balanceOf(address)(uint256)" $user_address --rpc-url $subnet_a_url)
echo "Native ERC20 token balance for user account is $actual_native_token_default_account_balance"
if [[ "$actual_native_token_default_account_balance" != "9999999996000000000000000000" ]]; then
    echo "Native token balance for $user_address did not match expected."
    echo "Actual balance: $actual_native_token_default_account_balance, Expected: 9999999996000000000000000000"
    exit 1
fi
echo "Native token balance matches expected for user account."

# Check the balance of the native token for the relayer, which should have received the fee rewards
if [ ! -z "$relayer_address" ]; then
    actual_relayer_redeemable_balance=$(cast call $teleporter_contract_address "checkRelayerRewardAmount(address,address)(uint256)" $relayer_address $native_erc20_contract_address --rpc-url $subnet_a_url)
    echo "Redeemable ERC20 token reward balance for relayer account is $actual_relayer_redeemable_balance"
    if [[ "$actual_relayer_redeemable_balance" != "2000000000000000000" ]]; then
        echo "Redeemable reward balance for relayer account ($relayer_address) did not match expected."
        echo "Actual balance: $actual_relayer_redeemable_balance, Expected: 2000000000000000000"
        exit 1
    fi
    echo "Redeemable rewards balance matches expected for relayer on Subnet A."
fi


exit 0