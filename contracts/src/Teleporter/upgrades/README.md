# Upgradeability

## Overview

The `TeleporterMessenger` contract is non-upgradable, once a version of the contract is deployed it cannot be changed. This is with the intention of preventing any changes to the deployed contract that could potentially introduce bugs or vulnerabilities.

However, there should still be new versions of the `TeleporterMessenger` contract that provide new features, and `TeleporterRegistry` supports the integration with these new versions.

The `TeleporterRegistry` keeps track of a mapping of `TeleporterMessenger` contract versions to their addresses. This allows us to deploy new versions of the `TeleporterMessenger` contract and then update the `TeleporterRegistry` to point to different versions. The `TeleporterRegistry` can only be updated through a Warp out-of-band message that meets the following requirements:

- `sourceChainAddress` must match `VALIDATORS_SOURCE_ADDRESS = address(0)`
- `sourceChainID` must match the blockchain ID that the registry is deployed on
- `destinationChainID` must match the blockchain ID that the registry is deployed on
- `destinationAddress` must match the address of the registry

## Design

- `TeleporterRegistry` is deployed on each blockchain that needs to keep track of `TeleporterMessenger` contract versions.
- Each registry's mapping of version to contract address is independent of registries on other blockchains, and chains can decide on their own registry mapping entries.
- `TeleporterRegistry` contract can be initialized through a list of initial registry entries, which are `TeleporterMessenger` contract versions and their addresses.
- The registry keeps track of a mapping of `TeleporterMessenger` contract versions to their addresses, and vice versa, a mapping of `TeleporterMessenger` contract addresses to their versions.
- Version zero is an invalid version, and is used to indicate that a `TeleporterMessenger` contract has not been registered yet.
- Once a version number is registered in the registry, it cannot be changed, but a previous registered protocol address can be added to the registry with a new version. This is especially important in the case of a rollback to a previous version, the previous Teleporter contract address can be registered with a new version to the registry.

## How to use `TeleporterRegistry`

`TeleporterUpgradeable` is an abstract contract that helps integrate the `TeleporterRegistry` into a dapp. The dapp contract can inherit `TeleporterUpgradeable`, and pass in the Teleporter registry address inside the constructor, for example in [ERC20Bridge](../ERC20Bridge.sol):

```solidity
contract ERC20Bridge is
    IERC20Bridge,
    ITeleporterReceiver,
    ReentrancyGuard,
    TeleporterUpgradeable
{
    ...
    constructor(
        address teleporterRegistryAddress
    ) TeleporterUpgradeable(teleporterRegistryAddress) {
        currentChainID = WarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();
    }
    ...
}
```

The `TeleporterUpgradeable` contract saves the Teleporter registry in a state variable used by the inheriting contract, and initializes a `_minTeleporterVersion` to the latest `TeleporterMessenger` version registered in `TeleporterRegistry`. The `onlyAllowedTeleporter` modifier ensures that `msg.sender` is a `TeleporterMessenger` contract with a version greater than or equal to `_minTeleporterVersion`. This modifier is used to restrict access to functions that should only be called by a `TeleporterMessenger` contract, i.e. `ITeleporterReceiver.receiveTeleporterMessage`. This is to support the case where a dapp wants to upgrade to a new version of the `TeleporterMessenger` contract, but still wants to be able to receive messages from the old Teleporter contract. Then once the dapp completes delivery of messages from the old Teleporter contract, it can call `TeleporterUpgradeable.updateMinTeleporterVersion` to update the `_minTeleporterVersion` to the new version.

For sending messages with the Teleporter registry, dapps should generally use `TeleporterRegistry.getLatestTeleporter` for the latest version, but if the dapp wants to send a message to a specific version, it can use `TeleporterRegistry.getTeleporterFromVersion` to get the specific Teleporter version.

Using latest version:

```solidity
        ITeleporterMessenger teleporterMessenger = teleporterRegistry
            .getLatestTeleporter();
```

Using specific version:

```solidity
        ITeleporterMessenger teleporterMessenger = teleporterRegistry
            .getTeleporterFromVersion(version);
```