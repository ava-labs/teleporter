// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

import {IACP99SecurityModule} from "./interfaces/IACP99SecurityModule.sol";
import {ValidatorManager} from "./ValidatorManager.sol";
import {IACP99ValidatorManager, ConversionData, ValidatorRegistrationInput} from "./interfaces/IACP99ValidatorManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {ValidatorManagerSettings} from "./interfaces/IValidatorManager.sol";

pragma solidity 0.8.25;

contract ACP99ValidatorManager is IACP99ValidatorManager, ValidatorManager {
    struct ACP99ValidatorManagerStorage {
        IACP99SecurityModule securityModule;
    }
    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.ACP99ValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant ACP_99_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x6d7896c90f86967e463241c430aa4c1ef638639857cb8d4f18c905fa5443d600;

    function _getACP99ValidatorManagerStorage()
        private
        pure
        returns (ACP99ValidatorManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := ACP_99_VALIDATOR_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    function initialize(ValidatorManagerSettings calldata settings, IACP99SecurityModule securityModule) external initializer {
        __ACP99ValidatorManager_init(settings, securityModule);
    }

    // TODO: calling this should be restricted to...who?
    function setSecurityModule(IACP99SecurityModule securityModule) external {
        _getACP99ValidatorManagerStorage().securityModule = securityModule;
    }

    function getSecurityModule() external view returns (IACP99SecurityModule) {
        return _getACP99ValidatorManagerStorage().securityModule;
    }

    function __ACP99ValidatorManager_init(ValidatorManagerSettings calldata settings, IACP99SecurityModule securityModule) internal onlyInitializing {
        __ValidatorManager_init(settings);
        __ACP99ValidatorManager_init_unchained(securityModule);
    }

    function __ACP99ValidatorManager_init_unchained(IACP99SecurityModule securityModule) internal onlyInitializing {
        ACP99ValidatorManagerStorage storage $ = _getACP99ValidatorManagerStorage();
        $.securityModule = securityModule;
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
        ACP99ValidatorManagerStorage storage $ = _getACP99ValidatorManagerStorage();
        // address securityModule = address($.securityModule);

        // // Delegate call so that the sender can approval ERC20 transfers by the validator manager
        // (bool success, bytes memory data) = securityModule.delegatecall(abi.encodeWithSelector($.securityModule.handleInitializeValidatorRegistration.selector, validationID, weight, args));
        // require(success, string(data));
        
        $.securityModule.handleInitializeValidatorRegistration(validationID, weight, args);
        return validationID;
    }

    function completeValidatorRegistration(uint32 messageIndex) external{
        bytes32 validationID = _completeValidatorRegistration(messageIndex);
        _getACP99ValidatorManagerStorage().securityModule.handleCompleteValidatorRegistration(validationID);
    }

    function initializeEndValidation(bytes32 validationID, bytes calldata args) external{
        _initializeEndValidation(validationID);
        _getACP99ValidatorManagerStorage().securityModule.handleInitializeEndValidation(validationID, args);
    }

    function completeEndValidation(uint32 messageIndex) external{
        bytes32 validationID = _completeEndValidation(messageIndex);
        ACP99ValidatorManagerStorage storage $ = _getACP99ValidatorManagerStorage();
        // address securityModule = address($.securityModule);

        // (bool success, bytes memory data) = securityModule.delegatecall(abi.encodeWithSelector($.securityModule.handleCompleteEndValidation.selector, validationID));
        // require(success, string(data));

        $.securityModule.handleCompleteEndValidation(validationID);
    }

    function initializeValidatorWeightChange(bytes32 validationID, uint64 weight, bytes calldata args) external{
        (uint64 nonce, ) = _setValidatorWeight(validationID, weight);
        _getACP99ValidatorManagerStorage().securityModule.handleInitializeValidatorWeightChange(validationID, weight, nonce, args);
    }

    function completeValidatorWeightChange(bytes32 validationID, bytes calldata args) external{
        _getACP99ValidatorManagerStorage().securityModule.handleCompleteValidatorWeightChange(validationID, args);
    }
}