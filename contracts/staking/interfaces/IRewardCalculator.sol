// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

interface IRewardCalculator {
    function calculateReward(
        uint256 stakeAmount,
        uint64 validatorStartTime,
        uint64 stakingStartTime,
        uint64 stakingEndTime,
        uint64 uptimeSeconds,
        uint256 initialSupply,
        uint256 endSupply
    ) external view returns (uint256);
}
