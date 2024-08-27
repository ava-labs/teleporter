// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {PoSValidatorManager} from "../PoSValidatorManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";

abstract contract PoSValidatorManagerTest is ValidatorManagerTest {
    uint64 public constant DEFAULT_UPTIME = uint64(100);
    uint64 public constant DEFAULT_DELEGATOR_WEIGHT = uint64(1e5);
    uint64 public constant DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP = uint64(2000);
    uint64 public constant DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP = uint64(3000);
    uint64 public constant DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP = uint64(4000);
    address public constant DEFAULT_DELEGATOR_ADDRESS = address(0x1234123412341234123412341234123412341234);

    PoSValidatorManager public posValidatorManager;

    event ValidationUptimeUpdated(bytes32 indexed validationID, uint64 uptime);

    event DelegatorAdded(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        address indexed delegator,
        uint256 weight,
        uint256 startTime
    );

    event DelegatorRegistered(
        bytes32 indexed validationID,
        address indexed delegator,
        uint256 weight,
        uint256 startTime
    );

    event DelegatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        address indexed delegator,
        uint256 endTime
    );

    event DelegationEnded(
        bytes32 indexed validationID,
        address indexed delegator
    );

    function testInitializeEndValidationWithUptimeProof() public {
        // TODO: implement
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _mockGetBlockchainID();
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getVerifiedWarpMessage.selector, uint32(0)),
            abi.encode(
                WarpMessage({
                    sourceChainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    originSenderAddress: address(0),
                    payload: ValidatorMessages.packValidationUptimeMessage(validationID, DEFAULT_UPTIME)
                }),
                true
            )
        );
        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeCall(IWarpMessenger.getVerifiedWarpMessage, 0)
        );

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit ValidationUptimeUpdated(validationID, DEFAULT_UPTIME);
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeWarpMessage() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _mockGetVerifiedWarpMessage(new bytes(0), false);
        vm.expectRevert(_formatErrorMessage("invalid warp message"));
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeChainID() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        _mockGetVerifiedWarpMessage(new bytes(0), true);
        _mockGetBlockchainID();
        vm.expectRevert(_formatErrorMessage("invalid source chain ID"));
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeSenderAddress() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

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

        vm.expectRevert(_formatErrorMessage("invalid origin sender address"));
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInvalidUptimeValidationID() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

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

        vm.expectRevert(_formatErrorMessage("invalid uptime validation ID"));
        posValidatorManager.initializeEndValidation(validationID, true, 0);
    }

    function testInitializeDelegatorRegistration() public {
        _setUpInitializeDelegatorRegistration(DEFAULT_DELEGATOR_ADDRESS, DEFAULT_DELEGATOR_WEIGHT, DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP);
    }

    function testResendDelegatorRegistration() public {
        bytes32 validationID = _setUpInitializeDelegatorRegistration(DEFAULT_DELEGATOR_ADDRESS, DEFAULT_DELEGATOR_WEIGHT, DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP);
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(
                validationID,
                1,
                DEFAULT_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
            );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendDelegatorRegistration(validationID, DEFAULT_DELEGATOR_ADDRESS);
    }

    function testCompleteDelegatorRegistration() public {
        _setUpCompleteDelegatorRegistration(DEFAULT_DELEGATOR_ADDRESS, DEFAULT_DELEGATOR_WEIGHT, DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP, DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP);
    }

    function testInitializeEndDelegation() public {
        _setUpInitializeEndDelegation(
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP
        );
    }

    function testResendEndDelegation() public {
        bytes32 validationID = _setUpInitializeEndDelegation(
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP
        );
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 2, DEFAULT_WEIGHT);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendEndDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);
    }

    function testCompleteEndDelegation() public {
        bytes32 validationID = _setUpInitializeEndDelegation(
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP
        );

        bytes memory weightUpdateMessage =
            ValidatorMessages.packSubnetValidatorWeightUpdateMessage(
                validationID,
                2,
                DEFAULT_WEIGHT
            );
        _mockGetVerifiedWarpMessage(weightUpdateMessage, true);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegationEnded(validationID, DEFAULT_DELEGATOR_ADDRESS);
        posValidatorManager.completeEndDelegation(0, DEFAULT_DELEGATOR_ADDRESS);
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

    function _initializeEndValidation(bytes32 validationID) internal virtual override {
        return posValidatorManager.initializeEndValidation(validationID, false, 0);
    }

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegator,
        uint64 weight
    ) internal virtual;

    function _setUpInitializeDelegatorRegistration(
        address delegator,
        uint64 weight,
        uint64 registrationTimestamp
    ) internal returns (bytes32) {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(
                validationID,
                1,
                DEFAULT_WEIGHT + weight
            );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.warp(registrationTimestamp);

        _beforeSend(weight, delegator);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorAdded(validationID, bytes32(0), delegator, weight, registrationTimestamp);

        _initializeDelegatorRegistration(validationID, delegator, weight);
        return validationID;
    }

    function _setUpCompleteDelegatorRegistration(
        address delegator,
        uint64 weight,
        uint64 initRegistrationTimestamp,
        uint64 completeRegistrationTimestamp
    ) internal returns (bytes32) {
        bytes32 validationID = _setUpInitializeDelegatorRegistration(delegator, weight, initRegistrationTimestamp);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSubnetValidatorWeightUpdateMessage(
                validationID,
                1,
                DEFAULT_WEIGHT + weight
            );
        _mockGetVerifiedWarpMessage(setValidatorWeightPayload, true);

        vm.warp(completeRegistrationTimestamp);
        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRegistered(validationID, delegator, weight, completeRegistrationTimestamp);
        posValidatorManager.completeDelegatorRegistration(0, delegator);
        return validationID;
    }

    function _setUpInitializeEndDelegation(
        address delegator,
        uint64 weight,
        uint64 initRegistrationTimestamp,
        uint64 completeRegistrationTimestamp,
        uint64 endDelegationTimestamp
    ) internal returns (bytes32) {
        bytes32 validationID = _setUpCompleteDelegatorRegistration(
            delegator,
            weight,
            initRegistrationTimestamp,
            completeRegistrationTimestamp
        );

        vm.warp(endDelegationTimestamp);
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 2, DEFAULT_WEIGHT);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRemovalInitialized(validationID, bytes32(0), delegator, endDelegationTimestamp);
        vm.prank(delegator);
        posValidatorManager.initializeEndDelegation(validationID);
        return validationID;
    }

    function _formatErrorMessage(bytes memory errorMessage) internal pure returns (bytes memory) {
        return abi.encodePacked("PoSValidatorManager: ", errorMessage);
    }
}
