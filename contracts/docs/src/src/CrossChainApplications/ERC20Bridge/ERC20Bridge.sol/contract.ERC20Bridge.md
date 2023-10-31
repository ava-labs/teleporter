# ERC20Bridge
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/CrossChainApplications/ERC20Bridge/ERC20Bridge.sol)

**Inherits:**
[IERC20Bridge](/src/CrossChainApplications/ERC20Bridge/IERC20Bridge.sol/interface.IERC20Bridge.md), [ITeleporterReceiver](/src/Teleporter/ITeleporterReceiver.sol/interface.ITeleporterReceiver.md), ReentrancyGuard, [TeleporterUpgradeable](/src/Teleporter/upgrades/TeleporterUpgradeable.sol/abstract.TeleporterUpgradeable.md), Ownable

*Implementation of the {IERC20Bridge} interface.
This implementation uses the {BridgeToken} contract to represent tokens on this chain, and uses
{ITeleporterMessenger} to send and receive messages to other chains.*


## State Variables
### WARP_PRECOMPILE_ADDRESS

```solidity
address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
```


### currentChainID

```solidity
bytes32 public immutable currentChainID;
```


### submittedBridgeTokenCreations

```solidity
mapping(bytes32 => mapping(address => mapping(address => bool))) public submittedBridgeTokenCreations;
```


### bridgedBalances

```solidity
mapping(bytes32 => mapping(address => mapping(address => uint256))) public bridgedBalances;
```


### wrappedTokenContracts

```solidity
mapping(address => bool) public wrappedTokenContracts;
```


### nativeToWrappedTokens

```solidity
mapping(bytes32 => mapping(address => mapping(address => address))) public nativeToWrappedTokens;
```


### CREATE_BRIDGE_TOKENS_REQUIRED_GAS

```solidity
uint256 public constant CREATE_BRIDGE_TOKENS_REQUIRED_GAS = 2_000_000;
```


### MINT_BRIDGE_TOKENS_REQUIRED_GAS

```solidity
uint256 public constant MINT_BRIDGE_TOKENS_REQUIRED_GAS = 200_000;
```


### TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS

```solidity
uint256 public constant TRANSFER_BRIDGE_TOKENS_REQUIRED_GAS = 300_000;
```


## Functions
### constructor

*Initializes the Teleporter Messenger used for sending and receiving messages,
and initializes the current chain ID.*


```solidity
constructor(address teleporterRegistryAddress) TeleporterUpgradeable(teleporterRegistryAddress);
```

### bridgeTokens

*See {IERC20Bridge-bridgeTokens}.
Requirements:
- `destinationChainID` cannot be the same as the current chain ID.
- For wrapped tokens, `totalAmount` must be greater than the sum of the primary and secondary fee amounts.
- For native tokens, `adjustedAmount` after safe transfer must be greater than the primary fee amount.*


```solidity
function bridgeTokens(
    bytes32 destinationChainID,
    address destinationBridgeAddress,
    address tokenContractAddress,
    address recipient,
    uint256 totalAmount,
    uint256 primaryFeeAmount,
    uint256 secondaryFeeAmount
) external nonReentrant;
```

### submitCreateBridgeToken

*See {IERC20Bridge-submitCreateBridgeToken}.
We allow for `submitCreateBridgeToken` to be called multiple times with the same bridge and token
information because a previous message may have been dropped or otherwise selectively not delivered.
If the bridge token already exists on the destination, we are sending a message that will
simply have no effect on the destination.
Emits a {SubmitCreateBridgeToken} event.*


```solidity
function submitCreateBridgeToken(
    bytes32 destinationChainID,
    address destinationBridgeAddress,
    ERC20 nativeToken,
    address messageFeeAsset,
    uint256 messageFeeAmount
) external nonReentrant;
```

### receiveTeleporterMessage

*See {ITeleporterReceiver-receiveTeleporterMessage}.
Receives a Teleporter message and routes to the appropriate internal function call.*


```solidity
function receiveTeleporterMessage(bytes32 nativeChainID, address nativeBridgeAddress, bytes calldata message)
    external
    onlyAllowedTeleporter;
```

### updateMinTeleporterVersion

*See {TeleporterUpgradeable-updateMinTeleporterVersion}
Updates the minimum Teleporter version allowed for receiving on this contract
to the latest version registered in the {TeleporterRegistry}.
Restricted to only owners of the contract.
Emits a {MinTeleporterVersionUpdated} event.*


```solidity
function updateMinTeleporterVersion() external override onlyOwner;
```

### encodeCreateBridgeTokenData

*Encodes the parameters for the Create action to be decoded and executed on the destination.*


```solidity
function encodeCreateBridgeTokenData(
    address nativeContractAddress,
    string memory nativeName,
    string memory nativeSymbol,
    uint8 nativeDecimals
) public pure returns (bytes memory);
```

### encodeMintBridgeTokensData

*Encodes the parameters for the Mint action to be decoded and executed on the destination.*


```solidity
function encodeMintBridgeTokensData(address nativeContractAddress, address recipient, uint256 bridgeAmount)
    public
    pure
    returns (bytes memory);
```

### encodeTransferBridgeTokensData

*Encodes the parameters for the Transfer action to be decoded and executed on the destination.*


```solidity
function encodeTransferBridgeTokensData(
    bytes32 destinationChainID,
    address destinationBridgeAddress,
    address nativeContractAddress,
    address recipient,
    uint256 amount,
    uint256 feeAmount
) public pure returns (bytes memory);
```

### _createBridgeToken

*Teleporter message receiver for creating a new bridge token on this chain.
Emits a {CreateBridgeToken} event.
Note: This function is only called within `receiveTeleporterMessage`, which can only be
called by the Teleporter Messenger.*


```solidity
function _createBridgeToken(
    bytes32 nativeChainID,
    address nativeBridgeAddress,
    address nativeContractAddress,
    string memory nativeName,
    string memory nativeSymbol,
    uint8 nativeDecimals
) private;
```

### _mintBridgeTokens

*Teleporter message receiver for minting of an existing bridge token on this chain.
Emits a {MintBridgeTokens} event.
Note: This function is only called within `receiveTeleporterMessage`, which can only be
called by the Teleporter Messenger.*


```solidity
function _mintBridgeTokens(
    bytes32 nativeChainID,
    address nativeBridgeAddress,
    address nativeContractAddress,
    address recipient,
    uint256 amount
) private nonReentrant;
```

### _transferBridgeTokens

*Teleporter message receiver for handling bridge tokens transfers back from another chain
and optionally routing them to a different third chain.
Note: This function is only called within `receiveTeleporterMessage`, which can only be
called by the Teleporter Messenger.*


```solidity
function _transferBridgeTokens(
    bytes32 sourceChainID,
    address sourceBridgeAddress,
    bytes32 destinationChainID,
    address destinationBridgeAddress,
    address nativeContractAddress,
    address recipient,
    uint256 totalAmount,
    uint256 secondaryFeeAmount
) private nonReentrant;
```

### _processNativeTokenTransfer

*Increments the balance of the native tokens bridged to the specified bridge instance and
sends a Teleporter message to have the destination bridge mint the new tokens. The tokens to be
bridge must already be locked in this contract before calling.
Emits a {BridgeTokens} event.
Requirements:
- `destinationChainID` cannot be the same as the current chain ID.
- can not do nested bridging of wrapped tokens.*


```solidity
function _processNativeTokenTransfer(
    bytes32 destinationChainID,
    address destinationBridgeAddress,
    address nativeContractAddress,
    address recipient,
    uint256 totalAmount,
    uint256 feeAmount
) private;
```

### _processWrappedTokenTransfer

*Processes a wrapped token transfer by burning the tokens and sending a Teleporter message
to the native chain and bridge of the wrapped asset that was burned.
It is the caller's responsibility to ensure that the wrapped token contract is supported by this bridge instance.
Emits a {BridgeTokens} event.*


```solidity
function _processWrappedTokenTransfer(WrappedTokenTransferInfo memory wrappedTransferInfo) private;
```

## Structs
### WrappedTokenTransferInfo

```solidity
struct WrappedTokenTransferInfo {
    bytes32 destinationChainID;
    address destinationBridgeAddress;
    address wrappedContractAddress;
    address recipient;
    uint256 totalAmount;
    uint256 primaryFeeAmount;
    uint256 secondaryFeeAmount;
}
```

