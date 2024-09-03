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
        /// @notice The minimum fee rate in basis points charged to the delegators by the validator.
        uint256 _minimumDelegationFeeRate;
        /// @notice The reward calculator for this validator manager.
        IRewardCalculator _rewardCalculator;
        /// @notice Maps the validationID to the uptime of the validator.
        mapping(bytes32 validationID => uint64) _uptimes;
        /// @notice Maps the validationID to a mapping of delegator address to delegator information.
        mapping(bytes32 validationID => mapping(address delegator => Delegator)) _delegatorStakes;
        /// @notice Maps the validationID to a mapping of delegator address to pending register delegator messages.
        mapping(bytes32 validationID => mapping(address delegator => bytes))
            _pendingRegisterDelegatorMessages;
        /// @notice Maps the validationID to a mapping of delegator address to pending end delegator messages.
        mapping(bytes32 validationID => mapping(address delegator => bytes))
            _pendingEndDelegatorMessages;
        /// @notice Maps the validationID to the uptime of the validator.
        mapping(bytes32 validationID => uint64) _validatorUptimes;
        /// @notice Maps the validationID to the delegation fee rate.
        mapping(bytes32 validationID => uint256) _validatorDelegationFeeRates;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.PoSValidatorManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Unit test for storage slot and update slot
    bytes32 private constant _POS_VALIDATOR_MANAGER_STORAGE_LOCATION =
        0x4317713f7ecbdddd4bc99e95d903adedaa883b2e7c2551610bd13e2c7e473d00;

    // The maximum delegation fee rate in basis points. This is 99.99%
    uint256 private constant _MAXIMUM_DELEGATION_FEE_RATE = 1e4 - 1;

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
            settings.minimumDelegationFeeRate,
            settings.rewardCalculator
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __POS_Validator_Manager_init_unchained(
        uint256 minimumStakeAmount,
        uint256 maximumStakeAmount,
        uint64 minimumStakeDuration,
        uint256 minimumDelegationFeeRate,
        IRewardCalculator rewardCalculator
    ) internal onlyInitializing {
        require(
            minimumDelegationFeeRate > 0 && minimumDelegationFeeRate < _MAXIMUM_DELEGATION_FEE_RATE,
            "PoSValidatorManager: invalid delegation fee rate"
        );
        PoSValidatorManagerStorage storage s = _getPoSValidatorManagerStorage();
        s._minimumStakeAmount = minimumStakeAmount;
        s._maximumStakeAmount = maximumStakeAmount;
        s._minimumStakeDuration = minimumStakeDuration;
        s._rewardCalculator = rewardCalculator;
        s._minimumDelegationFeeRate = minimumDelegationFeeRate;
    }

    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        if (includeUptimeProof) {
            PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
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
                validationID == uptimeValidationID,
                "PoSValidatorManager: invalid uptime validation ID"
            );

            $._validatorUptimes[validationID] = uptime;
            emit ValidationUptimeUpdated(validationID, uptime);
        }

        _initializeEndValidation(validationID);
    }

    function _setDelegationFeeRate(bytes32 validationID, uint256 feeRate) internal {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(
            feeRate >= $._minimumDelegationFeeRate && feeRate <= _MAXIMUM_DELEGATION_FEE_RATE,
            "PoSValidatorManager: invalid delegation fee rate"
        );
        $._validatorDelegationFeeRates[validationID] = feeRate;
        emit DelegationFeeRateSet(validationID, feeRate);
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
    ) internal nonReentrant {
        uint64 weight = valueToWeight(_lock(delegationAmount));
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Ensure the validation period is active
        Validator memory validator = _getValidator(validationID);
        require(
            validator.status == ValidatorStatus.Active, "ValidatorManager: validator not active"
        );

        // Ensure the delegator is not already registered
        require(
            $._delegatorStakes[validationID][delegator].weight == 0,
            "ValidatorManager: delegator already registered"
        );

        _checkAndUpdateChurnTracker(weight);

        // Update the validator weight
        validator.weight += weight;
        _setValidator(validationID, validator);

        // Submit the message to the Warp precompile.
        uint64 nonce = _getAndIncrementNonce(validationID);
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, nonce, validator.weight);
        $._pendingRegisterDelegatorMessages[validationID][delegator] = setValidatorWeightPayload;
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);

        // Store the delegator information
        $._delegatorStakes[validationID][delegator] = Delegator({
            weight: weight,
            startedAt: 0,
            endedAt: 0,
            status: DelegatorStatus.PendingAdded
        });

        emit DelegatorAdded({
            validationID: validationID,
            setWeightMessageID: messageID,
            delegator: delegator,
            delegatorWeight: weight,
            validatorWeight: validator.weight,
            nonce: nonce
        });
    }

    function resendDelegatorRegistration(bytes32 validationID, address delegator) external {
        _checkPendingRegisterDelegatorMessages(validationID, delegator);
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        // Submit the message to the Warp precompile.
        WARP_MESSENGER.sendWarpMessage($._pendingRegisterDelegatorMessages[validationID][delegator]);
    }

    function completeDelegatorRegistration(uint32 messageIndex, address delegator) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Unpack the Warp message
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce,) =
            ValidatorMessages.unpackSubnetValidatorWeightUpdateMessage(warpMessage.payload);
        _checkPendingRegisterDelegatorMessages(validationID, delegator);
        delete $._pendingRegisterDelegatorMessages[validationID][delegator];

        // Update the validator weight and delegator status
        Validator memory validator = _getValidator(validationID);
        require(validator.messageNonce >= nonce, "PoSValidatorManager: invalid nonce");

        $._delegatorStakes[validationID][delegator].status = DelegatorStatus.Active;
        $._delegatorStakes[validationID][delegator].startedAt = uint64(block.timestamp);

        emit DelegatorRegistered({
            validationID: validationID,
            delegator: delegator,
            nonce: nonce,
            startTime: block.timestamp
        });
    }

    function initializeEndDelegation(bytes32 validationID) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Ensure the delegator is active
        Delegator memory delegator = $._delegatorStakes[validationID][_msgSender()];
        require(
            delegator.status == DelegatorStatus.Active, "ValidatorManager: delegator not active"
        );
        delegator.status = DelegatorStatus.PendingRemoved;
        delegator.endedAt = uint64(block.timestamp);

        $._delegatorStakes[validationID][_msgSender()] = delegator;

        Validator memory validator = _getValidator(validationID);
        require(validator.weight > delegator.weight, "PoSValidatorManager: Invalid weight");
        validator.weight -= delegator.weight;
        _setValidator(validationID, validator);

        // Submit the message to the Warp precompile.
        uint64 nonce = _getAndIncrementNonce(validationID);
        bytes memory setValidatorWeightPayload = ValidatorMessages
            .packSetSubnetValidatorWeightMessage(validationID, nonce, validator.weight);
        $._pendingEndDelegatorMessages[validationID][_msgSender()] = setValidatorWeightPayload;
        bytes32 messageID = WARP_MESSENGER.sendWarpMessage(setValidatorWeightPayload);

        emit DelegatorRemovalInitialized({
            validationID: validationID,
            setWeightMessageID: messageID,
            delegator: _msgSender(),
            validatorWeight: validator.weight,
            nonce: nonce,
            endTime: block.timestamp
        });
    }

    function resendEndDelegation(bytes32 validationID, address delegator) external {
        _checkPendingEndDelegatorMessage(validationID, delegator);
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        WARP_MESSENGER.sendWarpMessage($._pendingEndDelegatorMessages[validationID][delegator]);
    }

    function completeEndDelegation(uint32 messageIndex, address delegator) external {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();

        // Unpack the Warp message
        WarpMessage memory warpMessage = _getPChainWarpMessage(messageIndex);
        (bytes32 validationID, uint64 nonce,) =
            ValidatorMessages.unpackSubnetValidatorWeightUpdateMessage(warpMessage.payload);
        _checkPendingEndDelegatorMessage(validationID, delegator);
        delete $._pendingEndDelegatorMessages[validationID][delegator];

        // Update the validator weight and delegator status
        Validator memory validator = _getValidator(validationID);
        require(validator.messageNonce >= nonce, "PoSValidatorManager: invalid nonce");
        $._delegatorStakes[validationID][delegator].status = DelegatorStatus.Completed;

        emit DelegationEnded(validationID, delegator, nonce);
    }

    function _checkPendingEndDelegatorMessage(
        bytes32 validationID,
        address delegator
    ) private view {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(
            $._pendingEndDelegatorMessages[validationID][delegator].length > 0
                && $._delegatorStakes[validationID][delegator].status == DelegatorStatus.PendingRemoved,
            "PoSValidatorManager: delegator removal not pending"
        );
    }

    function _checkPendingRegisterDelegatorMessages(
        bytes32 validationID,
        address delegator
    ) private view {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        require(
            $._pendingRegisterDelegatorMessages[validationID][delegator].length > 0
                && $._delegatorStakes[validationID][delegator].status == DelegatorStatus.PendingAdded,
            "PoSValidatorManager: delegator registration not pending"
        );
    }
}
