#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error

# Needed for submodules
git config --global --add safe.directory '*'
if [[ $# -eq 1 ]] && [[ "$1" == "--local-tests-only" ]]; then
    dir_prefix=.
else
    dir_prefix=/code
fi

rm -f $dir_prefix/NETWORK_READY

until cat $dir_prefix/NETWORK_READY &> /dev/null
do
    if [[ retry_count -ge 300 ]]; then
        echo "Subnets didn't start up quickly enough."
        exit 1
    fi
    echo "Waiting for subnets to start up. Retry count: $retry_count"
    retry_count=$((retry_count+1))
    sleep 2
done

# Source all variables set in run_setup.sh
set -a
source $dir_prefix/vars.sh || true # ignore readonly variable errors
set +a

# Relayer constants
reward_address=0xA100fF48a37cab9f87c8b5Da933DA46ea1a5fb80
account_private_key=C2CE4E001B7585F543982A01FBC537CFF261A672FA8BD1FAFC08A207098FE2DE

# Populate the relayer configuration
relayerConfigFile=./relayerConfig.json
rm -f $relayerConfigFile

# Avoid using sed -i due to Docker + MacOS M1 issues
sed "s/<SUBNET_A_BLOCKCHAIN_ID>/$subnet_a_blockchain_id/g
     s/<SUBNET_A_SUBNET_ID>/$subnet_a_subnet_id/g
     s/<SUBNET_B_BLOCKCHAIN_ID>/$subnet_b_blockchain_id/g
     s/<SUBNET_B_SUBNET_ID>/$subnet_b_subnet_id/g
     s/<SUBNET_C_BLOCKCHAIN_ID>/$subnet_c_blockchain_id/g
     s/<SUBNET_C_SUBNET_ID>/$subnet_c_subnet_id/g
     s/<TELEPORTER_CONTRACT_ADDRESS>/$teleporter_contract_address/g
     s/<C_CHAIN_BLOCKCHAIN_ID>/$c_chain_blockchain_id/g
     s/<REWARD_ADDRESS>/$reward_address/g
     s/<ACCOUNT_PRIVATE_KEY>/$account_private_key/g" ./docker/relayerConfigTemplate.json > $relayerConfigFile

/usr/bin/awm-relayer --config-file $relayerConfigFile
