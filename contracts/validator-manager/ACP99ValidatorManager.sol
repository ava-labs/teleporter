// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

import {IACP99SecurityModule} from "./interfaces/IACP99SecurityModule.sol";
import {ValidatorManager} from "./ValidatorManager.sol";
import {ConversionData,ValidatorRegistrationInput} from "./interfaces/IValidatorManager.sol";

pragma solidity 0.8.25;

abstract contract ACP99ValidatorManager is ValidatorManager {
    IACP99SecurityModule public securityModule;

    // TODO: calling this should be restricted to...who?
    function setSecurityModule(IACP99SecurityModule _securityModule) external {
        securityModule = _securityModule;
    }

    function initializeValidatorSet(
        ConversionData calldata conversionData,
        uint32 messageIndex
    ) external {
        _initializeValidatorSet(conversionData, messageIndex);
    }

    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata input,
        uint64 weight,
        bytes calldata args
    ) external returns (bytes32){
        bytes32 validationID = _initializeValidatorRegistration(input, weight);
        securityModule.handleInitializeValidatorRegistration(validationID, weight, args);
        return validationID;
    }

    function completeValidatorRegistration(uint32 messageIndex) external{
        bytes32 validationID = _completeValidatorRegistration(messageIndex);
        securityModule.handleCompleteValidatorRegistration(validationID);
    }

    function initializeEndValidation(bytes32 validationID, bytes calldata args) external{
        _initializeEndValidation(validationID);
        securityModule.handleInitializeEndValidation(validationID, args);
    }

    function completeEndValidation(uint32 messageIndex) external{
        bytes32 validationID = _completeEndValidation(messageIndex);
        securityModule.handleCompleteEndValidation(validationID);
    }

    function initializeValidatorWeightChange() external{}

    function completeValidatorWeightChange() external{}
}