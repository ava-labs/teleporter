// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.0;

/**
 * @dev Interface that describes functionalities for a cross-chain ERC20 bridge.
 */
interface INativeTokenReceiver {

    /**
     * @dev Emitted when tokens are locked in this bridge contract to be bridged to another chain.
     */
    event BridgeTokens(
        address indexed tokenContractAddress,
        uint256 indexed teleporterMessageID,
        bytes32 destinationChainID,
        address destinationBridgeAddress,
        address recipient,
        uint256 transferAmount,
        uint256 feeAmount
    );

    /**
     * @dev Emitted when tokens are unlocked on this chain.
     */
    event UnlockTokens(
        address recipient,
        uint256 amount
    );

    /**
     * @dev Transfers native tokens to another chain as that chain's native token.
     */
    function bridgeTokens(
        address recipient,
        address feeTokenContractAddress,
        uint256 feeAmount
    ) external payable;
}
