// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

interface IStakingManager {
    /**
     * @notice Emitted when a new validation period is created by stake being locked in the manager contract.
     * Note that this event does not mean that the validation period has been successfully registered on the P-Chain,
     * and rewards for this validation period will not begin accruing until the {ValidationPeriodRegistered} event is
     * emitted.
     */
    event ValidationPeriodCreated(
        bytes32 indexed validationID,
        bytes32 indexed nodeID,
        bytes32 indexed registerValidationMessageID,
        uint256 weight,
        uint64 registrationExpiry
    );

    /**
     * @notice Emitted when the staking manager learns that the validation period has been successfully registered
     * on the P-Chain. Rewards for this validation period will begin accruing when this event is emitted.
     */
    event ValidationPeriodRegistered(
        bytes32 indexed validationID, uint256 stakeAmount, uint256 timestamp
    );

    /**
     * @notice Emitted when the process of ending a registered validation period is started by calling
     * {initializeEndValidation}. Note that the stake for this validation period remains locked until
     * a {ValidationPeriodRemoved} event is emitted.
     */
    event ValidatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        uint256 stakeAmount,
        uint256 endTime,
        uint64 uptime
    );

    /**
     * @notice Emitted when the stake for a validation period is unlocked and returned to the staker.
     * This is done by calling {completeEndValidation}, which provides proof from the P-Chain that the
     * validation period is not active and will never be active in the future.
     */
    event ValidationPeriodEnded(bytes32 indexed validationID);

    /**
     * @notice Resubmits a validator registration message to be sent to P-Chain to the Warp precompile.
     * Only necessary if the original message can't be delivered due to validator churn.
     */
    function resendRegisterValidatorMessage(bytes32 validationID) external;

    /**
     * @notice Completes the validator registration process by returning an acknowledgement of the registration of a
     * validationID from the P-Chain.
     * @param messageIndex The index of the Warp message to be received providing the acknowledgement.
     */
    function completeValidatorRegistration(uint32 messageIndex) external;

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
     * @notice Resubmits a validator end message to be sent to P-Chain to the Warp precompile.
     * Only necessary if the original message can't be delivered due to validator churn.
     */
    function resendEndValidatorMessage(bytes32 validationID) external;

    /**
     * @notice Completes the process of ending a validation period by receiving an acknowledgement from the P-Chain
     * that the validation ID is not active and will never be active in the future. Returns the the stake associated
     * with the validation. Note that this function can be used for successful validation periods that have been explicitly
     * ended by calling {initializeEndValidation} or for validation periods that never began on the P-Chain due to the
     * {registrationExpiry} being reached.
     * @param messageIndex The index of the Warp message to be received providing the proof the validation is not active
     * and never will be active on the P-Chain.
     * @param setWeightMessageType Whether or not the message type is a SetValidatorWeight message, or a
     * SubnetValidatorRegistration message (with valid set to false). Both message types are valid for ending
     * a validation period.
     */
    function completeEndValidation(uint32 messageIndex, bool setWeightMessageType) external;
}
