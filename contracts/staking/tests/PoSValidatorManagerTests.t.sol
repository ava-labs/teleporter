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

    PoSValidatorManager public posValidatorManager;

    // Used to create unique validator IDs in {registerValidators}
    uint64 public validatorCounter = 0;

    event ValidationUptimeUpdated(bytes32 indexed validationID, uint64 uptime);

    function registerValidators(uint64 n) public {
        for (uint64 i = validatorCounter; validatorCounter <= i + n; validatorCounter++) {
            _setUpCompleteValidatorRegistration({
                nodeID: sha256(new bytes(validatorCounter)),
                subnetID: DEFAULT_SUBNET_ID,
                weight: DEFAULT_WEIGHT,
                registrationExpiry: DEFAULT_EXPIRY,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                registrationTimestamp: DEFAULT_REGISTRATION_TIMESTAMP
            });
        }
    }

    function testInvalidChurnRegistration() public {
        // First registration should work
        _setUpCompleteValidatorRegistration({
            nodeID: sha256(new bytes(1)),
            subnetID: DEFAULT_SUBNET_ID,
            weight: DEFAULT_WEIGHT,
            registrationExpiry: DEFAULT_CHURN_TRACKER_START_TIME + 1 days,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
            registrationTimestamp: DEFAULT_CHURN_TRACKER_START_TIME - 1
        });

        vm.warp(DEFAULT_CHURN_TRACKER_START_TIME + 1);

        bytes32 nodeID = sha256(new bytes(2));
        (, bytes memory registerSubnetValidatorMessage) = ValidatorMessages
            .packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                nodeID: nodeID,
                subnetID: DEFAULT_SUBNET_ID,
                weight: DEFAULT_WEIGHT,
                registrationExpiry: DEFAULT_CHURN_TRACKER_START_TIME + 1 days,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
            })
        );
        _beforeSend(DEFAULT_WEIGHT);

        uint256 value = posValidatorManager.weightToValue(DEFAULT_WEIGHT);
        vm.expectRevert(_formatErrorMessage("maximum hourly churn rate exceeded"));
        _initializeValidatorRegistrationWithValue(nodeID, DEFAULT_CHURN_TRACKER_START_TIME + 1 days, DEFAULT_BLS_PUBLIC_KEY, value);
    }

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

    function _formatErrorMessage(bytes memory errorMessage) internal pure returns (bytes memory) {
        return abi.encodePacked("PoSValidatorManager: ", errorMessage);
    }

    function _initializeValidatorRegistrationWithValue(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory blsPublicKey,
        uint256 value
    ) internal virtual returns (bytes32);
}
