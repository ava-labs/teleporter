// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorRegistrationInput} from "./IValidatorManager.sol";
import {IPoSValidatorManager, PoSValidatorRequirements} from "./IPoSValidatorManager.sol";

interface IERC20TokenStakingManager is IPoSValidatorManager {
    /**
     * @notice Begins the validator registration process. Locks the {stakeAmount} of the managers specified ERC20 token.
     * @param registrationInput The inputs for a validator registration.
     * @param requirements The requirements for the validator being registered.
     * @param stakeAmount The amount of the ERC20 token to lock in the contract as the stake.
     */
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        PoSValidatorRequirements calldata requirements,
        uint256 stakeAmount
    ) external returns (bytes32 validationID);

    function initializeDelegatorRegistration(
        bytes32 validationID,
        uint256 stakeAmount
    ) external returns (bytes32);
}
