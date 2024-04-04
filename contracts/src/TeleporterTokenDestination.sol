// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {
    ITeleporterTokenBridge,
    SendTokensInput,
    SendAndCallInput,
    BridgeMessageType,
    BridgeMessage,
    SingleHopSendMessage,
    SingleHopCallMessage,
    MultiHopSendMessage,
    MultiHopCallMessage
} from "./interfaces/ITeleporterTokenBridge.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

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
    address public immutable feeTokenAddress;

    // TODO: these are values brought from the example ERC20Bridge contract.
    // Need to figure out appropriate values.
    uint256 public constant SEND_TOKENS_REQUIRED_GAS = 300_000;

    /**
     * @notice Initializes this destination token bridge instance to receive
     * tokens from the specified source blockchain and token bridge instance.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID_,
        address tokenSourceAddress_,
        address feeTokenAddress_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress, teleporterManager) {
        blockchainID = IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
        require(
            sourceBlockchainID_ != bytes32(0),
            "TeleporterTokenDestination: zero source blockchain ID"
        );
        require(
            sourceBlockchainID_ != blockchainID,
            "TeleporterTokenDestination: cannot deploy to same blockchain as source"
        );
        require(
            tokenSourceAddress_ != address(0),
            "TeleporterTokenDestination: zero token source address"
        );
        require(
            feeTokenAddress_ != address(0), "TeleporterTokenDestination: zero fee token address"
        );
        sourceBlockchainID = sourceBlockchainID_;
        tokenSourceAddress = tokenSourceAddress_;
        feeTokenAddress = feeTokenAddress_;
    }

    /**
     * @notice Sends tokens to the specified destination token bridge instance.
     *
     * @dev Burns the bridged amount, and uses Teleporter to send a cross chain message.
     * Tokens can be sent to the same blockchain this bridge instance is deployed on,
     * to another destination bridge instance.
     * Requirements:
     *
     * - `input.destinationBridgeAddress` cannot be the zero address
     * - `input.recipient` cannot be the zero address
     * - `amount` must be greater than 0
     * - `amount` must be greater than `input.primaryFee`
     */
    function _send(SendTokensInput calldata input, uint256 amount) internal virtual {
        require(input.recipient != address(0), "TeleporterTokenDestination: zero recipient address");
        amount = _prepareSend(
            input.destinationBlockchainID,
            input.destinationBridgeAddress,
            amount,
            input.primaryFee,
            input.secondaryFee
        );

        BridgeMessage memory message;
        if (input.destinationBlockchainID == sourceBlockchainID) {
            // If the destination blockchain is the source bridge instance's blockchain,
            // the destination bridge address must match the token source address.
            require(
                input.destinationBridgeAddress == tokenSourceAddress,
                "TeleporterTokenDestination: invalid destination bridge address"
            );
            message = BridgeMessage({
                messageType: BridgeMessageType.SINGLE_HOP_SEND,
                amount: amount,
                payload: abi.encode(SingleHopSendMessage({recipient: input.recipient}))
            });
        } else {
            // If the destination blockchain ID is this blockchian, the destination
            // bridge address must be a differet contract. This is a multi-hop case to
            // a different bridge contract on this chain.
            if (input.destinationBlockchainID == blockchainID) {
                require(
                    input.destinationBridgeAddress != address(this),
                    "TeleporterTokenDestination: invalid destination bridge address"
                );
            }
            message = BridgeMessage({
                messageType: BridgeMessageType.MULTI_HOP_SEND,
                amount: amount,
                payload: abi.encode(
                    MultiHopSendMessage({
                        destinationBlockchainID: input.destinationBlockchainID,
                        destinationBridgeAddress: input.destinationBridgeAddress,
                        recipient: input.recipient,
                        secondaryFee: input.secondaryFee
                    })
                    )
            });
        }

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: tokenSourceAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeTokenAddress, amount: input.primaryFee}),
                // TODO: placeholder value
                requiredGasLimit: SEND_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: input.allowedRelayerAddresses,
                message: abi.encode(message)
            })
        );

        emit SendTokens(messageID, msg.sender, amount);
    }

    function _sendAndCall(SendAndCallInput memory input, uint256 amount) internal virtual {
        require(
            input.recipientContract != address(0),
            "TeleporterTokenDestination: zero recipient contract address"
        );
        require(
            input.recipientGasLimit >= 21_000,
            "TeleporterTokenDestination: invalid recipient gas limit"
        );
        require(
            input.fallbackRecipient != address(0),
            "TeleporterTokenDestination: zero fallback recipient address"
        );
        amount = _prepareSend(
            input.destinationBlockchainID,
            input.destinationBridgeAddress,
            amount,
            input.primaryFee,
            input.secondaryFee
        );

        BridgeMessage memory message;
        if (input.destinationBlockchainID == sourceBlockchainID) {
            // If the destination blockchain is the source bridge instance's blockchain,
            // the destination bridge address must match the token source address.
            require(
                input.destinationBridgeAddress == tokenSourceAddress,
                "TeleporterTokenDestination: invalid destination bridge address"
            );

            message = BridgeMessage({
                messageType: BridgeMessageType.SINGLE_HOP_CALL,
                amount: amount,
                payload: abi.encode(
                    SingleHopCallMessage({
                        recipientContract: input.recipientContract,
                        recipientPayload: input.recipientPayload,
                        recipientGasLimit: input.recipientGasLimit,
                        fallbackRecipient: input.fallbackRecipient
                    })
                    )
            });
        } else {
            // If the destination blockchain ID is this blockchian, the destination
            // bridge address must be a differet contract. This is a multi-hop case to
            // a different bridge contract on this chain.
            if (input.destinationBlockchainID == blockchainID) {
                require(
                    input.destinationBridgeAddress != address(this),
                    "TeleporterTokenDestination: invalid destination bridge address"
                );
            }

            message = BridgeMessage({
                messageType: BridgeMessageType.MULTI_HOP_CALL,
                amount: amount,
                payload: abi.encode(
                    MultiHopCallMessage({
                        destinationBlockchainID: input.destinationBlockchainID,
                        destinationBridgeAddress: input.destinationBridgeAddress,
                        recipientContract: input.recipientContract,
                        recipientPayload: input.recipientPayload,
                        recipientGasLimit: input.recipientGasLimit,
                        fallbackRecipient: input.fallbackRecipient,
                        secondaryFee: input.secondaryFee
                    })
                    )
            });
        }

        // Send message to the destination bridge address
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeTokenAddress, amount: input.primaryFee}),
                // TODO: Set requiredGasLimit
                requiredGasLimit: SEND_TOKENS_REQUIRED_GAS + input.recipientGasLimit,
                allowedRelayerAddresses: input.allowedRelayerAddresses,
                message: abi.encode(message)
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
        BridgeMessage memory bridgeMessage = abi.decode(message, (BridgeMessage));

        // Destination contracts should only ever receive single-hop messages because
        // multi-hop messages are always routed through the source contract.
        if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_SEND) {
            SingleHopSendMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopSendMessage));

            _withdraw(payload.recipient, bridgeMessage.amount);
        } else if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_CALL) {
            SingleHopCallMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopCallMessage));
            _handleSendAndCall(payload, bridgeMessage.amount);
        } else {
            revert("TeleporterTokenDestination: invalid message type");
        }
    }

    /**
     * @notice Deposits tokens from the sender to this contract,
     * and returns the adjusted amount of tokens deposited.
     * @param amount is initial amount sent to this contract.
     * @return The actual amount deposited to this contract.
     */
    function _deposit(uint256 amount) internal virtual returns (uint256);

    /**
     * @notice Withdraws tokens to the recipient address.
     * @param recipient The address to withdraw tokens to
     * @param amount The amount of tokens to withdraw
     */
    function _withdraw(address recipient, uint256 amount) internal virtual;

    /**
     * @notice Burns a fee adjusted amount of tokens that the user
     * has deposited to this token bridge instance.
     * @param amount The amount of tokens to burn
     */
    function _burn(uint256 amount) internal virtual;

    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual;

    function _prepareSend(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 amount,
        uint256 primaryFee,
        uint256 secondaryFee
    ) private returns (uint256) {
        require(
            destinationBlockchainID != bytes32(0),
            "TeleporterTokenDestination: zero destination blockchain ID"
        );
        require(
            destinationBridgeAddress != address(0),
            "TeleporterTokenDestination: zero destination bridge address"
        );

        // Deposit the funds sent from the user to the bridge,
        // and set to adjusted amount after deposit
        amount = _deposit(amount);
        require(
            amount > primaryFee + secondaryFee,
            "TeleporterTokenDestination: insufficient amount to cover fees"
        );

        amount -= primaryFee;
        _burn(amount);

        return amount;
    }
}
