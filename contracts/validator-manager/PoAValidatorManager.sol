// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    ValidatorManagerSettings
} from "./interfaces/IValidatorManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {IACP99SecurityModule} from "./interfaces/IACP99SecurityModule.sol";
import {ACP99ValidatorManager} from "./ACP99ValidatorManager.sol";

/**
 * @dev Implementation of the {IPoAValidatorManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
contract PoAValidatorManager is IACP99SecurityModule, OwnableUpgradeable {
    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    function initialize(
        ValidatorManagerSettings calldata settings,
        address initialOwner
    ) external initializer {
        __PoAValidatorManager_init(settings, initialOwner);
    }

    // solhint-disable func-name-mixedcase, ordering
    function __PoAValidatorManager_init(
        ValidatorManagerSettings calldata,
        address initialOwner
    ) internal onlyInitializing {
        __Ownable_init(initialOwner);
    }

    // solhint-disable-next-line no-empty-blocks
    function __PoAValidatorManager_init_unchained() internal onlyInitializing {}

    // solhint-enable ordering
    /**
     * @notice See {IPoAValidatorManager-initializeEndValidation}.
     */

    /**
     * @notice See {IValidatorManager-completeEndValidation}.
     */
    function handleInitializeValidatorRegistration(bytes32 validationID, uint64 weight, bytes calldata args) external {
    }

    function handleCompleteValidatorRegistration(bytes32 validationID) external {
        // No-op
    }

    function handleInitializeEndValidation(bytes32 validationID, bytes calldata args) external {
        // No-op
    }

    function handleCompleteEndValidation(bytes32 validationID) external {
        // No-op
    }

    function handleInitializeValidatorWeightChange(bytes32 validationID, uint64 weight, uint64 nonce, bytes calldata args) external {
        // No-op
    }

    function handleCompleteValidatorWeightChange(bytes32 validationID, bytes calldata args) external {
        // No-op
    }
}
