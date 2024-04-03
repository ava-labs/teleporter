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
 * @param secondaryFee amount of tokens to pay for Teleporter fee if a multihop is needed.
 * @param allowedRelayerAddresses addresses of relayers allowed to send the message
 */
struct SendTokensInput {
    bytes32 destinationBlockchainID;
    address destinationBridgeAddress;
    address recipient;
    uint256 primaryFee;
    uint256 secondaryFee;
    address[] allowedRelayerAddresses;
}

/**
 * @notice Parameters for briding tokens to another chain and calling a contract on that chain
 * @param destinationBlockchainID blockchainID of the destination
 * @param destinationBridgeAddress address of the destination token bridge instance
 * @param recipientContract the contract on the destination chain that will be called
 * @param recipientPayload the payload that will be provided to the recipient contract on the destination chain
 * @param recipientGasLimit the amount of gas that will provided to the recipient contract on the destination chain
 * @param fallbackRecipient address where the bridged tokens are sent if the call to the recipient contract fails.
 * @param primaryFee amount of tokens to pay for Teleporter fee on the source chain
 * @param secondaryFee amount of tokens to pay for Teleporter fee if a multihop is needed
 * @param allowedRelayerAddresses addresses of relayers allowed to send the message
 */
struct SendAndCallInput {
    bytes32 destinationBlockchainID;
    address destinationBridgeAddress;
    address recipientContract;
    bytes recipientPayload;
    uint256 recipientGasLimit;
    address fallbackRecipient;
    uint256 primaryFee;
    uint256 secondaryFee;
    address[] allowedRelayerAddresses;
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
}

struct MultiHopCallMessage {
    bytes32 destinationBlockchainID;
    address destinationBridgeAddress;
    address recipientContract;
    bytes recipientPayload;
    uint256 recipientGasLimit;
    address fallbackRecipient;
    uint256 secondaryFee;
}

/**
 * @notice Interface for a Teleporter token bridge that sends tokens to another chain.
 */
interface ITeleporterTokenBridge is ITeleporterReceiver {
    /**
     * @notice Emitted when tokens are sent to another chain.
     * TODO: might want to add SendTokensInput as a parameter
     */
    event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, uint256 amount);
}
