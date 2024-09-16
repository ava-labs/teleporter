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
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
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
                    churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                    maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS,
                maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
                rewardCalculator: IRewardCalculator(new ExampleRewardCalculator(DEFAULT_REWARD_RATE))
            })
        );
        validatorManager = app;
        posValidatorManager = app;
    }

    function testCompleteEndValidation() public override {
        // TODO: get native token staking rewards working, then remove this
        // method and let the implementation in PosValidatorManagerTests do the
        // test, and remove the `virtual` modifier from that implementation.
    }

    function testCompleteEndDelegation() public override {
        // TODO: get native token staking rewards working, then remove this
        // method and let the implementation in PosValidatorManagerTests do the
        // test, and remove the `virtual` modifier from that implementation.
    }

    function testZeroMinimumDelegationFee() public {
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(_formatErrorMessage("zero delegation fee"));
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                    maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFeeBips: 0,
                maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
    }

    function testMaxMinimumDelegationFee() public {
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        uint16 minimumDelegationFeeBips = app.MAXIMUM_DELEGATION_FEE_BIPS() + 1;
        vm.expectRevert(_formatErrorMessage("invalid delegation fee"));
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                    maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFeeBips: minimumDelegationFeeBips,
                maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
    }

    function testInvalidStakeAmountRange() public {
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(_formatErrorMessage("invalid stake amount range"));
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                    maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
                }),
                minimumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
                maximumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS,
                maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
    }

    function testZeroMaxStakeMultiplier() public {
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(_formatErrorMessage("zero maximum stake multiplier"));
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                    maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS,
                maximumStakeMultiplier: 0,
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
    }

    function testMaxStakeMultiplierOverLimit() public {
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        uint8 maximumStakeMultiplier = app.MAXIMUM_STAKE_MULTIPLIER_LIMIT() + 1;
        vm.expectRevert(_formatErrorMessage("invalid maximum stake multiplier"));
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                    maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS,
                maximumStakeMultiplier: maximumStakeMultiplier,
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
    }

    // Helpers
    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory registrationInput,
        PoSValidatorRequirements memory requirements,
        uint256 stakeAmount
    ) internal virtual override returns (bytes32) {
        return
            app.initializeValidatorRegistration{value: stakeAmount}(registrationInput, requirements);
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration{value: _weightToValue(weight)}(
            input,
            PoSValidatorRequirements({
                minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                delegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS
            })
        );
    }

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        uint256 value = _weightToValue(weight);
        vm.prank(delegatorAddress);
        vm.deal(delegatorAddress, value);
        return app.initializeDelegatorRegistration{value: value}(validationID);
    }

    function _beforeSend(uint256 amount, address spender) internal override {
        // Native tokens no need pre approve
    }

    function _expectStakeUnlock(address account, uint256 amount) internal override {
        // empty calldata implies the receive function will be called
        vm.expectCall(account, amount, "");
    }

    function _expectRewardIssuance(address account, uint256 amount) internal override {}

    function _getStakeAssetBalance(address account) internal view override returns (uint256) {
        return account.balance;
    }
}
// TODO: Remove this once all unit tests implemented
// solhint-enable no-empty-blocks
