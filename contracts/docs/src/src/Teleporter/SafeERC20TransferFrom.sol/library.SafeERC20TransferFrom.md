# SafeERC20TransferFrom
[Git Source](https://github.com/ava-labs/teleporter/blob/4e46f28c075e9bfc858fb8bbe266f5b4cb45a0be/src/Teleporter/SafeERC20TransferFrom.sol)

*Provides a wrapper used for calling an ERC20 transferFrom method
to receive tokens to a contract from msg.sender.
Checks the balance of the recipient before and after the call to transferFrom, and
returns balance increase. Designed for safely handling ERC20 "fee on transfer" and "burn on transfer" implementations.
Note: A reentrancy guard must always be used when calling token.safeTransferFrom in order to
prevent against possible "before-after" pattern vulnerabilities.*


## Functions
### safeTransferFrom


```solidity
function safeTransferFrom(IERC20 erc20, uint256 amount) internal returns (uint256);
```

