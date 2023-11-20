#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Set up the versions to be used
AWM_RELAYER_VERSION=${AWM_RELAYER_VERSION:-'v0.2.4'}
SUBNET_EVM_VERSION=${SUBNET_EVM_VERSION:-'v0.5.9'}
# Don't export them as they're used in the context of other calls
AVALANCHE_VERSION=${AVALANCHE_VERSION:-'v1.10.15'}
GINKGO_VERSION=${GINKGO_VERSION:-'v2.13.1'}
