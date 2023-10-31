# ExampleUpgradeableApp
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/upgrades/tests/TeleporterUpgradeableTests.t.sol)

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

