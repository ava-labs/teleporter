#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

CHECK=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -c | --check) CHECK=true ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/lint.sh [OPTIONS]"
    echo "Format Solidity contracts."
    echo ""
    echo "Options:"
    echo "  -c, --check                       Check for proper linted files. Exits with code 1 if not."
    echo "  -h, --help                        Print this help message"
    exit 0
fi

TELEPORTER_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

if ! command -v forge &> /dev/null; then
    echo "forge not found, installing"
    $TELEPORTER_PATH/scripts/install_foundry.sh
fi

if [ "$CHECK" = true ] ; then
    forge fmt --check --root $TELEPORTER_PATH/contracts $TELEPORTER_PATH/contracts/src/**
else
    forge fmt --root $TELEPORTER_PATH/contracts $TELEPORTER_PATH/contracts/src/**
fi

exit 0