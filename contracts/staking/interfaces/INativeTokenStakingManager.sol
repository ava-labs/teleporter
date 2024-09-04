// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IPoSValidatorManager} from "./IPoSValidatorManager.sol";

interface INativeTokenStakingManager is IPoSValidatorManager {
    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The Unix timestamp after which the reigistration is no longer valid on the P-Chain.
     * @param delegationFeeRate The fee rate in basis points charged by this validator for delegations.
     * @param blsPublicKey The BLS public key of the validator.
     */
    function initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        uint256 delegationFeeRate,
        bytes memory blsPublicKey
    ) external payable returns (bytes32 validationID);

    function initializeDelegatorRegistration(bytes32 validationID) external payable;
}
