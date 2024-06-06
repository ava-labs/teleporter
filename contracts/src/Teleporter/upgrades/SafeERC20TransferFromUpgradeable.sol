// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20Upgradeable} from
    "@openzeppelin/contracts-upgradeable@4.9.6/token/ERC20/ERC20Upgradeable.sol";

import {SafeERC20Upgradeable} from
    "@openzeppelin/contracts-upgradeable@4.9.6/token/ERC20/utils/SafeERC20Upgradeable.sol";
/**
 * @dev Provides a wrapper used for calling an ERC20 transferFrom method
 * to receive tokens to a contract from msg.sender.
 *
 * Checks the balance of the contract using the library before and after the call to safeTransferFrom, and
 * returns balance increase. Designed for safely handling ERC20 "fee on transfer" and "burn on transfer" implementations.
 *
 * Note: A reentrancy guard must always be used when calling token.safeTransferFrom in order to
 * prevent against possible "before-after" pattern vulnerabilities.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */

library SafeERC20TransferFromUpgradeable {
    using SafeERC20Upgradeable for IERC20Upgradeable;

    /**
     * @dev Checks the balance of the contract before and after the call to safeTransferFrom, and returns the balance
     * increase. Designed for safely handling ERC20 "fee on transfer" and "burn on transfer" implementations.
     */
    // solhint-disable private-vars-leading-underscore
    function safeTransferFrom(IERC20Upgradeable erc20, uint256 amount) internal returns (uint256) {
        uint256 balanceBefore = erc20.balanceOf(address(this));
        erc20.safeTransferFrom(msg.sender, address(this), amount);
        uint256 balanceAfter = erc20.balanceOf(address(this));

        require(balanceAfter > balanceBefore, "SafeERC20TransferFrom: balance not increased");

        return balanceAfter - balanceBefore;
    }
    // solhint-enable private-vars-leading-underscore
}
