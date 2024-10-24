// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorRegistrationInput} from "./IValidatorManager.sol";
import {IPoSValidatorManager} from "./IPoSValidatorManager.sol";

/**
 * Proof of Stake Validator Manager that stakes the blockchain's native tokens.
 */
interface INativeTokenStakingManager is IPoSValidatorManager {
    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param registrationInput The inputs for a validator registration.
     * @param delegationFeeBips The fee that delegators must pay to delegate to this validator.
     * @param minStakeDuration The minimum amount of time this validator must be staked for in seconds.
     */
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint16 delegationFeeBips,
        uint64 minStakeDuration
    ) external payable returns (bytes32 validationID);

    /**
     * @notice Begins the delegator registration process. Locks the provided native asset in the contract as the stake.
     * @param validationID The ID of the validator to stake to.
     */
    function initializeDelegatorRegistration(bytes32 validationID)
        external
        payable
        returns (bytes32);
}
