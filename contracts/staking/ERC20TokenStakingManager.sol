// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {IERC20TokenStakingManager} from "./interfaces/IERC20TokenStakingManager.sol";
import {StakingManager, StakingManagerSettings} from "./StakingManager.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {IERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/IERC20.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";

contract ERC20TokenStakingManager is Initializable, StakingManager, IERC20TokenStakingManager {
    using SafeERC20 for IERC20;
    using SafeERC20TransferFrom for IERC20;

    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avalanche-icm.storage.ERC20TokenStakingManager
    struct ERC20TokenStakingManagerStorage {
        IERC20 _token;
        uint8 _tokenDecimals;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avalanche-icm.storage.ERC20TokenStakingManager")) - 1)) & ~bytes32(uint256(0xff));
    // TODO: Update to correct storage slot
    bytes32 private constant _ERC20_STAKING_MANAGER_STORAGE_LOCATION =
        0x8568826440873e37a96cb0aab773b28d8154d963d2f0e41bd9b5c15f63625f91;

    // solhint-disable ordering
    function _getERC20StakingManagerStorage()
        private
        pure
        returns (ERC20TokenStakingManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := _ERC20_STAKING_MANAGER_STORAGE_LOCATION
        }
    }

    constructor() {
        _disableInitializers();
    }

    function initialize(
        StakingManagerSettings calldata settings,
        IERC20 token,
        uint8 tokenDecimals
    ) external initializer {
        __ERC20TokenStakingManager_init(settings, token, tokenDecimals);
    }

    function __ERC20TokenStakingManager_init(
        StakingManagerSettings calldata settings,
        IERC20 token,
        uint8 tokenDecimals
    ) internal onlyInitializing {
        __StakingManager_init(settings, tokenDecimals);
        __ERC20TokenStakingManager_init_unchained(token);
    }

    function __ERC20TokenStakingManager_init_unchained(IERC20 token) internal onlyInitializing {
        ERC20TokenStakingManagerStorage storage $ = _getERC20StakingManagerStorage();
        require(address(token) != address(0), "ERC20TokenStakingManager: zero token address");
        $._token = token;
    }

    function initializeValidatorRegistration(
        uint256 stakeAmount,
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature
    ) external override returns (bytes32 validationID) {
        return _initializeValidatorRegistration(nodeID, stakeAmount, registrationExpiry, signature);
    }

    // Must be guarded with reentrancy guard for safe transfer from
    function _lock(uint256 value) internal virtual override returns (uint256) {
        return _getERC20StakingManagerStorage()._token.safeTransferFrom(value);
    }

    function _unlock(uint256 value, address to) internal virtual override {
        _getERC20StakingManagerStorage()._token.safeTransfer(to, value);
    }
}
