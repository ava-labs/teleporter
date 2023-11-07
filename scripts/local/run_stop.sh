#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error
source ./scripts/utils.sh

if [[ $# -gt 1 ]]; then
    echo "Invalid number of arguments. Usage:"
    echo "   ./scripts/local/run_stop.sh             # stop the running containers and preserve the network for subsequent runs"
    echo "   ./scripts/local/run_stop.sh -c          # stop the running containers and clean the network"
    exit 1
fi

clean=false
while getopts c flag
do
    case "${flag}" in
        c ) clean=true;;
    esac
done

# If clean flag is set we remove network running file to start the network from scratch
if $clean
then
    echo "Cleaning network"
    rm -f NETWORK_RUNNING
fi

docker compose -f docker/docker-compose-run.yml --project-directory ./ stop