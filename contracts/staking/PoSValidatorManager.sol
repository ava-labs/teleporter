// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IPoSValidatorManager} from "./interfaces/IPoSValidatorManager.sol";
import {StakingManager} from "./StakingManager.sol";
import {ICMInitializable} from "../utilities/ICMInitializable.sol";
import {WarpMessage} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {StakingMessages} from "./StakingMessages.sol";

abstract contract PoSValidatorManager is IPoSValidatorManager, StakingManager {
    function initializeEndValidation(
        bytes32 validationID,
        bool includeUptimeProof,
        uint32 messageIndex
    ) external {
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
        _initializeEndValidation(validationID, uptimeSeconds);
    }
}
