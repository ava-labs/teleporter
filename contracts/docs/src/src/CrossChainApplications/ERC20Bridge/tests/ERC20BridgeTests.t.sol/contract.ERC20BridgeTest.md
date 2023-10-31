# ERC20BridgeTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/CrossChainApplications/ERC20Bridge/tests/ERC20BridgeTests.t.sol)

**Inherits:**
Test


## State Variables
### MOCK_TELEPORTER_MESSENGER_ADDRESS

```solidity
address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS = 0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
```


### MOCK_TELEPORTER_REGISTRY_ADDRESS

```solidity
address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS = 0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
```


### WARP_PRECOMPILE_ADDRESS

```solidity
address public constant WARP_PRECOMPILE_ADDRESS = address(0x0200000000000000000000000000000000000005);
```


### _MOCK_BLOCKCHAIN_ID

```solidity
bytes32 private constant _MOCK_BLOCKCHAIN_ID = bytes32(uint256(123456));
```


### _DEFAULT_OTHER_CHAIN_ID

```solidity
bytes32 private constant _DEFAULT_OTHER_CHAIN_ID =
    bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
```


### _DEFAULT_OTHER_BRIDGE_ADDRESS

```solidity
address private constant _DEFAULT_OTHER_BRIDGE_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
```


### _DEFAULT_TOKEN_NAME

```solidity
string private constant _DEFAULT_TOKEN_NAME = "Test Token";
```


### _DEFAULT_SYMBOL

```solidity
string private constant _DEFAULT_SYMBOL = "TSTTK";
```


### _DEFAULT_DECIMALS

```solidity
uint8 private constant _DEFAULT_DECIMALS = 18;
```


### _DEFAULT_RECIPIENT

```solidity
address private constant _DEFAULT_RECIPIENT = 0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;
```


### erc20Bridge

```solidity
ERC20Bridge public erc20Bridge;
```


### mockERC20

```solidity
UnitTestMockERC20 public mockERC20;
```


## Functions
### setUp


```solidity
function setUp() public virtual;
```

### testSameChainID


```solidity
function testSameChainID() public;
```

### testInvalidFeeAmountsNativeTransfer


```solidity
function testInvalidFeeAmountsNativeTransfer() public;
```

### testInvalidFeeAmountsWrappedTransfer


```solidity
function testInvalidFeeAmountsWrappedTransfer() public;
```

### testNativeTokenTransferFailure


```solidity
function testNativeTokenTransferFailure() public;
```

### testBridgeNativeTokensNoFee


```solidity
function testBridgeNativeTokensNoFee() public;
```

### testFeeApprovalFails


```solidity
function testFeeApprovalFails() public;
```

### testBridgeNativeTokensWithFee


```solidity
function testBridgeNativeTokensWithFee() public;
```

### testBridgeNativeFeeOnTransferTokens


```solidity
function testBridgeNativeFeeOnTransferTokens() public;
```

### testBridgeNativeFeeOnTransferAmountTooHigh


```solidity
function testBridgeNativeFeeOnTransferAmountTooHigh() public;
```

### testNewBridgeTokenMint


```solidity
function testNewBridgeTokenMint() public;
```

### testMintExistingBridgeToken


```solidity
function testMintExistingBridgeToken() public;
```

### testZeroTeleporterRegistryAddress


```solidity
function testZeroTeleporterRegistryAddress() public;
```

### testUpdateMinTeleporterVersion


```solidity
function testUpdateMinTeleporterVersion() public;
```

### _initMockTeleporterRegistry


```solidity
function _initMockTeleporterRegistry() internal;
```

### _setUpExpectedTransferFromCall


```solidity
function _setUpExpectedTransferFromCall(address tokenContract, uint256 amount) private;
```

### _submitCreateBridgeToken


```solidity
function _submitCreateBridgeToken(
    bytes32 destinationChainID,
    address destinationBridgeAddress,
    address nativeContractAddress
) private;
```

### _setUpBridgeToken


```solidity
function _setUpBridgeToken(
    bytes32 nativeChainID,
    address nativeBridgeAddress,
    address nativeContractAddress,
    string memory nativeName,
    string memory nativeSymbol,
    uint8 nativeDecimals,
    uint8 contractNonce
) private returns (address);
```

### _setUpMockERC20ContractValues


```solidity
function _setUpMockERC20ContractValues(address tokenContract) private;
```

### _deriveExpectedContractAddress


```solidity
function _deriveExpectedContractAddress(address creator, uint8 nonce) private pure returns (address);
```

### _formatERC20BridgeErrorMessage


```solidity
function _formatERC20BridgeErrorMessage(string memory errorMessage) private pure returns (bytes memory);
```

## Events
### BridgeTokens

```solidity
event BridgeTokens(
    address indexed tokenContractAddress,
    bytes32 indexed destinationChainID,
    uint256 indexed teleporterMessageID,
    address destinationBridgeAddress,
    address recipient,
    uint256 amount
);
```

### SubmitCreateBridgeToken

```solidity
event SubmitCreateBridgeToken(
    bytes32 indexed destinationChainID,
    address indexed destinationBridgeAddress,
    address indexed nativeContractAddress,
    uint256 teleporterMessageID
);
```

### CreateBridgeToken

```solidity
event CreateBridgeToken(
    bytes32 indexed nativeChainID,
    address indexed nativeBridgeAddress,
    address indexed nativeContractAddress,
    address bridgeTokenAddress
);
```

### MintBridgeTokens

```solidity
event MintBridgeTokens(address indexed contractAddress, address recipient, uint256 amount);
```

### MinTeleporterVersionUpdated

```solidity
event MinTeleporterVersionUpdated(uint256 indexed oldMinTeleporterVersion, uint256 indexed newMinTeleporterVersion);
```

