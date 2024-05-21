// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {IERC20TokenBridge} from "../src/interfaces/IERC20TokenBridge.sol";
import {SendTokensInput, SendAndCallInput} from "../src/interfaces/ITokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

abstract contract ERC20BridgeTest is TeleporterTokenBridgeTest {
    using SafeERC20 for IERC20;

    IERC20TokenBridge public erc20Bridge;

    function testZeroSendAmount() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        _setUpRegisteredDestination(
            input.destinationBlockchainID, input.destinationBridgeAddress, 0
        );
        _setUpExpectedZeroAmountRevert();
        _send(input, 0);
    }

    function _send(SendTokensInput memory input, uint256 amount) internal virtual override {
        erc20Bridge.send(input, amount);
    }

    function _sendAndCall(
        SendAndCallInput memory input,
        uint256 amount
    ) internal virtual override {
        erc20Bridge.sendAndCall(input, amount);
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal virtual override {
        // Transfer the fee to the bridge if it is greater than 0
        if (feeAmount > 0) {
            bridgedToken.safeIncreaseAllowance(address(tokenBridge), feeAmount);
            if (address(bridgedToken) != address(tokenBridge)) {
                vm.expectCall(
                    address(bridgedToken),
                    abi.encodeCall(
                        IERC20.transferFrom, (address(this), address(tokenBridge), feeAmount)
                    )
                );
            }
        }
        // Increase the allowance of the bridge to transfer the funds from the user
        bridgedToken.safeIncreaseAllowance(address(tokenBridge), amount);

        // Check that transferFrom is called to deposit the funds sent from the user to the bridge
        // This is only the case if the bridge is not the fee token itself, in which case this is an internal call.
        if (address(bridgedToken) != address(tokenBridge)) {
            vm.expectCall(
                address(bridgedToken),
                abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenBridge), amount))
            );
        }
        vm.expectEmit(true, true, true, true, address(bridgedToken));
        emit Transfer(address(this), address(tokenBridge), amount);
    }
}
