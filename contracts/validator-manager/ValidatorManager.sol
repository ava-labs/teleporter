// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    IValidatorManager,
    ValidatorManagerSettings,
    ValidatorChurnPeriod,
    ValidatorStatus,
    Validator,
    ValidatorChurnPeriod,
    ValidatorWeightUpdate,
    SubnetConversionData,
    InitialValidator,
    ValidatorRegistrationInput
} from "./interfaces/IValidatorManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

abstract contract ValidatorManager is Initializable, ContextUpgradeable, IValidatorManager {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.ValidatorManager

    struct ValidatorManagerStorage {
        /// @notice The subnetID associated with this validator manager.
        bytes32 _subnetID;
        /// @notice The number of seconds after which to reset the churn tracker.
        uint64 _churnPeriodSeconds;
        /// @notice The maximum churn rate allowed per churn period.
        uint8 _maximumChurnPercentage;
        /// @notice The churn tracker used to track the amount of stake added or removed in the churn period.
        ValidatorChurnPeriod _churnTracker;
        /// @notice Maps the validationID to the registration message such that the message can be re-sent if needed.
        mapping(bytes32 => bytes) _pendingRegisterValidationMessages;
        /// @notice Maps the validationID to the validator information.
        mapping(bytes32 => Validator) _validationPeriods;
        /// @notice Maps the nodeID to the validationID for validation periods that have not ended.
        mapping(bytes32 => bytes32) _registeredValidators;
        /// @notice Boolean that indicates if the initial validator set has been set.
        bool _initializedValidatorSet;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.ValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant VALIDATOR_MANAGER_STORAGE_LOCATION =
        0xe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00;

    uint8 public constant MAXIMUM_CHURN_PERCENTAGE_LIMIT = 20;
    uint64 public constant MAXIMUM_REGISTRATION_EXPIRY_LENGTH = 2 days;
    uint32 public constant ADDRESS_LENGTH = 20; // This is only used as a packed uint32
    uint8 public constant BLS_PUBLIC_KEY_LENGTH = 48;
    bytes32 public constant P_CHAIN_BLOCKCHAIN_ID = bytes32(0);

    error InvalidValidatorManagerAddress(address validatorManagerAddress);
    error InvalidWarpOriginSenderAddress(address senderAddress);
    error InvalidValidatorManagerBlockchainID(bytes32 blockchainID);
    error InvalidWarpSourceChainID(bytes32 sourceChainID);
    error InvalidRegistrationExpiry(uint64 registrationExpiry);
    error InvalidInitializationStatus();
    error InvalidMaximumChurnPercentage(uint8 maximumChurnPercentage);
    error InvalidBLSKeyLength(uint256 length);
    error InvalidNodeID(bytes32 nodeID);
    error InvalidSubnetConversionID(
        bytes32 encodedSubnetConversionID, bytes32 expectedSubnetConversionID
    );
    error InvalidValidationID(bytes32 validationID);
    error InvalidValidatorStatus(ValidatorStatus status);
    error InvalidWarpMessage();
    error MaxChurnRateExceeded(uint64 churnAmount);
    error NodeAlreadyRegistered(bytes32 nodeID);
    error UnexpectedRegistrationStatus(bool validRegistration);

    // solhint-disable ordering
    function _getValidatorManagerStorage()
        private
        pure
        returns (ValidatorManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := VALIDATOR_MANAGER_STORAGE_LOCATION
        }
    }

    /**
     * @notice Warp precompile used for sending and receiving Warp messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    // solhint-disable-next-line func-name-mixedcase
    function __ValidatorManager_init(ValidatorManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        __Context_init();
        __ValidatorManager_init_unchained(settings);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ValidatorManager_init_unchained(ValidatorManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        $._subnetID = settings.subnetID;

        if (
            settings.maximumChurnPercentage > MAXIMUM_CHURN_PERCENTAGE_LIMIT
                || settings.maximumChurnPercentage == 0
        ) {
            revert InvalidMaximumChurnPercentage(settings.maximumChurnPercentage);
        }

        $._maximumChurnPercentage = settings.maximumChurnPercentage;
        $._churnPeriodSeconds = settings.churnPeriodSeconds;
    }

    modifier initializedValidatorSet() {
        if (!_getValidatorManagerStorage()._initializedValidatorSet) {
            revert InvalidInitializationStatus();
        }
        _;
    }

    /**
     * @notice See {IValidatorManager-initializeValidatorSet}.
     */
    function initializeValidatorSet(
        SubnetConversionData calldata subnetConversionData,
        uint32 messageIndex
    ) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        if ($._initializedValidatorSet) {
            revert InvalidInitializationStatus();
        }
        // Check that the blockchainID and validator manager address in the subnetConversionData correspond to this contract.
        // Other validation checks are done by the P-Chain when converting the subnet, so are not required here.
        if (subnetConversionData.validatorManagerBlockchainID != WARP_MESSENGER.getBlockchainID()) {
            revert InvalidValidatorManagerBlockchainID(
                subnetConversionData.validatorManagerBlockchainID
            );
        }
        if (address(subnetConversionData.validatorManagerAddress) != address(this)) {
            revert InvalidValidatorManagerAddress(
                address(subnetConversionData.validatorManagerAddress)
            );
        }

        uint256 numInitialValidators = subnetConversionData.initialValidators.length;

        uint256 totalWeight;
        for (uint32 i; i < numInitialValidators; ++i) {
            InitialValidator memory initialValidator = subnetConversionData.initialValidators[i];
            bytes32 nodeID = initialValidator.nodeID;

            if ($._registeredValidators[nodeID] != bytes32(0)) {
                revert NodeAlreadyRegistered(nodeID);
            }

            // Validation ID of the initial validators is the sha256 hash of the
            // convert Subnet tx ID and the index of the initial validator.
            bytes32 validationID =
                sha256(abi.encodePacked(subnetConversionData.convertSubnetTxID, i));

            // Save the initial validator as an active validator.

            $._registeredValidators[nodeID] = validationID;
            $._validationPeriods[validationID] = Validator({
                status: ValidatorStatus.Active,
                nodeID: initialValidator.nodeID,
                startingWeight: initialValidator.weight,
                messageNonce: 0,
                weight: initialValidator.weight,
                startedAt: uint64(block.timestamp),
                endedAt: 0
            });
            totalWeight += initialValidator.weight;

            emit InitialValidatorCreated(
                validationID, initialValidator.nodeID, initialValidator.weight
            );
        }
        $._churnTracker.totalWeight = totalWeight;

        // Verify that the sha256 hash of the Subnet conversion data matches with the Warp message's subnetConversionID.
        bytes32 subnetConversionID = ValidatorMessages.unpackSubnetConversionMessage(
            _getPChainWarpMessage(messageIndex).payload
        );
        bytes memory encodedConversion =
            ValidatorMessages.packSubnetConversionData(subnetConversionData);
        bytes32 encodedSubnetConversionID = sha256(encodedConversion);
        if (encodedSubnetConversionID != subnetConversionID) {
            revert InvalidSubnetConversionID(encodedSubnetConversionID, subnetConversionID);
        }

        $._initializedValidatorSet = true;
    }

    /**
     * @notice Begins the validator registration process, and sets the initial weight for the validator.
     * This is the only method related to validator registration and removal that needs the initializedValidatorSet
     * modifier. All others are guarded by checking the validator status changes initialized in this function.
     * @param input The inputs for a validator registration.
     * @param weight The weight of the validator being registered.
     */
    function _initializeValidatorRegistration(
        ValidatorRegistrationInput calldata input,
        uint64 weight
    ) internal virtual initializedValidatorSet returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        if (
            input.registrationExpiry <= block.timestamp
                || input.registrationExpiry >= block.timestamp + MAXIMUM_REGISTRATION_EXPIRY_LENGTH
        ) {
            revert InvalidRegistrationExpiry(input.registrationExpiry);
        }

        // Ensure the nodeID is not the zero address, and is not already an active validator.

        if (input.blsPublicKey.length != BLS_PUBLIC_KEY_LENGTH) {
            revert InvalidBLSKeyLength(input.blsPublicKey.length);
        }
        if (input.nodeID == bytes32(0)) {
            revert InvalidNodeID(input.nodeID);
        }
        if ($._registeredValidators[input.nodeID] != bytes32(0)) {
            revert NodeAlreadyRegistered(input.nodeID);
        }

        // Check that adding this validator would not exceed the maximum churn rate.
        _checkAndUpdateChurnTracker(weight, 0);

        (bytes32 validationID, bytes memory registerSubnetValidatorMessage) = ValidatorMessages
            .packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                subnetID: $._subnetID,
                nodeID: input.nodeID,
                weight: weight,
                blsPublicKey: input.blsPublicKey,
                registrationExpiry: input.registrationExpiry
            })
        );
        $._pendingRegisterValidationMessages[validationID] = registerSubnetValidatorMessage;
        $._registeredValidators[input.nodeID] = validationID;

        // Submit the message to the Warp precompile.
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(registerSubnetValidatorMessage);
        $._validationPeriods[validationID] = Validator({
            status: ValidatorStatus.PendingAdded,
            nodeID: input.nodeID,
            startingWeight: weight,
            messageNonce: 0,
            weight: weight,
            startedAt: 0, // The validation period only starts once the registration is acknowledged.
            endedAt: 0
        });

        emit ValidationPeriodCreated(
            validationID, input.nodeID, messageID, weight, input.registrationExpiry
        );

        return validationID;
    }

    /**
     * @notice Resubmits a validator registration message to be sent to P-Chain to the Warp precompile.
     * Only necessary if the original message can't be delivered due to validator churn.
     */
    function resendRegisterValidatorMessage(bytes32 validationID) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        // The initial validator set must have been set already to have pending register validation messages.
        if ($._pendingRegisterValidationMessages[validationID].length == 0) {
            revert InvalidValidationID(validationID);
        }
        if ($._validationPeriods[validationID].status != ValidatorStatus.PendingAdded) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        // Submit the message to the Warp precompile.
        WARP_MESSENGER.sendWarpMessage($._pendingRegisterValidationMessages[validationID]);
    }

    /**
     * @notice Completes the validator registration process by returning an acknowledgement of the registration of a
     * validationID from the P-Chain.
     * @param messageIndex The index of the Warp message to be received providing the acknowledgement.
     */
    function completeValidatorRegistration(uint32 messageIndex) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        (bytes32 validationID, bool validRegistration) = ValidatorMessages
            .unpackSubnetValidatorRegistrationMessage(_getPChainWarpMessage(messageIndex).payload);

        if (!validRegistration) {
            revert UnexpectedRegistrationStatus(validRegistration);
        }
        // The initial validator set must have been set already to have pending register validation messages.
        if ($._pendingRegisterValidationMessages[validationID].length == 0) {
            revert InvalidValidationID(validationID);
        }
        if ($._validationPeriods[validationID].status != ValidatorStatus.PendingAdded) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        delete $._pendingRegisterValidationMessages[validationID];
        $._validationPeriods[validationID].status = ValidatorStatus.Active;
        $._validationPeriods[validationID].startedAt = uint64(block.timestamp);
        emit ValidationPeriodRegistered(
            validationID, $._validationPeriods[validationID].weight, block.timestamp
        );
    }

    function registeredValidators(bytes32 nodeID) public view returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._registeredValidators[nodeID];
    }

    function getValidator(bytes32 validationID) public view returns (Validator memory) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._validationPeriods[validationID];
    }

    /**
     * @notice Begins the process of ending an active validation period. The validation period must have been previously
     * started by a successful call to {completeValidatorRegistration} with the given validationID.
     * Any rewards for this validation period will stop accruing when this function is called.
     * @param validationID The ID of the validation being ended.
     */
    function _initializeEndValidation(bytes32 validationID)
        internal
        virtual
        returns (Validator memory)
    {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Ensure the validation period is active.
        // The initial validator set must have been set already to have active validators.
        Validator memory validator = $._validationPeriods[validationID];
        if (validator.status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        // Update the validator status to pending removal.
        // They are not removed from the active validators mapping until the P-Chain acknowledges the removal.
        validator.status = ValidatorStatus.PendingRemoved;

        // Set the end time of the validation period, since it is no longer known to be an active validator
        // on the P-Chain.
        validator.endedAt = uint64(block.timestamp);

        // Save the validator updates.
        // TODO: Optimize storage writes here (probably don't need to write the whole value).
        $._validationPeriods[validationID] = validator;

        (, bytes32 messageID) = _setValidatorWeight(validationID, 0);

        // Emit the event to signal the start of the validator removal process.
        emit ValidatorRemovalInitialized(validationID, messageID, validator.weight, block.timestamp);

        return validator;
    }

    /**
     * @notice Resubmits a validator end message to be sent to P-Chain to the Warp precompile.
     * Only necessary if the original message can't be delivered due to validator churn.
     */
    // solhint-disable-next-line
    function resendEndValidatorMessage(bytes32 validationID) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        Validator memory validator = $._validationPeriods[validationID];

        // The initial validator set must have been set already to have pending end validation messages.
        if (validator.status != ValidatorStatus.PendingRemoved) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packSetSubnetValidatorWeightMessage(
                validationID, validator.messageNonce, 0
            )
        );
    }

    /**
     * @notice Completes the process of ending a validation period by receiving an acknowledgement from the P-Chain
     * that the validation ID is not active and will never be active in the future.
     *  Note that this function can be used for successful validation periods that have been explicitly
     * ended by calling {initializeEndValidation} or for validation periods that never began on the P-Chain due to the
     * {registrationExpiry} being reached.
     * @return The Validator instance representing the completed validation period and the corresponding validation ID.
     */
    function _completeEndValidation(uint32 messageIndex)
        internal
        returns (bytes32, Validator memory)
    {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Get the Warp message.
        (bytes32 validationID, bool validRegistration) = ValidatorMessages
            .unpackSubnetValidatorRegistrationMessage(_getPChainWarpMessage(messageIndex).payload);
        if (validRegistration) {
            revert UnexpectedRegistrationStatus(validRegistration);
        }

        Validator memory validator = $._validationPeriods[validationID];

        // The validation status is PendingRemoved if validator removal was initiated with a call to {initiateEndValidation}.
        // The validation status is PendingAdded if the validator was never registered on the P-Chain.
        // The initial validator set must have been set already to have pending validation messages.
        if (
            validator.status != ValidatorStatus.PendingRemoved
                && validator.status != ValidatorStatus.PendingAdded
        ) {
            revert InvalidValidatorStatus(validator.status);
        }

        if (validator.status == ValidatorStatus.PendingRemoved) {
            validator.status = ValidatorStatus.Completed;
        } else {
            validator.status = ValidatorStatus.Invalidated;
        }
        // Remove the validator from the registered validators mapping.
        delete $._registeredValidators[validator.nodeID];

        // Update the validator.
        $._validationPeriods[validationID] = validator;

        // Emit event.
        emit ValidationPeriodEnded(validationID, validator.status);

        return (validationID, validator);
    }

    /**
     * @notice Returns the validator's weight. This weight is not guaranteed to be known by the P-Chain
     * @return The weight of the validator. If the validation ID does not exist, the weight will be 0.
     */
    function getWeight(bytes32 validationID) external view returns (uint64) {
        return getValidator(validationID).weight;
    }

    function _incrementAndGetNonce(bytes32 validationID) internal returns (uint64) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return ++$._validationPeriods[validationID].messageNonce;
    }

    function _getPChainWarpMessage(uint32 messageIndex)
        internal
        view
        returns (WarpMessage memory)
    {
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        if (!valid) {
            revert InvalidWarpMessage();
        }
        // Must match to P-Chain blockchain id, which is 0.
        if (warpMessage.sourceChainID != P_CHAIN_BLOCKCHAIN_ID) {
            revert InvalidWarpSourceChainID(warpMessage.sourceChainID);
        }
        if (warpMessage.originSenderAddress != address(0)) {
            revert InvalidWarpOriginSenderAddress(warpMessage.originSenderAddress);
        }

        return warpMessage;
    }

    function _setValidatorWeight(
        bytes32 validationID,
        uint64 newWeight
    ) internal returns (uint64, bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        uint64 validatorWeight = $._validationPeriods[validationID].weight;

        // Check that changing the validator weight would not exceed the maximum churn rate.
        _checkAndUpdateChurnTracker(newWeight, validatorWeight);

        uint64 nonce = _incrementAndGetNonce(validationID);

        $._validationPeriods[validationID].weight = newWeight;

        // Submit the message to the Warp precompile.
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packSetSubnetValidatorWeightMessage(validationID, nonce, newWeight)
        );

        emit ValidatorWeightUpdate({
            validationID: validationID,
            nonce: nonce,
            validatorWeight: newWeight,
            setWeightMessageID: messageID
        });

        return (nonce, messageID);
    }

    /**
     * @dev Helper function to check if the stake weight to be added or removed would exceed the maximum stake churn
     * rate for the past churn period. If the churn rate is exceeded, the function will revert. If the churn rate is
     * not exceeded, the function will update the churn tracker with the new weight.
     */
    function _checkAndUpdateChurnTracker(
        uint64 newValidatorWeight,
        uint64 oldValidatorWeight
    ) private {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        uint64 weightChange;
        if (newValidatorWeight > oldValidatorWeight) {
            weightChange = newValidatorWeight - oldValidatorWeight;
        } else {
            weightChange = oldValidatorWeight - newValidatorWeight;
        }

        uint256 currentTime = block.timestamp;
        ValidatorChurnPeriod memory churnTracker = $._churnTracker;

        if (
            churnTracker.startedAt == 0
                || currentTime >= churnTracker.startedAt + $._churnPeriodSeconds
        ) {
            churnTracker.churnAmount = weightChange;
            churnTracker.startedAt = currentTime;
            churnTracker.initialWeight = churnTracker.totalWeight;
        } else {
            // Churn is always additive whether the weight is being added or removed.
            churnTracker.churnAmount += weightChange;
        }

        // Rearranged equation of maximumChurnPercentage >= currentChurnPercentage to avoid integer division truncation.
        if ($._maximumChurnPercentage * churnTracker.initialWeight < churnTracker.churnAmount * 100)
        {
            revert MaxChurnRateExceeded(churnTracker.churnAmount);
        }

        // Two separate calculations because we're using uints and (newValidatorWeight - oldValidatorWeight) could underflow.
        churnTracker.totalWeight += newValidatorWeight;
        churnTracker.totalWeight -= oldValidatorWeight;

        $._churnTracker = churnTracker;
    }
}
