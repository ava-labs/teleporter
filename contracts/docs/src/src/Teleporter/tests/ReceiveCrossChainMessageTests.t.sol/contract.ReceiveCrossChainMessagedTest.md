# ReceiveCrossChainMessagedTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/tests/ReceiveCrossChainMessageTests.t.sol)

**Inherits:**
[TeleporterMessengerTest](/src/Teleporter/tests/TeleporterMessengerTest.t.sol/contract.TeleporterMessengerTest.md)


## State Variables
### DEFAULT_MESSAGE_PAYLOAD

```solidity
bytes public constant DEFAULT_MESSAGE_PAYLOAD =
    hex"cafebabe11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff11223344556677889900aabbccddeeff11223344556677889900aabbccddeeffdeadbeef";
```


## Functions
### setUp


```solidity
function setUp() public virtual override;
```

### testSuccess


```solidity
function testSuccess() public;
```

### testNoValidMessage


```solidity
function testNoValidMessage() public;
```

### testInvalidOriginSenderAddress


```solidity
function testInvalidOriginSenderAddress() public;
```

### testInvalidDestinationChainID


```solidity
function testInvalidDestinationChainID() public;
```

### testInvalidDestinationAddress


```solidity
function testInvalidDestinationAddress() public;
```

### testInvalidRelayerAddress


```solidity
function testInvalidRelayerAddress() public;
```

### testMessageAlreadyReceived


```solidity
function testMessageAlreadyReceived() public;
```

### testUnauthorizedRelayer


```solidity
function testUnauthorizedRelayer() public;
```

### testMessageSentToEOADoesNotExecute


```solidity
function testMessageSentToEOADoesNotExecute() public;
```

