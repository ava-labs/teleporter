#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error

source ./scripts/utils.sh

LOCAL_RELAYER_IMAGE=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -l | --local-relayer-image) LOCAL_RELAYER_IMAGE=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/local/run.sh [OPTIONS]"
    echo "Run Teleporter."
    echo ""
    echo "Options:"
    echo "  -l, --local-relayer-image <tag>   Use a local AWM Relayer image instead of pulling from dockerhub"
    echo "  -h, --help                        Print this help message"
    exit 0
fi

function cleanup {
    echo "Shutting down network before exiting..."
    ./scripts/local/run_stop.sh
    echo "Network stopped"
}

# Set up the trap to catch the SIGINT signal (CTRL+C)
# Note that the output of the cleanup function defined in run_setup.sh does not appear if CTRL+C is used to kill this script,
# but the function does in fact run as expected.
trap cleanup SIGTERM
trap cleanup SIGINT

# Set ARCH env so as a container executes without issues in a portable way
# Should be amd64 for linux/macos x86 hosts, and arm64 for macos M1
# It is referenced in the docker composer yaml, and then passed as a Dockerfile ARG
setARCH

if [ -z "$LOCAL_RELAYER_IMAGE" ]; then
    echo "Using published awm-relayer image"
    docker compose -f docker/docker-compose-run.yml --project-directory ./ up --build &
else
    echo "Using local awm-relayer image: $LOCAL_RELAYER_IMAGE"
    if [[ "$(docker images -q awm-relayer:$LOCAL_RELAYER_IMAGE 2> /dev/null)" == "" ]]; then
        echo "awm-relayer:$LOCAL_RELAYER_IMAGE does not exist locally. Exiting."
        exit 1
    fi
    rm -f docker/docker-compose-run-local.yml
    sed "s/<TAG>/$LOCAL_RELAYER_IMAGE/g" docker/docker-compose-run-local-template.yml > docker/docker-compose-run-local.yml
    docker compose -f docker/docker-compose-run-local.yml --project-directory ./ up --build &
fi

tail -f /dev/null
