// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for a cross-chain native token bridge.
 */
interface INativeTokenDestination {
    error InvalidTeleporterMessengerAddress();
    error InvalidRecipientAddress();
    error InvalidSourceChain();
    error InvalidPartnerContractAddress();
    error CannotBridgeTokenWithinSameChain();
    error Unauthorized();
    error InsufficientPayment();
    error InvalidSourceContractAddress();
    error InsufficientAdjustedAmount(uint256 adjustedAmount, uint256 feeAmount);

    /**
     * @dev Emitted when tokens are burned in this bridge contract to be bridged to the source chain.
     */
    event TransferToSource(
        uint256 indexed teleporterMessageID,
        address recipient,
        uint256 transferAmount,
        address feeContractAddress,
        uint256 feeAmount
    );

    /**
     * @dev Emitted when minting bridge tokens.
     */
    event MintNativeTokens(address recipient, uint256 amount);

    /**
     * @dev This burns native tokens on this chain, and sends a message to the source
     * chain to mint unlock tokens there.
     */
    function transferToSource(
        address recipient,
        address feeContractAddress,
        uint256 feeAmount
    ) external payable;
}
