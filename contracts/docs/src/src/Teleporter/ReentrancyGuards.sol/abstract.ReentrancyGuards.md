# ReentrancyGuards
[Git Source](https://github.com/ava-labs/teleporter/blob/cadc1420fd95195b094eea855b7496cc71b5be2a/src/Teleporter/ReentrancyGuards.sol)

*Abstract contract that helps implement reentrancy guards between functions for sending and receiving.
Consecutive calls for sending functions should work together, same for receive functions, but recursive calls
should be detected as a reentrancy and revert.
Calls between send and receive functions should also be allowed, but not in the case it ends up being a recursive
send or receive call. For example the following should fail: send -> receive -> send.*


## State Variables
### _NOT_ENTERED

```solidity
uint256 internal constant _NOT_ENTERED = 1;
```


### _ENTERED

```solidity
uint256 internal constant _ENTERED = 2;
```


### _sendEntered

```solidity
uint256 internal _sendEntered;
```


### _receiveEntered

```solidity
uint256 internal _receiveEntered;
```


## Functions
### senderNonReentrant


```solidity
modifier senderNonReentrant();
```

### receiverNonReentrant


```solidity
modifier receiverNonReentrant();
```

### constructor


```solidity
constructor();
```

