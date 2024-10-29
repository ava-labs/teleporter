// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {PoSValidatorManager} from "../PoSValidatorManager.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {ValidatorRegistrationInput, ValidatorStatus} from "../interfaces/IValidatorManager.sol";

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
    uint256 public constant SECONDS_IN_YEAR = 31536000;

    PoSValidatorManager public posValidatorManager;
    IRewardCalculator public rewardCalculator;

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
        uint64 validatorWeight,
        bytes32 setWeightMessageID
    );

    event DelegationEnded(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees
    );

    event UptimeUpdated(bytes32 indexed validationID, uint64 uptime);

    event ValidationRewardsClaimed(bytes32 indexed validationID, uint256 indexed reward);

    function testDelegationFeeBipsTooLow() public {
        ValidatorRegistrationInput memory registrationInput = ValidatorRegistrationInput({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationExpiry: DEFAULT_EXPIRY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegationFee.selector,
                DEFAULT_MINIMUM_DELEGATION_FEE_BIPS - 1
            )
        );
        _initializeValidatorRegistration(
            registrationInput,
            DEFAULT_MINIMUM_DELEGATION_FEE_BIPS - 1,
            DEFAULT_MINIMUM_STAKE_DURATION,
            DEFAULT_MINIMUM_STAKE_AMOUNT
        );
    }

    function testDelegationFeeBipsTooHigh() public {
        ValidatorRegistrationInput memory registrationInput = ValidatorRegistrationInput({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationExpiry: DEFAULT_EXPIRY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });
        uint16 delegationFeeBips = posValidatorManager.MAXIMUM_DELEGATION_FEE_BIPS() + 1;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidDelegationFee.selector, delegationFeeBips
            )
        );

        _initializeValidatorRegistration(
            registrationInput,
            delegationFeeBips,
            DEFAULT_MINIMUM_STAKE_DURATION,
            DEFAULT_MINIMUM_STAKE_AMOUNT
        );
    }

    function testInvalidMinStakeDuration() public {
        ValidatorRegistrationInput memory registrationInput = ValidatorRegistrationInput({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationExpiry: DEFAULT_EXPIRY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidMinStakeDuration.selector,
                DEFAULT_MINIMUM_STAKE_DURATION - 1
            )
        );
        _initializeValidatorRegistration(
            registrationInput,
            DEFAULT_DELEGATION_FEE_BIPS,
            DEFAULT_MINIMUM_STAKE_DURATION - 1,
            DEFAULT_MINIMUM_STAKE_AMOUNT
        );
    }

    function testStakeAmountTooLow() public {
        ValidatorRegistrationInput memory registrationInput = ValidatorRegistrationInput({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationExpiry: DEFAULT_EXPIRY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidStakeAmount.selector, DEFAULT_MINIMUM_STAKE_AMOUNT - 1
            )
        );
        _initializeValidatorRegistration(
            registrationInput,
            DEFAULT_DELEGATION_FEE_BIPS,
            DEFAULT_MINIMUM_STAKE_DURATION,
            DEFAULT_MINIMUM_STAKE_AMOUNT - 1
        );
    }

    function testStakeAmountTooHigh() public {
        ValidatorRegistrationInput memory registrationInput = ValidatorRegistrationInput({
            nodeID: DEFAULT_NODE_ID,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationExpiry: DEFAULT_EXPIRY,
            remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
            disableOwner: DEFAULT_P_CHAIN_OWNER
        });
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.InvalidStakeAmount.selector, DEFAULT_MAXIMUM_STAKE_AMOUNT + 1
            )
        );
        _initializeValidatorRegistration(
            registrationInput,
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
        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION);
        vm.expectRevert(ValidatorManager.InvalidWarpMessage.selector);
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeChainID() public {
        bytes32 validationID = _registerDefaultValidator();

        _mockGetUptimeWarpMessage(new bytes(0), true);
        _mockGetBlockchainID(posValidatorManager.P_CHAIN_BLOCKCHAIN_ID());
        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION);
        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidWarpSourceChainID.selector, DEFAULT_SOURCE_BLOCKCHAIN_ID
            )
        );
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeSenderAddress() public {
        bytes32 validationID = _registerDefaultValidator();

        _mockGetBlockchainID();
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

        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION);
        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorManager.InvalidWarpOriginSenderAddress.selector, address(this)
            )
        );
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeValidationID() public {
        bytes32 validationID = _registerDefaultValidator();

        _mockGetBlockchainID();
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

        vm.warp(DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION);
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
        bytes memory setValidatorWeightPayload = ValidatorMessages.packSubnetValidatorWeightMessage(
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
        bytes memory setValidatorWeightPayload = ValidatorMessages.packSubnetValidatorWeightMessage(
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
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID1,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            2
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
            force: false
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
            force: false
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
            force: false
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
            force: true
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
            force: true
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
            force: false
        });
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightMessage(validationID, 2, DEFAULT_WEIGHT);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendUpdateDelegation(delegationID);
    }

    function testResendEndValidation() public override {
        bytes32 validationID = _registerDefaultValidator();
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            expectedNonce: 1,
            includeUptime: true,
            force: false
        });

        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        validatorManager.resendEndValidatorMessage(validationID);
    }

    function testCompleteEndDelegation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);
        _initializeAndCompleteEndDelegationWithChecks(validationID, delegationID);
    }

    // Delegator registration is not allowed when Validator is pending removed.
    function testInitializeDelegatorRegistrationValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();

        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION,
            expectedNonce: 1,
            includeUptime: true,
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

        uint64 validatorEndTime = DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION;
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: validatorEndTime,
            expectedNonce: 2,
            includeUptime: true,
            force: false
        });

        _setUpCompleteDelegatorRegistrationWithChecks(
            validationID,
            delegationID,
            validatorEndTime + 1,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1
        );
    }

    // Delegator cannot initialize end delegation when validator is pending removed.
    function testInitializeEndDelegationValidatorPendingRemoved() public {
        bytes32 validationID = _registerDefaultValidator();

        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_REGISTRATION_TIMESTAMP + DEFAULT_MINIMUM_STAKE_DURATION,
            expectedNonce: 1,
            includeUptime: true,
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
            force: false
        });

        uint64 validationEndTime = DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1;
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: validationEndTime,
            expectedNonce: 3,
            includeUptime: true,
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });

        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedTotalReward: expectedTotalReward,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: 0,
            expectedNonce: 2
        });
    }

    function testInitializeDelegatorRegistrationValidatorCompleted() public {
        bytes32 validationID = _registerDefaultValidator();
        _endDefaultValidator(validationID, 1);

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

        _endDefaultValidator(validationID, 2);

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

        _endDefaultValidator(validationID, 2);

        uint64 delegationEndTime = DEFAULT_COMPLETION_TIMESTAMP + 1;

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
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
            force: false
        });

        _endDefaultValidator(validationID, 3);

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });

        uint256 expectedValidatorFees = expectedTotalReward * DEFAULT_DELEGATION_FEE_BIPS / 10000;
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
            force: false
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
            force: false
        });

        // Complete ending delegator2 with delegator1's nonce
        // Note that ending delegator1 with delegator2's nonce is valid
        uint64 nonce = 3;
        bytes memory setValidatorWeightPayload = ValidatorMessages.packSubnetValidatorWeightMessage(
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
            force: false
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
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });

        // Complete delegation1 by delivering the weight update from nonce 4 (delegator2's nonce)
        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID1,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedTotalReward: expectedTotalReward,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 4
        });
    }

    function testCompleteEndValidation() public virtual override {
        bytes32 validationID = _registerDefaultValidator();
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            expectedNonce: 1,
            includeUptime: true,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_COMPLETION_TIMESTAMP,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            expectedReward: expectedReward,
            validatorWeight: DEFAULT_WEIGHT
        });
    }

    function testInitializeEndValidation() public virtual override {
        bytes32 validationID = _registerDefaultValidator();
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            expectedNonce: 1,
            includeUptime: true,
            force: false
        });
    }

    function testInitializeEndValidationUseStoredUptime() public {
        bytes32 validationID = _registerDefaultValidator();

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        // Submit an uptime proof via submitUptime
        uint64 uptimePercentage1 = 80;
        uint64 uptime1 = (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP)
            * uptimePercentage1 / 100;
        bytes memory uptimeMsg1 =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime1);
        _mockGetUptimeWarpMessage(uptimeMsg1, true);
        _mockGetBlockchainID();

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit UptimeUpdated(validationID, uptime1);
        posValidatorManager.submitUptimeProof(validationID, 0);

        // Submit a second uptime proof via initializeEndValidation. This one is not sufficient for rewards
        // Submit an uptime proof via submitUptime
        uint64 uptimePercentage2 = 79;
        uint64 uptime2 = (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP)
            * uptimePercentage2 / 100;
        bytes memory uptimeMsg2 =
            ValidatorMessages.packValidationUptimeMessage(validationID, uptime2);
        _mockGetUptimeWarpMessage(uptimeMsg2, true);
        _mockGetBlockchainID();

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _initializeEndValidation(validationID, true);
    }

    function testInitializeEndValidationInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 79;

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );
        _mockGetUptimeWarpMessage(uptimeMsg, true);
        _mockGetBlockchainID();

        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.ValidatorIneligibleForRewards.selector, validationID
            )
        );

        _initializeEndValidation(validationID, true);
    }

    function testInitializeEndValidationPoAValidator() public {
        bytes32 defaultInitialValidationID = sha256(abi.encodePacked(DEFAULT_SUBNET_ID, uint32(1)));

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightMessage(defaultInitialValidationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(
            defaultInitialValidationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _initializeEndValidation(defaultInitialValidationID, false);
    }

    function testForceInitializeEndValidation() public {
        bytes32 validationID = _registerDefaultValidator();
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            expectedNonce: 1,
            includeUptime: true,
            force: true
        });
    }

    function testForceInitializeEndValidationInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 79;

        vm.warp(DEFAULT_COMPLETION_TIMESTAMP);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );
        _mockGetUptimeWarpMessage(uptimeMsg, true);
        _mockGetBlockchainID();

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(
            validationID, bytes32(0), DEFAULT_WEIGHT, DEFAULT_COMPLETION_TIMESTAMP
        );

        _forceInitializeEndValidation(validationID, true);
    }

    function testClaimValidationReward() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 80;
        uint64 claimTime = DEFAULT_COMPLETION_TIMESTAMP;
        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: claimTime,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, (claimTime - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );
        _mockGetUptimeWarpMessage(uptimeMsg, true);
        _mockGetBlockchainID();

        vm.warp(claimTime);
        _expectValidationRewardsIssuance(validationID, address(this), expectedReward);
        posValidatorManager.claimValidationRewards(validationID, 0);
    }

    function testClaimSubsequentValidationReward() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 80;
        uint64 firstClaimTime = DEFAULT_COMPLETION_TIMESTAMP;
        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, (firstClaimTime - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );
        uint256 claimedReward = _claimRewards({
            validationID: validationID,
            uptimeMsg: uptimeMsg,
            validationStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            claimTime: firstClaimTime,
            claimedReward: 0,
            success: true
        });

        // Claim rewards again
        uint64 secondClaimTime = firstClaimTime + 1000;
        uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            (secondClaimTime - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );
        _claimRewards({
            validationID: validationID,
            uptimeMsg: uptimeMsg,
            validationStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            claimTime: secondClaimTime,
            claimedReward: claimedReward,
            success: true
        });
    }

    function testClaimValidationRewardStaleUptime() public {
        uint64 validationStartTime = 0;
        bytes32 validationID = _registerValidator({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: validationStartTime
        });
        uint64 uptimePercentage = 80;
        uint64 firstClaimTime = DEFAULT_MINIMUM_VALIDATION_DURATION + 1; // 24 hours

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, ((firstClaimTime - validationStartTime) * uptimePercentage / 100) + 1
        );
        uint256 claimedReward = _claimRewards({
            validationID: validationID,
            uptimeMsg: uptimeMsg,
            validationStartTime: validationStartTime,
            claimTime: firstClaimTime,
            claimedReward: 0,
            success: true
        });

        // Attempt to claim rewards again, but by submitting the same uptime
        uint64 secondClaimTime = firstClaimTime + 5 hours;
        _claimRewards({
            validationID: validationID,
            uptimeMsg: uptimeMsg,
            validationStartTime: validationStartTime,
            claimTime: secondClaimTime,
            claimedReward: claimedReward,
            success: false
        });
    }

    function testClaimValidationRewardInsufficientUptime() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 79;

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            (DEFAULT_COMPLETION_TIMESTAMP - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );
        _claimRewards({
            validationID: validationID,
            uptimeMsg: uptimeMsg,
            validationStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            claimTime: DEFAULT_COMPLETION_TIMESTAMP,
            claimedReward: 0,
            success: false
        });
    }

    function testClaimValidationRewardWithDelegation() public {
        bytes32 validationID = _registerDefaultValidator();
        bytes32 delegationID = _registerDefaultDelegator(validationID);

        // Claim validator rewards
        uint64 uptimePercentage = 80;
        uint64 claimTime = DEFAULT_COMPLETION_TIMESTAMP;
        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, (claimTime - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );

        _claimRewards({
            validationID: validationID,
            uptimeMsg: uptimeMsg,
            validationStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            claimTime: claimTime,
            claimedReward: 0,
            success: true
        });

        // End the delegation and claim rewards
        _initializeAndCompleteEndDelegationWithChecks(validationID, delegationID);
    }

    function testEndValidationPreviouslyClaimedReward() public {
        bytes32 validationID = _registerDefaultValidator();
        uint64 uptimePercentage = 80;
        uint64 firstClaimTime = DEFAULT_COMPLETION_TIMESTAMP;

        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, (firstClaimTime - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );
        uint256 firstExpectedReward = _claimRewards({
            validationID: validationID,
            uptimeMsg: uptimeMsg,
            validationStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            claimTime: firstClaimTime,
            claimedReward: 0,
            success: true
        });

        // Claim rewards again
        uint64 secondClaimTime = firstClaimTime + 5 hours;
        uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID,
            (secondClaimTime - DEFAULT_REGISTRATION_TIMESTAMP) * uptimePercentage / 100
        );
        uint256 secondExpectedReward = _claimRewards({
            validationID: validationID,
            uptimeMsg: uptimeMsg,
            validationStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            claimTime: secondClaimTime,
            claimedReward: firstExpectedReward,
            success: true
        });

        // Try to end the validation with a stale uptime
        uint64 thirdClaimTime = secondClaimTime + 5 hours;
        vm.expectRevert(
            abi.encodeWithSelector(
                PoSValidatorManager.ValidatorIneligibleForRewards.selector, validationID
            )
        );
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: thirdClaimTime,
            expectedNonce: 1,
            includeUptime: false,
            force: false
        });

        // Try again with the correct uptime
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: thirdClaimTime,
            expectedNonce: 1,
            includeUptime: true,
            force: false
        });

        uint256 thirdExpectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: thirdClaimTime,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });
        thirdExpectedReward = thirdExpectedReward - firstExpectedReward - secondExpectedReward;

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            expectedReward: thirdExpectedReward,
            validatorWeight: DEFAULT_WEIGHT
        });

        // Confirm the pro-rated rewards match the total expected reward
        uint256 totalExpectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_WEIGHT),
            validatorStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingStartTime: DEFAULT_REGISTRATION_TIMESTAMP,
            stakingEndTime: thirdClaimTime,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });
        assertEq(
            totalExpectedReward, firstExpectedReward + secondExpectedReward + thirdExpectedReward
        );
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
        bool includeUptime
    ) internal virtual override {
        posValidatorManager.initializeEndValidation(validationID, includeUptime, 0);
    }

    function _forceInitializeEndValidation(
        bytes32 validationID,
        bool includeUptime
    ) internal virtual override {
        posValidatorManager.forceInitializeEndValidation(validationID, includeUptime, 0);
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
        bytes memory setValidatorWeightPayload = ValidatorMessages.packSubnetValidatorWeightMessage(
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
        bytes32 validationID,
        bytes32 delegationID,
        uint64 completeRegistrationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packSubnetValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
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
        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRegistered({
            delegationID: delegationID,
            validationID: validationID,
            startTime: completeRegistrationTimestamp
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID,
            completeRegistrationTimestamp,
            expectedValidatorWeight,
            expectedNonce
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
        _setUpCompleteDelegatorRegistration({
            validationID: validationID,
            delegationID: delegationID,
            completeRegistrationTimestamp: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            expectedValidatorWeight: DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            expectedNonce: 1
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
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegationID,
            completeRegistrationTimestamp,
            expectedValidatorWeight,
            expectedNonce
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
        bool force
    ) internal {
        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit ValidatorWeightUpdate({
            validationID: validationID,
            nonce: expectedNonce,
            validatorWeight: expectedValidatorWeight,
            setWeightMessageID: bytes32(0)
        });

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRemovalInitialized({delegationID: delegationID, validationID: validationID});

        _initializeEndDelegationValidatorActive({
            validationID: validationID,
            sender: sender,
            delegationID: delegationID,
            startDelegationTimestamp: startDelegationTimestamp,
            endDelegationTimestamp: endDelegationTimestamp,
            expectedValidatorWeight: expectedValidatorWeight,
            expectedNonce: expectedNonce,
            includeUptime: includeUptime,
            force: force
        });
    }

    function _initializeEndDelegationValidatorActive(
        bytes32 validationID,
        address sender,
        bytes32 delegationID,
        uint64 startDelegationTimestamp,
        uint64 endDelegationTimestamp,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce,
        bool includeUptime,
        bool force
    ) internal {
        bytes memory setValidatorWeightPayload = ValidatorMessages.packSubnetValidatorWeightMessage(
            validationID, expectedNonce, expectedValidatorWeight
        );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
            validationID, endDelegationTimestamp - startDelegationTimestamp
        );
        if (includeUptime) {
            _mockGetUptimeWarpMessage(uptimeMsg, true);
            _mockGetBlockchainID();
        }
        _initializeEndDelegation(sender, delegationID, endDelegationTimestamp, includeUptime, force);
    }

    function _initializeEndDelegation(
        address sender,
        bytes32 delegationID,
        uint64 endDelegationTimestamp,
        bool includeUptime,
        bool force
    ) internal {
        vm.warp(endDelegationTimestamp);
        vm.prank(sender);
        if (force) {
            posValidatorManager.forceInitializeEndDelegation(delegationID, includeUptime, 0);
        } else {
            posValidatorManager.initializeEndDelegation(delegationID, includeUptime, 0);
        }
    }

    function _endDefaultValidator(bytes32 validationID, uint64 expectedNonce) internal {
        _endValidationWithChecks({
            validationID: validationID,
            validatorOwner: address(this),
            completeRegistrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP,
            validatorWeight: DEFAULT_WEIGHT,
            expectedNonce: expectedNonce
        });
    }

    function _endValidationWithChecks(
        bytes32 validationID,
        address validatorOwner,
        uint64 completeRegistrationTimestamp,
        uint64 completionTimestamp,
        uint64 validatorWeight,
        uint64 expectedNonce
    ) internal {
        _initializeEndValidation({
            validationID: validationID,
            registrationTimestamp: completeRegistrationTimestamp,
            completionTimestamp: completionTimestamp,
            expectedNonce: expectedNonce,
            includeUptime: true,
            force: false
        });

        uint256 expectedReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(validatorWeight),
            validatorStartTime: 0,
            stakingStartTime: completeRegistrationTimestamp,
            stakingEndTime: completionTimestamp,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });

        _completeEndValidationWithChecks({
            validationID: validationID,
            validatorOwner: validatorOwner,
            expectedReward: expectedReward,
            validatorWeight: validatorWeight
        });
    }

    function _expectValidationRewardsIssuance(
        bytes32 validationID,
        address account,
        uint256 amount
    ) internal {
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationRewardsClaimed(validationID, amount);
        _expectRewardIssuance(account, amount);
    }

    function _completeEndValidationWithChecks(
        bytes32 validationID,
        address validatorOwner,
        uint256 expectedReward,
        uint64 validatorWeight
    ) internal {
        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit ValidationPeriodEnded(validationID, ValidatorStatus.Completed);
        uint256 balanceBefore = _getStakeAssetBalance(validatorOwner);

        _expectStakeUnlock(validatorOwner, _weightToValue(validatorWeight));
        _expectValidationRewardsIssuance(validationID, validatorOwner, expectedReward);

        _completeEndValidation(validationID);

        assertEq(
            _getStakeAssetBalance(validatorOwner),
            balanceBefore + _weightToValue(validatorWeight) + expectedReward
        );
    }

    function _completeEndValidation(bytes32 validationID) internal {
        bytes memory subnetValidatorRegistrationMessage =
            ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, false);
        _mockGetPChainWarpMessage(subnetValidatorRegistrationMessage, true);

        posValidatorManager.completeEndValidation(0);
    }

    function _completeEndDelegationWithChecks(
        bytes32 validationID,
        bytes32 delegationID,
        address delegator,
        uint64 delegatorWeight,
        uint256 expectedTotalReward,
        uint64 delegationFeeBips,
        uint64 validatorWeight,
        uint64 expectedValidatorWeight,
        uint64 expectedNonce
    ) internal {
        uint256 expectedValidatorFees = expectedTotalReward * delegationFeeBips / 10000;
        uint256 expectedDelegatorReward = expectedTotalReward - expectedValidatorFees;

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegationEnded(
            delegationID, validationID, expectedDelegatorReward, expectedValidatorFees
        );
        uint256 balanceBefore = _getStakeAssetBalance(delegator);

        _expectStakeUnlock(delegator, _weightToValue(delegatorWeight));
        _expectRewardIssuance(delegator, expectedDelegatorReward);

        _completeEndDelegation(validationID, delegationID, validatorWeight, expectedNonce);

        assertEq(posValidatorManager.getWeight(validationID), expectedValidatorWeight);
        assertEq(
            _getStakeAssetBalance(delegator),
            balanceBefore + _weightToValue(delegatorWeight) + expectedDelegatorReward
        );
    }

    function _completeEndDelegation(
        bytes32 validationID,
        bytes32 delegationID,
        uint64 validatorWeight,
        uint64 expectedNonce
    ) internal {
        bytes memory weightUpdateMessage = ValidatorMessages.packSubnetValidatorWeightMessage(
            validationID, expectedNonce, validatorWeight
        );
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
            force: false
        });

        uint256 expectedTotalReward = rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(DEFAULT_DELEGATOR_WEIGHT),
            validatorStartTime: 0,
            stakingStartTime: DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            stakingEndTime: DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP,
            uptimeSeconds: 0,
            initialSupply: 0,
            endSupply: 0
        });

        _completeEndDelegationWithChecks({
            validationID: validationID,
            delegationID: delegationID,
            delegator: DEFAULT_DELEGATOR_ADDRESS,
            delegatorWeight: DEFAULT_DELEGATOR_WEIGHT,
            expectedTotalReward: expectedTotalReward,
            delegationFeeBips: DEFAULT_DELEGATION_FEE_BIPS,
            validatorWeight: DEFAULT_WEIGHT,
            expectedValidatorWeight: DEFAULT_WEIGHT,
            expectedNonce: 2
        });
    }

    function _claimRewards(
        bytes32 validationID,
        bytes memory uptimeMsg,
        uint64 validationStartTime,
        uint64 claimTime,
        uint256 claimedReward,
        bool success
    ) internal returns (uint256) {
        uint256 expectedReward;
        if (success) {
            expectedReward = rewardCalculator.calculateReward({
                stakeAmount: _weightToValue(DEFAULT_WEIGHT),
                validatorStartTime: validationStartTime,
                stakingStartTime: validationStartTime,
                stakingEndTime: claimTime,
                uptimeSeconds: 0,
                initialSupply: 0,
                endSupply: 0
            });
            expectedReward -= claimedReward;
            _expectValidationRewardsIssuance(validationID, address(this), expectedReward);
        } else {
            vm.expectRevert(
                abi.encodeWithSelector(
                    PoSValidatorManager.ValidatorIneligibleForRewards.selector, validationID
                )
            );
        }

        _mockGetUptimeWarpMessage(uptimeMsg, true);
        _mockGetBlockchainID();
        vm.warp(claimTime);
        posValidatorManager.claimValidationRewards(validationID, 0);
        return expectedReward;
    }

    function _getStakeAssetBalance(address account) internal virtual returns (uint256);
    function _expectStakeUnlock(address account, uint256 amount) internal virtual;
    function _expectRewardIssuance(address account, uint256 amount) internal virtual;
}
