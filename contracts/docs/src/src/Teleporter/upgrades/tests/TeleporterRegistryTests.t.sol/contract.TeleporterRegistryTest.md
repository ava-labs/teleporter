# TeleporterRegistryTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/upgrades/tests/TeleporterRegistryTests.t.sol)

**Inherits:**
Test


## State Variables
### teleporterRegistry

```solidity
TeleporterRegistry public teleporterRegistry;
```


### teleporterAddress

```solidity
address public teleporterAddress;
```


### MOCK_BLOCK_CHAIN_ID

```solidity
bytes32 public constant MOCK_BLOCK_CHAIN_ID = bytes32(uint256(123456));
```


### WARP_PRECOMPILE_ADDRESS

```solidity
address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
```


## Functions
### setUp


```solidity
function setUp() public;
```

### testAddProtocolVersionBasic


```solidity
function testAddProtocolVersionBasic() public;
```

### testAddNonContractAddress


```solidity
function testAddNonContractAddress() public;
```

### testAddOldVersion


```solidity
function testAddOldVersion() public;
```

### testAddExistingVersion


```solidity
function testAddExistingVersion() public;
```

### testAddZeroProtocolAddress


```solidity
function testAddZeroProtocolAddress() public;
```

### testAddZeroVersion


```solidity
function testAddZeroVersion() public;
```

### testGetAddressFromVersion


```solidity
function testGetAddressFromVersion() public;
```

### testGetVersionFromAddress


```solidity
function testGetVersionFromAddress() public;
```

### testInvalidWarpMessage


```solidity
function testInvalidWarpMessage() public;
```

### _addProtocolVersion


```solidity
function _addProtocolVersion(TeleporterRegistry registry) internal;
```

### _createWarpOutofBandMessage


```solidity
function _createWarpOutofBandMessage(uint256 version, address protocolAddress, address registryAddress)
    internal
    view
    returns (WarpMessage memory);
```

### _formatErrorMessage


```solidity
function _formatErrorMessage(string memory errorMessage) private pure returns (bytes memory);
```

## Events
### AddProtocolVersion

```solidity
event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress);
```

