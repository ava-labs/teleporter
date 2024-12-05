// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManager} from "./ValidatorManager.sol";
import {IPoASecurityModule} from "./interfaces/IPoASecurityModule.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {IACP99ValidatorManager, ValidatorRegistrationInput} from "./interfaces/IACP99ValidatorManager.sol";

/**
 * @dev Implementation of the {IPoASecurityModule} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
contract PoASecurityModule is IPoASecurityModule, OwnableUpgradeable {

    struct PoASecurityModuleStorage {
        IACP99ValidatorManager validatorManager;
    }

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoASecurityModule")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant POA_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x81773fca73a14ca21edf1cadc6ec0b26d6a44966f6e97607e90422658d423500;

    function _getPoASecurityModuleStorage()
        private
        pure
        returns (PoASecurityModuleStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := POA_VALIDATOR_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    function initialize(
        address initialOwner,
        IACP99ValidatorManager validatorManager
    ) external initializer {
        __PoASecurityModule_init(initialOwner, validatorManager);
    }

    // solhint-disable func-name-mixedcase, ordering
    function __PoASecurityModule_init(
        address initialOwner,
        IACP99ValidatorManager validatorManager
    ) internal onlyInitializing {
        __Ownable_init(initialOwner);
        __PoASecurityModule_init_unchained(validatorManager);
    }

    // solhint-disable-next-line no-empty-blocks
    function __PoASecurityModule_init_unchained(IACP99ValidatorManager validatorManager) internal onlyInitializing {
        PoASecurityModuleStorage storage $ = _getPoASecurityModuleStorage();
        $.validatorManager = validatorManager;
    }

    // solhint-enable func-name-mixedcase

    /**
     * @notice See {IPoASecurityModule-initializeValidatorRegistration}.
     */
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint64 weight
    ) external onlyOwner returns (bytes32 validationID) {
        PoASecurityModuleStorage storage $ = _getPoASecurityModuleStorage();
        return $.validatorManager.initializeValidatorRegistration(registrationInput, weight);
    }

    function completeValidatorRegistration(uint32 messageIndex) external {
        PoASecurityModuleStorage storage $ = _getPoASecurityModuleStorage();
        $.validatorManager.completeValidatorRegistration(messageIndex);
    }

    // solhint-enable ordering
    /**
     * @notice See {IPoASecurityModule-initializeEndValidation}.
     */
    function initializeEndValidation(bytes32 validationID) external override onlyOwner {
        PoASecurityModuleStorage storage $ = _getPoASecurityModuleStorage();
        $.validatorManager.initializeEndValidation(validationID);
    }

    /**
     * @notice See {IValidatorManager-completeEndValidation}.
     */
    function completeEndValidation(uint32 messageIndex) external {
        PoASecurityModuleStorage storage $ = _getPoASecurityModuleStorage();
        $.validatorManager.completeEndValidation(messageIndex);
    }
}
