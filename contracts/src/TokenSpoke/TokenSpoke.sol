// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenSpoke, TokenSpokeSettings} from "./interfaces/ITokenSpoke.sol";
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
import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {SafeERC20TransferFrom} from "@teleporter/SafeERC20TransferFrom.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/ERC20.sol";
import {SendReentrancyGuard} from "../utils/SendReentrancyGuard.sol";
import {TokenScalingUtils} from "../utils/TokenScalingUtils.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title TokenSpoke
 * @dev Abstract contract for a token bridge spoke that receives tokens from its specified token hub instance, and
 * allows for burning that token to redeem the backing asset on the hub chain, or bridging to other spokes.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
abstract contract TokenSpoke is ITokenSpoke, TeleporterOwnerUpgradeable, SendReentrancyGuard {
    /// @notice The blockchain ID of the chain this contract is deployed on.
    bytes32 public immutable blockchainID;

    /// @notice The blockchain ID of the hub instance this contract receives tokens from.
    bytes32 public immutable tokenHubBlockchainID;

    /// @notice The address of the hub instance this contract receives tokens from on the {tokenHubBlockchainID}.
    address public immutable tokenHubAddress;

    /**
     * @notice tokenMultiplier allows this contract to scale the number of tokens it sends/receives to/from
     * its token hub instance.
     *
     * @dev This can be used to normalize the number of decimals places between the tokens on
     * the two subnets. Is calculated as 10^d, where d is decimalsShift specified in the constructor.
     */
    uint256 public immutable tokenMultiplier;

    /**
     * @notice If {multiplyOnSpoke} is true, the hub token amount will be multiplied by {tokenMultiplier} when tokens
     * are transferred from the hub into this spoke, and divided by {tokenMultiplier} when tokens are transferred from
     * this spoke back to its hub.
     * If {multiplyOnSpoke} is false, the hub token amount value will be divided by {tokenMultiplier} when tokens
     * are transferred from the hub into this spoke, and multiplied by {tokenMultiplier} when tokens are transferred
     * from this spoke back to its hub.
     */
    bool public immutable multiplyOnSpoke;

    /**
     * @notice Initial reserve imbalance that the token for this spoke instance starts with. The hub contract must
     * collateralize a corresonding amount of hub tokens before tokens can be minted on this contract.
     */
    uint256 public immutable initialReserveImbalance;

    /**
     * @notice Whether or not the contract is known to be fully collateralized.
     */
    bool public isCollateralized;

    /**
     * @notice Whether or not the contract is known to be registered with its specified hub contract.
     * This is set to true when the first message is received from the hub contract. Note that {isRegistered}
     * will still be false after the spoke contract is registered on the hub contract until the first
     * message is received back from that contract.
     */
    bool public isRegistered;

    /**
     * @notice Fixed gas cost for performing a multi-hop transfer on the token hub contract,
     * before forwarding to the destination token spoke instance.
     */
    uint256 public constant MULTI_HOP_REQUIRED_GAS = 240_000;

    /**
     * @notice The amount of gas added to the required gas limit for a multi-hop call message
     * for each 32-byte word of the recipient payload.
     */
    uint256 public constant MULTI_HOP_CALL_GAS_PER_WORD = 8_500;

    /**
     * @notice Fixed gas cost for registering the spoke contract on the hub contract.
     */
    uint256 public constant REGISTER_SPOKE_REQUIRED_GAS = 150_000;

    /**
     * @notice Initializes this token spoke instance.
     * @param settings The settings for the token spoke instance.
     * @param initialReserveImbalance_ The initial reserve imbalance that must be collateralized before minting.
     * @param decimalsShift The number of decimal places to shift the token amount by.
     * @param multiplyOnSpoke_ If true, the hub token amount will be multiplied by {tokenMultiplier} when tokens
     * are transferred from the hub into this spoke, and divided by {tokenMultiplier} when tokens are transferred
     * from this spoke back to its hub. If false, the hub token amount value will be divided by {tokenMultiplier}
     * when tokens are transferred from the hub into this spoke, and multiplied by {tokenMultiplier} when tokens
     * are transferred from this spoke back to its hub.
     */
    constructor(
        TokenSpokeSettings memory settings,
        uint256 initialReserveImbalance_,
        uint8 decimalsShift,
        bool multiplyOnSpoke_
    ) TeleporterOwnerUpgradeable(settings.teleporterRegistryAddress, settings.teleporterManager) {
        blockchainID = IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
        require(
            settings.tokenHubBlockchainID != bytes32(0), "TokenSpoke: zero token hub blockchain ID"
        );
        require(
            settings.tokenHubBlockchainID != blockchainID,
            "TokenSpoke: cannot deploy to same blockchain as token hub"
        );
        require(settings.tokenHubAddress != address(0), "TokenSpoke: zero token hub address");
        require(decimalsShift <= 18, "TokenSpoke: invalid decimalsShift");
        tokenHubBlockchainID = settings.tokenHubBlockchainID;
        tokenHubAddress = settings.tokenHubAddress;
        initialReserveImbalance = initialReserveImbalance_;
        isCollateralized = initialReserveImbalance_ == 0;
        tokenMultiplier = 10 ** decimalsShift;
        multiplyOnSpoke = multiplyOnSpoke_;
    }

    /**
     * @notice Sends a message to the contract's specified token hub instance to register this spoke
     * instance. Spoke instances must be registered with their hub contract prior to being able to receive
     * tokens from them.
     */
    function registerWithHub(TeleporterFeeInfo calldata feeInfo) external virtual {
        require(!isRegistered, "TokenSpoke: already registered");

        // Send a message to the token hub instance to register this spoke instance.
        RegisterSpokeMessage memory registerMessage = RegisterSpokeMessage({
            initialReserveImbalance: initialReserveImbalance,
            tokenMultiplier: tokenMultiplier,
            multiplyOnSpoke: multiplyOnSpoke
        });
        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.REGISTER_SPOKE,
            payload: abi.encode(registerMessage)
        });

        uint256 feeAmount = _handleFees(feeInfo.feeTokenAddress, feeInfo.amount);
        _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: tokenHubBlockchainID,
                destinationAddress: tokenHubAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeInfo.feeTokenAddress, amount: feeAmount}),
                requiredGasLimit: REGISTER_SPOKE_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );
    }

    /**
     * @dev Calculates the number of 32-byte words required to fit a payload of a given length.
     * The payloads are padded to have a length that is a multiple of 32.
     */
    function calculateNumWords(uint256 payloadSize) public pure returns (uint256) {
        // Add 31 to effectively round up to the nearest multiple of 32.
        // Right-shift by 5 bits to divide by 32.
        return (payloadSize + 31) >> 5;
    }

    /**
     * @notice Sends tokens to the specified destination.
     *
     * @dev Burns the bridged amount, and uses Teleporter to send a cross chain message to the token hub instance.
     * Tokens can be sent the token hub instance, or to any spoke instance registered with the hub other than this one.
     */
    function _send(SendTokensInput calldata input, uint256 amount) internal sendNonReentrant {
        _validateSendTokensInput(input);
        if (input.destinationBlockchainID == tokenHubBlockchainID) {
            _processSend(input, amount);
        } else {
            _processSendMultiHop(input, amount);
        }
    }

    /**
     * @notice Sends tokens to the specified recipient contract on the destination blockchain ID by
     * calling the {receiveTokens} method of the respective recipient.
     *
     * @dev Burns the bridged amount, and uses Teleporter to send a cross chain message.
     * Tokens and data can be sent to the token hub instance, or to any spoke instance registered with the hub
     * other than this one.
     */
    function _sendAndCall(
        SendAndCallInput calldata input,
        uint256 amount
    ) internal sendNonReentrant {
        _validateSendAndCallInput(input);

        if (input.destinationBlockchainID == tokenHubBlockchainID) {
            return _processSendAndCall(input, amount);
        } else {
            return _processSendAndCallMultiHop(input, amount);
        }
    }

    /**
     * @notice Handles receiving messages from the token hub instance. The supported message types are
     * single-hop sends, and single-hop calls.
     *
     * @dev See {ITeleporterUpgradeable-_receiveTeleporterMessage}
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        require(
            sourceBlockchainID == tokenHubBlockchainID, "TokenSpoke: invalid source blockchain ID"
        );
        require(originSenderAddress == tokenHubAddress, "TokenSpoke: invalid origin sender address");
        BridgeMessage memory bridgeMessage = abi.decode(message, (BridgeMessage));

        // If the contract was not previously known to be registered or collateralized, it is now given that
        // the hub has sent a message to mint funds.
        if (!isRegistered || !isCollateralized) {
            isRegistered = true;
            isCollateralized = true;
        }

        // Spoke contracts should only ever receive single-hop messages because
        // multi-hop messages are always routed through the hub contract.
        if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_SEND) {
            SingleHopSendMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopSendMessage));
            _withdraw(payload.recipient, payload.amount);
        } else if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_CALL) {
            // The {sourceBlockchainID}, and {originSenderAddress} specified in the message
            // payload will not match the sender of this Teleporter message in the case of a
            // multi-hop message. Since Teleporter messages are only received from the specified
            // token hub instance, no additional authentication is needed on the payload values.
            SingleHopCallMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopCallMessage));
            _handleSendAndCall(payload, payload.amount);
        } else {
            revert("TokenSpoke: invalid message type");
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
     * @notice Burns a fee adjusted amount of tokens that the user
     * has deposited to this token bridge instance.
     * @param amount The amount of tokens to burn
     */
    function _burn(uint256 amount) internal virtual;

    /**
     * @notice Processes a send and call message by calling the recipient contract.
     * @param message The send and call message include recipient calldata
     * @param amount The amount of tokens to be sent to the recipient. This amount is assumed to be
     * already scaled to the local denomination of this contract.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal virtual;

    /**
     * @dev Prepares tokens to be sent to another chain by handling the
     * deposit, burning, and checking that the corresonding amount of
     * hub tokens is greater than zero.
     */
    function _prepareSend(
        uint256 amount,
        address primaryFeeTokenAddress,
        uint256 primaryFee,
        uint256 secondaryFee
    ) private returns (uint256, uint256) {
        // Deposit the funds sent from the user to the bridge,
        // and set to adjusted amount after deposit
        amount = _deposit(amount);

        // Transfer the primary fee to pay for fees on the first hop.
        // The user can specify this contract as {primaryFeeTokenAddress},
        // in which case the fee will be paid on top of the bridged amount.
        primaryFee = _handleFees(primaryFeeTokenAddress, primaryFee);

        // Burn the amount of tokens that will be bridged.
        _burn(amount);

        // The bridged amount must cover the secondary fee, because the secondary fee
        // is directly subtracted from the bridged amount on the intermediate (hub) chain
        // performing the multi-hop, before forwarding to the final destination spoke instance.
        require(
            TokenScalingUtils.removeTokenScale(tokenMultiplier, multiplyOnSpoke, amount)
                > TokenScalingUtils.removeTokenScale(tokenMultiplier, multiplyOnSpoke, secondaryFee),
            "TokenSpoke: insufficient tokens to transfer"
        );

        // Return the amount in this contract's local denomination and the primary fee.
        return (amount, primaryFee);
    }

    /**
     * @dev Sends tokens to the specified destination.
     */
    function _processSend(SendTokensInput calldata input, uint256 amount) private {
        _validateSingleHopInput(
            input.destinationBridgeAddress, input.secondaryFee, input.multiHopFallback
        );

        uint256 primaryFee;
        (amount, primaryFee) = _prepareSend({
            amount: amount,
            primaryFeeTokenAddress: input.primaryFeeTokenAddress,
            primaryFee: input.primaryFee,
            secondaryFee: input.secondaryFee
        });

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_SEND,
            payload: abi.encode(SingleHopSendMessage({recipient: input.recipient, amount: amount}))
        });

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: tokenHubBlockchainID,
                destinationAddress: tokenHubAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: primaryFee
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensSent(messageID, _msgSender(), input, amount);
    }

    /**
     * @dev Sends tokens to the specified destination via multi-hop.
     */
    function _processSendMultiHop(SendTokensInput calldata input, uint256 amount) private {
        _validateMultiHopInput(
            input.destinationBlockchainID, input.destinationBridgeAddress, input.multiHopFallback
        );

        uint256 primaryFee;
        (amount, primaryFee) = _prepareSend({
            amount: amount,
            primaryFeeTokenAddress: input.primaryFeeTokenAddress,
            primaryFee: input.primaryFee,
            secondaryFee: input.secondaryFee
        });

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.MULTI_HOP_SEND,
            payload: abi.encode(
                MultiHopSendMessage({
                    destinationBlockchainID: input.destinationBlockchainID,
                    destinationBridgeAddress: input.destinationBridgeAddress,
                    recipient: input.recipient,
                    amount: amount,
                    secondaryFee: input.secondaryFee,
                    secondaryGasLimit: input.requiredGasLimit,
                    multiHopFallback: input.multiHopFallback
                })
                )
        });

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: tokenHubBlockchainID,
                destinationAddress: tokenHubAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: primaryFee
                }),
                requiredGasLimit: MULTI_HOP_REQUIRED_GAS,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensSent(messageID, _msgSender(), input, amount);
    }

    /**
     * @dev Sends tokens to the specified recipient contract on the destination blockchain ID by
     * calling the {receiveTokens} method of the respective recipient.
     */
    function _processSendAndCall(SendAndCallInput calldata input, uint256 amount) private {
        _validateSingleHopInput(
            input.destinationBridgeAddress, input.secondaryFee, input.multiHopFallback
        );

        uint256 primaryFee;
        (amount, primaryFee) = _prepareSend({
            amount: amount,
            primaryFeeTokenAddress: input.primaryFeeTokenAddress,
            primaryFee: input.primaryFee,
            secondaryFee: input.secondaryFee
        });

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_CALL,
            payload: abi.encode(
                SingleHopCallMessage({
                    sourceBlockchainID: blockchainID,
                    originBridgeAddress: address(this),
                    originSenderAddress: _msgSender(),
                    recipientContract: input.recipientContract,
                    amount: amount,
                    recipientPayload: input.recipientPayload,
                    recipientGasLimit: input.recipientGasLimit,
                    fallbackRecipient: input.fallbackRecipient
                })
                )
        });

        // Send message to the token hub instance.
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: tokenHubBlockchainID,
                destinationAddress: tokenHubAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: primaryFee
                }),
                requiredGasLimit: input.requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensAndCallSent(messageID, _msgSender(), input, amount);
    }

    /**
     * @dev Sends tokens to the specified recipient contract on the destination blockchain ID by
     * calling the {receiveTokens} method of the respective recipient.
     */
    function _processSendAndCallMultiHop(SendAndCallInput calldata input, uint256 amount) private {
        _validateMultiHopInput(
            input.destinationBlockchainID, input.destinationBridgeAddress, input.multiHopFallback
        );

        uint256 primaryFee;
        (amount, primaryFee) = _prepareSend({
            amount: amount,
            primaryFeeTokenAddress: input.primaryFeeTokenAddress,
            primaryFee: input.primaryFee,
            secondaryFee: input.secondaryFee
        });

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.MULTI_HOP_CALL,
            payload: abi.encode(
                MultiHopCallMessage({
                    originSenderAddress: _msgSender(),
                    destinationBlockchainID: input.destinationBlockchainID,
                    destinationBridgeAddress: input.destinationBridgeAddress,
                    recipientContract: input.recipientContract,
                    amount: amount,
                    recipientPayload: input.recipientPayload,
                    recipientGasLimit: input.recipientGasLimit,
                    fallbackRecipient: input.fallbackRecipient,
                    multiHopFallback: input.multiHopFallback,
                    secondaryRequiredGasLimit: input.requiredGasLimit,
                    secondaryFee: input.secondaryFee
                })
                )
        });

        // The required gas limit for the first message sent back to the hub instance
        // needs to account for the number of words in the payload. Each word uses additional
        // gas to include in the message to the final destination chain.
        uint256 messageRequiredGasLimit = MULTI_HOP_REQUIRED_GAS
            + (calculateNumWords(input.recipientPayload.length) * MULTI_HOP_CALL_GAS_PER_WORD);

        // Send message to the token hub instance
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: tokenHubBlockchainID,
                destinationAddress: tokenHubAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: primaryFee
                }),
                requiredGasLimit: messageRequiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensAndCallSent(messageID, _msgSender(), input, amount);
    }

    /**
     * @notice Handles fees sent to this contract
     * @param feeTokenAddress The address of the fee token
     * @param feeAmount The amount of the fee
     */
    function _handleFees(address feeTokenAddress, uint256 feeAmount) private returns (uint256) {
        if (feeAmount == 0) {
            return 0;
        }
        // If the {feeTokenAddress} is this contract, then just deposit the tokens directly.
        if (feeTokenAddress == address(this)) {
            return _deposit(feeAmount);
        }
        return SafeERC20TransferFrom.safeTransferFrom(IERC20(feeTokenAddress), feeAmount);
    }

    function _validateSingleHopInput(
        address destinationBridgeAddress,
        uint256 secondaryFee,
        address multiHopFallback
    ) private view {
        require(
            destinationBridgeAddress == tokenHubAddress,
            "TokenSpoke: invalid destination bridge address"
        );
        require(secondaryFee == 0, "TokenSpoke: non-zero secondary fee");
        require(multiHopFallback == address(0), "TokenSpoke: non-zero multi-hop fallback");
    }

    function _validateMultiHopInput(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address multiHopFallback
    ) private view {
        // If the destination blockchain ID is this blockchain, the destination
        // bridge address must be a different contract. This is a multi-hop case to
        // a different bridge contract on this chain.
        if (destinationBlockchainID == blockchainID) {
            require(
                destinationBridgeAddress != address(this),
                "TokenSpoke: invalid destination bridge address"
            );
        }
        require(multiHopFallback != address(0), "TokenSpoke: zero multi-hop fallback");
    }

    function _validateSendTokensInput(SendTokensInput calldata input) private pure {
        require(input.recipient != address(0), "TokenSpoke: zero recipient address");
        require(input.requiredGasLimit > 0, "TokenSpoke: zero required gas limit");
        require(
            input.destinationBlockchainID != bytes32(0),
            "TokenSpoke: zero destination blockchain ID"
        );
        require(
            input.destinationBridgeAddress != address(0),
            "TokenSpoke: zero destination bridge address"
        );
    }

    function _validateSendAndCallInput(SendAndCallInput calldata input) private pure {
        require(
            input.destinationBlockchainID != bytes32(0),
            "TokenSpoke: zero destination blockchain ID"
        );
        require(
            input.destinationBridgeAddress != address(0),
            "TokenSpoke: zero destination bridge address"
        );
        require(
            input.recipientContract != address(0), "TokenSpoke: zero recipient contract address"
        );
        require(input.requiredGasLimit > 0, "TokenSpoke: zero required gas limit");
        require(input.recipientGasLimit > 0, "TokenSpoke: zero recipient gas limit");
        require(
            input.recipientGasLimit < input.requiredGasLimit,
            "TokenSpoke: invalid recipient gas limit"
        );
        require(
            input.fallbackRecipient != address(0), "TokenSpoke: zero fallback recipient address"
        );
    }
}
