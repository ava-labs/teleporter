// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

library StakingMessages {
    // The information that uniquely identifies a subnet validation period.
    // The SHA-256 hash of the concatenation of these field is the validationID.
    struct ValidationInfo {
        bytes32 subnetID;
        bytes32 nodeID;
        uint64 weight;
        uint64 registrationExpiry;
        bytes signature;
    }

    // Subnets send a RegisterSubnetValidator message to the P-Chain to register a validator.
    // The P-Chain responds with a RegisterSubnetValidator message indicating whether the registration was successful
    // for the given validation ID.
    uint32 internal constant SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID = 1;
    // Subnets can send a SetSubnetValidatorWeight message to the P-Chain to update a validator's weight.
    // The P-Chain responds with a SetSubnetValidatorWeight message acknowledging the weight update.
    uint32 internal constant SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID = 2;
    // The Subnet will self-sign a ValidationUptimeMessage to be provided when a validator is initiating
    // the end of their validation period.
    uint32 internal constant VALIDATION_UPTIME_MESSAGE_TYPE_ID = 3;

    // TODO: The implemenation of these packing and unpacking functions is neither tested or optimzied at all.
    // Full test coverage should be provided, and the implementation should be optimized for gas efficiency.

    /**
     * @notice Packs a RegisterSubnetValidatorMessage message into a byte array.
     * The message format specification is:
     * +-----------+----------+-----------+
     * |    typeID :   uint32 |   4 bytes |
     * +-----------+----------+-----------+
     * |  subnetID : [32]byte |  32 bytes |
     * +-----------+----------+-----------+
     * |    nodeID : [32]byte |  32 bytes |
     * +-----------+----------+-----------+
     * |    weight :   uint64 |   8 bytes |
     * +-----------+----------+-----------+
     * |    expiry :   uint64 |   8 bytes |
     * +-----------+----------+-----------+
     * | signature : [64]byte |  64 bytes |
     * +-----------+----------+-----------+
     *                        | 148 bytes |
     *                        +-----------+
     *
     * @param valiationInfo The information to pack into the message.
     * @return The validationID and the packed message.
     */
    function packRegisterSubnetValidatorMessage(ValidationInfo memory validationInfo)
        internal
        pure
        returns (bytes32, bytes memory)
    {
        (bytes32 validationID, bytes memory serializedValidationInfo) =
            packValidationInfo(validationInfo);

        bytes memory res = new bytes(148);
        // Pack the message type
        for (uint256 i; i < 4; ++i) {
            res[i] = bytes1(uint8(SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID >> (8 * (3 - i))));
        }
        // Pack the validation info
        for (uint256 i; i < 144; ++i) {
            res[i + 4] = serializedValidationInfo[i];
        }

        return (validationID, res);
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
        require(input.length == 148, "StakingMessages: Invalid message length");

        // Unpack the type ID
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i])) << uint32((8 * (3 - i)));
        }
        require(
            typeID == SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID,
            "StakingMessages: Invalid message type"
        );

        // Unpack the subnet ID
        bytes32 subnetID;
        for (uint256 i; i < 32; ++i) {
            subnetID |= bytes32(uint256(uint8(input[i + 4])) << (8 * (31 - i)));
        }

        // Unpack the node ID
        bytes32 nodeID;
        for (uint256 i; i < 32; ++i) {
            nodeID |= bytes32(uint256(uint8(input[i + 36])) << (8 * (31 - i)));
        }

        // Unpack the weight
        uint64 weight;
        for (uint256 i; i < 8; ++i) {
            weight |= uint64(uint8(input[i + 68])) << uint64((8 * (7 - i)));
        }

        // Unpack the expiry
        uint64 expiry;
        for (uint256 i; i < 8; ++i) {
            expiry |= uint64(uint8(input[i + 76])) << uint64((8 * (7 - i)));
        }

        // Unpack the signature
        bytes memory signature = new bytes(64);
        for (uint256 i; i < 64; ++i) {
            signature[i] = input[i + 84];
        }

        return ValidationInfo({
            subnetID: subnetID,
            nodeID: nodeID,
            weight: weight,
            registrationExpiry: expiry,
            signature: signature
        });
    }

    /**
     * @notice Packs a SubnetValidatorRegistrationMessage into a byte array.
     * The message format specification is:
     * +--------------+----------+----------+
     * |       typeID :   uint32 |  4 bytes |
     * +--------------+----------+----------+
     * | validationID : [32]byte | 32 bytes |
     * +--------------+----------+----------+
     * |        valid :     bool +  1 byte  |
     * +--------------+----------+----------+
     *                           | 37 bytes |
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
        bytes memory res = new bytes(37);
        // Pack the type ID.
        for (uint256 i; i < 4; ++i) {
            res[i] = bytes1(uint8(SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID >> (8 * (3 - i))));
        }
        // Pack the validation ID.
        for (uint256 i; i < 32; ++i) {
            res[i + 4] = bytes1(uint8(uint256(validationID >> (8 * (31 - i)))));
        }
        // Pack the validity.
        res[36] = bytes1(valid ? 1 : 0);
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
        require(input.length == 37, "StakingMessages: Invalid message length");

        // Unpack the type ID
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i])) << uint32((8 * (3 - i)));
        }
        require(
            typeID == SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID,
            "StakingMessages: Invalid message type"
        );

        // Unpack the validation ID.
        bytes32 validationID;
        for (uint256 i; i < 32; ++i) {
            validationID |= bytes32(uint256(uint8(input[i + 4])) << (8 * (31 - i)));
        }

        // Unpack the validity
        bool valid = input[36] != 0;

        return (validationID, valid);
    }

    /**
     * @notice Packs a SetSubnetValidatorWeightMessage message into a byte array.
     * The message format specification is:
     * +--------------+----------+----------+
     * |       typeID :   uint32 |  4 bytes |
     * +--------------+----------+----------+
     * | validationID : [32]byte | 32 bytes |
     * +--------------+----------+----------+
     * |        nonce :   uint64 |  8 bytes |
     * +--------------+----------+----------+
     * |       weight :   uint64 |  8 bytes |
     * +--------------+----------+----------+
     *                           | 52 bytes |
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
        bytes memory res = new bytes(52);
        // Pack the type ID.
        for (uint256 i; i < 4; ++i) {
            res[i] = bytes1(uint8(SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID >> (8 * (3 - i))));
        }
        // Pack the validation ID.
        for (uint256 i; i < 32; ++i) {
            res[i + 4] = bytes1(uint8(uint256(validationID >> (8 * (31 - i)))));
        }
        // Pack the nonce.
        for (uint256 i; i < 8; ++i) {
            res[i + 36] = bytes1(uint8(nonce >> (8 * (7 - i))));
        }
        // Pack the weight.
        for (uint256 i; i < 8; ++i) {
            res[i + 44] = bytes1(uint8(weight >> (8 * (7 - i))));
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
        require(input.length == 52, "StakingMessages: Invalid message length");

        // Unpack the type ID.
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i])) << uint32((8 * (3 - i)));
        }
        require(
            typeID == SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID,
            "StakingMessages: Invalid message type"
        );

        // Unpack the validation ID.
        bytes32 validationID;
        for (uint256 i; i < 32; ++i) {
            validationID |= bytes32(uint256(uint8(input[i + 4])) << (8 * (31 - i)));
        }

        // Unpack the nonce.
        uint64 nonce;
        for (uint256 i; i < 8; ++i) {
            nonce |= uint64(uint8(input[i + 36])) << uint64((8 * (7 - i)));
        }

        // Unpack the weight.
        uint64 weight;
        for (uint256 i; i < 8; ++i) {
            weight |= uint64(uint8(input[i + 44])) << uint64((8 * (7 - i)));
        }

        return (validationID, nonce, weight);
    }

    /**
     * @notice Packs a ValidationUptimeMessage into a byte array.
     * The message format specification is:
     * +--------------+----------+----------+
     * |       typeID :   uint32 |  4 bytes |
     * +--------------+----------+----------+
     * | validationID : [32]byte | 32 bytes |
     * +--------------+----------+----------+
     * |       uptime :   uint64 |  8 bytes |
     * +--------------+----------+----------+
     *                           | 44 bytes |
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
        bytes memory res = new bytes(44);
        // Pack the type ID.
        for (uint256 i; i < 4; ++i) {
            res[i] = bytes1(uint8(VALIDATION_UPTIME_MESSAGE_TYPE_ID >> (8 * (3 - i))));
        }
        // Pack the validation ID.
        for (uint256 i; i < 32; ++i) {
            res[i + 4] = bytes1(uint8(uint256(validationID >> (8 * (31 - i)))));
        }
        // Pack the uptime.
        for (uint256 i; i < 8; ++i) {
            res[i + 36] = bytes1(uint8(uptime >> (8 * (7 - i))));
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
        require(input.length == 44, "StakingMessages: Invalid message length");

        // Unpack the type ID.
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i])) << uint32((8 * (3 - i)));
        }
        require(
            typeID == VALIDATION_UPTIME_MESSAGE_TYPE_ID, "StakingMessages: Invalid message type"
        );

        // Unpack the validation ID.
        bytes32 validationID;
        for (uint256 i; i < 32; ++i) {
            validationID |= bytes32(uint256(uint8(input[i + 4])) << (8 * (31 - i)));
        }

        // Unpack the uptime.
        uint64 uptime;
        for (uint256 i; i < 8; ++i) {
            uptime |= uint64(uint8(input[i + 36])) << uint64((8 * (7 - i)));
        }

        return (validationID, uptime);
    }

    /**
     * @notice Packs all of the information pertaining to a validation period into a byte array.
     * The packed data is used to calculate the validationID as the SHA-256 hash of the packed data.
     * The specification of the packed data is:
     * +-----------+----------+-----------+
     * |  subnetID : [32]byte |  32 bytes |
     * +-----------+----------+-----------+
     * |    nodeID : [32]byte |  32 bytes |
     * +-----------+----------+-----------+
     * |    weight :   uint64 |   8 bytes |
     * +-----------+----------+-----------+
     * |    expiry :   uint64 |   8 bytes |
     * +-----------+----------+-----------+
     * | signature : [64]byte |  64 bytes |
     * +-----------+----------+-----------+
     *                        | 144 bytes |
     *                        +-----------+
     *
     * @param validationInfo The information to pack.
     * @return The validationID and the packed data.
     */
    function packValidationInfo(ValidationInfo memory validationInfo)
        internal
        pure
        returns (bytes32, bytes memory)
    {
        require(validationInfo.signature.length == 64, "StakingMessages: Invalid signature length");
        bytes memory res = new bytes(144);
        // Pack the subnetID
        for (uint256 i; i < 32; ++i) {
            res[i] = validationInfo.subnetID[i];
        }
        // Pack the nodeID
        for (uint256 i; i < 32; ++i) {
            res[i + 32] = validationInfo.nodeID[i];
        }
        // Pack the weight
        for (uint256 i; i < 8; ++i) {
            res[i + 64] = bytes1(uint8(validationInfo.weight >> uint8((8 * (7 - i)))));
        }
        // Pack the registration expiry
        for (uint256 i; i < 8; ++i) {
            res[i + 72] = bytes1(uint8(validationInfo.registrationExpiry >> uint64((8 * (7 - i)))));
        }
        // Pack the signature
        for (uint256 i; i < 64; ++i) {
            res[i + 80] = validationInfo.signature[i];
        }
        return (sha256(res), res);
    }

    /**
     * @notice Unpacks a byte array as a ValidationInfo.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return The unpacked ValidationInfo.
     */
    function unpackValidationInfo(bytes memory input)
        internal
        pure
        returns (ValidationInfo memory)
    {
        require(input.length == 144, "StakingMessages: Invalid message length");

        // Unpack the subnetID
        bytes32 subnetID;
        for (uint256 i; i < 32; ++i) {
            subnetID |= bytes32(uint256(uint8(input[i])) << (8 * (31 - i)));
        }

        // Unpack the nodeID
        bytes32 nodeID;
        for (uint256 i; i < 32; ++i) {
            nodeID |= bytes32(uint256(uint8(input[i + 32])) << (8 * (31 - i)));
        }

        // Unpack the weight
        uint64 weight;
        for (uint256 i; i < 8; ++i) {
            weight |= uint64(uint8(input[i + 64])) << uint64((8 * (7 - i)));
        }

        // Unpack the registration expiry
        uint64 expiry;
        for (uint256 i; i < 8; ++i) {
            expiry |= uint64(uint8(input[i + 72])) << uint64((8 * (7 - i)));
        }

        // Unpack the signature
        bytes memory signature = new bytes(64);
        for (uint256 i; i < 64; ++i) {
            signature[i] = input[i + 80];
        }

        return ValidationInfo({
            subnetID: subnetID,
            nodeID: nodeID,
            weight: weight,
            registrationExpiry: expiry,
            signature: signature
        });
    }
}
