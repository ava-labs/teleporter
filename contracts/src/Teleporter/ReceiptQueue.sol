// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./ITeleporterMessenger.sol";

/**
 * @dev ReceiptQueue is a convenience contract that creates a queue-like interface of
 * TeleporterMessageReceipt structs. It provides FIFO properties.
 */
library ReceiptQueue {
    struct TeleporterMessageReceiptQueue {
        address owner;
        uint256 begin;
        uint256 end;
        mapping(uint256 index => TeleporterMessageReceipt) data;
    }

    /**
     * @dev Emitted when a new receipt is added to the queue.
     */
    event Enqueue(
        uint256 indexed messageID,
        address indexed relayerRewardAddress
    );

    /**
     * @dev Emitted when a receipt is taken off the front of the queue.
     */
    event Dequeue(
        uint256 indexed messageID,
        address indexed relayerRewardAddress
    );

    // Errors
    error Unauthorized();
    error EmptyQueue();

    /**
     * @dev Adds a receipt to the queue.
     *
     * Requirements:
     *
     * - `msg.sender` must be the owner.
     */
    function enqueue(TeleporterMessageReceiptQueue storage queue, TeleporterMessageReceipt memory receipt) internal {
        unchecked {
            if (queue.owner != msg.sender) {
                revert Unauthorized();
            }

            uint256 backIndex = queue.end;
            queue.data[backIndex] = receipt;
            queue.end = backIndex + 1;

            emit Enqueue(receipt.receivedMessageID, receipt.relayerRewardAddress);
        }
    }

    /**
     * @dev Removes the oldest open receipt from the queue.
     *
     * Requirements:
     *
     * - `msg.sender` must be the owner.
     * - The queue must be non-empty.
     */
    function dequeue(TeleporterMessageReceiptQueue storage queue) internal returns (TeleporterMessageReceipt memory result) {
        unchecked {
            if (queue.owner != msg.sender) {
                revert Unauthorized();
            }

            uint256 frontIndex = queue.begin;
            if (frontIndex == queue.end) revert EmptyQueue();
            result = queue.data[frontIndex];
            delete queue.data[frontIndex];
            queue.begin = frontIndex + 1;

            emit Dequeue(result.receivedMessageID, result.relayerRewardAddress);
        }
    }

    /**
     * @dev Returns the number of open receipts in the queue.
     */
    function size(TeleporterMessageReceiptQueue storage queue) internal view returns (uint256) {
        return queue.end - queue.begin;
    }
}
