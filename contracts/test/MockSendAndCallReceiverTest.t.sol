// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {MockERC20SendAndCallReceiver} from "../src/mocks/MockERC20SendAndCallReceiver.sol";
import {MockNativeSendAndCallReceiver} from "../src/mocks/MockNativeSendAndCallReceiver.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";

contract MockERC20SendAndCallReceiverTest is Test {
    IERC20 public erc20;
    MockERC20SendAndCallReceiver public erc20SendAndCallReceiver;

    bytes32 public constant DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address internal _originSenderAddress = vm.addr(0x1);
    address internal _originBridgeAddress = vm.addr(0x2);

    event TokensReceived(
        bytes32 indexed sourceBlockchainID,
        address indexed originBridgeAddress,
        address indexed originSenderAddress,
        address token,
        uint256 amount,
        bytes payload
    );

    function setUp() public virtual {
        erc20 = new ExampleERC20();
        erc20SendAndCallReceiver = new MockERC20SendAndCallReceiver();
    }

    function testRevert() public {
        vm.expectRevert("MockERC20SendAndCallReceiver: empty payload");
        erc20SendAndCallReceiver.receiveTokens({
            sourceBlockchainID: bytes32(0),
            originBridgeAddress: address(this),
            originSenderAddress: address(this),
            token: address(erc20),
            amount: 10,
            payload: new bytes(0)
        });
    }

    function testSuccess() public {
        uint256 amount = 100;
        bytes memory payload = hex"9876543210";
        erc20.approve(address(erc20SendAndCallReceiver), amount);
        vm.expectEmit(true, true, true, true, address(erc20SendAndCallReceiver));
        emit TokensReceived({
            sourceBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            originBridgeAddress: _originBridgeAddress,
            originSenderAddress: _originSenderAddress,
            token: address(erc20),
            amount: amount,
            payload: payload
        });
        erc20SendAndCallReceiver.receiveTokens({
            sourceBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            originBridgeAddress: _originBridgeAddress,
            originSenderAddress: _originSenderAddress,
            token: address(erc20),
            amount: amount,
            payload: payload
        });
        assertEq(erc20.balanceOf(address(erc20SendAndCallReceiver)), amount);
    }

    function testReceiveSenderCheck() public {
        uint256 amount = 100;
        bytes memory payload = hex"9876543210";
        erc20.approve(address(erc20SendAndCallReceiver), amount);

        erc20SendAndCallReceiver.blockSender(DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, _originSenderAddress);
        assertTrue(
            erc20SendAndCallReceiver.blockedSenders(
                DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, _originSenderAddress
            )
        );

        vm.expectRevert("MockERC20SendAndCallReceiver: sender blocked");
        erc20SendAndCallReceiver.receiveTokens({
            sourceBlockchainID: DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            originBridgeAddress: _originBridgeAddress,
            originSenderAddress: _originSenderAddress,
            token: address(erc20),
            amount: amount,
            payload: payload
        });
        assertEq(erc20.balanceOf(address(erc20SendAndCallReceiver)), 0);
    }
}

contract MockNativeSendAndCallReceiverTest is Test {
    MockNativeSendAndCallReceiver public nativeSendAndCallReceiver;

    bytes32 public constant DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address internal _originSenderAddress = vm.addr(0x1);
    address internal _originBridgeAddress = vm.addr(0x2);

    event TokensReceived(
        bytes32 indexed sourceBlockchainID,
        address indexed originBridgeAddress,
        address indexed originSenderAddress,
        uint256 amount,
        bytes payload
    );

    function setUp() public virtual {
        nativeSendAndCallReceiver = new MockNativeSendAndCallReceiver();
    }

    function testRevert() public {
        vm.expectRevert("MockNativeSendAndCallReceiver: empty payload");
        nativeSendAndCallReceiver.receiveTokens(
            bytes32(0), address(this), address(this), new bytes(0)
        );
    }

    function testSuccess() public {
        uint256 amount = 10;
        bytes memory payload = hex"1234567890";
        vm.expectEmit(true, true, true, true, address(nativeSendAndCallReceiver));
        emit TokensReceived(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID,
            _originBridgeAddress,
            _originSenderAddress,
            amount,
            payload
        );
        nativeSendAndCallReceiver.receiveTokens{value: amount}(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, _originBridgeAddress, _originSenderAddress, payload
        );
        assertEq(address(nativeSendAndCallReceiver).balance, amount);
    }

    function testReceiveSenderCheck() public {
        uint256 amount = 10;
        bytes memory payload = hex"1234567890";

        nativeSendAndCallReceiver.blockSender(DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, _originSenderAddress);
        assertTrue(
            nativeSendAndCallReceiver.blockedSenders(
                DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, _originSenderAddress
            )
        );

        vm.expectRevert("MockNativeSendAndCallReceiver: sender blocked");
        nativeSendAndCallReceiver.receiveTokens{value: amount}(
            DEFAULT_TOKEN_HOME_BLOCKCHAIN_ID, _originBridgeAddress, _originSenderAddress, payload
        );
        assertEq(address(nativeSendAndCallReceiver).balance, 0);
    }
}
