// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

import {IStakingManager} from "./IStakingManager.sol";

pragma solidity 0.8.18;

interface IRewardCalculator {
    function calculateReward(
        uint64 stakeAmount,
        uint64 startTime,
        uint64 endTime,
        uint256 initialSupply,
        uint256 endSupply
    ) external returns (uint256);
}
