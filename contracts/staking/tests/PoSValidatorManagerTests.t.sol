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
    address public constant DEFAULT_DELEGATOR_ADDRESS =
        address(0x1234123412341234123412341234123412341234);

    PoSValidatorManager public posValidatorManager;

    event ValidationUptimeUpdated(bytes32 indexed validationID, uint64 uptime);

    event DelegatorAdded(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        address indexed delegator,
        uint64 delegatorWeight,
        uint64 validatorWeight,
        uint64 nonce
    );

    event DelegatorRegistered(
        bytes32 indexed validationID,
        address indexed delegator,
        uint64 delegatorWeight,
        uint64 validatorWeight,
        uint64 nonce,
        uint256 startTime
    );

    event DelegatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        address indexed delegator,
        uint64 validatorWeight,
        uint64 nonce,
        uint256 endTime
    );

    event DelegationEnded(
        bytes32 indexed validationID,
        address indexed delegator,
        uint64 indexed validatorWeight,
        uint64 nonce
    );

    function testInitializeEndValidationWithUptimeProof() public {
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
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        _setUpInitializeDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP
        );
    }

    function testResendDelegatorRegistration() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        _setUpInitializeDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP
        );
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(
            validationID, 1, DEFAULT_WEIGHT + DEFAULT_DELEGATOR_WEIGHT
        );
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendDelegatorRegistration(validationID, DEFAULT_DELEGATOR_ADDRESS);
    }

    function testCompleteDelegatorRegistration() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP
        );
    }

    function testInitializeEndDelegation() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP
        );
        _setUpInitializeEndDelegation(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_WEIGHT,
            2,
            DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP
        );
    }

    function testResendEndDelegation() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP
        );
        _setUpInitializeEndDelegation(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_WEIGHT,
            2,
            DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP
        );
        bytes memory setValidatorWeightPayload =
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, 2, DEFAULT_WEIGHT);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        posValidatorManager.resendEndDelegation(validationID, DEFAULT_DELEGATOR_ADDRESS);
    }

    function testCompleteEndDelegation() public {
        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });
        _setUpCompleteDelegatorRegistration(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_DELEGATOR_WEIGHT,
            DEFAULT_DELEGATOR_WEIGHT + DEFAULT_WEIGHT,
            1,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP
        );
        _setUpInitializeEndDelegation(
            validationID,
            DEFAULT_DELEGATOR_ADDRESS,
            DEFAULT_WEIGHT,
            2,
            DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP
        );

        bytes memory weightUpdateMessage = ValidatorMessages.packSubnetValidatorWeightUpdateMessage(
            validationID, 2, DEFAULT_WEIGHT
        );
        _mockGetVerifiedWarpMessage(weightUpdateMessage, true);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegationEnded(validationID, DEFAULT_DELEGATOR_ADDRESS, DEFAULT_WEIGHT, 2);
        posValidatorManager.completeEndDelegation(0, DEFAULT_DELEGATOR_ADDRESS);
    }

    function testCompleteEndMultipleDelegations() public {}

    // Ensures that the behavior is the same regardless of delisting order
    function testCompleteEndMultipleDelegationsOutOfOrder() public {
        uint64 weight1 = 2e5;
        address delegator1 = address(0x1234123412341234123412341234123412341234);
        uint64 weight2 = 3e5;
        address delegator2 = address(0x5678567856785678567856785678567856785678);

        bytes32 validationID = _setUpCompleteValidatorRegistration({
            nodeID: DEFAULT_NODE_ID,
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_EXPIRY,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
        });

        // Register two delegators
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegator1,
            weight1,
            DEFAULT_WEIGHT + weight1,
            1,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP
        );
        _setUpCompleteDelegatorRegistration(
            validationID,
            delegator2,
            weight2,
            DEFAULT_WEIGHT + weight1 + weight2,
            2,
            DEFAULT_DELEGATOR_INIT_REGISTRATION_TIMESTAMP + 1,
            DEFAULT_DELEGATOR_COMPLETE_REGISTRATION_TIMESTAMP + 1
        );

        // Delist them in the opposite order,
        _setUpInitializeEndDelegation(
            validationID,
            delegator2,
            DEFAULT_WEIGHT + weight1,
            3,
            DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP
        );
        _setUpInitializeEndDelegation(
            validationID,
            delegator1,
            DEFAULT_WEIGHT,
            4,
            DEFAULT_DELEGATOR_END_DELEGATION_TIMESTAMP + 1
        );

        // Receive the weight update message for the second delegator first
        // This should include the weight updates for both delegators delisting
        bytes memory weightUpdateMessage = ValidatorMessages.packSubnetValidatorWeightUpdateMessage(
            validationID, 4, DEFAULT_WEIGHT
        );
        _mockGetVerifiedWarpMessage(weightUpdateMessage, true);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegationEnded(validationID, delegator2, DEFAULT_WEIGHT, 4);
        posValidatorManager.completeEndDelegation(0, delegator2);
        require(posValidatorManager.getWeight(validationID) == DEFAULT_WEIGHT, "PoSValidatorManagerTest: invalid weight");

        // Receive the weight update message for the first delegator.
        // The specified weight should still include the second delegator's weight,
        // but the validator weight will not be updated since we've already received a higher nonce
        weightUpdateMessage = ValidatorMessages.packSubnetValidatorWeightUpdateMessage(
            validationID, 3, DEFAULT_WEIGHT + weight2
        );
        _mockGetVerifiedWarpMessage(weightUpdateMessage, true);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegationEnded(validationID, delegator1, DEFAULT_WEIGHT + weight2, 3);
        posValidatorManager.completeEndDelegation(0, delegator1);
        require(posValidatorManager.getWeight(validationID) == DEFAULT_WEIGHT, "PoSValidatorManagerTest: invalid weight");
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
        bytes32 validationID,
        address delegator,
        uint64 weight,
        uint64 newValidatorWeight,
        uint64 nonce,
        uint64 registrationTimestamp
    ) internal returns (bytes32) {
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, nonce, newValidatorWeight);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));
        vm.warp(registrationTimestamp);

        _beforeSend(weight, delegator);

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorAdded({
            validationID: validationID,
            setWeightMessageID: bytes32(0),
            delegator: delegator,
            delegatorWeight: weight,
            validatorWeight: newValidatorWeight,
            nonce: nonce
        });

        _initializeDelegatorRegistration(validationID, delegator, weight);
        return validationID;
    }

    function _setUpCompleteDelegatorRegistration(
        bytes32 validationID,
        address delegator,
        uint64 weight,
        uint64 newValidatorWeight,
        uint64 nonce,
        uint64 initRegistrationTimestamp,
        uint64 completeRegistrationTimestamp
    ) internal returns (bytes32) {
        // TODONOW: decouple this
        _setUpInitializeDelegatorRegistration(
            validationID, delegator, weight, newValidatorWeight, nonce, initRegistrationTimestamp
        );
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSubnetValidatorWeightUpdateMessage(validationID, nonce, newValidatorWeight);
        _mockGetVerifiedWarpMessage(setValidatorWeightPayload, true);

        vm.warp(completeRegistrationTimestamp);
        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRegistered({
            validationID: validationID,
            delegator: delegator,
            delegatorWeight: weight,
            validatorWeight: newValidatorWeight,
            nonce: nonce,
            startTime: completeRegistrationTimestamp
        });
        posValidatorManager.completeDelegatorRegistration(0, delegator);
        return validationID;
    }

    function _setUpInitializeEndDelegation(
        bytes32 validationID,
        address delegator,
        uint64 newValidatorWeight,
        uint64 nonce,
        uint64 endDelegationTimestamp
    ) internal returns (bytes32) {
        require(nonce > 0, "PoSValidatorManagerTest: invalid nonce");
        // _setUpCompleteDelegatorRegistration(
        //     validationID, delegator, delegatorWeight, nonce-1, initRegistrationTimestamp, completeRegistrationTimestamp
        // );

        vm.warp(endDelegationTimestamp);
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, nonce, newValidatorWeight);
        _mockSendWarpMessage(setValidatorWeightPayload, bytes32(0));

        vm.expectEmit(true, true, true, true, address(posValidatorManager));
        emit DelegatorRemovalInitialized({
            validationID: validationID,
            setWeightMessageID: bytes32(0),
            delegator: delegator,
            validatorWeight: newValidatorWeight,
            nonce: nonce,
            endTime: endDelegationTimestamp
        });
        vm.prank(delegator);
        posValidatorManager.initializeEndDelegation(validationID);
        return validationID;
    }

    function _formatErrorMessage(bytes memory errorMessage) internal pure returns (bytes memory) {
        return abi.encodePacked("PoSValidatorManager: ", errorMessage);
    }
}
