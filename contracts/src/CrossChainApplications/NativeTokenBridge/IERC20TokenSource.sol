// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for a contract that will lock ERC20 tokens and send a
 * Teleporter message to a {INativeTokenDestination} contract to mint native tokens on that chain.
 */
interface IERC20TokenSource {
    /**
     * @dev Emitted when ERC20 tokens are locked in this source contract to be transferred to the destination chain.
     */
    event TransferToDestination(
        address indexed sender,
        address indexed recipient,
        uint256 transferAmount,
        uint256 feeAmount,
        uint256 teleporterMessageID
    );

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
