// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    IValidatorManager, ValidatorStatus, ValidatorManagerSettings
} from "./IValidatorManager.sol";
import {IRewardCalculator} from "./IRewardCalculator.sol";

struct PoSValidatorManagerSettings {
    ValidatorManagerSettings baseSettings;
    uint256 minimumStakeAmount;
    uint256 maximumStakeAmount;
    uint64 minimumStakeDuration;
    IRewardCalculator rewardCalculator;
}

struct Delegator {
    uint64 weight;
    uint64 startedAt;
    uint64 endedAt;
    ValidatorStatus status;
}

interface IPoSValidatorManager is IValidatorManager {
    /**
     * @notice Event emitted when a validator's uptime is updated.
     * @param validationID The ID of the validation period
     * @param uptime The new uptime of the validator
     */
    event ValidationUptimeUpdated(bytes32 indexed validationID, uint64 uptime);

    event DelegatorAdded(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        address indexed delegator,
        uint256 weight,
        uint256 startTime
    );

    event DelegatorRegistered(
        bytes32 indexed validationID, address indexed delegator, uint256 weight, uint256 startTime
    );

    event DelegatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        address indexed delegator,
        uint256 endTime
    );

    event DelegationEnded(bytes32 indexed validationID, address indexed delegator);

    /**
     * @notice Begins the process of ending an active validation period. The validation period must have been previously
     * started by a successful call to {completeValidatorRegistration} with the given validationID.
     * Any rewards for this validation period will stop accruing when this function is called.
     * @param validationID The ID of the validation being ended.
     * @param includeUptimeProof Whether or not an uptime proof is provided for the validation period.
     * If no uptime proof is provided, the validation uptime will be assumed to be 0.
     * @param messageIndex If {includeUptimeProof} is true, the index of the Warp message to be received providing the
     * uptime proof.
     */
    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external;

    function resendDelegatorRegistration(bytes32 validationID, address delegator) external;

    function completeDelegatorRegistration(uint32 messageIndex, address delegator) external;

    function initializeEndDelegation(bytes32 validationID) external;

    function resendEndDelegation(bytes32 validationID, address delegator) external;

    function completeEndDelegation(uint32 messageIndex, address delegator) external;
}
