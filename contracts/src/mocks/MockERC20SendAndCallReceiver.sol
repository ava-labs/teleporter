// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20SendAndCallReceiver} from "../interfaces/IERC20SendAndCallReceiver.sol";
import {SafeERC20TransferFrom} from "@teleporter/SafeERC20TransferFrom.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice This is mock implementation of {receiveTokens} to be used in tests.
 * This contract DOES NOT provide a mechanism for accessing the tokens transfered to it.
 * Real implementations must ensure that tokens are properly handled and not incorrectly locked.
 */
contract MockERC20SendAndCallReceiver is IERC20SendAndCallReceiver {
    using SafeERC20 for IERC20;

    /**
     * @dev Emitted when receiveTokens is called.
     */
    event TokensReceived(address token, uint256 amount, bytes payload);

    /**
     * @dev See {IERC20SendAndCallReceiver-receiveTokens}
     */
    function receiveTokens(address token, uint256 amount, bytes calldata payload) external {
        emit TokensReceived(token, amount, payload);

        require(payload.length > 0, "MockERC20SendAndCallReceiver: empty payload");

        SafeERC20TransferFrom.safeTransferFrom(IERC20(token), amount);
    }
}
