#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

FOUNDRY_VERSION=v0.1.0
curl -L https://raw.githubusercontent.com/ava-labs/foundry/${FOUNDRY_VERSION}/foundryup/install > /tmp/foundry-install-script
sed -i "s/\/ava-labs\/foundry\/master\/foundryup/\/ava-labs\/foundry\/${FOUNDRY_VERSION}\/foundryup/g" /tmp/foundry-install-script
cat /tmp/foundry-install-script | bash
echo "export PATH=\"$PATH:/$HOME/.foundry/bin\"">> ~/.bashrc
source ~/.bashrc
export PATH=$PATH:$HOME/.foundry/bin:$HOME/.cargo/bin
foundryup --version ${FOUNDRY_VERSION}