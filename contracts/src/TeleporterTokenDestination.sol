// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {ITeleporterTokenBridge} from "./interfaces/ITeleporterTokenBridge.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title TeleporterTokenDestination
 * @dev Abstract contract for a Teleporter token bridge that receives tokens from a {TeleporterTokenSource} in exchange for the tokens of this token bridge instance.
 */
abstract contract TeleporterTokenDestination is
    ITeleporterTokenBridge,
    TeleporterOwnerUpgradeable
{
    /// @notice The blockchain ID of the chain this contract is deployed on.
    bytes32 public immutable blockchainID;

    /// @notice The blockchain ID of the source chain this contract receives tokens from.
    bytes32 public immutable sourceBlockchainID;
    /// @notice The address of the source token bridge instance this contract receives tokens from.
    address public immutable tokenSourceAddress;
    /// @notice The ERC20 token this contract uses to pay for Teleporter fees.
    IERC20 public immutable feeToken;

    /**
     * @notice Initializes this destination token bridge instance to receive
     * tokens from the specified source blockchain and token bridge instance.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID_,
        address tokenSourceAddress_,
        address feeToken_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress, teleporterManager) {
        sourceBlockchainID = sourceBlockchainID_;
        tokenSourceAddress = tokenSourceAddress_;
        // TODO: figure out if NativeTokenDestination passes in token or not.
        // NativeTokenDestination will pass in erc20 token it deposits native tokens to pay for fees.
        // ERC20Destination will pass in the erc20 token it's bridging.
        feeToken = IERC20(feeToken_);
        blockchainID = IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
    }

    /**
     * @notice Sends tokens to the specified destination token bridge instance.
     *
     * @dev Burns the bridged amount, and uses Teleporter to send a cross chain message.
     * TODO: Determine if this can be abstracted to a common function with {TeleporterTokenSource}
     * Requirements:
     *
     * - `input.destinationBlockchainID` cannot be the same as the current blockchainID
     * - `input.destinationBridgeAddress` cannot be the zero address
     * - `input.recipient` cannot be the zero address
     * - `amount` must be greater than 0
     * - `amount` must be greater than `input.primaryFee`
     */
    function _send(SendTokensInput calldata input, uint256 amount) internal virtual {
        require(
            input.destinationBlockchainID != blockchainID,
            "TeleporterTokenDestination: cannot bridge to same chain"
        );
        require(
            input.destinationBridgeAddress != address(0),
            "TeleporterTokenDestination: zero destination bridge address"
        );
        require(input.recipient != address(0), "TeleporterTokenDestination: zero recipient address");
        require(amount > 0, "TeleporterTokenDestination: zero send amount");
        require(
            amount > input.primaryFee,
            "TeleporterTokenDestination: insufficient amount to cover fees"
        );

        // TODO: For NativeTokenDestination before this _send, we should exchange the fee amount
        // in native tokens for the fee amount in erc20 tokens. For ERC20Destination, we simply
        // safeTransferFrom the full amount.
        amount -= input.primaryFee;
        _burn(amount);

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: tokenSourceAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: address(feeToken),
                    amount: input.primaryFee
                }),
                // TODO: placeholder value
                requiredGasLimit: 0,
                allowedRelayerAddresses: input.allowedRelayerAddresses,
                message: abi.encode(
                    SendTokensInput({
                        destinationBlockchainID: input.destinationBlockchainID,
                        destinationBridgeAddress: input.destinationBridgeAddress,
                        recipient: input.recipient,
                        primaryFee: input.secondaryFee,
                        secondaryFee: 0,
                        // TODO: Does multihop allowed relayer need to be separate parameter?
                        allowedRelayerAddresses: input.allowedRelayerAddresses
                    }),
                    amount
                    )
            })
        );

        emit SendTokens(messageID, msg.sender, amount);
    }

    /**
     * @notice Verifies the source token bridge instance, and withdraws the amount to the recipient address.
     *
     * @dev See {ITeleporterUpgradeable-_receiveTeleporterMessage}
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID_,
        address originSenderAddress,
        bytes memory message
    ) internal virtual override {
        require(
            sourceBlockchainID_ == sourceBlockchainID,
            "TeleporterTokenDestination: invalid source chain"
        );
        require(
            originSenderAddress == tokenSourceAddress,
            "TeleporterTokenDestination: invalid token source address"
        );
        (address recipient, uint256 amount) = abi.decode(message, (address, uint256));

        _withdraw(recipient, amount);
    }

    /**
     * @notice Burns tokens from the sender's balance.
     * @param amount The amount of tokens to burn
     */
    // solhint-disable-next-line no-empty-blocks
    function _burn(uint256 amount) internal virtual;

    /**
     * @notice Withdraws tokens to the recipient address.
     * @param recipient The address to withdraw tokens to
     * @param amount The amount of tokens to withdraw
     */
    // solhint-disable-next-line no-empty-blocks
    function _withdraw(address recipient, uint256 amount) internal virtual;
}
