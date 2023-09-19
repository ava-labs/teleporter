#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides

TELEPORTER_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd ../.. && pwd
)

# Set the PATHS
GOPATH="$(go env GOPATH)"

# Current branch
current_branch=$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match || true)

git_commit=${RELAYER_COMMIT:-$( git rev-list -1 HEAD )}
