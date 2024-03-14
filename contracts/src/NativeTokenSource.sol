// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenSource} from "./TeleporterTokenSource.sol";
import {INativeTokenBridge} from "./interfaces/INativeTokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {SendTokensInput} from "./interfaces/ITeleporterTokenBridge.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title NativeTokenSource
 * @notice This contract is an {INativeTokenBridge} that sends native tokens to another chain's
 * {ITeleporterTokenBridge} instance, and gets represented by the tokens of that destination
 * token bridge instance.
 */
contract NativeTokenSource is INativeTokenBridge, TeleporterTokenSource {
    using SafeERC20 for IERC20;

    /**
     * @notice Initializes this source token bridge instance to send
     * tokens to the specified destination chain and token bridge instance.
     *
     * Teleporter fees are paid by the same token that is being bridged.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address feeTokenAddress
    ) TeleporterTokenSource(teleporterRegistryAddress, teleporterManager, feeTokenAddress) {}

    /**
     * @dev See {IERC20Bridge-send}
     */
    function send(SendTokensInput calldata input) external payable {
        _send(input, msg.value, false);
    }

    /**
     * @dev See {TeleportTokenSource-_deposit}
     */
    function _deposit(uint256 amount) internal virtual override returns (uint256) {
        // TODO: Deposit native token for fee token and transfer to bridge
        // Return the amount after fee is transferred to bridge
        return amount;
    }

    /**
     * @dev See {TeleportTokenSource-_withdraw}
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        // TODO: Withdraw erc20 token amount into native tokens
        // Transfer the native tokens to the recipient
        payable(recipient).transfer(amount);
    }
}
