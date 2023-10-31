# HandleInitialMessageExecutionTest
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/tests/HandleInitialMessageExecutionTests.t.sol)

**Inherits:**
[TeleporterMessengerTest](/src/Teleporter/tests/TeleporterMessengerTest.t.sol/contract.TeleporterMessengerTest.md)


## State Variables
### destinationContract

```solidity
SampleMessageReceiver public destinationContract;
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

### testInsufficientGasProvided


```solidity
function testInsufficientGasProvided() public;
```

### testCannotReceiveMessageRecursively


```solidity
function testCannotReceiveMessageRecursively() public;
```

### testStoreHashOfFailedMessageExecution


```solidity
function testStoreHashOfFailedMessageExecution() public;
```

