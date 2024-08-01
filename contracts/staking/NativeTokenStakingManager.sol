// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {INativeTokenStakingManager} from "./interfaces/INativeTokenStakingManager.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {StakingManager} from "./StakingManager.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

contract NativeTokenStakingManager is Initializable, StakingManager, INativeTokenStakingManager {
    using Address for address payable;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The time at which the reigistration is no longer valid on the P-Chain.
     * @param signature The raw bytes of the Ed25519 signature over the concatenated bytes of
     * [subnetID]+[nodeID]+[blsPublicKey]+[weight]+[balance]+[expiry]. This signature must correspond to the Ed25519
     * public key that is used for the nodeID. This approach prevents NodeIDs from being unwillingly added to Subnets.
     * balance is the minimum initial $nAVAX balance that must be attached to the validator serialized as a uint64.
     * The signature field will be validated by the P-Chain. Implementations may choose to validate that the signature
     * field is well-formed but it is not required.
     */
    function initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature
    ) external payable onlyWhenInitialStakeProvided returns (bytes32) {
        return _initializeValidatorRegistration(nodeID, msg.value, registrationExpiry, signature);
    }

    function _lock(uint256 value) internal virtual override returns (uint256) {
        return value;
    }

    function _unlock(uint256 value, address to) internal virtual override {
        payable(to).sendValue(value);
    }
}
