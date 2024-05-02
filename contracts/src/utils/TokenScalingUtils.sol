// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.18;

library TokenScalingUtils {
    /**
     * @dev Scales `value` based on `tokenMultiplier` and the direction of the transfer.
     * Should be used for all tokens being transferred to/from other subnets.
     */
    function scaleTokens(
        uint256 tokenMultiplier,
        bool multiplyOnReceive,
        uint256 amount,
        bool isReceive
    ) internal pure returns (uint256) {
        // Multiply when multiplyOnReceive and isReceive are both true or both false.
        if (multiplyOnReceive == isReceive) {
            return amount * tokenMultiplier;
        }
        // Otherwise divide.
        return amount / tokenMultiplier;
    }
}
