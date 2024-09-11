// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {PoSValidatorManagerTest} from "./PoSValidatorManagerTests.t.sol";
import {ERC20TokenStakingManager} from "../ERC20TokenStakingManager.sol";
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
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
contract ERC20TokenStakingManagerTest is PoSValidatorManagerTest {
    using SafeERC20 for IERC20;

    ERC20TokenStakingManager public app;
    IERC20 public token;

    function setUp() public virtual {
        // Construct the object under test
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        token = new ExampleERC20();
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
            }),
            token
        );
        validatorManager = app;
        posValidatorManager = app;
    }

    function testZeroDelegationFee() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(_formatErrorMessage("zero delegation fee"));
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
                minimumDelegationFee: 0,
                maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
                rewardCalculator: IRewardCalculator(address(0))
            }),
            token
        );
    }

    function testInvalidStakeAmountRange() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(_formatErrorMessage("invalid stake amount range"));
        app.initialize(
            PoSValidatorManagerSettings({
                baseSettings: ValidatorManagerSettings({
                    pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                    subnetID: DEFAULT_SUBNET_ID,
                    maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN
                }),
                minimumStakeAmount: DEFAULT_MAXIMUM_STAKE,
                maximumStakeAmount: DEFAULT_MINIMUM_STAKE,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                minimumDelegationFee: DEFAULT_MINIMUM_DELEGATION_FEE,
                maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
                rewardCalculator: IRewardCalculator(address(0))
            }),
            token
        );
    }

    function testZeroMaxStakeMultiplier() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        vm.expectRevert(_formatErrorMessage("invalid maximum stake multiplier"));
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
                maximumStakeMultiplier: 0,
                rewardCalculator: IRewardCalculator(address(0))
            }),
            token
        );
    }

    function testMaxStakeMultiplierOverLimit() public {
        app = new ERC20TokenStakingManager(ICMInitializable.Allowed);
        uint8 maximumStakeMultiplier = app.MAXIMUM_STAKE_MULTIPLIER_LIMIT() + 1;
        vm.expectRevert(_formatErrorMessage("invalid maximum stake multiplier"));
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
                maximumStakeMultiplier: maximumStakeMultiplier,
                rewardCalculator: IRewardCalculator(address(0))
            }),
            token
        );
    }

    function testInvalidValidatorMinStakeDuration() public {
        ValidatorRegistrationInput memory input =
            ValidatorRegistrationInput(DEFAULT_NODE_ID, DEFAULT_EXPIRY, DEFAULT_BLS_PUBLIC_KEY);
        uint256 stakeAmount = app.weightToValue(DEFAULT_WEIGHT);
        vm.expectRevert(_formatErrorMessage("invalid min stake duration"));
        app.initializeValidatorRegistration(
            input,
            PoSValidatorRequirements({
                minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION - 1,
                delegationFee: DEFAULT_MINIMUM_DELEGATION_FEE
            }),
            stakeAmount
        );
    }

    function testInvalidValidatorDelegationFee() public {
        ValidatorRegistrationInput memory input =
            ValidatorRegistrationInput(DEFAULT_NODE_ID, DEFAULT_EXPIRY, DEFAULT_BLS_PUBLIC_KEY);
        uint256 stakeAmount = app.weightToValue(DEFAULT_WEIGHT);
        vm.expectRevert(_formatErrorMessage("invalid delegation fee"));
        app.initializeValidatorRegistration(
            input,
            PoSValidatorRequirements({
                minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                delegationFee: DEFAULT_MINIMUM_DELEGATION_FEE - 1
            }),
            stakeAmount
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration(
            input,
            PoSValidatorRequirements({
                minStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                delegationFee: DEFAULT_MINIMUM_DELEGATION_FEE
            }),
            app.weightToValue(weight)
        );
    }

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        uint256 value = app.weightToValue(weight);
        vm.startPrank(delegatorAddress);
        bytes32 delegationID = app.initializeDelegatorRegistration(validationID, value);
        vm.stopPrank();
        return delegationID;
    }

    function _beforeSend(uint64 weight, address spender) internal override {
        uint256 value = app.weightToValue(weight);
        token.safeIncreaseAllowance(spender, value);
        token.safeTransfer(spender, value);

        // ERC20 tokens need to be pre-approved
        vm.startPrank(spender);
        token.safeIncreaseAllowance(address(app), value);
        vm.stopPrank();
    }

    function _expectStakeUnlock(address account, uint256 amount) internal override {
        vm.expectCall(address(token), abi.encodeCall(IERC20.transfer, (account, amount)));
    }

    function _getStakeAssetBalance(address account) internal view override returns (uint256) {
        return token.balanceOf(account);
    }
}
