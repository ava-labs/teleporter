// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem
pragma solidity 0.8.25;

import {SubnetConversionData} from "./interfaces/IValidatorManager.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";

contract Codec {
    function packSubnetConversionMessage(
        bytes32 subnetConversionID
    ) external pure returns (bytes memory) {
        return ValidatorMessages.packSubnetConversionMessage(subnetConversionID);
    }

    function unpackSubnetConversionMessage(
        bytes memory input
    ) external pure returns (bytes32) {
        return ValidatorMessages.unpackSubnetConversionMessage(input);
    }

    function packSubnetConversionData(
        SubnetConversionData calldata subnetConversionData
    ) external pure returns (bytes memory) {
        return ValidatorMessages.packSubnetConversionData(subnetConversionData);
    }

    function packRegisterSubnetValidatorMessage(
        ValidatorMessages.ValidationPeriod memory validationPeriod
    ) external pure returns (bytes32, bytes memory) {
        return ValidatorMessages.packRegisterSubnetValidatorMessage(validationPeriod);
    }

    function unpackRegisterSubnetValidatorMessage(
        bytes memory input
    ) external pure returns (ValidatorMessages.ValidationPeriod memory) {
        return ValidatorMessages.unpackRegisterSubnetValidatorMessage(input);
    }

    function packSubnetValidatorRegistrationMessage(
        bytes32 validationID,
        bool valid
    ) external pure returns (bytes memory) {
        return ValidatorMessages.packSubnetValidatorRegistrationMessage(validationID, valid);
    }

    function unpackSubnetValidatorRegistrationMessage(
        bytes memory input
    ) external pure returns (bytes32, bool) {
        return ValidatorMessages.unpackSubnetValidatorRegistrationMessage(input);
    }

    function packSubnetValidatorWeightMessage(
        bytes32 validationID,
        uint64 nonce,
        uint64 weight
    ) external pure returns (bytes memory) {
        return ValidatorMessages.packSubnetValidatorWeightMessage(validationID, nonce, weight);
    }

    function unpackSubnetValidatorWeightMessage(
        bytes memory input
    ) external pure returns (bytes32, uint64, uint64) {
        return ValidatorMessages.unpackSubnetValidatorWeightMessage(input);
    }

    function packValidationUptimeMessage(
        bytes32 validationID,
        uint64 uptime
    ) external pure returns (bytes memory) {
        return ValidatorMessages.packValidationUptimeMessage(validationID, uptime);
    }

    function unpackValidationUptimeMessage(
        bytes memory input
    ) external pure returns (bytes32, uint64) {
        return ValidatorMessages.unpackValidationUptimeMessage(input);
    }
}
