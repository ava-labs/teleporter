#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd ../.. && pwd
)

function printHelp() {
    echo "Usage: ./scripts/local/e2e_test.sh [--component component]"
    echo ""
    printUsage
}

function printUsage() {
    cat << EOF
Arguments:
    --component <component>          Component test suite to run. Must be a valid directory in tests/local/
Options:
    --help                           Print this help message
EOF
}

component=

while [ $# -gt 0 ]; do
    case "$1" in
        --component)  
            if [[ $2 != --* ]]; then
                component=$2
            else 
                echo "Invalid component $2" && printHelp && exit 1
            fi 
            shift;;
        --help) 
            printHelp && exit 0 ;;
        *) 
            echo "Invalid option: $1" && printHelp && exit 1;;
    esac
    shift
done

# Exit if no component is provided
if [ -z "$component" ]; then
    echo "No component provided" && exit 1
fi

if [ ! -d "./tests/local/$component" ]; then
    echo "Component test suite not found" && exit 1
fi

source "$TELEPORTER_PATH"/scripts/constants.sh
source "$TELEPORTER_PATH"/scripts/versions.sh

BASEDIR=${BASEDIR:-"$HOME/.teleporter-deps"}

cwd=$(pwd)
# Install the avalanchego and subnet-evm binaries
rm -rf $BASEDIR/avalanchego
BASEDIR=$BASEDIR AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego "${TELEPORTER_PATH}/scripts/install_avalanchego_release.sh"
BASEDIR=$BASEDIR "${TELEPORTER_PATH}/scripts/install_subnetevm_release.sh"

cp ${BASEDIR}/subnet-evm/subnet-evm ${BASEDIR}/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy
echo "Copied ${BASEDIR}/subnet-evm/subnet-evm binary to ${BASEDIR}/avalanchego/plugins/"

export AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego

cd $TELEPORTER_PATH
if command -v forge &> /dev/null; then
  forge build --skip test 
else
  echo "Forge command not found, attempting to use from $HOME"
  $HOME/.foundry/bin/forge build
fi

cd "$TELEPORTER_PATH"
# Build ginkgo
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo@${GINKGO_VERSION}

ginkgo build ./tests/local/$component

# Run the tests
echo "Running e2e tests $RUN_E2E"
RUN_E2E=true ./tests/local/$component/$component.test \
  --ginkgo.vv \
  --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
  --ginkgo.focus=${GINKGO_FOCUS:-""} \
  --ginkgo.trace

echo "e2e tests passed"
exit 0
