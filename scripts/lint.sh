#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

cd $TELEPORTER_PATH/contracts/src
# "solhint **/*.sol" runs differently than "solhint '**/*.sol'", where the latter checks sol files
# in subdirectories. The former only checks sol files in the current directory and directories one level down.
solhint '**/*.sol' --config ./.solhint.json --ignore-path ./.solhintignore --max-warnings 0

exit 0