#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -ex

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

if ! command -v forge &> /dev/null; then
    echo "forge not found, installing"
    $TELEPORTER_PATH/scripts/install_foundry.sh
fi

forge fmt $TELEPORTER_PATH/contracts/src/**

exit 0