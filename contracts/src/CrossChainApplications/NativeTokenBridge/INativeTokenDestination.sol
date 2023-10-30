// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "../../Teleporter/ITeleporterMessenger.sol";

/**
 * @dev Interface that describes functionalities for a contract that can mint native tokens when
 * paired with a {INativeTokenSource} contract that will lock tokens on another chain.
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
     * @dev Burns native tokens on the destination contract chain, and sends a message to the source
     * contract to unlock corresponding tokens.
     */
    function transferToSource(
        address recipient,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external payable;

    /**
     * @dev Returns true if currentTokenImbalance is 0. When this is true, all tokens sent to
     * this chain will be minted, and sending tokens to the source chain is allowed.
     */
    function isCollateralized() external view returns (bool);

    /**
     * @dev Returns the total number of tokens minted + initialReserveImbalance -
     * total tokens burned through fees and transferring back to the source chain.
     */
    function totalSupply() external view returns (uint256);
}

