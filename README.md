# Teleporter Token Bridge

## Status

Please note that `teleporter-token-bridge` is still under active development and should not be used in production. The code is not yet audited and will be subject to active changes.

## Overview

Teleporter token bridge is an application that allows users to transfer tokens between Subnets. The bridge is a set of smart contracts that are deployed across multiple Subnets, and leverages [Teleporter](https://github.com/ava-labs/teleporter) for cross-chain communication.

Each bridge instance consists of one "hub" contract and at least one but possibly many "spoke" contracts. Each hub contract instance manages one asset to be bridged out to spoke instances. The hub contract lives on the Subnet where the asset to be bridged exists and locks that asset as collateral to be bridged to other Subnets. The spoke contracts, each of which has a single specified hub contract, live on other Subnets that want to import the asset bridged by their specified hub. The token bridges are designed to be permissionless: anyone can register compatible spoke instances to allow for bridging tokens from the hub instance to that new spoke. The token bridge hub contract keeps track of token balances bridged to each spoke instance, and handles returning the original tokens back to the user when bridged back to the hub instance. Spoke instances are registered with their hub contract via a Teleporter message upon creation.

Hub contract instances specify either an ERC20 token or the native token to be bridged, and allow for transferring the token to any registered spoke instances. The token representation on the spoke chain can also either be an ERC20 or native token, allowing users to have any combination of ERC20 and native tokens between hub and spoke chains:

- `ERC20` -> `ERC20`
- `ERC20` -> `Native`
- `Native` -> `ERC20`
- `Native` -> `Native`

The spoke tokens are designed to by default have compatibility with the token bridge on the hub chain, and allow custom logic to be implemented in addition. For example, developers can inherit and extend the `ERC20TokenSpoke` contract to add additional functionality, such as a custom minting, burning, or transfer logic.

The token bridge also supports "multi-hop" transfers, where tokens can be transferred between spoke chains. To illustrate, consider two spokes _S<sub>a</sub>_ and _S<sub>b</sub>_ that are both connected to the same hub _H_. A multi-hop transfer from _S<sub>a</sub>_ to _S<sub>b</sub>_ first gets routed from _S<sub>a</sub>_ to _H_, where spoke balances are updated, and then _H_ automatically routes the transfer on to _S<sub>b</sub>_.

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
[⠔] Compiling 78 files with 0.8.18
[⠘] Solc 0.8.18 finished in 5.51s
Compiler run successful!
Analysing contracts...
Running tests...
| File                                        | % Lines           | % Statements      | % Branches        | % Funcs         |
|---------------------------------------------|-------------------|-------------------|-------------------|-----------------|
| src/TokenHub/ERC20TokenHub.sol              | 100.00% (16/16)   | 100.00% (19/19)   | 100.00% (4/4)     | 100.00% (6/6)   |
| src/TokenHub/NativeTokenHub.sol             | 100.00% (14/14)   | 100.00% (16/16)   | 100.00% (2/2)     | 100.00% (6/6)   |
| src/TokenHub/TokenHub.sol                   | 100.00% (135/135) | 100.00% (159/159) | 100.00% (84/84)   | 100.00% (15/15) |
| src/TokenSpoke/ERC20TokenSpoke.sol          | 100.00% (20/20)   | 100.00% (22/22)   | 100.00% (4/4)     | 100.00% (7/7)   |
| src/TokenSpoke/NativeTokenSpoke.sol         | 100.00% (39/39)   | 100.00% (48/48)   | 100.00% (8/8)     | 100.00% (12/12) |
| src/TokenSpoke/TokenSpoke.sol               | 100.00% (81/81)   | 100.00% (102/102) | 100.00% (56/56)   | 100.00% (15/15) |
| src/WrappedNativeToken.sol                  | 100.00% (6/6)     | 100.00% (6/6)     | 100.00% (0/0)     | 100.00% (3/3)   |
| src/mocks/ExampleERC20Decimals.sol          | 100.00% (1/1)     | 100.00% (1/1)     | 100.00% (0/0)     | 100.00% (1/1)   |
| src/mocks/MockERC20SendAndCallReceiver.sol  | 100.00% (5/5)     | 100.00% (5/5)     | 100.00% (4/4)     | 100.00% (2/2)   |
| src/mocks/MockNativeSendAndCallReceiver.sol | 100.00% (4/4)     | 100.00% (4/4)     | 100.00% (4/4)     | 100.00% (2/2)   |
| src/utils/CallUtils.sol                     | 100.00% (8/8)     | 100.00% (9/9)     | 100.00% (6/6)     | 100.00% (2/2)   |
| src/utils/SafeWrappedNativeTokenDeposit.sol | 100.00% (5/5)     | 100.00% (8/8)     | 100.00% (2/2)     | 100.00% (1/1)   |
| src/utils/TokenScalingUtils.sol             | 100.00% (8/8)     | 100.00% (14/14)   | 100.00% (2/2)     | 100.00% (4/4)   |
| Total                                       | 100.00% (342/342) | 100.00% (413/413) | 100.00% (176/176) | 100.00% (76/76) |
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

The E2E tests also supports `GINKGO_LABEL_FILTER`, making it easy to group test cases and run them together. For example, to run all `ERC20TokenHub` E2E tests:

```bash
	ginkgo.It("Bridge an ERC20 token between two Subnets",
		ginkgo.Label(erc20TokenHubLabel, erc20TokenSpokeLabel),
		func() {
			flows.ERC20TokenHubERC20TokenSpoke(LocalNetworkInstance)
		})
```

```bash
GINKGO_LABEL_FILTER="ERC20TokenHub" ./scripts/e2e_test.sh
```
