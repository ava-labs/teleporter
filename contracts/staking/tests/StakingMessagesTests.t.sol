// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";

contract StakingMessagesTest is Test {
    bytes32 public constant DEFAULT_SUBNET_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant DEFAULT_NODE_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_ED25519_SIGNATURE = bytes(
        hex"12345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    bytes32 public constant DEFAULT_VALIDATION_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    uint64 public constant DEFAULT_WEIGHT = 1e6;
    uint64 public constant DEFAULT_EXPIRY = 1000;

    function testRegisterSubnetValidatorMessage() public pure {
        (bytes32 validationID, bytes memory packed) = ValidatorMessages
            .packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationInfo({
                subnetID: DEFAULT_SUBNET_ID,
                nodeID: DEFAULT_NODE_ID,
                weight: DEFAULT_WEIGHT,
                registrationExpiry: DEFAULT_EXPIRY,
                signature: DEFAULT_ED25519_SIGNATURE
            })
        );

        StakingMessages.ValidationInfo memory info =
            ValidatorMessages.unpackRegisterSubnetValidatorMessage(packed);
        assertEq(info.subnetID, DEFAULT_SUBNET_ID);
        assertEq(info.nodeID, DEFAULT_NODE_ID);
        assertEq(info.weight, DEFAULT_WEIGHT);
        assertEq(info.registrationExpiry, DEFAULT_EXPIRY);
        assertEq(info.signature, DEFAULT_ED25519_SIGNATURE);

        (bytes32 recoveredID,) = ValidatorMessages.packValidationInfo(info);
        assertEq(recoveredID, validationID);
    }

    function testSubnetValidatorRegistrationMessage() public pure {
        bytes memory packed =
            ValidatorMessages.packSubnetValidatorRegistrationMessage(DEFAULT_VALIDATION_ID, true);
        (bytes32 validationID, bool valid) =
            ValidatorMessages.unpackSubnetValidatorRegistrationMessage(packed);
        assertEq(validationID, DEFAULT_VALIDATION_ID);
        assert(valid);
    }

    function testSetSubnetValidatorWeightMessage() public pure {
        bytes memory packed = ValidatorMessages.packSetSubnetValidatorWeightMessage(
            DEFAULT_VALIDATION_ID, 100, DEFAULT_WEIGHT
        );
        (bytes32 validationID, uint64 nonce, uint64 weight) =
            ValidatorMessages.unpackSetSubnetValidatorWeightMessage(packed);
        assertEq(validationID, DEFAULT_VALIDATION_ID);
        assertEq(nonce, 100);
        assertEq(weight, DEFAULT_WEIGHT);
    }

    function testValidationUptimeMessage() public pure {
        bytes memory packed =
            ValidatorMessages.packValidationUptimeMessage(DEFAULT_VALIDATION_ID, 100);
        (bytes32 validationID, uint64 uptime) =
            ValidatorMessages.unpackValidationUptimeMessage(packed);
        assertEq(validationID, DEFAULT_VALIDATION_ID);
        assertEq(uptime, 100);
    }

    function testValidationInfo() public pure {
        (, bytes memory packed) = ValidatorMessages.packValidationInfo(
            ValidatorMessages.ValidationInfo({
                subnetID: DEFAULT_SUBNET_ID,
                nodeID: DEFAULT_NODE_ID,
                weight: DEFAULT_WEIGHT,
                registrationExpiry: DEFAULT_EXPIRY,
                signature: DEFAULT_ED25519_SIGNATURE
            })
        );
        StakingMessages.ValidationInfo memory info = ValidatorMessages.unpackValidationInfo(packed);
        assertEq(info.subnetID, DEFAULT_SUBNET_ID);
        assertEq(info.nodeID, DEFAULT_NODE_ID);
        assertEq(info.weight, DEFAULT_WEIGHT);
        assertEq(info.registrationExpiry, DEFAULT_EXPIRY);
        assertEq(info.signature, DEFAULT_ED25519_SIGNATURE);
    }
}
