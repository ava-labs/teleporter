// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ValidatorManager, SubnetConversionData, InitialValidator} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {ValidatorStatus, ValidatorRegistrationInput} from "../interfaces/IValidatorManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

// TODO: Remove this once all unit tests implemented
// solhint-disable no-empty-blocks
abstract contract ValidatorManagerTest is Test {
    bytes32 public constant DEFAULT_SUBNET_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant DEFAULT_NODE_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant DEFAULT_INITIAL_VALIDATOR_NODE_ID =
        bytes32(hex"2345678123456781234567812345678123456781234567812345678123456781");
    bytes public constant DEFAULT_BLS_PUBLIC_KEY = bytes(
        hex"123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_SUBNET_CONVERSION_TX_ID =
        bytes32(hex"1156c1551ee48fa5c2e4916b6ab24b526d7fcff07f3923077ad596325482cb4d");
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    uint64 public constant DEFAULT_WEIGHT = 1e6;
    uint64 public constant DEFAULT_INITIAL_VALIDATOR_WEIGHT = DEFAULT_WEIGHT * 1e4;
    uint256 public constant DEFAULT_MINIMUM_STAKE_AMOUNT = 20e12;
    uint256 public constant DEFAULT_MAXIMUM_STAKE_AMOUNT = 1e22;
    uint64 public constant DEFAULT_CHURN_PERIOD = 1 hours;
    uint8 public constant DEFAULT_MAXIMUM_CHURN_PERCENTAGE = 20;
    uint64 public constant DEFAULT_EXPIRY = 1000;
    uint8 public constant DEFAULT_MAXIMUM_HOURLY_CHURN = 0;
    uint64 public constant DEFAULT_REGISTRATION_TIMESTAMP = 1000;
    uint256 public constant DEFAULT_STARTING_TOTAL_WEIGHT = 1e10;
    uint64 public constant DEFAULT_COMPLETION_TIMESTAMP = 100_000;

    ValidatorManager public validatorManager;

    // Used to create unique validator IDs in {_newNodeID}
    uint64 public nodeIDCounter = 0;

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

    function testInitializeEndValidation() public virtual {
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

    function testResendEndValidation() public virtual {
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
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        validatorManager.resendEndValidatorMessage(validationID);
    }

    function testCompleteEndValidation() public virtual {
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

        _mockGetPChainWarpMessage(subnetValidatorRegistrationMessage, true);

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

        _mockGetPChainWarpMessage(subnetValidatorRegistrationMessage, true);

        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodEnded(validationID, ValidatorStatus.Invalidated);

        validatorManager.completeEndValidation(0);
    }

    function testCummulativeChurnRegistration() public {
        uint64 churnThreshold =
            uint64(DEFAULT_STARTING_TOTAL_WEIGHT) * DEFAULT_MAXIMUM_CHURN_PERCENTAGE / 100;
        _beforeSend(_weightToValue(churnThreshold), address(this));

        // First registration should succeed
        _setUpCompleteValidatorRegistration({
            nodeID: _newNodeID(),
            subnetID: DEFAULT_SUBNET_ID,
            weight: churnThreshold,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _beforeSend(DEFAULT_MINIMUM_STAKE_AMOUNT, address(this)); // TODO may need to be updated with minimum stake amount

        // Second call should fail
        vm.expectRevert("ValidatorManager: maximum churn rate exceeded");
        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: DEFAULT_NODE_ID,
                registrationExpiry: DEFAULT_REGISTRATION_TIMESTAMP + 1,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
            }),
            _valueToWeight(DEFAULT_MINIMUM_STAKE_AMOUNT)
        );
    }

    function testCummulativeChurnRegistrationAndEndValidation() public {
        // Registration should succeed
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: _valueToWeight(DEFAULT_MINIMUM_STAKE_AMOUNT),
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        uint64 churnThreshold =
            uint64(DEFAULT_STARTING_TOTAL_WEIGHT) * DEFAULT_MAXIMUM_CHURN_PERCENTAGE / 100;
        _beforeSend(_weightToValue(churnThreshold), address(this));

        // Registration should succeed
        _setUpCompleteValidatorRegistration({
            nodeID: _newNodeID(),
            subnetID: DEFAULT_SUBNET_ID,
            weight: churnThreshold,
            registrationExpiry: DEFAULT_EXPIRY + 25 hours,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP + 25 hours
        });

        // Second call should fail
        vm.expectRevert("ValidatorManager: maximum churn rate exceeded");
        _initializeEndValidation(validationID, false);
    }

    function _oldSubnetConversionDataPack(SubnetConversionData memory scd)
        public
        pure
        returns (bytes memory)
    {
        bytes memory res = abi.encodePacked(
            scd.convertSubnetTxID,
            scd.validatorManagerBlockchainID,
            uint32(20),
            scd.validatorManagerAddress,
            uint32(scd.initialValidators.length)
        );
        for (uint256 i = 0; i < scd.initialValidators.length; i++) {
            res = abi.encodePacked(
                res,
                scd.initialValidators[i].nodeID,
                scd.initialValidators[i].weight,
                scd.initialValidators[i].blsPublicKey
            );
        }
        return res;
    }

    function _newNodeID() internal returns (bytes32) {
        nodeIDCounter++;
        return sha256(new bytes(nodeIDCounter));
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

        _beforeSend(_weightToValue(weight), address(this));
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidationPeriodCreated(validationID, nodeID, bytes32(0), weight, registrationExpiry);

        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: nodeID,
                registrationExpiry: registrationExpiry,
                blsPublicKey: blsPublicKey
            }),
            weight
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
            ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, true);

        _mockGetPChainWarpMessage(subnetValidatorRegistrationMessage, true);

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
        return _setUpInitializeEndValidation({
            nodeID: nodeID,
            subnetID: subnetID,
            weight: weight,
            registrationExpiry: registrationExpiry,
            blsPublicKey: blsPublicKey,
            registrationTimestamp: registrationTimestamp,
            completionTimestamp: completionTimestamp,
            includeUptime: false
        });
    }

    function _setUpInitializeEndValidation(
        bytes32 nodeID,
        bytes32 subnetID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory blsPublicKey,
        uint64 registrationTimestamp,
        uint64 completionTimestamp,
        bool includeUptime
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
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 1, 0);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        if (includeUptime) {
            bytes memory uptimeMsg = ValidatorMessages.packValidationUptimeMessage(
                validationID, completionTimestamp - registrationTimestamp
            );
            _mockGetUptimeWarpMessage(uptimeMsg, true);
            _mockGetBlockchainID();
        }
        vm.expectEmit(true, true, true, true, address(validatorManager));
        emit ValidatorRemovalInitialized(validationID, bytes32(0), weight, completionTimestamp);

        _initializeEndValidation(validationID, includeUptime);
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

    function _mockGetPChainWarpMessage(bytes memory expectedPayload, bool valid) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: validatorManager.P_CHAIN_BLOCKCHAIN_ID(),
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

    function _mockGetUptimeWarpMessage(bytes memory expectedPayload, bool valid) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
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
        _mockGetBlockchainID(DEFAULT_SOURCE_BLOCKCHAIN_ID);
    }

    function _mockGetBlockchainID(bytes32 blockchainID) internal {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(blockchainID)
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );
    }

    function _mockInitializeValidatorSet() internal {
        _mockGetPChainWarpMessage(
            ValidatorMessages.packSubnetConversionMessage(DEFAULT_SUBNET_CONVERSION_TX_ID), true
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual returns (bytes32);

    function _initializeEndValidation(bytes32 validationID, bool includeUptime) internal virtual;

    function _beforeSend(uint256 amount, address spender) internal virtual;

    function _defaultSubnetConversionData(uint8 numValidators)
        internal
        view
        returns (SubnetConversionData memory)
    {
        InitialValidator[] memory initialValidators = new InitialValidator[](numValidators);
        for (uint8 i = 0; i < numValidators; ++i) {
            bytes32 nodeID = bytes32(uint256(DEFAULT_INITIAL_VALIDATOR_NODE_ID) + i);
            bytes memory blsPublicKey = _generateBlsPublicKey(i);
            initialValidators[i] = InitialValidator({
                nodeID: nodeID,
                weight: DEFAULT_INITIAL_VALIDATOR_WEIGHT,
                blsPublicKey: blsPublicKey
            });
        }
        return SubnetConversionData({
            convertSubnetTxID: bytes32(0),
            validatorManagerBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            validatorManagerAddress: address(validatorManager),
            initialValidators: initialValidators
        });
    }

    // TODO this needs to be kept in line with the contract conversions, but we can't make external calls
    // to the contract and use vm.expectRevert at the same time.
    // These are okay to use for PoA as well, because they're just used for conversions inside the tests.
    function _valueToWeight(uint256 value) internal pure returns (uint64) {
        return uint64(value / 1e12);
    }

    // TODO this needs to be kept in line with the contract conversions, but we can't make external calls
    // to the contract and use vm.expectRevert at the same time.
    // These are okay to use for PoA as well, because they're just used for conversions inside the tests.
    function _weightToValue(uint64 weight) internal pure returns (uint256) {
        return uint256(weight) * 1e12;
    }

    // Function to generate a new BLS public key based on an index
    // It just sets the last byte to the value of the index
    function _generateBlsPublicKey(uint8 index) internal pure returns (bytes memory) {
        bytes memory res = DEFAULT_BLS_PUBLIC_KEY;
        res[47] = bytes1(index);
        return res;
    }
}
// solhint-enable no-empty-blocks
