// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

import {Unpack} from "./Unpack.sol";

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
     * @notice Packs a RegisterSubnetValidator message into a byte array.
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
     */
    function packRegisterSubnetValidatorMessage(ValidationInfo memory validationInfo)
        internal
        pure
        returns (bytes32, bytes memory)
    {
        (bytes32 validationID, bytes memory serialized) = packAndHashValidationInfo(validationInfo);
        return (
            validationID,
            abi.encodePacked(SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID, serialized)
        );
    }

    /**
     * @notice Unpacks a byte array as a RegisterSubnetValidator message.
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
     * @return The validationID and whether or the validation period was registered
     * or is not a validator and never will be a validator to do the expiry time passing.
     */
    function unpackSubnetValidatorRegistrationMessage(bytes memory input)
        internal
        pure
        returns (bytes32, bool)
    {
        (bytes4 typeID, bytes32 validationID, bytes1 valid) = Unpack.unpack_4_32_1(input);
        require(
            uint32(typeID) == SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID,
            "StakingMessages: Invalid message type"
        );
        return (validationID, uint8(valid) != 0);
    }

    /**
     * @notice Packs a SetSubnetValidatorWeight message into a byte array.
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
     */
    function packSetSubnetValidatorWeightMessage(
        bytes32 validationID,
        uint64 nonce,
        uint64 weight
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID, validationID, nonce, weight
        );
    }

    /**
     * @notice Unpacks a byte array as a SetSubnetValidatorWeight message.
     * The message format specification is the same as the one used in above for packing.
     */
    function unpackSetSubnetValidatorWeightMessage(bytes memory input)
        internal
        pure
        returns (bytes32, uint64, uint64)
    {
        (bytes4 typeID, bytes32 validationID, bytes8 nonce, bytes8 weight) =
            Unpack.unpack_4_32_8_8(input);
        require(
            uint32(typeID) == SET_SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID,
            "StakingMessages: Invalid message type"
        );
        return (validationID, uint64(nonce), uint64(weight));
    }

    /**
     * @notice Unpacks a byte array as a ValidationUptimeMessage.
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
     */
    function unpackValidationUptimeMessage(bytes memory input)
        internal
        pure
        returns (bytes32, uint64)
    {
        (bytes4 typeID, bytes32 validationID, bytes8 uptime) = Unpack.unpack_4_32_8(input);
        require(
            uint32(typeID) == VALIDATION_UPTIME_MESSAGE_TYPE_ID,
            "StakingMessages: Invalid message type"
        );
        return (validationID, uint64(uptime));
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
     */
    function packValidationInfo(ValidationInfo memory info) internal pure returns (bytes memory) {
        require(info.signature.length == 64, "StakingMessages: Invalid signature length");
        return abi.encodePacked(
            info.subnetID, info.nodeID, info.weight, info.registrationExpiry, info.signature
        );
    }

    /// @dev Equivalent to returning `packValidationInfo(info)` and the SHA256 hash thereof.
    function packAndHashValidationInfo(ValidationInfo memory info)
        internal
        pure
        returns (bytes32, bytes memory)
    {
        bytes memory res = packValidationInfo(info);
        return (sha256(res), res);
    }
}
