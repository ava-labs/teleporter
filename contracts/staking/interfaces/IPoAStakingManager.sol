// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IStakingManager} from "./IStakingManager.sol";

interface IPoAStakingManager is IStakingManager {
    /**
     * @notice Begins the validator registration process. Locks the {stakeAmount} of the managers specified ERC20 token.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The time at which the reigistration is no longer valid on the P-Chain.
     */
    function initializeValidatorRegistration(
        uint256 stakeAmount,
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature
    ) external returns (bytes32 validationID);
}
