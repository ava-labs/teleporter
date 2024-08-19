// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {INativeTokenStakingManager} from "./interfaces/INativeTokenStakingManager.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {ICMInitializable} from "../utilities/ICMInitializable.sol";
import {PoSValidatorManager} from "./PoSValidatorManager.sol";
import {ValidatorManagerSettings} from "./interfaces/IValidatorManager.sol";
import {IRewardCalculator} from "./interfaces/IRewardCalculator.sol";

contract NativeTokenStakingManager is
    Initializable,
    PoSValidatorManager,
    INativeTokenStakingManager
{
    using Address for address payable;

    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    // solhint-disable ordering
    function initialize(
        ValidatorManagerSettings calldata settings,
        uint256 minimumStakeAmount,
        uint256 maximumStakeAmount,
        uint64 minimumStakeDuration,
        IRewardCalculator rewardCalculator
    ) external initializer {
        __NativeTokenStakingManager_init(
            settings, minimumStakeAmount, maximumStakeAmount, minimumStakeDuration, rewardCalculator
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __NativeTokenStakingManager_init(
        ValidatorManagerSettings calldata settings,
        uint256 minimumStakeAmount,
        uint256 maximumStakeAmount,
        uint64 minimumStakeDuration,
        IRewardCalculator rewardCalculator
    ) internal onlyInitializing {
        __POS_Validator_Manager_init(
            settings, minimumStakeAmount, maximumStakeAmount, minimumStakeDuration, rewardCalculator
        );
    }

    // solhint-disable-next-line func-name-mixedcase, no-empty-blocks
    function __NativeTokenStakingManager_init_unchained() internal onlyInitializing {}

    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The time at which the registration is no longer valid on the P-Chain.
     * @param signature The raw bytes of the Ed25519 signature over the concatenated bytes of
     * [subnetID]+[nodeID]+[blsPublicKey]+[weight]+[balance]+[expiry]. This signature must correspond to the Ed25519
     * public key that is used for the nodeID. This approach prevents nodeIDs from being unwillingly added to Subnets.
     * balance is the minimum initial $nAVAX balance that must be attached to the validator serialized as a uint64.
     * The signature field will be validated by the P-Chain. Implementations may choose to validate that the signature
     * field is well-formed but it is not required.
     */
    function initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature
    ) external payable returns (bytes32) {
        uint64 weight = _processStake(msg.value);

        return _initializeValidatorRegistration(nodeID, weight, registrationExpiry, signature);
    }

    // solhint-enable ordering
    function _lock(uint256 value) internal virtual override returns (uint256) {
        return value;
    }

    function _unlock(uint256 value, address to) internal virtual override {
        payable(to).sendValue(value);
    }
}