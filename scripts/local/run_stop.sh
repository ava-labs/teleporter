#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error
source ./scripts/utils.sh

if [[ $# -gt 0 ]]; then
    echo "Invalid number of arguments. Usage:"
    echo "   ./scripts/local/run_stop.sh             # stop the running containers and clean the network"
    exit 1
fi

echo "Cleaning network"
rm -f NETWORK_RUNNING
docker compose -f docker/docker-compose-run.yml --project-directory ./ stop
echo "Network cleaned"
