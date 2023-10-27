// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for a contract that can mint native tokens when
 * paired with a NativeTokenSource contract that will lock tokens on another chain.
 */
interface INativeTokenDestination {
    /**
     * @dev Emitted when tokens are burned in the destination contract and to be unlocked on the source contract.
     */
    event TransferToSource(
        address indexed sender,
        address indexed recipient,
        uint256 amount,
        uint256 teleporterMessageID
    );

    /**
     * @dev Emitted when tokens are not minted in order to collateralize the source contract.
     */
    event CollateralAdded(uint256 amount, uint256 remaining);

    /**
     * @dev Emitted when minting native tokens.
     */
    event NativeTokensMinted(address indexed recipient, uint256 amount);

    /**
     * @dev This burns native tokens on this chain, and sends a message to the source
     * chain to unlock tokens there.
     */
    function transferToSource(
        address recipient,
        address feeContractAddress,
        uint256 feeAmount,
        address[] calldata allowedRelayerAddresses
    ) external payable;
}
