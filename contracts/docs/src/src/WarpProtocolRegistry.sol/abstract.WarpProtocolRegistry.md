# WarpProtocolRegistry
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/WarpProtocolRegistry.sol)

*Implementation of an abstract `WarpProtocolRegistry` contract.
This implementation is a contract that can be used as a base contract for protocols that are
built on top of Warp. It allows the protocol to be upgraded through a Warp out-of-band message.*


## State Variables
### VALIDATORS_SOURCE_ADDRESS

```solidity
address public constant VALIDATORS_SOURCE_ADDRESS = address(0);
```


### WARP_MESSENGER

```solidity
WarpMessenger public constant WARP_MESSENGER = WarpMessenger(0x0200000000000000000000000000000000000005);
```


### _blockchainID

```solidity
bytes32 internal immutable _blockchainID;
```


### _latestVersion

```solidity
uint256 internal _latestVersion;
```


### _versionToAddress

```solidity
mapping(uint256 => address) internal _versionToAddress;
```


### _addressToVersion

```solidity
mapping(address => uint256) internal _addressToVersion;
```


## Functions
### constructor

*Initializes the contract by setting `_blockchainID` and `_latestVersion`.
Also adds the initial protocol versions to the registry.*


```solidity
constructor(ProtocolRegistryEntry[] memory initialEntries);
```

### addProtocolVersion

*Gets and verifies a warp out-of-band message, and adds the new protocol version
address to the registry.
If a version is greater than the current latest version, it will be set as the latest version.
If a version is less than the current latest version, it is added to the registry, but
doesn't change the latest version.
Emits a {AddProtocolVersion} event when successful.
Requirements:
- a valid Warp out-of-band message must be provided.
- source chain ID must be the same as the blockchain ID of the registry.
- origin sender address must be the same as the `VALIDATORS_SOURCE_ADDRESS`.
- destination chain ID must be the same as the blockchain ID of the registry.
- destination address must be the same as the address of the registry.
- version must not be zero.
- version must not already be registered.
- protocol address must not be zero address.*


```solidity
function addProtocolVersion(uint32 messageIndex) external virtual;
```

### getAddressFromVersion

*Gets the address of a protocol version.
Requirements:
- the version must be a valid version.*


```solidity
function getAddressFromVersion(uint256 version) external view virtual returns (address);
```

### getVersionFromAddress

*Gets the version of the given `protocolAddress`.
If `protocolAddress` is not a registered protocol address, returns 0, which is an invalid version.*


```solidity
function getVersionFromAddress(address protocolAddress) external view virtual returns (uint256);
```

### getLatestVersion

*Gets the latest protocol version.
If the registry has no versions, we return 0, which is an invalid version.*


```solidity
function getLatestVersion() external view virtual returns (uint256);
```

### _addToRegistry

*Adds the new protocol version address to the registry.
Updates latest version if the version is greater than the current latest version.
Emits a {AddProtocolVersion} event when successful.
Note: `protocolAddress` doesn't have to be a contract address, this is primarily
to support the case we want to register a new protocol address meant for a security patch
before the contract is deployed, to prevent the vulnerabilitiy from being exposed before registry update.
Requirements:
- `version` is not zero
- `version` is not already registered
- `protocolAddress` is not zero address*


```solidity
function _addToRegistry(ProtocolRegistryEntry memory entry) internal virtual;
```

### _getAddressFromVersion

*Gets the address of a protocol version.
Requirements:
- `version` must be a valid version, i.e. greater than 0 and not greater than the latest version.*


```solidity
function _getAddressFromVersion(uint256 version) internal view virtual returns (address);
```

## Events
### AddProtocolVersion
*Emitted when a new protocol version is added to the registry.*


```solidity
event AddProtocolVersion(uint256 indexed version, address indexed protocolAddress);
```

