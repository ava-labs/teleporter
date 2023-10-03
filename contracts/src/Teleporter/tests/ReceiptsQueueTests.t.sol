// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "forge-std/Test.sol";
import "../ReceiptQueue.sol";

contract ReceiptQueueTest is Test {
    using ReceiptQueue for ReceiptQueue.TeleporterMessageReceiptQueue;
    // The state of the contract gets reset before each
    // test is run, with the `setUp()` function being called
    // each time after deployment.
    ReceiptQueue.TeleporterMessageReceiptQueue private _queue;

     // Add 3 elements to the queue.
    TeleporterMessageReceipt private _receipt1 = TeleporterMessageReceipt({
        receivedMessageID: 543,
        relayerRewardAddress: 0x10eB43ef5982628728E3E4bb9F78834f67Fbb40b
    });
    TeleporterMessageReceipt private _receipt2 = TeleporterMessageReceipt({
        receivedMessageID: 684384,
        relayerRewardAddress: 0x10eB43ef5982628728E3E4bb9F78834f67Fbb40b
    });
    TeleporterMessageReceipt private _receipt3 = TeleporterMessageReceipt({
        receivedMessageID: 654351,
        relayerRewardAddress: 0xcC8E718045817AebA89592C72Ae1C9917f5D0894
    });

    function testReceiptQueue() public {
        // Check the initial size is zero.
        assertEq(_queue.size(), 0);

        _queue.enqueue(_receipt1);
        _queue.enqueue(_receipt2);
        _queue.enqueue(_receipt3);

        // Check the size is now three.
        assertEq(_queue.size(), 3);

        // Dequeue 3 elements and check they are given in the correct order (FIFO).
        TeleporterMessageReceipt memory result = _queue.dequeue();
        assertEq(result.receivedMessageID, _receipt1.receivedMessageID);
        assertEq(result.relayerRewardAddress, _receipt1.relayerRewardAddress);
        result = _queue.dequeue();
        assertEq(result.receivedMessageID, _receipt2.receivedMessageID);
        assertEq(result.relayerRewardAddress, _receipt2.relayerRewardAddress);
        result = _queue.dequeue();
        assertEq(result.receivedMessageID, _receipt3.receivedMessageID);
        assertEq(result.relayerRewardAddress, _receipt3.relayerRewardAddress);

        // Check the size is now 0 again.
        assertEq(_queue.size(), 0);

        // Enqueue two more of the same item to check you can have duplicates, followed by the second and third.
        _queue.enqueue(_receipt1);
        _queue.enqueue(_receipt1);
        _queue.enqueue(_receipt2);
        _queue.enqueue(_receipt3);

        // Check the size again.
        assertEq(_queue.size(), 4);

        // Finally dequeue the elements and once again check they are returned in the correct order.
        result = _queue.dequeue();
        assertEq(result.receivedMessageID, _receipt1.receivedMessageID);
        assertEq(result.relayerRewardAddress, _receipt1.relayerRewardAddress);
        result = _queue.dequeue();
        assertEq(result.receivedMessageID, _receipt1.receivedMessageID);
        assertEq(result.relayerRewardAddress, _receipt1.relayerRewardAddress);
        result = _queue.dequeue();
        assertEq(result.receivedMessageID, _receipt2.receivedMessageID);
        assertEq(result.relayerRewardAddress, _receipt2.relayerRewardAddress);
        result = _queue.dequeue();
        assertEq(result.receivedMessageID, _receipt3.receivedMessageID);
        assertEq(result.relayerRewardAddress, _receipt3.relayerRewardAddress);
    }

    function testDequeueEmptyQueue() public {
        // Check that you can't dequeue from empty queue.
        vm.expectRevert(ReceiptQueue.EmptyQueue.selector);
        TeleporterMessageReceipt memory result = _queue.dequeue();
        assertEq(result.receivedMessageID, 0);
        assertEq(result.relayerRewardAddress, address(0));
    }

    function testGetReceiptAtIndex() public {
        _queue.enqueue(_receipt1);
        _queue.enqueue(_receipt1);
        _queue.enqueue(_receipt2);
        _queue.enqueue(_receipt3);

        // Check the size again.
        assertEq(_queue.size(), 4);

        // Check that you can get receipts at specific indexes.
        TeleporterMessageReceipt memory result = _queue.getReceiptAtIndex(2);
        assertEq(result.receivedMessageID, _receipt2.receivedMessageID);
        assertEq(result.relayerRewardAddress, _receipt2.relayerRewardAddress);

        // Check  can't get an out of index element.
        vm.expectRevert(ReceiptQueue.OutofIndex.selector);
        result = _queue.getReceiptAtIndex(4);
    }
}
