// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for a cross-chain native token bridge.
 */
interface INativeTokenDestination {
    /**
     * @dev Emitted when tokens are burned in this bridge contract to be bridged to the source chain.
     */
    event TransferToSource(
        address indexed tokenContractAddress,
        uint256 indexed teleporterMessageID,
        bytes32 destinationChainID,
        address destinationBridgeAddress,
        address recipient,
        uint256 transferAmount,
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
        address feeTokenContractAddress,
        uint256 feeAmount
    ) external payable;
}
