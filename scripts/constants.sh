#!/usr/bin/env bash
# Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides

TELEPORTER_TOKEN_BRIDGE_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

TELEPORTER_PATH="$TELEPORTER_TOKEN_BRIDGE_PATH"/contracts/lib/teleporter
source $TELEPORTER_PATH/scripts/constants.sh
