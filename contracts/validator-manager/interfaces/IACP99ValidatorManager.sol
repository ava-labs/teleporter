// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

/**
 * @dev Validator status
 */
enum ValidatorStatus {
    Unknown,
    PendingAdded,
    Active,
    PendingRemoved,
    Completed,
    Invalidated
}

/**
 * @dev Contains the active state of a Validator
 */
struct Validator {
    ValidatorStatus status;
    bytes nodeID;
    uint64 startingWeight;
    uint64 messageNonce;
    uint64 weight;
    uint64 startedAt;
    uint64 endedAt;
}

/**
 * @dev Specifies the owner of a validator's remaining balance or disable owner on the P-Chain.
 * P-Chain addresses are also 20-bytes, so we use the address type to represent them.
 */
struct PChainOwner {
    uint32 threshold;
    address[] addresses;
}

/**
 * @dev Specifies an initial validator, used in the conversion data.
 */
struct InitialValidator {
    bytes nodeID;
    bytes blsPublicKey;
    uint64 weight;
}

/**
 * @dev Description of the conversion data used to convert
 * a subnet to an L1 on the P-Chain.
 * This data is the pre-image of a hash that is authenticated by the P-Chain
 * and verified by the Validator Manager.
 */
struct ConversionData {
    bytes32 subnetID;
    bytes32 validatorManagerBlockchainID;
    address validatorManagerAddress;
    InitialValidator[] initialValidators;
}

/**
 * @dev Specifies a validator to register.
 */
struct ValidatorRegistrationInput {
    bytes nodeID;
    bytes blsPublicKey;
    uint64 registrationExpiry;
    PChainOwner remainingBalanceOwner;
    PChainOwner disableOwner;
}

interface IACP99ValidatorManager {
    function initializeValidatorSet(
        ConversionData calldata conversionData,
        uint32 messageIndex
    ) external;

    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata input,
        uint64 weight
    ) external returns (bytes32);

    function completeValidatorRegistration(uint32 messageIndex) external  returns (bytes32) ;

    function initializeEndValidation(bytes32 validationID) external;

    function completeEndValidation(uint32 messageIndex) external returns (bytes32);

    function initializeValidatorWeightChange(bytes32 validationID, uint64 weight) external returns (uint64) ;

    function completeValidatorWeightChange(bytes32 validationID) external;

    // These shouldn't be in this interface, but put them here for now
    function getChurnPeriodSeconds() external view returns (uint64);
    function getValidator(bytes32 validationID) external view returns (Validator memory);
}