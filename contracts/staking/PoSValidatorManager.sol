// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {
    IPoSValidatorManager, Delegator, DelegatorStatus
} from "./interfaces/IPoSValidatorManager.sol";
import {
    PoSValidatorManagerSettings,
    PoSValidatorRequirements
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
        uint256 _minimumDelegationFeeBips;
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
        mapping(bytes32 validationID => PoSValidatorRequirements) _validatorRequirements;
        /// @notice Maps the delegation ID to the delegator information.
        mapping(bytes32 delegationID => Delegator) _delegatorStakes;
        /// @notice Maps the delegation ID to its pending staking rewards.
        mapping(bytes32 delegationID => uint256) _redeemableDelegatorRewards;
        /// @notice Maps the validation ID to its pending staking rewards.
        mapping(bytes32 validationID => uint256) _redeemableValidatorRewards;
        /// @notice Saves the uptime of a pending completed or completed validation period so that delegators can collect rewards.
        mapping(bytes32 validationID => uint64) _completedValidationUptimeSeconds;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoSValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Unit test for storage slot and update slot
    bytes32 private constant _POS_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00;

    uint8 public constant MAXIMUM_STAKE_MULTIPLIER_LIMIT = 10;

    uint16 public constant MAXIMUM_DELEGATION_FEE_BIPS = 10000;

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
        uint256 minimumDelegationFeeBips,
        uint8 maximumStakeMultiplier,
        IRewardCalculator rewardCalculator
    ) internal onlyInitializing {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(minimumDelegationFeeBips > 0, "PoSValidatorManager: zero delegation fee");
        require(
            minimumDelegationFeeBips <= MAXIMUM_DELEGATION_FEE_BIPS,
            "PoSValidatorManager: invalid delegation fee"
        );
        require(
            minimumStakeAmount <= maximumStakeAmount,
            "PoSValidatorManager: invalid stake amount range"
        );
        require(maximumStakeMultiplier > 0, "PoSValidatorManager: zero maximum stake multiplier");
        require(
            maximumStakeMultiplier <= MAXIMUM_STAKE_MULTIPLIER_LIMIT,
            "PoSValidatorManager: invalid maximum stake multiplier"
        );

        $._minimumStakeAmount = minimumStakeAmount;
        $._maximumStakeAmount = maximumStakeAmount;
        $._minimumStakeDuration = minimumStakeDuration;
        $._minimumDelegationFeeBips = minimumDelegationFeeBips;
        $._maximumStakeMultiplier = maximumStakeMultiplier;
        $._rewardCalculator = rewardCalculator;
    }

    function claimDelegationFees(bytes32 validationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Validator memory validator = getValidator(validationID);

        require(
            validator.status == ValidatorStatus.Completed,
            "PoSValidatorManager: validation period not completed"
        );
        require(
            $._validatorRequirements[validationID].owner == _msgSender(),
            "PoSValidatorManager: validator not owned by sender"
        );

        uint256 rewards = $._redeemableValidatorRewards[validationID];
        delete $._redeemableValidatorRewards[validationID];
        _reward($._validatorRequirements[validationID].owner, rewards);
    }

    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        Validator memory validator = _initializeEndValidation(validationID);

        if (!_isPoSValidator(validationID)) {
            return;
        }

        // Check that minimum stake duration has passed.
        require(
            validator.endedAt
                >= validator.startedAt + $._validatorRequirements[validationID].minStakeDuration,
            "PoSValidatorManager: minimum stake duration not met"
        );

        if (includeUptimeProof) {
            // Uptime proofs include the absolute number of seconds the validator has been active.
            uint64 uptimeSeconds = _getUptime(validationID, messageIndex);
            // Save this value for use by this validator's delegators.
            $._completedValidationUptimeSeconds[validationID] = uptimeSeconds;

            $._redeemableValidatorRewards[validationID] += $._rewardCalculator.calculateReward({
                stakeAmount: weightToValue(validator.startingWeight),
                validatorStartTime: validator.startedAt,
                stakingStartTime: validator.startedAt,
                stakingEndTime: validator.endedAt,
                uptimeSeconds: uptimeSeconds,
                initialSupply: 0,
                endSupply: 0
            });
        }
    }

    function completeEndValidation(uint32 messageIndex) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        (bytes32 validationID, Validator memory validator) = _completeEndValidation(messageIndex);

        // Return now if this was originally a PoA validator that was later migrated to this PoS manager
        if (!_isPoSValidator(validationID)) {
            return;
        }

        address owner = $._validatorRequirements[validationID].owner;
        // The validator can either be Completed or Invalidated here. We only grant rewards for Completed.
        if (validator.status == ValidatorStatus.Completed) {
            uint256 rewards = $._redeemableValidatorRewards[validationID];
            delete $._redeemableValidatorRewards[validationID];
            _reward(owner, rewards);
        }

        // We unlock the stake whether the validation period is completed or invalidated.
        _unlock(owner, weightToValue(validator.startingWeight));
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

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput calldata registrationInput,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount
    ) internal virtual returns (bytes32) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        // Validate and save the validator requirements
        require(
            delegationFeeBips >= $._minimumDelegationFeeBips
                && delegationFeeBips <= MAXIMUM_DELEGATION_FEE_BIPS,
            "PoSValidatorManager: invalid delegation fee"
        );
        require(
            minStakeDuration >= $._minimumStakeDuration,
            "PoSValidatorManager: invalid min stake duration"
        );

        // Ensure the weight is within the valid range.
        require(stakeAmount >= $._minimumStakeAmount, "PoSValidatorManager: stake amount too low");
        require(stakeAmount <= $._maximumStakeAmount, "PoSValidatorManager: stake amount too high");

        // Lock the stake in the contract.
        uint256 lockedValue = _lock(stakeAmount);

        uint64 weight = valueToWeight(lockedValue);
        bytes32 validationID = _initializeValidatorRegistration(registrationInput, weight);

        $._validatorRequirements[validationID] = PoSValidatorRequirements({
            owner: _msgSender(),
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration
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
        require(_isPoSValidator(validationID), "PoSValidatorManager: not a PoS validator");
        require(
            validator.status == ValidatorStatus.Active, "PoSValidatorManager: validator not active"
        );

        // Update the validator weight
        uint64 newValidatorWeight = validator.weight + weight;
        require(
            newValidatorWeight <= validator.startingWeight * $._maximumStakeMultiplier,
            "PoSValidatorManager: maximum validator weight reached"
        );

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
            endedAt: 0,
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

        // Unpack the Warp message
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce,) =
            ValidatorMessages.unpackSubnetValidatorWeightUpdateMessage(warpMessage.payload);

        Validator memory validator = getValidator(validationID);

        // The received nonce should be no greater than the highest sent nonce. This should never
        // happen since the staking manager is the only entity that can trigger a weight update
        // on the P-Chain.
        require(validator.messageNonce >= nonce, "PoSValidatorManager: invalid nonce");

        // The nonce should also be greater than or equal to the delegationID's starting nonce. This allows
        // a weight update using a higher nonce (which implicitly includes the delegation's weight update)
        // to be used to complete registration for an earlier delegation. This is necessary because the P-Chain
        // is only willing to sign the latest weight update.
        require(
            $._delegatorStakes[delegationID].startingNonce <= nonce,
            "PoSValidatorManager: nonce does not match"
        );

        // Ensure the delegator is pending added. Since anybody can call this function once
        // delegator registration has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        require(
            $._delegatorStakes[delegationID].status == DelegatorStatus.PendingAdded,
            "PoSValidatorManager: delegationID not pending added"
        );
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

    function initializeEndDelegation(
        bytes32 delegationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        bytes32 validationID = $._delegatorStakes[delegationID].validationID;

        Delegator memory delegator = $._delegatorStakes[delegationID];
        Validator memory validator = getValidator(validationID);

        // Ensure the delegator is active
        require(
            delegator.status == DelegatorStatus.Active, "PoSValidatorManager: delegation not active"
        );
        // Only the delegation owner can end the delegation.
        require(
            delegator.owner == _msgSender(), "PoSValidatorManager: delegation not owned by sender"
        );

        // Set the delegator status to pending removed, so that it can be properly removed in
        // the complete step, even if the delivered nonce is greater than the nonce used to
        // initialize the removal.
        delegator.status = DelegatorStatus.PendingRemoved;

        uint64 validatorUptimeSeconds;
        if (validator.status == ValidatorStatus.Active) {
            if (includeUptimeProof) {
                // Uptime proofs include the absolute number of seconds the validator has been active.
                validatorUptimeSeconds = _getUptime(validationID, messageIndex);
            }
            uint64 newValidatorWeight = validator.weight - delegator.weight;
            (delegator.endingNonce,) = _setValidatorWeight(validationID, newValidatorWeight);

            delegator.endedAt = uint64(block.timestamp);
        } else {
            // If the validation period has already ended, we have saved the uptime.
            // Further, it is impossible to retrieve an uptime proof for an already ended validation,
            // so there's no need to check any uptime proof provided in this function call.
            validatorUptimeSeconds = $._completedValidationUptimeSeconds[validationID];

            delegator.endingNonce = validator.messageNonce;
            delegator.endedAt = validator.endedAt;
        }

        $._redeemableDelegatorRewards[delegationID] = $._rewardCalculator.calculateReward({
            stakeAmount: weightToValue(delegator.weight),
            validatorStartTime: validator.startedAt,
            stakingStartTime: delegator.startedAt,
            stakingEndTime: delegator.endedAt,
            uptimeSeconds: validatorUptimeSeconds,
            initialSupply: 0,
            endSupply: 0
        });

        $._delegatorStakes[delegationID] = delegator;

        emit DelegatorRemovalInitialized({
            delegationID: delegationID,
            validationID: validationID,
            endTime: delegator.endedAt
        });
    }

    /**
     * @dev Resending the latest validator weight with the latest nonce is safe because all weight changes are
     * cumulative, so the latest weight change will always include the weight change for any added delegators.
     */
    function resendUpdateDelegation(bytes32 delegationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        Delegator memory delegator = $._delegatorStakes[delegationID];
        require(
            delegator.status == DelegatorStatus.PendingAdded
                || delegator.status == DelegatorStatus.PendingRemoved,
            "PoSValidatorManager: delegation status not pending"
        );

        Validator memory validator = getValidator(delegator.validationID);
        require(
            validator.messageNonce != 0,
            "PoSValidatorManager: could not find validator for delegation ID"
        );

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

        // Unpack the Warp message
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce,) =
            ValidatorMessages.unpackSubnetValidatorWeightUpdateMessage(warpMessage.payload);

        Validator memory validator = getValidator(validationID);
        Delegator memory delegator = $._delegatorStakes[delegationID];

        // The received nonce should be no greater than the highest sent nonce. This should never
        // happen since the staking manager is the only entity that can trigger a weight update
        // on the P-Chain.
        require(validator.messageNonce >= nonce, "PoSValidatorManager: invalid nonce");

        // The nonce should also be greater than or equal to the delegationID's ending nonce. This allows
        // a weight update using a higher nonce (which implicitly includes the delegation's weight update)
        // to be used to complete delisting for an earlier delegation. This is necessary because the P-Chain
        // is only willing to sign the latest weight update.
        require(delegator.endingNonce <= nonce, "PoSValidatorManager: nonce does not match");

        // Ensure the delegator is pending removed. Since anybody can call this function once
        // end delegation has been initialized, we need to make sure that this function is only
        // callable after that has been done.
        require(
            delegator.status == DelegatorStatus.PendingRemoved,
            "PoSValidatorManager: delegation not pending added"
        );

        // Update the delegator status
        $._delegatorStakes[delegationID].status = DelegatorStatus.Completed;

        uint256 rewards = $._redeemableDelegatorRewards[delegationID];
        delete $._redeemableDelegatorRewards[delegationID];

        uint256 validatorFees =
            rewards * $._validatorRequirements[validationID].delegationFeeBips / 10000;
        $._redeemableValidatorRewards[validationID] += validatorFees;

        _reward(delegator.owner, rewards - validatorFees);
        _unlock(delegator.owner, weightToValue(delegator.weight));
        delete $._delegatorStakes[delegationID];

        emit DelegationEnded(delegationID, validationID, nonce);
    }

    function _reward(address account, uint256 amount) internal virtual;

    function _isPoSValidator(bytes32 validationID) internal view returns (bool) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        return $._validatorRequirements[validationID].owner != address(0);
    }
}
