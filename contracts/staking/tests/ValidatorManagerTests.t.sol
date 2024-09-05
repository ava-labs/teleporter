// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ValidatorManager} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {ValidatorStatus} from "../interfaces/IValidatorManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
abstract contract ValidatorManagerTest is Test {
    bytes32 public constant P_CHAIN_BLOCKCHAIN_ID =
        bytes32(hex"0000000000000000000000000000000000000000000000000000000000000000");
    bytes32 public constant DEFAULT_SUBNET_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant DEFAULT_NODE_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_BLS_PUBLIC_KEY = bytes(
        hex"123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    uint64 public constant DEFAULT_WEIGHT = 1e6;
    uint256 public constant DEFAULT_MINIMUM_STAKE = 20;
    uint256 public constant DEFAULT_MAXIMUM_STAKE = 1e10;
    uint64 public constant DEFAULT_MINIMUM_STAKE_DURATION = 24 hours;
    uint8 public constant DEFAULT_MAXIMUM_HOURLY_CHURN = 0;
    uint64 public constant DEFAULT_EXPIRY = 1000;
    uint64 public constant DEFAULT_REGISTRATION_TIMESTAMP = 1000;
    uint64 public constant DEFAULT_COMPLETION_TIMESTAMP = 2000;

    ValidatorManager public validatorManager;

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
        uint256 endTime
    );

    event ValidationPeriodEnded(bytes32 indexed validationID, ValidatorStatus indexed status);

    receive() external payable {}
    fallback() external payable {}

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

    // The following tests call functions that are  implemented in ValidatorManager, but access state that's
    // only set in NativeTokenValidatorManager. Therefore we call them via the concrete type, rather than a
    // reference to the abstract type.
    function testResendRegisterValidatorMessage() public {
        bytes32 validationID = _setUpInitializeValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_EXPIRY,
            DEFAULT_BLS_PUBLIC_KEY
        );
        (, bytes memory registerSubnetValidatorMessage) = ValidatorMessages
            .packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                subnetID: DEFAULT_SUBNET_ID,
                nodeID: DEFAULT_NODE_ID,
                weight: DEFAULT_WEIGHT,
                registrationExpiry: DEFAULT_EXPIRY,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
            })
        );
        _mockSendWarpMessage(registerSubnetValidatorMessage, bytes32(0));
        validatorManager.resendRegisterValidatorMessage(validationID);
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
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 0, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        validatorManager.resendEndValidatorMessage(validationID);
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
            ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, false);

        _mockGetVerifiedWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodEnded(validationID, ValidatorStatus.Completed);

        validatorManager.completeEndValidation(0);
    }

    function testCompleteInvalidatedValidation() public {
        bytes32 validationID = _setUpInitializeValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_SUBNET_ID,
            DEFAULT_WEIGHT,
            DEFAULT_EXPIRY,
            DEFAULT_BLS_PUBLIC_KEY
        );
        bytes memory subnetValidatorRegistrationMessage =
            ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, false);

        _mockGetVerifiedWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodEnded(validationID, ValidatorStatus.Invalidated);

        validatorManager.completeEndValidation(0);
    }

    function _setUpInitializeValidatorRegistration(
        bytes32 nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory blsPublicKey
    ) internal returns (bytes32 validationID) {
        (validationID,) = ValidatorMessages.packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                nodeID: nodeID,
                subnetID: subnetID,
                weight: weight,
                registrationExpiry: registrationExpiry,
                blsPublicKey: blsPublicKey
            })
        );
        (, bytes memory registerSubnetValidatorMessage) = ValidatorMessages
            .packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                subnetID: subnetID,
                nodeID: nodeID,
                weight: weight,
                registrationExpiry: registrationExpiry,
                blsPublicKey: blsPublicKey
            })
        );
        vm.warp(registrationExpiry - 1);
        _mockSendWarpMessage(registerSubnetValidatorMessage, bytes32(0));

        _beforeSend(weight);
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodCreated(validationID, nodeID, bytes32(0), weight, registrationExpiry);

        _initializeValidatorRegistration(nodeID, registrationExpiry, blsPublicKey, weight);
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
            ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, true);

        _mockGetVerifiedWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.warp(registrationTimestamp);
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodRegistered(validationID, weight, registrationTimestamp);

        validatorManager.completeValidatorRegistration(0);
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
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 0, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(validationID, bytes32(0), weight, completionTimestamp);

        _initializeEndValidation(validationID);
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

    function _mockGetBlockchainID() internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_SOURCE_BLOCKCHAIN_ID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );
    }

    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory blsPublicKey,
        uint64 weight
    ) internal virtual returns (bytes32);

    function _initializeEndValidation(bytes32 validationID) internal virtual;

    function _beforeSend(uint64 weight) internal virtual;
}
// solhint-enable no-empty-blocks
