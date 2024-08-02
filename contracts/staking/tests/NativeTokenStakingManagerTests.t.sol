// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {StakingManagerTest} from "./StakingManagerTests.t.sol";
import {NativeTokenStakingManager} from "../NativeTokenStakingManager.sol";
import {StakingManagerSettings, InitialStakerInfo} from "../interfaces/IStakingManager.sol";
import {IRewardCalculator} from "../interfaces/IRewardCalculator.sol";
import {IWarpMessenger, WarpMessage} from "../StakingManager.sol";
import {StakingMessages} from "../StakingMessages.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
contract NativeTokenStakingManagerTest is StakingManagerTest {
    NativeTokenStakingManager public app;

    function setUp() public virtual {
        // Construct the object under test
        app = new NativeTokenStakingManager(
            StakingManagerSettings({
                pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                subnetID: DEFAULT_SUBNET_ID,
                minimumStakeAmount: DEFAULT_MINIMUM_STAKE,
                maximumStakeAmount: DEFAULT_MAXIMUM_STAKE,
                minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
                maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN,
                initialStakers: new InitialStakerInfo[](0),
                rewardCalculator: IRewardCalculator(address(0))
            })
        );
        stakingManager = app;
    }

    function testInitializeValidatorRegistrationSuccess() public {
        _setUpInitializeValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_EXPIRY,
            DEFAULT_ED25519_SIGNATURE
        );
    }

    function testInitializeValidatorRegistrationExcessiveChurn() public {
        // TODO: implement
    }
    function testInitializeValidatorRegistrationInsufficientStake() public {
        // TODO: implement
    }
    function testInitializeValidatorRegistrationExcessiveStake() public {
        // TODO: implement
    }
    function testInitializeValidatorRegistrationInsufficientDuration() public {
        // TODO: implement
    }

    // The following tests call functions that are  implemented in StakingManager, but access state that's
    // only set in NativeTokenStakingManager. Therefore we call them via the concrete type, rather than a
    // reference to the abstract type.
    function testResendRegisterValidatorMessage() public {
        bytes32 validationID = _setUpInitializeValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_EXPIRY,
            DEFAULT_ED25519_SIGNATURE
        );
        (, bytes memory registerSubnetValidatorMessage) = StakingMessages
            .packRegisterSubnetValidatorMessage(
            StakingMessages.ValidationInfo({
                subnetID: DEFAULT_SUBNET_ID,
                nodeID: DEFAULT_NODE_ID,
                weight: DEFAULT_WEIGHT,
                registrationExpiry: DEFAULT_EXPIRY,
                signature: DEFAULT_ED25519_SIGNATURE
            })
        );
        _mockSendWarpMessage(registerSubnetValidatorMessage, bytes32(0));
        app.resendRegisterValidatorMessage(validationID);
    }

    function testCompleteValidatorRegistration() public {
        _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            signature: DEFAULT_ED25519_SIGNATURE,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
    }

    function testInitializeEndValidation() public {
        _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            signature: DEFAULT_ED25519_SIGNATURE,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP
        });
    }

    function testInitializeEndValidationWithUptimeProof() public {
        // TODO: implement
    }

    function testInitializeEndValidationExcessiveChurn() public {
        // TODO: implement
    }

    function testResendEndValidation() public {
        bytes32 validationID = _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            signature: DEFAULT_ED25519_SIGNATURE,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP
        });
        bytes memory setValidatorWeightPayload =
            StakingMessages.packSetSubnetValidatorWeightMessage(validationID, 0, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        app.resendEndValidatorMessage(validationID);
    }

    function testCompleteEndValidation() public {
        bytes32 validationID = _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            signature: DEFAULT_ED25519_SIGNATURE,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP
        });

        bytes memory subnetValidatorRegistrationMessage =
            StakingMessages.packSubnetValidatorRegistrationMessage(validationID, false);

        _mockGetVerifiedWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(app));
        emit ValidationPeriodEnded(validationID);

        app.completeEndValidation(0, false);
    }

    function testCompleteEndValidationSetWeightMessageType() public {
        bytes32 validationID = _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            signature: DEFAULT_ED25519_SIGNATURE,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP
        });

        bytes memory setSubnetValidatorWeightMessage =
            StakingMessages.packSetSubnetValidatorWeightMessage(validationID, 1, 0);

        _mockGetVerifiedWarpMessage(setSubnetValidatorWeightMessage, true);

        vm.expectEmit(true, true, true, true, address(app));
        emit ValidationPeriodEnded(validationID);

        app.completeEndValidation(0, true);
    }

    // Helpers
    function _mockGetVerifiedWarpMessage(bytes memory expectedPayload, bool valid) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: P_CHAIN_BLOCKCHAIN_ID,
                    originSenderAddress: address(0),
                    payload: expectedPayload
                }),
                valid
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );
    }

    function _mockSendWarpMessage(bytes memory payload, bytes32 expectedMessageID) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encode(IWarpMessenger.sendWarpMessage.selector),
            abi.encode(expectedMessageID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.sendWarpMessage, payload)
        );
    }

    function _setUpInitializeValidatorRegistration(
        bytes32 nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory signature
    ) internal returns (bytes32 validationID) {
        (validationID,) = StakingMessages.packValidationInfo(
            StakingMessages.ValidationInfo({
                nodeID: nodeID,
                subnetID: subnetID,
                weight: weight,
                registrationExpiry: registrationExpiry,
                signature: signature
            })
        );
        (, bytes memory registerSubnetValidatorMessage) = StakingMessages
            .packRegisterSubnetValidatorMessage(
            StakingMessages.ValidationInfo({
                subnetID: subnetID,
                nodeID: nodeID,
                weight: weight,
                registrationExpiry: registrationExpiry,
                signature: signature
            })
        );
        vm.warp(DEFAULT_EXPIRY - 1);
        _mockSendWarpMessage(registerSubnetValidatorMessage, bytes32(0));
        vm.expectEmit(true, true, true, true, address(app));
        emit ValidationPeriodCreated(
            validationID, DEFAULT_NODE_ID, bytes32(0), weight, DEFAULT_EXPIRY
        );

        app.initializeValidatorRegistration{value: app.weightToValue(weight)}(
            nodeID, registrationExpiry, signature
        );
    }

    function _setUpCompleteValidatorRegistration(
        bytes32 nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory signature,
        uint64 registrationTimestamp
    ) internal returns (bytes32 validationID) {
        validationID = _setUpInitializeValidatorRegistration(
            nodeID, subnetID, weight, registrationExpiry, signature
        );
        bytes memory subnetValidatorRegistrationMessage =
            StakingMessages.packSubnetValidatorRegistrationMessage(validationID, true);

        _mockGetVerifiedWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.warp(registrationTimestamp);
        vm.expectEmit(true, true, true, true, address(app));
        emit ValidationPeriodRegistered(validationID, weight, registrationTimestamp);

        app.completeValidatorRegistration(0);
    }

    function _setUpInitializeEndValidation(
        bytes32 nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory signature,
        uint64 registrationTimestamp,
        uint64 completionTimestamp
    ) internal returns (bytes32 validationID) {
        validationID = _setUpCompleteValidatorRegistration({
            nodeID: nodeID,
            subnetID: subnetID,
            weight: weight,
            registrationExpiry: registrationExpiry,
            signature: signature,
            registrationTimestamp: registrationTimestamp
        });

        vm.warp(completionTimestamp);
        bytes memory setValidatorWeightPayload =
            StakingMessages.packSetSubnetValidatorWeightMessage(validationID, 0, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.expectEmit(true, true, true, true, address(app));
        emit ValidatorRemovalInitialized(validationID, bytes32(0), weight, completionTimestamp, 0);

        app.initializeEndValidation(validationID, false, 0);
    }
}
// TODO: Remove this once all unit tests implemented
// solhint-enable no-empty-blocks
