// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {INativeTokenStakingManager} from "./interfaces/INativeTokenStakingManager.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {ICMInitializable} from "../utilities/ICMInitializable.sol";
import {PoSValidatorManager} from "./PoSValidatorManager.sol";
import {PoSValidatorManagerSettings} from "./interfaces/IPoSValidatorManager.sol";

contract NativeTokenStakingManager is
    Initializable,
    PoSValidatorManager,
    INativeTokenStakingManager
{
    using Address for address payable;

    INativeMinter public constant NATIVE_MINTER =
        INativeMinter(0x0200000000000000000000000000000000000001);

    constructor(ICMInitializable init) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initialize the ERC20 token staking manager
     * @dev Uses reinitializer(2) on the PoS staking contracts to make sure after migration from PoA, the PoS contracts can reinitialize with its needed values.
     * @param settings Initial settings for the PoS validator manager
     */
    // solhint-disable ordering
    function initialize(PoSValidatorManagerSettings calldata settings) external reinitializer(2) {
        __NativeTokenStakingManager_init(settings);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __NativeTokenStakingManager_init(PoSValidatorManagerSettings calldata settings)
        internal
        onlyInitializing
    {
        __POS_Validator_Manager_init(settings);
    }

    // solhint-disable-next-line func-name-mixedcase, no-empty-blocks
    function __NativeTokenStakingManager_init_unchained() internal onlyInitializing {}

    /**
     * @notice Begins the validator registration process. Locks the provided native asset in the contract as the stake.
     * @param nodeID The node ID of the validator being registered.
     * @param registrationExpiry The Unix timestamp after which the registration is no longer valid on the P-Chain.
     * @param blsPublicKey The BLS public key of the validator.
     */
    function initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory blsPublicKey
    ) external payable returns (bytes32) {
        uint64 weight = _processStake(msg.value);
        return _initializeValidatorRegistration(nodeID, weight, registrationExpiry, blsPublicKey);
    }

    /**
     * @notice Begins the delegator registration process. Locks the provided native asset in the contract as the delegated stake.
     * @param validationID The ID of the validation period being delegated to.
     */
    function initializeDelegatorRegistration(bytes32 validationID)
        external
        payable
        returns (bytes32)
    {
        return _initializeDelegatorRegistration(validationID, _msgSender(), msg.value);
    }

    // solhint-enable ordering
    function _lock(uint256 value) internal virtual override returns (uint256) {
        return value;
    }

    function _unlock(uint256 value, address to) internal virtual override {
        payable(to).sendValue(value);
    }

    // solhint-disable-next-line no-empty-blocks
    function _reward(address account, uint256 amount) internal virtual override {
        NATIVE_MINTER.mintNativeCoin(account, amount);
    }
}
