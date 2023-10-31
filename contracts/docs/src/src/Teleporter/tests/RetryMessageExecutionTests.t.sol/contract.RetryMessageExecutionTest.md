# RetryMessageExecutionTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/tests/RetryMessageExecutionTests.t.sol)

**Inherits:**
[TeleporterMessengerTest](/src/Teleporter/tests/TeleporterMessengerTest.t.sol/contract.TeleporterMessengerTest.md)


## State Variables
### destinationContract

```solidity
FlakyMessageReceiver public destinationContract;
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

### testExecutionFailsAgain


```solidity
function testExecutionFailsAgain() public;
```

### testMessageHashNotFound


```solidity
function testMessageHashNotFound() public;
```

### testInvalidMessageHash


```solidity
function testInvalidMessageHash() public;
```

### testCanNotRetryAgainAfterSuccess


```solidity
function testCanNotRetryAgainAfterSuccess() public;
```

### testCanNotReceiveMessageWhileRetrying


```solidity
function testCanNotReceiveMessageWhileRetrying() public;
```

### testEOAFailsThenRetrySucceeds


```solidity
function testEOAFailsThenRetrySucceeds() public;
```

### testEOAFailsAgainOnRetry


```solidity
function testEOAFailsAgainOnRetry() public;
```

### _receiveFailedMessage


```solidity
function _receiveFailedMessage(bool retryReceive) internal returns (bytes32, TeleporterMessage memory, string memory);
```

### _successfullyRetryMessage


```solidity
function _successfullyRetryMessage() internal returns (bytes32, TeleporterMessage memory);
```

