// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {PChainOwner} from "../interfaces/IValidatorManager.sol";

contract ValidatorMessagesTest is Test {
    bytes32 public constant DEFAULT_L1_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_NODE_ID =
        bytes(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_BLS_PUBLIC_KEY = bytes(
        hex"123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    bytes32 public constant DEFAULT_VALIDATION_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    uint64 public constant DEFAULT_WEIGHT = 1e6;
    uint64 public constant DEFAULT_EXPIRY = 1000;
    // solhint-disable-next-line var-name-mixedcase
    PChainOwner public DEFAULT_P_CHAIN_OWNER;

    function setUp() public {
        address[] memory addresses = new address[](1);
        addresses[0] = 0x1234567812345678123456781234567812345678;
        DEFAULT_P_CHAIN_OWNER = PChainOwner({threshold: 1, addresses: addresses});
    }

    function testRegisterSubnetValidatorMessage() public view {
        (bytes32 validationID, bytes memory packed) = ValidatorMessages
            .packRegisterL1ValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                l1ID: DEFAULT_L1_ID,
                nodeID: DEFAULT_NODE_ID,
                registrationExpiry: DEFAULT_EXPIRY,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                weight: DEFAULT_WEIGHT
            })
        );

        ValidatorMessages.ValidationPeriod memory info =
            ValidatorMessages.unpackRegisterL1ValidatorMessage(packed);
        assertEq(info.l1ID, DEFAULT_L1_ID);
        assertEq(info.nodeID, DEFAULT_NODE_ID);
        assertEq(info.weight, DEFAULT_WEIGHT);
        assertEq(info.registrationExpiry, DEFAULT_EXPIRY);
        assertEq(info.blsPublicKey, DEFAULT_BLS_PUBLIC_KEY);

        (bytes32 recoveredID,) = ValidatorMessages.packRegisterL1ValidatorMessage(info);
        assertEq(recoveredID, validationID);
    }

    function testSubnetValidatorRegistrationMessage() public pure {
        bytes memory packed =
            ValidatorMessages.packL1ValidatorRegistrationMessage(DEFAULT_VALIDATION_ID, true);
        (bytes32 validationID, bool valid) =
            ValidatorMessages.unpackL1ValidatorRegistrationMessage(packed);
        assertEq(validationID, DEFAULT_VALIDATION_ID);
        assert(valid);
    }

    function testSetSubnetValidatorWeightMessage() public pure {
        bytes memory packed = ValidatorMessages.packL1ValidatorWeightMessage(
            DEFAULT_VALIDATION_ID, 100, DEFAULT_WEIGHT
        );
        (bytes32 validationID, uint64 nonce, uint64 weight) =
            ValidatorMessages.unpackL1ValidatorWeightMessage(packed);
        assertEq(validationID, DEFAULT_VALIDATION_ID);
        assertEq(nonce, 100);
        assertEq(weight, DEFAULT_WEIGHT);
    }

    function testSubnetValidatorWeightUpdateMessag() public pure {
        bytes memory packed = ValidatorMessages.packL1ValidatorWeightMessage(
            DEFAULT_VALIDATION_ID, 100, DEFAULT_WEIGHT
        );
        (bytes32 validationID, uint64 nonce, uint64 weight) =
            ValidatorMessages.unpackL1ValidatorWeightMessage(packed);
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
}
