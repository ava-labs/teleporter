// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {Test} from "@forge-std/Test.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {PChainOwner, ConversionData, InitialValidator} from "../interfaces/IValidatorManager.sol";

contract ValidatorMessagesTest is Test {
    bytes32 public constant DEFAULT_L1_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_NODE_ID =
        bytes(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes32 public constant DEFAULT_SUBNET_CONVERSION_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    bytes public constant DEFAULT_BLS_PUBLIC_KEY = bytes(
        hex"123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
    );
    bytes32 public constant DEFAULT_VALIDATION_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    uint64 public constant DEFAULT_WEIGHT = 1e6;
    uint64 public constant DEFAULT_EXPIRY = 1000;
    // solhint-disable-next-line var-name-mixedcase
    PChainOwner public DEFAULT_P_CHAIN_OWNER;
    address public constant DEFAULT_OWNER = 0x1234567812345678123456781234567812345678;

    function setUp() public {
        address[] memory addresses = new address[](1);
        addresses[0] = DEFAULT_OWNER;
        DEFAULT_P_CHAIN_OWNER = PChainOwner({threshold: 1, addresses: addresses});
    }

    function testSubnetConversionMessageInvalidInputLength() public {
        bytes memory packed =
            ValidatorMessages.packSubnetToL1ConversionMessage(DEFAULT_SUBNET_CONVERSION_ID);
        // Invalid length
        bytes memory invalidPacked = new bytes(packed.length - 1);
        for (uint256 i = 0; i < packed.length - 1; i++) {
            invalidPacked[i] = packed[i];
        }
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidMessageLength.selector, 37, 38)
        );
        ValidatorMessages.unpackSubnetToL1ConversionMessage(invalidPacked);
    }

    function testSubnetConversionMessageInvalidCodecID() public {
        bytes memory packed =
            ValidatorMessages.packSubnetToL1ConversionMessage(DEFAULT_SUBNET_CONVERSION_ID);

        // Invalid codec ID
        bytes memory invalidPacked2 = packed;
        invalidPacked2[1] = 0x01;
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidCodecID.selector, uint32(1))
        );
        ValidatorMessages.unpackSubnetToL1ConversionMessage(invalidPacked2);
    }

    function testSubnetConversionMessageInvalidTypeID() public {
        bytes memory packed =
            ValidatorMessages.packSubnetToL1ConversionMessage(DEFAULT_SUBNET_CONVERSION_ID);
        // Invalid message type
        bytes memory invalidPacked3 = packed;
        invalidPacked3[5] = 0x01;
        vm.expectRevert(ValidatorMessages.InvalidMessageType.selector);
        ValidatorMessages.unpackSubnetToL1ConversionMessage(invalidPacked3);
    }

    function testRegisterSubnetValidatorMessageInvalidBLSKey() public {
        vm.expectRevert(ValidatorMessages.InvalidBLSPublicKey.selector);
        // 47 bytes
        bytes memory invalidBLSKey = bytes(
            hex"3456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
        );
        ValidatorMessages.packRegisterL1ValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                l1ID: DEFAULT_L1_ID,
                nodeID: DEFAULT_NODE_ID,
                registrationExpiry: DEFAULT_EXPIRY,
                blsPublicKey: invalidBLSKey,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER,
                weight: DEFAULT_WEIGHT
            })
        );
    }

    function testRegisterSubnetValidatorMessageInvalidInputLength() public {
        bytes memory packed = _getPackedRegisterL1ValidatorMessage();
        // Invalid length
        bytes memory invalidPacked = new bytes(packed.length - 1);
        for (uint256 i = 0; i < packed.length - 1; i++) {
            invalidPacked[i] = packed[i];
        }
        vm.expectRevert(
            abi.encodeWithSelector(
                ValidatorMessages.InvalidMessageLength.selector, uint32(193), uint32(194)
            )
        );
        ValidatorMessages.unpackRegisterL1ValidatorMessage(invalidPacked);
    }

    function testRegisterSubnetValidatorMessageInvalidCodecID() public {
        bytes memory packed = _getPackedRegisterL1ValidatorMessage();

        // Invalid codec ID
        bytes memory invalidPacked2 = packed;
        invalidPacked2[1] = 0x01;
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidCodecID.selector, uint32(1))
        );
        ValidatorMessages.unpackRegisterL1ValidatorMessage(invalidPacked2);
    }

    function testRegisterSubnetValidatorMessageInvalidTypeID() public {
        bytes memory packed = _getPackedRegisterL1ValidatorMessage();

        // Invalid message type
        bytes memory invalidPacked3 = packed;
        invalidPacked3[5] = 0x00;
        vm.expectRevert(ValidatorMessages.InvalidMessageType.selector);
        ValidatorMessages.unpackRegisterL1ValidatorMessage(invalidPacked3);
    }

    function testSubnetValidatorRegistrationMessageInvalidInputLength() public {
        bytes memory packed =
            ValidatorMessages.packL1ValidatorRegistrationMessage(DEFAULT_VALIDATION_ID, true);

        // Invalid length
        bytes memory invalidPacked = new bytes(packed.length - 1);
        for (uint256 i = 0; i < packed.length - 1; i++) {
            invalidPacked[i] = packed[i];
        }
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidMessageLength.selector, 38, 39)
        );
        ValidatorMessages.unpackL1ValidatorRegistrationMessage(invalidPacked);
    }

    function testSubnetValidatorRegistrationMessageInvalidCodecID() public {
        bytes memory packed =
            ValidatorMessages.packL1ValidatorRegistrationMessage(DEFAULT_VALIDATION_ID, true);

        // Invalid codec ID
        bytes memory invalidPacked2 = packed;
        invalidPacked2[1] = 0x01;
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidCodecID.selector, uint32(1))
        );
        ValidatorMessages.unpackL1ValidatorRegistrationMessage(invalidPacked2);
    }

    function testSubnetValidatorRegistrationMessageInvalidTypeID() public {
        bytes memory packed =
            ValidatorMessages.packL1ValidatorRegistrationMessage(DEFAULT_VALIDATION_ID, true);

        // Invalid message type
        bytes memory invalidPacked3 = packed;
        invalidPacked3[5] = 0x01;
        vm.expectRevert(ValidatorMessages.InvalidMessageType.selector);
        ValidatorMessages.unpackL1ValidatorRegistrationMessage(invalidPacked3);
    }

    function testValidationUptimeMessageInvalidInputLength() public {
        bytes memory packed =
            ValidatorMessages.packValidationUptimeMessage(DEFAULT_VALIDATION_ID, 100);

        // Invalid length
        bytes memory invalidPacked = new bytes(packed.length - 1);
        for (uint256 i = 0; i < packed.length - 1; i++) {
            invalidPacked[i] = packed[i];
        }
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidMessageLength.selector, 45, 46)
        );
        ValidatorMessages.unpackValidationUptimeMessage(invalidPacked);
    }

    function testValidationUptimeMessageInvalidCodecID() public {
        bytes memory packed =
            ValidatorMessages.packValidationUptimeMessage(DEFAULT_VALIDATION_ID, 100);

        // Invalid codec ID
        bytes memory invalidPacked2 = packed;
        invalidPacked2[1] = 0x01;
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidCodecID.selector, uint32(1))
        );
        ValidatorMessages.unpackValidationUptimeMessage(invalidPacked2);
    }

    function testValidationUptimeMessageInvalidTypeID() public {
        bytes memory packed =
            ValidatorMessages.packValidationUptimeMessage(DEFAULT_VALIDATION_ID, 100);

        // Invalid message type
        bytes memory invalidPacked3 = packed;
        invalidPacked3[5] = 0x01;
        vm.expectRevert(ValidatorMessages.InvalidMessageType.selector);
        ValidatorMessages.unpackValidationUptimeMessage(invalidPacked3);
    }

    function testSetSubnetValidatorWeightMessageInvalidInputLength() public {
        bytes memory packed = ValidatorMessages.packL1ValidatorWeightMessage(
            DEFAULT_VALIDATION_ID, 100, DEFAULT_WEIGHT
        );

        // Invalid length
        bytes memory invalidPacked = new bytes(packed.length - 1);
        for (uint256 i = 0; i < packed.length - 1; i++) {
            invalidPacked[i] = packed[i];
        }
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidMessageLength.selector, 53, 54)
        );
        ValidatorMessages.unpackL1ValidatorWeightMessage(invalidPacked);
    }

    function testSetSubnetValidatorWeightMessageInvalidCodecID() public {
        bytes memory packed = ValidatorMessages.packL1ValidatorWeightMessage(
            DEFAULT_VALIDATION_ID, 100, DEFAULT_WEIGHT
        );

        // Invalid codec ID
        bytes memory invalidPacked2 = packed;
        invalidPacked2[1] = 0x01;
        vm.expectRevert(
            abi.encodeWithSelector(ValidatorMessages.InvalidCodecID.selector, uint32(1))
        );
        ValidatorMessages.unpackL1ValidatorWeightMessage(invalidPacked2);
    }

    function testSetSubnetValidatorWeightMessageInvalidTypeID() public {
        bytes memory packed = ValidatorMessages.packL1ValidatorWeightMessage(
            DEFAULT_VALIDATION_ID, 100, DEFAULT_WEIGHT
        );

        // Invalid message type
        bytes memory invalidPacked3 = packed;
        invalidPacked3[5] = 0x01;
        vm.expectRevert(ValidatorMessages.InvalidMessageType.selector);
        ValidatorMessages.unpackL1ValidatorWeightMessage(invalidPacked3);
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

    function testSubnetConversionMessage() public pure {
        bytes memory packed =
            ValidatorMessages.packSubnetToL1ConversionMessage(DEFAULT_SUBNET_CONVERSION_ID);
        bytes32 conversionID = ValidatorMessages.unpackSubnetToL1ConversionMessage(packed);
        assertEq(conversionID, DEFAULT_SUBNET_CONVERSION_ID);
    }

    function testPackL1ConversionData() public pure {
        InitialValidator[] memory initialValidators = new InitialValidator[](1);
        initialValidators[0] = InitialValidator({
            nodeID: DEFAULT_NODE_ID,
            weight: DEFAULT_WEIGHT,
            blsPublicKey: DEFAULT_BLS_PUBLIC_KEY
        });
        bytes memory packed = ValidatorMessages.packConversionData(
            ConversionData({
                l1ID: DEFAULT_L1_ID,
                validatorManagerBlockchainID: DEFAULT_SUBNET_CONVERSION_ID,
                validatorManagerAddress: DEFAULT_OWNER,
                initialValidators: initialValidators
            })
        );

        assertEq(packed.length, 186);
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

    function testSubnetValidatorWeightUpdateMessage() public pure {
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

    function _getPackedRegisterL1ValidatorMessage() internal returns (bytes memory) {
        (, bytes memory packed) = ValidatorMessages.packRegisterL1ValidatorMessage(
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
        return packed;
    }
}
