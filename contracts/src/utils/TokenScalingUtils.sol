// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.18;

library TokenScalingUtils {
    function applyTokenScale(
        uint256 tokenMultiplier,
        bool multiplyOnDestination,
        uint256 amount
    ) internal pure returns (uint256) {
        return _scaleTokens(tokenMultiplier, multiplyOnDestination, amount, true);
    }

    function removeTokenScale(
        uint256 tokenMultiplier,
        bool multiplyOnDestination,
        uint256 amount
    ) internal pure returns (uint256) {
        return _scaleTokens(tokenMultiplier, multiplyOnDestination, amount, false);
    }

    /**
     * @dev Scales `value` based on `tokenMultiplier` and the direction of the transfer.
     * Should be used for all tokens being transferred to/from other subnets.
     */
    function _scaleTokens(
        uint256 tokenMultiplier,
        bool multiplyOnDestination,
        uint256 amount,
        bool isSendToDestination
    ) private pure returns (uint256) {
        // Multiply when multiplyOnDestination and isReceive are both true or both false.
        if (multiplyOnDestination == isSendToDestination) {
            return amount * tokenMultiplier;
        }
        // Otherwise divide.
        return amount / tokenMultiplier;
    }
}
