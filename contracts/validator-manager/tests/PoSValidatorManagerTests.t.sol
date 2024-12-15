// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {PoSValidatorManager} from "../PoSValidatorManager.sol";
import {
    DelegatorStatus, PoSValidatorManagerSettings
} from "../interfaces/IPoSValidatorManager.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import {
    ValidatorManagerSettings,
    ValidatorRegistrationInput,
    ValidatorStatus
} from "../interfaces/IValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

abstract contract PoSValidatorManagerTest is ValidatorManagerTest {
    uint64 public constant DEFAULT_UPTIME = uint64(100);
    uint64 public constant DEFAULT_DELEGATOR_WEIGHT = uint64(1e5);
    uint64 public constant DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP =
        DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_EXPIRY;
    uint64 public constant DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP =
        DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + DEFAULT_EXPIRY;
    uint64 public constant DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP =
        DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION;
    address public constant DEFAULT_DELEGATOR_ADDRESS =
        address(0x1234123412341234123412341234123412341234);
    address public constant DEFAULT_VALIDATOR_OWNER_ADDRESS =
        address(0x2345234523452345234523452345234523452345);
    uint64 public constant DEFAULT_REWARD_RATE = uint64(10);
    uint64 public constant DEFAULT_MINIMUM_STAKE_DURATION = 24 hours;
    uint16 public constant DEFAULT_MINIMUM_DELEGATION_FEE_BIPS = 100;
    uint16 public constant DEFAULT_DELEGATION_FEE_BIPS = 150;
    uint8 public constant DEFAULT_MAXIMUM_STAKE_MULTIPLIER = 4;
    uint256 public constant DEFAULT_WEIGHT_TO_VALUE_FACTOR = 1e12;
    uint256 public constant SECONDS_IN_YEAR = 31536000;

    PoSValidatorManager public posValidatorManager;
    IRewardCalculator public rewardCalculator;

    ValidatorRegistrationInput public defaultRegistrationInput = ValidatorRegistrationInput({
        nodeID: DEFAULT_NODE_ID,
        blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
        registrationExpiry: DEFAULT_EXPIRY,
        remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
        disableOwner: DEFAULT_P_CHAIN_OWNER
    });

    event ValidationUptimeUpdated(bytes32 indexed validationID, uint64 uptime);

    event DelegatorAdded(
        bytes32 indexed delegationID,
        bytes32 indexed validationID,
        address indexed delegatorAddress,
        uint64 nonce,
        uint64 validatorWeight,
        uint64 delegatorWeight,
        bytes32 setWeightMessageID
    );

    event DelegatorRegistered(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime
    );

    event DelegatorRemovalInitialized(bytes32 indexed delegationID, bytes32 indexed validationID);

    event ValidatorWeightUpdate(
        bytes32 indexed validationID,
        uint64 indexed nonce,
        uint64 weight,
        bytes32 setWeightMessageID
    );

    event DelegationEnded(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees
    );

    event UptimeUpdated(bytes32 indexed validationID, uint64 uptime);

    function testDelegationFeeBipsTooLow() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegationFee.selector,
                DEFAULT_MINIMUM_DELEGATION_FEE_BIPS - 1
            )
        );
        _initializeValidatorRegistration(
            defaultRegistrationInput,
            DEFAULT_MINIMUM_DELEGATION_FEE_BIPS - 1,
            DEFAULT_MINIMUM_STAKE_DURATION,
            DEFAULT_MINIMUM_STAKE_AMOUNT
        );
    }

    function testDelegationFeeBipsTooHigh() public {
        uint16 delegationFeeBips = posValidatorManager.MAXIMUM_DELEGATION_FEE_BIPS() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegationFee.selector, delegationFeeBips
            )
        );

        _initializeValidatorRegistration(
            defaultRegistrationInput,
            delegationFeeBips,
            DEFAULT_MINIMUM_STAKE_DURATION,
            DEFAULT_MINIMUM_STAKE_AMOUNT
        );
    }

    function testInvalidMinStakeDuration() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidMinStakeDuration.selector,
                DEFAULT_MINIMUM_STAKE_DURATION - 1
            )
        );
        _initializeValidatorRegistration(
            defaultRegistrationInput,
            DEFAULT_DELEGATION_FEE_BIPS,
            DEFAULT_MINIMUM_STAKE_DURATION - 1,
            DEFAULT_MINIMUM_STAKE_AMOUNT
        );
    }

    function testStakeAmountTooLow() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidStakeAmount.selector, DEFAULT_MINIMUM_STAKE_AMOUNT - 1
            )
        );
        _initializeValidatorRegistration(
            defaultRegistrationInput,
            DEFAULT_DELEGATION_FEE_BIPS,
            DEFAULT_MINIMUM_STAKE_DURATION,
            DEFAULT_MINIMUM_STAKE_AMOUNT - 1
        );
    }

    function testStakeAmountTooHigh() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidStakeAmount.selector, DEFAULT_MAXIMUM_STAKE_AMOUNT + 1
            )
        );
        _initializeValidatorRegistration(
            defaultRegistrationInput,
            DEFAULT_DELEGATION_FEE_BIPS,
            DEFAULT_MINIMUM_STAKE_DURATION,
            DEFAULT_MAXIMUM_STAKE_AMOUNT + 1
        );
    }

    function testInvalidInitializeEndTime() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.MinStakeDurationNotPassed.selector, block.timestamp
            )
        );
        posValidatorManager.initializeEndValidation(validationID, false, 0);
    }

    function testInvalidUptimeWarpMessage() public {
        bytes32 validationID = _registerDefaultValidator();

        _mockGetUptimeWarpMessage(new bytes(0), false);
        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        vm.expectRevert(ValidatorManager.InvalidWarpMessage.selector);
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeSenderAddress() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    originSenderAddress: address(this),
                    payload: new bytes(0)
                }),
                true
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidWarpOriginSenderAddress.selector, address(this)
            )
        );
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeValidationID() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    originSenderAddress: address(0),
                    payload: ValidatorMessages.packValidationUptimeMessage(bytes32(0), 0)
                }),
                true
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorManager.InvalidValidationID.selector, validationID)
        );
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInitializeDelegatorRegistration() public {
        bytes32 validationID = _registerDefaultValidator();

        _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
    }

    function testResendDelegatorRegistration() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes32 delegationID = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, 1, DEFAULT_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
        );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendUpdateDelegation(delegationID);
    }

    function testCompleteDelegatorRegistration() public {
        bytes32 validationID = _registerDefaultValidator();

        _registerDefaultDelegator(validationID);
    }

    function testCompleteDelegatorRegistrationWrongNonce() public {
        bytes32 validationID = _registerDefaultValidator();

        // Initialize two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
                + DEFAULT_WEIGHT,
            expectedNonce: 2
        });

        // Complete registration of delegator2 with delegator1's nonce
        // Note that registering delegator1 with delegator2's nonce is valid
        uint64 nonce = 1;
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, nonce, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.warp(DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP);
        vm.expectRevert(abi.encodeWithSelector(PoSValidatorManager.InvalidNonce.selector, nonce));
        posValidatorManager.completeDelegatorRegistration(delegationID2, 0);
    }

    function testCompleteDelegatorRegistrationImplicitNonce() public {
        bytes32 validationID = _registerDefaultValidator();

        // Initialize two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
        address delegator2 = address(0x5678567856785678567856785678567856785678);
        _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
                + DEFAULT_WEIGHT,
            expectedNonce: 2
        });
        // Mark delegator1 as registered by delivering the weight update from nonce 2 (delegator 2's nonce)
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, 2, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );

        _setUpCompleteDelegatorRegistration(
            delegationID1,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            setValidatorWeightPayload
        );
    }

    function testInitializeEndValidationNotOwner() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.prank(address(1));
        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.UnauthorizedOwner.selector, address(1))
        );
        posValidatorManager.initializeEndValidation(validationID, false, 0);
    }

    function testInitializeEndDelegation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });
    }

    function testInitializeEndDelegationByValidator() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: address(this),
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });
    }

    function testInitializeEndDelegationByValidatorMinStakeDurationNotPassed() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        uint64 invalidEndTime = DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1 hours;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.MinStakeDurationNotPassed.selector, invalidEndTime
            )
        );
        _initializeEndDelegation({
            sender: address(this),
            delegationID: delegationID,
            endDelegationTimestamp: invalidEndTime,
            includeUptime: false,
            force: false,
            rewardRecipient: address(0)
        });
    }

    function testInitializeEndDelegationMinStakeDurationNotPassed() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        uint64 invalidEndTime =
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION - 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.MinStakeDurationNotPassed.selector, invalidEndTime
            )
        );
        _initializeEndDelegation({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: invalidEndTime,
            includeUptime: false,
            force: false,
            rewardRecipient: address(0)
        });
    }

    function testCompleteEndDelegationChurnPeriodSecondsNotPassed() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 delegatorRegistrationTime =
            DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION + 1;
        bytes32 delegationID = _registerDelegator({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: delegatorRegistrationTime - 1,
            completeRegistrationTimestamp: delegatorRegistrationTime,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });

        address validatorOwner = address(this);

        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: delegatorRegistrationTime + 1,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: validatorOwner
        });

        uint64 invalidEndTime = delegatorRegistrationTime + DEFAULT_CHURN_PERIOD - 1;

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.MinStakeDurationNotPassed.selector, invalidEndTime
            )
        );

        // Initialize end delegation will also call _completeEndDelegation because the validator is copmleted.
        _initializeEndDelegation({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: invalidEndTime,
            includeUptime: false,
            force: false,
            rewardRecipient: address(0)
        });
    }

    function testInitializeEndDelegationInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.DelegatorIneligibleForRewards.selector, delegationID
            )
        );
        vm.warp(DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP);
        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        posValidatorManager.initializeEndDelegation(delegationID, false, 0);
    }

    function testForceInitializeEndDelegation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: true,
            rewardRecipient: address(0)
        });
    }

    function testForceInitializeEndDelegationInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: false,
            force: true,
            rewardRecipient: address(0)
        });
    }

    function testResendEndDelegation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 2, DEFAULT_WEIGHT);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendUpdateDelegation(delegationID);
    }

    function testResendEndValidation() public override {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        validatorManager.resendEndValidatorMessage(validationID);
    }

    function testCompleteEndDelegation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _completeDefaultDelegator(validationID, delegationID);
    }

    function testClaimDelegationFeesInvalidValidatorStatus() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _completeDefaultDelegator(validationID, delegationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.Active
            )
        );

        posValidatorManager.claimDelegationFees(validationID);
    }

    function testClaimDelegationFeesInvalidSender() public {
        bytes32 validationID = _registerDefaultValidator();
        _registerDefaultDelegator(validationID);

        _endDefaultValidatorWithChecks(validationID, 2);

        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.UnauthorizedOwner.selector, address(123))
        );

        vm.prank(address(123));
        posValidatorManager.claimDelegationFees(validationID);
    }

    function testClaimDelegationFees() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);
        address rewardRecipient = address(42);

        _endDefaultValidatorWithChecks(validationID, 2);

        // Validator is Completed, so this will also complete the delegation.
        _initializeEndDelegation({
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            includeUptime: true,
            force: false,
            rewardRecipient: rewardRecipient
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        _expectRewardIssuance(
            address(this), expectedTotalReward * DEFAULT_DELEGATION_FEE_BIPS / 10000
        );
        posValidatorManager.claimDelegationFees(validationID);
    }

    function testCompleteEndDelegationWithNonDelegatorRewardRecipient() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);
        address rewardRecipient = address(42);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: rewardRecipient
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: rewardRecipient
        });
    }

    function testChangeDelegatorRewardRecipientWithNullAddress() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);
        address rewardRecipient = address(42);
        address newRewardRecipient = address(0);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: rewardRecipient
        });

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidRewardRecipient.selector, newRewardRecipient
            )
        );

        posValidatorManager.changeDelegatorRewardRecipient(delegationID, newRewardRecipient);
    }

    function testChangeDelegatorRewardRecipientByNonDelegator() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);
        address rewardRecipient = address(42);
        address badActor = address(43);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: rewardRecipient
        });

        vm.prank(badActor);

        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.UnauthorizedOwner.selector, badActor)
        );

        posValidatorManager.changeDelegatorRewardRecipient(delegationID, badActor);
    }

    function testChangeDelegatorRewardRecipientBackToSelf() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);
        address rewardRecipient = address(42);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: rewardRecipient
        });

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);

        posValidatorManager.changeDelegatorRewardRecipient(delegationID, DEFAULT_DELEGATOR_ADDRESS);

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: DEFAULT_DELEGATOR_ADDRESS
        });
    }

    function testChangeDelegatorRewardRecipient() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);
        address rewardRecipient = address(42);
        address newRewardRecipient = address(43);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: rewardRecipient
        });

        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        posValidatorManager.changeDelegatorRewardRecipient(delegationID, newRewardRecipient);

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: newRewardRecipient
        });
    }

    // Delegator registration is not allowed when Validator is pending removed.
    function testInitializeDelegatorRegistrationValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.PendingRemoved
            )
        );
        _initializeDelegatorRegistration(
            validationID, DEFAULT_DELEGATOR_ADDRESS, DEFAULT_DELEGATOR_WEIGHT
        );
    }

    // Complete delegator registration may be called when validator is pending removed.
    function testCompleteRegisterDelegatorValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes32 delegationID = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 2, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _setUpCompleteDelegatorRegistrationWithChecks(
            validationID,
            delegationID,
            DEFAULT_COMPLETION_TIMESTAMP + 1,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1
        );
    }

    // Delegator cannot initialize end delegation when validator is pending removed.
    function testInitializeEndDelegationValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.PendingRemoved
            )
        );
        _initializeDelegatorRegistration(
            validationID, DEFAULT_DELEGATOR_ADDRESS, DEFAULT_DELEGATOR_WEIGHT
        );
    }

    // Delegator may complete end delegation while validator is pending removed.
    function testCompleteEndDelegationValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });

        uint64 validationEndTime = DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1;
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 3, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, validationEndTime - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: validationEndTime,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: validationEndTime - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        address delegator = DEFAULT_DELEGATOR_ADDRESS;

        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: delegator,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: 0,
            expectedNonce: 2,
            rewardRecipient: delegator
        });
    }

    function testInitializeDelegatorRegistrationValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        _endDefaultValidatorWithChecks(validationID, 1);

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP + 1);
        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.Completed
            )
        );
        _initializeDelegatorRegistration(
            validationID, DEFAULT_DELEGATOR_ADDRESS, DEFAULT_DELEGATOR_WEIGHT
        );
    }

    function testCompleteDelegatorRegistrationValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initializeDefaultDelegatorRegistration(validationID);

        _endDefaultValidatorWithChecks(validationID, 2);

        // completeDelegatorRegistration should fall through to _completeEndDelegation and refund the stake
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit DelegationEnded(delegationID, validationID, 0, 0);

        uint256 balanceBefore = _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS);

        _expectStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, _weightToValue(DEFAULT_DELEGATOR_WEIGHT));

        // warp to right after validator ended
        vm.warp(DEFAULT_COMPLETION_TIMESTAMP + 1);
        posValidatorManager.completeDelegatorRegistration(delegationID, 0);

        assertEq(
            _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS),
            balanceBefore + _weightToValue(DEFAULT_DELEGATOR_WEIGHT)
        );
    }

    function testInitializeEndDelegationValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _endDefaultValidatorWithChecks(validationID, 2);

        uint64 delegationEndTime = DEFAULT_COMPLETION_TIMESTAMP + 1;

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint256 expectedValidatorFees = expectedTotalReward * DEFAULT_DELEGATION_FEE_BIPS / 10000;
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        // completeDelegatorRegistration should fall through to _completeEndDelegation and refund the stake
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit DelegationEnded(
            delegationID, validationID, expectedDelegatorReward, expectedValidatorFees
        );

        uint256 balanceBefore = _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS);

        _expectStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, _weightToValue(DEFAULT_DELEGATOR_WEIGHT));
        _expectRewardIssuance(DEFAULT_DELEGATOR_ADDRESS, expectedDelegatorReward);

        // warp to right after validator ended
        vm.warp(delegationEndTime);
        vm.prank(DEFAULT_DELEGATOR_ADDRESS);
        posValidatorManager.initializeEndDelegation(delegationID, false, 0);

        assertEq(
            _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS),
            balanceBefore + _weightToValue(DEFAULT_DELEGATOR_WEIGHT) + expectedDelegatorReward
        );
    }

    function testCompleteEndDelegationValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });

        _endDefaultValidatorWithChecks(validationID, 3);

        uint256 expectedTotalReward = _defaultDelegatorExpectedTotalReward();

        uint256 expectedValidatorFees = (expectedTotalReward * DEFAULT_DELEGATION_FEE_BIPS) / 10000;
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegationEnded(
            delegationID, validationID, expectedDelegatorReward, expectedValidatorFees
        );
        uint256 balanceBefore = _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS);

        _expectStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, _weightToValue(DEFAULT_DELEGATOR_WEIGHT));
        _expectRewardIssuance(DEFAULT_DELEGATOR_ADDRESS, expectedDelegatorReward);

        posValidatorManager.completeEndDelegation(delegationID, 0);

        assertEq(
            _getStakeAssetBalance(DEFAULT_DELEGATOR_ADDRESS),
            balanceBefore + _weightToValue(DEFAULT_DELEGATOR_WEIGHT) + expectedDelegatorReward
        );
    }

    function testCompleteEndDelegationWrongNonce() public {
        bytes32 validationID = _registerDefaultValidator();
        // Register two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _registerDelegator({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });

        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _registerDelegator({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT * 2 + DEFAULT_WEIGHT,
            expectedNonce: 2
        });

        // Initialize end delegation for both delegators
        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: delegator1,
            delegationID: delegationID1,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 3,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });
        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: delegator2,
            delegationID: delegationID2,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });

        // Complete ending delegator2 with delegator1's nonce
        // Note that ending delegator1 with delegator2's nonce is valid
        uint64 nonce = 3;
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, nonce, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.expectRevert(abi.encodeWithSelector(PoSValidatorManager.InvalidNonce.selector, nonce));
        posValidatorManager.completeEndDelegation(delegationID2, 0);
    }

    function testCompleteEndDelegationImplicitNonce() public {
        bytes32 validationID = _registerDefaultValidator();

        // Register two delegations
        address delegator1 = DEFAULT_DELEGATOR_ADDRESS;
        bytes32 delegationID1 = _registerDelegator({
            validationID: validationID,
            delegatorAddress: delegator1,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });

        address delegator2 = address(0x5678567856785678567856785678567856785678);
        bytes32 delegationID2 = _registerDelegator({
            validationID: validationID,
            delegatorAddress: delegator2,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT * 2 + DEFAULT_WEIGHT,
            expectedNonce: 2
        });

        // Initialize end delegation for both delegators
        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: delegator1,
            delegationID: delegationID1,
            startDelegationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 3,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });
        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: delegator2,
            delegationID: delegationID2,
            startDelegationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });

        uint256 expectedTotalReward = _defaultDelegatorExpectedTotalReward();

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        address delegator = DEFAULT_DELEGATOR_ADDRESS;

        // Complete delegation1 by delivering the weight update from nonce 4 (delegator2's nonce)
        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID1,
            delegator: delegator,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4,
            rewardRecipient: delegator
        });
    }

    function testCompleteEndValidation() public virtual override {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        address validatorOwner = address(this);

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: validatorOwner
        });
    }

    function testCompleteEndValidationWithNonValidatorRewardRecipient() public virtual {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address rewardRecipient = address(42);

        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false,
            recipientAddress: rewardRecipient
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        address validatorOwner = address(this);

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: rewardRecipient
        });
    }

    function testChangeValidatorRewardRecipient() public virtual {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address rewardRecipient = address(42);
        address newRecipient = address(43);

        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false,
            recipientAddress: rewardRecipient
        });

        posValidatorManager.changeValidatorRewardRecipient(validationID, newRecipient);

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: newRecipient
        });
    }

    function testChangeValidatorRewardRecipientBackToSelf() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address rewardRecipient = address(42);
        address newRecipient = address(this);

        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false,
            recipientAddress: rewardRecipient
        });

        posValidatorManager.changeValidatorRewardRecipient(validationID, newRecipient);

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        address validatorOwner = address(this);

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: validatorOwner
        });
    }

    function testChangeValidatorRewardRecipientWithNullAddress() public virtual {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address rewardRecipient = address(42);
        address newRecipient = address(0);

        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false,
            recipientAddress: rewardRecipient
        });

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidRewardRecipient.selector, newRecipient
            )
        );

        posValidatorManager.changeValidatorRewardRecipient(validationID, newRecipient);
    }

    function testChangeValidatorRewardRecipientByNonValidator() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        address rewardRecipient = address(42);
        address badActor = address(43);

        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false,
            recipientAddress: rewardRecipient
        });

        vm.prank(badActor);

        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.UnauthorizedOwner.selector, badActor)
        );

        posValidatorManager.changeValidatorRewardRecipient(validationID, badActor);
    }

    function testInitializeEndValidation() public virtual override {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });
    }

    function testInitializeEndValidationUseStoredUptime() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        // Submit an uptime proof via submitUptime
        uint64 uptimePercentage1 = 80;
        uint64 uptime1 = (
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage1
        ) / 100;
        bytes memory uptimeMsg1 =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime1);
        _mockGetUptimeWarpMessage(uptimeMsg1, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit UptimeUpdated(validationID, uptime1);
        posValidatorManager.submitUptimeProof(validationID, 0);

        // Submit a second uptime proof via initializeEndValidation. This one is not sufficient for rewards
        // Submit an uptime proof via submitUptime
        uint64 uptimePercentage2 = 79;
        uint64 uptime2 = (
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage2
        ) / 100;
        bytes memory uptimeMsg2 =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime2);
        _mockGetUptimeWarpMessage(uptimeMsg2, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _initializeEndValidation(validationID, true, address(0));
    }

    function testInitializeEndValidationWithoutNewUptime() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        // Submit an uptime proof via submitUptime
        uint64 uptimePercentage1 = 80;
        uint64 uptime1 = (
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage1
        ) / 100;
        bytes memory uptimeMsg1 =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime1);
        _mockGetUptimeWarpMessage(uptimeMsg1, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit UptimeUpdated(validationID, uptime1);
        posValidatorManager.submitUptimeProof(validationID, 0);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _initializeEndValidation(validationID, false, address(0));
    }

    function testInitializeEndValidationInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 79;

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            ((DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage)
                / 100
        );
        _mockGetUptimeWarpMessage(uptimeMsg, true);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.ValidatorIneligibleForRewards.selector, validationID
            )
        );

        _initializeEndValidation(validationID, true, address(0));
    }

    function testSubmitUptimeProofPoaValidator() public {
        bytes32 defaultInitialValidationID = sha256(abi.encodePacked(DEFAULT_L1_ID, uint32(1)));

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.ValidatorNotPoS.selector, defaultInitialValidationID
            )
        );
        posValidatorManager.submitUptimeProof(defaultInitialValidationID, 0);
    }

    function testSubmitUptimeProofInactiveValidator() public {
        bytes32 validationID = _registerDefaultValidator();

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );

        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.PendingRemoved
            )
        );
        posValidatorManager.submitUptimeProof(validationID, 0);
    }

    function testEndValidationPoAValidator() public {
        bytes32 validationID = sha256(abi.encodePacked(DEFAULT_L1_ID, uint32(1)));

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _initializeEndValidation(validationID, false, address(0));

        uint256 balanceBefore = _getStakeAssetBalance(address(this));

        bytes memory l1ValidatorRegistrationMessage =
            ValidatorMessages.packL1ValidatorRegistrationMessage(validationID, false);
        _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);

        posValidatorManager.completeEndValidation(0);

        assertEq(_getStakeAssetBalance(address(this)), balanceBefore);
    }

    function testDelegationToPoAValidator() public {
        bytes32 defaultInitialValidationID = sha256(abi.encodePacked(DEFAULT_L1_ID, uint32(1)));

        _beforeSend(_weightToValue(DEFAULT_DELEGATOR_WEIGHT), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.ValidatorNotPoS.selector, defaultInitialValidationID
            )
        );

        _initializeDelegatorRegistration(
            defaultInitialValidationID, DEFAULT_DELEGATOR_ADDRESS, DEFAULT_DELEGATOR_WEIGHT
        );
    }

    function testDelegationOverWeightLimit() public {
        bytes32 validationID = _registerDefaultValidator();

        uint64 delegatorWeight = DEFAULT_WEIGHT * DEFAULT_MAXIMUM_STAKE_MULTIPLIER + 1;

        _beforeSend(_weightToValue(delegatorWeight), DEFAULT_DELEGATOR_ADDRESS);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.MaxWeightExceeded.selector, delegatorWeight + DEFAULT_WEIGHT
            )
        );

        _initializeDelegatorRegistration(validationID, DEFAULT_DELEGATOR_ADDRESS, delegatorWeight);
    }

    function testCompleteDelegatorRegistrationAlreadyRegistered() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegatorStatus.selector, DelegatorStatus.Active
            )
        );

        posValidatorManager.completeDelegatorRegistration(delegationID, 0);
    }

    function testCompleteDelegatorRegistrationWrongValidationID() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initializeDefaultDelegatorRegistration(validationID);

        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            delegationID, 2, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.expectRevert(
            abi.encodeWithSelector(ValidatorManager.InvalidValidationID.selector, validationID)
        );

        vm.warp(DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP);
        posValidatorManager.completeDelegatorRegistration(delegationID, 0);
    }

    function testCompleteEndDelegationWrongValidationID() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: address(this)
        });

        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(delegationID, 2, DEFAULT_WEIGHT);
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.expectRevert(
            abi.encodeWithSelector(ValidatorManager.InvalidValidationID.selector, delegationID)
        );

        posValidatorManager.completeEndDelegation(delegationID, 0);
    }

    function testInitializeEndDelegationNotRegistered() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initializeDefaultDelegatorRegistration(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegatorStatus.selector, DelegatorStatus.PendingAdded
            )
        );

        posValidatorManager.initializeEndDelegation(delegationID, true, 0);
    }

    function testInitializeEndDelegationWrongSender() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.UnauthorizedOwner.selector, address(123))
        );

        vm.prank(address(123));
        posValidatorManager.initializeEndDelegation(delegationID, true, 0);
    }

    function testCompleteDelegatorRegistrationForDelegatorRegisteredWhileValidatorPendingRemoved()
        public
    {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initializeDefaultDelegatorRegistration(validationID);

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 2, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );

        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _setUpCompleteDelegatorRegistrationWithChecks(
            validationID, delegationID, DEFAULT_COMPLETION_TIMESTAMP + 1, 0, 2
        );

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT,
            rewardRecipient: address(this)
        });

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP + 1 + DEFAULT_MINIMUM_STAKE_DURATION);
        _expectStakeUnlock(DEFAULT_DELEGATOR_ADDRESS, _weightToValue(DEFAULT_DELEGATOR_WEIGHT));
        posValidatorManager.initializeEndDelegation(delegationID, true, 0);
    }

    function testCompleteEndDelegationWhileActive() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegatorStatus.selector, DelegatorStatus.Active
            )
        );

        posValidatorManager.completeEndDelegation(delegationID, 0);
    }

    function testCompleteDelegatorRegistrationValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _initializeDefaultDelegatorRegistration(validationID);

        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 2, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );

        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        _setUpCompleteDelegatorRegistrationWithChecks(
            validationID, delegationID, DEFAULT_COMPLETION_TIMESTAMP + 1, 0, 2
        );

        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidValidatorStatus.selector, ValidatorStatus.PendingRemoved
            )
        );

        posValidatorManager.initializeEndDelegation(delegationID, false, 0);
    }

    function testResendEndDelegationWhileActive() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegatorStatus.selector, DelegatorStatus.Active
            )
        );

        posValidatorManager.resendUpdateDelegation(delegationID);
    }

    function testForceInitializeEndValidation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: true
        });
    }

    function testForceInitializeEndValidationInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 79;

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            ((DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage)
                / 100
        );
        _mockGetUptimeWarpMessage(uptimeMsg, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _forceInitializeEndValidation(validationID, true, address(0));
    }

    function testValueToWeightTruncated() public {
        // default weightToValueFactor is 1e12
        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.InvalidStakeAmount.selector, 1e11)
        );
        posValidatorManager.valueToWeight(1e11);
    }

    function testValueToWeightExceedsUInt64Max() public {
        // default weightToValueFactor is 1e12
        vm.expectRevert(
            abi.encodeWithSelector(PoSValidatorManager.InvalidStakeAmount.selector, 1e40)
        );
        posValidatorManager.valueToWeight(1e40);
    }

    function testValueToWeight() public view {
        uint64 w1 = posValidatorManager.valueToWeight(1e12);
        uint64 w2 = posValidatorManager.valueToWeight(1e18);
        uint64 w3 = posValidatorManager.valueToWeight(1e27);

        assertEq(w1, 1);
        assertEq(w2, 1e6);
        assertEq(w3, 1e15);
    }

    function testWeightToValue() public view {
        uint256 v1 = posValidatorManager.weightToValue(1);
        uint256 v2 = posValidatorManager.weightToValue(1e6);
        uint256 v3 = posValidatorManager.weightToValue(1e15);

        assertEq(v1, 1e12);
        assertEq(v2, 1e18);
        assertEq(v3, 1e27);
    }

    function testPoSValidatorManagerStorageSlot() public view {
        assertEq(
            _erc7201StorageSlot("PoSValidatorManager"),
            posValidatorManager.POS_VALIDATOR_MANAGER_STORAGE_LOCATION()
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory registrationInput,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) internal virtual returns (bytes32);

    function _initializeEndValidation(
        bytes32 validationID,
        bool includeUptime,
        address recipientAddress
    ) internal virtual override {
        posValidatorManager.initializeEndValidation(
            validationID, includeUptime, 0, recipientAddress
        );
    }

    function _forceInitializeEndValidation(
        bytes32 validationID,
        bool includeUptime,
        address recipientAddress
    ) internal virtual override {
        posValidatorManager.forceInitializeEndValidation(
            validationID, includeUptime, 0, recipientAddress
        );
    }

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight
    ) internal virtual returns (bytes32);

    //
    // Delegation setup utilities
    //
    function _setUpInitializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight,
        uint64 registrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal returns (bytes32) {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.warp(registrationTimestamp);

        _beforeSend(_weightToValue(weight), delegatorAddress);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorAdded({
            delegationID: keccak256(abi.encodePacked(validationID, expectedNonce)),
            validationID: validationID,
            delegatorAddress: delegatorAddress,
            nonce: expectedNonce,
            validatorWeight: expectedValidatorWeight,
            delegatorWeight: weight,
            setWeightMessageID: bytes32(0)
        });

        return _initializeDelegatorRegistration(validationID, delegatorAddress, weight);
    }

    function _setUpCompleteDelegatorRegistration(
        bytes32 delegationID,
        uint64 completeRegistrationTimestamp,
        bytes memory setValidatorWeightPayload
    ) internal {
        _mockGetPChainWarpMessage(setValidatorWeightPayload, true);

        vm.warp(completeRegistrationTimestamp);
        posValidatorManager.completeDelegatorRegistration(delegationID, 0);
    }

    function _setUpCompleteDelegatorRegistrationWithChecks(
        bytes32 validationID,
        bytes32 delegationID,
        uint64 completeRegistrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRegistered({
            delegationID: delegationID,
            validationID: validationID,
            startTime: completeRegistrationTimestamp
        });
        _setUpCompleteDelegatorRegistration(
            delegationID, completeRegistrationTimestamp, setValidatorWeightPayload
        );
    }

    function _registerDefaultDelegator(bytes32 validationID)
        internal
        returns (bytes32 delegationID)
    {
        return _registerDelegator({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            initRegistrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
    }

    function _initializeDefaultDelegatorRegistration(bytes32 validationID)
        internal
        returns (bytes32)
    {
        return _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: DEFAULT_DELEGATOR_ADDRESS,
            weight: DEFAULT_DELEGATOR_WEIGHT,
            registrationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
        });
    }

    function _completeDefaultDelegatorRegistration(
        bytes32 validationID,
        bytes32 delegationID
    ) internal {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, 1, DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT
        );
        _setUpCompleteDelegatorRegistration({
            delegationID: delegationID,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            setValidatorWeightPayload: setValidatorWeightPayload
        });
    }

    function _completeDefaultDelegator(bytes32 validationID, bytes32 delegationID) internal {
        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: address(this)
        });

        uint256 expectedTotalReward = _defaultDelegatorExpectedTotalReward();
        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: address(this)
        });
    }

    function _registerDelegator(
        bytes32 validationID,
        address delegatorAddress,
        uint64 weight,
        uint64 initRegistrationTimestamp,
        uint64 completeRegistrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal returns (bytes32) {
        bytes32 delegationID = _setUpInitializeDelegatorRegistration({
            validationID: validationID,
            delegatorAddress: delegatorAddress,
            weight: weight,
            registrationTimestamp: initRegistrationTimestamp,
            expectedValidatorWeight: expectedValidatorWeight,
            expectedNonce: expectedNonce
        });
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );

        _setUpCompleteDelegatorRegistration(
            delegationID, completeRegistrationTimestamp, setValidatorWeightPayload
        );
        return delegationID;
    }

    function _initializeEndDelegationValidatorActiveWithChecks(
        bytes32 validationID,
        address sender,
        bytes32 delegationID,
        uint64 startDelegationTimestamp,
        uint64 endDelegationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce,
        bool includeUptime,
        bool force,
        address rewardRecipient
    ) internal {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, endDelegationTimestamp - startDelegationTimestamp
        );
        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit ValidatorWeightUpdate({
            validationID: validationID,
            nonce: expectedNonce,
            weight: expectedValidatorWeight,
            setWeightMessageID: bytes32(0)
        });

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRemovalInitialized({delegationID: delegationID, validationID: validationID});

        _initializeEndDelegationValidatorActive({
            sender: sender,
            delegationID: delegationID,
            endDelegationTimestamp: endDelegationTimestamp,
            includeUptime: includeUptime,
            force: force,
            setValidatorWeightPayload: setValidatorWeightPayload,
            uptimePayload: uptimeMsg,
            rewardRecipient: rewardRecipient
        });
    }

    function _initializeEndDelegationValidatorActive(
        address sender,
        bytes32 delegationID,
        uint64 endDelegationTimestamp,
        bool includeUptime,
        bool force,
        bytes memory setValidatorWeightPayload,
        bytes memory uptimePayload,
        address rewardRecipient
    ) internal {
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        if (includeUptime) {
            _mockGetUptimeWarpMessage(uptimePayload, true);
        }
        _initializeEndDelegation({
            sender: sender,
            delegationID: delegationID,
            endDelegationTimestamp: endDelegationTimestamp,
            includeUptime: includeUptime,
            force: force,
            rewardRecipient: rewardRecipient
        });
    }

    function _initializeEndDelegation(
        address sender,
        bytes32 delegationID,
        uint64 endDelegationTimestamp,
        bool includeUptime,
        bool force,
        address rewardRecipient
    ) internal {
        vm.warp(endDelegationTimestamp);
        vm.prank(sender);
        if (force) {
            posValidatorManager.forceInitializeEndDelegation(
                delegationID, includeUptime, 0, rewardRecipient
            );
        } else {
            posValidatorManager.initializeEndDelegation(
                delegationID, includeUptime, 0, rewardRecipient
            );
        }
    }

    function _endDefaultValidatorWithChecks(bytes32 validationID, uint64 expectedNonce) internal {
        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: expectedNonce,
            rewardRecipient: address(this)
        });
    }

    function _endDefaultValidator(bytes32 validationID, uint64 expectedNonce) internal {
        address validatorOwner = address(this);
        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: expectedNonce,
            rewardRecipient: validatorOwner
        });
    }

    function _endValidationWithChecks(
        bytes32 validationID,
        address validatorOwner,
        uint64 completeRegistrationTimestamp,
        uint64 completionTimestamp,
        uint64 validatorWeight,
        uint64 expectedNonce,
        address rewardRecipient
    ) internal {
        bytes memory setWeightMessage =
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, expectedNonce, 0);
        bytes memory uptimeMessage = ValidatorMessages.packValidationUptimeMessage(
            validationID, completionTimestamp - completeRegistrationTimestamp
        );
        _initializeEndValidation({
            validationID: validationID,
            completionTimestamp: completionTimestamp,
            setWeightMessage: setWeightMessage,
            includeUptime: true,
            uptimeMessage: uptimeMessage,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(validatorWeight),
            validatorStartTime: completeRegistrationTimestamp,
            stakingStartTime: completeRegistrationTimestamp,
            stakingEndTime: completionTimestamp,
            uptimeSeconds: completionTimestamp - completeRegistrationTimestamp
        });

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: validatorWeight,
            rewardRecipient: rewardRecipient
        });
    }

    function _completeEndValidationWithChecks(
        bytes32 validationID,
        address validatorOwner,
        uint256 expectedReward,
        uint64 validatorWeight,
        address rewardRecipient
    ) internal {
        bytes memory l1ValidatorRegistrationMessage =
            ValidatorMessages.packL1ValidatorRegistrationMessage(validationID, false);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit ValidationPeriodEnded(validationID, ValidatorStatus.Completed);
        uint256 balanceBefore = _getStakeAssetBalance(validatorOwner);
        uint256 rewardRecipientBalanceBefore = _getStakeAssetBalance(rewardRecipient);

        _expectStakeUnlock(validatorOwner, _weightToValue(validatorWeight));
        _expectRewardIssuance(rewardRecipient, expectedReward);

        _completeEndValidation(l1ValidatorRegistrationMessage);

        if (rewardRecipient == validatorOwner) {
            assertEq(
                _getStakeAssetBalance(validatorOwner),
                balanceBefore + _weightToValue(validatorWeight) + expectedReward
            );
        } else {
            assertEq(
                _getStakeAssetBalance(validatorOwner),
                balanceBefore + _weightToValue(validatorWeight)
            );

            assertEq(
                _getStakeAssetBalance(rewardRecipient),
                rewardRecipientBalanceBefore + expectedReward
            );
        }
    }

    function _completeEndValidation(bytes memory l1ValidatorRegistrationMessage) internal {
        _mockGetPChainWarpMessage(l1ValidatorRegistrationMessage, true);
        posValidatorManager.completeEndValidation(0);
    }

    function _completeEndDelegationWithChecks(
        bytes32 validationID,
        bytes32 delegationID,
        address delegator,
        uint64 delegatorWeight,
        uint256 expectedValidatorFees,
        uint256 expectedDelegatorReward,
        uint64 validatorWeight,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce,
        address rewardRecipient
    ) internal {
        bytes memory weightUpdateMessage = ValidatorMessages.packL1ValidatorWeightMessage(
            validationID, expectedNonce, validatorWeight
        );

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegationEnded(
            delegationID, validationID, expectedDelegatorReward, expectedValidatorFees
        );
        uint256 balanceBefore = _getStakeAssetBalance(delegator);
        uint256 rewardRecipientBalanceBefore = _getStakeAssetBalance(rewardRecipient);

        _expectStakeUnlock(delegator, _weightToValue(delegatorWeight));
        _expectRewardIssuance(rewardRecipient, expectedDelegatorReward);

        _completeEndDelegation(delegationID, weightUpdateMessage);

        assertEq(posValidatorManager.getValidator(validationID).weight, expectedValidatorWeight);

        if (rewardRecipient == delegator) {
            assertEq(
                _getStakeAssetBalance(delegator),
                balanceBefore + _weightToValue(delegatorWeight) + expectedDelegatorReward
            );
        } else {
            assertEq(
                _getStakeAssetBalance(delegator), balanceBefore + _weightToValue(delegatorWeight)
            );

            assertEq(
                _getStakeAssetBalance(rewardRecipient),
                rewardRecipientBalanceBefore + expectedDelegatorReward
            );
        }
    }

    function _completeEndDelegation(
        bytes32 delegationID,
        bytes memory weightUpdateMessage
    ) internal {
        _mockGetPChainWarpMessage(weightUpdateMessage, true);
        posValidatorManager.completeEndDelegation(delegationID, 0);
    }

    function _initializeAndCompleteEndDelegationWithChecks(
        bytes32 validationID,
        bytes32 delegationID
    ) internal {
        _initializeEndDelegationValidatorActiveWithChecks({
            validationID: validationID,
            sender: DEFAULT_DELEGATOR_ADDRESS,
            delegationID: delegationID,
            startDelegationTimestamp: DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            endDelegationTimestamp: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            includeUptime: true,
            force: false,
            rewardRecipient: address(0)
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: 0
        });

        uint256 expectedValidatorFees =
            _calculateValidatorFeesFromDelegator(expectedTotalReward, DEFAULT_DELEGATION_FEE_BIPS);
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;
        address delegator = DEFAULT_DELEGATOR_ADDRESS;

        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: delegator,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedValidatorFees: expectedValidatorFees,
            expectedDelegatorReward: expectedDelegatorReward,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2,
            rewardRecipient: delegator
        });
    }

    function _getStakeAssetBalance(address account) internal virtual returns (uint256);
    function _expectStakeUnlock(address account, uint256 amount) internal virtual;
    function _expectRewardIssuance(address account, uint256 amount) internal virtual;

    function _defaultDelegatorExpectedTotalReward() internal view returns (uint256) {
        return rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP
        });
    }

    function _defaultPoSSettings() internal pure returns (PoSValidatorManagerSettings memory) {
        return PoSValidatorManagerSettings({
            baseSettings: ValidatorManagerSettings({l1ID: DEFAULT_L1_ID, churnTracker: address(0)}),
            minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
            maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
            minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            minimumDelegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS,
            maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
            weightToValueFactor: DEFAULT_WEIGHT_TO_VALUE_FACTOR,
            rewardCalculator: IRewardCalculator(address(0)),
            uptimeBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID
        });
    }

    function _calculateValidatorFeesFromDelegator(
        uint256 totalReward,
        uint64 delegationFeeBips
    ) internal pure returns (uint256) {
        return totalReward * delegationFeeBips / 10000;
    }
}
