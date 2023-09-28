#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

cwd=$PWD
cd $SUBNET_EVM_PATH
BASEDIR=$DATA_DIRECTORY AVALANCHEGO_BUILD_PATH=$DATA_DIRECTORY/avalanchego ./scripts/install_avalanchego_release.sh
./scripts/build.sh $DATA_DIRECTORY/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy

cd $cwd
AVALANCHEGO_BUILD_PATH=$DATA_DIRECTORY/avalanchego DATA_DIR=$DATA_DIRECTORY/data ./scripts/local/e2e_test.sh