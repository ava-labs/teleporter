#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

if ! command -v forge &> /dev/null; then
    echo "forge not found, installing"
    $TELEPORTER_PATH/scripts/install_foundry.sh
fi

function printHelp() {
    echo "Usage: ./scripts/deploy_teleporter.sh [OPTIONS]"
    echo "Deploys a selected TeleporterMessenger contract to the specified chain"
    echo "For a list of releases, go to https://github.com/ava-labs/teleporter/releases"
    printUsage
}

function printUsage() {
    echo "Options:"
    echo "  --fund-deployer <private_key>    Optional. Funds the deployer address with the account held by <private_key>"
    echo "  --version <version>              Required. Specify the release version to deploy"
    echo "  --rpc-url <url>                  Required. Specify the rpc url of the node to use"
    echo "  --help                           Print this help message"
}

teleporter_version=
user_private_key=
rpc_url=

while [ $# -gt 0 ]; do
    case "$1" in
        --fund-deployer) 
            if [[ $2 != --* ]]; then
                user_private_key=$2
            else 
                echo "Invalid private key $2" && printHelp && exit 1
            fi 
            shift;;
        --version)  
            if [[ $2 != --* ]]; then
                teleporter_version=$2
            else 
                echo "Invalid teleporter version $2" && printHelp && exit 1
            fi 
            shift;;
        --rpc-url)  
            if [[ $2 != --* ]]; then
                rpc_url=$2
            else 
                echo "Invalid rpc url $2" && printHelp && exit 1
            fi 
            shift;;
        --help) 
            printHelp && exit 0 ;;
        *) 
            echo "Invalid option: -$1" && printHelp && exit 1;;
    esac
    shift
done

# Tokens required to deploy the contract.
# Equal to contractCreationGasLimit * contractCreationGasPrice
# from utils/deployment-utils/deployment_utils.go
gas_tokens_required=10000000000000000000 # 10^19 wei = 10 eth

teleporter_contract_address=$(curl -sL https://github.com/ava-labs/teleporter/releases/download/$teleporter_version/TeleporterMessenger_ContractAddress_$teleporter_version.txt)
echo "TeleporterMessenger $teleporter_version contract address: $teleporter_contract_address"
teleporter_deployer_address=$(curl -sL https://github.com/ava-labs/teleporter/releases/download/$teleporter_version/TeleporterMessenger_DeployerAddress_$teleporter_version.txt)
echo "TeleporterMessenger $teleporter_version deployer address: $teleporter_deployer_address"
teleporter_deploy_tx=$(curl -sL https://github.com/ava-labs/teleporter/releases/download/$teleporter_version/TeleporterMessenger_DeployerTransaction_$teleporter_version.txt)

if [[ $(cast code --rpc-url $rpc_url $teleporter_contract_address) != "0x" ]]; then
    echo "TeleporterMessenger $teleporter_version has already been deployed on this chain." && exit 1
fi

deployer_balance=$(cast balance --rpc-url $rpc_url $teleporter_deployer_address)

if [[ $(echo "$deployer_balance>=$gas_tokens_required" | bc) == 1 ]]; then
    echo "Deployer address already funded"
else 
    transfer_amount=$(echo "$gas_tokens_required-$deployer_balance" | bc)
    if [[ $user_private_key == "" ]]; then
        echo "No private key provided. Deployer address must be funded with $transfer_amount wei to deploy contract" && exit 1
    fi
    echo "Funding deployer address with $transfer_amount wei"
    cast send --rpc-url $rpc_url --private-key $user_private_key --value $transfer_amount $teleporter_deployer_address
fi

echo "Deploying TeleporterMessenger $teleporter_version"
cast publish --rpc-url $rpc_url $teleporter_deploy_tx

echo "Success! TeleporterMessenger $teleporter_version deployed to $teleporter_deployer_address"
exit 0
