// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

import {IACP99SecurityModule} from "./interfaces/IACP99SecurityModule.sol";
import {PoSValidatorManager} from "./PoSValidatorManager.sol";
import {IERC20Mintable} from "./interfaces/IERC20Mintable.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {ValidatorRegistrationInput} from "./interfaces/IValidatorManager.sol";



pragma solidity 0.8.25;

contract ERC20ACP99Manager is IACP99SecurityModule, PoSValidatorManager {
    using SafeERC20 for IERC20Mintable;
    using SafeERC20TransferFrom for IERC20Mintable;

    IERC20Mintable public token;
    uint8 public tokenDecimals;

    constructor(IERC20Mintable token_, uint8 tokenDecimals_) {
        token = token_;
        tokenDecimals = tokenDecimals_;
    }

    function handleValidatorRegistration(bytes32 validationID, uint64 weight, bytes calldata args) external {
        uint256 stakeAmount = weightToValue(weight);
        (uint16 delegationFeeBips, uint64 minStakeDuration) = abi.decode(args, (uint16, uint64));
        _acp99InitValidatorRegistration(validationID, stakeAmount, delegationFeeBips, minStakeDuration);
    }

    function handleEndValidation() external {
    }

    function handleValidatorWeightChange() external {
    }

    // unused
    function initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) external nonReentrant returns (bytes32 validationID) {
        return _initializeValidatorRegistration(
            registrationInput, delegationFeeBips, minStakeDuration, stakeAmount
        );
    }

    /**
     * @notice See {PoSValidatorManager-_lock}
     * Note: Must be guarded with reentrancy guard for safe transfer from.
     */
    function _lock(uint256 value) internal virtual override returns (uint256) {
        return token.safeTransferFrom(value);
    }

    /**
     * @notice See {PoSValidatorManager-_unlock}
     * Note: Must be guarded with reentrancy guard for safe transfer.
     */
    function _unlock(address to, uint256 value) internal virtual override {
        token.safeTransfer(to, value);
    }

    /**
     * @notice See {PoSValidatorManager-_reward}
     */
    function _reward(address account, uint256 amount) internal virtual override {
        token.mint(account, amount);
    }

    
}