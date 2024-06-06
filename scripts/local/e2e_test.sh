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

BASEDIR=${BASEDIR:-"$HOME/.teleporter-deps"}

cwd=$(pwd)

AVALANCHE_CLI_DIR=$HOME/.avalanche-cli

export AVALANCHEGO_BUILD_PATH=${AVALANCHE_CLI_DIR}/bin/avalanchego/avalanchego-${AVALANCHEGO_VERSION}

# ensure that avalanche-cli has installed the version of avalanchego that we need:
if [[ ! -x "$AVALANCHEGO_BUILD_PATH/avalanchego" ]]; then
  avalanche network clean
  avalanche network start --avalanchego-version "$AVALANCHEGO_VERSION"
  avalanche network clean
fi

subnetevm_bin=${AVALANCHE_CLI_DIR}/bin/subnet-evm/subnet-evm-${SUBNET_EVM_VERSION}/subnet-evm

# ensure that avalanche-cli has installed the version of subnetevm that we need:
if [[ ! -x "$subnetevm_bin" ]]; then
  avalanche subnet create dummy --evm --vm-version "$SUBNET_EVM_VERSION" \
    --force \
    --teleporter --relayer --evm-chain-id 1 --evm-token DUM --evm-defaults # this line is just to avoid the dialog prompts
fi

# Install the subnet-evm binary as a plugin within the avalanche-cli resources.
# This is a hack that can probably be undone when we stop setting up the
# avalanche network and subnet ourselves (via BeforeSuite/AfterSuite in
# tests/local/e2e_test.go) and start having avalanche-cli do that for us
# instead.
mkdir -p "${AVALANCHEGO_BUILD_PATH}/plugins"
cp "$subnetevm_bin" \
  "${AVALANCHEGO_BUILD_PATH}/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy"

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

ginkgo build ./tests/local/

# Run the tests
echo "Running e2e tests $RUN_E2E"
RUN_E2E=true ./tests/local/local.test \
  --ginkgo.vv \
  --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
  --ginkgo.focus=${GINKGO_FOCUS:-""} \
  --ginkgo.trace

echo "e2e tests passed"
exit 0
