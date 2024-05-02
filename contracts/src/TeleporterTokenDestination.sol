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
 * @title TeleporterTokenDestination
 * @dev Abstract contract for a Teleporter token bridge that receives tokens from a {TeleporterTokenSource} in exchange for the tokens of this token bridge instance.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
abstract contract TeleporterTokenDestination is
    ITeleporterTokenBridge,
    TeleporterOwnerUpgradeable,
    SendReentrancyGuard
{
    /// @notice The blockchain ID of the chain this contract is deployed on.
    bytes32 public immutable blockchainID;

    /// @notice The blockchain ID of the source chain this contract receives tokens from.
    bytes32 public immutable sourceBlockchainID;
    /// @notice The address of the source token bridge instance this contract receives tokens from.
    address public immutable tokenSourceAddress;

    /**
     * @notice tokenMultiplier allows this contract to scale the number of tokens it sends/receives to/from
     * the source chain.
     *
     * @dev This can be used to normalize the number of decimals places between the tokens on
     * the two subnets. Is calculated as 10^d, where d is decimalsShift specified in the constructor.
     */
    uint256 public immutable tokenMultiplier;

    /**
     * @notice If multiplyOnReceive is true, the raw token amount value will be multiplied by `tokenMultiplier` when tokens
     * are transferred from the source chain into this destination chain, and divided by `tokenMultiplier` when
     * tokens are transferred from this destination chain back to the source chain. This is intended
     * when the "decimals" value on the source chain is less than the native EVM denomination of 18.
     * If multiplyOnReceive is false, the raw token amount value will be divided by `tokenMultiplier` when tokens
     * are transferred from the source chain into this destination chain, and multiplied by `tokenMultiplier` when
     * tokens are transferred from this destination chain back to the source chain.
     */
    bool public immutable multiplyOnReceive;

    /**
     * @notice Initial reserve imbalance that must be collateralized on the source before minting.
     *
     * @dev The first `initialReserveImbalance` tokens sent to this subnet will not be minted.
     * `initialReserveImbalance` should be constructed to match the initial token supply of this subnet.
     * This means tokens will not be minted until the source contact is collateralized.
     */
    uint256 public immutable initialReserveImbalance;

    /**
     * @notice Whether or not the contract is known to be fully collateralized. The contract initially
     * starts undercollateralized, and the initialReserveImbalance must be added to the source contract
     * prior to it sending messages to mint tokens.
     */
    bool public isCollateralized;

    /**
     * @notice Fixed gas cost for performing a multi-hop transfer on the `sourceBlockchainID`,
     * before forwarding to the final destination bridge instance.
     */
    uint256 public constant MULTI_HOP_REQUIRED_GAS = 220_000;

    /**
     * @notice The amount of gas added to the required gas limit for a multi-hop call message
     * for each 32-byte word of the recipient payload.
     */
    uint256 public constant MULTI_HOP_CALL_GAS_PER_WORD = 8_500;

    /**
     * @notice Initializes this destination token bridge instance to receive
     * tokens from the specified source blockchain and token bridge instance.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID_,
        address tokenSourceAddress_,
        uint256 initialReserveImbalance_,
        uint8 decimalsShift,
        bool multiplyOnReceive_
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
        require(decimalsShift <= 18, "TeleporterTokenDestination: invalid decimalsShift");
        sourceBlockchainID = sourceBlockchainID_;
        tokenSourceAddress = tokenSourceAddress_;
        initialReserveImbalance = initialReserveImbalance_;
        isCollateralized = initialReserveImbalance_ == 0;
        tokenMultiplier = 10 ** decimalsShift;
        multiplyOnReceive = multiplyOnReceive_;

        // Send a message to the source token bridge instance to register this destination instance.
        RegisterDestinationMessage memory registerMessage = RegisterDestinationMessage({
            initialReserveImbalance: initialReserveImbalance_,
            tokenMultiplier: tokenMultiplier,
            // Invert the value of multiplyOnReceive to match the source contract's perspective.
            multiplyOnReceive: !multiplyOnReceive
        });
        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.REGISTER_DESTINATION,
            payload: abi.encode(registerMessage)
        });
        _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: tokenSourceAddress,
                // TODO: No fee is provided by default when registering a destination, since it
                // is assumed that the deployer will relay the message themselves. However, also
                // need to consider the fee token address such that an amount could be added via addFeeAmount
                // in the future if desired.
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(this), amount: 0}),
                requiredGasLimit: 0,
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
        uint256 numWords = payloadSize / 32;
        if (payloadSize % 32 != 0) {
            ++numWords;
        }
        return numWords;
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
    function _send(SendTokensInput calldata input, uint256 amount) internal sendNonReentrant {
        require(input.recipient != address(0), "TeleporterTokenDestination: zero recipient address");
        require(input.requiredGasLimit > 0, "TeleporterTokenDestination: zero required gas limit");
        amount = _prepareSend(
            input.destinationBlockchainID,
            input.destinationBridgeAddress,
            amount,
            input.primaryFee,
            input.secondaryFee
        );

        // If the destination blockchain is the source blockchain,
        // no multi-hop is needed. Only the required gas limit for the Teleporter message back to
        // `sourceBlockchainID` is needed, which is provided by `input.requiredGasLimit`.
        // Else, there will be a multi-hop transfer to the final destination.
        // The first hop back to `sourceBlockchainID` requires `MULTI_HOP_REQUIRED_GAS`,
        // and the second hop to the final destination requires `input.requiredGasLimit`.
        BridgeMessage memory message;
        uint256 messageRequiredGasLimit = input.requiredGasLimit;
        if (input.destinationBlockchainID == sourceBlockchainID) {
            // If the destination blockchain is the source bridge instance's blockchain,
            // the destination bridge address must match the token source address,
            // and no secondary fee is needed.
            require(
                input.destinationBridgeAddress == tokenSourceAddress,
                "TeleporterTokenDestination: invalid destination bridge address"
            );
            require(input.secondaryFee == 0, "TeleporterTokenDestination: non-zero secondary fee");
            message = BridgeMessage({
                messageType: BridgeMessageType.SINGLE_HOP_SEND,
                payload: abi.encode(SingleHopSendMessage({recipient: input.recipient, amount: amount}))
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
                payload: abi.encode(
                    MultiHopSendMessage({
                        destinationBlockchainID: input.destinationBlockchainID,
                        destinationBridgeAddress: input.destinationBridgeAddress,
                        recipient: input.recipient,
                        amount: amount,
                        secondaryFee: input.secondaryFee,
                        secondaryGasLimit: input.requiredGasLimit,
                        fallbackRecipient: input.fallbackRecipient
                    })
                    )
            });
            messageRequiredGasLimit = MULTI_HOP_REQUIRED_GAS;
        }

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: tokenSourceAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(this), amount: input.primaryFee}),
                requiredGasLimit: messageRequiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensSent(messageID, msg.sender, input, amount);
    }

    /**
     * @notice Sends tokens to the specified recipient contract on the destination chain by
     * calling the {receiveTokens} method of the respective recipient.
     *
     * @dev Burns the bridged amount, and uses Teleporter to send a cross chain message.
     * Tokens and data can be sent to the same blockchain this bridge instance is deployed on,
     * to another destination bridge instance.
     */
    function _sendAndCall(
        SendAndCallInput memory input,
        uint256 amount
    ) internal sendNonReentrant {
        require(
            input.recipientContract != address(0),
            "TeleporterTokenDestination: zero recipient contract address"
        );
        require(input.requiredGasLimit > 0, "TeleporterTokenDestination: zero required gas limit");
        require(input.recipientGasLimit > 0, "TeleporterTokenDestination: zero recipient gas limit");
        require(
            input.recipientGasLimit < input.requiredGasLimit,
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
        uint256 messageRequiredGasLimit = input.requiredGasLimit;
        if (input.destinationBlockchainID == sourceBlockchainID) {
            // If the destination blockchain is the source bridge instance's blockchain,
            // the destination bridge address must match the token source address.
            require(
                input.destinationBridgeAddress == tokenSourceAddress,
                "TeleporterTokenDestination: invalid destination bridge address"
            );

            message = BridgeMessage({
                messageType: BridgeMessageType.SINGLE_HOP_CALL,
                payload: abi.encode(
                    SingleHopCallMessage({
                        sourceBlockchainID: blockchainID,
                        originSenderAddress: msg.sender,
                        recipientContract: input.recipientContract,
                        amount: amount,
                        recipientPayload: input.recipientPayload,
                        recipientGasLimit: input.recipientGasLimit,
                        fallbackRecipient: input.fallbackRecipient
                    })
                    )
            });
        } else {
            // If the destination blockchain ID is this blockchian, the destination
            // bridge address must be a different contract. This is a multi-hop case to
            // a different bridge contract on this chain.
            if (input.destinationBlockchainID == blockchainID) {
                require(
                    input.destinationBridgeAddress != address(this),
                    "TeleporterTokenDestination: invalid destination bridge address"
                );
            }

            message = BridgeMessage({
                messageType: BridgeMessageType.MULTI_HOP_CALL,
                payload: abi.encode(
                    MultiHopCallMessage({
                        originSenderAddress: msg.sender,
                        destinationBlockchainID: input.destinationBlockchainID,
                        destinationBridgeAddress: input.destinationBridgeAddress,
                        recipientContract: input.recipientContract,
                        amount: amount,
                        recipientPayload: input.recipientPayload,
                        recipientGasLimit: input.recipientGasLimit,
                        fallbackRecipient: input.fallbackRecipient,
                        secondaryRequiredGasLimit: input.requiredGasLimit,
                        secondaryFee: input.secondaryFee
                    })
                    )
            });

            // The required gas limit for the first message sent back to the source chain
            // needs to account for the number of words in the payload, which each use additional
            // gas to send in a message to the final destination chain.
            messageRequiredGasLimit = MULTI_HOP_REQUIRED_GAS
                + (calculateNumWords(input.recipientPayload.length) * MULTI_HOP_CALL_GAS_PER_WORD);
        }

        // Send message to the destination bridge address
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: tokenSourceAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(this), amount: input.primaryFee}),
                requiredGasLimit: messageRequiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit TokensAndCallSent(messageID, msg.sender, input, amount);
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
    ) internal override {
        require(
            sourceBlockchainID_ == sourceBlockchainID,
            "TeleporterTokenDestination: invalid source chain"
        );
        require(
            originSenderAddress == tokenSourceAddress,
            "TeleporterTokenDestination: invalid token source address"
        );
        BridgeMessage memory bridgeMessage = abi.decode(message, (BridgeMessage));

        // If the contract was not previously known to be collateralized, it is now given that
        // the source has sent a message to mint funds.
        if (!isCollateralized) {
            isCollateralized = true;
        }

        // Destination contracts should only ever receive single-hop messages because
        // multi-hop messages are always routed through the source contract.
        if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_SEND) {
            SingleHopSendMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopSendMessage));
            _withdraw(payload.recipient, payload.amount);
        } else if (bridgeMessage.messageType == BridgeMessageType.SINGLE_HOP_CALL) {
            SingleHopCallMessage memory payload =
                abi.decode(bridgeMessage.payload, (SingleHopCallMessage));
            _handleSendAndCall(payload, payload.amount);
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
     * deposit, burning, and scaling of the token amount.
     */
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

        // Ensure that the scaled amount on the source chain is non-zero.
        require(
            TokenScalingUtils.scaleTokens(tokenMultiplier, multiplyOnReceive, amount, false) > 0,
            "TeleporterTokenDestination: insufficient tokens to transfer"
        );

        // Return the non-scaled amount to be sent in the message to the source.
        return amount;
    }
}
