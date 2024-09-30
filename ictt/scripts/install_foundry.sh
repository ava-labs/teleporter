#!/usr/bin/env bash
# Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e
set -o pipefail

AVALANCHE_INTERCHAIN_TOKEN_TRANSFER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $AVALANCHE_INTERCHAIN_TOKEN_TRANSFER_PATH/scripts/constants.sh
$TELEPORTER_PATH/scripts/install_foundry.sh
