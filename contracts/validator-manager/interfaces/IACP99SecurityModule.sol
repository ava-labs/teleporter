// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

interface IACP99SecurityModule {
    // Called by the ValidatorManager on initializeRegisterValidator
    function handleInitializeValidatorRegistration(
        bytes32 validationID,
        uint64 weight,
        bytes calldata args
    ) external;

    function handleCompleteValidatorRegistration(bytes32 validationID) external;

    function handleInitializeEndValidation(bytes32 validationID, bytes calldata args) external;

    function handleCompleteEndValidation(bytes32 validationID) external;

    // Called by the ValidatorManager on initializeSetValidatorWeight
    function handleInitializeValidatorWeightChange() external;

    function handleCompleteValidatorWeightChange() external;
}