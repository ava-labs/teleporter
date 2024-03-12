// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenSource} from "./TeleporterTokenSource.sol";
import {IERC20Bridge} from "./interfaces/IERC20Bridge.sol";
import {SafeERC20TransferFrom} from "@teleporter/SafeERC20TransferFrom.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {SendTokensInput} from "./interfaces/ITeleporterTokenBridge.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title ERC20Source
 * @notice This contract is an {IERC20Bridge} that sends ERC20 tokens to another chain's
 * {ITeleporterTokenBridge} instance, and gets represented by the tokens of that destination
 * token bridge instance.
 */
contract ERC20Source is IERC20Bridge, TeleporterTokenSource {
    using SafeERC20 for IERC20;

    /// @notice The ERC20 token this source contract bridges to destination instances.
    IERC20 public immutable token;

    /**
     * @notice Initializes this source token bridge instance to send
     * tokens to the specified destination chain and token bridge instance.
     *
     * Teleporter fees are paid by the same token that is being bridged.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address tokenAddress
    ) TeleporterTokenSource(teleporterRegistryAddress, teleporterManager, tokenAddress) {
        token = IERC20(tokenAddress);
    }

    /**
     * @dev See {IERC20Bridge-send}
     */
    function send(SendTokensInput calldata input, uint256 amount) external {
        _send(input, amount);
    }

    /**
     * @dev See {TeleportTokenSource-_deposit}
     */
    function _deposit(uint256 amount) internal virtual override returns (uint256) {
        return SafeERC20TransferFrom.safeTransferFrom(token, amount);
    }

    /**
     * @dev See {TeleportTokenSource-_withdraw}
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        token.safeTransfer(recipient, amount);
    }
}
