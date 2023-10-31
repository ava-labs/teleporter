# TeleporterUpgradeableTest
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/upgrades/tests/TeleporterUpgradeableTests.t.sol)

**Inherits:**
Test


## State Variables
### MOCK_TELEPORTER_REGISTRY_ADDRESS

```solidity
address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS = 0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
```


### MOCK_TELEPORTER_MESSENGER_ADDRESS

```solidity
address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS = 0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
```


## Functions
### setUp


```solidity
function setUp() public;
```

### testInvalidRegistryAddress


```solidity
function testInvalidRegistryAddress() public;
```

### testOnlyAllowedTeleporter


```solidity
function testOnlyAllowedTeleporter() public;
```

### testUpdateMinTeleporterVersion


```solidity
function testUpdateMinTeleporterVersion() public;
```

### _formatTeleporterUpgradeableErrorMessage


```solidity
function _formatTeleporterUpgradeableErrorMessage(string memory errorMessage) private pure returns (bytes memory);
```

## Events
### MinTeleporterVersionUpdated

```solidity
event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion);
```

