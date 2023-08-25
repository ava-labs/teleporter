#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

echo "Using testnet environment variables"
set -a
source .env.testnet # config of testnet
source .env # set the user private key and address
source ./scripts/utils.sh
set +a

set -e # Stop on first error

setARCH
setGrep

TEST_TARGET=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -t | --test) TEST_TARGET=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/fuji/test.sh [OPTIONS]"
    echo "Run tests for Teleporter on Fuji testnet."
    echo ""
    echo "Options:"
    echo "  -t, --test <test_name>            Run a specific test. If empty, runs all tests in the ./scripts/fuji/example-workflows/"
    echo "  -t, --test "test1 test2"          Run multiple tests. Test names must be space delimited and enclosed in quotes"
    echo "  -h, --help                        Print this help message"
    exit 0
fi

# before running tests, check if user private key and address are provided
if [ -z "$user_private_key" ]; then
    echo "User private key not provided. Please set in .env file."
    exit 1
fi

if [ -z "$user_address" ]; then
    echo "User address not provided. Please set in .env file."
    exit 1
fi

# Run the selected unit tests.
tests_to_run=($TEST_TARGET)

# If TEST_TARGET is empty, run all tests in ./scripts/fuji/example-workflows/
if [[ -z "${TEST_TARGET}" ]]; then
    for entry in "./scripts/fuji/example-workflows"/*
    do
        fname="$(basename $entry .sh)" # strip the path and file extension
        tests_to_run+=($fname)
    done

fi

for test_name in "${tests_to_run[@]}"
do
    echo "Running test $test_name..."
    ./scripts/fuji/example-workflows/$test_name.sh
    echo "Done running test $test_name."
    echo ""
done

exit 0

