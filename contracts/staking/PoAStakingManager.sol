// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IPoAStakingManager} from "./interfaces/IPoAStakingManager.sol";
import {StakingManager} from "./StakingManager.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {ICMInitializable} from "../utilities/ICMInitializable.sol";
import {StakingManager, StakingManagerSettings} from "./StakingManager.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";

contract PoAStakingManager is IPoAStakingManager, StakingManager, OwnableUpgradeable {
    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    function initialize(
        StakingManagerSettings calldata settings,
        address initialOwner
    ) external initializer {
        __PoAStakingManager_init(settings, initialOwner);
    }

    // solhint-disable func-name-mixedcase
    function __PoAStakingManager_init(
        StakingManagerSettings calldata settings,
        address initialOwner
    ) internal onlyInitializing {
        __StakingManager_init(settings);
        __Ownable_init(initialOwner);
        __PoAStakingManager_init_unchained();
    }

    // solhint-disable func-name-mixedcase
    function __PoAStakingManager_init_unchained() internal onlyInitializing {}

    function initializeValidatorRegistration(
        uint256 stakeAmount,
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature
    ) external override onlyOwner returns (bytes32 validationID) {
        return _initializeValidatorRegistration(nodeID, stakeAmount, registrationExpiry, signature);
    }

    function _lock(uint256 value) internal virtual override returns (uint256) {
        return value;
    }

    function _unlock(uint256 value, address to) internal virtual override {}
}
