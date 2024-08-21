// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IValidatorManager} from "./IValidatorManager.sol";

interface IPoAValidatorManager is IValidatorManager {
    /**
     * @notice Begins the validator registration process, and sets the {weight} of the validator.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The time at which the reigistration is no longer valid on the P-Chain.
     */
    function initializeValidatorRegistration(
        uint64 weight,
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature
    ) external returns (bytes32 validationID);

    /**
     * @notice Begins the process of ending an active validation period. The validation period must have been previously
     * started by a successful call to {completeValidatorRegistration} with the given validationID.
     * Any rewards for this validation period will stop accruing when this function is called.
     * @param validationID The ID of the validation being ended.
     */
    function initializeEndValidation(bytes32 validationID) external;
}
