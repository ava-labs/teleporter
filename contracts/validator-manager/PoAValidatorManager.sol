// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManager} from "./ValidatorManager.sol";
import {IPoAValidatorManager} from "./interfaces/IPoAValidatorManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {IACP99ValidatorManager, ValidatorRegistrationInput} from "./interfaces/IACP99ValidatorManager.sol";

/**
 * @dev Implementation of the {IPoAValidatorManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
contract PoAValidatorManager is IPoAValidatorManager, OwnableUpgradeable {

    struct PoAValidatorManagerStorage {
        IACP99ValidatorManager validatorManager;
    }

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoAValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant POA_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x81773fca73a14ca21edf1cadc6ec0b26d6a44966f6e97607e90422658d423500;

    function _getPoAValidatorManagerStorage()
        private
        pure
        returns (PoAValidatorManagerStorage storage $)
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
        __PoAValidatorManager_init(initialOwner, validatorManager);
    }

    // solhint-disable func-name-mixedcase, ordering
    function __PoAValidatorManager_init(
        address initialOwner,
        IACP99ValidatorManager validatorManager
    ) internal onlyInitializing {
        __Ownable_init(initialOwner);
        __PoAValidatorManager_init_unchained(validatorManager);
    }

    // solhint-disable-next-line no-empty-blocks
    function __PoAValidatorManager_init_unchained(IACP99ValidatorManager validatorManager) internal onlyInitializing {
        PoAValidatorManagerStorage storage $ = _getPoAValidatorManagerStorage();
        $.validatorManager = validatorManager;
    }

    // solhint-enable func-name-mixedcase

    /**
     * @notice See {IPoAValidatorManager-initializeValidatorRegistration}.
     */
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint64 weight
    ) external onlyOwner returns (bytes32 validationID) {
        PoAValidatorManagerStorage storage $ = _getPoAValidatorManagerStorage();
        return $.validatorManager.initializeValidatorRegistration(registrationInput, weight);
    }

    function completeValidatorRegistration(uint32 messageIndex) external {
        PoAValidatorManagerStorage storage $ = _getPoAValidatorManagerStorage();
        $.validatorManager.completeValidatorRegistration(messageIndex);
    }

    // solhint-enable ordering
    /**
     * @notice See {IPoAValidatorManager-initializeEndValidation}.
     */
    function initializeEndValidation(bytes32 validationID) external override onlyOwner {
        PoAValidatorManagerStorage storage $ = _getPoAValidatorManagerStorage();
        $.validatorManager.initializeEndValidation(validationID);
    }

    /**
     * @notice See {IValidatorManager-completeEndValidation}.
     */
    function completeEndValidation(uint32 messageIndex) external {
        PoAValidatorManagerStorage storage $ = _getPoAValidatorManagerStorage();
        $.validatorManager.completeEndValidation(messageIndex);
    }
}
