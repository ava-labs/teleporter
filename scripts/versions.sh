#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e
set -o pipefail

# The go version for this project is set from a combination of major.minor from go.mod and the patch version set here.
GO_PATCH_VERSION=10

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

# Pass in the full name of the dependency.
# Parses go.mod for a matching entry and extracts the version number.
function getDepVersion() {
    grep -m1 "^\s*$1" $TELEPORTER_PATH/go.mod | cut -d ' ' -f2
}

# AWM_RELAYER is needed for the bash-script integration tests, which should eventually be removed.
AWM_RELAYER_VERSION=${AWM_RELAYER_VERSION:-'v0.2.3'}

# This needs to be exported to be picked up by the dockerfile.
GO_VERSION=${GO_VERSION:-$(getDepVersion go).$GO_PATCH_VERSION}

# Don't export them as they're used in the context of other calls
AVALANCHE_VERSION=${AVALANCHE_VERSION:-$(getDepVersion github.com/ava-labs/avalanchego)}
GINKGO_VERSION=${GINKGO_VERSION:-$(getDepVersion github.com/onsi/ginkgo/v2)}
