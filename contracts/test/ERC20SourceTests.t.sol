// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ERC20BridgeTest} from "./ERC20BridgeTests.t.sol";
import {TeleporterTokenSourceTest} from "./TeleporterTokenSourceTests.t.sol";
import {IERC20SendAndCallReceiver} from "../src/interfaces/IERC20SendAndCallReceiver.sol";
import {ERC20Source} from "../src/ERC20Source.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";

contract ERC20SourceTest is ERC20BridgeTest, TeleporterTokenSourceTest {
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

    function testReceiveSendAndCallSuccess() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 200_000;
        _sendSingleHopSendSuccess(amount, 0);

        bytes memory payload = hex"DEADBEEF";

        // Then recipient contract will be approved to spend the tokens from the source bridge contract.
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        // Then a call to the recipient is made.
        bytes memory expectedCalldata = abi.encodeCall(
            IERC20SendAndCallReceiver.receiveTokens, (address(mockERC20), amount, payload)
        );
        vm.etch(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, new bytes(1));
        vm.mockCall(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, expectedCalldata, new bytes(0));
        vm.expectCall(
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            0,
            uint64(DEFAULT_RECIPIENT_GAS_LIMIT),
            expectedCalldata
        );

        // Then recipient contract approval is reset
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, 0);

        // The call should have succeeded.
        vm.expectEmit(true, true, true, true, address(app));
        emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        bytes memory message = _encodeSingleHopCallMessage(
            amount,
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            payload,
            DEFAULT_RECIPIENT_GAS_LIMIT,
            DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailure() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSingleHopSendSuccess(amount, 0);

        bytes memory payload = hex"DEADBEEF";

        // Then recipient contract will be approved to spend the tokens from the source bridge contract.
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        // Then a call to the recipient is made. Mock it reverting
        bytes memory expectedCalldata = abi.encodeCall(
            IERC20SendAndCallReceiver.receiveTokens, (address(mockERC20), amount, payload)
        );
        vm.etch(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, new bytes(1));
        vm.mockCallRevert(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, expectedCalldata, new bytes(0));
        vm.expectCall(
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            0,
            uint64(DEFAULT_RECIPIENT_GAS_LIMIT),
            expectedCalldata
        );

        // Then recipient contract approval is reset
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, 0);

        // The call should have failed.
        vm.expectEmit(true, true, true, true, address(app));
        emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        // The tokens should be transfered to the fallback recipient.
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Transfer(address(app), DEFAULT_FALLBACK_RECIPIENT_ADDRESS, amount);

        bytes memory message = _encodeSingleHopCallMessage(
            amount,
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            payload,
            DEFAULT_RECIPIENT_GAS_LIMIT,
            DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        );

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenSource.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS, message
        );
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectCall(
            address(mockERC20), abi.encodeCall(IERC20.transfer, (address(recipient), amount))
        );
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Transfer(address(app), recipient, amount);
    }
}
