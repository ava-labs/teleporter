// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorMessages} from "./ValidatorMessages.sol";
import {
    InitialValidator,
    IValidatorManager,
    IChurnTracker,
    PChainOwner,
    ConversionData,
    Validator,
    ValidatorManagerSettings,
    ValidatorRegistrationInput,
    ValidatorStatus
} from "./interfaces/IValidatorManager.sol";
import {
    IWarpMessenger,
    WarpMessage
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

/**
 * @dev Implementation of the {IValidatorManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
abstract contract ValidatorManager is Initializable, ContextUpgradeable, IValidatorManager {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.ValidatorManager

    struct ValidatorManagerStorage {
        /// @notice The l1ID associated with this validator manager.
        bytes32 _l1ID;
        uint64 _totalWeight;
        /// @notice The churn tracker used to track the amount of stake added or removed in the churn period.
        IChurnTracker _churnTracker;
        /// @notice Maps the validationID to the registration message such that the message can be re-sent if needed.
        mapping(bytes32 => bytes) _pendingRegisterValidationMessages;
        /// @notice Maps the validationID to the validator information.
        mapping(bytes32 => Validator) _validationPeriods;
        /// @notice Maps the nodeID to the validationID for validation periods that have not ended.
        mapping(bytes => bytes32) _registeredValidators;
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
    error InvalidNodeID(bytes nodeID);
    error InvalidConversionID(bytes32 encodedConversionID, bytes32 expectedConversionID);
    error InvalidTotalWeight(uint64 weight);
    error InvalidValidationID(bytes32 validationID);
    error InvalidValidatorStatus(ValidatorStatus status);
    error InvalidWarpMessage();
    error MaxChurnRateExceeded(uint64 churnAmount);
    error NodeAlreadyRegistered(bytes nodeID);
    error UnexpectedRegistrationStatus(bool validRegistration);
    error InvalidPChainOwnerThreshold(uint256 threshold, uint256 addressesLength);
    error PChainOwnerAddressesNotSorted();

    // solhint-disable ordering
    /**
     * @dev This storage is visible to child contracts for convenience.
     *      External getters would be better practice, but code size limitations are preventing this.
     *      Child contracts should probably never write to this storage.
     */
    function _getValidatorManagerStorage()
        internal
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
        $._l1ID = settings.l1ID;
        $._churnTracker = IChurnTracker(settings.churnTracker);
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
        ConversionData calldata conversionData,
        uint32 messageIndex
    ) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        if ($._initializedValidatorSet) {
            revert InvalidInitializationStatus();
        }
        // Check that the blockchainID and validator manager address in the ConversionData correspond to this contract.
        // Other validation checks are done by the P-Chain when converting the L1, so are not required here.
        if (conversionData.validatorManagerBlockchainID != WARP_MESSENGER.getBlockchainID()) {
            revert InvalidValidatorManagerBlockchainID(conversionData.validatorManagerBlockchainID);
        }
        if (address(conversionData.validatorManagerAddress) != address(this)) {
            revert InvalidValidatorManagerAddress(address(conversionData.validatorManagerAddress));
        }

        uint256 numInitialValidators = conversionData.initialValidators.length;

        uint64 totalWeight;
        for (uint32 i; i < numInitialValidators; ++i) {
            InitialValidator memory initialValidator = conversionData.initialValidators[i];
            if ($._registeredValidators[initialValidator.nodeID] != bytes32(0)) {
                revert NodeAlreadyRegistered(initialValidator.nodeID);
            }

            // Validation ID of the initial validators is the sha256 hash of the
            // convert subnet to L1 tx ID and the index of the initial validator.
            bytes32 validationID = sha256(abi.encodePacked(conversionData.l1ID, i));

            // Save the initial validator as an active validator.

            $._registeredValidators[initialValidator.nodeID] = validationID;
            $._validationPeriods[validationID].status = ValidatorStatus.Active;
            $._validationPeriods[validationID].nodeID = initialValidator.nodeID;
            $._validationPeriods[validationID].startingWeight = initialValidator.weight;
            $._validationPeriods[validationID].messageNonce = 0;
            $._validationPeriods[validationID].weight = initialValidator.weight;
            $._validationPeriods[validationID].startedAt = uint64(block.timestamp);
            $._validationPeriods[validationID].endedAt = 0;
            totalWeight += initialValidator.weight;

            emit InitialValidatorCreated(
                validationID, initialValidator.nodeID, initialValidator.weight
            );
        }
        $._churnTracker.setTotalWeight(totalWeight);

        // Verify that the sha256 hash of the L1 conversion data matches with the Warp message's conversionID.
        bytes32 conversionID = ValidatorMessages.unpackSubnetToL1ConversionMessage(
            _getPChainWarpMessage(messageIndex).payload
        );
        bytes memory encodedConversion = ValidatorMessages.packConversionData(conversionData);
        bytes32 encodedConversionID = sha256(encodedConversion);
        if (encodedConversionID != conversionID) {
            revert InvalidConversionID(encodedConversionID, conversionID);
        }

        $._initializedValidatorSet = true;
    }

    function _validatePChainOwner(PChainOwner calldata pChainOwner) internal pure {
        // If threshold is 0, addresses must be empty.
        if (pChainOwner.threshold == 0 && pChainOwner.addresses.length != 0) {
            revert InvalidPChainOwnerThreshold(pChainOwner.threshold, pChainOwner.addresses.length);
        }
        // Threshold must be less than or equal to the number of addresses.
        if (pChainOwner.threshold > pChainOwner.addresses.length) {
            revert InvalidPChainOwnerThreshold(pChainOwner.threshold, pChainOwner.addresses.length);
        }
        // Addresses must be sorted in ascending order
        for (uint256 i = 1; i < pChainOwner.addresses.length; i++) {
            // Compare current address with the previous one
            if (pChainOwner.addresses[i] < pChainOwner.addresses[i - 1]) {
                revert PChainOwnerAddressesNotSorted();
            }
        }
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

        _validatePChainOwner(input.remainingBalanceOwner);
        _validatePChainOwner(input.disableOwner);

        // Ensure the nodeID is not the zero address, and is not already an active validator.

        if (input.blsPublicKey.length != BLS_PUBLIC_KEY_LENGTH) {
            revert InvalidBLSKeyLength(input.blsPublicKey.length);
        }
        if (input.nodeID.length == 0) {
            revert InvalidNodeID(input.nodeID);
        }
        if ($._registeredValidators[input.nodeID] != bytes32(0)) {
            revert NodeAlreadyRegistered(input.nodeID);
        }

        // Check that changing the validator weight would not exceed the maximum churn rate.
        $._churnTracker.checkAndUpdateChurnTracker(weight, 0);

        (bytes32 validationID, bytes memory registerL1ValidatorMessage) = ValidatorMessages
            .packRegisterL1ValidatorMessage(
            ValidatorMessages.ValidationPeriod({
                l1ID: $._l1ID,
                nodeID: input.nodeID,
                blsPublicKey: input.blsPublicKey,
                remainingBalanceOwner: input.remainingBalanceOwner,
                disableOwner: input.disableOwner,
                registrationExpiry: input.registrationExpiry,
                weight: weight
            })
        );
        $._pendingRegisterValidationMessages[validationID] = registerL1ValidatorMessage;
        $._registeredValidators[input.nodeID] = validationID;

        // Submit the message to the Warp precompile.
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(registerL1ValidatorMessage);
        $._validationPeriods[validationID].status = ValidatorStatus.PendingAdded;
        $._validationPeriods[validationID].nodeID = input.nodeID;
        $._validationPeriods[validationID].startingWeight = weight;
        $._validationPeriods[validationID].messageNonce = 0;
        $._validationPeriods[validationID].weight = weight;
        $._validationPeriods[validationID].startedAt = 0; // The validation period only starts once the registration is acknowledged.
        $._validationPeriods[validationID].endedAt = 0;

        emit ValidationPeriodCreated(
            validationID, input.nodeID, messageID, weight, input.registrationExpiry
        );

        return validationID;
    }

    /**
     * @notice See {IValidatorManager-resendRegisterValidatorMessage}.
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
     * @notice See {IValidatorManager-completeValidatorRegistration}.
     */
    function completeValidatorRegistration(uint32 messageIndex) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        (bytes32 validationID, bool validRegistration) = ValidatorMessages
            .unpackL1ValidatorRegistrationMessage(_getPChainWarpMessage(messageIndex).payload);

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

    /**
     * @notice Returns a validation ID registered to the given nodeID
     * @param nodeID ID of the node associated with the validation ID
     */
    function registeredValidators(bytes calldata nodeID) public view returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._registeredValidators[nodeID];
    }

    /**
     * @notice Returns a validator registered to the given validationID
     * @param validationID ID of the validation period associated with the validator
     */
    function getValidator(bytes32 validationID) public view returns (Validator memory) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._validationPeriods[validationID];
    }

    /**
     * @notice Begins the process of ending an active validation period. The validation period must have been previously
     * started by a successful call to {completeValidatorRegistration} with the given validationID.
     * Any rewards for this validation period will stop accruing when this function is called.
     * @param validationID The ID of the validation period being ended.
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
        $._validationPeriods[validationID] = validator;

        (, bytes32 messageID) = _setValidatorWeight(validationID, 0);

        // Emit the event to signal the start of the validator removal process.
        emit ValidatorRemovalInitialized(validationID, messageID, validator.weight, block.timestamp);

        return validator;
    }

    /**
     * @notice See {IValidatorManager-resendEndValidatorMessage}.
     */
    function resendEndValidatorMessage(bytes32 validationID) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        Validator memory validator = $._validationPeriods[validationID];

        // The initial validator set must have been set already to have pending end validation messages.
        if (validator.status != ValidatorStatus.PendingRemoved) {
            revert InvalidValidatorStatus($._validationPeriods[validationID].status);
        }

        WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, validator.messageNonce, 0)
        );
    }

    /**
     * @notice Completes the process of ending a validation period by receiving an acknowledgement from the P-Chain
     * that the validation ID is not active and will never be active in the future.
     * Note: that this function can be used for successful validation periods that have been explicitly
     * ended by calling {initializeEndValidation} or for validation periods that never began on the P-Chain due to the
     * {registrationExpiry} being reached.
     * @return (Validation ID, Validator instance) representing the completed validation period.
     */
    function _completeEndValidation(uint32 messageIndex)
        internal
        returns (bytes32, Validator memory)
    {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Get the Warp message.
        (bytes32 validationID, bool validRegistration) = ValidatorMessages
            .unpackL1ValidatorRegistrationMessage(_getPChainWarpMessage(messageIndex).payload);
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

    function _getChurnPeriodSeconds() internal returns (uint64) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._churnTracker.getChurnPeriodSeconds();
    }

    function _setValidatorWeight(
        bytes32 validationID,
        uint64 newWeight
    ) internal returns (uint64, bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        uint64 validatorWeight = $._validationPeriods[validationID].weight;

        // Check that changing the validator weight would not exceed the maximum churn rate.
        $._churnTracker.checkAndUpdateChurnTracker(newWeight, validatorWeight);

        uint64 nonce = _incrementAndGetNonce(validationID);

        $._validationPeriods[validationID].weight = newWeight;

        // Submit the message to the Warp precompile.
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packL1ValidatorWeightMessage(validationID, nonce, newWeight)
        );

        emit ValidatorWeightUpdate({
            validationID: validationID,
            nonce: nonce,
            weight: newWeight,
            setWeightMessageID: messageID
        });

        return (nonce, messageID);
    }
}
