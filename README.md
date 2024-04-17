# Teleporter Token Bridge

## Status

Please note that `teleporter-token-bridge` is still under active development and should not be used in production. The code is not yet audited and will be subject to active changes.

## Overview

Teleporter token bridge is an application that allows users to transfer tokens between subnets. The bridge is a set of smart contracts that are deployed on both the source and destination subnets, and leverages [Teleporter](https://github.com/ava-labs/teleporter) for cross-chain communication. The token bridges are designed to be permissionless: as long as there is a compatible token bridge instance on a destination chain, users can transfer tokens from a home chain to that destination chain. The token bridge source on the home chain keeps track of token balances bridged to each destination token bridge instance, and handles returning the original tokens back to the user when bridged back to the home chain.

The token bridge contracts take in either an ERC20 or native token to be bridged from a source chain, which can be referred to as the "home chain", and transfers the token to a destination chain to be represented as a new token. The new token representation on the destination chain can also either be an ERC20 or native token, allowing users to have any combination of ERC20 and native tokens between source and destination chains:

- `ERC20` -> `ERC20`
- `ERC20` -> `Native`
- `Native` -> `ERC20`
- `Native` -> `Native`

The destination tokens are designed to by default have compatibility with the token bridges on the source chain, and allow any custom logic to be implemented in addition. For example, developers can inherit and extend the destination ERC20 token contract to add additional functionality, such as a custom minting, burning, or transfer logic.

The token bridge also supports "multi-hop" transfers, where tokens can be transferred between destination chains. The multi-hop transfer first transfers the token from the origin destination chain to the home chain, where token balances are updated, and then triggers a second transfer to the final destination chain.

In addition to supporting basic token transfers, the token bridge contracts offer a `sendAndCall` interface for atomically bridging tokens and using them to interact with a smart contract on the destination chain. If the call to the recipient smart contract fails, the bridged tokens are sent to a fallback recipient address. The `sendAndCall` interfaces enables the direct use of bridged tokens in dApps on other chains, such as performing swaps, using the tokens to pay for fees when invoking services, etc.

A breakdown of the structure of the contracts that implement this function can be found under `./contracts` [here](./contracts/README.md).

## Setup

### Initialize the repository

- Get all submodules: `git submodule update --init --recursive`

### Dependencies

- [Foundry](https://book.getfoundry.sh/getting-started/installation)
- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo) for running the end-to-end tests

## Structure

- `contracts/` is a Foundry project that includes the implementation of the token bridge contracts and Solidity unit tests
- `scripts/` includes various bash utility scripts
- `tests/` includes integration tests for the contracts in `contracts/`, written using the [Ginkgo](https://onsi.github.io/ginkgo/) testing framework.

## Solidity Unit Tests

Unit tests are written under `contracts/test/` and can be run with `forge`:

```
cd contracts
forge test -vvv
```

Unit test coverage of the contracts can be viewed using `forge coverage`:
```
$ forge coverage
[⠢] Compiling...
[⠒] Compiling 70 files with 0.8.18
[⠆] Solc 0.8.18 finished in 4.98s
Compiler run successful!
Analysing contracts...
Running tests...
| File                                        | % Lines           | % Statements      | % Branches       | % Funcs         |
|---------------------------------------------|-------------------|-------------------|------------------|-----------------|
| src/ERC20Destination.sol                    | 100.00% (20/20)   | 100.00% (22/22)   | 100.00% (4/4)    | 100.00% (7/7)   |
| src/ERC20Source.sol                         | 100.00% (14/14)   | 100.00% (17/17)   | 100.00% (4/4)    | 100.00% (5/5)   |
| src/NativeTokenDestination.sol              | 100.00% (56/56)   | 100.00% (68/68)   | 100.00% (14/14)  | 100.00% (14/14) |
| src/NativeTokenSource.sol                   | 100.00% (14/14)   | 100.00% (15/15)   | 100.00% (2/2)    | 100.00% (5/5)   |
| src/TeleporterTokenDestination.sol          | 100.00% (62/62)   | 100.00% (71/71)   | 88.46% (46/52)   | 100.00% (6/6)   |
| src/TeleporterTokenSource.sol               | 100.00% (50/50)   | 100.00% (55/55)   | 97.22% (35/36)   | 100.00% (4/4)   |
| src/mocks/ExampleWAVAX.sol                  | 100.00% (6/6)     | 100.00% (6/6)     | 100.00% (0/0)    | 100.00% (3/3)   |
| src/mocks/MockERC20SendAndCallReceiver.sol  | 100.00% (3/3)     | 100.00% (3/3)     | 100.00% (2/2)    | 100.00% (1/1)   |
| src/mocks/MockNativeSendAndCallReceiver.sol | 100.00% (2/2)     | 100.00% (2/2)     | 100.00% (2/2)    | 100.00% (1/1)   |
| src/utils/CallUtils.sol                     | 100.00% (8/8)     | 100.00% (9/9)     | 100.00% (6/6)    | 100.00% (2/2)   |
| src/utils/SafeWrappedNativeTokenDeposit.sol | 100.00% (5/5)     | 100.00% (8/8)     | 100.00% (2/2)    | 100.00% (1/1)   |
| Total                                       | 100.00% (240/240) | 100.00% (276/276) | 94.35% (117/124) | 100.00% (49/49) |
```

## E2E tests

End-to-end integration tests written using Ginkgo are provided in the `tests/` directory. E2E tests are run as part of CI, but can also be run locally. Any new features or cross-chain example applications checked into the repository should be accompanied by an end-to-end tests.

To run the E2E tests locally, you'll need to install Gingko following the instructions [here](https://onsi.github.io/ginkgo/#installing-ginkgo).

Then run the following command from the root of the repository:

```bash
./scripts/e2e_test.sh
```

### Run specific E2E tests

To run a specific E2E test, specify the environment variable `GINKGO_FOCUS`, which will then look for test descriptions that match the provided input. For example, to run the `Bridge an ERC20 token between two Subnets` test:

```bash
GINKGO_FOCUS="Bridge an ERC20 token between two Subnets" ./scripts/e2e_test.sh
```

A substring of the full test description can be used as well:

```bash
GINKGO_FOCUS="Bridge an ERC20 token" ./scripts/e2e_test.sh
```

The E2E tests also supports `GINKGO_LABEL_FILTER`, making it easy to group test cases and run them together. For example, to run all `ERC20Source` E2E tests:

```bash
	ginkgo.It("Bridge an ERC20 token between two Subnets",
		ginkgo.Label(erc20SourceLabel, erc20DestinationLabel),
		func() {
			flows.BasicERC20SendReceive(LocalNetworkInstance)
		})
```

```bash
GINKGO_LABEL_FILTER="ERC20Source" ./scripts/e2e_test.sh
```
