# TeleporterUpgradeable
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/upgrades/TeleporterUpgradeable.sol)

*TeleporterUpgradeable provides upgrade utility for applications built on top
of the Teleporter protocol by integrating with the {TeleporterRegistry}.
This contract is intended to be inherited by other contracts that wish to use the
upgrade mechanism. It provides a modifier that restricts access to only Teleporter
versions that are greater than or equal to `minTeleporterVersion`.*


## State Variables
### teleporterRegistry

```solidity
TeleporterRegistry public immutable teleporterRegistry;
```


### minTeleporterVersion

```solidity
uint256 public minTeleporterVersion;
```


## Functions
### onlyAllowedTeleporter

*Throws if called by a `msg.sender` that is not an allowed Teleporter version.
Checks that `msg.sender` matches a Teleporter version greater than or equal to `minTeleporterVersion`.*


```solidity
modifier onlyAllowedTeleporter();
```

### constructor

*Initializes the {TeleporterUpgradeable} contract by getting `teleporterRegistry`
instance and setting `_minTeleporterVersion`.*


```solidity
constructor(address teleporterRegistryAddress);
```

### updateMinTeleporterVersion

*This is a virtual function that should be overridden to update the `minTeleporterVersion`
allowed for modifier `onlyAllowedTeleporter`, and emit {MinTeleporterVersionUpdated} event after.
Note: To prevent anyone from being able to call this function, which would disallow messages
from old Teleporter versions from being received, this function should be safeguarded with access
controls. For example, if the derived contract has an owner/admin, only they can call this function.*


```solidity
function updateMinTeleporterVersion() external virtual;
```

## Events
### MinTeleporterVersionUpdated

```solidity
event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion);
```

