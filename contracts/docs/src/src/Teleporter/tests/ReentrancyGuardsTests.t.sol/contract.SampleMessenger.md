# SampleMessenger
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/tests/ReentrancyGuardsTests.t.sol)

**Inherits:**
[ReentrancyGuards](/src/Teleporter/ReentrancyGuards.sol/abstract.ReentrancyGuards.md)


## State Variables
### nonce

```solidity
uint256 public nonce;
```


## Functions
### constructor


```solidity
constructor();
```

### sendAndCall


```solidity
function sendAndCall(bytes memory message) public senderNonReentrant;
```

### sendMessage


```solidity
function sendMessage() public senderNonReentrant;
```

### sendRecursive


```solidity
function sendRecursive() public senderNonReentrant;
```

### receiveAndCall


```solidity
function receiveAndCall(bytes memory message) public receiverNonReentrant;
```

### receiveMessage


```solidity
function receiveMessage() public receiverNonReentrant;
```

### receiveRecursive


```solidity
function receiveRecursive() public receiverNonReentrant;
```

