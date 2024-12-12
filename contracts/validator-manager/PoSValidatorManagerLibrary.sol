pragma solidity 0.8.25;

import {PoSValidatorManager} from "./PoSValidatorManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";
import {
    Validator,
    ValidatorStatus
} from "./interfaces/IValidatorManager.sol";
import {IWarpMessenger, WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {IPoSValidatorManager} from "./interfaces/IPoSValidatorManager.sol";


library PoSValidatorManagerLibrary {
    function weightToValue(PoSValidatorManager.PoSValidatorManagerStorage storage self, uint64 weight) external view returns (uint256) {
        return _weightToValue(self, weight);
    }

    function valueToWeight(PoSValidatorManager.PoSValidatorManagerStorage storage self, uint256 value) external view returns (uint64) {
        return _valueToWeight(self, value);
    }

    function isPoSValidator(PoSValidatorManager.PoSValidatorManagerStorage storage self, bytes32 validationID) external view returns (bool) {
        return _isPoSValidator(self, validationID);
    }

    function updateUptime(PoSValidatorManager.PoSValidatorManagerStorage storage self, IWarpMessenger WARP_MESSENGER, bytes32 validationID, uint32 messageIndex) external returns (uint64) {
        return _updateUptime(self, WARP_MESSENGER, validationID, messageIndex);
    }

    function _weightToValue(PoSValidatorManager.PoSValidatorManagerStorage storage self, uint64 weight) internal view returns (uint256) {
        return uint256(weight) * self._weightToValueFactor;
    }

    function _valueToWeight(PoSValidatorManager.PoSValidatorManagerStorage storage self, uint256 value) internal view returns (uint64) {
        uint256 weight = value / self._weightToValueFactor;
        if (weight == 0 || weight > type(uint64).max) {
            revert PoSValidatorManager.InvalidStakeAmount(value);
        }
        return uint64(weight);
    }

    function _isPoSValidator(PoSValidatorManager.PoSValidatorManagerStorage storage self,bytes32 validationID) internal view returns (bool) {
        return self._posValidatorInfo[validationID].owner != address(0);
    }

    function _updateUptime(PoSValidatorManager.PoSValidatorManagerStorage storage self, IWarpMessenger WARP_MESSENGER, bytes32 validationID, uint32 messageIndex) internal returns (uint64) {
        (WarpMessage memory warpMessage, bool valid) =
            WARP_MESSENGER.getVerifiedWarpMessage(messageIndex);
        if (!valid) {
            revert ValidatorManager.InvalidWarpMessage();
        }

        // The uptime proof must be from the specifed uptime blockchain
        if (warpMessage.sourceChainID != self._uptimeBlockchainID) {
            revert ValidatorManager.InvalidWarpSourceChainID(warpMessage.sourceChainID);
        }

        // The sender is required to be the zero address so that we know the validator node
        // signed the proof directly, rather than as an arbitrary on-chain message
        if (warpMessage.originSenderAddress != address(0)) {
            revert ValidatorManager.InvalidWarpOriginSenderAddress(warpMessage.originSenderAddress);
        }
        if (warpMessage.originSenderAddress != address(0)) {
            revert ValidatorManager.InvalidWarpOriginSenderAddress(warpMessage.originSenderAddress);
        }

        (bytes32 uptimeValidationID, uint64 uptime) =
            ValidatorMessages.unpackValidationUptimeMessage(warpMessage.payload);
        if (validationID != uptimeValidationID) {
            revert ValidatorManager.InvalidValidationID(validationID);
        }

        if (uptime > self._posValidatorInfo[validationID].uptimeSeconds) {
            self._posValidatorInfo[validationID].uptimeSeconds = uptime;
            emit IPoSValidatorManager.UptimeUpdated(validationID, uptime);
        } else {
            uptime = self._posValidatorInfo[validationID].uptimeSeconds;
        }

        return uptime;
    }

    function submitUptimeProof(PoSValidatorManager.PoSValidatorManagerStorage storage self, IWarpMessenger WARP_MESSENGER, Validator memory validator, bytes32 validationID, uint32 messageIndex) external {
        if (!_isPoSValidator(self, validationID)) {
            revert PoSValidatorManager.ValidatorNotPoS(validationID);
        }
        ValidatorStatus status = validator.status;
        if (status != ValidatorStatus.Active) {
            revert ValidatorManager.InvalidValidatorStatus(status);
        }

        // Uptime proofs include the absolute number of seconds the validator has been active.
        _updateUptime(self, WARP_MESSENGER, validationID, messageIndex);
    }

    function initializeEndPoSValidation(
        PoSValidatorManager.PoSValidatorManagerStorage storage self,
        IWarpMessenger WARP_MESSENGER,
        Validator memory validator,
        address sender,
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex,
        address rewardRecipient
    ) external returns (bool) {
        // Non-PoS validators are required to boostrap the network, but are not eligible for rewards.
        if (!_isPoSValidator(self, validationID)) {
            return true;
        }

        // PoS validations can only be ended by their owners.
        if (self._posValidatorInfo[validationID].owner != sender) {
            revert PoSValidatorManager.UnauthorizedOwner(sender);
        }

        // Check that minimum stake duration has passed.
        if (
            validator.endedAt
                < validator.startedAt + self._posValidatorInfo[validationID].minStakeDuration
        ) {
            revert PoSValidatorManager.MinStakeDurationNotPassed(validator.endedAt);
        }

        // Uptime proofs include the absolute number of seconds the validator has been active.
        uint64 uptimeSeconds;
        if (includeUptimeProof) {
            uptimeSeconds = _updateUptime(self, WARP_MESSENGER, validationID, messageIndex);
        } else {
            uptimeSeconds = self._posValidatorInfo[validationID].uptimeSeconds;
        }

        uint256 reward = self._rewardCalculator.calculateReward({
            stakeAmount: _weightToValue(self, validator.startingWeight),
            validatorStartTime: validator.startedAt,
            stakingStartTime: validator.startedAt,
            stakingEndTime: validator.endedAt,
            uptimeSeconds: uptimeSeconds
        });

        if (rewardRecipient == address(0)) {
            rewardRecipient = self._posValidatorInfo[validationID].owner;
        }

        self._redeemableValidatorRewards[validationID] += reward;
        self._rewardRecipients[validationID] = rewardRecipient;

        return (reward > 0);
    }
}