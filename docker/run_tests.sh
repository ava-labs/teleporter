#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e # Stop on first error

# Needed for submodules
git config --global --add safe.directory '*'
if [[ $# -eq 1 ]] && [[ "$1" == "--local-tests-only" ]]; then
    dir_prefix=.
else
    dir_prefix=/code
fi

rm -f $dir_prefix/NETWORK_READY

until cat $dir_prefix/NETWORK_READY &> /dev/null
do
    if [[ retry_count -ge 60 ]]; then
        echo "Subnets didn't start up quickly enough."
        exit 1
    fi
    echo "Waiting for subnets to start up. Retry count: $retry_count"
    retry_count=$((retry_count+1))
    sleep 10
done

# Wait for relayers to start and subscribe
sleep 3

# Source all variables set in run_setup.sh
set -a
source $dir_prefix/vars.sh || true # ignore readonly variable errors
source $dir_prefix/scripts/utils.sh
set +a

# Run the selected unit tests.
tests_to_run=($TEST_TARGET)

# If TEST_TARGET is empty, run all tests in teleporter/integration-tests/
if [[ -z "${TEST_TARGET}" ]]; then
    for entry in "./scripts/local/integration-tests"/*
    do
        fname="$(basename $entry .sh)" # strip the path and file extension
        tests_to_run+=($fname)
    done

fi

for test_name in "${tests_to_run[@]}"
do
    echo "Running test $test_name..."
    ./scripts/local/integration-tests/$test_name.sh
    echo "Done running test $test_name."
    echo ""
done

exit 0
