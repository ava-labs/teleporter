// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    IPoSValidatorManager, Delegator, DelegatorStatus
} from "./interfaces/IPoSValidatorManager.sol";
import {
    PoSValidatorManagerSettings, PoSValidatorInfo
} from "./interfaces/IPoSValidatorManager.sol";
import {Validator} from "./interfaces/IValidatorManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";
import {
    Validator,
    ValidatorStatus,
    ValidatorRegistrationInput
} from "./interfaces/IValidatorManager.sol";
import {WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";

abstract contract PoSValidatorManager is
    IPoSValidatorManager,
    ValidatorManager,
    ReentrancyGuardUpgradeable
{
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.PoSValidatorManager
    struct PoSValidatorManagerStorage {
        /// @notice The minimum amount of stake required to be a validator.
        uint256 _minimumStakeAmount;
        /// @notice The maximum amount of stake allowed to be a validator.
        uint256 _maximumStakeAmount;
        /// @notice The minimum amount of time a validator must be staked for.
        uint64 _minimumStakeDuration;
        /// @notice The minimum delegation fee percentage, in basis points, required to delegate to a validator.
        uint16 _minimumDelegationFeeBips;
        /**
         * @notice A multiplier applied to validator's initial stake amount to determine
         * the maximum amount of stake a validator can have with delegations.
         *
         * Note: Setting this value to 1 would disable delegations to validators, since
         * the maximum stake would be equal to the initial stake.
         */
        uint64 _maximumStakeMultiplier;
        /// @notice The reward calculator for this validator manager.
        IRewardCalculator _rewardCalculator;
        /// @notice Maps the validation ID to its requirements.
        mapping(bytes32 validationID => PoSValidatorInfo) _posValidatorInfo;
        /// @notice Maps the delegation ID to the delegator information.
        mapping(bytes32 delegationID => Delegator) _delegatorStakes;
        /// @notice Maps the delegation ID to its pending staking rewards.
        mapping(bytes32 delegationID => uint256) _redeemableDelegatorRewards;
        /// @notice Maps the validation ID to its pending staking rewards.
        mapping(bytes32 validationID => uint256) _redeemableValidatorRewards;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoSValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Unit test for storage slot and update slot
    bytes32 private constant _POS_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00;

    uint8 public constant MAXIMUM_STAKE_MULTIPLIER_LIMIT = 10;

    uint16 public constant MAXIMUM_DELEGATION_FEE_BIPS = 10000;

    error InvalidDelegatorStatus();
    error InvalidNonce();
    error InvalidDelegationID();
    error ValidatorNotPoS();
    error MaxWeightExceeded();
    error InvalidDelegationFee();
    error InvalidStakeDuration();
    error InvalidStakeAmount();
    error InvalidStakeMultiplier();
    error ValidatorIneligibleForRewards();

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
        __ReentrancyGuard_init();
        __POS_Validator_Manager_init_unchained({
            minimumStakeAmount: settings.minimumStakeAmount,
            maximumStakeAmount: settings.maximumStakeAmount,
            minimumStakeDuration: settings.minimumStakeDuration,
            minimumDelegationFeeBips: settings.minimumDelegationFeeBips,
            maximumStakeMultiplier: settings.maximumStakeMultiplier,
            rewardCalculator: settings.rewardCalculator
        });
    }

    // solhint-disable-next-line func-name-mixedcase
    function __POS_Validator_Manager_init_unchained(
        uint256 minimumStakeAmount,
        uint256 maximumStakeAmount,
        uint64 minimumStakeDuration,
        uint16 minimumDelegationFeeBips,
        uint8 maximumStakeMultiplier,
        IRewardCalculator rewardCalculator
    ) internal onlyInitializing {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        if (minimumDelegationFeeBips == 0 || minimumDelegationFeeBips > MAXIMUM_DELEGATION_FEE_BIPS)
        {
            revert InvalidDelegationFee();
        }
        if (minimumStakeAmount > maximumStakeAmount) {
            revert InvalidStakeAmount();
        }
        if (maximumStakeMultiplier == 0 || maximumStakeMultiplier > MAXIMUM_STAKE_MULTIPLIER_LIMIT)
        {
            revert InvalidStakeMultiplier();
        }

        $._minimumStakeAmount = minimumStakeAmount;
        $._maximumStakeAmount = maximumStakeAmount;
        $._minimumStakeDuration = minimumStakeDuration;
        $._minimumDelegationFeeBips = minimumDelegationFeeBips;
        $._maximumStakeMultiplier = maximumStakeMultiplier;
        $._rewardCalculator = rewardCalculator;
    }

    function submitUptimeProof(bytes32 validationID, uint32 messageIndex) external {
        if (!_isPoSValidator(validationID)) {
            revert ValidatorNotPoS();
        }
        if (getValidator(validationID).status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus();
        }

        // Uptime proofs include the absolute number of seconds the validator has been active.
        _updateUptime(validationID, messageIndex);
    }

    function claimDelegationFees(bytes32 validationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        if (getValidator(validationID).status != ValidatorStatus.Completed) {
            revert InvalidValidatorStatus();
        }

        if ($._posValidatorInfo[validationID].owner != _msgSender()) {
            revert InvalidAddress();
        }

        uint256 rewards = $._redeemableValidatorRewards[validationID];
        delete $._redeemableValidatorRewards[validationID];
        _reward($._posValidatorInfo[validationID].owner, rewards);
    }

    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        if(!_initializeEndPoSValidation(validationID, includeUptimeProof, messageIndex)) {
            revert ValidatorIneligibleForRewards();
        }
    }

    function forceInitializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        _initializeEndPoSValidation(validationID, includeUptimeProof, messageIndex);
    }

    function _initializeEndPoSValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) internal returns (bool) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Validator memory validator = _initializeEndValidation(validationID);

        // Non-PoS validators are required to boostrap the network, but are not eligible for rewards.
        if (!_isPoSValidator(validationID)) {
            return true;
        }

        // PoS validations can only be ended by their owners.
        if ($._posValidatorInfo[validationID].owner != _msgSender()) {
            revert InvalidAddress();
        }

        // Check that minimum stake duration has passed.
        if (
            validator.endedAt
                < validator.startedAt + $._posValidatorInfo[validationID].minStakeDuration
        ) {
            revert InvalidStakeDuration();
        }

        // Uptime proofs include the absolute number of seconds the validator has been active.
        uint64 uptimeSeconds;
        if (includeUptimeProof) {
            uptimeSeconds = _updateUptime(validationID, messageIndex);
        } else {
            uptimeSeconds = $._posValidatorInfo[validationID].uptimeSeconds;
        }

        uint256 reward = $._rewardCalculator.calculateReward({
            stakeAmount: weightToValue(validator.startingWeight),
            validatorStartTime: validator.startedAt,
            stakingStartTime: validator.startedAt,
            stakingEndTime: validator.endedAt,
            uptimeSeconds: uptimeSeconds,
            initialSupply: 0,
            endSupply: 0
        });
        $._redeemableValidatorRewards[validationID] += reward;
        return (reward > 0);
    }

    function completeEndValidation(uint32 messageIndex) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        (bytes32 validationID, Validator memory validator) = _completeEndValidation(messageIndex);

        // Return now if this was originally a PoA validator that was later migrated to this PoS manager,
        // or the validator was part of the initial validator set.
        if (!_isPoSValidator(validationID)) {
            return;
        }

        address owner = $._posValidatorInfo[validationID].owner;
        // The validator can either be Completed or Invalidated here. We only grant rewards for Completed.
        if (validator.status == ValidatorStatus.Completed) {
            uint256 rewards = $._redeemableValidatorRewards[validationID];
            delete $._redeemableValidatorRewards[validationID];
            _reward(owner, rewards);
        }

        // We unlock the stake whether the validation period is completed or invalidated.
        _unlock(owner, weightToValue(validator.startingWeight));
    }

    // Helper function that extracts the uptime from a ValidationUptimeMessage Warp message
    // If the uptime is greater than the stored uptime, update the stored uptime
    function _updateUptime(bytes32 validationID, uint32 messageIndex) internal returns (uint64) {
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        if (!valid) {
            revert InvalidWarpMessage();
        }

        if (warpMessage.sourceChainID != WARP_MESSENGER.getBlockchainID()) {
            revert InvalidBlockchainID();
        }
        if (warpMessage.originSenderAddress != address(0)) {
            revert InvalidAddress();
        }

        (bytes32 uptimeValidationID, uint64 uptime) =
            ValidatorMessages.unpackValidationUptimeMessage(warpMessage.payload);
        if (validationID != uptimeValidationID) {
            revert InvalidValidationID();
        }

        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        if (uptime > $._posValidatorInfo[validationID].uptimeSeconds) {
            $._posValidatorInfo[validationID].uptimeSeconds = uptime;
            emit UptimeUpdated(validationID, uptime);
        } else {
            uptime = $._posValidatorInfo[validationID].uptimeSeconds;
        }

        return uptime;
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) internal virtual returns (bytes32) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        // Validate and save the validator requirements
        if (
            delegationFeeBips < $._minimumDelegationFeeBips
                || delegationFeeBips > MAXIMUM_DELEGATION_FEE_BIPS
        ) {
            revert InvalidDelegationFee();
        }

        if (minStakeDuration < $._minimumStakeDuration) {
            revert InvalidStakeDuration();
        }

        // Ensure the weight is within the valid range.
        if (stakeAmount < $._minimumStakeAmount || stakeAmount > $._maximumStakeAmount) {
            revert InvalidStakeAmount();
        }

        // Lock the stake in the contract.
        uint256 lockedValue = _lock(stakeAmount);

        uint64 weight = valueToWeight(lockedValue);
        bytes32 validationID = _initializeValidatorRegistration(registrationInput, weight);

        $._posValidatorInfo[validationID] = PoSValidatorInfo({
            owner: _msgSender(),
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration,
            uptimeSeconds: 0
        });
        return validationID;
    }

    function valueToWeight(uint256 value) public pure returns (uint64) {
        return uint64(value / 1e12);
    }

    function weightToValue(uint64 weight) public pure returns (uint256) {
        return uint256(weight) * 1e12;
    }

    function _lock(uint256 value) internal virtual returns (uint256);
    function _unlock(address to, uint256 value) internal virtual;

    function _initializeDelegatorRegistration(
        bytes32 validationID,
        address delegatorAddress,
        uint256 delegationAmount
    ) internal returns (bytes32) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        uint64 weight = valueToWeight(_lock(delegationAmount));

        // Ensure the validation period is active
        Validator memory validator = getValidator(validationID);
        // Check that the validation ID is a PoS validator
        if (!_isPoSValidator(validationID)) {
            revert ValidatorNotPoS();
        }
        if (validator.status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus();
        }

        // Update the validator weight
        uint64 newValidatorWeight = validator.weight + weight;
        if (newValidatorWeight > validator.startingWeight * $._maximumStakeMultiplier) {
            revert MaxWeightExceeded();
        }

        (uint64 nonce, bytes32 messageID) = _setValidatorWeight(validationID, newValidatorWeight);

        bytes32 delegationID = keccak256(abi.encodePacked(validationID, nonce));

        // Store the delegation information. Set the delegator status to pending added,
        // so that it can be properly started in the complete step, even if the delivered
        // nonce is greater than the nonce used to initialize registration.
        $._delegatorStakes[delegationID] = Delegator({
            status: DelegatorStatus.PendingAdded,
            owner: delegatorAddress,
            validationID: validationID,
            weight: weight,
            startedAt: 0,
            startingNonce: nonce,
            endingNonce: 0
        });

        emit DelegatorAdded({
            delegationID: delegationID,
            validationID: validationID,
            delegatorAddress: delegatorAddress,
            nonce: nonce,
            validatorWeight: newValidatorWeight,
            delegatorWeight: weight,
            setWeightMessageID: messageID
        });
        return delegationID;
    }

    function completeDelegatorRegistration(uint32 messageIndex, bytes32 delegationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        Validator memory validator = getValidator(delegator.validationID);

        // Ensure the delegator is pending added. Since anybody can call this function once
        // delegator registration has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        if ($._delegatorStakes[delegationID].status != DelegatorStatus.PendingAdded) {
            revert InvalidDelegatorStatus();
        }

        // In the case where the validator has initiated ending its validation period, we can no
        // longer stake, but the validator's weight may not have been removed from the p-chain yet,
        // so we have to wait for confirmation that the validator has been removed before returning the stake.
        if (validator.status == ValidatorStatus.PendingRemoved) {
            delegator.status = DelegatorStatus.PendingRemoved;
            delegator.startedAt = validator.endedAt;
            delegator.endingNonce = validator.messageNonce;

            // Write back the delegator
            $._delegatorStakes[delegationID] = delegator;

            emit DelegatorRemovalInitialized({
                delegationID: delegationID,
                validationID: delegator.validationID,
                endTime: validator.endedAt
            });

            return;
        }

        // In the case where the validator has completed its validation period, we can no
        // longer stake and should move our status directly to completed and return the stake.
        if (validator.status == ValidatorStatus.Completed) {
            _completeEndDelegation(delegationID);
            return;
        }

        // Unpack the Warp message
        (bytes32 validationID, uint64 nonce,) = ValidatorMessages
            .unpackSubnetValidatorWeightUpdateMessage(_getPChainWarpMessage(messageIndex).payload);

        if (delegator.validationID != validationID) {
            revert InvalidValidationID();
        }

        // The received nonce should be no greater than the highest sent nonce, and at least as high as
        // the delegation's starting nonce. This allows a weight update using a higher nonce
        // (which implicitly includes the delegation's weight update) to be used to complete delisting
        // for an earlier delegation. This is necessary because the P-Chain is only willing to sign the latest weight update.
        if (
            validator.messageNonce < nonce || $._delegatorStakes[delegationID].startingNonce > nonce
        ) {
            revert InvalidNonce();
        }

        // Update the delegation status
        $._delegatorStakes[delegationID].status = DelegatorStatus.Active;
        $._delegatorStakes[delegationID].startedAt = uint64(block.timestamp);

        emit DelegatorRegistered({
            delegationID: delegationID,
            validationID: validationID,
            nonce: nonce,
            startTime: block.timestamp
        });
    }

    function endDelegationCompletedValidator(bytes32 delegationID) external nonReentrant {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;
        Validator memory validator = getValidator(validationID);

        // Ensure the delegation is found
        if (delegator.status == DelegatorStatus.Unknown) {
            revert InvalidDelegationID();
        }

        // Ensure the validation is completed
        if (validator.status != ValidatorStatus.Completed) {
            revert InvalidValidatorStatus();
        }

        // If the validator completes before the delegator is added, let them exit from here.
        if (delegator.status == DelegatorStatus.PendingAdded) {
            _completeEndDelegation(delegationID);
            return;
        }

        // Calculate and set rewards if the delegator is active.
        if (delegator.status == DelegatorStatus.Active) {
            $._redeemableDelegatorRewards[delegationID] = $._rewardCalculator.calculateReward({
                stakeAmount: weightToValue(delegator.weight),
                validatorStartTime: validator.startedAt,
                stakingStartTime: delegator.startedAt,
                stakingEndTime: validator.endedAt,
                uptimeSeconds: $._posValidatorInfo[validationID].uptimeSeconds,
                initialSupply: 0,
                endSupply: 0
            });
        }

        _completeEndDelegation(delegationID);
    }

    function initializeEndDelegation(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;
        Validator memory validator = getValidator(validationID);

        // Ensure the delegator is active
        if (delegator.status != DelegatorStatus.Active) {
            revert InvalidDelegatorStatus();
        }
        // Only the delegation owner can end the delegation.
        if (delegator.owner != _msgSender()) {
            revert InvalidAddress();
        }

        // Set the delegator status to pending removed, so that it can be properly removed in
        // the complete step, even if the delivered nonce is greater than the nonce used to
        // initialize the removal.
        delegator.status = DelegatorStatus.PendingRemoved;

        uint64 validatorUptimeSeconds;
        uint64 delegationEndTime;
        if (validator.status == ValidatorStatus.Active) {
            if (includeUptimeProof) {
                // Uptime proofs include the absolute number of seconds the validator has been active.
                validatorUptimeSeconds = _updateUptime(validationID, messageIndex);
            }
            uint64 newValidatorWeight = validator.weight - delegator.weight;
            (delegator.endingNonce,) = _setValidatorWeight(validationID, newValidatorWeight);

            delegationEndTime = uint64(block.timestamp);
        } else {
            // If the validation period has already ended, we have saved the uptime.
            // Further, it is impossible to retrieve an uptime proof for an already ended validation,
            // so there's no need to check any uptime proof provided in this function call.
            validatorUptimeSeconds = $._posValidatorInfo[validationID].uptimeSeconds;

            delegator.endingNonce = validator.messageNonce;
            delegationEndTime = validator.endedAt;
        }

        $._redeemableDelegatorRewards[delegationID] = $._rewardCalculator.calculateReward({
            stakeAmount: weightToValue(delegator.weight),
            validatorStartTime: validator.startedAt,
            stakingStartTime: delegator.startedAt,
            stakingEndTime: delegationEndTime,
            uptimeSeconds: validatorUptimeSeconds,
            initialSupply: 0,
            endSupply: 0
        });

        $._delegatorStakes[delegationID] = delegator;

        emit DelegatorRemovalInitialized({
            delegationID: delegationID,
            validationID: validationID,
            endTime: delegationEndTime
        });
    }

    /**
     * @dev Resending the latest validator weight with the latest nonce is safe because all weight changes are
     * cumulative, so the latest weight change will always include the weight change for any added delegators.
     */
    function resendUpdateDelegation(bytes32 delegationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID];
        if (
            delegator.status != DelegatorStatus.PendingAdded
                && delegator.status != DelegatorStatus.PendingRemoved
        ) {
            revert InvalidDelegatorStatus();
        }

        Validator memory validator = getValidator(delegator.validationID);
        if (validator.messageNonce == 0) {
            revert InvalidDelegationID();
        }

        // Submit the message to the Warp precompile.
        WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packSetSubnetValidatorWeightMessage(
                delegator.validationID, validator.messageNonce, validator.weight
            )
        );
    }

    function completeEndDelegation(
        uint32 messageIndex,
        bytes32 delegationID
    ) external nonReentrant {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID];

        // Unpack the Warp message
        (bytes32 validationID, uint64 nonce,) = ValidatorMessages
            .unpackSubnetValidatorWeightUpdateMessage(_getPChainWarpMessage(messageIndex).payload);

        if (delegator.validationID != validationID) {
            revert InvalidValidationID();
        }

        // The received nonce should be at least as high as the delegation's ending nonce. This allows a weight
        // update using a higher nonce (which implicitly includes the delegation's weight update) to be used to
        // complete delisting for an earlier delegation. This is necessary because the P-Chain is only willing
        // to sign the latest weight update.
        if (delegator.endingNonce > nonce) {
            revert InvalidNonce();
        }

        // Ensure the delegator is pending removed. Since anybody can call this function once
        // end delegation has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        if (delegator.status != DelegatorStatus.PendingRemoved) {
            revert InvalidDelegatorStatus();
        }

        _completeEndDelegation(delegationID);
    }

    function _completeEndDelegation(bytes32 delegationID) internal {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;

        // Once this function completes, the delegation is completed so we can clear it from state now.
        delete $._delegatorStakes[delegationID];

        uint256 rewards = $._redeemableDelegatorRewards[delegationID];
        delete $._redeemableDelegatorRewards[delegationID];

        uint256 validatorFees =
            rewards * $._posValidatorInfo[validationID].delegationFeeBips / 10000;

        // Allocate the delegation fees to the validator.
        $._redeemableValidatorRewards[validationID] += validatorFees;

        // Reward the remaining tokens to the delegator.
        uint256 delegatorRewards = rewards - validatorFees;
        _reward(delegator.owner, delegatorRewards);

        // Unlock the delegator's stake.
        _unlock(delegator.owner, weightToValue(delegator.weight));

        emit DelegationEnded(delegationID, validationID, delegatorRewards, validatorFees);
    }

    /**
     * @dev This function must be implemented to mint rewards to validators and delegators.
     */
    function _reward(address account, uint256 amount) internal virtual;

    /**
     * @dev Return true if this is a PoS validator with locked stake. Returns false if this was originally a PoA
     * validator that was later migrated to this PoS manager, or the validator was part of the initial validator set.
     */
    function _isPoSValidator(bytes32 validationID) internal view returns (bool) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        return $._posValidatorInfo[validationID].owner != address(0);
    }
}
