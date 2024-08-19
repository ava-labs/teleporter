// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    IValidatorManager,
    ValidatorManagerSettings,
    ValidatorStatus,
    Validator,
    ValidatorChurnPeriod
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
        bytes32 _pChainBlockchainID;
        bytes32 _subnetID;
        uint8 _maximumHourlyChurn;
        ValidatorChurnPeriod _churnTracker;
        // Maps the validationID to the registration message such that the message can be re-sent if needed.
        mapping(bytes32 => bytes) _pendingRegisterValidationMessages;
        // Maps the validationID to the validator information.
        mapping(bytes32 => Validator) _validationPeriods;
        // Maps the nodeID to the validationID for active validation periods.
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
        $._maximumHourlyChurn = settings.maximumHourlyChurn;
    }

    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The time at which the reigistration is no longer valid on the P-Chain.
     * @param signature The raw bytes of the Ed25519 signature over the concatenated bytes of
     * [subnetID]+[nodeID]+[blsPublicKey]+[weight]+[balance]+[expiry]. This signature must correspond to the Ed25519
     * public key that is used for the nodeID. This approach prevents NodeIDs from being unwillingly added to Subnets.
     * balance is the minimum initial $nAVAX balance that must be attached to the validator serialized as a uint64.
     * The signature field will be validated by the P-Chain. Implementations may choose to validate that the signature
     * field is well-formed but it is not required.
     */
    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 weight,
        uint64 registrationExpiry,
        bytes memory signature
    ) internal nonReentrant returns (bytes32) {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Ensure the registration expiry is in a valid range.
        require(
            registrationExpiry > block.timestamp && block.timestamp + 2 days > registrationExpiry,
            "ValidatorManager: Invalid registration expiry"
        );

        // Ensure the nodeID is not the zero address, and is not already an active validator.
        require(nodeID != bytes32(0), "ValidatorManager: Invalid node ID");
        require(
            $._activeValidators[nodeID] == bytes32(0), "ValidatorManager: Node ID already active"
        );

        // Ensure the signature is the proper length. The EVM does not provide an Ed25519 precompile to
        // validate the signature, but the P-Chain will validate the signature. If the signature is invalid,
        // the P-Chain will reject the registration, and the stake can be returned to the staker after the registration
        // expiry has passed.
        require(signature.length == 64, "ValidatorManager: Invalid signature length");

        _checkAndUpdateChurnTracker(weight);

        (bytes32 validationID, bytes memory registerSubnetValidatorMessage) = ValidatorMessages
            .packRegisterSubnetValidatorMessage(
            ValidatorMessages.ValidationInfo({
                subnetID: $._subnetID,
                nodeID: nodeID,
                weight: weight,
                registrationExpiry: registrationExpiry,
                signature: signature
            })
        );
        $._pendingRegisterValidationMessages[validationID] = registerSubnetValidatorMessage;

        // Submit the message to the Warp precompile.
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(registerSubnetValidatorMessage);

        $._validationPeriods[validationID] = Validator({
            status: ValidatorStatus.PendingAdded,
            nodeID: nodeID,
            weight: weight,
            startedAt: 0, // The validation period only starts once the registration is acknowledged.
            endedAt: 0,
            uptimeSeconds: 0,
            owner: _msgSender(),
            rewarded: false,
            messageNonce: 0
        });
        emit ValidationPeriodCreated(validationID, nodeID, messageID, weight, registrationExpiry);

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
            "ValidatorManager: Invalid validation ID"
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
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "ValidatorManager: Invalid warp message");

        require(
            warpMessage.sourceChainID == $._pChainBlockchainID,
            "ValidatorManager: Invalid source chain ID"
        );
        require(
            warpMessage.originSenderAddress == address(0),
            "ValidatorManager: Invalid origin sender address"
        );

        (bytes32 validationID, bool validRegistration) =
            ValidatorMessages.unpackSubnetValidatorRegistrationMessage(warpMessage.payload);
        require(validRegistration, "ValidatorManager: Registration not valid");
        require(
            $._pendingRegisterValidationMessages[validationID].length > 0
                && $._validationPeriods[validationID].status == ValidatorStatus.PendingAdded,
            "ValidatorManager: Invalid validation ID"
        );

        delete $._pendingRegisterValidationMessages[validationID];
        $._validationPeriods[validationID].status = ValidatorStatus.Active;
        $._validationPeriods[validationID].startedAt = uint64(block.timestamp);
        $._activeValidators[$._validationPeriods[validationID].nodeID] = validationID;
        emit ValidationPeriodRegistered(
            validationID, $._validationPeriods[validationID].weight, block.timestamp
        );
    }

    /**
     * @notice Begins the process of ending an active validation period. The validation period must have been previously
     * started by a successful call to {completeValidatorRegistration} with the given validationID.
     * Any rewards for this validation period will stop accruing when this function is called.
     * @param validationID The ID of the validation being ended.
     */
    function _initializeEndValidation(
        bytes32 validationID,
        uint64 uptimeSeconds
    ) internal virtual {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Ensure the validation period is active.
        Validator memory validator = $._validationPeriods[validationID];
        require(
            validator.status == ValidatorStatus.Active, "ValidatorManager: Validator not active"
        );
        require(_msgSender() == validator.owner, "ValidatorManager: Sender not validator owner");

        // Check that removing this validator would not exceed the maximum churn rate.
        _checkAndUpdateChurnTracker(validator.weight);

        // Update the validator status to pending removal.
        // They are not removed from the active validators mapping until the P-Chain acknowledges the removal.
        validator.status = ValidatorStatus.PendingRemoved;

        // Set the end time of the validation period, since it is no longer known to be an active validator
        // on the P-Chain.
        validator.endedAt = uint64(block.timestamp);
        validator.uptimeSeconds = uptimeSeconds;

        // Save the validator updates.
        // TODO: Optimize storage writes here (probably don't need to write the whole value).
        $._validationPeriods[validationID] = validator;

        // Submit the message to the Warp precompile.
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, validator.messageNonce, 0);
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);

        // Emit the event to signal the start of the validator removal process.
        emit ValidatorRemovalInitialized(
            validationID, messageID, validator.weight, block.timestamp, uptimeSeconds
        );
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

        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, validator.messageNonce, 0);
        WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);
    }

    /**
     * @notice Completes the process of ending a validation period by receiving an acknowledgement from the P-Chain
     * that the validation ID is not active and will never be active in the future. Returns the the stake associated
     * with the validation. Note that this function can be used for successful validation periods that have been explicitly
     * ended by calling {initializeEndValidation} or for validation periods that never began on the P-Chain due to the
     * {registrationExpiry} being reached.
     * @param setWeightMessageType Whether or not the message type is a SetValidatorWeight message, or a
     * SubnetValidatorRegistration message (with valid set to false). Both message types are valid for ending
     * a validation period.
     */
    function completeEndValidation(uint32 messageIndex, bool setWeightMessageType) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();

        // Get the Warp message.
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "ValidatorManager: Invalid warp message");

        bytes32 validationID;
        if (setWeightMessageType) {
            uint64 weight;
            (validationID,, weight) =
                ValidatorMessages.unpackSetSubnetValidatorWeightMessage(warpMessage.payload);
            require(weight == 0, "ValidatorManager: Weight must be 0");
        } else {
            bool validRegistration;
            (validationID, validRegistration) =
                ValidatorMessages.unpackSubnetValidatorRegistrationMessage(warpMessage.payload);
            require(!validRegistration, "ValidatorManager: Registration still valid");
        }

        // Note: The validation status is not necessarily PendingRemoved here. It could be PendingAdded if the
        // registration was never successfully registered on the P-Chain, or it could be Active if the validator
        // removed themselves directly on the P-Chain without going through the contract.
        Validator memory validator = $._validationPeriods[validationID];
        require(_msgSender() == validator.owner, "ValidatorManager: Sender not validator owner");

        ValidatorStatus endStatus;
        if (
            validator.status == ValidatorStatus.PendingRemoved
                || validator.status == ValidatorStatus.Active
        ) {
            // Remove the validator from the active validators mapping.
            delete $._activeValidators[validator.nodeID];
            endStatus = ValidatorStatus.Completed;
        } else {
            endStatus = ValidatorStatus.Invalidated;
        }

        $._validationPeriods[validationID] = validator;

        // Unlock the stake.

        // Calculate the reward for the validator.

        // Emit event.
        emit ValidationPeriodEnded(validationID);
    }

    /**
     * @notice Helper function to check if the stake amount to be added or removed would exceed the maximum stake churn
     * rate for the past hour. If the churn rate is exceeded, the function will revert. If the churn rate is not exceeded,
     * the function will update the churn tracker with the new amount.
     */
    function _checkAndUpdateChurnTracker(uint64 amount) private {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        if ($._maximumHourlyChurn == 0) {
            return;
        }

        ValidatorChurnPeriod memory churnTracker = $._churnTracker;
        uint256 currentTime = block.timestamp;
        if (currentTime - churnTracker.startedAt >= 1 hours) {
            churnTracker.churnAmount = amount;
            churnTracker.startedAt = currentTime;
        } else {
            churnTracker.churnAmount += amount;
        }

        uint8 churnPercentage = uint8((churnTracker.churnAmount * 100) / churnTracker.initialStake);
        require(
            churnPercentage <= $._maximumHourlyChurn,
            "ValidatorManager: Maximum hourly churn rate exceeded"
        );
        $._churnTracker = churnTracker;
    }
}
