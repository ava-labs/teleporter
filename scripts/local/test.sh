#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error
source ./scripts/utils.sh

TEST_TARGET=
LOCAL_RELAYER_IMAGE=
TESTNET=
RUN_STOP_FLAG="-c"
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -t | --test) export TEST_TARGET=$2 ;;
        -l | --local-relayer-image) LOCAL_RELAYER_IMAGE=$2 ;;
        -p | --pause) RUN_STOP_FLAG= ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/local/test.sh [OPTIONS]"
    echo "Run integration tests for Teleporter."
    echo ""
    echo "Options:"
    echo "  -t, --test <test_name>            Run a specific test. If empty, runs all tests in the ./scripts/local/integration-tests/"
    echo "  -t, --test "test1 test2"          Run multiple tests. Test names must be space delimited and enclosed in quotes"
    echo "  -l, --local-relayer-image <tag>   Use a local AWM Relayer image instead of pulling from dockerhub"
    echo "  -p, --pause                       Pause the network on stop. Will attempt to restart the paused network on subsequent runs"
    echo "  -h, --help                        Print this help message"
    exit 0
fi

function cleanup {
    echo "Shutting down network before exiting..."
    ./scripts/local/run_stop.sh $RUN_STOP_FLAG
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
    docker compose -f docker/docker-compose-test.yml --project-directory ./ up --abort-on-container-exit --build &
    dockerPid=$!
else
    echo "Using local awm-relayer image: $LOCAL_RELAYER_IMAGE"
    if [[ "$(docker images -q awm-relayer:$LOCAL_RELAYER_IMAGE 2> /dev/null)" == "" ]]; then
        echo "awm-relayer:$LOCAL_RELAYER_IMAGE does not exist locally. Exiting."
        exit 1
    fi
    rm -f docker/docker-compose-test-local.yml
    sed "s/<TAG>/$LOCAL_RELAYER_IMAGE/g" docker/docker-compose-test-local-template.yml > docker/docker-compose-test-local.yml
    docker compose -f docker/docker-compose-test-local.yml --project-directory ./ up --abort-on-container-exit --build &
    dockerPid=$!
fi

# Wait for the containers to start up
until [ "$( docker container inspect --format '{{.State.Running}}' test_runner )" == "true" ]
do
    if ! ps -p $dockerPid > /dev/null
    then
        echo "Docker process is dead"
        exit 1
    fi

    echo "Waiting for containers to start..."
    sleep 1
done
echo "Containers started."

# Wait for the test runner to exit, then grab the exit code
until [ "$( docker container inspect --format '{{.State.Running}}' test_runner )" == "false" ]
do
    sleep 10
done

# Stop the containers gracefully
./scripts/local/run_stop.sh $RUN_STOP_FLAG
code=$(docker inspect --format='{{.State.ExitCode}}' test_runner)
exit $code
