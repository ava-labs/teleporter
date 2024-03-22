// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20BridgeTest} from "./IERC20BridgeTests.t.sol";
import {TeleporterTokenSourceTest} from "./TeleporterTokenSourceTests.t.sol";
import {ITeleporterTokenBridgeTest} from "./ITeleporterTokenBridgeTests.t.sol";
import {ERC20Source} from "../src/ERC20Source.sol";
import {
    ITeleporterTokenBridge, SendTokensInput
} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";

contract ERC20SourceTest is IERC20BridgeTest, TeleporterTokenSourceTest {
    using SafeERC20 for IERC20;

    ERC20Source public app;
    IERC20 public mockERC20;

    function setUp() public override {
        TeleporterTokenSourceTest.setUp();

        mockERC20 = new ExampleERC20();
        app = new ERC20Source(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            address(mockERC20)
        );
        erc20Bridge = app;
        tokenSource = app;
        tokenBridge = app;

        feeToken = mockERC20;
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20Source(address(0), address(this), address(mockERC20));
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20Source(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            address(mockERC20)
        );
    }

    function testZeroFeeTokenAddress() public {
        vm.expectRevert(_formatErrorMessage("zero fee token address"));
        new ERC20Source(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(0)
        );
    }

    function _checkWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectCall(
            address(mockERC20), abi.encodeCall(IERC20.transfer, (address(recipient), amount))
        );
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Transfer(address(app), recipient, amount);
    }
}
