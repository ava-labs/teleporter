// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

import {PChainOwner, SubnetConversionData} from "./interfaces/IValidatorManager.sol";

/**
 * @dev Packing utilities for the Warp message types used by the Validator Manager contracts, as specified in ACP-77:
 * https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets
 */
library ValidatorMessages {
    // The information that uniquely identifies a subnet validation period.
    // The validationID is the SHA-256 hash of the concatenation of the CODEC_ID,
    // REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID, and the concatenated ValidationPeriod fields.
    struct ValidationPeriod {
        bytes32 subnetID;
        bytes nodeID;
        bytes blsPublicKey;
        uint64 registrationExpiry;
        PChainOwner remainingBalanceOwner;
        PChainOwner disableOwner;
        uint64 weight;
    }

    // The P-Chain uses a hardcoded codecID of 0 for all messages.
    uint16 internal constant CODEC_ID = 0;

    // The P-Chain signs a SubnetConversion message that is used to verify the Subnet's initial validators.
    uint32 internal constant SUBNET_CONVERSION_MESSAGE_TYPE_ID = 0;

    // Subnets send a RegisterSubnetValidator message to the P-Chain to register a validator.
    uint32 internal constant REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID = 1;

    // The P-Chain responds with a RegisterSubnetValidator message indicating whether the registration was successful
    // for the given validation ID.
    uint32 internal constant SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID = 2;

    // Subnets can send a SubnetValidatorWeight message to the P-Chain to update a validator's weight.
    // The P-Chain responds with another SubnetValidatorWeight message acknowledging the weight update.
    uint32 internal constant SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID = 3;

    // The Subnet will self-sign a ValidationUptimeMessage to be provided when a validator is initiating
    // the end of their validation period.
    uint32 internal constant VALIDATION_UPTIME_MESSAGE_TYPE_ID = 0;

    error InvalidMessageLength(uint32 actual, uint32 expected);
    error InvalidCodecID(uint32 id);
    error InvalidMessageType();
    error InvalidBLSPublicKey();

    /**
     * @notice Packs a SubnetConversionMessage message into a byte array.
     * The message format specification is:
     * +--------------------+----------+----------+
     * |            codecID :   uint16 |  2 bytes |
     * +--------------------+----------+----------+
     * |             typeID :   uint32 |  4 bytes |
     * +--------------------+----------+----------+
     * | subnetConversionID : [32]byte | 32 bytes |
     * +--------------------+----------+----------+
     *                                 | 38 bytes |
     *                                 +----------+
     *
     * @param subnetConversionID The subnet conversion ID to pack into the message.
     * @return The packed message.
     */
    function packSubnetConversionMessage(bytes32 subnetConversionID)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encodePacked(CODEC_ID, SUBNET_CONVERSION_MESSAGE_TYPE_ID, subnetConversionID);
    }

    /**
     * @notice Unpacks a byte array as a SubnetConversionMessage message.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return The unpacked subnetConversionID.
     */
    function unpackSubnetConversionMessage(bytes memory input) internal pure returns (bytes32) {
        if (input.length != 38) {
            revert InvalidMessageLength(uint32(input.length), 38);
        }

        // Unpack the codec ID
        uint16 codecID;
        for (uint256 i; i < 2; ++i) {
            codecID |= uint16(uint8(input[i])) << uint16((8 * (1 - i)));
        }
        if (codecID != CODEC_ID) {
            revert InvalidCodecID(codecID);
        }

        // Unpack the type ID
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i + 2])) << uint32((8 * (3 - i)));
        }
        if (typeID != SUBNET_CONVERSION_MESSAGE_TYPE_ID) {
            revert InvalidMessageType();
        }

        // Unpack the subnetConversionID
        bytes32 subnetConversionID;
        for (uint256 i; i < 32; ++i) {
            subnetConversionID |= bytes32(uint256(uint8(input[i + 6])) << (8 * (31 - i)));
        }

        return subnetConversionID;
    }

    /**
     * @notice Packs SubnetConversionData into a byte array.
     * This byte array is the SHA256 pre-image of the subnetConversionID hash
     * The message format specification is:
     *
     * SubnetConversionData:
     * +----------------+-----------------+--------------------------------------------------------+
     * |       codecID  :          uint16 |                                                2 bytes |
     * +----------------+-----------------+--------------------------------------------------------+
     * |       subnetID :        [32]byte |                                               32 bytes |
     * +----------------+-----------------+--------------------------------------------------------+
     * | managerChainID :        [32]byte |                                               32 bytes |
     * +----------------+-----------------+--------------------------------------------------------+
     * | managerAddress :          []byte |                          4 + len(managerAddress) bytes |
     * +----------------+-----------------+--------------------------------------------------------+
     * |     validators : []ValidatorData |                        4 + sum(validatorLengths) bytes |
     * +----------------+-----------------+--------------------------------------------------------+
     *                                    | 74 + len(managerAddress) + len(validatorLengths) bytes |
     *                                    +--------------------------------------------------------+
     * ValidatorData:
     * +--------------+----------+------------------------+
     * |       nodeID :   []byte |  4 + len(nodeID) bytes |
     * +--------------+----------+------------------------+
     * | blsPublicKey : [48]byte |               48 bytes |
     * +--------------+----------+------------------------+
     * |       weight :   uint64 |                8 bytes |
     * +--------------+----------+------------------------+
     *                           | 60 + len(nodeID) bytes |
     *                           +------------------------+
     *
     * @dev Input validation is skipped, since the returned value is intended to be compared
     * directly with an authenticated Warp message.
     * @param subnetConversionData The struct representing data to pack into the message.
     * @return The packed message.
     */
    function packSubnetConversionData(SubnetConversionData memory subnetConversionData)
        internal
        pure
        returns (bytes memory)
    {
        // Hardcoded 20 is for length of the managerAddress on EVM chains
        // solhint-disable-next-line func-named-parameters
        bytes memory res = abi.encodePacked(
            CODEC_ID,
            subnetConversionData.subnetID,
            subnetConversionData.validatorManagerBlockchainID,
            uint32(20),
            subnetConversionData.validatorManagerAddress,
            uint32(subnetConversionData.initialValidators.length)
        );
        // The approach below of encoding initialValidators using `abi.encodePacked` in a loop
        // was tested against pre-allocating the array and doing manual byte by byte packing and
        // it was found to be more gas efficient.
        for (uint256 i; i < subnetConversionData.initialValidators.length; ++i) {
            res = abi.encodePacked(
                res,
                uint32(subnetConversionData.initialValidators[i].nodeID.length),
                subnetConversionData.initialValidators[i].nodeID,
                subnetConversionData.initialValidators[i].blsPublicKey,
                subnetConversionData.initialValidators[i].weight
            );
        }
        return res;
    }

    /**
     * @notice Packs a RegisterSubnetValidatorMessage message into a byte array.
     * The message format specification is:
     *
     * RegisterSubnetValidatorMessage:
     * +-----------------------+-------------+--------------------------------------------------------------------+
     * |               codecID :      uint16 |                                                            2 bytes |
     * +-----------------------+-------------+--------------------------------------------------------------------+
     * |                typeID :      uint32 |                                                            4 bytes |
     * +-----------------------+-------------+-------------------------------------------------------------------+
     * |              subnetID :    [32]byte |                                                           32 bytes |
     * +-----------------------+-------------+--------------------------------------------------------------------+
     * |                nodeID :      []byte |                                              4 + len(nodeID) bytes |
     * +-----------------------+-------------+--------------------------------------------------------------------+
     * |          blsPublicKey :    [48]byte |                                                           48 bytes |
     * +-----------------------+-------------+--------------------------------------------------------------------+
     * |                expiry :      uint64 |                                                            8 bytes |
     * +-----------------------+-------------+--------------------------------------------------------------------+
     * | remainingBalanceOwner : PChainOwner |                                      8 + len(addresses) * 20 bytes |
     * +-----------------------+-------------+--------------------------------------------------------------------+
     * |          disableOwner : PChainOwner |                                      8 + len(addresses) * 20 bytes |
     * +-----------------------+-------------+--------------------------------------------------------------------+
     * |                weight :      uint64 |                                                            8 bytes |
     * +-----------------------+-------------+--------------------------------------------------------------------+
     *                                       | 122 + len(nodeID) + (len(addresses1) + len(addresses2)) * 20 bytes |
     *                                       +--------------------------------------------------------------------+
     *
     * PChainOwner:
     * +-----------+------------+-------------------------------+
     * | threshold :     uint32 |                       4 bytes |
     * +-----------+------------+-------------------------------+
     * | addresses : [][20]byte | 4 + len(addresses) * 20 bytes |
     * +-----------+------------+-------------------------------+
     *                          | 8 + len(addresses) * 20 bytes |
     *                          +-------------------------------+
     *
     * @param validationPeriod The information to pack into the message.
     * @return The validationID and the packed message.
     */
    function packRegisterSubnetValidatorMessage(ValidationPeriod memory validationPeriod)
        internal
        pure
        returns (bytes32, bytes memory)
    {
        if (validationPeriod.blsPublicKey.length != 48) {
            revert InvalidBLSPublicKey();
        }

        // solhint-disable-next-line func-named-parameters
        bytes memory res = abi.encodePacked(
            CODEC_ID,
            REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID,
            validationPeriod.subnetID,
            uint32(validationPeriod.nodeID.length),
            validationPeriod.nodeID,
            validationPeriod.blsPublicKey,
            validationPeriod.registrationExpiry,
            validationPeriod.remainingBalanceOwner.threshold,
            uint32(validationPeriod.remainingBalanceOwner.addresses.length)
        );
        for (uint256 i; i < validationPeriod.remainingBalanceOwner.addresses.length; ++i) {
            res = abi.encodePacked(res, validationPeriod.remainingBalanceOwner.addresses[i]);
        }
        res = abi.encodePacked(
            res,
            validationPeriod.disableOwner.threshold,
            uint32(validationPeriod.disableOwner.addresses.length)
        );
        for (uint256 i; i < validationPeriod.disableOwner.addresses.length; ++i) {
            res = abi.encodePacked(res, validationPeriod.disableOwner.addresses[i]);
        }
        res = abi.encodePacked(res, validationPeriod.weight);

        return (sha256(res), res);
    }

    /**
     * @notice Unpacks a byte array as a RegisterSubnetValidatorMessage message.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return The unpacked ValidationPeriod.
     */
    function unpackRegisterSubnetValidatorMessage(bytes memory input)
        internal
        pure
        returns (ValidationPeriod memory)
    {
        uint32 index = 0;
        ValidationPeriod memory validation;

        // Unpack the codec ID
        // Individual fields are unpacked in their own scopes to avoid stack too deep errors.
        {
            uint16 codecID;
            for (uint256 i; i < 2; ++i) {
                codecID |= uint16(uint8(input[i + index])) << uint16((8 * (1 - i)));
            }
            if (codecID != CODEC_ID) {
                revert InvalidCodecID(codecID);
            }
            index += 2;
        }

        // Unpack the type ID
        {
            uint32 typeID;
            for (uint256 i; i < 4; ++i) {
                typeID |= uint32(uint8(input[i + index])) << uint32((8 * (3 - i)));
            }
            if (typeID != REGISTER_SUBNET_VALIDATOR_MESSAGE_TYPE_ID) {
                revert InvalidMessageType();
            }
            index += 4;
        }

        // Unpack the subnetID
        {
            bytes32 subnetID;
            for (uint256 i; i < 32; ++i) {
                subnetID |= bytes32(uint256(uint8(input[i + index])) << (8 * (31 - i)));
            }
            validation.subnetID = subnetID;
            index += 32;
        }

        // Unpack the nodeID length
        uint32 nodeIDLength;
        {
            for (uint256 i; i < 4; ++i) {
                nodeIDLength |= uint32(uint8(input[i + index])) << uint32((8 * (3 - i)));
            }
            index += 4;

            // Unpack the nodeID
            bytes memory nodeID = new bytes(nodeIDLength);
            for (uint256 i; i < nodeIDLength; ++i) {
                nodeID[i] = input[i + index];
            }
            validation.nodeID = nodeID;
            index += nodeIDLength;
        }

        // Unpack the blsPublicKey
        {
            bytes memory blsPublicKey = new bytes(48);
            for (uint256 i; i < 48; ++i) {
                blsPublicKey[i] = input[i + index];
            }
            validation.blsPublicKey = blsPublicKey;
            index += 48;
        }

        // Unpack the registration expiry
        {
            uint64 expiry;
            for (uint256 i; i < 8; ++i) {
                expiry |= uint64(uint8(input[i + index])) << uint64((8 * (7 - i)));
            }
            validation.registrationExpiry = expiry;
            index += 8;
        }

        // Unpack the remainingBalanceOwner threshold
        uint32 remainingBalanceOwnerAddressesLength;
        {
            uint32 remainingBalanceOwnerThreshold;
            for (uint256 i; i < 4; ++i) {
                remainingBalanceOwnerThreshold |=
                    uint32(uint8(input[i + index])) << uint32((8 * (3 - i)));
            }
            index += 4;

            // Unpack the remainingBalanceOwner addresses length
            for (uint256 i; i < 4; ++i) {
                remainingBalanceOwnerAddressesLength |=
                    uint32(uint8(input[i + index])) << uint32((8 * (3 - i)));
            }
            index += 4;

            // Unpack the remainingBalanceOwner addresses
            address[] memory remainingBalanceOwnerAddresses =
                new address[](remainingBalanceOwnerAddressesLength);
            for (uint256 i; i < remainingBalanceOwnerAddressesLength; ++i) {
                bytes memory addrBytes = new bytes(20);
                for (uint256 j; j < 20; ++j) {
                    addrBytes[j] = input[j + index];
                }
                address addr;
                // solhint-disable-next-line no-inline-assembly
                assembly {
                    addr := mload(add(addrBytes, 20))
                }
                remainingBalanceOwnerAddresses[i] = addr;
                index += 20;
            }
            validation.remainingBalanceOwner = PChainOwner({
                threshold: remainingBalanceOwnerThreshold,
                addresses: remainingBalanceOwnerAddresses
            });
        }

        // Unpack the disableOwner threshold
        uint32 disableOwnerAddressesLength;
        {
            uint32 disableOwnerThreshold;
            for (uint256 i; i < 4; ++i) {
                disableOwnerThreshold |= uint32(uint8(input[i + index])) << uint32((8 * (3 - i)));
            }
            index += 4;

            // Unpack the disableOwner addresses length
            for (uint256 i; i < 4; ++i) {
                disableOwnerAddressesLength |=
                    uint32(uint8(input[i + index])) << uint32((8 * (3 - i)));
            }
            index += 4;

            // Unpack the disableOwner addresses
            address[] memory disableOwnerAddresses = new address[](disableOwnerAddressesLength);
            for (uint256 i; i < disableOwnerAddressesLength; ++i) {
                bytes memory addrBytes = new bytes(20);
                for (uint256 j; j < 20; ++j) {
                    addrBytes[j] = input[j + index];
                }
                address addr;
                // solhint-disable-next-line no-inline-assembly
                assembly {
                    addr := mload(add(addrBytes, 20))
                }
                disableOwnerAddresses[i] = addr;
                index += 20;
            }
            validation.disableOwner =
                PChainOwner({threshold: disableOwnerThreshold, addresses: disableOwnerAddresses});
        }
        // Now that we have all the variable lengths, validate the input length
        uint32 expectedLength = 122 + nodeIDLength
            + (remainingBalanceOwnerAddressesLength + disableOwnerAddressesLength) * 20;
        if (input.length != expectedLength) {
            revert InvalidMessageLength(uint32(input.length), expectedLength);
        }
        // Unpack the weight
        {
            uint64 weight;
            for (uint256 i; i < 8; ++i) {
                weight |= uint64(uint8(input[i + index])) << uint64((8 * (7 - i)));
            }
            validation.weight = weight;
        }

        return validation;
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
     * @return The validationID and whether the validation period was registered or is not a
     * validator and never will be a validator due to the expiry time passing.
     */
    function unpackSubnetValidatorRegistrationMessage(bytes memory input)
        internal
        pure
        returns (bytes32, bool)
    {
        if (input.length != 39) {
            revert InvalidMessageLength(uint32(input.length), 39);
        }
        // Unpack the codec ID
        uint16 codecID;
        for (uint256 i; i < 2; ++i) {
            codecID |= uint16(uint8(input[i])) << uint16((8 * (1 - i)));
        }
        if (codecID != CODEC_ID) {
            revert InvalidCodecID(codecID);
        }

        // Unpack the type ID
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i + 2])) << uint32((8 * (3 - i)));
        }
        if (typeID != SUBNET_VALIDATOR_REGISTRATION_MESSAGE_TYPE_ID) {
            revert InvalidMessageType();
        }

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
     * @notice Packs a SubnetValidatorWeightMessage message into a byte array.
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
    function packSubnetValidatorWeightMessage(
        bytes32 validationID,
        uint64 nonce,
        uint64 weight
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            CODEC_ID, SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID, validationID, nonce, weight
        );
    }

    /**
     * @notice Unpacks a byte array as a SubnetValidatorWeight message.
     * The message format specification is the same as the one used in above for packing.
     *
     * @param input The byte array to unpack.
     * @return The validationID, nonce, and weight.
     */
    function unpackSubnetValidatorWeightMessage(bytes memory input)
        internal
        pure
        returns (bytes32, uint64, uint64)
    {
        if (input.length != 54) {
            revert InvalidMessageLength(uint32(input.length), 54);
        }

        // Unpack the codec ID.
        uint16 codecID;
        for (uint256 i; i < 2; ++i) {
            codecID |= uint16(uint8(input[i])) << uint16((8 * (1 - i)));
        }
        if (codecID != CODEC_ID) {
            revert InvalidCodecID(codecID);
        }

        // Unpack the type ID.
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i + 2])) << uint32((8 * (3 - i)));
        }
        if (typeID != SUBNET_VALIDATOR_WEIGHT_MESSAGE_TYPE_ID) {
            revert InvalidMessageType();
        }

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
        if (input.length != 46) {
            revert InvalidMessageLength(uint32(input.length), 46);
        }

        // Unpack the codec ID.
        uint16 codecID;
        for (uint256 i; i < 2; ++i) {
            codecID |= uint16(uint8(input[i])) << uint16((8 * (1 - i)));
        }
        if (codecID != CODEC_ID) {
            revert InvalidCodecID(codecID);
        }

        // Unpack the type ID.
        uint32 typeID;
        for (uint256 i; i < 4; ++i) {
            typeID |= uint32(uint8(input[i + 2])) << uint32((8 * (3 - i)));
        }
        if (typeID != VALIDATION_UPTIME_MESSAGE_TYPE_ID) {
            revert InvalidMessageType();
        }

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
