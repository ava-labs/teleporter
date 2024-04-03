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
    uint256 requiredGasLimit;
}

/**
 * @notice Interface for a Teleporter token bridge that sends tokens to another chain.
 */
interface ITeleporterTokenBridge is ITeleporterReceiver {
    /**
     * @notice Emitted when tokens are sent to another chain.
     */
    event SendTokens(
        bytes32 indexed teleporterMessageID,
        address indexed sender,
        SendTokensInput input,
        uint256 amount
    );

    /**
     * @notice Emitted when tokens are withdrawn from the token bridge contract.
     */
    event WithdrawTokens(address indexed recipient, uint256 amount);
}
