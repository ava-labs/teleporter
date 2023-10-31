# HandleInitialMessageExecutionTest
[Git Source](https://github.com/ava-labs/teleporter/blob/dde09fbf56cc395da6bfd76c7f894a3cf5b2cd9e/src/Teleporter/tests/HandleInitialMessageExecutionTests.t.sol)

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

