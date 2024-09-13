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
    ValidatorRegistrationInput
} from "./interfaces/IValidatorManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

abstract contract ValidatorManager is
    Initializable,
    ContextUpgradeable,
    ReentrancyGuardUpgradeable,
    IValidatorManager
{
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.ValidatorManager
    struct ValidatorManagerStorage {
        /// @notice The blockchainID of the P-Chain.
        bytes32 _pChainBlockchainID;
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
        /// @notice Maps the nodeID to the validationID for active validation periods.
        mapping(bytes32 => bytes32) _activeValidators;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.ValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Unit test for storage slot and update slot
    bytes32 private constant _VALIDATOR_MANAGER_STORAGE_LOCATION =
        0xe92546d698950ddd38910d2e15ed1d923cd0a7b3dde9e2a6a3f380565559cb00;

    // solhint-disable ordering
    function _getValidatorManagerStorage()
        private
        pure
        returns (ValidatorManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := _VALIDATOR_MANAGER_STORAGE_LOCATION
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
        __ReentrancyGuard_init();
        __Context_init();
        __ValidatorManager_init_unchained(settings);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ValidatorManager_init_unchained(ValidatorManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        $._pChainBlockchainID = settings.pChainBlockchainID;
        $._subnetID = settings.subnetID;

        require(
            settings.maximumChurnPercentage <= 20,
            "ValidatorManager: maximum churn percentage too high"
        );
        require(
            settings.maximumChurnPercentage > 0,
            "ValidatorManager: maximum churn percentage cannot be zero"
        );
        $._maximumChurnPercentage = settings.maximumChurnPercentage;
        $._churnPeriodSeconds = settings.churnPeriodSeconds;

        // TODO Remove - this is a hack to get a starting total weight before
        // adding an initial validator set is implemented.
        $._churnTracker.totalWeight = 1e10;
    }

    /**
     * @notice Begins the validator registration process, and sets the initial weight for the validator.
     * @param input The inputs for a validator registration.
     * @param weight The weight of the validator being registered.
     */
    function _initializeValidatorRegistration(
        ValidatorRegistrationInput calldata input,
        uint64 weight
    ) internal virtual returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        require(
            input.registrationExpiry > block.timestamp,
            "ValidatorManager: registration expiry not in future"
        );
        require(
            block.timestamp + 2 days > input.registrationExpiry,
            "ValidatorManager: registration expiry too far in future"
        );

        // Ensure the nodeID is not the zero address, and is not already an active validator.
        require(input.nodeID != bytes32(0), "ValidatorManager: invalid node ID");
        require(
            $._activeValidators[input.nodeID] == bytes32(0),
            "ValidatorManager: node ID already active"
        );
        require(input.blsPublicKey.length == 48, "ValidatorManager: invalid blsPublicKey length");

        // Check that adding this validator would not exceed the maximum churn rate.
        _checkAndUpdateChurnTrackerAddition(weight);

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

        // Submit the message to the Warp precompile.
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(registerSubnetValidatorMessage);
        $._validationPeriods[validationID] = Validator({
            status: ValidatorStatus.PendingAdded,
            nodeID: input.nodeID,
            owner: _msgSender(),
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
        require(
            $._pendingRegisterValidationMessages[validationID].length > 0
                && $._validationPeriods[validationID].status == ValidatorStatus.PendingAdded,
            "ValidatorManager: invalid validation ID"
        );

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
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, bool validRegistration) =
            ValidatorMessages.unpackSubnetValidatorRegistrationMessage(warpMessage.payload);

        require(validRegistration, "ValidatorManager: Registration not valid");
        require(
            $._pendingRegisterValidationMessages[validationID].length > 0
                && $._validationPeriods[validationID].status == ValidatorStatus.PendingAdded,
            "ValidatorManager: invalid validation ID"
        );

        delete $._pendingRegisterValidationMessages[validationID];
        $._validationPeriods[validationID].status = ValidatorStatus.Active;
        $._validationPeriods[validationID].startedAt = uint64(block.timestamp);
        $._activeValidators[$._validationPeriods[validationID].nodeID] = validationID;
        emit ValidationPeriodRegistered(
            validationID, $._validationPeriods[validationID].weight, block.timestamp
        );
    }

    function activeValidators(bytes32 nodeID) public view returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        return $._activeValidators[nodeID];
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
    function _initializeEndValidation(bytes32 validationID) internal virtual {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Ensure the validation period is active.
        Validator memory validator = $._validationPeriods[validationID];
        require(
            validator.status == ValidatorStatus.Active, "ValidatorManager: validator not active"
        );
        require(_msgSender() == validator.owner, "ValidatorManager: sender not validator owner");

        // Check that removing this delegator would not exceed the maximum churn rate.
        _checkAndUpdateChurnTrackerRemoval(validator.weight);

        // Update the validator status to pending removal.
        // They are not removed from the active validators mapping until the P-Chain acknowledges the removal.
        validator.status = ValidatorStatus.PendingRemoved;

        // Set the end time of the validation period, since it is no longer known to be an active validator
        // on the P-Chain.
        validator.endedAt = uint64(block.timestamp);

        // Save the validator updates.
        // TODO: Optimize storage writes here (probably don't need to write the whole value).
        $._validationPeriods[validationID] = validator;

        // Submit the message to the Warp precompile.
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, _incrementAndGetNonce(validationID), 0);

        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);

        // Emit the event to signal the start of the validator removal process.
        emit ValidatorRemovalInitialized(validationID, messageID, validator.weight, block.timestamp);
    }

    /**
     * @notice Resubmits a validator end message to be sent to P-Chain to the Warp precompile.
     * Only necessary if the original message can't be delivered due to validator churn.
     */
    // solhint-disable-next-line
    function resendEndValidatorMessage(bytes32 validationID) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        Validator memory validator = $._validationPeriods[validationID];

        require(
            validator.status == ValidatorStatus.PendingRemoved,
            "ValidatorManager: Validator not pending removal"
        );

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
     * @return The Validator instance representing the completed validation period
     */
    function _completeEndValidation(uint32 messageIndex) internal returns (Validator memory) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Get the Warp message.
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);

        (bytes32 validationID, bool validRegistration) =
            ValidatorMessages.unpackSubnetValidatorRegistrationMessage(warpMessage.payload);
        require(!validRegistration, "ValidatorManager: registration still valid");

        Validator memory validator = $._validationPeriods[validationID];
        ValidatorStatus endStatus;

        // The validation status is PendingRemoved if validator removal was initiated with a call to {initiateEndValidation}.
        // The validation status is PendingAdded if the validator was never registered on the P-Chain.
        require(
            validator.status == ValidatorStatus.PendingRemoved
                || validator.status == ValidatorStatus.PendingAdded,
            "ValidatorManager: invalid validator status"
        );

        if (validator.status == ValidatorStatus.PendingRemoved) {
            endStatus = ValidatorStatus.Completed;
        } else {
            endStatus = ValidatorStatus.Invalidated;
        }
        // Remove the validator from the active validators mapping.
        delete $._activeValidators[validator.nodeID];

        // Update the validator status.
        validator.status = endStatus;
        $._validationPeriods[validationID] = validator;

        // TODO: Unlock the stake.

        // Emit event.
        emit ValidationPeriodEnded(validationID, validator.status);

        return validator;
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
        $._validationPeriods[validationID].messageNonce++;
        return $._validationPeriods[validationID].messageNonce;
    }

    function _getPChainWarpMessage(uint32 messageIndex)
        internal
        view
        returns (WarpMessage memory)
    {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "ValidatorManager: Invinvalidalid warp message");
        require(
            warpMessage.sourceChainID == $._pChainBlockchainID,
            "ValidatorManager: invalid source chain ID"
        );
        require(
            warpMessage.originSenderAddress == address(0),
            "ValidatorManager: invalid origin sender address"
        );
        return warpMessage;
    }

    function _setValidatorWeight(bytes32 validationID, uint64 weight) internal {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        $._validationPeriods[validationID].weight = weight;
    }

    /**
     * @dev Helper function to check if the stake amount to be added exceeds churn thresholds.
     */
    function _checkAndUpdateChurnTrackerAddition(uint64 weight) internal {
        _checkAndUpdateChurnTracker(weight, true);
    }

    /**
     * @dev Helper function to check if the stake amount to be removed exceeds churn thresholds.
     */
    function _checkAndUpdateChurnTrackerRemoval(uint64 weight) internal {
        _checkAndUpdateChurnTracker(weight, false);
    }

    /**
     * @dev Helper function to check if the stake weight to be added or removed would exceed the maximum stake churn
     * rate for the past churn period. If the churn rate is exceeded, the function will revert. If the churn rate is
     * not exceeded, the function will update the churn tracker with the new weight.
     */
    function _checkAndUpdateChurnTracker(uint64 weight, bool addition) private {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        if ($._maximumChurnPercentage == 0) {
            return;
        }

        uint256 currentTime = block.timestamp;
        ValidatorChurnPeriod memory churnTracker = $._churnTracker;

        if (
            churnTracker.startedAt == 0
                || currentTime >= churnTracker.startedAt + $._churnPeriodSeconds
        ) {
            churnTracker.churnAmount = weight;
            churnTracker.startedAt = currentTime;
            churnTracker.initialWeight = churnTracker.totalWeight;
        } else {
            // Churn is always additive whether the weight is being added or removed.
            churnTracker.churnAmount += weight;
        }

        require(
            // Rearranged equation of maximumChurnPercentage >= currentChurnPercentage to avoid integer division truncation.
            $._maximumChurnPercentage * churnTracker.initialWeight >= churnTracker.churnAmount * 100,
            "ValidatorManager: maximum churn rate exceeded"
        );

        if (addition) {
            churnTracker.totalWeight += weight;
        } else {
            churnTracker.totalWeight -= weight;
        }

        $._churnTracker = churnTracker;
    }
}
