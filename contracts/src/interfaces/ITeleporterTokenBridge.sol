// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Parameters for delivery of tokens to another chain and destination recipient.
 * @param destinationBlockchainID blockchainID of the destination
 * @param destinationBridgeAddress address of the destination token bridge instance
 * @param recipient address of the recipient on the destination chain
 * @param primaryFee amount of tokens to pay for Teleporter fee on the source chain
 * @param secondaryFee amount of tokens to pay for Teleporter fee if a multihop is needed
 * @param requiredGasLimit gas limit requirement for sending to a token bridge.
 * This is required because the gas requirement varies based on the token bridge instance
 * specified by `destinationBlockchainID` and `destinationBridgeAddress`.
 */
struct SendTokensInput {
    bytes32 destinationBlockchainID;
    address destinationBridgeAddress;
    address recipient;
    uint256 primaryFee;
    uint256 secondaryFee;
    uint256 requiredGasLimit;
}

/**
 * @notice Parameters for briding tokens to another chain and calling a contract on that chain
 * @param destinationBlockchainID blockchainID of the destination
 * @param destinationBridgeAddress address of the destination token bridge instance
 * @param recipientContract the contract on the destination chain that will be called
 * @param recipientPayload the payload that will be provided to the recipient contract on the destination chain
 * @param requiredGasLimit the required amount of gas needed to deliver the message on its destination chain,
 *                         including token operations and the call the recipient contract.
 * @param recipientGasLimit the amount of gas that will provided to the recipient contract on the destination chain,
 *                          which must be less than the requiredGasLimit of the message as a whole.
 * @param fallbackRecipient address where the bridged tokens are sent if the call to the recipient contract fails.
 * @param primaryFee amount of tokens to pay for Teleporter fee on the source chain
 * @param secondaryFee amount of tokens to pay for Teleporter fee if a multihop is needed
 */
struct SendAndCallInput {
    bytes32 destinationBlockchainID;
    address destinationBridgeAddress;
    address recipientContract;
    bytes recipientPayload;
    uint256 requiredGasLimit;
    uint256 recipientGasLimit;
    address fallbackRecipient;
    uint256 primaryFee;
    uint256 secondaryFee;
}

enum BridgeMessageType {
    SINGLE_HOP_SEND,
    SINGLE_HOP_CALL,
    MULTI_HOP_SEND,
    MULTI_HOP_CALL
}

struct BridgeMessage {
    BridgeMessageType messageType;
    uint256 amount;
    bytes payload;
}

struct SingleHopSendMessage {
    address recipient;
}

struct SingleHopCallMessage {
    address recipientContract;
    bytes recipientPayload;
    uint256 recipientGasLimit;
    address fallbackRecipient;
}

struct MultiHopSendMessage {
    bytes32 destinationBlockchainID;
    address destinationBridgeAddress;
    address recipient;
    uint256 secondaryFee;
    uint256 secondaryGasLimit;
}

struct MultiHopCallMessage {
    bytes32 destinationBlockchainID;
    address destinationBridgeAddress;
    address recipientContract;
    bytes recipientPayload;
    uint256 requiredGasLimit;
    uint256 recipientGasLimit;
    address fallbackRecipient;
    uint256 secondaryFee;
}

/**
 * @notice Interface for a Teleporter token bridge that sends tokens to another chain.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
interface ITeleporterTokenBridge is ITeleporterReceiver {
    /**
     * @notice Emitted when tokens are sent to another chain.
     */
    event TokensSent(
        bytes32 indexed teleporterMessageID,
        address indexed sender,
        SendTokensInput input,
        uint256 amount
    );

    /**
     * @notice Emitted when tokens are sent to another chain with caldata for a contract recipient.
     */
    event TokensAndCallSent(
        bytes32 indexed teleporterMessageID,
        address indexed sender,
        SendAndCallInput input,
        uint256 amount
    );

    /**
     * @notice Emitted when tokens are withdrawn from the token bridge contract.
     */
    event TokensWithdrawn(address indexed recipient, uint256 amount);

    /**
     * @notice Emitted when a call to a recipient contract to receive token succeeds.
     */
    event CallSucceeded(address indexed recipientContract, uint256 amount);

    /**
     * @notice Emitted when a call to a recipient contract to receive token fails, and the tokens are sent
     * to a fallback recipient.
     */
    event CallFailed(address indexed recipientContract, uint256 amount);
}
