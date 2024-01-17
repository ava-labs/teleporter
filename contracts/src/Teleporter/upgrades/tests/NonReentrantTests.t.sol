// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeable} from "../TeleporterUpgradeable.sol";
import {TeleporterRegistryTest} from "./TeleporterRegistryTests.t.sol";
import {
    ITeleporterMessenger,
    TeleporterMessage,
    TeleporterMessageReceipt
} from "../../ITeleporterMessenger.sol";
import {WarpMessage} from "@subnet-evm-contracts/interfaces/IWarpMessenger.sol";
import {TeleporterMessenger} from "../../TeleporterMessenger.sol";

uint32 constant warpMessageIndex = 2;

contract NonreentrantUpgradeableApp is TeleporterUpgradeable {
    constructor(address teleporterRegistryAddress)
        TeleporterUpgradeable(teleporterRegistryAddress)
    {}

    function setMinTeleporterVersion(uint256 version) public {
        _setMinTeleporterVersion(version);
    }

    function getTeleporterMessenger() public view returns (ITeleporterMessenger) {
        return _getTeleporterMessenger();
    }

    // Calls receiveCrossChainMessage on the latest TeleporterMessenger Contract.
    // The Warp Precompile is mocked to return a message that will call
    // TeleporterUpgradeable.receiveTeleporterMessage which should revert because it is
    // non-reentrant.
    function _receiveTeleporterMessage(bytes32, address, bytes memory) internal override {
        // Call `receiveCrossChainMessage` of the latest version of Teleporter
        getTeleporterMessenger().receiveCrossChainMessage(warpMessageIndex, address(this));
    }

    // solhint-disable-next-line no-empty-blocks
    function _checkTeleporterUpgradeAccess() internal override {}
}

// The flow for the tests below is as follows:
// NonreentrantUpgradeableApp::receiveTeleporterMessage ->
// NonreentrantUpgradeableApp::_receiveTeleporterMessage ->
// TeleporterMessenger::receiveCrossChainMessage ->
// NonreentrantUpgradeableApp::receiveTeleporterMessage
// The last step should revert because receiveTeleporterMessage (contained in
// TeleporterUpgradeable) is non-reentrant.
contract TeleporterUpgradeableTest is TeleporterRegistryTest {
    address public constant DEFAULT_ORIGIN_SENDER_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_DESTINATION_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    bytes public constant DEFAULT_MESSAGE = bytes(hex"1234");
    bytes32 public constant DEFAULT_ORIGIN_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    uint256 public constant DEFAULT_REQUIRED_GAS_LIMIT = 1e6;

    NonreentrantUpgradeableApp public app;

    event MessageExecutionFailed(
        bytes32 indexed messageID, bytes32 indexed originBlockchainID, TeleporterMessage message
    );

    event MessageExecuted(bytes32 indexed messageID, bytes32 indexed originBlockchainID);

    function setUp() public virtual override {
        TeleporterRegistryTest.setUp();
        TeleporterMessenger(teleporterAddress).initializeBlockchainID();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);

        app = new NonreentrantUpgradeableApp(address(teleporterRegistry));
    }

    function testNonreentrantSameTeleporter() public {
        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 987,
            originSenderAddress: DEFAULT_ORIGIN_SENDER_ADDRESS,
            destinationBlockchainID: MOCK_BLOCK_CHAIN_ID,
            destinationAddress: address(app),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE
        });
        WarpMessage memory warpMessage = WarpMessage({
            sourceChainID: DEFAULT_ORIGIN_BLOCKCHAIN_ID,
            originSenderAddress: teleporterAddress,
            payload: abi.encode(messageToReceive)
        });

        // Same index as in NonreentrantUpgradeableApp._receiveTeleporterMessage()
        _mockGetVerifiedWarpMessage(warpMessageIndex, warpMessage, true);

        bytes32 messageID = TeleporterMessenger(teleporterAddress).calculateMessageID(
            DEFAULT_ORIGIN_BLOCKCHAIN_ID, MOCK_BLOCK_CHAIN_ID, 987
        );

        vm.expectEmit(true, true, true, true, address(teleporterAddress));
        emit MessageExecutionFailed(messageId, DEFAULT_ORIGIN_BLOCKCHAIN_ID, messageToReceive);

        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_BLOCKCHAIN_ID, DEFAULT_ORIGIN_SENDER_ADDRESS, DEFAULT_MESSAGE
        );
    }

    function testNonreentrantDifferentTeleporter() public {
        TeleporterMessenger teleporterV2 = new TeleporterMessenger();
        teleporterV2.initializeBlockchainID();
        _addProtocolVersion(teleporterRegistry, address(teleporterV2));

        TeleporterMessage memory messageToReceive = TeleporterMessage({
            messageNonce: 987,
            originSenderAddress: DEFAULT_ORIGIN_SENDER_ADDRESS,
            destinationBlockchainID: MOCK_BLOCK_CHAIN_ID,
            destinationAddress: address(app),
            requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: DEFAULT_MESSAGE
        });
        WarpMessage memory warpMessage = WarpMessage({
            sourceChainID: DEFAULT_ORIGIN_BLOCKCHAIN_ID,
            originSenderAddress: address(teleporterV2),
            payload: abi.encode(messageToReceive)
        });

        // Same index as in NonreentrantUpgradeableApp._receiveTeleporterMessage()
        _mockGetVerifiedWarpMessage(warpMessageIndex, warpMessage, true);

        vm.expectCall(
            address(teleporterV2),
            abi.encodeCall(
                ITeleporterMessenger.receiveCrossChainMessage, (warpMessageIndex, address(app))
            )
        );

        bytes32 messageId = TeleporterMessenger(teleporterV2).calculateMessageID(
            DEFAULT_ORIGIN_BLOCKCHAIN_ID, MOCK_BLOCK_CHAIN_ID, messageToReceive.messageNonce
        );

        vm.expectEmit(true, true, true, true, address(teleporterV2));
        emit MessageExecutionFailed(messageId, DEFAULT_ORIGIN_BLOCKCHAIN_ID, messageToReceive);

        vm.prank(teleporterAddress);
        app.receiveTeleporterMessage(
            DEFAULT_ORIGIN_BLOCKCHAIN_ID, DEFAULT_ORIGIN_SENDER_ADDRESS, DEFAULT_MESSAGE
        );
    }
}
