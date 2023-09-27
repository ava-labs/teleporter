#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

BASE_DIR=${XDG_CONFIG_HOME:-$HOME}
if ! command -v forge &> /dev/null; then
    echo "forge not found, installing"
    curl -L https://foundry.paradigm.xyz | bash
    source $HOME/.bashrc
    $BASE_DIR/.foundry/bin/foundryup
fi

echo "Building subnet-evm abigen"
go install $TELEPORTER_PATH/subnet-evm/cmd/abigen

echo "Building Contracts"
cd $TELEPORTER_PATH/contracts
$BASE_DIR/.foundry/bin/forge build --extra-output-files abi
cp $TELEPORTER_PATH/contracts/out/TeleporterMessenger.sol/TeleporterMessenger.abi.json $TELEPORTER_PATH/abis/

echo "Generating TeleporterMessenger Go bindings"
gen_file_path=$TELEPORTER_PATH/abis/go/teleport-messenger/teleport_messenger.go
cd $TELEPORTER_PATH
mkdir -p abis/go/teleport-messenger
$GOPATH/bin/abigen --abi abis/TeleporterMessenger.abi.json --pkg teleport_messenger --out $gen_file_path

echo "TeleporterMessenger Go bindings generated at $gen_file_path"

exit 0