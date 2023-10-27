// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

/**
 * @dev Interface that describes functionalities for a contract that will lock tokens and send a
 * Teleporter message to a NativeTokenDestination contract to mint native tokens on that chain.
 */
interface IERC20TokenSource {
    /**
     * @dev Emitted when tokens are locked in this bridge contract to be bridged to the destination chain.
     */
    event TransferToDestination(
        address indexed sender,
        address indexed recipient,
        address ERC20ContractAddress,
        uint256 transferAmount,
        uint256 feeAmount,
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
        address ERC20ContractAddress,
        uint256 totalAmount,
        uint256 feeAmount,
        address[] calldata allowedRelayerAddresses
    ) external;
}
