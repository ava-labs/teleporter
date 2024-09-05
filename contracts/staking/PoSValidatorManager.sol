// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    IPoSValidatorManager, Delegator, DelegatorStatus
} from "./interfaces/IPoSValidatorManager.sol";
import {PoSValidatorManagerSettings} from "./interfaces/IPoSValidatorManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";
import {Validator, ValidatorStatus} from "./interfaces/IValidatorManager.sol";
import {WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";

abstract contract PoSValidatorManager is IPoSValidatorManager, ValidatorManager {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.PoSValidatorManager
    struct PoSValidatorManagerStorage {
        /// @notice The minimum amount of stake required to be a validator.
        uint256 _minimumStakeAmount;
        /// @notice The maximum amount of stake allowed to be a validator.
        uint256 _maximumStakeAmount;
        /// @notice The minimum amount of time a validator must be staked for.
        uint64 _minimumStakeDuration;
        /// @notice The reward calculator for this validator manager.
        IRewardCalculator _rewardCalculator;
        /// @notice Maps the validationID to a mapping of delegator address to delegator information.
        mapping(bytes32 delegationID => Delegator) _delegatorStakes;
        /// @notice Maps the validationID to a mapping of delegator address to pending register delegator messages.
        mapping(bytes32 delegationID => bytes) _pendingRegisterDelegatorMessages;
        /// @notice Maps the validationID to a mapping of delegator address to pending end delegator messages.
        mapping(bytes32 delegationID => bytes) _pendingEndDelegatorMessages;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoSValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Unit test for storage slot and update slot
    bytes32 private constant _POS_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00;

    // solhint-disable ordering
    function _getPoSValidatorManagerStorage()
        private
        pure
        returns (PoSValidatorManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := _POS_VALIDATOR_MANAGER_STORAGE_LOCATION
        }
    }

    // solhint-disable-next-line func-name-mixedcase
    function __POS_Validator_Manager_init(PoSValidatorManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        __ValidatorManager_init(settings.baseSettings);
        __POS_Validator_Manager_init_unchained(
            settings.minimumStakeAmount,
            settings.maximumStakeAmount,
            settings.minimumStakeDuration,
            settings.rewardCalculator
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __POS_Validator_Manager_init_unchained(
        uint256 minimumStakeAmount,
        uint256 maximumStakeAmount,
        uint64 minimumStakeDuration,
        IRewardCalculator rewardCalculator
    ) internal onlyInitializing {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        $._minimumStakeAmount = minimumStakeAmount;
        $._maximumStakeAmount = maximumStakeAmount;
        $._minimumStakeDuration = minimumStakeDuration;
        $._rewardCalculator = rewardCalculator;
    }

    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        if (includeUptimeProof) {
            _getUptime(validationID, messageIndex);
        }
        // TODO: Calculate the reward for the validator, but do not unlock it

        _initializeEndValidation(validationID);
    }

    function _getUptime(bytes32 validationID, uint32 messageIndex) internal view returns (uint64) {
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        require(valid, "PoSValidatorManager: invalid warp message");

        require(
            warpMessage.sourceChainID == WARP_MESSENGER.getBlockchainID(),
            "PoSValidatorManager: invalid source chain ID"
        );
        require(
            warpMessage.originSenderAddress == address(0),
            "PoSValidatorManager: invalid origin sender address"
        );

        (bytes32 uptimeValidationID, uint64 uptime) =
            ValidatorMessages.unpackValidationUptimeMessage(warpMessage.payload);
        require(
            validationID == uptimeValidationID, "PoSValidatorManager: invalid uptime validation ID"
        );

        return uptime;
    }

    function _processStake(uint256 stakeAmount) internal virtual returns (uint64) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        // Lock the stake in the contract.
        uint256 lockedValue = _lock(stakeAmount);
        uint64 weight = valueToWeight(lockedValue);

        // Ensure the weight is within the valid range.
        require(
            weight >= $._minimumStakeAmount && weight <= $._maximumStakeAmount,
            "PoSValidatorManager: invalid stake amount"
        );
        return weight;
    }

    function valueToWeight(uint256 value) public pure returns (uint64) {
        return uint64(value / 1e12);
    }

    function weightToValue(uint64 weight) public pure returns (uint256) {
        return uint256(weight) * 1e12;
    }

    function _lock(uint256 value) internal virtual returns (uint256);
    function _unlock(uint256 value, address to) internal virtual;

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegator,
        uint256 delegationAmount
    ) internal nonReentrant returns (bytes32) {
        uint64 weight = valueToWeight(_lock(delegationAmount));
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Ensure the validation period is active
        Validator memory validator = _getValidator(validationID);
        require(
            validator.status == ValidatorStatus.Active, "PoSValidatorManager: validator not active"
        );
        // Update the validator weight
        validator.weight += weight;
        _setValidator(validationID, validator);

        uint64 nonce = _getAndIncrementNonce(validationID);
        bytes32 delegationID = sha256(abi.encodePacked(validationID, delegator, nonce));

        // Ensure the delegator is not already registered
        require(
            $._delegatorStakes[delegationID].weight == 0,
            "PoSValidatorManager: delegator already registered"
        );

        _checkAndUpdateChurnTracker(weight);

        // Submit the message to the Warp precompile.
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, nonce, validator.weight);
        $._pendingRegisterDelegatorMessages[delegationID] = setValidatorWeightPayload;
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);

        // Store the delegator information
        $._delegatorStakes[delegationID] = Delegator({
            status: DelegatorStatus.PendingAdded,
            owner: delegator,
            validationID: validationID,
            weight: weight,
            startedAt: 0,
            endedAt: 0,
            startingNonce: nonce,
            endingNonce: 0
        });

        emit DelegatorAdded({
            validationID: validationID,
            setWeightMessageID: messageID,
            delegationID: delegationID,
            delegatorWeight: weight,
            validatorWeight: validator.weight,
            nonce: nonce
        });
        return delegationID;
    }

    function resendDelegatorRegistration(bytes32 delegationID) external {
        _checkPendingRegisterDelegatorMessages(delegationID);
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        // Submit the message to the Warp precompile.
        WARP_MESSENGER.sendWarpMessage($._pendingRegisterDelegatorMessages[delegationID]);
    }

    function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Unpack the Warp message
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce,) =
            ValidatorMessages.unpackSubnetValidatorWeightUpdateMessage(warpMessage.payload);
        _checkPendingRegisterDelegatorMessages(delegationID);
        delete $._pendingRegisterDelegatorMessages[delegationID];

        Validator memory validator = _getValidator(validationID);

        // The received nonce should be no greater than the highest sent nonce
        require(validator.messageNonce >= nonce, "PoSValidatorManager: invalid nonce");
        // It should also be greater than or equal to the delegator's starting nonce
        require(
            $._delegatorStakes[delegationID].startingNonce <= nonce,
            "PoSValidatorManager: nonce does not match"
        );

        require(
            $._delegatorStakes[delegationID].status == DelegatorStatus.PendingAdded,
            "PoSValidatorManager: delegator not pending added"
        );
        // Update the delegator status
        $._delegatorStakes[delegationID].status = DelegatorStatus.Active;
        $._delegatorStakes[delegationID].startedAt = uint64(block.timestamp);
        emit DelegatorRegistered({
            validationID: validationID,
            delegationID: delegationID,
            nonce: nonce,
            startTime: block.timestamp
        });
    }

    function initializeEndDelegation(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        bytes32 validationID = $._delegatorStakes[delegationID].validationID;

        uint64 uptime;
        if (includeUptimeProof) {
            uptime = _getUptime(validationID, messageIndex);
        }

        // TODO: Calculate the delegator's reward, but do not unlock it

        // Ensure the delegator is active
        Delegator memory delegator = $._delegatorStakes[delegationID];
        require(
            delegator.status == DelegatorStatus.Active, "PoSValidatorManager: delegator not active"
        );
        uint64 nonce = _getAndIncrementNonce(validationID);
        delegator.status = DelegatorStatus.PendingRemoved;
        delegator.endedAt = uint64(block.timestamp);
        delegator.endingNonce = nonce;

        $._delegatorStakes[delegationID] = delegator;

        Validator memory validator = _getValidator(validationID);
        require(validator.weight > delegator.weight, "PoSValidatorManager: Invalid weight");
        validator.weight -= delegator.weight;
        _setValidator(validationID, validator);

        // Submit the message to the Warp precompile.
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, nonce, validator.weight);
        $._pendingEndDelegatorMessages[delegationID] = setValidatorWeightPayload;
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);

        emit DelegatorRemovalInitialized({
            validationID: validationID,
            setWeightMessageID: messageID,
            delegationID: delegationID,
            validatorWeight: validator.weight,
            nonce: nonce,
            endTime: block.timestamp
        });
    }

    function resendEndDelegation(bytes32 delegationID) external {
        _checkPendingEndDelegatorMessage(delegationID);
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        WARP_MESSENGER.sendWarpMessage($._pendingEndDelegatorMessages[delegationID]);
    }

    function completeEndDelegation(uint32 messageIndex, bytes32 delegationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Unpack the Warp message
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce,) =
            ValidatorMessages.unpackSubnetValidatorWeightUpdateMessage(warpMessage.payload);
        _checkPendingEndDelegatorMessage(delegationID);
        delete $._pendingEndDelegatorMessages[delegationID];

        Validator memory validator = _getValidator(validationID);
        // The received nonce should be no greater than the highest sent nonce
        require(validator.messageNonce >= nonce, "PoSValidatorManager: invalid nonce");
        // It should also be greater than or equal to the delegator's ending nonce
        require(
            $._delegatorStakes[delegationID].endingNonce <= nonce,
            "PoSValidatorManager: nonce does not match"
        );

        require(
            $._delegatorStakes[delegationID].status == DelegatorStatus.PendingRemoved,
            "PoSValidatorManager: delegator not pending added"
        );

        // Update the delegator status
        $._delegatorStakes[delegationID].status = DelegatorStatus.Completed;

        // TODO: Unlock the delegator's stake and their reward

        emit DelegationEnded(validationID, delegationID, nonce);
    }

    function _checkPendingEndDelegatorMessage(bytes32 delegationID) private view {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(
            $._pendingEndDelegatorMessages[delegationID].length > 0
                && $._delegatorStakes[delegationID].status == DelegatorStatus.PendingRemoved,
            "PoSValidatorManager: delegator removal not pending"
        );
    }

    function _checkPendingRegisterDelegatorMessages(bytes32 delegationID) private view {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(
            $._pendingRegisterDelegatorMessages[delegationID].length > 0
                && $._delegatorStakes[delegationID].status == DelegatorStatus.PendingAdded,
            "PoSValidatorManager: delegator registration not pending"
        );
    }
}
