// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessengerTest, TeleporterMessageReceipt} from "./TeleporterMessengerTest.t.sol";

contract GetRelayerRewardAddressTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testSuccess() public {
        // Before receiving the message, it returns the 0 address.
        uint256 mockNonce = 8;
        bytes32 mockMessageID = _createMessageID(DEFAULT_DESTINATION_BLOCKCHAIN_ID, mockNonce);
        assertEq(teleporterMessenger.getRelayerRewardAddress(mockMessageID), address(0));

        // Mock receiving the message
        address relayerRewardAddress = 0xCAFebAbeDc0D4D7B7EEdCf61eb863fF413BB6234;
        _receiveTestMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            mockNonce,
            relayerRewardAddress,
            new TeleporterMessageReceipt[](0)
        );

        // Now it has the relayer reward address.
        assertEq(teleporterMessenger.getRelayerRewardAddress(mockMessageID), relayerRewardAddress);
    }
}
