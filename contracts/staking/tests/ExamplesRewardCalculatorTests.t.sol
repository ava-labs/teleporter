// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";

contract ExampleRewardCalculatorTest is Test {
    ExampleRewardCalculator public exampleRewardCalculator;

    uint256 public constant DEFAULT_STAKE_AMOUNT = 1e12;
    uint64 public constant DEFAULT_START_TIME = 1000;
    uint64 public constant DEFAULT_END_TIME = 31537000; // a year + 1000 seonds
    uint64 public constant DEFAULT_REWARD_BASIS_POINTS = 42;

    function setUp() public {
        exampleRewardCalculator = new ExampleRewardCalculator(DEFAULT_REWARD_BASIS_POINTS);
    }

    function testRewardCalculation() public view {
        uint256 output = exampleRewardCalculator.calculateReward(
            DEFAULT_STAKE_AMOUNT, DEFAULT_END_TIME - DEFAULT_START_TIME, 0, 0
        );
        assertEq(output, 42e9);
    }
}
