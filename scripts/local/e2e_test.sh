#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd ../.. && pwd
)

source "$TELEPORTER_PATH"/scripts/local/constants.sh

source "$TELEPORTER_PATH"/scripts/local/versions.sh

# Build the teleporter and cross chain apps smart contracts
cwd=$(pwd)
cd $TELEPORTER_PATH/contracts
forge build
cd $cwd

# Build ginkgo
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo@${GINKGO_VERSION}

ginkgo build ./tests/

# Run the tests
echo "Running e2e tests $RUN_E2E"
RUN_E2E=true ./tests/tests.test \
  --ginkgo.vv \
  --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""}

echo "e2e tests passed"
exit 0