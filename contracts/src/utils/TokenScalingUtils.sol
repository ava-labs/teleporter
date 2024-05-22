// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.18;

library TokenScalingUtils {
    /**
     * @notice Scales the {amount} of hub tokens to a spoke instance's token scale.
     * @param tokenMultiplier The token multiplier of the spoke instance.
     * @param multiplyOnSpoke Whether the amount of hub tokens will be multiplied on the spoke, or divided.
     * @param hubTokenAmount The amount of hub tokens to scale.
     */
    function applyTokenScale(
        uint256 tokenMultiplier,
        bool multiplyOnSpoke,
        uint256 hubTokenAmount
    ) internal pure returns (uint256) {
        return _scaleTokens(tokenMultiplier, multiplyOnSpoke, hubTokenAmount, true);
    }

    /**
     * @notice Removes the spoke instance's token scaling, and returns the corresponding
     * amount of hub tokens.
     * @param tokenMultiplier The token multiplier of the spoke instance.
     * @param multiplyOnSpoke Whether the amount of hub tokens will be multiplied on the spoke, or divided.
     * @param spokeTokenAmount The amount of spoke tokens to remove scaling from.
     */
    function removeTokenScale(
        uint256 tokenMultiplier,
        bool multiplyOnSpoke,
        uint256 spokeTokenAmount
    ) internal pure returns (uint256) {
        return _scaleTokens(tokenMultiplier, multiplyOnSpoke, spokeTokenAmount, false);
    }

    /**
     * @dev Scales {value} based on {tokenMultiplier} and if the amount is applying or
     * removing the spoke instance's token scale.
     * Should be used for all tokens and fees being transferred to/from other subnets.
     * @param tokenMultiplier The token multiplier of the spoke instance.
     * @param multiplyOnSpoke Whether the amount of hub tokens will be multiplied on the spoke, or divided.
     * @param amount The amount of tokens to scale.
     * @param isSendToSpoke If true, indicates the amount is being sent to the
     * spoke instance, so applies token scale. If false, indicates the amount is being
     * sent back to the hub instance, so removes token scale.
     */
    function _scaleTokens(
        uint256 tokenMultiplier,
        bool multiplyOnSpoke,
        uint256 amount,
        bool isSendToSpoke
    ) private pure returns (uint256) {
        // Multiply when multiplyOnSpoke and isSendToSpoke are
        // both true or both false.
        if (multiplyOnSpoke == isSendToSpoke) {
            return amount * tokenMultiplier;
        }
        // Otherwise divide.
        return amount / tokenMultiplier;
    }
}
