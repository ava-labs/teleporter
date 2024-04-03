// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenDestination} from "./TeleporterTokenDestination.sol";
import {IERC20Bridge} from "./interfaces/IERC20Bridge.sol";
import {IERC20SendAndCallReceiver} from "./interfaces/IERC20SendAndCallReceiver.sol";
import {SafeERC20TransferFrom} from "@teleporter/SafeERC20TransferFrom.sol";
import {IERC20, ERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    SingleHopCallMessage
} from "./interfaces/ITeleporterTokenBridge.sol";
import {GasUtils} from "./utils/GasUtils.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title ERC20Destination
 * @notice This contract is an {IERC20Bridge} that receives tokens from another chain's
 * {ITeleporterTokenBridge} instance, and represents the received tokens with an ERC20 token
 * on this destination chain.
 */
contract ERC20Destination is IERC20Bridge, TeleporterTokenDestination, ERC20 {
    using SafeERC20 for IERC20;
    using GasUtils for *;

    uint8 private immutable _decimals;

    /**
     * @notice Initializes this destination token bridge instance to receive
     * tokens from the specified source chain and token bridge instance.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID,
        address tokenSourceAddress,
        string memory tokenName,
        string memory tokenSymbol,
        uint8 tokenDecimals
    )
        TeleporterTokenDestination(
            teleporterRegistryAddress,
            teleporterManager,
            sourceBlockchainID,
            tokenSourceAddress,
            address(this)
        )
        ERC20(tokenName, tokenSymbol)
    {
        _decimals = tokenDecimals;
    }

    /**
     * @notice For transfers to an `input.destinationBlockchainID` that is not the `sourceBlockchainID`,
     * a multihop transfer is performed, where the tokens are sent back to the token source chain
     * first to check for bridge balance, and then forwarded to the final destination chain.
     *
     * @dev See {IERC20Bridge-send}
     */
    function send(SendTokensInput calldata input, uint256 amount) external {
        _send(input, amount);
    }

    /**
     * @dev See {IERC20Bridge-sendAndCall}
     */
    function sendAndCall(SendAndCallInput calldata input, uint256 amount) external {
        _sendAndCall(input, amount);
    }

    /**
     * @dev See {IERC20-decimals}
     */
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }

    /**
     * @dev See {TeleportTokenDestination-_deposit}
     */
    function _deposit(uint256 amount) internal virtual override returns (uint256) {
        // TODO: can copy logic from SafeERC20TransferFrom.safeTransferFrom directly
        // figure out if has gas savings.
        return SafeERC20TransferFrom.safeTransferFrom(this, amount);
    }

    /**
     * @dev See {TeleportTokenDestination-_withdraw}
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        _mint(recipient, amount);
    }

    /**
     * @dev See {TeleportTokenDestination-_burn}
     *
     * Calls {ERC20-_burn} to burn tokens from this contract.
     */
    function _burn(uint256 amount) internal virtual override {
        _burn(address(this), amount);
    }

    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual override {
        // Mint the tokens to this contract address.
        _mint(address(this), amount);

        // Approve the destination contract to spend the amount.
        _approve(address(this), message.recipientContract, amount);

        // Encode the call to {IERC20SendAndCallReceiver-receiveTokens}
        bytes memory payload = abi.encodeCall(
            IERC20SendAndCallReceiver.receiveTokens,
            (address(this), amount, message.recipientPayload)
        );

        // Call the destination contract with the given payload and gas amount.
        bool success = GasUtils._callWithExactGas(
            message.recipientGasLimit, message.recipientContract, payload
        );

        // Reset the destination contract allowance to 0.
        _approve(address(this), message.recipientContract, 0);

        // If the call failed, send the funds to the fallback recipient.
        if (!success) {
            _transfer(address(this), message.fallbackRecipient, amount);
        }
    }
}
