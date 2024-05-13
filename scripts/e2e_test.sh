#!/usr/bin/env bash
# Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_TOKEN_BRIDGE_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source "$TELEPORTER_TOKEN_BRIDGE_PATH"/scripts/constants.sh
source "$TELEPORTER_TOKEN_BRIDGE_PATH"/scripts/versions.sh

BASEDIR=${BASEDIR:-"$HOME/.teleporter-token-bridge-deps"}

cwd=$(pwd)
# Install the avalanchego and subnet-evm binaries
rm -rf $BASEDIR/avalanchego
BASEDIR=$BASEDIR AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego $TELEPORTER_PATH/scripts/install_avalanchego_release.sh
BASEDIR=$BASEDIR $TELEPORTER_PATH/scripts/install_subnetevm_release.sh

cp ${BASEDIR}/subnet-evm/subnet-evm ${BASEDIR}/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy
echo "Copied ${BASEDIR}/subnet-evm/subnet-evm binary to ${BASEDIR}/avalanchego/plugins/"

export AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego

if command -v forge &> /dev/null; then
  FORGE_COMMAND="forge build"
else
  echo "Forge command not found, attempting to use from $HOME"
  FORGE_COMMAND="$HOME/.foundry/bin/forge build"
fi

cd $TELEPORTER_TOKEN_BRIDGE_PATH/contracts
$FORGE_COMMAND

cd $TELEPORTER_PATH/contracts
$FORGE_COMMAND

cd $cwd
# Build ginkgo
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo

ginkgo build ./tests/local/

# Run the tests
echo "Running e2e tests"
RUN_E2E=true ./tests/local/local.test \
  --ginkgo.vv \
  --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
  --ginkgo.focus=${GINKGO_FOCUS:-""} \
  --ginkgo.trace

echo "e2e tests passed"
exit 0
