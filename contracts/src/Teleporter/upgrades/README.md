# Upgradeability

## Overview

The `TeleporterMessenger` contract is non-upgradable, once a version of the contract is deployed it cannot be changed. This is with the intention of preventing any changes to the deployed contract that could potentially introduce bugs or vulnerabilities.

However, there could still be new versions of `TeleporterMessenger` contracts needed to be deployed in the future. `TeleporterRegistry` provides applications that use a `TeleporterMessenger` instance a minimal step process to integrate with new versions of `TeleporterMessenger`.

The `TeleporterRegistry` maintains a mapping of `TeleporterMessenger` contract versions to their addresses. When a new `TeleporterMessenger` version is deployed, its address can be added to the `TeleporterRegistry`. The `TeleporterRegistry` can only be updated through a Warp off-chain message that meets the following requirements:

- `sourceChainAddress` must match `VALIDATORS_SOURCE_ADDRESS = address(0)`
  - The zero address can only be set as the source chain address by a Warp off-chain message, and cannot be set by an on-chain Warp message.
- `sourceBlockchainID` must match the blockchain ID that the registry is deployed on
- `destinationBlockchainID` must match the blockchain ID that the registry is deployed on
- `destinationAddress` must match the address of the registry

In the `TeleporterRegistry` contract, the `latestVersion` state variable returns the highest version number that has been registered in the registry. The `getLatestTeleporter` function returns the `ITeleporterMessenger` that is registered with the corresponding version.

## Design

- `TeleporterRegistry` is deployed on each blockchain that needs to keep track of `TeleporterMessenger` contract versions.
- Each registry's mapping of version to contract address is independent of registries on other blockchains, and chains can decide on their own registry mapping entries.
- `TeleporterRegistry` contract can be initialized through a list of initial registry entries, which are `TeleporterMessenger` contract versions and their addresses.
- The registry keeps track of a mapping of `TeleporterMessenger` contract versions to their addresses, and vice versa, a mapping of `TeleporterMessenger` contract addresses to their versions.
- Version zero is an invalid version, and is used to indicate that a `TeleporterMessenger` contract has not been registered yet.
- Once a version number is registered in the registry, it cannot be changed, but a previous registered protocol address can be added to the registry with a new version. This is especially important in the case of a rollback to a previous Teleporter version, in which case the previous Teleporter contract address would need to be registered with a new version to the registry.

## How to use `TeleporterRegistry`

`TeleporterUpgradeable` is an abstract contract that helps integrate the `TeleporterRegistry` into a dapp. The dapp contract can inherit `TeleporterUpgradeable`, and pass in the Teleporter registry address inside the constructor, for example in [ERC20Bridge](../ERC20Bridge.sol):

```solidity
contract ERC20Bridge is
    IERC20Bridge,
    ITeleporterReceiver,
    ReentrancyGuard,
    TeleporterUpgradeable,
    Ownable
{
    ...
    constructor(
        address teleporterRegistryAddress
    ) TeleporterUpgradeable(teleporterRegistryAddress) {
        currentBlockchainID = IWarpMessenger(WARP_PRECOMPILE_ADDRESS)
            .getBlockchainID();
    }
    ...
}
```

The `TeleporterUpgradeable` contract saves the Teleporter registry in a state variable used by the inheriting contract, and initializes a `minTeleporterVersion` to the highest number `TeleporterMessenger` version registered in `TeleporterRegistry`. The `onlyAllowedTeleporter` modifier ensures that `msg.sender` is a `TeleporterMessenger` contract with a version greater than or equal to `minTeleporterVersion`. This modifier is used to restrict access to functions that should only be called by a `TeleporterMessenger` contract, i.e. `ITeleporterReceiver.receiveTeleporterMessage`. This is to support the case where a dapp wants to upgrade to a new version of the `TeleporterMessenger` contract, but still wants to be able to receive messages from the old Teleporter contract.

Every derived contract of `TeleporterUpgradeable` must implement `TeleporterUpgradeable.updateMinTeleporterVersion`, which updates the `minTeleporterVersion` used by the `onlyAllowedTeleporter` modifier and emits the `MinTeleporterVersionUpdated` event. The `updateMinTeleporterVersion` function should be called by the dapp when it completes delivery of messages from the old Teleporter contract, and now wants to update the `minTeleporterVersion` to only allow the new Teleporter version.

To prevent anyone from calling the dapp's `updateMinTeleporterVersion`, which would disallow messages from old Teleporter versions from being received, this function should be safeguarded with access controls. For example, [TeleporterOwnerUpgrade](./TeleporterOwnerUpgradeable.sol) is a contract that inherits `TeleporterUpgradeable` and restricts `updateMinTeleporterVersion` calls to the owner of the contract. The [ERC20Bridge](../../CrossChainApplications/ERC20Bridge/ERC20Bridge.sol) contract is an example of inheriting `TeleporterOwnerUpgradeable`.

```solidity
    function updateMinTeleporterVersion() external override onlyOwner {
        uint256 oldMinTeleporterVersion = minTeleporterVersion;
        minTeleporterVersion = teleporterRegistry.latestVersion();
        emit MinTeleporterVersionUpdated(
            oldMinTeleporterVersion,
            minTeleporterVersion
        );
    }
```

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
