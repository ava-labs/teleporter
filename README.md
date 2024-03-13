# Teleporter-token-bridge

## Status

Please note that `teleporter-token-bridge` is still under active development and should not be used in production. The code is not yet audited and will be subject to active changes.

## Overview

Teleporter token bridge is an application that allows users to transfer tokens between subnets. The bridge is a set of smart contracts that are deployed on both the source and destination subnets, and leverages [Teleporter](https://github.com/ava-labs/teleporter) for cross-chain communication. The token bridges are designed to be permissionless, as long as there is a compatible token bridge instance on a destination chain, users can transfer tokens from a home chain to that destination chain. The token bridge source on the home chain keeps track of token balances bridged to each destination token bridge instance, and handles returning the original tokens back to the user when bridged back to the home chain.

The token bridge takes in either an ERC20 or native token to be bridged from a source chain, which can be referred to as the "home chain", and transfers the token to a destination chain to be represented as a new token. The new token representation on the destination chain can also either be an ERC20 or native token, allowing users to have any combination of ERC20 and native tokens between source and destination chains:

- ERC20 -> ERC20
- ERC20 -> Native
- Native -> ERC20
- Native -> Native

The destination tokens are designed to by default have compatibility with the token bridges on the source chain, and allow any custom logic to be implemented in addition. For example, developers can inherit and extend the destination ERC20 token contract to add additional functionality, such as a custom minting or burning logic.

Other than simple transfers from source to destination chains, the token bridge also supports "multihop" transfers, where tokens can be transferred between destination chains. The multihop transfer first transfers the token from the first destination chain to the home chain, where token balances are updated, and then triggers a second transfer to the final destination chain.

## Setup

### Initialize the repository

- Get all submodules: `git submodule update --init --recursive`

### Dependencies

- [Foundry](https://book.getfoundry.sh/getting-started/installation)
- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo) for running the end-to-end tests

## Structure

- `contracts/` is a Foundry project that includes the implementation of the token bridge contracts
- `scripts` includes various bash scripts for utility
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.

## Solidity Unit Tests

Unit tests are written under `contracts/test/` and can be run with `forge`:

```
cd contracts/test
forge test -vvv
```

## E2E tests

End-to-end integration tests written using Ginkgo are provided in the `tests/` directory. E2E tests are run as part of CI, but can also be run locally. Any new features or cross-chain example applications checked into the repository should be accompanied by an end-to-end tests.

To run the E2E tests locally, you'll need to install Gingko following the instructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo).

Then run the following command from the root of the repository:

```bash
./scripts/local/e2e_test.sh
```

### Run specific E2E tests

To run a specific E2E test, specify the environment variable `GINKGO_FOCUS`, which will then look for test descriptions that match the provided input. For example, to run the `Calculate Teleporter message IDs` test:

```bash
GINKGO_FOCUS="TODO" ./scripts/local/e2e_test.sh
```

A substring of the full test description can be used as well:

```bash
GINKGO_FOCUS="TODO" ./scripts/local/e2e_test.sh
```

The E2E tests also supports `GINKGO_LABEL_FILTER`, making it easy to group test cases and run them together. For example, to run all E2E tests for the example cross chain applications:

```bash
	ginkgo.It("Send native tokens from subnet A to B and back",
		ginkgo.Label("cross chain apps"),
		func() {
			flows.NativeTokenBridge(LocalNetworkInstance)
		})
```

```bash
GINKGO_LABEL_FILTER="cross chain apps" ./scripts/local/e2e_test.sh
```
