// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {PoSSecurityModuleTest} from "./PoSSecurityModuleTests.t.sol";
import {NativeTokenSecurityModule} from "../NativeTokenSecurityModule.sol";
import {PoSSecurityModule, PoSSecurityModuleSettings} from "../PoSSecurityModule.sol";
import {ValidatorRegistrationInput, IValidatorManager} from "../interfaces/IValidatorManager.sol";
import {ExampleRewardCalculator} from "../ExampleRewardCalculator.sol";
import {ICMInitializable} from "../../utilities/ICMInitializable.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {Initializable} from "@openzeppelin/contracts@5.0.2/proxy/utils/Initializable.sol";

contract NativeTokenSecurityModuleTest is PoSSecurityModuleTest {
    NativeTokenSecurityModule public app;

    function setUp() public override {
        ValidatorManagerTest.setUp();

        _setUp();
        _mockGetBlockchainID();
        _mockInitializeValidatorSet();
        app.initializeValidatorSet(_defaultConversionData(), 0);
    }

    function testDisableInitialization() public {
        app = new NativeTokenSecurityModule(ICMInitializable.Disallowed);
        vm.expectRevert(abi.encodeWithSelector(Initializable.InvalidInitialization.selector));

        app.initialize(_defaultPoSSettings());
    }

    function testZeroMinimumDelegationFee() public {
        app = new NativeTokenSecurityModule(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(PoSSecurityModule.InvalidDelegationFee.selector, 0)
        );

        PoSSecurityModuleSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.minimumDelegationFeeBips = 0;
        app.initialize(defaultPoSSettings);
    }

    function testMaxMinimumDelegationFee() public {
        app = new NativeTokenSecurityModule(ICMInitializable.Allowed);
        uint16 minimumDelegationFeeBips = app.MAXIMUM_DELEGATION_FEE_BIPS() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSSecurityModule.InvalidDelegationFee.selector, minimumDelegationFeeBips
            )
        );

        PoSSecurityModuleSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.minimumDelegationFeeBips = minimumDelegationFeeBips;
        app.initialize(defaultPoSSettings);
    }

    function testInvalidStakeAmountRange() public {
        app = new NativeTokenSecurityModule(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSSecurityModule.InvalidStakeAmount.selector, DEFAULT_MAXIMUM_STAKE_AMOUNT
            )
        );

        PoSSecurityModuleSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.minimumStakeAmount = DEFAULT_MAXIMUM_STAKE_AMOUNT;
        defaultPoSSettings.maximumStakeAmount = DEFAULT_MINIMUM_STAKE_AMOUNT;
        app.initialize(defaultPoSSettings);
    }

    function testZeroMaxStakeMultiplier() public {
        app = new NativeTokenSecurityModule(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(PoSSecurityModule.InvalidStakeMultiplier.selector, 0)
        );

        PoSSecurityModuleSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.maximumStakeMultiplier = 0;
        app.initialize(defaultPoSSettings);
    }

    function testMaxStakeMultiplierOverLimit() public {
        app = new NativeTokenSecurityModule(ICMInitializable.Allowed);
        uint8 maximumStakeMultiplier = app.MAXIMUM_STAKE_MULTIPLIER_LIMIT() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSSecurityModule.InvalidStakeMultiplier.selector, maximumStakeMultiplier
            )
        );

        PoSSecurityModuleSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.maximumStakeMultiplier = maximumStakeMultiplier;
        app.initialize(defaultPoSSettings);
    }

    function testZeroWeightToValueFactor() public {
        app = new NativeTokenSecurityModule(ICMInitializable.Allowed);
        vm.expectRevert(
            abi.encodeWithSelector(PoSSecurityModule.ZeroWeightToValueFactor.selector)
        );

        PoSSecurityModuleSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.weightToValueFactor = 0;
        app.initialize(defaultPoSSettings);
    }

    function testMinStakeDurationTooLow() public {
        app = new NativeTokenSecurityModule(ICMInitializable.Allowed);
        uint64 minStakeDuration = DEFAULT_CHURN_PERIOD - 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSSecurityModule.InvalidMinStakeDuration.selector, minStakeDuration
            )
        );

        PoSSecurityModuleSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.minimumStakeDuration = minStakeDuration;
        app.initialize(defaultPoSSettings);
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
        address nativeMinter = address(app.NATIVE_MINTER());
        bytes memory callData = abi.encodeCall(INativeMinter.mintNativeCoin, (account, amount));
        vm.mockCall(nativeMinter, callData, "");
        vm.expectCall(nativeMinter, callData);
    }

    function _setUp() internal override returns (IValidatorManager) {
        // Construct the object under test
        app = new TestableNativeTokenSecurityModule(ICMInitializable.Allowed);
        rewardCalculator = new ExampleRewardCalculator(DEFAULT_REWARD_RATE);

        PoSSecurityModuleSettings memory defaultPoSSettings = _defaultPoSSettings();
        defaultPoSSettings.rewardCalculator = rewardCalculator;
        app.initialize(defaultPoSSettings);

        validatorManager = app;
        posValidatorManager = app;
        return app;
    }

    function _getStakeAssetBalance(address account) internal view override returns (uint256) {
        return account.balance;
    }
}

contract TestableNativeTokenSecurityModule is NativeTokenSecurityModule, Test {
    constructor(ICMInitializable init) NativeTokenSecurityModule(init) {}

    function _reward(address account, uint256 amount) internal virtual override {
        super._reward(account, amount);
        // Units tests don't have access to the native minter precompile, so use vm.deal instead.
        vm.deal(account, account.balance + amount);
    }
}
