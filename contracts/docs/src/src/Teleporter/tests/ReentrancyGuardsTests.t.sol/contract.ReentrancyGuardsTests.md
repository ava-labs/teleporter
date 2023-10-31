# ReentrancyGuardsTests
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/tests/ReentrancyGuardsTests.t.sol)

**Inherits:**
Test


## State Variables
### _sampleMessenger

```solidity
SampleMessenger internal _sampleMessenger;
```


## Functions
### setUp


```solidity
function setUp() public virtual;
```

### testConsecutiveSendSuccess


```solidity
function testConsecutiveSendSuccess() public;
```

### testConsecutiveReceiveSuccess


```solidity
function testConsecutiveReceiveSuccess() public;
```

### testRecursiveSendFails


```solidity
function testRecursiveSendFails() public;
```

### testRecursiveReceiveFails


```solidity
function testRecursiveReceiveFails() public;
```

### testSendCallsReceiveSuccess


```solidity
function testSendCallsReceiveSuccess() public;
```

### testReceiveCallsSendSuccess


```solidity
function testReceiveCallsSendSuccess() public;
```

### testRecursiveDirectSendFails


```solidity
function testRecursiveDirectSendFails() public;
```

### testRecursiveDirectReceiveFails


```solidity
function testRecursiveDirectReceiveFails() public;
```

