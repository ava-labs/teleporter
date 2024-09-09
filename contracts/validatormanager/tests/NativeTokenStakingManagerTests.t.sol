// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoSValidatorManagerTest} from "./PoSValidatorManagerTests.t.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";
import {
    ValidatorManagerSettings,
    ValidatorRegistrationInput
} from "../interfaces/IValidatorManager.sol";
import {
    PoSValidatorManagerSettings,
    PoSValidatorRequirements
} from "../interfaces/IPoSValidatorManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
contract NativeTokenStakingManagerTest is PoSValidatorManagerTest {
    NativeTokenStakingManager public app;

    function setUp() public virtual {
        // Construct the object under test
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFee: DEFAULT_MINIMUM_DELEGATION_FEE,
                maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
        validatorManager = app;
        posValidatorManager = app;
    }

    // Helpers
    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration{value: app.weightToValue(weight)}(
            input,
            PoSValidatorRequirements({
                minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                delegationFee: DEFAULT_MINIMUM_DELEGATION_FEE
            })
        );
    }

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        uint256 value = app.weightToValue(weight);
        vm.prank(delegatorAddress);
        vm.deal(delegatorAddress, value);
        return app.initializeDelegatorRegistration{value: value}(validationID);
    }

    function _beforeSend(uint64 weight, address spender) internal override {
        // Native tokens no need pre approve
    }
}
// TODO: Remove this once all unit tests implemented
// solhint-enable no-empty-blocks
