# ReentrancyGuardsTests
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/tests/ReentrancyGuardsTests.t.sol)

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

