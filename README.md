## Overview

Teleporter is an EVM compatible cross-subnet communication protocol built on top of [Avalanche Warp Messaging (AWM)](https://docs.avax.network/learn/avalanche/awm), and implemented as a Solidity smart contract. It provides a mechanism to asynchronously invoke smart contract functions on other EVM blockchains within Avalanche. Teleporter provides a handful of useful features on top of AWM, such as specifying relayer incentives for message delivery, replay protection, message delivery and execution retries, and a standard interface for sending and receiving messages within a dApp deployed across multiple subnets.

It's important to understand the distinction between Avalanche Warp Messaging and Teleporter. AWM allows subnets to communicate with each other via authenticated messages by providing signing and verification primitives in Avalanchego. These are used by the blockchain VMs to sign outgoing messages and verify incoming messages.

The Teleporter protocol, on the other hand, is implemented at the smart contract level, and is a user-friendly interface to AWM, aimed at dApp developers. All of the message signing and verification is abstracted away from developers. Instead, developers simply call `sendCrossChainMessage` on the `TeleporterMessenger` contract to send a message invoking a smart contract on another subnet, and implement the `ITeleporterReceiver` interface to receive messages on the destination subnet. Teleporter handles all of the Warp message construction and sending, as well as the message delivery and execution.

- [Overview](#overview)
- [Setup](#setup)
  - [Docker Setup](#docker-setup)
  - [General Setup](#general-setup)
- [Structure](#structure)
- [Build + Run + Test](#build--run--test)
  - [Run tests on Fuji Testnet](#run-tests-on-fuji-testnet)
  - [E2E tests](#e2e-tests)
- [Docs](#docs)
- [Resources](#resources)

## Setup

### Docker Setup

- Install Docker:

```
sudo apt-get install \
    ca-certificates \
    curl \
    gnupg \
    lsb-release
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io
sudo systemctl start docker # Docker will start automatically on future system boots
```

- Install Docker Compose v2:

```
sudo apt-get install docker-compose-plugin
docker compose version
```

- Set Docker to run without `sudo`:

```
sudo groupadd docker
sudo usermod -aG docker $USER
sudo reboot
docker run hello-world # This should work without sudo now
```

- Note that as you develop and continuously build Docker images, it will eat up continuously more disk space. Periodically you'll want to remedy this be removing everything Docker has built: `docker system prune --all --volumes --force`

### General Setup

The above steps are sufficient to run the included integration tests inside Docker containers. If you wish to run them outside docker, you'll need to install the following dependencies:
- [Foundry](https://book.getfoundry.sh/getting-started/installation)
- [Python3](https://www.python.org/downloads/)

## Structure

- `contracts/` is a [Foundry](https://github.com/foundry-rs/foundry) project that includes the implementation of the `TeleporterMessenger` contract and example dApps that demonstrate how to write contracts that interact with Teleporter.
- `abi-bindings/` includes Go ABI bindings for the contracts in `contracts/`.
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.
- `utils/` includes Go utility functions for interacting with the contracts in `contracts/`. Included are Golang scripts to derive the expected EVM contract address deployed from a given EOA at a specific nonce, and also construct a transaction to deploy provided byte code to the same address on any EVM chain using [Nick's method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c#).
- `subnet-evm/` is the public subnet-evm repository (included as a submodule) checked out on the `warp-contract` branch with our changes.
- `scripts/` includes bash scripts for interacting with Teleporter in various environments, as well as utility scripts.
  - `abi_bindings.sh` generates ABI bindings for the contracts in `contracts/` and outputs them to `abi-bindings/`.
  - `lint.sh` lints the contracts in `contracts/`.
  - `scripts/local/` includes scripts for running Teleporter in Docker containers and for running the tests in the `scripts/local/integration-tests` locally.
    - `scripts/local/integration-tests/` includes integration test scripts written in bash. The scripts use `foundry` to deploy smart contracts that use Teleporter, send transactions to interact with the contracts, and check that cross-chain messages have the expected effect on destination chains.
      - *Note* These tests will be deprecated in favor of the end to end tests in `tests/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.
    - `scripts/fuji/` includes scripts to interact with a live Teleporter deployment on Fuji subnets.
      - `scripts/fuji/example-workflows/` includes example workflows that send transactions to interact with Teleporter contracts on [Fuji](https://docs.avax.network/learn/avalanche/fuji) subnets.
- `docker/` includes configurations for a local, containerized setup of Teleporter.

## Run the Docker integration tests

- Get all submodules: `git submodule update --init --recursive`
- Install Docker as described in the setup section of the README in the root of this repository.
- If we are using a local version of the `awm-relayer` image, build it using `./scripts/build_local_image.sh` from the root of the `awm-relayer` repository.

### Start up the local testnet

- Run `./scripts/local/run.sh` to run the local testnet in Docker containers with the ability to interact with the nodes directly.
  - `./scripts/local/run.sh` usage is as follows:
  ```
    -l, --local-relayer-image <tag>   Use a local AWM Relayer image instead of pulling from dockerhub
    -h, --help                        Print this help message
  ```
  - Note that if `-l, --local` is not set, then the latest published `awm-relayer` image will be pulled from Dockerhub.
- Once it's running, try sending 1 AVAX on the local C-Chain:
```
cast balance --rpc-url http://127.0.0.1:9652/ext/bc/C/rpc 0x333d17d3b42bf7930dbc6e852ca7bcf560a69003
cast send --private-key 0x56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027 --value 1 0x333d17d3b42bf7930dbc6e852ca7bcf560a69003 --rpc-url http://127.0.0.1:9650/ext/bc/C/rpc
cast balance --rpc-url http://127.0.0.1:9652/ext/bc/C/rpc 0x333d17d3b42bf7930dbc6e852ca7bcf560a69003
```

- After calling `./scripts/local/run.sh`, you can directly send messages between the deployed subnets. As an example, `./scripts/integration-tests/basic_send_receive.sh` can be run manually like so:

```
# Open a shell in the container
docker exec -it relayer_run /bin/bash
# In the container:
set -a                        # export all variables so child processes can access
source vars.sh
./integration-tests/basic_send_receive.sh   # run a test manually
```

- Command line tools such as `cast` can also be used from the container.
- Similarly, you can send the above transaction on either of the subnets by replacing the above rpc-url with the subnet's corresponding URL.
- Run `./scripts/local/test.sh` to execute the integration test as it's run in Continuous Integration.
  - `./scripts/local/test.sh` usage is as follows:
  ```
    -t, --test <test_name>            Run a specific test. If empty, runs all tests in the ./scripts/local/integration-tests/
    -t, --test "test1 test2"          Run multiple tests. Test names must be space delimited and enclosed in quotes
    -l, --local-relayer-image <tag>   Use a local AWM Relayer image instead of pulling from dockerhub
    -h, --help                        Print this help message
  ```
  - Note that if `-l, --local` is not set, then the latest published `awm-relayer` image will be pulled from Dockerhub.
- The script `./scripts/local/run_stop.sh` should be used to gracefully shut down the containers, preserving the local network state between runs. This script is called automatically at the end of `./scripts/local/test.sh`, but can be called at any time from a separate terminal to pause the network. This script should also be used with `run.sh` from a separate terminal.
  - `./scripts/local/run_stop.sh` usage is as follows:
  ```
  ./run_stop.sh             # stop the running containers and preserve the network for subsequent runs
  ./run_stop.sh -c          # stop the running containers and clean the network
  ```

### Run the integration tests in Docker containers

- Run `./scripts/local/test.sh` to run the integration tests in Docker containers.
  - `./scripts/local/test.sh` usage is as follows:
  ```
    -t, --test <test_name>            Run a specific test. If empty, runs all tests in the ./scripts/local/integration-tests/
    -t, --test "test1 test2"          Run multiple tests. Test names must be space delimited and enclosed in quotes
    -l, --local-relayer-image <tag>   Use a local AWM Relayer image instead of pulling from dockerhub
    -h, --help                        Print this help message
  ```
  - Note that if `-l, --local` is not set, then the latest published `awm-relayer` image will be pulled from Dockerhub.
- This script performs the same setup steps as `scripts/local/run.sh` (described above), and then runs the tests in `./scripts/integration-tests/` automatically

### Additional notes

- Both the `./scripts/local/run.sh` and `./scripts/local/test.sh` scripts run local five node networks, with each of the nodes validating the primary network and three subnets (Subnet A, Subnet B, and Subnet C).
- Logs from the subnets on one of the five nodes are printed to stdout when run using either script.
- These logs can also be found at `~/.avalanche-cli/runs/network-runner-root-data_<DATE>_<TIMESTAMP>/node{1,5]/logs/<SUBNET_ID>.log`, or at `/var/lib/docker/overlay2/<CONTAINER_ID>/merged/root/.avalanche-cli/....` on your local machine. You will need to be the root user to access the logs under `/var/lib`.

### Run tests on Fuji Testnet

The following steps will allow you to run the integration tests described above and example workflows against three Fuji subnets that have AWM enabled. These workflows send transaction on each of these subnets, so you will need to fund your account with some native tokens to pay for gas fees. The three subnets are called [Amplify](https://subnets-test.avax.network/amplify), [Bulletin](https://subnets-test.avax.network/bulletin), and [Conduit](https://subnets-test.avax.network/conduit).

The configuration for all the subnets is in `.env.testnet`. If needed, you can change the subnet IDs, the chain IDs, urls and teleporter contract address to match your setup.

- Step 1: Fund your account on each of the three Fuji subnets using the faucet.
  - [Amplify](https://core.app/tools/testnet-faucet/?subnet=amplify&token=amplify)
  - [Bulletin](https://core.app/tools/testnet-faucet/?subnet=bulletin&token=bulletin)
  - [Conduit](https://core.app/tools/testnet-faucet/?subnet=conduit&token=conduit)
- Step 2: Set the private key and address of your testnet account.
  ```bash
  cp .env.example .env
  # Edit .env to set the private key and address of your testnet account
  ```
- Step 3: Run tests
  ```bash
  ./scripts/fuji/test.sh -t <test_name>
  ```

`./scripts/fuji/test.sh` usage is as follows:
```
  -t, --test <test_name>            Run a specific test. If empty, runs all tests in the ./scripts/fuji/example-workflows/
  -t, --test "test1 test2"          Run multiple tests. Test names must be space delimited and enclosed in quotes
  -h, --help                        Print this help message
```

To setup all environment variables manually, you can run the following commands:
```bash
# set the subnet IDs, chain IDs, urls and teleporter contract address
source .env.testnet
# set the private key and address of your testnet account
source .env
```

## E2E tests

E2E tests are ran as part of CI, but can also be ran locally with the `--local` flag. To run the E2E tests locally, you'll need to install Gingko following the intructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo)

Next, provide the path to the `subnet-evm` repository and the path to a writeable data directory (here we use the `subnet-evm` submodule and `~/tmp/e2e-test`) to use for the tests:
```bash
./scripts/local/e2e_test.sh --local --subnet-evm ./subnet-evm --data-dir ~/tmp/e2e-test
```

## ABI Bindings

To generate Golang ABI bindings for the Solidity smart contracts, run:
```bash
./scripts/abi_go_bindings.sh
```

The auto-generated bindings should be written under the `abi-bindings/` directory.

## Docs

- [Teleporter Protocol Overview](./contracts/src/Teleporter/README.md)
- [Getting Started](./contracts/src/CrossChainApplications/README.md)
- [Contract Deployment](./contract-deployment/README.md)
- [ERC20Bridge](./contracts/src/CrossChainApplications/ERC20Bridge/README.md)
- [VerifiedBlockHash](./contracts/src/CrossChainApplications/VerifiedBlockHash/README.md)

## Resources

- List of blockchain signing cryptography algorithms [here](http://ethanfast.com/top-crypto.html).
- Background on stateful precompiles [here](https://medium.com/avalancheavax/customizing-the-evm-with-stateful-precompiles-f44a34f39efd).
- Tutorial on stateful precompiles [here](https://github.com/ava-labs/subnet-evm/blob/precompile-tutorial-only/cmd/README.md).
- Background on BLS signature aggregation [here](https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html).
