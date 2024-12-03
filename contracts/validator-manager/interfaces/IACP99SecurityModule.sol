// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

interface IACP99SecurityModule {
    // Called by the ValidatorManager on initializeRegisterValidator
    function handleValidatorRegistration(
        bytes32 validationID,
        uint64 weight,
        bytes calldata args
    ) external;

    function handleEndValidation() external;

    // Called by the ValidatorManager on initializeSetValidatorWeight
    function handleValidatorWeightChange() external;
}