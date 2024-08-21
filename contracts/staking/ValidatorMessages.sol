// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

library ValidatorMessages {
    // The information that uniquely identifies a subnet validation period.
    // The validationID is the SHA-256 hash of the concatenation of the CODEC_ID,
    // REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID, and the concatenated ValidationInfo fields.
    struct ValidationInfo {
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

    // The P-Chain responds with a SetSubnetValidatorWeight message indicating whether the weight update was successful
    // for the given validation ID.
    uint32 internal constant SET_SUBNET_VALIDATOR_WEIGHT_UPDATE_MESSAGE_TYPE_ID = 3;

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
     * @param validationInfo The information to pack into the message.
     * @return The validationID and the packed message.
     */
    function packRegisterSubnetValidatorMessage(ValidationInfo memory validationInfo)
        internal
        pure
        returns (bytes32, bytes memory)
    {
        require(
            validationInfo.blsPublicKey.length == 48, "StakingMessages: Invalid signature length"
        );
        bytes memory res = new bytes(134);
        // Pack the codec ID
        for (uint256 i; i < 2; ++i) {
            res[i] = bytes1(uint8(CODEC_ID >> uint8((8 * (1 - i)))));
        }
        // Pack the type ID
        for (uint256 i; i < 4; ++i) {
            res[i + 2] =
                bytes1(uint8(REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID >> uint8((8 * (3 - i)))));
        }

        // Pack the subnetID
        for (uint256 i; i < 32; ++i) {
            res[i + 6] = validationInfo.subnetID[i];
        }
        // Pack the nodeID
        for (uint256 i; i < 32; ++i) {
            res[i + 38] = validationInfo.nodeID[i];
        }
        // Pack the weight
        for (uint256 i; i < 8; ++i) {
            res[i + 70] = bytes1(uint8(validationInfo.weight >> uint8((8 * (7 - i)))));
        }
        // Pack the blsPublicKey
        for (uint256 i; i < 48; ++i) {
            res[i + 78] = validationInfo.blsPublicKey[i];
        }
        // Pack the registration expiry
        for (uint256 i; i < 8; ++i) {
            res[i + 126] = bytes1(uint8(validationInfo.registrationExpiry >> uint64((8 * (7 - i)))));
        }
        return (sha256(res), res);
    }

    /**
     * @notice Unpacks a byte array as a RegisterSubnetValidatorMessage message.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return the unpacked ValidationInfo.
     */
    function unpackRegisterSubnetValidatorMessage(bytes memory input)
        internal
        pure
        returns (ValidationInfo memory)
    {
        require(input.length == 134, "ValidatorMessages: Invalid message length");

        // Unpack the codec ID
        uint16 codecID;
        for (uint256 i; i < 2; ++i) {
            codecID |= uint16(uint8(input[i])) << uint16((8 * (1 - i)));
        }
        require(codecID == CODEC_ID, "ValidatorMessages: Invalid codec ID");

        // Unpack the type ID
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i + 2])) << uint32((8 * (3 - i)));
        }
        require(
            typeID == REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID,
            "ValidatorMessages: Invalid message type"
        );

        // Unpack the subnetID
        bytes32 subnetID;
        for (uint256 i; i < 32; ++i) {
            subnetID |= bytes32(uint256(uint8(input[i + 6])) << (8 * (31 - i)));
        }

        // Unpack the nodeID
        bytes32 nodeID;
        for (uint256 i; i < 32; ++i) {
            nodeID |= bytes32(uint256(uint8(input[i + 38])) << (8 * (31 - i)));
        }

        // Unpack the weight
        uint64 weight;
        for (uint256 i; i < 8; ++i) {
            weight |= uint64(uint8(input[i + 70])) << uint64((8 * (7 - i)));
        }

        // Unpack the blsPublicKey
        bytes memory blsPublicKey = new bytes(48);
        for (uint256 i; i < 48; ++i) {
            blsPublicKey[i] = input[i + 78];
        }

        // Unpack the registration expiry
        uint64 expiry;
        for (uint256 i; i < 8; ++i) {
            expiry |= uint64(uint8(input[i + 126])) << uint64((8 * (7 - i)));
        }

        return ValidationInfo({
            subnetID: subnetID,
            nodeID: nodeID,
            weight: weight,
            registrationExpiry: expiry,
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
     * |        valid :     bool +  1 byte  |
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
        bytes memory res = new bytes(39);
        // Pack the codec ID.
        for (uint256 i; i < 2; ++i) {
            res[i] = bytes1(uint8(CODEC_ID >> (8 * (1 - i))));
        }
        // Pack the type ID.
        for (uint256 i; i < 4; ++i) {
            res[i + 2] =
                bytes1(uint8(SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID >> (8 * (3 - i))));
        }
        // Pack the validation ID.
        for (uint256 i; i < 32; ++i) {
            res[i + 6] = bytes1(uint8(uint256(validationID >> (8 * (31 - i)))));
        }
        // Pack the validity.
        res[38] = bytes1(valid ? 1 : 0);
        return res;
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
        require(input.length == 39, "ValidatorMessages: Invalid message length");
        // Unpack the codec ID
        uint16 codecID;
        for (uint256 i; i < 2; ++i) {
            codecID |= uint16(uint8(input[i])) << uint16((8 * (1 - i)));
        }
        require(codecID == CODEC_ID, "ValidatorMessages: Invalid codec ID");

        // Unpack the type ID
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i + 2])) << uint32((8 * (3 - i)));
        }
        require(
            typeID == SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID,
            "ValidatorMessages: Invalid message type"
        );

        // Unpack the validation ID.
        bytes32 validationID;
        for (uint256 i; i < 32; ++i) {
            validationID |= bytes32(uint256(uint8(input[i + 6])) << (8 * (31 - i)));
        }

        // Unpack the validity
        bool valid = input[38] != 0;

        return (validationID, valid);
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
        bytes memory res = new bytes(54);
        // Pack the codec ID.
        for (uint256 i; i < 2; ++i) {
            res[i] = bytes1(uint8(CODEC_ID >> (8 * (1 - i))));
        }
        // Pack the type ID.
        for (uint256 i; i < 4; ++i) {
            res[i + 2] = bytes1(uint8(SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID >> (8 * (3 - i))));
        }
        // Pack the validation ID.
        for (uint256 i; i < 32; ++i) {
            res[i + 6] = bytes1(uint8(uint256(validationID >> (8 * (31 - i)))));
        }
        // Pack the nonce.
        for (uint256 i; i < 8; ++i) {
            res[i + 38] = bytes1(uint8(nonce >> (8 * (7 - i))));
        }
        // Pack the weight.
        for (uint256 i; i < 8; ++i) {
            res[i + 46] = bytes1(uint8(weight >> (8 * (7 - i))));
        }
        return res;
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
        require(input.length == 54, "ValidatorMessages: Invalid message length");

        // Unpack the codec ID.
        uint16 codecID;
        for (uint256 i; i < 2; ++i) {
            codecID |= uint16(uint8(input[i])) << uint16((8 * (1 - i)));
        }
        require(codecID == CODEC_ID, "ValidatorMessages: Invalid codec ID");

        // Unpack the type ID.
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i + 2])) << uint32((8 * (3 - i)));
        }
        require(
            typeID == SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID,
            "ValidatorMessages: Invalid message type"
        );

        // Unpack the validation ID.
        bytes32 validationID;
        for (uint256 i; i < 32; ++i) {
            validationID |= bytes32(uint256(uint8(input[i + 6])) << (8 * (31 - i)));
        }

        // Unpack the nonce.
        uint64 nonce;
        for (uint256 i; i < 8; ++i) {
            nonce |= uint64(uint8(input[i + 38])) << uint64((8 * (7 - i)));
        }

        // Unpack the weight.
        uint64 weight;
        for (uint256 i; i < 8; ++i) {
            weight |= uint64(uint8(input[i + 46])) << uint64((8 * (7 - i)));
        }

        return (validationID, nonce, weight);
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
        bytes memory res = new bytes(46);

        // Pack the codec ID.
        for (uint256 i; i < 2; ++i) {
            res[i] = bytes1(uint8(CODEC_ID >> (8 * (1 - i))));
        }
        // Pack the type ID.
        for (uint256 i; i < 4; ++i) {
            res[i + 2] = bytes1(uint8(VALIDATION_UPTIME_MESSAGE_TYPE_ID >> (8 * (3 - i))));
        }
        // Pack the validation ID.
        for (uint256 i; i < 32; ++i) {
            res[i + 6] = bytes1(uint8(uint256(validationID >> (8 * (31 - i)))));
        }
        // Pack the uptime.
        for (uint256 i; i < 8; ++i) {
            res[i + 38] = bytes1(uint8(uptime >> (8 * (7 - i))));
        }
        return res;
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
        require(input.length == 46, "ValidatorMessages: Invalid message length");

        // Unpack the codec ID.
        uint16 codecID;
        for (uint256 i; i < 2; ++i) {
            codecID |= uint16(uint8(input[i])) << uint16((8 * (1 - i)));
        }
        require(codecID == CODEC_ID, "ValidatorMessages: Invalid codec ID");

        // Unpack the type ID.
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i + 2])) << uint32((8 * (3 - i)));
        }
        require(
            typeID == VALIDATION_UPTIME_MESSAGE_TYPE_ID, "ValidatorMessages: Invalid message type"
        );

        // Unpack the validation ID.
        bytes32 validationID;
        for (uint256 i; i < 32; ++i) {
            validationID |= bytes32(uint256(uint8(input[i + 6])) << (8 * (31 - i)));
        }

        // Unpack the uptime.
        uint64 uptime;
        for (uint256 i; i < 8; ++i) {
            uptime |= uint64(uint8(input[i + 38])) << uint64((8 * (7 - i)));
        }

        return (validationID, uptime);
    }
}
