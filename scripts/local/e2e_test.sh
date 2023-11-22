#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd ../.. && pwd
)

source "$TELEPORTER_PATH"/scripts/constants.sh
source "$TELEPORTER_PATH"/scripts/versions.sh

BASEDIR=${BASEDIR:-"/tmp/e2e-test"}

cwd=$(pwd)
# Build the teleporter and cross chain apps smart contracts
BASEDIR=$BASEDIR AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego ./scripts/install_avalanchego_release.sh
BASEDIR=$BASEDIR ./scripts/install_subnetevm_release.sh
echo "Copy ${BASEDIR}/subnet-evm/subnet-evm to ${BASEDIR}/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy"
cp ${BASEDIR}/subnet-evm/subnet-evm ${BASEDIR}/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy

export AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego
export DATA_DIR=$BASEDIR/data

cd $TELEPORTER_PATH/contracts
if command -v forge &> /dev/null; then
  forge build
else
  echo "Forge command not found, attempting to use from $HOME"
  $HOME/.foundry/bin/forge build
fi

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