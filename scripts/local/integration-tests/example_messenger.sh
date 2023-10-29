#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error

# Variables provided by run_setup.sh:
#   c_chain_url
#   user_private_key
#   user_address_bytes
#   user_address
#   relayer_address
#   subnet_a_chain_id
#   subnet_b_chain_id
#   subnet_a_subnet_id
#   subnet_b_subnet_id
#   subnet_a_url
#   subnet_b_url
#   subnet_a_chain_id_hex
#   subnet_b_chain_id_hex
#   subnet_a_subnet_id_hex
#   subnet_b_subnet_id_hex
#   teleporter_contract_address
#   warp_messenger_precompile_addr

# Deploy a test ERC20 on subnet A.
cd contracts
erc20_deploy_result=$(forge create --private-key $user_private_key src/Mocks/ExampleERC20.sol:ExampleERC20 --rpc-url $subnet_a_url)
erc20_contract_address=$(parseContractAddress "$erc20_deploy_result")
echo "Test ERC20 contract deployed to $erc20_contract_address on Subnet A"

# Deploy the example messenger application on subnet A
example_messenger_a_deploy_result=$(forge create --private-key $user_private_key --constructor-args $teleporter_contract_address \
    --rpc-url $subnet_a_url src/CrossChainApplications/ExampleMessenger/ExampleCrossChainMessenger.sol:ExampleCrossChainMessenger)
example_messenger_a_contract_address=$(parseContractAddress "$example_messenger_a_deploy_result")
echo "Example Messenger contract deployed to subnet A at $example_messenger_a_contract_address"

# Deploy the example messenger application on subnet B
example_messenger_b_deploy_result=$(forge create --private-key $user_private_key --constructor-args $teleporter_contract_address \
    --rpc-url $subnet_b_url src/CrossChainApplications/ExampleMessenger/ExampleCrossChainMessenger.sol:ExampleCrossChainMessenger)
example_messenger_b_contract_address=$(parseContractAddress "$example_messenger_b_deploy_result")
echo "Example Messenger contract deployed to subnet B at $example_messenger_b_contract_address"

# Approve the example messenger contract on subnet A spent ERC20 tokens from the user account we're using to send transactions
approve_amount=100000000000000000000000
cast send $erc20_contract_address "approve(address,uint256)(bool)" $example_messenger_a_contract_address \
    $approve_amount --private-key $user_private_key --rpc-url $subnet_a_url
result=$(cast call $erc20_contract_address "allowance(address,address)(uint256)" $user_address $example_messenger_a_contract_address --rpc-url $subnet_a_url)
if [[ $result != $approve_amount ]]; then
    echo $result
    echo "Error approving example messenger contract to spend ERC20 from user account."
    exit 1
fi
echo "Approved the subnet A example messenger to spend the test ERC20 token from the user account."

# Send a transaction to the smart contract that calls the sendCrossChainMessage precompile method.
send_cross_chain_message_fee_amount=100000
send_cross_chain_message_required_gas_limit=300000
send_cross_chain_message_message_string="hello world!"
cast send $example_messenger_a_contract_address "sendMessage(bytes32,address,address,uint256,uint256,string)(uint256)" \
    $subnet_b_chain_id_hex \
    $example_messenger_b_contract_address \
    $erc20_contract_address \
    $send_cross_chain_message_fee_amount \
    $send_cross_chain_message_required_gas_limit \
    "$send_cross_chain_message_message_string" \
    --private-key $user_private_key --rpc-url $subnet_a_url
echo "Sent a transaction to sendCrossChainMessage via contract."

# Wait for the cross chain message to be processed by a relayer.
sleep 10

# Check that the message has been delivered to the warp messenger contract on subnet B.
function stringToLower() {
    python3 -c "import sys; input = \"$1\"; sys.stdout.write(input.lower())"
}

origin_contract_address_lower=$(stringToLower $example_messenger_a_contract_address)
result=$(cast call $example_messenger_b_contract_address "getCurrentMessage(bytes32)(address,string)" $subnet_a_chain_id_hex --rpc-url $subnet_b_url)
echo "Raw result is: \"$result\""
result_arr=($result)

actual_sender_address=${result_arr[0]}
if [[ $(stringToLower $actual_sender_address) != $origin_contract_address_lower ]]; then
    echo $result
    echo "Get current message returned unexpected origin sender address."
    echo "Expected: $origin_contract_address_lower"
    exit 1
fi
# Check that the message matched the one submitted to subnet A.
# The message "hello world!" is returned without quotations, so it is split up as two different values by bash across indices 1 and 2 of the result.
if [[ "${result_arr[1]} ${result_arr[2]}" != "$send_cross_chain_message_message_string" ]]; then
    echo "Get current message returned unexpected message."
    echo "Actual: ${result_arr[1]} ${result_arr[2]}"
    echo "Expected: $send_cross_chain_message_message_string"
    exit 1
fi
echo "The latest message from chain ID $subnet_a_chain_id was: "
echo "    Message Sender: ${result_arr[0]}"
echo "    Message: \"$send_cross_chain_message_message_string\""
echo "Successfully called getCurrentMessage."
