// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterFeeInfo} from "../../Teleporter/ITeleporterMessenger.sol";

/**
 * @dev Interface that describes functionalities for a contract that will lock native tokens on the source chain and send a
 * Teleporter message to a {INativeTokenDestination} contract to mint native tokens on the destination chain.
 */
interface INativeTokenSource {
    /**
     * @dev Emitted when native tokens are locked in the source contract to be transferred to the destination chain.
     */
    event TransferToDestination(
        address indexed sender,
        address indexed recipient,
        uint256 amount,
        uint256 indexed teleporterMessageID
    );

    /**
     * @dev Locks native tokens on the source contract chain, and sends a message to the destination
     * contract to mint corresponding tokens.
     */
    function transferToDestination(
        address recipient,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external payable;
}
