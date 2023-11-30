## Overview

Teleporter is an EVM compatible cross-subnet communication protocol built on top of [Avalanche Warp Messaging (AWM)](https://docs.avax.network/learn/avalanche/awm), and implemented as a Solidity smart contract. It provides a mechanism to asynchronously invoke smart contract functions on other EVM blockchains within Avalanche. Teleporter provides a handful of useful features on top of AWM, such as specifying relayer incentives for message delivery, replay protection, message delivery and execution retries, and a standard interface for sending and receiving messages within a dApp deployed across multiple subnets.

It's important to understand the distinction between Avalanche Warp Messaging and Teleporter. AWM allows subnets to communicate with each other via authenticated messages by providing signing and verification primitives in Avalanchego. These are used by the blockchain VMs to sign outgoing messages and verify incoming messages.

The Teleporter protocol, on the other hand, is implemented at the smart contract level, and is a user-friendly interface to AWM, aimed at dApp developers. All of the message signing and verification is abstracted away from developers. Instead, developers simply call `sendCrossChainMessage` on the `TeleporterMessenger` contract to send a message invoking a smart contract on another subnet, and implement the `ITeleporterReceiver` interface to receive messages on the destination subnet. Teleporter handles all of the Warp message construction and sending, as well as the message delivery and execution.

- [Overview](#overview)
- [Setup](#setup)
  - [Initialize the repository](#initialize-the-repository)
  - [Dependencies](#dependencies)
- [Structure](#structure)
- [Run a local testnet in Docker](#run-a-local-testnet-in-docker)
  - [Start up the local testnet](#start-up-the-local-testnet)
  - [Additional notes](#additional-notes)
- [E2E tests](#e2e-tests)
- [ABI Bindings](#abi-bindings)
- [Docs](#docs)
- [Resources](#resources)

## Setup
### Initialize the repository
- Get all submodules: `git submodule update --init --recursive`

### Dependencies
- Docker and Docker Compose v2
  - The docker image installs the following:
    - [Foundry](https://book.getfoundry.sh/getting-started/installation)
    - [Python3](https://www.python.org/downloads/)
- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo)

## Structure
- `contracts/` is a [Foundry](https://github.com/foundry-rs/foundry) project that includes the implementation of the `TeleporterMessenger` contract and example dApps that demonstrate how to write contracts that interact with Teleporter.
- `abi-bindings/` includes Go ABI bindings for the contracts in `contracts/`.
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.
- `utils/` includes Go utility functions for interacting with the contracts in `contracts/`. Included are Golang scripts to derive the expected EVM contract address deployed from a given EOA at a specific nonce, and also construct a transaction to deploy provided byte code to the same address on any EVM chain using [Nick's method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c#).
- `scripts/` includes bash scripts for interacting with Teleporter in various environments, as well as utility scripts.
  - `abi_bindings.sh` generates ABI bindings for the contracts in `contracts/` and outputs them to `abi-bindings/`.
  - `lint.sh` performs Solidity and Golang linting.
  - `scripts/local/` includes scripts for running Teleporter in Docker.
- `docker/` includes configurations for a local, containerized setup of Teleporter.

## Run a local testnet in Docker
- Install Docker as described in the setup section of the README in the root of this repository.
- If using a local version of the `awm-relayer` image, build it using `./scripts/build_local_image.sh` from the root of the `awm-relayer` repository.

### Start up the local testnet

- Run `./scripts/local/run.sh` to run the local testnet in Docker containers with the ability to interact with the nodes directly.
  - `./scripts/local/run.sh` usage is as follows:
  ```
    -l, --local-relayer-image <tag>   Use a local AWM Relayer image instead of pulling from dockerhub
    -p, --pause                       Pause the network on stop. Will attempt to restart the paused network on subsequent runs
    -h, --help                        Print this help message
  ```
  - Note that if `-l, --local` is not set, then the latest published `awm-relayer` image will be pulled from Dockerhub.
- Once it's running, try sending 1 AVAX on the local C-Chain:
```
cast balance --rpc-url http://127.0.0.1:9652/ext/bc/C/rpc 0x333d17d3b42bf7930dbc6e852ca7bcf560a69003
cast send --private-key 0x56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027 --value 1 0x333d17d3b42bf7930dbc6e852ca7bcf560a69003 --rpc-url http://127.0.0.1:9650/ext/bc/C/rpc
cast balance --rpc-url http://127.0.0.1:9652/ext/bc/C/rpc 0x333d17d3b42bf7930dbc6e852ca7bcf560a69003
```

- After calling `./scripts/local/run.sh`, you can interact with the deployed Teleporter contracts directly, or deploy a cross-chain application contract such as `ExampleCrossChainMessenger`. To open a terminal in the container and initialize it with the environment variables needed to interact with the contracts, run:

```
# Open a shell in the container
docker exec -it local_network_run /bin/bash
# In the container:
set -a                        # export all variables so child processes can access
source vars.sh
```

- Command line tools such as `cast` can also be used from the container.
- Similarly, you can send the above transaction on either of the subnets by replacing the above rpc-url with the subnet's corresponding URL.
- The script `./scripts/local/run_stop.sh` should be used to gracefully shut down the containers, preserving the local network state between runs. This script is called automatically at the end of `./scripts/local/run.sh`, but can be called at any time from a separate terminal to pause the network.
  - `./scripts/local/run_stop.sh` usage is as follows:
  ```
  ./run_stop.sh             # stop the running containers and preserve the network for subsequent runs
  ./run_stop.sh -c          # stop the running containers and clean the network
  ```

### Additional notes

- The `./scripts/local/run.sh` script runs five local network nodes, with each of the nodes validating the primary network and three subnets (Subnet A, Subnet B, and Subnet C).
- Logs from the subnets on one of the five nodes are printed to stdout when run using either script.
- These logs can also be found at `~/.avalanche-cli/runs/network-runner-root-data_<DATE>_<TIMESTAMP>/node{1,5]/logs/<SUBNET_ID>.log` in the `local_network_run` container, or at `/var/lib/docker/overlay2/<CONTAINER_ID>/merged/root/.avalanche-cli/....` on your local machine. You will need to be the root user to access the logs under `/var/lib`.

## E2E tests

E2E tests are run as part of CI, but can also be run locally. To run the E2E tests locally, you'll need to install Gingko following the instructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo).

Then run the following command from the root of the repository:
```bash
./scripts/local/e2e_test.sh
```

### Run the E2E tests on another network

The E2E tests can be run on another network by implementing the `Network` interface in [`package network`](./tests/network/network.go). For example, the type `FujiNetwork` in [`example_fuji_network.go`](./tests/network/example_fuji_network.go) implements this interface, pointing to the [Amplify](https://subnets-test.avax.network/amplify), [Bulletin](https://subnets-test.avax.network/bulletin), and [Conduit](https://subnets-test.avax.network/conduit) Fuji subnets. After implementing this interface, you can run the E2E tests on this network by running a program such as:
```go
func main() {
  // Register a failure handler that panics
	gomega.RegisterFailHandler(func(message string, callerSkip ...int) {
		panic(message)
	})

  // Run the test, composing it with the Network implementation
  network := network.NewFujiNetwork()
  defer network.CloseNetworkConnections()
	tests.BasicOneWaySend(network)
}
```

## ABI Bindings

To generate Golang ABI bindings for the Solidity smart contracts, run:
```bash
./scripts/abi_go_bindings.sh
```

The auto-generated bindings should be written under the `abi-bindings/` directory.

## Docs

- [Teleporter Protocol Overview](./contracts/src/Teleporter/README.md)
- [Cross Chain Applications](./contracts/src/CrossChainApplications/README.md)
- [Getting Started](./contracts/src/CrossChainApplications/GETTING_STARTED.md)
- [Contract Deployment](./utils/contract-deployment/README.md)
- [ERC20Bridge](./contracts/src/CrossChainApplications/ERC20Bridge/README.md)
- [VerifiedBlockHash](./contracts/src/CrossChainApplications/VerifiedBlockHash/README.md)

## Resources

- List of blockchain signing cryptography algorithms [here](http://ethanfast.com/top-crypto.html).
- Background on stateful precompiles [here](https://medium.com/avalancheavax/customizing-the-evm-with-stateful-precompiles-f44a34f39efd).
- Background on BLS signature aggregation [here](https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html).
