// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import "./ITeleporterMessenger.sol";

/**
 * @dev ReceiptQueue is a convenience contract that creates a queue-like interface of
 * TeleporterMessageReceipt structs. It provides FIFO properties.
 */
contract ReceiptQueue {
    address public immutable owner;
    mapping(uint256 => TeleporterMessageReceipt) public queue;
    uint256 public first = 0;
    uint256 public last = 0;

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

    constructor() {
        owner = msg.sender;
    }

    /**
     * @dev Adds a receipt to the queue.
     *
     * Requirements:
     *
     * - `msg.sender` must be the owner.
     */
    function enqueue(TeleporterMessageReceipt calldata receipt) external {
        require(msg.sender == owner, "Unauthorized.");
        queue[last++] = receipt;

        emit Enqueue(receipt.receivedMessageID, receipt.relayerRewardAddress);
    }

    /**
     * @dev Removes the oldest open receipt from the queue.
     *
     * Requirements:
     *
     * - `msg.sender` must be the owner.
     * - The queue must be non-empty.
     */
    function dequeue()
        external
        returns (TeleporterMessageReceipt memory result)
    {
        require(msg.sender == owner, "Unauthorized.");
        uint256 first_ = first;
        require(last > first_, "Empty queue."); // non-empty queue

        result = queue[first_];

        delete queue[first_];
        first = first_ + 1;

        emit Dequeue(result.receivedMessageID, result.relayerRewardAddress);
    }

    /**
     * @dev Returns the number of open receipts in the queue.
     */
    function size() external view returns (uint256) {
        return last - first;
    }
}
