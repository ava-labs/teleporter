#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

source ./scripts/versions.sh

LINTER=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -l | --linter) LINTER=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/lint.sh [OPTIONS]"
    echo "Lint Teleporter Solidity contracts and E2E tests Golang code."
    echo ""
    echo "Options:"
    echo "  -l, --linter                      Specify which linter (sol/go) to run. If empty, runs both solidity and go linters."
    echo "  -h, --help                        Print this help message"
    exit 0
fi

function sollinter() {
    # lint solidity contracts
    echo "Linting Teleporter Solidity contracts..."
    cd $TELEPORTER_PATH/contracts/src
    # "solhint **/*.sol" runs differently than "solhint '**/*.sol'", where the latter checks sol files
    # in subdirectories. The former only checks sol files in the current directory and directories one level down.
    solhint '**/*.sol' --config ./.solhint.json --ignore-path ./.solhintignore --max-warnings 0
}

function golanglinter() {
    # lint e2e tests go code
    go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

    echo "Linting Teleporter E2E tests Golang code..."
    cd $TELEPORTER_PATH/tests
    golangci-lint run --config=$TELEPORTER_PATH/.golangci.yml ./...
}

if [ -z "$LINTER" ]; then
    sollinter
    golanglinter
elif [ "$LINTER" = "sol" ]; then
    sollinter
elif [ "$LINTER" = "go" ]; then
    golanglinter
else
    echo "Invalid linter option: $LINTER"
    exit 1
fi


exit 0