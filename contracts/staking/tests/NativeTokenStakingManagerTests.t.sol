// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {StakingManagerTest} from "./StakingManagerTests.t.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";
import {StakingManagerSettings} from "../interfaces/IStakingManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
contract NativeTokenStakingManagerTest is StakingManagerTest {
    NativeTokenStakingManager public app;

    function setUp() public virtual {
        // Construct the object under test
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        app.initialize(
            StakingManagerSettings({
                pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                subnetID: DEFAULT_SUBNET_ID,
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN,
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
        stakingManager = app;
    }

    // Helpers
    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature,
        uint256 stakeAmount
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration{value: stakeAmount}(
            nodeID, registrationExpiry, signature
        );
    }

    function _beforeSend(uint256 value) internal override {
        // Native tokens no need pre approve
    }
}
// TODO: Remove this once all unit tests implemented
// solhint-enable no-empty-blocks
