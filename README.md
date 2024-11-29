<p align="center">
  <img width="85%" alt="teleporter" src="resources/TeleporterLogo.png?raw=true"/>
</p>

To get started with building ICM contracts, refer to [the avalanche-starter-kit repository](https://github.com/ava-labs/avalanche-starter-kit). This README is focused on the development of the `TeleporterMessenger` contract itself.

`TeleporterMessenger` is a smart contract that serves as the interface for ICM contracts to [Avalanche Interchain Messaging (ICM)](https://academy.avax.network/course/interchain-messaging/04-icm-basics/01-icm-basics). It provides a mechanism to asynchronously invoke smart contract functions on other EVM blockchains within Avalanche. `TeleporterMessenger` provides a handful of useful features of ICM, such as specifying relayer incentives for message delivery, replay protection, message delivery and execution retries, and a standard interface for sending and receiving messages within a dApp deployed across multiple Avalanche L1s.

The `TeleporterMessenger` contract is a user-friendly interface to ICM, aimed at dApp developers. All of the message signing and verification is abstracted away from developers. Instead, developers simply call `sendCrossChainMessage` on the `TeleporterMessenger` contract to send a message invoking a smart contract on another Avalanche L1, and implement the `ITeleporterReceiver` interface to receive messages on the destination Avalanche L1. `TeleporterMessenger` handles all of the ICM message construction and sending, as well as the message delivery and execution.

To get started with using `TeleporterMessenger`, see [How to Deploy ICM Enabled Avalanche L1s on a Local Network](https://docs.avax.network/tooling/cli-cross-chain/teleporter-on-local-networks)

- [Deployed Addresses](#deployed-addresses)
- [A Note on Versioning](#a-note-on-versioning)
- [Setup](#setup)
  - [Initialize the repository](#initialize-the-repository)
  - [Dependencies](#dependencies)
- [Structure](#structure)
- [E2E tests](#e2e-tests)
  - [Run specific E2E tests](#run-specific-e2e-tests)
- [Upgradability](#upgradability)
- [Deploy TeleporterMessenger to an L1](#deploy-teleportermessenger-to-an-avalanche-l1)
- [Deploy TeleporterRegistry to an L1](#deploy-teleporterregistry-to-an-avalanche-l1)
- [ABI Bindings](#abi-bindings)
- [Docs](#docs)
- [Resources](#resources)

## Deployed Addresses

| Contract              | Address                                        | Chain                    |
| --------------------- | ---------------------------------------------- | ------------------------ |
| `TeleporterMessenger` | **0x253b2784c75e510dD0fF1da844684a1aC0aa5fcf** | All chains, all networks |
| `TeleporterRegistry`  | **0x7C43605E14F391720e1b37E49C78C4b03A488d98** | Mainnet C-Chain          |
| `TeleporterRegistry`  | **0xF86Cb19Ad8405AEFa7d09C778215D2Cb6eBfB228** | Fuji C-Chain             |

- Using [Nick's method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c#), `TeleporterMessenger` deploys at a universal address across all chains, varying with each `teleporter` Major release. **Compatibility exists only between same-version `TeleporterMessenger` instances.** See [TeleporterMessenger Contract Deployment](./utils/contract-deployment/README.md) and [Deploy TeleporterMessenger to an Avalanche L1](#deploy-teleportermessenger-to-an-avalanche-l1) for more details.

- `TeleporterRegistry` can be deployed to any address. See [Deploy TeleporterRegistry to an Avalanche L1](#deploy-teleporterregistry-to-an-avalanche-l1) for details. The table above enumerates the canonical registry addresses on the Mainnet and Fuji C-Chains.

## A Note on Versioning

Release versions follow the [semver](https://semver.org/) convention of incompatible Major releases. A new Major version is released whenever the `TeleporterMessenger` bytecode is changed, and a new version of `TeleporterMessenger` is meant to be deployed. Due to the use of Nick's method to deploy the contract to the same address on all chains (see [TeleporterMessenger Contract Deployment](./utils/contract-deployment/README.md) for details), this also means that new release versions would result in different `TeleporterMessenger` contract addresses. Minor and Patch versions may pertain to contract changes that do not change the `TeleporterMessenger` bytecode, or to changes in the test frameworks, and will only be included in tags.

## Setup

### Initialize the repository

- Get all submodules: `git submodule update --init --recursive`

### Dependencies

- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo) for running the end-to-end tests.
- Docker and Docker Compose v2 for running the local test network.
  - The docker image installs the following:
    - [Foundry](https://book.getfoundry.sh/) Use `/scripts/install_foundry.sh` to install the Ava Labs [fork](https://github.com/ava-labs/foundry).
    - [Python3](https://www.python.org/downloads/)

## Structure

- `contracts/` is a [Foundry](https://github.com/foundry-rs/foundry) project that includes the implementation of the `TeleporterMessenger` contract and example dApps that demonstrate how to write contracts that interact with Teleporter.
- `abi-bindings/` includes Go ABI bindings for the contracts in `contracts/`.
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.
- `utils/` includes Go utility functions for interacting with the contracts in `contracts/`. Included are Golang scripts to derive the expected EVM contract address deployed from a given EOA at a specific nonce, and also construct a transaction to deploy provided byte code to the same address on any EVM chain using [Nick's method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c#).
- `scripts/` includes bash scripts for interacting with TeleporterMessenger in various environments, as well as utility scripts.
  - `abi_bindings.sh` generates ABI bindings for the contracts in `contracts/` and outputs them to `abi-bindings/`.
  - `lint.sh` performs Solidity and Golang linting.

## E2E tests

In addition to the docker setup, end-to-end integration tests written using Ginkgo are provided in the `tests/` directory. E2E tests are run as part of CI, but can also be run locally. Any new features or cross-chain example applications checked into the repository should be accompanied by an end-to-end tests. See the [Contribution Guide](./CONTRIBUTING.md) for additional details.

To run the E2E tests locally, you'll need to install Gingko following the instructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo).

Then run the following command from the root of the repository:

```bash
./scripts/e2e_test.sh
```

### Run specific E2E tests

To run a specific E2E test, specify the environment variable `GINKGO_FOCUS`, which will then look for test descriptions that match the provided input. For example, to run the `Calculate Teleporter message IDs` test:

```bash
GINKGO_FOCUS="Calculate Teleporter message IDs" ./scripts/e2e_test.sh
```

A substring of the full test description can be used as well:

```bash
GINKGO_FOCUS="Calculate Teleporter" ./scripts/e2e_test.sh
```

The E2E tests also supports `GINKGO_LABEL_FILTER`, making it easy to group test cases and run them together. For example, to run all E2E tests for the example cross chain applications:

```bash
	ginkgo.It("Send native tokens from L1 A to B and back",
		ginkgo.Label("cross chain apps"),
		func() {
			flows.NativeTokenBridge(LocalNetworkInstance)
		})
```

```bash
GINKGO_LABEL_FILTER="cross chain apps" ./scripts/e2e_test.sh
```

## Upgradability

`TeleporterMessenger` is a non-upgradeable contract and can not be changed once it is deployed. This provides immutability to the contracts, and ensures that the contract's behavior at each address is unchanging. However, to allow for new features and potential bug fixes, new versions of `TeleporterMessenger` can be deployed to different addresses. The [TeleporterRegistry](./contracts/teleporter/TeleporterRegistry.sol) is used to keep track of the deployed versions of Teleporter, and to provide a standard interface for dApps to interact with the different `TeleporterMessenger` versions.

`TeleporterRegistry` **is not mandatory** for dApps built on top of ICM, but dApp's are recommended to leverage the registry to ensure they use the latest `TeleporterMessenger` version available. Another recommendation standard is to have a single canonical `TeleporterRegistry` for each Avalanche L1, and unlike the `TeleporterMessenger` contract, the registry does not need to be deployed to the same address on every chain. This means the registry does not need a Nick's method deployment, and can be at different contract addresses on different chains.

For more information on the registry and how to integrate with ICM contracts, see the [Upgradability doc](./contracts/teleporter/registry/README.md).

## Deploy TeleporterMessenger to an Avalanche L1

From the root of the repo, the TeleporterMessenger contract can be deployed by calling

```bash
./scripts/deploy_teleporter.sh --version <version> --rpc-url <url> [OPTIONS]
```

Required arguments:

- `--version <version>` Specify the release version to deploy. These will all be of the form `v1.X.0`. Each `TeleporterMessenger` version can only send and receive messages from the **same** `TeleporterMessenger` version on another chain. You can see a list of released versions at https://github.com/ava-labs/teleporter/releases.
- `--rpc-url <url>` Specify the rpc url of the node to use.

Options:

- `--private-key <private_key>` Funds the deployer address with the account held by `<private_key>`

To ensure that `TeleporterMessenger` can be deployed to the same address on every EVM based chain, it uses [Nick's Method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c) to deploy from a static deployer address. Teleporter costs exactly `10eth` in the Avalanche L1's native gas token to deploy, which must be sent to the deployer address.

`deploy_teleporter.sh` will send the necessary native tokens to the deployer address if it is provided with a private key for an account with sufficient funds. Alternatively, the deployer address can be funded externally. The deployer address for each version can be found by looking up the appropriate version at https://github.com/ava-labs/teleporter/releases and downloading `TeleporterMessenger_Deployer_Address_<VERSION>.txt`.

Alternatively for new Avalanche L1s, the `TeleporterMessenger` contract can be directly included in the genesis file as documented [here](./contracts/teleporter/README.md#teleporter-messenger-contract-deployment).

## Deploy TeleporterRegistry to an Avalanche L1

There should only be one canonical `TeleporterRegistry` deployed for each chain, but if one does not exist, it is recommended to deploy the registry so ICM contracts can always use the most recent `TeleporterMessenger` version available. The registry does not need to be deployed to the same address on every chain, and therefore does not need a Nick's method transaction. To deploy, run the following command from the root of the repository:

```bash
./scripts/deploy_registry.sh --version <version> --rpc-url <url> --private-key <private_key> [OPTIONS]
```

Required arguments:

- `--version <version>` Specify the release version to deploy. These will all be of the form `v1.X.0`.
- `--rpc-url <url>` Specify the rpc url of the node to use.
- `--private-key <private_key>` Funds the deployer address with the account held by `<private_key>`

`deploy_registry.sh` will deploy a new `TeleporterRegistry` contract for the intended release version, and will also have the corresponding `TeleporterMessenger` contract registered as the initial protocol version.

## ABI Bindings

To generate Golang ABI bindings for the Solidity smart contracts, run:

```bash
./scripts/abi_go_bindings.sh
```

The auto-generated bindings should be written under the `abi-bindings/` directory.

## Docs

- [ICM Protocol Overview](./contracts/teleporter/README.md)
- [Teleporter Registry and Upgrades](./contracts/teleporter/registry/README.md)
- [Contract Deployment](./utils/contract-deployment/README.md)
- [Teleporter CLI](./cmd/teleporter-cli/README.md)

## Resources

- List of blockchain signing cryptography algorithms [here](http://ethanfast.com/top-crypto.html).
- Background on stateful precompiles [here](https://medium.com/avalancheavax/customizing-the-evm-with-stateful-precompiles-f44a34f39efd).
- Background on BLS signature aggregation [here](https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html).
