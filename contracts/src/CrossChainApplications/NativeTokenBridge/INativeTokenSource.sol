// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for a contract that will lock tokens and send a
 * Teleporter message to a {INativeTokenDestination} contract to mint native tokens on the destination chain.
 */
interface INativeTokenSource {
    /**
     * @dev Emitted when tokens are locked in the source contract to be transferred to the destination chain.
     */
    event TransferToDestination(
        address indexed sender,
        address indexed recipient,
        uint256 amount,
        uint256 teleporterMessageID
    );

    /**
     * @dev Emitted when tokens are unlocked on this chain.
     */
    event UnlockTokens(address recipient, uint256 amount);

    /**
     * @dev Transfers source chain native tokens to destination chain's native tokens.
     */
    function transferToDestination(
        address recipient,
        address feeContractAddress,
        uint256 feeAmount,
        address[] calldata allowedRelayerAddresses
    ) external payable;
}
