// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {PoAValidatorManager} from "../PoAValidatorManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {ValidatorManagerSettings} from "../interfaces/IValidatorManager.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

contract PoAValidatorManagerTest is ValidatorManagerTest {
    PoAValidatorManager public app;

    address public constant DEFAULT_OWNER = address(0x1);

    function setUp() public {
        app = new PoAValidatorManager(ICMInitializable.Allowed);
        app.initialize(
            ValidatorManagerSettings({
                pChainBlockchainID: P_CHAIN_BLOCKCHAIN_ID,
                subnetID: DEFAULT_SUBNET_ID,
                maximumHourlyChurn: DEFAULT_MAXIMUM_HOURLY_CHURN
            }),
            address(this)
        );
        validatorManager = app;
    }

    function testInvalidOwnerRegistration() public {
        vm.prank(vm.addr(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, vm.addr(1)
            )
        );
        _initializeValidatorRegistration(
            DEFAULT_NODE_ID, DEFAULT_EXPIRY, DEFAULT_BLS_PUBLIC_KEY, DEFAULT_WEIGHT
        );
    }

    function _initializeValidatorRegistration(
        bytes32 nodeID,
        uint64 registrationExpiry,
        bytes memory signature,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration(weight, nodeID, registrationExpiry, signature);
    }

    function _initializeEndValidation(bytes32 validationID) internal virtual override {
        return app.initializeEndValidation(validationID);
    }

    // solhint-disable-next-line no-empty-blocks
    function _beforeSend(uint64 weight, address spender) internal virtual override {}
}
