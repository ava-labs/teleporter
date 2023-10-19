// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./ITeleporterMessenger.sol";

/**
 * @dev ReceiptQueue is a convenience library that creates a queue-like interface of
 * TeleporterMessageReceipt structs. It provides FIFO properties.
 * Note: All functions in this library are internal so that the library is not deployed as a contract.
 */
library ReceiptQueue {
    struct TeleporterMessageReceiptQueue {
        uint256 first;
        uint256 last;
        mapping(uint256 index => TeleporterMessageReceipt) data;
    }

    // The maximum number of receipts to include in a single message.
    uint256 private constant _MAXIMUM_RECEIPT_COUNT = 5;

    /**
     * @dev Adds a receipt to the queue.
     */
    function enqueue(
        TeleporterMessageReceiptQueue storage queue,
        TeleporterMessageReceipt memory receipt
    ) internal {
        queue.data[queue.last++] = receipt;
    }

    /**
     * @dev Removes the oldest open receipt from the queue.
     *
     * Requirements:
     * - The queue must be non-empty.
     */
    function dequeue(
        TeleporterMessageReceiptQueue storage queue
    ) internal returns (TeleporterMessageReceipt memory result) {
        uint256 first_ = queue.first;
        require(queue.last != first_, "ReceiptQueue: empty queue");
        result = queue.data[first_];
        delete queue.data[first_];
        queue.first = first_ + 1;
    }

    /**
     * @dev Returns the outstanding receipts for the given chain ID that should be included in the next message sent.
     */
    function getOutstandingReceiptsToSend(
        TeleporterMessageReceiptQueue storage queue
    ) internal returns (TeleporterMessageReceipt[] memory result) {
        // Get the current outstanding receipts for the given chain ID.
        // If the queue contract doesn't exist, there are no outstanding receipts to send.
        uint256 resultSize = size(queue);
        if (resultSize == 0) {
            return new TeleporterMessageReceipt[](0);
        }

        // Calculate the result size as the minimum of the number of receipts and maximum batch size.
        if (resultSize > _MAXIMUM_RECEIPT_COUNT) {
            resultSize = _MAXIMUM_RECEIPT_COUNT;
        }

        result = new TeleporterMessageReceipt[](resultSize);
        for (uint256 i = 0; i < resultSize; ++i) {
            result[i] = dequeue(queue);
        }
    }

    /**
     * @dev Returns the number of open receipts in the queue.
     */
    function size(
        TeleporterMessageReceiptQueue storage queue
    ) internal view returns (uint256) {
        return queue.last - queue.first;
    }

    /**
     * @dev Returns the receipt at the given index in the queue.
     */
    function getReceiptAtIndex(
        TeleporterMessageReceiptQueue storage queue,
        uint256 index
    ) internal view returns (TeleporterMessageReceipt memory) {
        require(index < size(queue), "ReceiptQueue: index out of bounds");
        return queue.data[queue.first + index];
    }
}
