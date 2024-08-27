// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IPoSValidatorManager} from "./IPoSValidatorManager.sol";

interface IERC20TokenStakingManager is IPoSValidatorManager {
    /**
     * @notice Begins the validator registration process. Locks the {stakeAmount} of the managers specified ERC20 token.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The time at which the reigistration is no longer valid on the P-Chain.
     * @param blsPublicKey The BLS public key of the validator.
     */
    function initializeValidatorRegistration(
        uint256 stakeAmount,
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory blsPublicKey
    ) external returns (bytes32 validationID);

    function initializeDelegatorRegistration(
        bytes32 validationID,
        uint256 stakeAmount
    ) external;
}
