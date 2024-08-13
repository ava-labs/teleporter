// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {StakingManagerTest} from "./StakingManagerTests.t.sol";
import {PoAStakingManager} from "../PoAStakingManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {StakingManagerSettings} from "../interfaces/IStakingManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";

contract PoAStakingManagerTest is StakingManagerTest {
    PoAStakingManager public app;

    address public constant DEFAULT_OWNER = address(0x1);

    function setUp() public {
        app = new PoAStakingManager(ICMInitializable.Allowed);
        app.initialize(
            StakingManagerSettings({
                pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                subnetID: DEFAULT_SUBNET_ID,
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN,
                rewardCalculator: IRewardCalculator(address(0))
            }),
            address(this)
        );
        stakingManager = app;
    }

    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature,
        uint256 stakeAmount
    ) internal virtual override returns (bytes32) {
        return
            app.initializeValidatorRegistration(stakeAmount, nodeID, registrationExpiry, signature);
    }

    function _beforeSend(uint256 value) internal virtual override {}
}
