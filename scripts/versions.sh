#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e
set -o pipefail

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

function setGO_VERSION() {
    export GO_VERSION=$(getGO_VERSION)
}

function getGO_VERSION() {
    echo $(grep -m1 go $TELEPORTER_PATH/go.mod | cut -d ' ' -f2).$(grep -m1 GO_PATCH_VERSION $TELEPORTER_PATH/go.mod | cut -d ' ' -f3)
}

# Pass in the full name of the dependency
function getDepVersion() {
    grep -m1 $1 $TELEPORTER_PATH/go.mod | cut -d ' ' -f2
}

# Don't export them as they're used in the context of other calls
AVALANCHE_VERSION=${AVALANCHE_VERSION:-$(getDepVersion github.com/ava-labs/avalanchego)}
GINKGO_VERSION=${GINKGO_VERSION:-$(getDepVersion github.com/onsi/ginkgo/v2)}
