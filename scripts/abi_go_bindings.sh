#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

source ./scripts/utils.sh

setARCH

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

DEFAULT_CONTRACT_LIST="TeleporterMessenger ERC20Bridge ExampleCrossChainMessenger BlockHashPublisher BlockHashReceiver BridgeToken"

CONTRACT_LIST=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -c | --contract) CONTRACT_LIST=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/abi_go_bindings.sh [OPTIONS]"
    echo "Build contracts and generate Go bindings"
    echo ""
    echo "Options:"
    echo "  -c, --contract <contract_name>          Generate Go bindings for the contract. If empty, generate Go bindings for a default list of contracts"
    echo "  -c, --contract "contract1 contract2"    Generate Go bindings for multiple contracts"
    echo "  -h, --help                              Print this help message"
    exit 0
fi

if ! command -v forge &> /dev/null; then
    echo "forge not found, installing"
    curl -L https://foundry.paradigm.xyz | bash
    source $HOME/.bashrc
    foundryup
fi

echo "Building subnet-evm abigen"
go install $TELEPORTER_PATH/subnet-evm/cmd/abigen

echo "Building Contracts"
cd $TELEPORTER_PATH/contracts
forge build --extra-output-files abi

contract_names=($CONTRACT_LIST)

# If CONTRACT_LIST is empty, use DEFAULT_CONTRACT_LIST
if [[ -z "${CONTRACT_LIST}" ]]; then
    contract_names=($DEFAULT_CONTRACT_LIST)
fi

cd $TELEPORTER_PATH
for contract_name in "${contract_names[@]}"
do
    abi_file=$TELEPORTER_PATH/contracts/out/$contract_name.sol/$contract_name.abi.json
    if ! [ -f $abi_file ]; then
        echo "Error: Contract $contract_name abi file not found"
        exit 1
    fi

    echo "Generating Go bindings for $contract_name..."
    mkdir -p $TELEPORTER_PATH/abis-bings
    $GOPATH/bin/abigen --abi $abi_file \
                       --pkg abisbings \
                       --type $contract_name \
                       --out $TELEPORTER_PATH/abis-bings/$contract_name.go
    echo "Done generating Go bindings for $contract_name."
done

exit 0