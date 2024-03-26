// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenSourceTest} from "./TeleporterTokenSourceTests.t.sol";
import {INativeTokenBridgeTest} from "./INativeTokenBridgeTests.t.sol";
import {NativeTokenSource} from "../src/NativeTokenSource.sol";
import {IWrappedNativeToken} from "../src/interfaces/IWrappedNativeToken.sol";
import {ExampleWAVAX} from "../src/mocks/ExampleWAVAX.sol";

contract NativeTokenSourceTest is INativeTokenBridgeTest, TeleporterTokenSourceTest {
    NativeTokenSource public app;
    IWrappedNativeToken public mockWrappedToken;

    function setUp() public override {
        TeleporterTokenSourceTest.setUp();

        mockWrappedToken = new ExampleWAVAX();
        app = new NativeTokenSource(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            address(mockWrappedToken)
        );
        tokenSource = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        feeToken = mockWrappedToken;
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new NativeTokenSource(address(0), address(this), address(mockWrappedToken));
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new NativeTokenSource(MOCK_TELEPORTER_REGISTRY_ADDRESS, address(0), address(mockWrappedToken));
    }

    function testZeroFeeTokenAddress() public {
        vm.expectRevert(_formatErrorMessage("zero fee token address"));
        new NativeTokenSource(MOCK_TELEPORTER_REGISTRY_ADDRESS, address(this), address(0));
    }

    function _checkWithdrawal(address, uint256 amount) internal override {
        vm.expectCall(
            address(mockWrappedToken), abi.encodeCall(IWrappedNativeToken.withdraw, (amount))
        );
        vm.expectEmit(true, true, true, true, address(mockWrappedToken));
        emit Withdrawal(address(app), amount);
    }
}
