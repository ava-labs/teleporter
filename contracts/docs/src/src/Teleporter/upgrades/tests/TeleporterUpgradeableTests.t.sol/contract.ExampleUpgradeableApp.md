# ExampleUpgradeableApp
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/upgrades/tests/TeleporterUpgradeableTests.t.sol)

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

