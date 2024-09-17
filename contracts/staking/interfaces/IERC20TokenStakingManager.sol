// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorRegistrationInput} from "./IValidatorManager.sol";
import {IPoSValidatorManager} from "./IPoSValidatorManager.sol";

interface IERC20TokenStakingManager is IPoSValidatorManager {
    /**
     * @notice Begins the validator registration process. Locks the {stakeAmount} of the managers specified ERC20 token.
     * @param registrationInput The inputs for a validator registration.
     * @param delegationFeeBips The fee that delegators must pay to delegate to this validator.
     * @param minStakeDuration The minimum amount of time this validator must be staked for.
     */
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) external returns (bytes32 validationID);

    function initializeDelegatorRegistration(
        bytes32 validationID,
        uint256 stakeAmount
    ) external returns (bytes32);
}
