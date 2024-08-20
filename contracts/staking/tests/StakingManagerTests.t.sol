// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {StakingManager} from "../StakingManager.sol";
import {StakingMessages} from "../StakingMessages.sol";
import {IWarpMessenger, WarpMessage} from "../StakingManager.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
abstract contract StakingManagerTest is Test {
    bytes32 public constant P_CHAIN_BLOCKCHAIN_ID =
        bytes32(hex"0000000000000000000000000000000000000000000000000000000000000000");
    bytes32 public constant DEFAULT_SUBNET_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant DEFAULT_NODE_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_BLS_PUBLIC_KEY = bytes(
        hex"123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    uint64 public constant DEFAULT_WEIGHT = 1e6;
    uint256 public constant DEFAULT_MINIMUM_STAKE = 20;
    uint256 public constant DEFAULT_MAXIMUM_STAKE = 1e10;
    uint64 public constant DEFAULT_MINIMUM_STAKE_DURATION = 24 hours;
    uint8 public constant DEFAULT_MAXIMUM_HOURLY_CHURN = 0;
    uint64 public constant DEFAULT_EXPIRY = 1000;
    uint64 public constant DEFAULT_REGISTRATION_TIMESTAMP = 1000;
    uint64 public constant DEFAULT_COMPLETION_TIMESTAMP = 2000;

    StakingManager public stakingManager;

    event ValidationPeriodCreated(
        bytes32 indexed validationID,
        bytes32 indexed nodeID,
        bytes32 indexed registerValidationMessageID,
        uint256 weight,
        uint64 registrationExpiry
    );

    event ValidationPeriodRegistered(
        bytes32 indexed validationID, uint256 stakeAmount, uint256 timestamp
    );

    event ValidatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        uint256 stakeAmount,
        uint256 endTime,
        uint64 uptime
    );

    event ValidationPeriodEnded(bytes32 indexed validationID);

    function testInitializeValidatorRegistrationSuccess() public {
        _setUpInitializeValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_EXPIRY,
            DEFAULT_BLS_PUBLIC_KEY
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
            DEFAULT_BLS_PUBLIC_KEY
        );
        (, bytes memory registerSubnetValidatorMessage) = StakingMessages
            .packRegisterSubnetValidatorMessage(
            StakingMessages.ValidationInfo({
                subnetID: DEFAULT_SUBNET_ID,
                nodeID: DEFAULT_NODE_ID,
                weight: DEFAULT_WEIGHT,
                registrationExpiry: DEFAULT_EXPIRY,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
            })
        );
        _mockSendWarpMessage(registerSubnetValidatorMessage, bytes32(0));
        stakingManager.resendRegisterValidatorMessage(validationID);
    }

    function testCompleteValidatorRegistration() public {
        _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
    }

    function testInitializeEndValidation() public {
        _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
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
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP
        });
        bytes memory setValidatorWeightPayload =
            StakingMessages.packSetSubnetValidatorWeightMessage(validationID, 0, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        stakingManager.resendEndValidatorMessage(validationID);
    }

    function testCompleteEndValidation() public {
        bytes32 validationID = _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP
        });

        bytes memory subnetValidatorRegistrationMessage =
            StakingMessages.packSubnetValidatorRegistrationMessage(validationID, false);

        _mockGetVerifiedWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidationPeriodEnded(validationID);

        stakingManager.completeEndValidation(0, false);
    }

    function testCompleteEndValidationSetWeightMessageType() public {
        bytes32 validationID = _setUpInitializeEndValidation({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP,
            completionTimestamp: DEFAULT_COMPLETION_TIMESTAMP
        });

        bytes memory setSubnetValidatorWeightMessage =
            StakingMessages.packSetSubnetValidatorWeightMessage(validationID, 1, 0);

        _mockGetVerifiedWarpMessage(setSubnetValidatorWeightMessage, true);

        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidationPeriodEnded(validationID);

        stakingManager.completeEndValidation(0, true);
    }

    function testValueToWeight() public view {
        uint64 w1 = stakingManager.valueToWeight(1e12);
        uint64 w2 = stakingManager.valueToWeight(1e18);
        uint64 w3 = stakingManager.valueToWeight(1e27);

        assertEq(w1, 1);
        assertEq(w2, 1e6);
        assertEq(w3, 1e15);
    }

    function testWeightToValue() public view {
        uint256 v1 = stakingManager.weightToValue(1);
        uint256 v2 = stakingManager.weightToValue(1e6);
        uint256 v3 = stakingManager.weightToValue(1e15);

        assertEq(v1, 1e12);
        assertEq(v2, 1e18);
        assertEq(v3, 1e27);
    }

    function _setUpInitializeValidatorRegistration(
        bytes32 nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory blsPublicKey
    ) internal returns (bytes32 validationID) {
        (validationID,) = StakingMessages.packRegisterSubnetValidatorMessage(
            StakingMessages.ValidationInfo({
                nodeID: nodeID,
                subnetID: subnetID,
                weight: weight,
                registrationExpiry: registrationExpiry,
                blsPublicKey: blsPublicKey
            })
        );
        (, bytes memory registerSubnetValidatorMessage) = StakingMessages
            .packRegisterSubnetValidatorMessage(
            StakingMessages.ValidationInfo({
                subnetID: subnetID,
                nodeID: nodeID,
                weight: weight,
                registrationExpiry: registrationExpiry,
                blsPublicKey: blsPublicKey
            })
        );
        vm.warp(DEFAULT_EXPIRY - 1);
        _mockSendWarpMessage(registerSubnetValidatorMessage, bytes32(0));

        _beforeSend(stakingManager.weightToValue(weight));
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidationPeriodCreated(
            validationID, DEFAULT_NODE_ID, bytes32(0), weight, DEFAULT_EXPIRY
        );

        _initializeValidatorRegistration(
            nodeID, registrationExpiry, blsPublicKey, stakingManager.weightToValue(weight)
        );
    }

    function _setUpCompleteValidatorRegistration(
        bytes32 nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory blsPublicKey,
        uint64 registrationTimestamp
    ) internal returns (bytes32 validationID) {
        validationID = _setUpInitializeValidatorRegistration(
            nodeID, subnetID, weight, registrationExpiry, blsPublicKey
        );
        bytes memory subnetValidatorRegistrationMessage =
            StakingMessages.packSubnetValidatorRegistrationMessage(validationID, true);

        _mockGetVerifiedWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.warp(registrationTimestamp);
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidationPeriodRegistered(validationID, weight, registrationTimestamp);

        stakingManager.completeValidatorRegistration(0);
    }

    function _setUpInitializeEndValidation(
        bytes32 nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory blsPublicKey,
        uint64 registrationTimestamp,
        uint64 completionTimestamp
    ) internal returns (bytes32 validationID) {
        validationID = _setUpCompleteValidatorRegistration({
            nodeID: nodeID,
            subnetID: subnetID,
            weight: weight,
            registrationExpiry: registrationExpiry,
            blsPublicKey: blsPublicKey,
            registrationTimestamp: registrationTimestamp
        });

        vm.warp(completionTimestamp);
        bytes memory setValidatorWeightPayload =
            StakingMessages.packSetSubnetValidatorWeightMessage(validationID, 0, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.expectEmit(true, true, true, true, address(stakingManager));
        emit ValidatorRemovalInitialized(validationID, bytes32(0), weight, completionTimestamp, 0);

        stakingManager.initializeEndValidation(validationID, false, 0);
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

    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory blsPublicKey,
        uint256 stakeAmount
    ) internal virtual returns (bytes32);

    function _beforeSend(uint256 value) internal virtual;
}
// solhint-enable no-empty-blocks
