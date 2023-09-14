#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error
source ./scripts/utils.sh

# Set ARCH env so as a container executes without issues in a portable way
# Should be amd64 for linux/macos x86 hosts, and arm64 for macos M1
# It is referenced in the docker composer yaml, and then passed as a Dockerfile ARG
setARCH

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
    rm -f NETWORK_RUNNING
fi

docker compose -f scripts/local/native/docker-compose-run.yml --project-directory ./ stop