#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# set -e

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

if ! command -v forge &> /dev/null; then
    echo "forge not found, installing"
    curl -L https://foundry.paradigm.xyz | bash &&
    source $HOME/.bashrc
    foundryup
fi

echo "Building Contracts"
cd $TELEPORTER_PATH/contracts
forge build
python3 -c "import json; json.dump(json.load(open('out/TeleporterMessenger.sol/TeleporterMessenger.json'))['abi'], open('$TELEPORTER_PATH/abis/TeleporterMessenger.json', 'w'), indent=4)"

echo "Generating TeleporterMessenger Go bindings"
gen_file_path=$TELEPORTER_PATH/abis/go/teleport-messenger/teleport_messenger.go
cd $TELEPORTER_PATH
mkdir -p abis/go/teleport-messenger
abigen --abi abis/TeleporterMessenger.json --pkg teleport_messenger --out $gen_file_path

echo "TeleporterMessenger Go bindings generated at $gen_file_path"

exit 0