// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {ITokenHub} from "./interfaces/ITokenHub.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    BridgeMessageType,
    BridgeMessage,
    SingleHopSendMessage,
    SingleHopCallMessage,
    MultiHopSendMessage,
    MultiHopCallMessage,
    RegisterSpokeMessage
} from "../interfaces/ITokenBridge.sol";
import {SendReentrancyGuard} from "../utils/SendReentrancyGuard.sol";
import {TokenScalingUtils} from "../utils/TokenScalingUtils.sol";
import {SafeERC20TransferFrom} from "../utils/SafeERC20TransferFrom.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Each spoke instance registers with the hub contract, and provides settings for bridging
 * to the spoke bridge contract.
 * @param registered Whether the spoke bridge is registered
 * @param collateralNeeded The amount of tokens that must be first added as collateral,
 * through {addCollateral} calls, before tokens can be bridged to the spoke token bridge.
 * @param tokenMultiplier The scaling factor for the amount of tokens to be bridged to the spoke.
 * @param multiplyOnSpoke Whether the {tokenMultiplier} should be applied when transferring tokens to
 * the spoke (multiplyOnSpoke=true), or when transferring tokens back to the hub (multiplyOnSpoke=false).
 */
struct SpokeBridgeSettings {
    bool registered;
    uint256 collateralNeeded;
    uint256 tokenMultiplier;
    bool multiplyOnSpoke;
}

/**
 * @title TokenHub
 * @dev Abstract contract for a token bridge hub that sends its specified token to {TokenSpoke} instances.
 *
 * This contract also handles multi-hop transfers, where tokens sent from a {TokenSpoke}
 * instance are forwarded to another {TokenSpoke} instance.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
abstract contract TokenHub is ITokenHub, TeleporterOwnerUpgradeable, SendReentrancyGuard {
    /// @notice The blockchain ID of the chain this contract is deployed on.
    bytes32 public blockchainID;

    /**
     * @notice The token address this hub contract bridges to spoke instances.
     * For multi-hop transfers, this {tokenAddress} is always used to pay for the secondary message fees.
     * If the token is an ERC20 token, the contract address is directly passed in.
     * If the token is a native asset, the contract address is the wrapped token contract.
     */
    address public tokenAddress;

    uint8 public tokenDecimals;

    /**
     * @notice Tracks the settings for each spoke bridge instance. Spoke instances
     * must register with their {TokenHub} contracts via Teleporter message to be able to
     * receive tokens from this contract.
     */
    mapping(
        bytes32 spokeBlockchainID
            => mapping(address spokeBridgeAddress => SpokeBridgeSettings spokeSettings)
    ) public registeredSpokes;

    /**
     * @notice Tracks the balances of tokens sent to spoke instances.
     * Balances are represented in the spoke token's denomination,
     * and bridges are not allowed to unwrap more than has been sent to them.
     * @dev (spokeBlockchainID, spokeBridgeAddress) -> balance
     */
    mapping(bytes32 spokeBlockchainID => mapping(address spokeBridgeAddress => uint256 balance))
        public bridgedBalances;

    /**
     * @notice Initializes this hub bridge instance to send tokens to spoke instances on other chains.
     */
    function __TokenHub_init(
        address teleporterRegistryAddress,
        address teleporterManager,
        address tokenAddress_,
        uint8 tokenDecimals_
    ) public virtual onlyInitializing {
        __TeleporterOwnerUpgradeable_init(teleporterRegistryAddress, teleporterManager);
        __SendReentrancyGuard_init();
        blockchainID = IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
        require(tokenAddress_ != address(0), "TokenHub: zero token address");
        require(
            tokenDecimals_ <= TokenScalingUtils.MAX_TOKEN_DECIMALS,
            "TokenHub: token decimals too high"
        );
        tokenAddress = tokenAddress_;
        tokenDecimals = tokenDecimals_;
    }

    function _registerSpoke(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        RegisterSpokeMessage memory message
    ) internal {
        require(spokeBlockchainID != bytes32(0), "TokenHub: zero spoke blockchain ID");
        require(spokeBlockchainID != blockchainID, "TokenHub: cannot register spoke on same chain");
        require(spokeBridgeAddress != address(0), "TokenHub: zero spoke bridge address");
        require(
            !registeredSpokes[spokeBlockchainID][spokeBridgeAddress].registered,
            "TokenHub: spoke already registered"
        );
        require(
            message.spokeTokenDecimals <= TokenScalingUtils.MAX_TOKEN_DECIMALS,
            "TokenHub: spoke token decimals too high"
        );
        require(message.hubTokenDecimals == tokenDecimals, "TokenHub: invalid hub token decimals");

        (uint256 tokenMultiplier, bool multiplyOnSpoke) =
            TokenScalingUtils.deriveTokenMultiplierValues(tokenDecimals, message.spokeTokenDecimals);

        // Calculate the collateral needed in hub token denomination.
        uint256 collateralNeeded = TokenScalingUtils.removeTokenScale(
            tokenMultiplier, multiplyOnSpoke, message.initialReserveImbalance
        );

        // Round up the collateral needed by 1 in the case that {multiplyOnSpoke} is true and
        // {initialReserveImbalance} is not divisible by the {tokenMultiplier} to
        // ensure that the full amount is accounted for.
        if (multiplyOnSpoke && message.initialReserveImbalance % tokenMultiplier != 0) {
            collateralNeeded += 1;
        }

        registeredSpokes[spokeBlockchainID][spokeBridgeAddress] = SpokeBridgeSettings({
            registered: true,
            collateralNeeded: collateralNeeded,
            tokenMultiplier: tokenMultiplier,
            multiplyOnSpoke: multiplyOnSpoke
        });

        emit SpokeRegistered(
            spokeBlockchainID, spokeBridgeAddress, collateralNeeded, message.spokeTokenDecimals
        );
    }

    /**
     * @notice Sends tokens to the specified spoke token bridge instance.
     *
     * @dev Increases the bridge balance sent to the spoke bridge instance,
     * and uses Teleporter to send a cross chain message. The amount passed is assumed to
     * be already scaled to the local denomination for this token hub.
     * Requirements:
     *
     * - The spoke instance specified by {input.destinationBlockchainID} and {input.destinationBridgeAddress} must
     *   be registered with this contract.
     * - {input.recipient} cannot be the zero address
     * - {amount} must be greater than 0
     */
    function _send(SendTokensInput memory input, uint256 amount) internal sendNonReentrant {
        _validateSendTokensInput(input);
        // Require that a single hop transfer does not have a multi-hop fallback recipient.
        require(input.multiHopFallback == address(0), "TokenHub: non-zero multi-hop fallback");

        (uint256 adjustedAmount, uint256 feeAmount) = _prepareSend({
            spokeBlockchainID: input.destinationBlockchainID,
            spokeBridgeAddress: input.destinationBridgeAddress,
            amount: amount,
            primaryFeeTokenAddress: input.primaryFeeTokenAddress,
            feeAmount: input.primaryFee
        });

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_SEND,
            payload: abi.encode(
                SingleHopSendMessage({recipient: input.recipient, amount: adjustedAmount})
                )
        });

        // Send message to the spoke instance
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensSent(messageID, _msgSender(), input, adjustedAmount);
    }

    /**
     * @notice Routes tokens from a multi-hop message to the specified spoke token bridge instance.
     *
     * @dev Increases the bridge balance sent to the spoke token bridge instance,
     * and uses Teleporter to send a cross chain message. The amount passed is assumed to
     * be already scaled to the local denomination for this token hub.
     * Requirements:
     *
     * - The spoke instance specified by {input.destinationBlockchainID} and {input.destinationBridgeAddress} must
     *   be registered with this contract.
     * - {input.recipient} cannot be the zero address
     * - {amount} must be greater than 0
     */
    function _routeMultiHop(
        SendTokensInput memory input,
        uint256 amount
    ) internal sendNonReentrant {
        _validateSendTokensInput(input);

        uint256 adjustedAmount = _prepareMultiHopRouting(
            input.destinationBlockchainID, input.destinationBridgeAddress, amount, input.primaryFee
        );

        if (adjustedAmount == 0) {
            // If the adjusted amount is zero for any reason (i.e. unregistered spoke,
            // being scaled down to zero, etc.), send the tokens to the multi-hop fallback.
            _withdraw(input.multiHopFallback, amount);
            return;
        }

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_SEND,
            payload: abi.encode(
                SingleHopSendMessage({recipient: input.recipient, amount: adjustedAmount})
                )
        });

        // Send message to the spoke instance.
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: input.primaryFee
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensRouted(messageID, input, adjustedAmount);
    }

    function _sendAndCall(
        bytes32 sourceBlockchainID,
        address originBridgeAddress,
        address originSenderAddress,
        SendAndCallInput memory input,
        uint256 amount
    ) internal sendNonReentrant {
        _validateSendAndCallInput(input);

        // Require that a single hop transfer does not have a multi-hop fallback recipient.
        require(input.multiHopFallback == address(0), "TokenHub: non-zero multi-hop fallback");

        (uint256 adjustedAmount, uint256 feeAmount) = _prepareSend({
            spokeBlockchainID: input.destinationBlockchainID,
            spokeBridgeAddress: input.destinationBridgeAddress,
            amount: amount,
            primaryFeeTokenAddress: input.primaryFeeTokenAddress,
            feeAmount: input.primaryFee
        });

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_CALL,
            payload: abi.encode(
                SingleHopCallMessage({
                    sourceBlockchainID: sourceBlockchainID,
                    originBridgeAddress: originBridgeAddress,
                    originSenderAddress: originSenderAddress,
                    recipientContract: input.recipientContract,
                    amount: adjustedAmount,
                    recipientPayload: input.recipientPayload,
                    recipientGasLimit: input.recipientGasLimit,
                    fallbackRecipient: input.fallbackRecipient
                })
                )
        });

        // Send message to the spoke instance.
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: feeAmount
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensAndCallSent(messageID, originSenderAddress, input, adjustedAmount);
    }

    function _routeMultiHopSendAndCall(
        bytes32 sourceBlockchainID,
        address originBridgeAddress,
        address originSenderAddress,
        SendAndCallInput memory input,
        uint256 amount
    ) internal sendNonReentrant {
        _validateSendAndCallInput(input);
        uint256 adjustedAmount = _prepareMultiHopRouting(
            input.destinationBlockchainID, input.destinationBridgeAddress, amount, input.primaryFee
        );

        if (adjustedAmount == 0) {
            // If the adjusted amount is zero for any reason (i.e. unregistered spoke,
            // being scaled down to zero, etc.), send the tokens to the multi-hop fallback recipient.
            _withdraw(input.multiHopFallback, amount);
            return;
        }

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_CALL,
            payload: abi.encode(
                SingleHopCallMessage({
                    sourceBlockchainID: sourceBlockchainID,
                    originBridgeAddress: originBridgeAddress,
                    originSenderAddress: originSenderAddress,
                    recipientContract: input.recipientContract,
                    amount: adjustedAmount,
                    recipientPayload: input.recipientPayload,
                    recipientGasLimit: input.recipientGasLimit,
                    fallbackRecipient: input.fallbackRecipient
                })
                )
        });

        // Send message to the spoke instance.
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationAddress: input.destinationBridgeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: input.primaryFee
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensAndCallRouted(messageID, input, adjustedAmount);
    }

    /**
     * @dev See {INativeTokenHub-addCollateral}
     */
    function _addCollateral(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount
    ) internal sendNonReentrant {
        SpokeBridgeSettings memory spokeSettings =
            registeredSpokes[spokeBlockchainID][spokeBridgeAddress];
        require(spokeSettings.registered, "TokenHub: spoke not registered");
        require(spokeSettings.collateralNeeded > 0, "TokenHub: zero collateral needed");

        // Deposit the full amount, and withdraw back to the sender if there is excess.
        amount = _deposit(amount);

        // Calculate the remaining collateral needed, any excess amount, and adjust
        // {amount} to represent the amount of tokens added as collateral.
        uint256 remainingCollateralNeeded;
        uint256 excessAmount;
        if (amount >= spokeSettings.collateralNeeded) {
            remainingCollateralNeeded = 0;
            excessAmount = amount - spokeSettings.collateralNeeded;
            amount = spokeSettings.collateralNeeded;
        } else {
            remainingCollateralNeeded = spokeSettings.collateralNeeded - amount;
        }

        // Update the remaining collateral needed.
        registeredSpokes[spokeBlockchainID][spokeBridgeAddress].collateralNeeded =
            remainingCollateralNeeded;
        emit CollateralAdded(
            spokeBlockchainID, spokeBridgeAddress, amount, remainingCollateralNeeded
        );

        // If there is excess amount, send it back to the sender.
        if (excessAmount > 0) {
            _withdraw(_msgSender(), excessAmount);
        }
    }

    /**
     * @dev See {ITeleporterUpgradeable-_receiveTeleporterMessage}
     *
     * Handles the processing of Teleporter messages sent to this contract.
     * Supported message types include registering a spoke instance, single-hop sends,
     * single-hop calls, multi-hop sends, and multi-hop calls.
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

            uint256 hubAmount =
                _processSingleHopTransfer(sourceBlockchainID, originSenderAddress, payload.amount);

            // Send the tokens to the recipient.
            _withdraw(payload.recipient, hubAmount);
            return;
        } else if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_CALL) {
            SingleHopCallMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopCallMessage));

            uint256 hubAmount =
                _processSingleHopTransfer(sourceBlockchainID, originSenderAddress, payload.amount);

            // Verify that the payload's source blockchain ID and origin bridge address matches the source blockchain ID
            // and origin sender address passed from Teleporter.
            require(
                payload.sourceBlockchainID == sourceBlockchainID,
                "TokenHub: mismatched source blockchain ID"
            );
            require(
                payload.originBridgeAddress == originSenderAddress,
                "TokenHub: mismatched origin sender address"
            );

            _handleSendAndCall(payload, hubAmount);
            return;
        } else if (bridgeMessage.messageType == BridgeMessageType.MULTI_HOP_SEND) {
            MultiHopSendMessage memory payload =
                abi.decode(bridgeMessage.payload, (MultiHopSendMessage));

            (uint256 hubAmount, uint256 fee) = _processMultiHopTransfer(
                sourceBlockchainID, originSenderAddress, payload.amount, payload.secondaryFee
            );

            // For a multi-hop send, the fee token address has to be {tokenAddress},
            // because the fee is taken from the amount that has already been deposited.
            // For ERC20 tokens, the token address of the contract is directly passed.
            // For native assets, the contract address is the wrapped token contract.
            _routeMultiHop(
                SendTokensInput({
                    destinationBlockchainID: payload.destinationBlockchainID,
                    destinationBridgeAddress: payload.destinationBridgeAddress,
                    recipient: payload.recipient,
                    primaryFeeTokenAddress: tokenAddress,
                    primaryFee: fee,
                    secondaryFee: 0,
                    requiredGasLimit: payload.secondaryGasLimit,
                    multiHopFallback: payload.multiHopFallback
                }),
                hubAmount
            );
            return;
        } else if (bridgeMessage.messageType == BridgeMessageType.MULTI_HOP_CALL) {
            MultiHopCallMessage memory payload =
                abi.decode(bridgeMessage.payload, (MultiHopCallMessage));

            (uint256 hubAmount, uint256 fee) = _processMultiHopTransfer(
                sourceBlockchainID, originSenderAddress, payload.amount, payload.secondaryFee
            );

            // For a multi-hop send, the fee token address has to be {tokenAddress},
            // because the fee is taken from the amount that has already been deposited.
            // For ERC20 tokens, the token address of the contract is directly passed.
            // For native assets, the contract address is the wrapped token contract.
            _routeMultiHopSendAndCall({
                sourceBlockchainID: sourceBlockchainID,
                originBridgeAddress: originSenderAddress,
                originSenderAddress: payload.originSenderAddress,
                input: SendAndCallInput({
                    destinationBlockchainID: payload.destinationBlockchainID,
                    destinationBridgeAddress: payload.destinationBridgeAddress,
                    recipientContract: payload.recipientContract,
                    recipientPayload: payload.recipientPayload,
                    requiredGasLimit: payload.secondaryRequiredGasLimit,
                    recipientGasLimit: payload.recipientGasLimit,
                    multiHopFallback: payload.multiHopFallback,
                    fallbackRecipient: payload.fallbackRecipient,
                    primaryFeeTokenAddress: tokenAddress,
                    primaryFee: fee,
                    secondaryFee: 0
                }),
                amount: hubAmount
            });
            return;
        } else if (bridgeMessage.messageType == BridgeMessageType.REGISTER_SPOKE) {
            RegisterSpokeMessage memory payload =
                abi.decode(bridgeMessage.payload, (RegisterSpokeMessage));
            _registerSpoke(sourceBlockchainID, originSenderAddress, payload);
        }
    }

    /**
     * @notice Deposits tokens from the sender to this contract,
     * and returns the adjusted amount of tokens deposited.
     * @param amount The initial amount sent to this contract.
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
     * already scaled to the local denomination for this token hub.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual;

    /**
     * @notice Processes a received single hop transfer from a spoke instance.
     * Validates that the message is sent from a registered spoke instance,
     * and is already collateralized.
     * @param spokeBlockchainID The blockchain ID of the spoke instance.
     * @param spokeBridgeAddress The address of the spoke instance.
     * @param amount The amount of tokens sent back from spoke, denominated by the
     * spoke's token scale amount.
     */
    function _processSingleHopTransfer(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount
    ) private returns (uint256) {
        SpokeBridgeSettings memory spokeSettings =
            registeredSpokes[spokeBlockchainID][spokeBridgeAddress];

        return
            _processReceivedTransfer(spokeSettings, spokeBlockchainID, spokeBridgeAddress, amount);
    }

    /**
     * @notice Processes a received multi-hop transfer from a spoke instance.
     * Validates that the message is sent from a registered spoke instance,
     * and is already collateralized.
     * @param spokeBlockchainID The blockchain ID of the spoke instance.
     * @param spokeBridgeAddress The address of the spoke instance.
     * @param amount The amount of tokens sent back from spoke, denominated by the
     * spoke's token scale amount.
     * @param secondaryFee The Teleporter fee for the second hop of the mutihop transfer
     */
    function _processMultiHopTransfer(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount,
        uint256 secondaryFee
    ) private returns (uint256, uint256) {
        SpokeBridgeSettings memory spokeSettings =
            registeredSpokes[spokeBlockchainID][spokeBridgeAddress];

        uint256 transferAmount =
            _processReceivedTransfer(spokeSettings, spokeBlockchainID, spokeBridgeAddress, amount);

        uint256 fee = TokenScalingUtils.removeTokenScale(
            spokeSettings.tokenMultiplier, spokeSettings.multiplyOnSpoke, secondaryFee
        );

        return (transferAmount, fee);
    }

    /**
     * @notice Processes a received transfer from a spoke instance.
     * Deducts the balance bridged to the given spoke instance.
     * Removes the token scaling of the spoke, checks the associated hub token
     * amount is greater than zero, and returns the hub token amount.
     * @param spokeSettings The bridge settings for the spoke instance we received the transfer from.
     * @param spokeBlockchainID The blockchain ID of the spoke instance.
     * @param spokeBridgeAddress The address of the spoke instance.
     * @param amount The amount of tokens sent back from spoke, denominated by the
     * spoke's token scale amount.
     */
    function _processReceivedTransfer(
        SpokeBridgeSettings memory spokeSettings,
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount
    ) private returns (uint256) {
        // Require that the spoke is registered and has no collateral needed.
        require(spokeSettings.registered, "TokenHub: spoke not registered");
        require(spokeSettings.collateralNeeded == 0, "TokenHub: spoke not collateralized");

        // Deduct the balance bridged to the given spoke instance prior to scaling the amount.
        _deductSenderBalance(spokeBlockchainID, spokeBridgeAddress, amount);

        // Remove the token scaling of the spoke and get hub token amount.
        uint256 hubAmount = TokenScalingUtils.removeTokenScale(
            spokeSettings.tokenMultiplier, spokeSettings.multiplyOnSpoke, amount
        );

        // Require that the hub token amount is greater than zero after removed scaling.
        require(hubAmount > 0, "TokenHub: zero token amount");
        return hubAmount;
    }

    /**
     * @notice Prepares a multi-hop send by checking the spoke instance settings
     * and adjusting the amount to be sent.
     * @return The scaled amount to be sent to the spoke instance. Zero can be returned if the
     * spoke instance is not registered, needs collateral, or the scaled amount is zero.
     */
    function _prepareMultiHopRouting(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount,
        uint256 fee
    ) private returns (uint256) {
        SpokeBridgeSettings memory spokeSettings =
            registeredSpokes[spokeBlockchainID][spokeBridgeAddress];
        if (!spokeSettings.registered || spokeSettings.collateralNeeded > 0) {
            return 0;
        }

        // Subtract fee amount from amount prior to scaling.
        require(amount > fee, "TokenHub: insufficient amount to cover fees");
        amount -= fee;

        // Scale the amount based on the token multiplier for the given spoke instance.
        uint256 scaledAmount = TokenScalingUtils.applyTokenScale(
            spokeSettings.tokenMultiplier, spokeSettings.multiplyOnSpoke, amount
        );
        if (scaledAmount == 0) {
            return 0;
        }

        // Increase the balance of the spoke instance by the scaled amount.
        bridgedBalances[spokeBlockchainID][spokeBridgeAddress] += scaledAmount;

        return scaledAmount;
    }

    /**
     * @dev Prepares tokens to be sent to another chain by handling the
     * locking of the token amount in this contract and updating the accounting
     * balances.
     */
    function _prepareSend(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount,
        address primaryFeeTokenAddress,
        uint256 feeAmount
    ) private returns (uint256, uint256) {
        SpokeBridgeSettings memory spokeSettings =
            registeredSpokes[spokeBlockchainID][spokeBridgeAddress];
        require(spokeSettings.registered, "TokenHub: spoke not registered");
        require(spokeSettings.collateralNeeded == 0, "TokenHub: collateral needed for spoke");

        // Deposit the funds sent from the user to the bridge,
        // and set to adjusted amount after deposit.
        amount = _deposit(amount);

        if (feeAmount > 0) {
            feeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(primaryFeeTokenAddress), _msgSender(), feeAmount
            );
        }

        // Scale the amount based on the token multiplier for the given spoke instance.
        uint256 scaledAmount = TokenScalingUtils.applyTokenScale(
            spokeSettings.tokenMultiplier, spokeSettings.multiplyOnSpoke, amount
        );
        require(scaledAmount > 0, "TokenHub: zero scaled amount");

        // Increase the balance of the spoke instance by the scaled amount.
        bridgedBalances[spokeBlockchainID][spokeBridgeAddress] += scaledAmount;

        return (scaledAmount, feeAmount);
    }

    function _deductSenderBalance(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount
    ) private {
        uint256 senderBalance = bridgedBalances[spokeBlockchainID][spokeBridgeAddress];
        require(senderBalance >= amount, "TokenHub: insufficient bridge balance");
        bridgedBalances[spokeBlockchainID][spokeBridgeAddress] = senderBalance - amount;
    }

    function _validateSendAndCallInput(SendAndCallInput memory input) private pure {
        require(input.recipientContract != address(0), "TokenHub: zero recipient contract address");
        require(input.requiredGasLimit > 0, "TokenHub: zero required gas limit");
        require(input.recipientGasLimit > 0, "TokenHub: zero recipient gas limit");
        require(
            input.recipientGasLimit < input.requiredGasLimit,
            "TokenHub: invalid recipient gas limit"
        );
        require(input.fallbackRecipient != address(0), "TokenHub: zero fallback recipient address");
        require(input.secondaryFee == 0, "TokenHub: non-zero secondary fee");
    }

    function _validateSendTokensInput(SendTokensInput memory input) private pure {
        require(input.recipient != address(0), "TokenHub: zero recipient address");
        require(input.requiredGasLimit > 0, "TokenHub: zero required gas limit");
        require(input.secondaryFee == 0, "TokenHub: non-zero secondary fee");
    }
}
