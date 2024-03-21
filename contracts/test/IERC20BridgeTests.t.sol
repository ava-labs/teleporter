// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenSourceTest} from "./TeleporterTokenSourceTests.t.sol";
import {IERC20Bridge} from "../src/interfaces/IERC20Bridge.sol";
import {SendTokensInput} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

abstract contract IERC20BridgeTest is TeleporterTokenSourceTest {
    using SafeERC20 for IERC20;

    IERC20Bridge public erc20Bridge;
    IERC20 public feeToken;

    // constructor(IERC20Bridge _app, IERC20 _feeToken) {
    //     erc20Bridge = _app;
    //     feeToken = _feeToken;
    // }

    function setUp() public virtual override {
        super.setUp();
    }

    function testZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination bridge address"));
        erc20Bridge.send(input, 0);
    }

    function testZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        vm.expectRevert(_formatTokenSourceErrorMessage("zero recipient address"));
        erc20Bridge.send(input, 0);
    }

    function testZeroSendAmount() public {
        vm.expectRevert("SafeERC20TransferFrom: balance not increased");
        erc20Bridge.send(_createDefaultSendTokensInput(), 0);
    }

    function testInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = 1;
        mockERC20.safeIncreaseAllowance(address(erc20Bridge), input.primaryFee);
        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient amount to cover fees"));
        erc20Bridge.send(input, input.primaryFee);
    }

    function testSendWithFees() public {
        uint256 amount = 2;
        uint256 primaryFee = 1;
        _sendSuccess(amount, primaryFee);
    }

    function testSendNoFees() public {
        uint256 amount = 2;
        uint256 primaryFee = 0;
        _sendSuccess(amount, primaryFee);
    }

    function _sendSuccess(uint256 amount, uint256 feeAmount) internal {
        feeToken.safeIncreaseAllowance(address(erc20Bridge), amount);

        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        // Check that transferFrom is called to deposit the funds sent from the user to the bridge
        vm.expectCall(
            address(feeToken),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(erc20Bridge), amount))
        );

        _checkExpectedTeleporterCalls(input, bridgedAmount);
        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit SendTokens(_MOCK_MESSAGE_ID, address(this), bridgedAmount);
        erc20Bridge.send(input, amount);
    }

    function _checkExpectedTeleporterCalls(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal {
        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: input.primaryFee}),
            requiredGasLimit: _requiredGasLimit(),
            allowedRelayerAddresses: input.allowedRelayerAddresses,
            message: abi.encode(DEFAULT_RECIPIENT_ADDRESS, bridgeAmount)
        });

        if (input.primaryFee > 0) {
            vm.expectCall(
                address(mockERC20),
                abi.encodeCall(
                    IERC20.allowance,
                    (address(erc20Bridge), address(MOCK_TELEPORTER_MESSENGER_ADDRESS))
                )
            );
        }
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );
    }

    function _requiredGasLimit() internal view virtual returns (uint256);
}
