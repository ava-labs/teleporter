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
#   subnet_c_chain_id
#   subnet_a_subnet_id
#   subnet_b_subnet_id
#   subnet_c_subnet_id
#   subnet_a_url
#   subnet_b_url
#   subnet_c_url
#   subnet_a_chain_id_hex
#   subnet_b_chain_id_hex
#   subnet_c_chain_id_hex
#   subnet_a_subnet_id_hex
#   subnet_b_subnet_id_hex
#   subnet_c_subnet_id_hex
#   teleporter_contract_address
#   registry_address_a
#   registry_address_b
#   registry_address_c

latestVersion=$(cast call $registry_address_a "getLatestVersion()(uint256)" --rpc-url $subnet_a_url)
echo "Got latest Teleporter version $latestVersion"

result=$(cast call $registry_address_a "getAddressToVersion(address)(uint256)" $teleporter_contract_address --rpc-url $subnet_a_url)
echo "Teleporter address $teleporter_contract_address is at version $result"

latestAddress=$(cast call $registry_address_a "getVersionToAddress(uint256)(address)" $latestVersion --rpc-url $subnet_a_url)
echo "Latest Teleporter address is $latestAddress"