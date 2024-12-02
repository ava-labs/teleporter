#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

ICM_CONTRACTS_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $ICM_CONTRACTS_PATH/scripts/versions.sh

function solFormat() {
    # format solidity contracts
    echo "Formatting Solidity contracts..."
    forge fmt --root $ICM_CONTRACTS_PATH $ICM_CONTRACTS_PATH/contracts/**
}

function solFormatCheck() {
    # format solidity contracts
    echo "Checking formatting of Solidity contracts..."
    forge fmt --check --root $ICM_CONTRACTS_PATH $ICM_CONTRACTS_PATH/contracts/**
}

function solLinter() {
    # lint solidity contracts
    echo "Linting Solidity contracts..."
    cd $ICM_CONTRACTS_PATH
    # "solhint **/*.sol" runs differently than "solhint '**/*.sol'", where the latter checks sol files
    # in subdirectories. The former only checks sol files in the current directory and directories one level down.
    solhint '**/*.sol' --config ./.solhint.json --ignore-path ./.solhintignore --max-warnings 0
}

function golangLinter() {
    # lint e2e tests go code
    go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

    echo "Linting Golang code..."
    cd $ICM_CONTRACTS_PATH
    golangci-lint run --config=$ICM_CONTRACTS_PATH/.golangci.yml ./...
}

function runAll() {
    solFormat
    solLinter
    golangLinter
}

function printHelp() {
    echo "Usage: ./scripts/lint.sh [OPTIONS]"
    echo "Lint/Format Teleporter Solidity contracts and E2E tests Golang code."
    echo "Pass no parameters to perform all checks"
    printUsage
}

function printUsage() {
    echo "Options:"
    echo "  -sfc, --sol-format-check            Check for proper formatted Solidity files. Exits with code 1 if not."
    echo "  -sf,  --sol-format                  Format Solidity contracts"
    echo "  -sl,  --sol-lint                    Run the Solidity linter"
    echo "  -gl,  --go-lint                     Run the Golang linter"
    echo "  -h,   --help                        Print this help message"
}

# if we have no args, perform all checks
if [ $# -eq 0 ]; then
    runAll
    exit 0
fi

while [ $# -gt 0 ]; do
    case "$1" in
        -sfc | --sol-format-check) 
            solFormatCheck ;;
        -sf  | --sol-format) 
            solFormat ;;
        -sl  | --sol-lint) 
            solLinter ;;
        -gl  | --go-lint) 
            golangLinter ;;
        -h   | --help) 
            printHelp ;;
        *) 
          echo "Invalid option: $1" && printHelp && exit 1;;
    esac
    shift
done


exit 0
