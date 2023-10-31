# TeleporterRegistry
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/upgrades/TeleporterRegistry.sol)

**Inherits:**
[WarpProtocolRegistry](/src/WarpProtocolRegistry.sol/abstract.WarpProtocolRegistry.md)

*TeleporterRegistry contract is a {WarpProtocolRegistry} and provides an upgrade
mechanism for {ITeleporterMessenger} contracts.*


## Functions
### constructor


```solidity
constructor(ProtocolRegistryEntry[] memory initialEntries) WarpProtocolRegistry(initialEntries);
```

### getTeleporterFromVersion

*Gets the {ITeleporterMessenger} contract of the given `version`.
Requirements:
- `version` must be a valid version, i.e. greater than 0 and not greater than the latest version.*


```solidity
function getTeleporterFromVersion(uint256 version) external view returns (ITeleporterMessenger);
```

### getLatestTeleporter

*Gets the latest {ITeleporterMessenger} contract.*


```solidity
function getLatestTeleporter() external view returns (ITeleporterMessenger);
```

