#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# use ggrep on arm64 otherwise grep -P returns error.
# set ARCH before calling this function
function setGrep() {
    if [ "$ARCH" = 'arm64' ]; then
        export grepcmd="ggrep"
    else
        export grepcmd="grep"
    fi
}

function convertToLower() {
    if [ "$ARCH" = 'arm64' ]; then
        echo $1 | perl -ne 'print lc'
    else
        echo $1 | sed -e 's/\(.*\)/\L\1/'
    fi
}
