# Teleporter Token Bridge

## Upgradeability

The teleporter-token-bridge contracts are non-upgradeable and cannot be changed once it is deployed. This provides immutability to the contracts, and ensures that the contract's behavior at each address is unchanging.

## Overview

Teleporter token bridge is an application that allows users to transfer tokens between Subnets. The bridge is a set of smart contracts that are deployed across multiple Subnets, and leverages [Teleporter](https://github.com/ava-labs/teleporter) for cross-chain communication.

Each bridge instance consists of one "home" contract and at least one but possibly many "remote" contracts. Each home contract instance manages one asset to be bridged out to `TokenRemote` instances. The home contract lives on the Subnet where the asset to be bridged exists and locks that asset as collateral to be bridged to other Subnets. The remote contracts, each of which has a single specified home contract, live on other Subnets that want to import the asset bridged by their specified home. The token bridges are designed to be permissionless: anyone can register compatible `TokenRemote` instances to allow for bridging tokens from the `TokenHome` instance to that new `TokenRemote` instance. The home contract keeps track of token balances bridged to each `TokenRemote` instance, and handles returning the original tokens back to the user when assets are bridged back to the `TokenHome` instance. `TokenRemote` instances are registered with their home contract via a Teleporter message upon creation.

Home contract instances specify the asset to be bridged as either an ERC20 token or the native token, and they allow for transferring the token to any registered `TokenRemote` instances. The token representation on the remote chain can also either be an ERC20 or native token, allowing users to have any combination of ERC20 and native tokens between home and remote chains:

- `ERC20` -> `ERC20`
- `ERC20` -> `Native`
- `Native` -> `ERC20`
- `Native` -> `Native`

The remote tokens are designed to have compatibility with the token bridge on the home chain by default, and they allow custom logic to be implemented in addition. For example, developers can inherit and extend the `ERC20TokenRemote` contract to add additional functionality, such as a custom minting, burning, or transfer logic.

The token bridge also supports "multi-hop" transfers, where tokens can be transferred between remote chains. To illustrate, consider two remotes _R<sub>a</sub>_ and _R<sub>b</sub>_ that are both connected to the same home _H_. A multi-hop transfer from _R<sub>a</sub>_ to _R<sub>b</sub>_ first gets routed from _R<sub>a</sub>_ to _H_, where remote balances are updated, and then _H_ automatically routes the transfer on to _R<sub>b</sub>_.

In addition to supporting basic token transfers, the token bridge contracts offer a `sendAndCall` interface for bridging tokens and using them in a smart contract interaction all within a single Teleporter message. If the call to the recipient smart contract fails, the bridged tokens are sent to a fallback recipient address on the destination chain of the transfer. The `sendAndCall` interface enables the direct use of bridged tokens in dApps on other chains, such as performing swaps, using the tokens to pay for fees when invoking services, etc.

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
[⠒] Compiling 78 files with 0.8.18
[⠆] Solc 0.8.18 finished in 3.92s
Compiler run successful!
Analysing contracts...
Running tests...
| File                                        | % Lines           | % Statements      | % Branches        | % Funcs         |
|---------------------------------------------|-------------------|-------------------|-------------------|-----------------|
| src/TokenHome/ERC20TokenHome.sol              | 100.00% (16/16)   | 100.00% (19/19)   | 100.00% (4/4)     | 100.00% (6/6)   |
| src/TokenHome/NativeTokenHome.sol             | 100.00% (14/14)   | 100.00% (16/16)   | 100.00% (2/2)     | 100.00% (6/6)   |
| src/TokenHome/TokenHome.sol                   | 100.00% (135/135) | 100.00% (159/159) | 100.00% (84/84)   | 100.00% (15/15) |
| src/TokenRemote/ERC20TokenRemote.sol          | 100.00% (27/27)   | 100.00% (31/31)   | 100.00% (8/8)     | 100.00% (8/8)   |
| src/TokenRemote/NativeTokenRemote.sol         | 100.00% (46/46)   | 100.00% (57/57)   | 100.00% (12/12)   | 100.00% (13/13) |
| src/TokenRemote/TokenRemote.sol               | 100.00% (76/76)   | 100.00% (94/94)   | 100.00% (52/52)   | 100.00% (14/14) |
| src/WrappedNativeToken.sol                  | 100.00% (6/6)     | 100.00% (6/6)     | 100.00% (0/0)     | 100.00% (3/3)   |
| src/mocks/ExampleERC20Decimals.sol          | 100.00% (1/1)     | 100.00% (1/1)     | 100.00% (0/0)     | 100.00% (1/1)   |
| src/mocks/MockERC20SendAndCallReceiver.sol  | 100.00% (5/5)     | 100.00% (5/5)     | 100.00% (4/4)     | 100.00% (2/2)   |
| src/mocks/MockNativeSendAndCallReceiver.sol | 100.00% (4/4)     | 100.00% (4/4)     | 100.00% (4/4)     | 100.00% (2/2)   |
| src/utils/CallUtils.sol                     | 100.00% (8/8)     | 100.00% (9/9)     | 100.00% (6/6)     | 100.00% (2/2)   |
| src/utils/SafeWrappedNativeTokenDeposit.sol | 100.00% (5/5)     | 100.00% (8/8)     | 100.00% (2/2)     | 100.00% (1/1)   |
| src/utils/TokenScalingUtils.sol             | 100.00% (8/8)     | 100.00% (14/14)   | 100.00% (2/2)     | 100.00% (4/4)   |
| Total                                       | 100.00% (351/351) | 100.00% (423/423) | 100.00% (180/180) | 100.00% (77/77) |
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

The E2E tests also supports `GINKGO_LABEL_FILTER`, making it easy to group test cases and run them together. For example, to run all `ERC20TokenHome` E2E tests:

```go
	ginkgo.It("Bridge an ERC20 token between two Subnets",
		ginkgo.Label(erc20TokenHomeLabel, erc20TokenRemoteLabel),
		func() {
			flows.ERC20TokenHomeERC20TokenRemote(LocalNetworkInstance)
		})
```

```bash
GINKGO_LABEL_FILTER="ERC20TokenHome" ./scripts/e2e_test.sh
```
