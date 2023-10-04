// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for a cross-chain native token bridge.
 */
interface INativeTokenSource {
    /**
     * @dev Emitted when tokens are locked in this bridge contract to be bridged to the destination chain.
     */
    event TransferToDestination(
        address indexed sender,
        address indexed recipient,
        uint256 amount,
        address feeContractAddress,
        uint256 feeAmount,
        uint256 teleporterMessageID
    );

    /**
     * @dev Emitted when tokens are unlocked on this chain.
     */
    event UnlockTokens(address recipient, uint256 amount);

    /**
     * @dev Transfers native tokens to the destination chain as that chain's native token.
     */
    function transferToDestination(
        address recipient,
        address feeContractAddress,
        uint256 feeAmount
    ) external payable;
}
