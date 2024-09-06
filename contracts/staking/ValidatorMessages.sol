// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

import {Unpack} from "./Unpack.sol";

library ValidatorMessages {
    // The information that uniquely identifies a subnet validation period.
    // The validationID is the SHA-256 hash of the concatenation of the CODEC_ID,
    // REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID, and the concatenated ValidationPeriod fields.
    struct ValidationPeriod {
        bytes32 subnetID;
        bytes32 nodeID;
        uint64 weight;
        uint64 registrationExpiry;
        bytes blsPublicKey;
    }

    // The P-Chain uses a hardcoded codecID of 0 for all messages.
    uint16 internal constant CODEC_ID = 0;

    // Subnets send a RegisterSubnetValidator message to the P-Chain to register a validator.
    uint32 internal constant REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID = 0;

    // Subnets can send a SetSubnetValidatorWeight message to the P-Chain to update a validator's weight.
    // The P-Chain responds with a SetSubnetValidatorWeight message acknowledging the weight update.
    uint32 internal constant SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID = 1;

    // The P-Chain responds with a RegisterSubnetValidator message indicating whether the registration was successful
    // for the given validation ID.
    uint32 internal constant SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID = 2;

    // The P-Chain responds with a SubnetValidatorWeightUpdateMessage message indicating whether the weight update was successful
    // for the given validation ID.
    uint32 internal constant SUBNET_VALIDATOR_WEIGHT_UPDATE_MESSAGE_TYPE_ID = 3;

    // The Subnet will self-sign a ValidationUptimeMessage to be provided when a validator is initiating
    // the end of their validation period.
    uint32 internal constant VALIDATION_UPTIME_MESSAGE_TYPE_ID = 4;

    // TODO: The implementation of these packing and unpacking functions is neither tested nor optimized at all.
    // Full test coverage should be provided, and the implementation should be optimized for gas efficiency.

    /**
     * @notice Packs a RegisterSubnetValidatorMessage message into a byte array.
     * The message format specification is:
     * +--------------+----------+-----------+
     * |      codecID :   uint16 |   2 bytes |
     * +--------------+----------+-----------+
     * |       typeID :   uint32 |   4 bytes |
     * +--------------+----------+-----------+
     * |     subnetID : [32]byte |  32 bytes |
     * +--------------+----------+-----------+
     * |       nodeID : [32]byte |  32 bytes |
     * +--------------+----------+-----------+
     * |       weight :   uint64 |   8 bytes |
     * +--------------+----------+-----------+
     * | blsPublicKey : [48]byte |  48 bytes |
     * +--------------+----------+-----------+
     * |       expiry :   uint64 |   8 bytes |
     * +--------------+----------+-----------+
     *                           | 134 bytes |
     *                           +-----------+
     *
     * @param validationPeriod The information to pack into the message.
     * @return The validationID and the packed message.
     */
    function packRegisterSubnetValidatorMessage(ValidationPeriod memory validationPeriod)
        internal
        pure
        returns (bytes32, bytes memory)
    {
        require(
            validationPeriod.blsPublicKey.length == 48, "StakingMessages: invalid signature length"
        );
        bytes memory packed = abi.encodePacked(
            CODEC_ID,
            REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID,
            validationPeriod.subnetID,
            validationPeriod.nodeID,
            validationPeriod.weight,
            validationPeriod.blsPublicKey,
            validationPeriod.registrationExpiry
        );
        return (sha256(packed), packed);
    }

    /**
     * @notice Unpacks a byte array as a RegisterSubnetValidatorMessage message.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return the unpacked ValidationPeriod.
     */
    function unpackRegisterSubnetValidatorMessage(bytes memory input)
        internal
        pure
        returns (ValidationPeriod memory)
    {
        require(input.length == 134, "ValidatorMessages: invalid message length");

        (
            bytes2 codecID,
            bytes4 typeID,
            bytes32 subnetID,
            bytes32 nodeID,
            bytes8 weight,
            bytes memory blsPublicKey,
            bytes8 expiry
        ) = Unpack.unpack_2_4_32_32_8_Dyn_8_Destructive(input);

        require(uint16(codecID) == CODEC_ID, "ValidatorMessages: invalid codec ID");
        require(
            uint32(typeID) == REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID,
            "ValidatorMessages: invalid message type"
        );

        return ValidationPeriod({
            subnetID: subnetID,
            nodeID: nodeID,
            weight: uint64(weight),
            registrationExpiry: uint64(expiry),
            blsPublicKey: blsPublicKey
        });
    }

    /**
     * @notice Packs a SubnetValidatorRegistrationMessage into a byte array.
     * The message format specification is:
     * +--------------+----------+----------+
     * |      codecID :   uint16 |  2 bytes |
     * +--------------+----------+----------+
     * |       typeID :   uint32 |  4 bytes |
     * +--------------+----------+----------+
     * | validationID : [32]byte | 32 bytes |
     * +--------------+----------+----------+
     * |        valid :     bool |  1 byte  |
     * +--------------+----------+----------+
     *                           | 39 bytes |
     *                           +----------+
     *
     * @param validationID The ID of the validation period.
     * @param valid true if the validation period was registered, false if it was not and never will be.
     * @return The packed message.
     *
     */
    function packSubnetValidatorRegistrationMessage(
        bytes32 validationID,
        bool valid
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            CODEC_ID, SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID, validationID, valid
        );
    }

    /**
     * @notice Unpacks a byte array as a SubnetValidatorRegistrationMessage message.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return The validationID and whether or the validation period was registered
     * or is not a validator and never will be a validator to do the expiry time passing.
     */
    function unpackSubnetValidatorRegistrationMessage(bytes memory input)
        internal
        pure
        returns (bytes32, bool)
    {
        require(input.length == 39, "ValidatorMessages: invalid message length");

        (bytes2 codecID, bytes4 typeID, bytes32 validationID, bytes1 valid) =
            Unpack.unpack_2_4_32_1(input);

        require(uint16(codecID) == CODEC_ID, "ValidatorMessages: invalid codec ID");
        require(
            uint32(typeID) == SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID,
            "ValidatorMessages: invalid message type"
        );

        return (validationID, uint8(valid) != 0);
    }

    /**
     * @notice Packs a SetSubnetValidatorWeightMessage message into a byte array.
     * The message format specification is:
     * +--------------+----------+----------+
     * |      codecID :   uint16 |  2 bytes |
     * +--------------+----------+----------+
     * |       typeID :   uint32 |  4 bytes |
     * +--------------+----------+----------+
     * | validationID : [32]byte | 32 bytes |
     * +--------------+----------+----------+
     * |        nonce :   uint64 |  8 bytes |
     * +--------------+----------+----------+
     * |       weight :   uint64 |  8 bytes |
     * +--------------+----------+----------+
     *                           | 54 bytes |
     *                           +----------+
     *
     * @param validationID The ID of the validation period.
     * @param nonce The nonce of the validation ID.
     * @param weight The new weight of the validator.
     * @return The packed message.
     */
    function packSetSubnetValidatorWeightMessage(
        bytes32 validationID,
        uint64 nonce,
        uint64 weight
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            CODEC_ID, SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID, validationID, nonce, weight
        );
    }

    /**
     * @notice Unpacks a byte array as a SetSubnetValidatorWeight message.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return The validationID, nonce, and weight.
     */
    function unpackSetSubnetValidatorWeightMessage(bytes memory input)
        internal
        pure
        returns (bytes32, uint64, uint64)
    {
        require(input.length == 54, "ValidatorMessages: invalid message length");

        (bytes2 codecID, bytes4 typeID, bytes32 validationID, bytes8 nonce, bytes8 weight) =
            Unpack.unpack_2_4_32_8_8(input);

        require(uint16(codecID) == CODEC_ID, "ValidatorMessages: invalid codec ID");
        require(
            uint32(typeID) == SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID,
            "ValidatorMessages: invalid message type"
        );

        return (validationID, uint64(nonce), uint64(weight));
    }

    /**
     * @notice Packs a SubnetValidatorWeightUpdateMessage into a byte array.
     * The message format specification is:
     * +--------------+----------+----------+
     * |      codecID :   uint16 |  2 bytes |
     * +--------------+----------+----------+
     * |       typeID :   uint32 |  4 bytes |
     * +--------------+----------+----------+
     * | validationID : [32]byte | 32 bytes |
     * +--------------+----------+----------+
     * |        nonce :   uint64 |  8 bytes |
     * +--------------+----------+----------+
     * |       weight :   uint64 |  8 bytes |
     * +--------------+----------+----------+
     *                           | 54 bytes |
     *                           +----------+
     */
    function packSubnetValidatorWeightUpdateMessage(
        bytes32 validationID,
        uint64 nonce,
        uint64 weight
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            CODEC_ID, SUBNET_VALIDATOR_WEIGHT_UPDATE_MESSAGE_TYPE_ID, validationID, nonce, weight
        );
    }

    /**
     * @notice Unpacks a byte array as a SubnetValidatorWeightUpdateMessag.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return The validationID, weight, and nonce.
     */
    function unpackSubnetValidatorWeightUpdateMessage(bytes memory input)
        internal
        pure
        returns (bytes32, uint64, uint64)
    {
        require(input.length == 54, "ValidatorMessages: invalid message length");

        (bytes2 codecID, bytes4 typeID, bytes32 validationID, bytes8 nonce, bytes8 weight) =
            Unpack.unpack_2_4_32_8_8(input);

        require(uint16(codecID) == CODEC_ID, "ValidatorMessages: invalid codec ID");
        require(
            uint32(typeID) == SUBNET_VALIDATOR_WEIGHT_UPDATE_MESSAGE_TYPE_ID,
            "ValidatorMessages: invalid message type"
        );

        return (validationID, uint64(nonce), uint64(weight));
    }

    /**
     * @notice Packs a ValidationUptimeMessage into a byte array.
     * The message format specification is:
     * +--------------+----------+----------+
     * |      codecID :   uint16 |  2 bytes |
     * +--------------+----------+----------+
     * |       typeID :   uint32 |  4 bytes |
     * +--------------+----------+----------+
     * | validationID : [32]byte | 32 bytes |
     * +--------------+----------+----------+
     * |       uptime :   uint64 |  8 bytes |
     * +--------------+----------+----------+
     *                           | 46 bytes |
     *                           +----------+
     *
     * @param validationID The ID of the validation period.
     * @param uptime The uptime of the validator.
     * @return The packed message.
     */
    function packValidationUptimeMessage(
        bytes32 validationID,
        uint64 uptime
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(CODEC_ID, VALIDATION_UPTIME_MESSAGE_TYPE_ID, validationID, uptime);
    }

    /**
     * @notice Unpacks a byte array as a ValidationUptimeMessage.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return The validationID and uptime.
     */
    function unpackValidationUptimeMessage(bytes memory input)
        internal
        pure
        returns (bytes32, uint64)
    {
        require(input.length == 46, "ValidatorMessages: invalid message length");

        (bytes2 codecID, bytes4 typeID, bytes32 validationID, bytes8 uptime) =
            Unpack.unpack_2_4_32_8(input);

        require(uint16(codecID) == CODEC_ID, "ValidatorMessages: invalid codec ID");
        require(
            uint32(typeID) == VALIDATION_UPTIME_MESSAGE_TYPE_ID,
            "ValidatorMessages: invalid message type"
        );

        return (validationID, uint64(uptime));
    }
}
