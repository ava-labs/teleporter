#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

function getBlockChainIDHex() {
    python3 -c "import base58,sys; sys.stdout.write(base58.b58decode(b'$1').hex()[:-8])";
}

function parseContractAddress() {
    echo $1 | $(getGrep) -o -P 'Deployed to: .{42}' | sed 's/^.\{13\}//';
}

# use ggrep on arm64 otherwise grep -P returns error.
function getGrep() {
    if [ $(getARCH) = 'arm64' ]; then
        echo ggrep
    else
        echo grep
    fi
}

# Get ARCH so a container can execute without issues in a portable way
# Should be amd64 for linux/macos x86 hosts, and arm64 for macos M1
function getARCH() {
    ARCH=$(uname -m)
    [ $ARCH = 'x86_64' ] && ARCH=amd64
    [ $ARCH = 'aarch64' ] && ARCH=arm64
    echo $ARCH
}

function convertToLower() {
    if [ $(getARCH) = 'arm64' ] ; then
        echo $1 | perl -ne 'print lc'
    else
        echo $1 | sed -e 's/\(.*\)/\L\1/'
    fi
}