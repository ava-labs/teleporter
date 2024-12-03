#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

ICM_CONTRACTS_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

# Check that foundry (specifically cast) is installed.
if ! command -v cast &> /dev/null; then
    echo "cast not found. You can install by calling $ICM_CONTRACTS_PATH/scripts/install_foundry.sh" && exit 1
fi

# Check that jq is installed.
if ! command -v jq &> /dev/null; then
    echo "jq not found. It is required to be installed before proceeding." && exit 1
fi

function printHelp() {
    echo "Usage: ./scripts/deploy_teleporter.sh --version <version> --rpc-url <url> [OPTIONS]"
    echo ""
    echo "Deploys a selected TeleporterMessenger contract to the specified chain"
    echo "For a list of releases, go to https://github.com/ava-labs/icm-contracts/releases"
    printUsage
}

function printUsage() {
    cat << EOF
Arguments:
    --version <version>              Specify the release version to deploy
    --rpc-url <url>                  Specify the rpc url of the node to use
Options:
    --private-key <private_key>      Private key of account to use to fund the Teleporter deployer address, if necessary.
    --help                           Print this help message
EOF
}

teleporter_version=
user_private_key=
rpc_url=

while [ $# -gt 0 ]; do
    case "$1" in
        --version)  
            if [[ $2 != --* ]]; then
                teleporter_version=$2
            else 
                echo "Invalid Teleporter version $2" && printHelp && exit 1
            fi 
            shift;;
        --rpc-url)  
            if [[ $2 != --* ]]; then
                rpc_url=$2
            else 
                echo "Invalid rpc url $2" && printHelp && exit 1
            fi 
            shift;;
        --private-key) 
            if [[ $2 != --* ]]; then
                user_private_key=$2
            else 
                echo "Invalid private key $2" && printHelp && exit 1
            fi 
            shift;;
        --help) 
            printHelp && exit 0 ;;
        *) 
            echo "Invalid option: $1" && printHelp && exit 1;;
    esac
    shift
done

if [[ $teleporter_version == "" || $rpc_url == "" ]]; then
    echo "Invalid usage. Teleporter version and RPC URL required."
    printHelp && exit 1
fi

# Tokens required to deploy the contract.
# Equal to contractCreationGasLimit * contractCreationGasPrice
# from utils/deployment-utils/deployment_utils.go
gas_tokens_required=10000000000000000000 # 10^19 wei = 10 eth

# Download the artifacts for this release.
teleporter_contract_address=$(curl -sL https://github.com/ava-labs/icm-contracts/releases/download/$teleporter_version/TeleporterMessenger_Contract_Address_$teleporter_version.txt)
echo "TeleporterMessenger $teleporter_version contract address: $teleporter_contract_address"
teleporter_deployer_address=$(curl -sL https://github.com/ava-labs/icm-contracts/releases/download/$teleporter_version/TeleporterMessenger_Deployer_Address_$teleporter_version.txt)
echo "TeleporterMessenger $teleporter_version deployer address: $teleporter_deployer_address"
teleporter_deploy_tx=$(curl -sL https://github.com/ava-labs/icm-contracts/releases/download/$teleporter_version/TeleporterMessenger_Deployment_Transaction_$teleporter_version.txt)
teleporter_messenger_bytecode=$(curl -sL https://github.com/ava-labs/icm-contracts/releases/download/$teleporter_version/TeleporterMessenger_Bytecode_$teleporter_version.txt)
if [ "$teleporter_contract_address" == "Not Found" ]; then
    echo "Error: TeleporterMessenger $teleporter_version contract address not found."
    exit 1
fi

# Check if this TeleporterMessenger version has already been deployed on this chain.
teleporter_contract_code=$(cast codesize $teleporter_contract_address --rpc-url $rpc_url)
if [[ $teleporter_contract_code != "0" ]]; then
    echo "TeleporterMessenger $teleporter_version has already been deployed on this chain." && exit 0
fi

# Estimate the amount of gas required to deploy the TeleporterMessenger bytecode from the Teleporter
# deployer address in order to simulate the transaction. This will error if the TeleporterMessenger
# contract is unable to be deployed from the deployer address.
cast estimate --rpc-url $rpc_url \
    --from $teleporter_deployer_address \
    --create $teleporter_messenger_bytecode > /dev/null

# Check the current balance of the deployer address.
deployer_balance=$(cast balance --rpc-url $rpc_url $teleporter_deployer_address)

if [[ $(echo "$deployer_balance>=$gas_tokens_required" | bc) == 1 ]]; then
    echo "Deployer address already funded"
else
    # Calculate how many wei the deployer address needs to create the contract.
    transfer_amount=$(echo "$gas_tokens_required-$deployer_balance" | bc)
    if [[ $user_private_key == "" ]]; then
        echo "No private key provided. Deployer address must be funded with $transfer_amount wei to deploy contract" && exit 1
    fi
    echo "Funding deployer address with $transfer_amount wei"
    cast send --rpc-url $rpc_url --private-key $user_private_key --value $transfer_amount $teleporter_deployer_address
fi

# Deploy the TeleporterMessenger contract by publishing the raw Nick's method transaction.
echo "Deploying TeleporterMessenger $teleporter_version"
deployment_result=$(cast publish --rpc-url $rpc_url $teleporter_deploy_tx)
deployment_status=$(echo $deployment_result | jq -r .status)
deployment_tx_id=$(echo $deployment_result | jq -r .transactionHash)
if [[ $deployment_status != "0x1" ]]; then 
    echo "TeleporterMessenger deployment transaction failed. Transaction ID: $deployment_tx_id"
    echo "Check failure reason and, if necessary, investigate deployment through state upgrade."
    echo "See https://github.com/ava-labs/subnet-evm/tree/master/stateupgrade"
    exit 1
fi

echo "Success! TeleporterMessenger $teleporter_version deployed to $teleporter_contract_address in transaction $deployment_tx_id."

exit 0
