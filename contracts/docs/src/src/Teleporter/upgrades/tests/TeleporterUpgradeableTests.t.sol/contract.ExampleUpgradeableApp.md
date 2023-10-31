# ExampleUpgradeableApp
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/upgrades/tests/TeleporterUpgradeableTests.t.sol)

**Inherits:**
[TeleporterUpgradeable](/src/Teleporter/upgrades/TeleporterUpgradeable.sol/abstract.TeleporterUpgradeable.md)


## Functions
### constructor


```solidity
constructor(address teleporterRegistryAddress) TeleporterUpgradeable(teleporterRegistryAddress);
```

### updateMinTeleporterVersion


```solidity
function updateMinTeleporterVersion() external override;
```

### teleporterCall


```solidity
function teleporterCall() public onlyAllowedTeleporter;
```

