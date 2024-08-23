// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IPoSValidatorManager} from "./interfaces/IPoSValidatorManager.sol";
import {PoSValidatorManagerSettings} from "./interfaces/IPoSValidatorManager.sol";
import {ValidatorManager} from "./ValidatorManager.sol";
import {WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ValidatorMessages} from "./ValidatorMessages.sol";
import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";

abstract contract PoSValidatorManager is IPoSValidatorManager, ValidatorManager {
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.PoSValidatorManager
    struct PoSValidatorManagerStorage {
        uint256 _minimumStakeAmount;
        uint256 _maximumStakeAmount;
        uint64 _minimumStakeDuration;
        IRewardCalculator _rewardCalculator;
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
        PoSValidatorManagerStorage storage s = _getPoSValidatorManagerStorage();
        s._minimumStakeAmount = minimumStakeAmount;
        s._maximumStakeAmount = maximumStakeAmount;
        s._minimumStakeDuration = minimumStakeDuration;
        s._rewardCalculator = rewardCalculator;
    }

    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
        ValidatorManagerStorage storage $ = _getValidatorManagerStorage();
        uint64 uptimeSeconds;
        if (includeUptimeProof) {
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

            (bytes32 uptimeValidationID, uint64 uptime) =
                ValidatorMessages.unpackValidationUptimeMessage(warpMessage.payload);
            require(
                validationID == uptimeValidationID, "StakingManager: Invalid uptime validation ID"
            );
            uptimeSeconds = uptime;
        }
        _initializeEndValidation(validationID, uptimeSeconds);
    }

    function _processStake(uint256 stakeAmount) internal virtual returns (uint64) {
        PoSValidatorManagerStorage storage $ = _getPoSValidatorManagerStorage();
        // Lock the stake in the contract.
        uint256 lockedValue = _lock(stakeAmount);

        // Ensure the stake churn doesn't exceed the maximum churn rate.
        uint64 weight = valueToWeight(lockedValue);
        // Ensure the weight is within the valid range.

        require(
            weight >= $._minimumStakeAmount && weight <= $._maximumStakeAmount,
            "PoSValidatorManager: Invalid stake amount"
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
}
