// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.25;

import {ValidatorManagerTest} from "./ValidatorManagerTests.t.sol";
import {PoAValidatorManager} from "../PoAValidatorManager.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {SubnetConversionData} from "../ValidatorManager.sol";
import {ValidatorMessages} from "../ValidatorMessages.sol";
import {
    ValidatorManagerSettings,
    ValidatorRegistrationInput
} from "../interfaces/IValidatorManager.sol";
import {OwnableUpgradeable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/access/OwnableUpgradeable.sol";

contract PoAValidatorManagerTest is ValidatorManagerTest {
    PoAValidatorManager public app;

    address public constant DEFAULT_OWNER = address(0x1);

    function setUp() public {
        app = new PoAValidatorManager(ICMInitializable.Allowed);
        app.initialize(
            ValidatorManagerSettings({
                subnetID: DEFAULT_SUBNET_ID,
                churnPeriodSeconds: DEFAULT_CHURN_PERIOD,
                maximumChurnPercentage: DEFAULT_MAXIMUM_CHURN_PERCENTAGE
            }),
            address(this)
        );
        validatorManager = app;
        _mockGetBlockchainID();
        _mockInitializeValidatorSet();
        app.initializeValidatorSet(_defaultSubnetConversionData(1), 0);
    }

    function testInvalidOwnerRegistration() public {
        vm.prank(vm.addr(1));
        vm.expectRevert(
            abi.encodeWithSelector(
                OwnableUpgradeable.OwnableUnauthorizedAccount.selector, vm.addr(1)
            )
        );
        _initializeValidatorRegistration(
            ValidatorRegistrationInput(DEFAULT_NODE_ID, DEFAULT_EXPIRY, DEFAULT_BLS_PUBLIC_KEY),
            DEFAULT_WEIGHT
        );
    }

    function testSubnetConversionDataPackingOld() public view {
        uint8 numValidators = 50;
        SubnetConversionData memory scd = _defaultSubnetConversionData(numValidators);
        bytes memory packedOld = _oldSubnetConversionDataPack(scd);
        uint256 length = 92 + scd.initialValidators.length * 88;

        require(packedOld.length == length, "Old Packed data length does not match");
    }

    function testSubnetConversionDataPackingNew() public view {
        uint8 numValidators = 50;
        SubnetConversionData memory scd = _defaultSubnetConversionData(numValidators);
        bytes memory packedNew = ValidatorMessages.packSubnetConversionData(scd);
        uint256 length = 92 + scd.initialValidators.length * 88;

        require(packedNew.length == length, "New Packed data length does not match");
    }

    function testSubnetConversionDataPacking() public view {
        uint8 numValidators = 50;
        SubnetConversionData memory scd = _defaultSubnetConversionData(numValidators);
        bytes memory packedOld = _oldSubnetConversionDataPack(scd);
        bytes memory packedNew = ValidatorMessages.packSubnetConversionData(scd);
        uint256 length = 92 + scd.initialValidators.length * 88;

        require(packedOld.length == length, "Old Packed data length does not match");
        require(packedNew.length == length, "New Packed data length does not match");
        require(sha256(packedOld) == sha256(packedNew), "Packed data does not match");
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

    // solhint-disable-next-line no-empty-blocks
    function _beforeSend(uint256 amount, address spender) internal virtual override {}
}
