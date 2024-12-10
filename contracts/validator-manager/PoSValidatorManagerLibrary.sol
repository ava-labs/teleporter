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
    function isPoSValidator(PoSValidatorManager.PoSValidatorManagerStorage storage self, bytes32 validationID) external view returns (bool) {
        return _isPoSValidator(self, validationID);
    }

    function updateUptime(PoSValidatorManager.PoSValidatorManagerStorage storage self, IWarpMessenger WARP_MESSENGER, bytes32 validationID, uint32 messageIndex) external returns (uint64) {
        return _updateUptime(self, WARP_MESSENGER, validationID, messageIndex);
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
}