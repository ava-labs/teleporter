// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IStakingManager} from "./interfaces/IStakingManager.sol";
import {
    WarpMessage,
    IWarpMessenger
} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";
import {StakingMessages} from "./StakingMessages.sol";
import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";
import {ContextUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ContextUpgradeable.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

abstract contract StakingManager is
    Initializable,
    ContextUpgradeable,
    ReentrancyGuardUpgradeable,
    IStakingManager
{
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
        uint64 weight;
        uint64 startedAt;
        uint64 endedAt;
        uint64 uptimeSeconds;
        address owner;
        bool rewarded;
        uint64 messageNonce;
    }

    struct ValidatorChurnPeriod {
        uint256 startedAt;
        uint64 initialStake;
        uint64 churnAmount;
    }

    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.StakingManager
    struct StakingManagerStorage {
        bytes32 _pChainBlockchainID;
        bytes32 _subnetID;
        uint256 _minimumStakeAmount;
        uint256 _maximumStakeAmount;
        uint64 _minimumStakeDuration;
        IRewardCalculator _rewardCalculator;
        uint8 _maximumHourlyChurn;
        uint64 _remainingInitialStake;
        ValidatorChurnPeriod _churnTracker;
        // Maps the validationID to the registration message such that the message can be re-sent if needed.
        mapping(bytes32 => bytes) _pendingRegisterValidationMessages;
        // Maps the validationID to the validator information.
        mapping(bytes32 => Validator) _validationPeriods;
        // Maps the nodeID to the validationID for active validation periods.
        mapping(bytes32 => bytes32) _activeValidators;
    }

    // solhint-enable private-vars-leading-underscore
    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.StakingManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Update to correct storage slot
    bytes32 private constant _STAKING_MANAGER_STORAGE_LOCATION =
        0x8568826440873e37a96cb0aab773b28d8154d963d2f0e41bd9b5c15f63625f91;

    // solhint-disable ordering
    function _getStakingManagerStorage() private pure returns (StakingManagerStorage storage $) {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := _STAKING_MANAGER_STORAGE_LOCATION
        }
    }

    struct InitialStakerInfo {
        StakingMessages.ValidationInfo validationInfo;
        address owner;
    }

    struct StakingManagerSettings {
        bytes32 pChainBlockchainID;
        bytes32 subnetID;
        uint256 minimumStakeAmount;
        uint256 maximumStakeAmount;
        uint64 minimumStakeDuration;
        uint8 maximumHourlyChurn;
        InitialStakerInfo[] initialStakers;
        IRewardCalculator rewardCalculator;
    }

    /**
     * @notice Warp precompile used for sending and receiving Warp messages.
     */
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    function initialize(StakingManagerSettings calldata settings) public initializer {
        __StakingManager_init(settings);
    }

    function __StakingManager_init(StakingManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        __ReentrancyGuard_init();
        __Context_init();
        __StakingManager_init_unchained(settings);
    }

    function __StakingManager_init_unchained(StakingManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        $._pChainBlockchainID = settings.pChainBlockchainID;
        $._subnetID = settings.subnetID;
        $._minimumStakeAmount = settings.minimumStakeAmount;
        $._maximumStakeAmount = settings.maximumStakeAmount;
        $._minimumStakeDuration = settings.minimumStakeDuration;
        $._maximumHourlyChurn = settings.maximumHourlyChurn;
        // Add each of the initial stakers as validators
        uint64 initialStake;
        for (uint256 i; i < settings.initialStakers.length; ++i) {
            (bytes32 validationID,) =
                StakingMessages.packValidationInfo(settings.initialStakers[i].validationInfo);
            $._validationPeriods[validationID] = Validator({
                status: ValidatorStatus.Active,
                nodeID: settings.initialStakers[i].validationInfo.nodeID,
                weight: settings.initialStakers[i].validationInfo.weight,
                startedAt: uint64(block.timestamp),
                endedAt: 0,
                uptimeSeconds: 0,
                owner: settings.initialStakers[i].owner,
                rewarded: false,
                messageNonce: 0
            });
            initialStake += settings.initialStakers[i].validationInfo.weight;
        }
        $._remainingInitialStake = initialStake;
        $._rewardCalculator = settings.rewardCalculator;
    }

    /**
     * @notice Modifier to ensure that the initial stake has been provided.
     */
    modifier onlyWhenInitialStakeProvided() {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        require($._remainingInitialStake > 0, "StakingManager: Initial stake not provided");
        _;
    }

    /**
     * @notice Called to provide initial stake amount for original validators added prior to the contract's initialization.
     */
    function provideInitialStake() external payable {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        uint64 remainingInitialStake = $._remainingInitialStake;
        require(
            msg.value <= remainingInitialStake,
            "StakingManager: Provided stake exceeds remaining initial stake"
        );
        $._remainingInitialStake = remainingInitialStake - uint64(msg.value);
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
        uint256 value,
        uint64 registrationExpiry,
        bytes memory signature
    ) internal nonReentrant onlyWhenInitialStakeProvided returns (bytes32) {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        // Ensure the registration expiry is in a valid range.
        require(
            registrationExpiry > block.timestamp && block.timestamp + 2 days > registrationExpiry,
            "StakingManager: Invalid registration expiry"
        );

        // Ensure the nodeID is not the zero address, and is not already an active validator.
        require(nodeID != bytes32(0), "StakingManager: Invalid node ID");
        require($._activeValidators[nodeID] == bytes32(0), "StakingManager: Node ID already active");

        // Ensure the signature is the proper length. The EVM does not provide an Ed25519 precompile to
        // validate the signature, but the P-Chain will validate the signature. If the signature is invalid,
        // the P-Chain will reject the registration, and the stake can be returned to the staker after the registration
        // expiry has passed.
        require(signature.length == 64, "StakingManager: Invalid signature length");

        // Lock the stake in the contract.
        uint256 lockedValue = _lock(value);

        // Ensure the stake churn doesn't exceed the maximum churn rate.
        uint64 weight = valueToWeight(lockedValue);
        _checkAndUpdateChurnTracker(weight);

        // Ensure the weight is within the valid range.
        require(
            weight >= $._minimumStakeAmount && weight <= $._maximumStakeAmount,
            "StakingManager: Invalid stake amount"
        );

        (bytes32 validationID, bytes memory registerSubnetValidatorMessage) = StakingMessages
            .packRegisterSubnetValidatorMessage(
            StakingMessages.ValidationInfo({
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
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        require(
            $._pendingRegisterValidationMessages[validationID].length > 0
                && $._validationPeriods[validationID].status == ValidatorStatus.PendingAdded,
            "StakingManager: Invalid validation ID"
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
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "StakingManager: Invalid warp message");

        require(
            warpMessage.sourceChainID == $._pChainBlockchainID,
            "StakingManager: Invalid source chain ID"
        );
        require(
            warpMessage.originSenderAddress == address(0),
            "StakingManager: Invalid origin sender address"
        );

        (bytes32 validationID, bool validRegistration) =
            StakingMessages.unpackSubnetValidatorRegistrationMessage(warpMessage.payload);
        require(validRegistration, "StakingManager: Registration not valid");
        require(
            $._pendingRegisterValidationMessages[validationID].length > 0
                && $._validationPeriods[validationID].status == ValidatorStatus.PendingAdded,
            "StakingManager: Invalid validation ID"
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
    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external onlyWhenInitialStakeProvided {
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        // Ensure the validation period is active.
        Validator memory validator = $._validationPeriods[validationID];
        require(validator.status == ValidatorStatus.Active, "StakingManager: Validator not active");
        require(_msgSender() == validator.owner, "StakingManager: Sender not validator owner");

        // Check that removing this validator would not exceed the maximum churn rate.
        _checkAndUpdateChurnTracker(validator.weight);

        // Update the validator status to pending removal.
        // They are not removed from the active validators mapping until the P-Chain acknowledges the removal.
        validator.status = ValidatorStatus.PendingRemoved;

        // Set the end time of the validation period, since it is no longer known to be an active validator
        // on the P-Chain.
        validator.endedAt = uint64(block.timestamp);

        uint64 uptimeSeconds;
        if (includeUptimeProof) {
            (WarpMessage memory warpMessage, bool valid) =
                WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
            require(valid, "StakingManager: Invalid warp message");

            require(
                warpMessage.sourceChainID == WARP_MESSENGER.getBlockchainID(),
                "StakingManager: Invalid source chain ID"
            );
            require(
                warpMessage.originSenderAddress == address(0),
                "StakingManager: Invalid origin sender address"
            );

            (bytes32 uptimeValidationID, uint64 uptime) =
                StakingMessages.unpackValidationUptimeMessage(warpMessage.payload);
            require(
                validationID == uptimeValidationID, "StakingManager: Invalid uptime validation ID"
            );
            uptimeSeconds = uptime;
        }
        validator.uptimeSeconds = uptimeSeconds;

        // Save the validator updates.
        // TODO: Optimize storage writes here (probably don't need to write the whole value).
        $._validationPeriods[validationID] = validator;

        // Submit the message to the Warp precompile.
        bytes memory setValidatorWeightPayload = StakingMessages.packSetSubnetValidatorWeightMessage(
            validationID, validator.messageNonce, 0
        );
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
    // solhint-disable-next-line no-empty-blocks
    function resendEndValidatorMessage(bytes32 validationID) external {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        Validator memory validator = $._validationPeriods[validationID];

        require(
            validator.status == ValidatorStatus.PendingRemoved,
            "StakingManager: Validator not pending removal"
        );

        bytes memory setValidatorWeightPayload = StakingMessages.packSetSubnetValidatorWeightMessage(
            validationID, validator.messageNonce, 0
        );
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
        StakingManagerStorage storage $ = _getStakingManagerStorage();

        // Get the Warp message.
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "StakingManager: Invalid warp message");

        bytes32 validationID;
        if (setWeightMessageType) {
            uint64 weight;
            (validationID,, weight) =
                StakingMessages.unpackSetSubnetValidatorWeightMessage(warpMessage.payload);
            require(weight == 0, "StakingManager: Weight must be 0");
        } else {
            bool validRegistration;
            (validationID, validRegistration) =
                StakingMessages.unpackSubnetValidatorRegistrationMessage(warpMessage.payload);
            require(!validRegistration, "StakingManager: Registration still valid");
        }

        // Note: The validation status is not necessarily PendingRemoved here. It could be PendingAdded if the
        // registration was never successfully registered on the P-Chain, or it could be Active if the validator
        // removed themselves directly on the P-Chain without going through the contract.
        Validator memory validator = $._validationPeriods[validationID];
        require(_msgSender() == validator.owner, "StakingManager: Sender not validator owner");

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
    }

    /**
     * @notice Helper function to check if the stake amount to be added or removed would exceed the maximum stake churn
     * rate for the past hour. If the churn rate is exceeded, the function will revert. If the churn rate is not exceeded,
     * the function will update the churn tracker with the new amount.
     */
    function _checkAndUpdateChurnTracker(uint64 amount) private {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        ValidatorChurnPeriod storage churnTracker = $._churnTracker;
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
            "StakingManager: Maximum hourly churn rate exceeded"
        );
        $._churnTracker = churnTracker;
    }

    function valueToWeight(uint256 value) public pure returns (uint64) {
        return uint64(value / 1e12);
    }

    function weightToValue(uint64 weight) public pure returns (uint256) {
        return uint256(weight) * 1e12;
    }

    function _lock(uint256 value) internal virtual returns (uint256);
    function _unlock(uint256 value, address to) internal virtual;
}
