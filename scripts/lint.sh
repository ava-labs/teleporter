#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

cd $TELEPORTER_PATH/contracts/src
solhint '**/*.sol' --config ./.solhint.json --ignore-path ./.solhintignore
exit 0