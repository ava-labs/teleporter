#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Set up the versions to be used
awm_relayer_version=${AWM_RELAYER_VERSION:-'v0.2.1'}
subnet_evm_version=${SUBNET_EVM_VERSION:-'v0.5.4'}
# Don't export them as they're used in the context of other calls
avalanche_version=${AVALANCHE_VERSION:-'v1.10.9'}
GINKGO_VERSION=${GINKGO_VERSION:-'v2.2.0'}
