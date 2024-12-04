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
    echo "Usage: ./scripts/deploy_registry.sh --version <version> --rpc-url <url> --private-key <private_key> [OPTIONS]"
    echo ""
    echo "Deploys a selected TeleporterRegistry contract to the specified chain"
    echo "For a list of releases, go to https://github.com/ava-labs/icm-contracts/releases"
    printUsage
}

function printUsage() {
    cat << EOF
Arguments:
    --version <version>              Specify the release version to deploy.
    --rpc-url <url>                  Specify the rpc url of the node to use
    --private-key <private_key>      Private key of account to deploy TeleporterRegistry
Options:
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

if [[ $teleporter_version == "" || $rpc_url == "" || $user_private_key == "" ]]; then
    echo "Invalid usage. Missing required command line arguments."
    printHelp && exit 1
fi

teleporter_registry_bytecode=$(curl -sL https://github.com/ava-labs/icm-contracts/releases/download/$teleporter_version/TeleporterRegistry_Bytecode_$teleporter_version.txt)
teleporter_contract_address=$(curl -sL https://github.com/ava-labs/icm-contracts/releases/download/$teleporter_version/TeleporterMessenger_Contract_Address_$teleporter_version.txt)
if [ "$teleporter_registry_bytecode" == "Not Found" ]; then
    echo "Error: TeleporterRegistry $teleporter_version byte code not found."
    exit 1
fi

# Encode the constructor arguments
# TODO: Update to iterate through all release major versions once we have multiple Teleporter versions.
constructor_encoding=$(cast abi-encode "constructor((uint256,address)[])" "[(1, $teleporter_contract_address)]")

# remove the 0x prefix
constructor_encoding=${constructor_encoding:2}

# Deploy the TeleporterRegistry contract 
deployment_result=$(cast send --private-key $user_private_key --rpc-url $rpc_url --json --create $teleporter_registry_bytecode$constructor_encoding)
teleporter_registry_address=$(echo $deployment_result | jq -r .contractAddress)
deployment_status=$(echo $deployment_result | jq -r .status)
deployment_tx_id=$(echo $deployment_result | jq -r .transactionHash)
if [[ $deployment_status != "0x1" ]]; then 
    echo "TeleporterRegistry deployment transaction failed. Transaction ID: $deployment_tx_id"
    exit 1
fi
echo "Success! TeleporterRegistry deployed to $teleporter_registry_address in transaction $deployment_tx_id."

exit 0
