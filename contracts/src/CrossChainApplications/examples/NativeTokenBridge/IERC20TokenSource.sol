// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Interface that describes functionalities for a contract that will lock ERC20 tokens and send a
 * Teleporter message to a {INativeTokenDestination} contract to mint native tokens on the destination chain.
 */
interface IERC20TokenSource {
    /**
     * @dev Locks ERC20 tokens on the source contract chain, and sends a message to mint the
     * corresponding native tokens on the destination chain.
     */
    function transferToDestination(
        address recipient,
        uint256 totalAmount,
        uint256 feeAmount,
        address[] calldata allowedRelayerAddresses
    ) external;
}
