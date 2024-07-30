// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "forge-std/Test.sol";
import {StakingManagerTest} from "./StakingManagerTests.t.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";
import {StakingManagerSettings, InitialStakerInfo} from "../interfaces/IStakingManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";

contract NativeTokenStakingManagerTest is StakingManagerTest {
    NativeTokenStakingManager public nativeTokenStakingManager;

    function setup() public override {
        StakingManagerTest.setup();
        nativeTokenStakingManager = new NativeTokenStakingManager();
        // TODONOW: Use real values
        nativeTokenStakingManager.initialize(
            StakingManagerSettings({
                pChainBlockchainID: "P-Chain",
                subnetID: "X",
                minimumStakeAmount: 1000,
                maximumStakeAmount: 1000000,
                minimumStakeDuration: 24 hours,
                maximumHourlyChurn: 10,
                initialStakers: new InitialStakerInfo[](0),
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
    }

    function testInitializeValidatorRegistration() public {}
}
