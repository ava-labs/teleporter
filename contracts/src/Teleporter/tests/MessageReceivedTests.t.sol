// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: BSD-3-Clause

pragma solidity 0.8.18;

import "./TeleporterMessengerTest.t.sol";

contract MessageReceivedTest is TeleporterMessengerTest {
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
    }

    function testReceivedMessage() public {
        // Mock receiving a message from another subnet.
        address relayerRewardAddress = 0xA66884fAdC0D4d7B7eedcF61Eb863Ff413bB6234;
        _receiveTestMessage(
            DEFAULT_DESTINATION_CHAIN_ID,
            1,
            relayerRewardAddress,
            new TeleporterMessageReceipt[](0)
        );

        assertEq(
            teleporterMessenger.messageReceived(
                DEFAULT_DESTINATION_CHAIN_ID,
                1
            ),
            true
        );
    }

    function testUnreceivedMessage() public {
        assertEq(
            teleporterMessenger.messageReceived(
                DEFAULT_DESTINATION_CHAIN_ID,
                1
            ),
            false
        );
    }
}
