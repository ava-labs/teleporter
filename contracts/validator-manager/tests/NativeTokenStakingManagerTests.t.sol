// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoSValidatorManagerTest} from "./PoSValidatorManagerTests.t.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";
import {PoSValidatorManager} from "../PoSValidatorManager.sol";
import {
    ValidatorManagerSettings,
    ValidatorRegistrationInput
} from "../interfaces/IValidatorManager.sol";
import {PoSValidatorManagerSettings} from "../interfaces/IPoSValidatorManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";

contract NativeTokenStakingManagerTest is PoSValidatorManagerTest {
    NativeTokenStakingManager public app;

    function setUp() public virtual {
        // Construct the object under test
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        rewardCalculator = new ExampleRewardCalculator(DEFAULT_REWARD_RATE);
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    subnetID: DEFAULT_SUBNET_ID,
                    churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                    maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
                }),
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS,
                maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
                rewardCalculator: rewardCalculator
            })
        );
        validatorManager = app;
        posValidatorManager = app;
        _mockGetBlockchainID();
        _mockInitializeValidatorSet();
        app.initializeValidatorSet(_defaultSubnetConversionData(), 0);
    }

    function testZeroMinimumDelegationFee() public {
        app = new NativeTokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.InvalidDelegationFee.selector, 0)
        );
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
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
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegationFee.selector, minimumDelegationFeeBips
            )
        );
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
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
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidStakeAmount.selector, DEFAULT_MAXIMUM_STAKE_AMOUNT
            )
        );
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
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
        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.InvalidStakeMultiplier.selector, 0)
        );
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
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
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidStakeMultiplier.selector, maximumStakeMultiplier
            )
        );
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
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
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration{value: stakeAmount}(
            registrationInput, delegationFeeBips, minStakeDuration
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration{value: _weightToValue(weight)}(
            input, DEFAULT_DELEGATION_FEE_BIPS, DEFAULT_MINIMUM_STAKE_DURATION
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

    // solhint-disable no-empty-blocks
    function _beforeSend(uint256 amount, address spender) internal override {
        // Native tokens no need pre approve
    }
    // solhint-enable no-empty-blocks

    function _expectStakeUnlock(address account, uint256 amount) internal override {
        // empty calldata implies the receive function will be called
        vm.expectCall(account, amount, "");
    }

    function _expectRewardIssuance(address account, uint256 amount) internal override {
        vm.mockCall(
            address(app.NATIVE_MINTER()),
            abi.encodeCall(INativeMinter.mintNativeCoin, (account, amount)),
            ""
        );
        // empty calldata implies the receive function will be called:
        vm.mockCall({
            callee: account,
            msgValue: amount,
            data: "", // implies receive()
            returnData: ""
        });
        // Units tests don't have access to the native minter precompile, so use vm.deal instead.
        vm.deal(account, account.balance + amount);
    }

    function _getStakeAssetBalance(address account) internal view override returns (uint256) {
        return account.balance;
    }
}
