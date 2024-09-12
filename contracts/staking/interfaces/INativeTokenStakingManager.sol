// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorRegistrationInput} from "./IValidatorManager.sol";
import {IPoSValidatorManager, PoSValidatorRequirements} from "./IPoSValidatorManager.sol";

interface INativeTokenStakingManager is IPoSValidatorManager {
    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param registrationInput The inputs for a validator registration.
     * @param delegationFee The fee that delegators must pay to delegate to this validator.
     * @param minStakeDuration The minimum amount of time this validator must be staked for.
     */
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint256 delegationFee,
        uint256 minStakeDuration
    ) external payable returns (bytes32 validationID);

    function initializeDelegatorRegistration(bytes32 validationID)
        external
        payable
        returns (bytes32);
}
