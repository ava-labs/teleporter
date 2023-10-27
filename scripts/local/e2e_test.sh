#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

SUBNET_EVM_PATH=
LOCAL=
DATA_DIRECTORY=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -l | --local) LOCAL=true ;;
        -s | --subnet-evm) SUBNET_EVM_PATH=$2 ;;
        -d | --data-dir) DATA_DIRECTORY=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/local/e2e_test.sh [OPTIONS]"
    echo "Run E2E tests for Teleporter."
    echo ""
    echo "Options:"
    echo "  -l, --local                       Run the test locally. Requires --subnet-evm and --data-dir"
    echo "  -s, --subnet-evm <path>           Path to subnet-evm repo"
    echo "  -d, --data-dir <path>             Path to data directory"
    echo "  -h, --help                        Print this help message"
    exit 0
fi

if [ "$LOCAL" = true ]; then
    if [ -z "$DATA_DIRECTORY" ]; then
        echo "Must specify data directory when running local"
        exit 1
    fi
    if [ -z "$SUBNET_EVM_PATH" ]; then
        echo "Must specify subnet-evm path when running local"
        exit 1
    fi
    cwd=$PWD
    cd $SUBNET_EVM_PATH
    BASEDIR=$DATA_DIRECTORY AVALANCHEGO_BUILD_PATH=$DATA_DIRECTORY/avalanchego ./scripts/install_avalanchego_release.sh
    ./scripts/build.sh $DATA_DIRECTORY/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy

    cd $cwd
    export AVALANCHEGO_BUILD_PATH=$DATA_DIRECTORY/avalanchego
    export DATA_DIR=$DATA_DIRECTORY/data
fi

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd ../.. && pwd
)

source "$TELEPORTER_PATH"/scripts/constants.sh

source "$TELEPORTER_PATH"/scripts/versions.sh

# Build the teleporter and cross chain apps smart contracts
cwd=$(pwd)
cd $TELEPORTER_PATH/contracts
if [[ ":$PATH:" == *".foundry/bin"* ]]; then
  forge build
else
  echo "Foundry not found in PATH, attempting to use from HOME"
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