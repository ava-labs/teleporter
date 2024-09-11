// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";

contract ExampleRewardCalculator is IRewardCalculator {
    uint256 public constant SECONDS_IN_YEAR = 31536000;

    uint64 public immutable rewardBasisPoints;

    constructor(uint64 rewardBasisPoints_) {
        rewardBasisPoints = rewardBasisPoints_;
    }

    /**
     * @notice A linear, non-compounding reward calculation that rewards a set percentage of tokens per year.
     */
    function calculateReward(
        uint256 stakeAmount,
        uint64 uptime,
        uint256, // initialSupply
        uint256 // endSupply
    ) external view returns (uint256) {
        return (stakeAmount * rewardBasisPoints * uptime) / SECONDS_IN_YEAR / 1000;
    }
}
