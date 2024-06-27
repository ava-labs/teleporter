// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TokenHome} from "./TokenHome.sol";
import {IERC20TokenHome} from "./interfaces/IERC20TokenHome.sol";
import {IERC20SendAndCallReceiver} from "../interfaces/IERC20SendAndCallReceiver.sol";
import {
    SendTokensInput, SendAndCallInput, SingleHopCallMessage
} from "../interfaces/ITokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {SafeERC20TransferFrom} from "../utils/SafeERC20TransferFrom.sol";
import {CallUtils} from "../utils/CallUtils.sol";

/**
 * @title ERC20TokenHome
 * @notice An {IERC20TokenHome} implementation that locks a specified ERC20 token to be sent to
 * TokenRemote instances on other chains.
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
contract ERC20TokenHome is IERC20TokenHome, TokenHome {
    using SafeERC20 for IERC20;

    /// @notice The ERC20 token this home contract bridges to TokenRemote instances.
    IERC20 public immutable token;

    /**
     * @notice Initializes the token TokenHome instance to send ERC20 tokens to TokenRemote instances on other chains.
     * @param teleporterRegistryAddress The current blockchain ID's Teleporter registry
     * address. See here for details: https://github.com/ava-labs/teleporter/tree/main/contracts/src/Teleporter/upgrades
     * @param teleporterManager Address that manages this contract's integration with the
     * Teleporter registry and Teleporter versions.
     * @param tokenAddress_ The ERC20 token contract address to be bridged by the home.
     * @param tokenDecimals_ The number of decimals for the ERC20 token
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address tokenAddress_,
        uint8 tokenDecimals_
    ) TokenHome(teleporterRegistryAddress, teleporterManager, tokenAddress_, tokenDecimals_) {
        token = IERC20(tokenAddress);
    }

    /**
     * @dev See {IERC20TokenBridge-send}
     */
    function send(SendTokensInput calldata input, uint256 amount) external {
        _send(input, amount);
    }

    /**
     * @dev See {IERC20TokenBridge-sendAndCall}
     */
    function sendAndCall(SendAndCallInput calldata input, uint256 amount) external {
        _sendAndCall({
            sourceBlockchainID: blockchainID,
            originBridgeAddress: address(this),
            originSenderAddress: _msgSender(),
            input: input,
            amount: amount
        });
    }

    /**
     * @dev See {IERC20TokenHome-addCollateral}
     */
    function addCollateral(
        bytes32 remoteBlockchainID,
        address remoteBridgeAddress,
        uint256 amount
    ) external {
        _addCollateral(remoteBlockchainID, remoteBridgeAddress, amount);
    }

    /**
     * @dev See {TokenHome-_deposit}
     */
    function _deposit(uint256 amount) internal virtual override returns (uint256) {
        return SafeERC20TransferFrom.safeTransferFrom(token, _msgSender(), amount);
    }

    /**
     * @dev See {TokenHome-_withdraw}
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        emit TokensWithdrawn(recipient, amount);
        token.safeTransfer(recipient, amount);
    }

    /**
     * @dev See {TokenHome-_handleSendAndCall}
     *
     * Approves the recipient contract to spend the amount of tokens from this contract,
     * and calls {IERC20SendAndCallReceiver-receiveTokens} on the recipient contract.
     * If the call fails or doesn't spend all of the tokens, the remaining amount is
     * sent to the fallback recipient.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual override {
        // Approve the recipient contract to spend the amount from the collateral.
        SafeERC20.safeIncreaseAllowance(token, message.recipientContract, amount);

        // Encode the call to {IERC20SendAndCallReceiver-receiveTokens}
        bytes memory payload = abi.encodeCall(
            IERC20SendAndCallReceiver.receiveTokens,
            (
                message.sourceBlockchainID,
                message.originBridgeAddress,
                message.originSenderAddress,
                address(token),
                amount,
                message.recipientPayload
            )
        );

        // Call the recipient contract with the given payload and gas amount.
        bool success = CallUtils._callWithExactGas(
            message.recipientGasLimit, message.recipientContract, payload
        );

        uint256 remainingAllowance = token.allowance(address(this), message.recipientContract);

        // Reset the recipient contract allowance to 0.
        // Use of {safeApprove} is okay to reset the allowance to 0.
        SafeERC20.safeApprove(token, message.recipientContract, 0);

        if (success) {
            emit CallSucceeded(message.recipientContract, amount);
        } else {
            emit CallFailed(message.recipientContract, amount);
        }

        // Transfer any remaining allowance to the fallback recipient. This will be the
        // full amount if the call failed.
        if (remainingAllowance > 0) {
            token.safeTransfer(message.fallbackRecipient, remainingAllowance);
        }
    }
}
