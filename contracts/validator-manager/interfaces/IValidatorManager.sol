// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

enum ValidatorStatus {
    Unknown,
    PendingAdded,
    Active,
    PendingRemoved,
    Completed,
    Invalidated
}

struct Validator {
    ValidatorStatus status;
    bytes32 nodeID;
    uint64 startingWeight;
    uint64 messageNonce;
    uint64 weight;
    uint64 startedAt;
    uint64 endedAt;
}

struct ValidatorChurnPeriod {
    uint256 startedAt;
    uint256 initialWeight;
    uint256 totalWeight; // TODO add initial validator set to total weight.
    uint64 churnAmount;
}

/**
 * @notice Event emitted when validator weight is updated.
 * @param validationID The ID of the validation period
 * @param nonce The message nonce used to update the validator weight
 * @param validatorWeight The updated validator weight that is sent to the P-Chain
 * @param setWeightMessageID The ID of the Warp message that updates the validator's weight on the P-Chain
 */
event ValidatorWeightUpdate(
    bytes32 indexed validationID,
    uint64 indexed nonce,
    uint64 validatorWeight,
    bytes32 setWeightMessageID
);

struct ValidatorManagerSettings {
    bytes32 pChainBlockchainID;
    bytes32 subnetID;
    uint64 churnPeriodSeconds;
    uint8 maximumChurnPercentage;
}

struct SubnetConversionData {
    bytes32 convertSubnetTxID;
    bytes32 validatorManagerBlockchainID;
    address validatorManagerAddress;
    InitialValidator[] initialValidators;
}

struct InitialValidator {
    bytes32 nodeID;
    uint64 weight;
    bytes blsPublicKey;
}

struct ValidatorRegistrationInput {
    bytes32 nodeID;
    uint64 registrationExpiry;
    bytes blsPublicKey;
}

interface IValidatorManager {
    /**
     * @notice Emitted when a new validation period is created by stake being locked in the manager contract.
     * Note that this event does not mean that the validation period has been successfully registered on the P-Chain,
     * and rewards for this validation period will not begin accruing until the {ValidationPeriodRegistered} event is
     * emitted.
     * @param validationID The ID of the validation period being created.
     * @param nodeID The node ID of the validator being registered.
     * @param registerValidationMessageID The ID of the Warp message that will be sent to the P-Chain to register the
     * validation period.
     * @param weight The weight of the validator being registered.
     * @param registrationExpiry The Unix timestamp after which the reigistration is no longer valid on the P-Chain.
     */
    event ValidationPeriodCreated(
        bytes32 indexed validationID,
        bytes32 indexed nodeID,
        bytes32 indexed registerValidationMessageID,
        uint256 weight,
        uint64 registrationExpiry
    );

    event InitialValidatorCreated(
        bytes32 indexed validationID, bytes32 indexed nodeID, uint256 weight
    );

    /**
     * @notice Emitted when the staking manager learns that the validation period has been successfully registered
     * on the P-Chain. Rewards for this validation period will begin accruing when this event is emitted.
     * @param validationID The ID of the validation period being registered.
     * @param weight The weight of the validator being registered.
     * @param timestamp The time at which the validation period was registered with the contract.
     */
    event ValidationPeriodRegistered(
        bytes32 indexed validationID, uint256 weight, uint256 timestamp
    );

    /**
     * @notice Emitted when the process of ending a registered validation period is started by calling
     * {initializeEndValidation}. Note that the stake for this validation period remains locked until
     * a {ValidationPeriodRemoved} event is emitted.
     * @param validationID The ID of the validation period being removed.
     * @param setWeightMessageID The ID of the Warp message that updates the validator's weight on the P-Chain.
     * @param weight The weight of the validator being removed.
     * @param endTime The time at which the removal was initiated.
     */
    event ValidatorRemovalInitialized(
        bytes32 indexed validationID,
        bytes32 indexed setWeightMessageID,
        uint256 weight,
        uint256 endTime
    );

    /**
     * @notice Emitted when the stake for a validation period is unlocked and returned to the staker.
     * This is done by calling {completeEndValidation}, which provides proof from the P-Chain that the
     * validation period is not active and will never be active in the future.
     * @param validationID The ID of the validation period being removed.
     */
    event ValidationPeriodEnded(bytes32 indexed validationID, ValidatorStatus indexed status);

    /**
     * @notice Verifies and sets the initial validator set for the chain through a P-Chain
     * SubnetConversionMessage.
     * @param subnetConversionData The subnet conversion message data used to recompute and verify against the subnetConversionID.
     * @param messsageIndex The index that contains the SubnetConversionMessage Warp message containing the subnetConversionID to be verified against the provided {subnetConversionData}
     */
    function initializeValidatorSet(
        SubnetConversionData calldata subnetConversionData,
        uint32 messsageIndex
    ) external;

    /**
     * @notice Resubmits a validator registration message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being registered.
     */
    function resendRegisterValidatorMessage(bytes32 validationID) external;

    /**
     * @notice Completes the validator registration process by returning an acknowledgement of the registration of a
     * validationID from the P-Chain.
     * @param messageIndex The index of the Warp message to be received providing the acknowledgement.
     */
    function completeValidatorRegistration(uint32 messageIndex) external;

    /**
     * @notice Resubmits a validator end message to be sent to the P-Chain.
     * Only necessary if the original message can't be delivered due to validator churn.
     * @param validationID The ID of the validation period being ended.
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
     */
    function completeEndValidation(uint32 messageIndex) external;
}
