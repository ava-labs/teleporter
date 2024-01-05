// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterUpgradeable} from "../TeleporterUpgradeable.sol";
import {TeleporterRegistryTest} from "./TeleporterRegistryTests.t.sol";
import {ITeleporterMessenger} from "../../ITeleporterMessenger.sol";
import {TeleporterMessenger} from "../../TeleporterMessenger.sol";

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

    function _receiveTeleporterMessage(
        // solhint-disable-next-line no-unused-vars
        bytes32 originBlockchainID,
        // solhint-disable-next-line no-unused-vars
        address originSenderAddress,
        // solhint-disable-next-line no-unused-vars
        bytes memory message
    ) internal override {
        // Call `receiveCrossChainMessage` of the latest version of `ModifiedTeleporterMessenger`
        getTeleporterMessenger().receiveCrossChainMessage(1, address(this));
    }

    // solhint-disable-next-line no-empty-blocks
    function _checkTeleporterUpgradeAccess() internal override {}
}

contract ModifiedTeleporterMessenger is TeleporterMessenger {
    bytes32 public constant DEFAULT_ORIGIN_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address public constant DEFAULT_ORIGIN_SENDER_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    bytes public constant DEFAULT_MESSAGE = bytes(hex"1234");

    function receiveCrossChainMessage(
        // solhint-disable-next-line no-unused-vars
        uint32 messageIndex,
        address relayerRewardAddress
    ) external override {
        // Use the relayerRewardAddress to pass in the address of the contract we want to call
        address addressToCall = relayerRewardAddress;

        // Calls the `TeleporterUpgradeable` version of `receiveTeleporterMessage` of the `NonreentrantUpgradeableApp`
        NonreentrantUpgradeableApp(addressToCall).receiveTeleporterMessage(
            DEFAULT_ORIGIN_BLOCKCHAIN_ID, DEFAULT_ORIGIN_SENDER_ADDRESS, DEFAULT_MESSAGE
        );
    }
}

contract TeleporterUpgradeableTest is TeleporterRegistryTest {
    NonreentrantUpgradeableApp public app;
    bytes32 public constant DEFAULT_ORIGIN_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");

    address public constant DEFAULT_ORIGIN_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;

    function setUp() public virtual override {
        TeleporterRegistryTest.setUp();
        _addProtocolVersion(teleporterRegistry, teleporterAddress);
        app = new NonreentrantUpgradeableApp(address(teleporterRegistry));
    }

    function testNonreentrantSameTeleporter() public {
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage("zero teleporter registry address")
        );
        new NonreentrantUpgradeableApp(address(0));
    }

    function testNonreentrantDifferentTeleporter() public {
        teleporterAddress = address(new ModifiedTeleporterMessenger());
        vm.expectRevert(
            _formatTeleporterUpgradeableErrorMessage("zero teleporter registry address")
        );
        new NonreentrantUpgradeableApp(address(0));
    }
}
