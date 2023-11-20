// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

interface ITokenSource {
    /**
     * @dev Enum representing the action to take on receiving a Teleporter message.
     */
    enum SourceAction {
        Unlock,
        Burn
    }

    /**
     * @dev Emitted when tokens are unlocked on this chain.
     */
    event UnlockTokens(address recipient, uint256 amount);

    /**
     * @dev Emitted when tokens are burned on this chain.
     */
    event BurnTokens(uint256 amount);
}