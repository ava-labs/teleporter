// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Interface that describes functionalities for a contract that will lock native tokens on the source chain and send a
 * Teleporter message to a {INativeTokenDestination} contract to mint native tokens on the destination chain.
 */
interface INativeTokenSource {
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
