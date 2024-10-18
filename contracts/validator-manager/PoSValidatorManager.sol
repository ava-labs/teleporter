// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManager} from "./ValidatorManager.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {
    Delegator,
    DelegatorStatus,
    IPoSValidatorManager,
    PoSValidatorInfo,
    PoSValidatorManagerSettings
} from "./interfaces/IPoSValidatorManager.sol";
import {
    Validator,
    ValidatorRegistrationInput,
    ValidatorStatus
} from "./interfaces/IValidatorManager.sol";
import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";
import {WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ReentrancyGuardUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/utils/ReentrancyGuardUpgradeable.sol";

/**
 * @dev Implementation of the {IPoSValidatorManager} interface.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter/blob/main/SECURITY.md
 */
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
        /// @notice The minimum amount of time in seconds a validator must be staked for. Must be at least {_churnPeriodSeconds}.
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
    bytes32 public constant POS_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00;

    uint8 public constant MAXIMUM_STAKE_MULTIPLIER_LIMIT = 10;

    uint16 public constant MAXIMUM_DELEGATION_FEE_BIPS = 10000;

    error InvalidDelegationFee(uint16 delegationFeeBips);
    error InvalidDelegationID(bytes32 delegationID);
    error InvalidDelegatorStatus(DelegatorStatus status);
    error InvalidNonce(uint64 nonce);
    error InvalidStakeAmount(uint256 stakeAmount);
    error InvalidMinStakeDuration(uint64 minStakeDuration);
    error InvalidStakeMultiplier(uint8 maximumStakeMultiplier);
    error MaxWeightExceeded(uint64 newValidatorWeight);
    error MinStakeDurationNotPassed(uint64 endTime);
    error UnauthorizedOwner(address sender);
    error ValidatorNotPoS(bytes32 validationID);
    error ValidatorIneligibleForRewards(bytes32 validationID);
    error DelegatorIneligibleForRewards(bytes32 delegationID);

    // solhint-disable ordering
    function _getPoSValidatorManagerStorage()
        private
        pure
        returns (PoSValidatorManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := POS_VALIDATOR_MANAGER_STORAGE_LOCATION
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
            revert InvalidDelegationFee(minimumDelegationFeeBips);
        }
        if (minimumStakeAmount > maximumStakeAmount) {
            revert InvalidStakeAmount(minimumStakeAmount);
        }
        if (maximumStakeMultiplier == 0 || maximumStakeMultiplier > MAXIMUM_STAKE_MULTIPLIER_LIMIT)
        {
            revert InvalidStakeMultiplier(maximumStakeMultiplier);
        }
        // Minimum stake duration should be at least one churn period in order to prevent churn tracker abuse.
        if (minimumStakeDuration < _getChurnPeriodSeconds()) {
            revert InvalidMinStakeDuration(minimumStakeDuration);
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
            revert ValidatorNotPoS(validationID);
        }
        ValidatorStatus status = getValidator(validationID).status;
        if (status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus(status);
        }

        // Uptime proofs include the absolute number of seconds the validator has been active.
        _updateUptime(validationID, messageIndex);
    }

    function claimDelegationFees(bytes32 validationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        ValidatorStatus status = getValidator(validationID).status;
        if (status != ValidatorStatus.Completed) {
            revert InvalidValidatorStatus(status);
        }

        if ($._posValidatorInfo[validationID].owner != _msgSender()) {
            revert UnauthorizedOwner(_msgSender());
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
        if (!_initializeEndPoSValidation(validationID, includeUptimeProof, messageIndex)) {
            revert ValidatorIneligibleForRewards(validationID);
        }
    }

    function forceInitializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        // Ignore the return value here to force end validation, regardless of possible missed rewards
        _initializeEndPoSValidation(validationID, includeUptimeProof, messageIndex);
    }

    // Helper function that initializes the end of a PoS validation period.
    // Returns false if it is possible for the validator to claim rewards, but it is not eligible.
    // Returns true otherwise.
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
            revert UnauthorizedOwner(_msgSender());
        }

        // Check that minimum stake duration has passed.
        if (
            validator.endedAt
                < validator.startedAt + $._posValidatorInfo[validationID].minStakeDuration
        ) {
            revert MinStakeDurationNotPassed(validator.endedAt);
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

    function completeEndValidation(uint32 messageIndex) external nonReentrant {
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
            revert InvalidWarpSourceChainID(warpMessage.sourceChainID);
        }
        if (warpMessage.originSenderAddress != address(0)) {
            revert InvalidWarpOriginSenderAddress(warpMessage.originSenderAddress);
        }

        (bytes32 uptimeValidationID, uint64 uptime) =
            ValidatorMessages.unpackValidationUptimeMessage(warpMessage.payload);
        if (validationID != uptimeValidationID) {
            revert InvalidValidationID(validationID);
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
            revert InvalidDelegationFee(delegationFeeBips);
        }

        if (minStakeDuration < $._minimumStakeDuration) {
            revert InvalidMinStakeDuration(minStakeDuration);
        }

        // Ensure the weight is within the valid range.
        if (stakeAmount < $._minimumStakeAmount || stakeAmount > $._maximumStakeAmount) {
            revert InvalidStakeAmount(stakeAmount);
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
            revert ValidatorNotPoS(validationID);
        }
        if (validator.status != ValidatorStatus.Active) {
            revert InvalidValidatorStatus(validator.status);
        }

        // Update the validator weight
        uint64 newValidatorWeight = validator.weight + weight;
        if (newValidatorWeight > validator.startingWeight * $._maximumStakeMultiplier) {
            revert MaxWeightExceeded(newValidatorWeight);
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
        bytes32 validationID = delegator.validationID;
        Validator memory validator = getValidator(validationID);

        // Ensure the delegator is pending added. Since anybody can call this function once
        // delegator registration has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        if (delegator.status != DelegatorStatus.PendingAdded) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        // In the case where the validator has completed its validation period, we can no
        // longer stake and should move our status directly to completed and return the stake.
        if (validator.status == ValidatorStatus.Completed) {
            return _completeEndDelegation(delegationID);
        }

        // Unpack the Warp message
        (bytes32 messageValidationID, uint64 nonce,) = ValidatorMessages
            .unpackSubnetValidatorWeightMessage(_getPChainWarpMessage(messageIndex).payload);

        if (validationID != messageValidationID) {
            revert InvalidValidationID(delegator.validationID);
        }

        // The received nonce should be no greater than the highest sent nonce, and at least as high as
        // the delegation's starting nonce. This allows a weight update using a higher nonce
        // (which implicitly includes the delegation's weight update) to be used to complete delisting
        // for an earlier delegation. This is necessary because the P-Chain is only willing to sign the latest weight update.
        if (validator.messageNonce < nonce || delegator.startingNonce > nonce) {
            revert InvalidNonce(nonce);
        }

        // Update the delegation status
        $._delegatorStakes[delegationID].status = DelegatorStatus.Active;
        $._delegatorStakes[delegationID].startedAt = uint64(block.timestamp);

        emit DelegatorRegistered({
            delegationID: delegationID,
            validationID: validationID,
            startTime: uint64(block.timestamp)
        });
    }

    function initializeEndDelegation(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        if (!_initializeEndDelegation(delegationID, includeUptimeProof, messageIndex)) {
            revert DelegatorIneligibleForRewards(delegationID);
        }
    }

    function forceInitializeEndDelegation(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        // Ignore the return value here to force end delegation, regardless of possible missed rewards
        _initializeEndDelegation(delegationID, includeUptimeProof, messageIndex);
    }

    // Helper function that initializes the end of a PoS delegation period.
    // Returns false if it is possible for the delegator to claim rewards, but it is not eligible.
    // Returns true otherwise.
    function _initializeEndDelegation(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) internal returns (bool) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Delegator memory delegator = $._delegatorStakes[delegationID];
        bytes32 validationID = delegator.validationID;
        Validator memory validator = getValidator(validationID);

        // Ensure the delegator is active
        if (delegator.status != DelegatorStatus.Active) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        // Only the delegation owner or parent validator can end the delegation.
        if (delegator.owner != _msgSender()) {
            // Validators can only remove delegations after the minimum stake duration has passed.
            if ($._posValidatorInfo[validationID].owner != _msgSender()) {
                revert UnauthorizedOwner(_msgSender());
            }

            if (
                block.timestamp
                    < validator.startedAt + $._posValidatorInfo[validationID].minStakeDuration
            ) {
                revert MinStakeDurationNotPassed(uint64(block.timestamp));
            }
        }

        if (validator.status == ValidatorStatus.Active) {
            // Check that minimum stake duration has passed.
            if (block.timestamp < delegator.startedAt + $._minimumStakeDuration) {
                revert MinStakeDurationNotPassed(uint64(block.timestamp));
            }

            if (includeUptimeProof) {
                // Uptime proofs include the absolute number of seconds the validator has been active.
                _updateUptime(validationID, messageIndex);
            }

            // Set the delegator status to pending removed, so that it can be properly removed in
            // the complete step, even if the delivered nonce is greater than the nonce used to
            // initialize the removal.
            $._delegatorStakes[delegationID].status = DelegatorStatus.PendingRemoved;

            ($._delegatorStakes[delegationID].endingNonce,) =
                _setValidatorWeight(validationID, validator.weight - delegator.weight);

            uint256 reward = _calculateDelegationReward(delegator);
            $._redeemableDelegatorRewards[delegationID] = reward;

            emit DelegatorRemovalInitialized({
                delegationID: delegationID,
                validationID: validationID
            });
            return (reward > 0);
        } else if (validator.status == ValidatorStatus.Completed) {
            $._redeemableDelegatorRewards[delegationID] = _calculateDelegationReward(delegator);

            _completeEndDelegation(delegationID);
            // If the validator has completed, then no further uptimes may be submitted, so we always
            // end the delegation
            return true;
        } else {
            revert InvalidValidatorStatus(validator.status);
        }
    }

    /// @dev Calculates the reward owed to the delegator based on the state of the delegator and its corresponding validator.
    function _calculateDelegationReward(Delegator memory delegator)
        private
        view
        returns (uint256)
    {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Validator memory validator = getValidator(delegator.validationID);

        uint64 delegationEndTime;
        if (
            validator.status == ValidatorStatus.PendingRemoved
                || validator.status == ValidatorStatus.Completed
        ) {
            delegationEndTime = validator.endedAt;
        } else if (validator.status == ValidatorStatus.Active) {
            delegationEndTime = uint64(block.timestamp);
        } else {
            revert InvalidValidatorStatus(validator.status);
        }

        // Only give rewards in the case that the delegation started before the validator exited.
        if (delegationEndTime <= delegator.startedAt) {
            return 0;
        }

        return $._rewardCalculator.calculateReward({
            stakeAmount: weightToValue(delegator.weight),
            validatorStartTime: validator.startedAt,
            stakingStartTime: delegator.startedAt,
            stakingEndTime: delegationEndTime,
            uptimeSeconds: $._posValidatorInfo[delegator.validationID].uptimeSeconds,
            initialSupply: 0,
            endSupply: 0
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
            revert InvalidDelegatorStatus(delegator.status);
        }

        Validator memory validator = getValidator(delegator.validationID);
        if (validator.messageNonce == 0) {
            revert InvalidDelegationID(delegationID);
        }

        // Submit the message to the Warp precompile.
        WARP_MESSENGER.sendWarpMessage(
            ValidatorMessages.packSubnetValidatorWeightMessage(
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

        // Ensure the delegator is pending removed. Since anybody can call this function once
        // end delegation has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        if (delegator.status != DelegatorStatus.PendingRemoved) {
            revert InvalidDelegatorStatus(delegator.status);
        }

        if (getValidator(delegator.validationID).status != ValidatorStatus.Completed) {
            // Unpack the Warp message
            WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
            (bytes32 validationID, uint64 nonce,) =
                ValidatorMessages.unpackSubnetValidatorWeightMessage(warpMessage.payload);

            if (delegator.validationID != validationID) {
                revert InvalidValidationID(validationID);
            }

            // The received nonce should be at least as high as the delegation's ending nonce. This allows a weight
            // update using a higher nonce (which implicitly includes the delegation's weight update) to be used to
            // complete delisting for an earlier delegation. This is necessary because the P-Chain is only willing
            // to sign the latest weight update.
            if (delegator.endingNonce > nonce) {
                revert InvalidNonce(nonce);
            }
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

        uint256 validatorFees;
        uint256 delegatorRewards;
        if (rewards > 0) {
            validatorFees = (rewards * $._posValidatorInfo[validationID].delegationFeeBips) / 10000;

            // Allocate the delegation fees to the validator.
            $._redeemableValidatorRewards[validationID] += validatorFees;

            // Reward the remaining tokens to the delegator.
            delegatorRewards = rewards - validatorFees;
            _reward(delegator.owner, delegatorRewards);
        }

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
