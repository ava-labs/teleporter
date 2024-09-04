// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {TeleporterRegistryOwnableApp} from "../teleporter/registry/TeleporterRegistryOwnableApp.sol";
import {TeleporterMessageInput, TeleporterFeeInfo} from "../teleporter/ITeleporterMessenger.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Sequential messages must be received in incrementing nonce order.
 */
struct SequentialMessage {
    uint256 nonce;
    bytes payload;
}

/**
 * @notice This contract is meant as an example to show how a cross-chain application could enforce the
 * requirement of in-order message execution.
 */
contract SequentialDeliverAppExample is TeleporterRegistryOwnableApp {
    // Blockchain ID and contract address of the partner app to send/receive messages to/from.
    bytes32 public immutable partnerAppBlockchainID;
    address public immutable partnerAppAddress;

    // Latest nonces sent and received.
    uint256 public latestSentMessageNonce;
    uint256 public latestReceivedMessageNonce;

    // Amount of gas needed to execute the message on the destination chain and contract.
    uint256 public constant DEFAULT_MESSAGE_REQUIRED_GAS_LIMIT = 75_000;

    event SequentialMessageSent(uint256 nonce, bytes payload);
    event SequentialMessageReceived(uint256 nonce, bytes payload);

    constructor(
        address teleporterRegistryAddress,
        address initialOwner,
        uint256 minTeleporterVersion,
        bytes32 partnerAppBlockchainID_,
        address partnerAppAddress_
    ) TeleporterRegistryOwnableApp(teleporterRegistryAddress, initialOwner, minTeleporterVersion) {
        partnerAppBlockchainID = partnerAppBlockchainID_;
        partnerAppAddress = partnerAppAddress_;
    }

    /**
     * @notice Example external function to send a message to the partner app. All messages sent
     * include a nonce that is incremented by one from the previous message.
     * @param payload The payload to send to the partner app.
     */
    function sendMessage(bytes memory payload) external {
        _sendSequentialMessage(payload, DEFAULT_MESSAGE_REQUIRED_GAS_LIMIT);
    }

    /**
     * @notice Required messages to be received in order. This function is called by the TeleporterMessenger
     * contract when a message is received from the partner app.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        // Only allow for receiving messages from the partner app.
        require(
            sourceBlockchainID == partnerAppBlockchainID,
            "SequentialDeliverAppExample: Invalid source blockchain ID"
        );
        require(
            originSenderAddress == partnerAppAddress,
            "SequentialDeliverAppExample: Invalid origin sender address"
        );

        // Require the message nonce to be sequential.
        // In the case of a message being delivered out of order, the execution of the message will revert.
        // It can later be re-executed once all messages with lower nonces have been executed. This would be
        // done by calling {TeleporterMessenger-retryMessageExecution}.
        SequentialMessage memory sequentialMessage = abi.decode(message, (SequentialMessage));
        require(
            sequentialMessage.nonce == latestReceivedMessageNonce + 1,
            "SequentialDeliverAppExample: Invalid nonce"
        );

        // Applications would add their own logic here to process the message payload.
        emit SequentialMessageReceived(sequentialMessage.nonce, sequentialMessage.payload);

        // Update the latest received nonce.
        latestReceivedMessageNonce = sequentialMessage.nonce;
    }

    /**
     * @notice Adds the sequential message to the TeleporterMessenger to be sent to the partner app.
     */
    function _sendSequentialMessage(bytes memory payload, uint256 requiredGasLimit) internal {
        // Increment the latest sent nonce.
        uint256 nonce = ++latestSentMessageNonce;

        // Encode the message with its nonce and send it.
        bytes memory message = abi.encode(SequentialMessage({nonce: nonce, payload: payload}));
        _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: partnerAppBlockchainID,
                destinationAddress: partnerAppAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
                requiredGasLimit: requiredGasLimit,
                allowedRelayerAddresses: new address[](0), // any relayer allowed
                message: message
            })
        );

        emit SequentialMessageSent(nonce, payload);
    }
}
