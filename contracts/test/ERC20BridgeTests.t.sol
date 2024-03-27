// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {IERC20Bridge} from "../src/interfaces/IERC20Bridge.sol";
import {SendTokensInput} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

abstract contract ERC20BridgeTest is TeleporterTokenBridgeTest {
    using SafeERC20 for IERC20;

    IERC20Bridge public erc20Bridge;

    function testZeroSendAmount() public {
        vm.expectRevert("SafeERC20TransferFrom: balance not increased");
        _send(_createDefaultSendTokensInput(), 0);
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        erc20Bridge.send(input, amount);
    }

    function _setUpExpectedDeposit(uint256 amount) internal virtual override {
        // Increase the allowance of the bridge to transfer the funds from the user
        feeToken.safeIncreaseAllowance(address(tokenBridge), amount);
        // Check that transferFrom is called to deposit the funds sent from the user to the bridge
        vm.expectCall(
            address(feeToken),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenBridge), amount))
        );
        vm.expectEmit(true, true, true, true, address(feeToken));
        emit Transfer(address(this), address(tokenBridge), amount);
    }
}
