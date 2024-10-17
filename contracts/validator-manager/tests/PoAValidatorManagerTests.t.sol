// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {PoAValidatorManager} from "../PoAValidatorManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {
    ValidatorManagerSettings,
    ValidatorRegistrationInput
} from "../interfaces/IValidatorManager.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";
import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {Codec} from "../Codec.sol";

contract PoAValidatorManagerTest is ValidatorManagerTest {
    PoAValidatorManager public app;

    address public constant DEFAULT_OWNER = address(0x1);

    function setUp() public override {
        ValidatorManagerTest.setUp();

        app = new PoAValidatorManager(ICMInitializable.Allowed);
        codec = new Codec();
        app.initialize(
            ValidatorManagerSettings({
                subnetID: DEFAULT_SUBNET_ID,
                churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE,
                codec: codec
            }),
            address(this)
        );
        validatorManager = app;
        _mockGetBlockchainID();
        _mockInitializeValidatorSet();
        app.initializeValidatorSet(_defaultSubnetConversionData(), 0);
    }

    function testInvalidOwnerRegistration() public {
        vm.prank(vm.addr(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, vm.addr(1)
            )
        );
        _initializeValidatorRegistration(
            ValidatorRegistrationInput({
                nodeID: DEFAULT_NODE_ID,
                blsPublicKey: DEFAULT_BLS_PUBLIC_KEY,
                registrationExpiry: DEFAULT_EXPIRY,
                remainingBalanceOwner: DEFAULT_P_CHAIN_OWNER,
                disableOwner: DEFAULT_P_CHAIN_OWNER
            }),
            DEFAULT_WEIGHT
        );
    }

    function _initializeValidatorRegistration(
        ValidatorRegistrationInput memory input,
        uint64 weight
    ) internal virtual override returns (bytes32) {
        return app.initializeValidatorRegistration(input, weight);
    }

    function _initializeEndValidation(bytes32 validationID, bool) internal virtual override {
        return app.initializeEndValidation(validationID);
    }

    function _forceInitializeEndValidation(bytes32 validationID, bool) internal virtual override {
        return app.initializeEndValidation(validationID);
    }

    // solhint-disable-next-line no-empty-blocks
    function _beforeSend(uint256 amount, address spender) internal virtual override {}
}
