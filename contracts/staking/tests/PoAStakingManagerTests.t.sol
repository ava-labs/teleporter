// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {StakingManagerTest} from "./StakingManagerTests.t.sol";
import {PoAStakingManager} from "../PoAStakingManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {StakingManagerSettings} from "../interfaces/IStakingManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

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

    function testInvalidOwnerRegistration() public {
        vm.prank(vm.addr(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, vm.addr(1)
            )
        );
        _initializeValidatorRegistration(
            DEFAULT_NODE_ID, DEFAULT_EXPIRY, DEFAULT_ED25519_SIGNATURE, DEFAULT_WEIGHT
        );
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

    // solhint-disable-next-line no-empty-blocks
    function _beforeSend(uint256 value) internal virtual override {}
}
