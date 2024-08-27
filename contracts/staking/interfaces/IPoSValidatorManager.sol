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

    /**
     * @notice Event emitted when a delegator registration is initiated
     * @param validationID The ID of the validation period
     * @param delegator The address of the delegator
     * @param weight The weight of the delegator
     * @param startTime The time at which the registration was initiated
     */
    event DelegatorAdded(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        address indexed delegator,
        uint256 weight,
        uint256 startTime
    );

    /**
     * @notice Event emitted when a delegator registration is completed
     * @param validationID The ID of the validation period
     * @param delegator The address of the delegator
     * @param weight The weight of the delegator
     * @param validatorWeight The updated weight of the validator the delegator is delegating to
     * @param startTime The time at which the registration was completed
     */
    event DelegatorRegistered(
        bytes32 indexed validationID,
        address indexed delegator,
        uint256 weight,
        uint256 validatorWeight,
        uint256 startTime
    );

    /**
     * @notice Event emitted when delegator removal is initiated
     * @param validationID The ID of the validation period
     * @param setWeightMessageID The ID of the Warp message that updates the validator's weight on the P-Chain
     * @param delegator The address of the delegator
     * @param endTime The time at which the removal was initiated
     */
    event DelegatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        address indexed delegator,
        uint256 endTime
    );

    /**
     * @notice Event emitted when delegator removal is completed
     * @param validationID The ID of the validation period
     * @param delegator The address of the delegator
     * @param validatorWeight The updated weight of the validator the delegator was delegating to
     */
    event DelegationEnded(
        bytes32 indexed validationID, address indexed delegator, uint256 validatorWeight
    );

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

    /**
     * @notice Resubmits a delegator registration message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being registered.
     * @param delegator The address of the delegator being registered.
     */
    function resendDelegatorRegistration(bytes32 validationID, address delegator) external;

    /**
     * @notice Completes the delegator registration process by returning an acknowledgement of the registration of a
     * validationID from the P-Chain. After this function is called, the validator's weight is updated in the contract state.
     * @param messageIndex The index of the Warp message to be received providing the acknowledgement.
     * @param delegator The address of the delegator being registered.
     */
    function completeDelegatorRegistration(uint32 messageIndex, address delegator) external;

    /**
     * @notice Begins the process of removing a delegator from a validation period. The delegator must have been previously
     * registered with the given validationID.
     * @param validationID The ID of the validation period being removed.
     */
    function initializeEndDelegation(bytes32 validationID) external;

    /**
     * @notice Resubmits a delegator end message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being ended.
     * @param delegator The address of the delegator being removed.
     */
    function resendEndDelegation(bytes32 validationID, address delegator) external;

    /**
     * @notice Completes the process of ending a delegation by receiving an acknowledgement from the P-Chain.
     * After this function is called, the validator's weight is updated in the contract state.
     * @param messageIndex The index of the Warp message to be received providing the acknowledgement.
     * @param delegator The address of the delegator being removed.
     */
    function completeEndDelegation(uint32 messageIndex, address delegator) external;
}
