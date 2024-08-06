// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";

contract ExampleRewardCalculator is IRewardCalculator {
    uint256 public constant SECONDS_IN_YEAR = 31536000;

    function calculateReward(
        uint64 stakeAmount,
        uint64 startTime,
        uint64 endTime,
        uint256, // initialSupply
        uint256, // endSupply
        uint256 rewardBasisPoints
    ) external pure returns (uint256) {
        return stakeAmount * rewardBasisPoints * (endTime - startTime) / SECONDS_IN_YEAR / 1000;
    }
}
