#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd ../.. && pwd
)

set -a

source "$TELEPORTER_PATH"/scripts/versions.sh
source $TELEPORTER_PATH/.env
source $TELEPORTER_PATH/.env.testnet

go run tests/testnet/main/run_testnet_flows.go
