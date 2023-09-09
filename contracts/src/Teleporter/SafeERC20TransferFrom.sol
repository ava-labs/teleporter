// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @dev Provides a wrapper used for calling an ERC20 transferFrom method
 * to receive tokens to a contract from msg.sender.
 *
 * Checks the balance of the recipient before and after the call to transferFrom, and
 * returns balance increase. Designed for safely handling ERC20 "fee on transfer" and "burn on transfer" implementations.
 *
 * Note: A reentrancy guard must always be used when calling token.safeTransferFrom in order to
 * prevent against possible "before-after" pattern vulnerabilities.
 */
library SafeERC20TransferFrom {
    using SafeERC20 for IERC20;

    // Errors
    error BalanceUnchanged();

    function safeTransferFrom(
        IERC20 erc20,
        uint256 amount
    ) internal returns (uint256) {
        uint256 balanceBefore = erc20.balanceOf(address(this));
        erc20.safeTransferFrom(msg.sender, address(this), amount);
        uint256 balanceAfter = erc20.balanceOf(address(this));

        if (balanceAfter <= balanceBefore) {
            revert BalanceUnchanged();
        }

        return balanceAfter - balanceBefore;
    }
}
