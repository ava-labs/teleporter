#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error

# Variables provided by run_test.sh:
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

# Test covers:
# - Deploying block hash publisher smart contract, which is built on top of teleporter.
# - Publish block hash information cross chain through teleporter, and checking its delivery on destination chain.

#Catch errors for undefined environment variables necessary for this script
# Variable names to check
variables_to_check=(
   "c_chain_url"
   "user_private_key"
   "user_address_bytes"
   "user_address"
   "relayer_address"
   "subnet_a_chain_id"
   "subnet_b_chain_id"
   "subnet_a_subnet_id"
   "subnet_b_subnet_id"
   "subnet_a_url"
   "subnet_b_url"
   "subnet_a_chain_id_hex"
   "subnet_b_chain_id_hex"
   "subnet_a_subnet_id_hex"
   "subnet_b_subnet_id_hex"
   "teleporter_contract_address"
)

missing_variables=()

# Iterate through the array and check if each variable is defined
for var_name in "${variables_to_check[@]}"; do
    if [ -z "${!var_name}" ]; then
        missing_variables+=("$var_name")
    fi
done

# Check if there are missing variables and display an error message
if [ ${#missing_variables[@]} -gt 0 ]; then
    echo -e "\e[1;31mError:\e[0m The following variables are not defined: ${missing_variables[*]}"
    echo
    echo -e "Before running this script, please wait until the network setup finishes and run \e[1mset -a\e[0m and \e[1msource vars.sh\e[0m to ensure the necessary variables are set."
    exit 1  # Exit with an error code
fi

# Deploy the block hash publisher to subnet A
cd contracts
block_hash_publisher_deploy_result=$(forge create --private-key $user_private_key --constructor-args $teleporter_contract_address  \
    --rpc-url $subnet_a_url src/CrossChainApplications/VerifiedBlockHash/BlockHashPublisher.sol:BlockHashPublisher)
block_hash_publisher_contract_address=$(parseContractAddress "$block_hash_publisher_deploy_result")
echo "Block hash publisher contract deployed to subnet A at $block_hash_publisher_contract_address"

# Deploy the example messenger application on subnet B
block_hash_receiver_deploy_result=$(forge create --private-key $user_private_key \
    --constructor-args $teleporter_contract_address $subnet_a_chain_id_hex $block_hash_publisher_contract_address \
    --rpc-url $subnet_b_url src/CrossChainApplications/VerifiedBlockHash/BlockHashReceiver.sol:BlockHashReceiver)
block_hash_receiver_contract_address=$(parseContractAddress "$block_hash_receiver_deploy_result")
echo "Block hash receiver contract deployed to subnet B at $block_hash_receiver_contract_address"

# Send a transaction to the publisher contract to publish the current block hash to subnet B
cast send $block_hash_publisher_contract_address "publishLatestBlockHash(bytes32,address)(uint256)" $subnet_b_chain_id_hex $block_hash_receiver_contract_address \
     --private-key $user_private_key --rpc-url $subnet_a_url
echo "Sent a transaction to publish Subnet A's latest block hash to Subnet B."

# Wait for the cross chain message to be processed by a relayer.
sleep 5

# Call the receiver to see if the block hash was received.
result=$(cast call $block_hash_receiver_contract_address "getLatestBlockInfo()(uint256,bytes32)" --rpc-url $subnet_b_url)
result_arr=($result)
result_block_num=${result_arr[0]}
result_block_hash=${result_arr[1]}

if [[ $result_block_num == 0 ]]; then
    echo "Receiver did not have latest block number."
    exit 1
fi

if [[ "$result_block_hash" == "0x0000000000000000000000000000000000000000000000000000000000000000" ]]; then
    echo "Receiver did not have latest block hash."
    exit 1
fi

# TODO: Check that the block number and hash match on the origin chain by getting the block info.
echo "Successfully published block hash between subnets"
echo "Latest received Subnet A block number: $result_block_num"
echo "Latest received Subnet A block hash: $result_block_hash"

exit 0