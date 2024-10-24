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
    --components component1,component2            Comma separated list of test suites to run. Valid components are:
                                                  $(echo $valid_components | tr ' ' '\n' | sort | tr '\n' ' ')
                                                  (default: all)
Options:
    --help                                        Print this help message
EOF
}

valid_components=$(ls -d $TELEPORTER_PATH/tests/local/*/ | xargs -n 1 basename)
components=

while [ $# -gt 0 ]; do
    case "$1" in
        --components)  
            if [[ $2 != --* ]]; then
                components=$2
            else 
                echo "Invalid components $2" && printHelp && exit 1
            fi 
            shift;;
        --help) 
            printHelp && exit 0 ;;
        *) 
            echo "Invalid option: $1" && printHelp && exit 1;;
    esac
    shift
done

# Run all suites if no component is provided
if [ -z "$components" ]; then
    components=$valid_components
fi

# Exit if invalid component is provided
for component in $(echo $components | tr ',' ' '); do
    if [[ $valid_components != *$component* ]]; then
        echo "Invalid component $component" && exit 1
    fi
done

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

for component in $(echo $components | tr ',' ' '); do
    echo "Building e2e tests for $component"
    ginkgo build ./tests/local/$component

    echo "Running e2e tests for $component"

    # If component is validator-manager, run each flow separately
    if [ "$component" == "validator-manager" ]; then
        for flow in "ERC20" "Native" "PoA"; do
            echo "Running $flow flow"
            RUN_E2E=true ./tests/local/$component/$component.test \
            --ginkgo.vv \
            --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
            --ginkgo.trace \
            --ginkgo.focus=$flow
        done
        continue
    fi

    RUN_E2E=true ./tests/local/$component/$component.test \
    --ginkgo.vv \
    --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
    --ginkgo.focus=${GINKGO_FOCUS:-""} \
    --ginkgo.trace

    echo "$component e2e tests passed"
    echo ""
done
exit 0
