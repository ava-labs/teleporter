# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

FROM ubuntu:20.04

ARG DEBIAN_FRONTEND=noninteractive

# defaults to amd64. needs to be arm64 for Macos M1
# configured using either docker --build-arg flag or docker composer yaml service.build.args
ARG ARCH=amd64

# defaults to grep running in a container.
ENV grepcmd=grep

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get clean

# Install dev dependencies
# python is used to parse JSON in run_tests.sh, and to find the proper log directory in run_setup.sh
RUN apt-get install -y wget curl git python3 make gcc build-essential

# Install Golang
RUN apt-get install -y bison
RUN wget https://go.dev/dl/go1.20.8.linux-${ARCH}.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.8.linux-${ARCH}.tar.gz
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH $HOME/go
ENV GOROOT /usr/local/go
RUN go version

# Set locale
RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/* \
    && localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

# Install Rust (needed to install specific commit of Foundry from source.)
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

# Install foundry from specific commit
SHELL ["/bin/bash", "-c"]
RUN curl -L https://raw.githubusercontent.com/foundry-rs/foundry/70d00222c2ef74fc484b8d98c5705d131ab31871/foundryup/install > /tmp/foundry-install-script && \
    sed -i 's/\/foundry-rs\/foundry\/master\/foundryup/\/foundry-rs\/foundry\/70d00222c2ef74fc484b8d98c5705d131ab31871\/foundryup/g' /tmp/foundry-install-script && \
    cat /tmp/foundry-install-script | bash && \
    echo "export PATH=\"$PATH:/$HOME/.foundry/bin\"">> ~/.bashrc && \
    source ~/.bashrc && \
    export PATH=$PATH:$HOME/.foundry/bin:$HOME/.cargo/bin && \
    foundryup -C 70d0022

# Install python base58 decode library
RUN apt-get update && \
    apt-get install -y python3-pip
RUN pip3 install base58

# Install avalanche-cli
RUN curl -sSfL https://raw.githubusercontent.com/ava-labs/avalanche-cli/main/scripts/install.sh | sh -s -- -b /usr/local/bin v1.3.3
RUN avalanche --version
RUN echo '{"MetricsEnabled}":false}' > ~/.avalanche-cli/config
