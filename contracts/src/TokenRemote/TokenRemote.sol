// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenRemote, TokenRemoteSettings} from "./interfaces/ITokenRemote.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    BridgeMessageType,
    BridgeMessage,
    SingleHopSendMessage,
    SingleHopCallMessage,
    MultiHopSendMessage,
    MultiHopCallMessage,
    RegisterRemoteMessage
} from "../interfaces/ITokenBridge.sol";
import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {SendReentrancyGuard} from "../utils/SendReentrancyGuard.sol";
import {TokenScalingUtils} from "../utils/TokenScalingUtils.sol";

/**
 * @title TokenRemote
 * @dev Abstract contract for a token bridge remote that receives tokens from its specified token TokenHome instance, and
 * allows for burning that token to redeem the backing asset on the home chain, or bridging to other remotes.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
abstract contract TokenRemote is ITokenRemote, TeleporterOwnerUpgradeable, SendReentrancyGuard {
    /// @notice The blockchain ID of the chain this contract is deployed on.
    bytes32 public immutable blockchainID;

    /// @notice The blockchain ID of the TokenHome instance this contract receives tokens from.
    bytes32 public immutable tokenHomeBlockchainID;

    /// @notice The address of the TokenHome instance this contract receives tokens from on the {tokenHomeBlockchainID}.
    address public immutable tokenHomeAddress;

    /**
     * @notice The number of decimal places in the denomination of the home
     * token.
     * @dev Used to derive tokenMultiplier and multiplyOnRemote.
     */
    uint8 public immutable homeTokenDecimals;

    /**
     * @notice The number of decimal places in the denomination of the remote token.
     * @dev Used to derive tokenMultiplier and multiplyOnRemote.
     */
    uint8 public immutable tokenDecimals;

    /**
     * @notice tokenMultiplier allows this contract to scale the number of tokens it sends/receives to/from
     * its token TokenHome instance.
     *
     * @dev This is used to normalize the number of decimal places between the token home asset and the
     * token remote asset. It is derived from home and remote token decimals values passed in the constructor.
     */
    uint256 public immutable tokenMultiplier;

    /**
     * @notice If {multiplyOnRemote} is true, the home token amount will be multiplied by {tokenMultiplier} when tokens
     * are transferred from the home into this remote, and divided by {tokenMultiplier} when tokens are transferred from
     * this remote back to its home.
     * If {multiplyOnRemote} is false, the home token amount value will be divided by {tokenMultiplier} when tokens
     * are transferred from the home into this remote, and multiplied by {tokenMultiplier} when tokens are transferred
     * from this remote back to its home.
     */
    bool public immutable multiplyOnRemote;

    /**
     * @notice Initial reserve imbalance that the token for this TokenRemote instance starts with. The home contract must
     * collateralize a corresonding amount of home tokens before tokens can be minted on this contract.
     */
    uint256 public immutable initialReserveImbalance;

    /**
     * @notice Whether or not the contract is known to be fully collateralized.
     */
    bool public isCollateralized;

    /**
     * @notice Whether or not the contract is known to be registered with its specified home contract.
     * This is set to true when the first message is received from the home contract. Note that {isRegistered}
     * will still be false after the remote contract is registered on the home contract until the first
     * message is received back from that contract.
     */
    bool public isRegistered;

    /**
     * @notice Fixed gas cost for executing a multi-hop send message on the token home contract,
     * before forwarding to the destination token TokenRemote instance. Note that for certain home implementations,
     * a higher required gas limit may be needed depending on the gas expenditure of the external call to the
     * token contract to approve the spending of the optional secondary fee amount.
     */
    uint256 public constant MULTI_HOP_SEND_REQUIRED_GAS = 340_000;

    /**
     * @notice Fixed gas cost for executing a multi-hop call message on the token home contract,
     * before forwarding to the destination token TokenRemote instance. Note that for certain home implementations,
     * a higher required gas limit may be needed depending on the gas expenditure of the external call to the
     * token contract to approve the spending of the optional secondary fee amount.
     */
    uint256 public constant MULTI_HOP_CALL_REQUIRED_GAS = 350_000;

    /**
     * @notice The amount of gas added to the required gas limit for a multi-hop call message
     * for each 32-byte word of the recipient payload.
     */
    uint256 public constant MULTI_HOP_CALL_GAS_PER_WORD = 1_500;

    /**
     * @notice Fixed gas cost for registering the remote contract on the home contract.
     */
    uint256 public constant REGISTER_REMOTE_REQUIRED_GAS = 115_000;

    /**
     * @notice Initializes this token TokenRemote instance.
     * @param settings The settings for the token TokenRemote instance.
     * @param initialReserveImbalance_ The initial reserve imbalance that must be collateralized before minting.
     * @param tokenDecimals_ The number of decimal places in the denomination of the remote token.
     */
    constructor(
        TokenRemoteSettings memory settings,
        uint256 initialReserveImbalance_,
        uint8 tokenDecimals_
    ) TeleporterOwnerUpgradeable(settings.teleporterRegistryAddress, settings.teleporterManager) {
        blockchainID = IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
        require(
            settings.tokenHomeBlockchainID != bytes32(0),
            "TokenRemote: zero token home blockchain ID"
        );
        require(
            settings.tokenHomeBlockchainID != blockchainID,
            "TokenRemote: cannot deploy to same blockchain as token home"
        );
        require(settings.tokenHomeAddress != address(0), "TokenRemote: zero token home address");
        require(
            settings.tokenHomeDecimals <= TokenScalingUtils.MAX_TOKEN_DECIMALS,
            "TokenRemote: token home decimals too high"
        );
        require(
            tokenDecimals_ <= TokenScalingUtils.MAX_TOKEN_DECIMALS,
            "TokenRemote: token decimals too high"
        );
        tokenHomeBlockchainID = settings.tokenHomeBlockchainID;
        tokenHomeAddress = settings.tokenHomeAddress;
        initialReserveImbalance = initialReserveImbalance_;
        isCollateralized = initialReserveImbalance_ == 0;
        homeTokenDecimals = settings.tokenHomeDecimals;
        tokenDecimals = tokenDecimals_;
        (tokenMultiplier, multiplyOnRemote) =
            TokenScalingUtils.deriveTokenMultiplierValues(homeTokenDecimals, tokenDecimals);
    }

    /**
     * @notice Sends a message to the contract's specified token TokenHome instance to register this remote
     * instance. TokenRemote instances must be registered with their home contract prior to being able to receive
     * tokens from them.
     */
    function registerWithHome(TeleporterFeeInfo calldata feeInfo) external virtual {
        require(!isRegistered, "TokenRemote: already registered");

        // Send a message to the token TokenHome instance to register this TokenRemote instance.
        RegisterRemoteMessage memory registerMessage = RegisterRemoteMessage({
            initialReserveImbalance: initialReserveImbalance,
            homeTokenDecimals: homeTokenDecimals,
            remoteTokenDecimals: tokenDecimals
        });
        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.REGISTER_REMOTE,
            payload: abi.encode(registerMessage)
        });

        uint256 feeAmount = _handleFees(feeInfo.feeTokenAddress, feeInfo.amount);
        _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: tokenHomeBlockchainID,
                destinationAddress: tokenHomeAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeInfo.feeTokenAddress, amount: feeAmount}),
                requiredGasLimit: REGISTER_REMOTE_REQUIRED_GAS,
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
     * @dev Burns the bridged amount, and uses Teleporter to send a cross chain message to the token TokenHome instance.
     * Tokens can be sent the token TokenHome instance, or to any TokenRemote instance registered with the home other than this one.
     */
    function _send(SendTokensInput calldata input, uint256 amount) internal sendNonReentrant {
        _validateSendTokensInput(input);
        if (input.destinationBlockchainID == tokenHomeBlockchainID) {
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
     * Tokens and data can be sent to the token TokenHome instance, or to any TokenRemote instance registered with the home
     * other than this one.
     */
    function _sendAndCall(
        SendAndCallInput calldata input,
        uint256 amount
    ) internal sendNonReentrant {
        _validateSendAndCallInput(input);

        if (input.destinationBlockchainID == tokenHomeBlockchainID) {
            return _processSendAndCall(input, amount);
        } else {
            return _processSendAndCallMultiHop(input, amount);
        }
    }

    /**
     * @notice Handles receiving messages from the token TokenHome instance. The supported message types are
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
            sourceBlockchainID == tokenHomeBlockchainID, "TokenRemote: invalid source blockchain ID"
        );
        require(
            originSenderAddress == tokenHomeAddress, "TokenRemote: invalid origin sender address"
        );
        BridgeMessage memory bridgeMessage = abi.decode(message, (BridgeMessage));

        // If the contract was not previously known to be registered or collateralized, it is now given that
        // the home has sent a message to mint funds.
        if (!isRegistered || !isCollateralized) {
            isRegistered = true;
            isCollateralized = true;
        }

        // Remote contracts should only ever receive single-hop messages because
        // multi-hop messages are always routed through the home contract.
        if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_SEND) {
            SingleHopSendMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopSendMessage));
            _withdraw(payload.recipient, payload.amount);
        } else if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_CALL) {
            // The {sourceBlockchainID}, and {originSenderAddress} specified in the message
            // payload will not match the sender of this Teleporter message in the case of a
            // multi-hop message. Since Teleporter messages are only received from the specified
            // token TokenHome instance, no additional authentication is needed on the payload values.
            SingleHopCallMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopCallMessage));
            _handleSendAndCall(payload, payload.amount);
        } else {
            revert("TokenRemote: invalid message type");
        }
    }

    /**
     * @notice Withdraws tokens to the recipient address.
     * @param recipient The address to withdraw tokens to
     * @param amount The amount of tokens to withdraw
     */
    function _withdraw(address recipient, uint256 amount) internal virtual;

    /**
     * @notice Burns the user's tokens to initiate a bridge transfer.
     * @param amount The amount of tokens to burn
     * @return The amount of tokens burned, which is the amount to credit
     * for the bridge transfer.
     */
    function _burn(uint256 amount) internal virtual returns (uint256);

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
     * @notice Handles fees sent to this contract for a bridge transfer.
     * The fee is expected to be approved by the sender for this bridge contract to use,
     * and will be transferred to this contract to allocate for the bridge transfer.
     * @param feeTokenAddress The address of the fee token
     * @param feeAmount The amount of the fee
     */
    function _handleFees(
        address feeTokenAddress,
        uint256 feeAmount
    ) internal virtual returns (uint256);

    /**
     * @dev Prepares tokens to be sent to another chain by handling the
     * deposit, burning, and checking that the corresonding amount of
     * home tokens is greater than zero.
     */
    function _prepareSend(
        uint256 amount,
        address primaryFeeTokenAddress,
        uint256 primaryFee,
        uint256 secondaryFee
    ) private returns (uint256, uint256) {
        // Transfer the primary fee to pay for fees on the first hop.
        // The user can specify this contract as {primaryFeeTokenAddress},
        // in which case the fee will be paid on top of the bridged amount.
        primaryFee = _handleFees(primaryFeeTokenAddress, primaryFee);

        // Burn the amount of tokens that will be bridged.
        amount = _burn(amount);

        // The bridged amount must cover the secondary fee, because the secondary fee
        // is directly subtracted from the bridged amount on the intermediate (home) chain
        // performing the multi-hop, before forwarding to the final destination TokenRemote instance.
        require(
            TokenScalingUtils.removeTokenScale(tokenMultiplier, multiplyOnRemote, amount)
                > TokenScalingUtils.removeTokenScale(tokenMultiplier, multiplyOnRemote, secondaryFee),
            "TokenRemote: insufficient tokens to transfer"
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
                destinationBlockchainID: tokenHomeBlockchainID,
                destinationAddress: tokenHomeAddress,
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
                destinationBlockchainID: tokenHomeBlockchainID,
                destinationAddress: tokenHomeAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: input.primaryFeeTokenAddress,
                    amount: primaryFee
                }),
                requiredGasLimit: MULTI_HOP_SEND_REQUIRED_GAS,
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

        // Send message to the token TokenHome instance.
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: tokenHomeBlockchainID,
                destinationAddress: tokenHomeAddress,
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

        // The required gas limit for the first message sent back to the TokenHome instance
        // needs to account for the number of words in the payload. Each word uses additional
        // gas to include in the message to the final destination chain.
        uint256 messageRequiredGasLimit = MULTI_HOP_CALL_REQUIRED_GAS
            + (calculateNumWords(input.recipientPayload.length) * MULTI_HOP_CALL_GAS_PER_WORD);

        // Send message to the token TokenHome instance
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: tokenHomeBlockchainID,
                destinationAddress: tokenHomeAddress,
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

    function _validateSingleHopInput(
        address destinationBridgeAddress,
        uint256 secondaryFee,
        address multiHopFallback
    ) private view {
        require(
            destinationBridgeAddress == tokenHomeAddress,
            "TokenRemote: invalid destination bridge address"
        );
        require(secondaryFee == 0, "TokenRemote: non-zero secondary fee");
        require(multiHopFallback == address(0), "TokenRemote: non-zero multi-hop fallback");
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
                "TokenRemote: invalid destination bridge address"
            );
        }
        require(multiHopFallback != address(0), "TokenRemote: zero multi-hop fallback");
    }

    function _validateSendTokensInput(SendTokensInput calldata input) private pure {
        require(input.recipient != address(0), "TokenRemote: zero recipient address");
        require(input.requiredGasLimit > 0, "TokenRemote: zero required gas limit");
        require(
            input.destinationBlockchainID != bytes32(0),
            "TokenRemote: zero destination blockchain ID"
        );
        require(
            input.destinationBridgeAddress != address(0),
            "TokenRemote: zero destination bridge address"
        );
    }

    function _validateSendAndCallInput(SendAndCallInput calldata input) private pure {
        require(
            input.destinationBlockchainID != bytes32(0),
            "TokenRemote: zero destination blockchain ID"
        );
        require(
            input.destinationBridgeAddress != address(0),
            "TokenRemote: zero destination bridge address"
        );
        require(
            input.recipientContract != address(0), "TokenRemote: zero recipient contract address"
        );
        require(input.requiredGasLimit > 0, "TokenRemote: zero required gas limit");
        require(input.recipientGasLimit > 0, "TokenRemote: zero recipient gas limit");
        require(
            input.recipientGasLimit < input.requiredGasLimit,
            "TokenRemote: invalid recipient gas limit"
        );
        require(
            input.fallbackRecipient != address(0), "TokenRemote: zero fallback recipient address"
        );
    }
}
