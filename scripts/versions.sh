#!/usr/bin/env bash
# Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e
set -o pipefail

TELEPORTER_TOKEN_BRIDGE_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $TELEPORTER_TOKEN_BRIDGE_PATH/scripts/constants.sh
# Use the versions supplied by the teleporter submodule
source $TELEPORTER_PATH/scripts/versions.sh
