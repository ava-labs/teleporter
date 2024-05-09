// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {ITeleporterTokenSource} from "./interfaces/ITeleporterTokenSource.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    BridgeMessageType,
    BridgeMessage,
    SingleHopSendMessage,
    SingleHopCallMessage,
    MultiHopSendMessage,
    MultiHopCallMessage,
    RegisterDestinationMessage
} from "./interfaces/ITeleporterTokenBridge.sol";
import {SendReentrancyGuard} from "./utils/SendReentrancyGuard.sol";
import {TokenScalingUtils} from "./utils/TokenScalingUtils.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Each destination bridge instance registers with the source token bridge contract,
 * and provides settings for bridging to the destination bridge.
 * @param registered whether the destination bridge is registered
 * @param collateralNeeded the amount of tokens that must be first added as collateral,
 * through `addCollateral` calls, before tokens can be bridged to the destination token bridge.
 * @param tokenMultiplier the scaling factor for the amount of tokens to be bridged to the destination.
 * @param multiplyOnSend whether the scaling factor is multiplied or divided when sending to the destination.
 */
struct DestinationBridgeSettings {
    bool registered;
    uint256 collateralNeeded;
    uint256 tokenMultiplier;
    bool multiplyOnSend;
}

/**
 * @title TeleporterTokenSource
 * @dev Abstract contract for a Teleporter token bridge that sends tokens to {TeleporterTokenDestination} instances.
 *
 * This contract also handles multi-hop transfers, where tokens sent from a {TeleporterTokenDestination}
 * instance are forwarded to another {TeleporterTokenDestination} instance.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
abstract contract TeleporterTokenSource is
    ITeleporterTokenSource,
    TeleporterOwnerUpgradeable,
    SendReentrancyGuard
{
    /// @notice The blockchain ID of the chain this contract is deployed on.
    bytes32 public immutable blockchainID;

    /// @notice The ERC20 token this contract uses to pay for Teleporter fees.
    address public immutable feeTokenAddress;

    /**
     * @notice Tracks the settings for each destination bridge instance. Destination bridge instances
     * must register with their {TeleporterTokenSource} contracts via Teleporter message to be able to
     * receive tokens from this contract.
     */
    mapping(
        bytes32 destinationBlockchainID
            => mapping(
                address destinationBridgeAddress => DestinationBridgeSettings destinationSettings
            )
    ) public registeredDestinations;

    /**
     * @notice Tracks the balances of tokens sent to other bridge instances.
     * Bridges are not allowed to unwrap more than has been sent to them.
     * @dev (destinationBlockchainID, destinationBridgeAddress) -> balance
     */
    mapping(
        bytes32 destinationBlockchainID
            => mapping(address destinationBridgeAddress => uint256 balance)
    ) public bridgedBalances;

    /**
     * @notice Initializes this source token bridge instance to send
     * tokens to the specified destination chain and token bridge instance.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address feeTokenAddress_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress, teleporterManager) {
        blockchainID = IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
        require(feeTokenAddress_ != address(0), "TeleporterTokenSource: zero fee token address");
        feeTokenAddress = feeTokenAddress_;
    }

    function _registerDestination(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 initialReserveImbalance,
        uint256 tokenMultiplier,
        bool multiplyOnSend
    ) internal {
        require(
            destinationBlockchainID != bytes32(0),
            "TeleporterTokenSource: zero destination blockchain ID"
        );
        require(
            destinationBlockchainID != blockchainID,
            "TeleporterTokenSource: cannot register bridge on same chain"
        );
        require(
            destinationBridgeAddress != address(0),
            "TeleporterTokenSource: zero destination bridge address"
        );
        require(
            tokenMultiplier > 0 && tokenMultiplier < 1e18,
            "TeleporterTokenSource: invalid token multiplier"
        );
        require(
            !registeredDestinations[destinationBlockchainID][destinationBridgeAddress].registered,
            "TeleporterTokenSource: destination already registered"
        );

        // Calculate the collateral needed in source token denomination.
        uint256 collateralNeeded = TokenScalingUtils.removeTokenScale(
            tokenMultiplier, multiplyOnSend, initialReserveImbalance
        );

        // Round up the collateral needed by 1 in the case that `multiplyOnSend` is true and
        // `initialReserveImbalance` is not divisible by the `tokenMultiplier` to
        // ensure that the full amount is account for.
        if (multiplyOnSend && initialReserveImbalance % tokenMultiplier != 0) {
            collateralNeeded += 1;
        }

        registeredDestinations[destinationBlockchainID][destinationBridgeAddress] =
        DestinationBridgeSettings({
            registered: true,
            collateralNeeded: collateralNeeded,
            tokenMultiplier: tokenMultiplier,
            multiplyOnSend: multiplyOnSend
        });

        emit DestinationRegistered(
            destinationBlockchainID,
            destinationBridgeAddress,
            collateralNeeded,
            tokenMultiplier,
            multiplyOnSend
        );
    }

    /**
     * @notice Sends tokens to the specified destination token bridge instance.
     *
     * @dev Increases the bridge balance sent to each destination token bridge instance,
     * and uses Teleporter to send a cross chain message. The amount passed is assumed to
     * be already scaled to the local denomination for this token source.
     * Requirements:
     *
     * - `input.destinationBlockchainID` cannot be the same as the current blockchainID
     * - `input.destinationBridgeAddress` cannot be the zero address
     * - `input.recipient` cannot be the zero address
     * - `amount` must be greater than 0
     * - `amount` must be greater than `input.primaryFee`
     */
    function _send(
        SendTokensInput memory input,
        uint256 amount,
        bool isMultihop
    ) internal sendNonReentrant {
        require(input.recipient != address(0), "TeleporterTokenSource: zero recipient address");
        require(input.requiredGasLimit > 0, "TeleporterTokenSource: zero required gas limit");
        require(input.secondaryFee == 0, "TeleporterTokenSource: non-zero secondary fee");

        uint256 adjustedAmount;
        if (isMultihop) {
            adjustedAmount = _prepareMultiHopRouting(
                input.destinationBlockchainID,
                input.destinationBridgeAddress,
                amount,
                input.primaryFee
            );

            if (adjustedAmount == 0) {
                // If the adjusted amount is zero for any reason (i.e. unsupported destination,
                // being scaled down to zero, etc.), send the tokens to the fallback recipient.
                _withdraw(input.fallbackRecipient, amount);
                return;
            }
        } else {
            adjustedAmount = _prepareSend(
                input.destinationBlockchainID,
                input.destinationBridgeAddress,
                amount,
                input.primaryFee
            );
        }

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_SEND,
            payload: abi.encode(
                SingleHopSendMessage({recipient: input.recipient, amount: adjustedAmount})
                )
        });

        // Send message to the destination bridge address
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeTokenAddress, amount: input.primaryFee}),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        if (isMultihop) {
            emit TokensRouted(messageID, input, adjustedAmount);
        } else {
            emit TokensSent(messageID, msg.sender, input, adjustedAmount);
        }
    }

    function _sendAndCall(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        SendAndCallInput memory input,
        uint256 amount,
        bool isMultihop
    ) internal sendNonReentrant {
        require(
            input.recipientContract != address(0),
            "TeleporterTokenSource: zero recipient contract address"
        );
        require(input.requiredGasLimit > 0, "TeleporterTokenSource: zero required gas limit");
        require(input.recipientGasLimit > 0, "TeleporterTokenSource: zero recipient gas limit");
        require(
            input.recipientGasLimit < input.requiredGasLimit,
            "TeleporterTokenSource: invalid recipient gas limit"
        );
        require(
            input.fallbackRecipient != address(0),
            "TeleporterTokenSource: zero fallback recipient address"
        );

        uint256 adjustedAmount;
        if (isMultihop) {
            adjustedAmount = _prepareMultiHopRouting(
                input.destinationBlockchainID,
                input.destinationBridgeAddress,
                amount,
                input.primaryFee
            );

            if (adjustedAmount == 0) {
                // If the adjusted amount is zero for any reason (i.e. unsupported destination,
                // being scaled down to zero, etc.), send the tokens to the fallback recipient.
                _withdraw(input.fallbackRecipient, amount);
                return;
            }
        } else {
            adjustedAmount = _prepareSend(
                input.destinationBlockchainID,
                input.destinationBridgeAddress,
                amount,
                input.primaryFee
            );
        }

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_CALL,
            payload: abi.encode(
                SingleHopCallMessage({
                    sourceBlockchainID: sourceBlockchainID,
                    originSenderAddress: originSenderAddress,
                    recipientContract: input.recipientContract,
                    amount: adjustedAmount,
                    recipientPayload: input.recipientPayload,
                    recipientGasLimit: input.recipientGasLimit,
                    fallbackRecipient: input.fallbackRecipient
                })
                )
        });

        // Send message to the destination bridge address
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeTokenAddress, amount: input.primaryFee}),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        if (isMultihop) {
            emit TokensAndCallRouted(messageID, input, adjustedAmount);
        } else {
            emit TokensAndCallSent(messageID, originSenderAddress, input, adjustedAmount);
        }
    }

    /**
     * @dev See {INativeTokenSource-addCollateral}
     */
    function _addCollateral(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 amount
    ) internal sendNonReentrant {
        DestinationBridgeSettings memory destinationSettings =
            registeredDestinations[destinationBlockchainID][destinationBridgeAddress];
        require(
            destinationSettings.registered,
            "TeleporterTokenSource: destination bridge not registered"
        );
        require(
            destinationSettings.collateralNeeded > 0,
            "TeleporterTokenSource: zero collateral needed"
        );

        // Deposit the full amount, and withdraw back to the sender if there is excess.
        amount = _deposit(amount);

        // Calculate the amount remaining.
        uint256 remainingCollateralNeeded;
        uint256 excessAmount;
        if (amount >= destinationSettings.collateralNeeded) {
            remainingCollateralNeeded = 0;
            excessAmount = amount - destinationSettings.collateralNeeded;
            amount = destinationSettings.collateralNeeded;
        } else {
            remainingCollateralNeeded = destinationSettings.collateralNeeded - amount;
        }

        // Update the remaining collateral needed.
        registeredDestinations[destinationBlockchainID][destinationBridgeAddress].collateralNeeded =
            remainingCollateralNeeded;
        emit CollateralAdded(
            destinationBlockchainID, destinationBridgeAddress, amount, remainingCollateralNeeded
        );

        // If there is excess amount, send it back to the sender.
        if (excessAmount > 0) {
            _withdraw(msg.sender, excessAmount);
        }
    }

    /**
     * @dev See {ITeleporterUpgradeable-_receiveTeleporterMessage}
     *
     * Verifies the Teleporter token bridge sending back tokens has enough balance,
     * and adjusts the bridge balance accordingly. If the final destination for this token
     * is this contract, the tokens are withdrawn and sent to the recipient. Otherwise,
     * a multi-hop is performed, and the tokens are forwarded to the destination token bridge.
     * Requirements:
     *
     * - `sourceBlockchainID` and `originSenderAddress` have enough bridge balance to send back.
     * - `input.destinationBridgeAddress` is this contract is this chain is the final destination.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        BridgeMessage memory bridgeMessage = abi.decode(message, (BridgeMessage));
        if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_SEND) {
            SingleHopSendMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopSendMessage));

            uint256 sourceAmount =
                _processReceivedTransfer(sourceBlockchainID, originSenderAddress, payload.amount);

            // Send the tokens to the recipient.
            _withdraw(payload.recipient, sourceAmount);
            return;
        } else if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_CALL) {
            SingleHopCallMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopCallMessage));

            uint256 sourceAmount =
                _processReceivedTransfer(sourceBlockchainID, originSenderAddress, payload.amount);

            // Verify that the payload's source blockchain ID matches the source blockchain ID passed from Teleporter.
            // Prevents a destination bridge from accessing tokens attributed to another destination bridge instance.
            require(
                payload.sourceBlockchainID == sourceBlockchainID,
                "TeleporterTokenSource: mismatched source blockchain ID"
            );

            _handleSendAndCall(payload, sourceAmount);
            return;
        } else if (bridgeMessage.messageType == BridgeMessageType.MULTI_HOP_SEND) {
            MultiHopSendMessage memory payload =
                abi.decode(bridgeMessage.payload, (MultiHopSendMessage));

            uint256 sourceAmount =
                _processReceivedTransfer(sourceBlockchainID, originSenderAddress, payload.amount);

            _send(
                SendTokensInput({
                    destinationBlockchainID: payload.destinationBlockchainID,
                    destinationBridgeAddress: payload.destinationBridgeAddress,
                    recipient: payload.recipient,
                    primaryFee: payload.secondaryFee,
                    secondaryFee: 0,
                    requiredGasLimit: payload.secondaryGasLimit,
                    fallbackRecipient: payload.fallbackRecipient
                }),
                sourceAmount,
                true
            );
            return;
        } else if (bridgeMessage.messageType == BridgeMessageType.MULTI_HOP_CALL) {
            MultiHopCallMessage memory payload =
                abi.decode(bridgeMessage.payload, (MultiHopCallMessage));

            uint256 sourceAmount =
                _processReceivedTransfer(sourceBlockchainID, originSenderAddress, payload.amount);

            _sendAndCall(
                sourceBlockchainID,
                payload.originSenderAddress,
                SendAndCallInput({
                    destinationBlockchainID: payload.destinationBlockchainID,
                    destinationBridgeAddress: payload.destinationBridgeAddress,
                    recipientContract: payload.recipientContract,
                    recipientPayload: payload.recipientPayload,
                    requiredGasLimit: payload.secondaryRequiredGasLimit,
                    recipientGasLimit: payload.recipientGasLimit,
                    fallbackRecipient: payload.fallbackRecipient,
                    primaryFee: payload.secondaryFee,
                    secondaryFee: 0
                }),
                sourceAmount,
                true
            );
            return;
        } else if (bridgeMessage.messageType == BridgeMessageType.REGISTER_DESTINATION) {
            RegisterDestinationMessage memory payload =
                abi.decode(bridgeMessage.payload, (RegisterDestinationMessage));
            _registerDestination(
                sourceBlockchainID,
                originSenderAddress,
                payload.initialReserveImbalance,
                payload.tokenMultiplier,
                payload.multiplyOnSend
            );
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
     * @notice Processes a send and call message by calling the recipient contract.
     * @param message The send and call message include recipient calldata
     * @param amount The amount of tokens to be sent to the recipient. This amount is assumed to be
     * already scaled to the local denomination for this token source.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual;

    /**
     * @notice Processes a received transfer from a destination bridge instance.
     * Validates that the message is sent from a registered destination bridge instance,
     * and is already collateralized.
     * Deducts the balance bridged to the given destination.
     * Removes the token scaling of the destination, checks the associated source token
     * amount is greater than zero, and returns the source token amount.
     * @param destinationBlockchainID The blockchain ID of the destination bridge instance.
     * @param destinationBridgeAddress The address of the destination bridge instance.
     * @param amount The amount of tokens sent back from destination, denominated by the
     * destination's token scale amount.
     */
    function _processReceivedTransfer(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 amount
    ) private returns (uint256) {
        DestinationBridgeSettings memory destinationSettings =
            registeredDestinations[destinationBlockchainID][destinationBridgeAddress];

        // Require that the destination bridge is registered and has no collateral needed.
        require(
            destinationSettings.registered,
            "TeleporterTokenSource: destination bridge not registered"
        );
        require(
            destinationSettings.collateralNeeded == 0,
            "TeleporterTokenSource: destination bridge not collateralized"
        );

        // Deduct the balance bridged to the given destination bridge address prior to scaling the amount.
        _deductSenderBalance(destinationBlockchainID, destinationBridgeAddress, amount);

        // Remove the token scaling of the destination and get source token amount.
        uint256 sourceAmount = TokenScalingUtils.removeTokenScale(
            destinationSettings.tokenMultiplier, destinationSettings.multiplyOnSend, amount
        );

        // Require that the source token amount is greater than zero after removed scaling.
        require(sourceAmount > 0, "TeleporterTokenSource: zero token amount");
        return sourceAmount;
    }

    /**
     * @notice Prepares a multi-hop send by checking the destination bridge settings
     * and adjusting the amount to be sent.
     * @return The scaled amount to be sent to the destination bridge. If zero is returned,
     * the tokens are sent to the fallback recipient. Zero can be returned if the
     * destination is not registered, needs collateral, or the scaled amount is zero.
     */
    function _prepareMultiHopRouting(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 amount,
        uint256 fee
    ) private returns (uint256) {
        DestinationBridgeSettings memory destinationSettings =
            registeredDestinations[destinationBlockchainID][destinationBridgeAddress];
        if (!destinationSettings.registered || destinationSettings.collateralNeeded > 0) {
            return 0;
        }

        // Subtract fee amount from amount prior to scaling.
        require(amount > fee, "TeleporterTokenSource: insufficient amount to cover fees");
        amount -= fee;

        // Scale the amount based on the token multiplier for the given destination.
        uint256 scaledAmount = TokenScalingUtils.applyTokenScale(
            destinationSettings.tokenMultiplier, destinationSettings.multiplyOnSend, amount
        );
        if (scaledAmount == 0) {
            return 0;
        }

        // Increase the balance of the destination bridge by the scaled amount.
        bridgedBalances[destinationBlockchainID][destinationBridgeAddress] += scaledAmount;

        return scaledAmount;
    }

    /**
     * @dev Prepares tokens to be sent to another chain by handling the
     * locking of the token amount in this contract and updating the accounting
     * balances.
     */
    function _prepareSend(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 amount,
        uint256 fee
    ) private returns (uint256) {
        DestinationBridgeSettings memory destinationSettings =
            registeredDestinations[destinationBlockchainID][destinationBridgeAddress];
        require(destinationSettings.registered, "TeleporterTokenSource: destination not registered");
        require(
            destinationSettings.collateralNeeded == 0,
            "TeleporterTokenSource: collateral needed for destination"
        );

        // Lock the amount in this contract to be sent.
        amount = _deposit(amount);
        require(amount > fee, "TeleporterTokenSource: insufficient amount to cover fees");

        // Subtract fee amount from amount prior to scaling.
        amount -= fee;

        // Scale the amount based on the token multiplier for the given destination.
        uint256 scaledAmount = TokenScalingUtils.applyTokenScale(
            destinationSettings.tokenMultiplier, destinationSettings.multiplyOnSend, amount
        );
        require(scaledAmount > 0, "TeleporterTokenSource: zero scaled amount");

        // Increase the balance of the destination bridge by the scaled amount.
        bridgedBalances[destinationBlockchainID][destinationBridgeAddress] += scaledAmount;

        return scaledAmount;
    }

    function _deductSenderBalance(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        uint256 amount
    ) private {
        uint256 senderBalance = bridgedBalances[destinationBlockchainID][destinationBridgeAddress];
        require(senderBalance >= amount, "TeleporterTokenSource: insufficient bridge balance");
        bridgedBalances[destinationBlockchainID][destinationBridgeAddress] = senderBalance - amount;
    }
}
