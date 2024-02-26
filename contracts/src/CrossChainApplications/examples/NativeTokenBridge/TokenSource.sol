// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ITokenSource} from "./ITokenSource.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Implementation of the {ITokenSource} interface.
 *
 * This abstract contract represents the shared functionality for source contracts.
 */
abstract contract TokenSource is ITokenSource, TeleporterOwnerUpgradeable {
    uint256 public constant MINT_NATIVE_TOKENS_REQUIRED_GAS = 100_000;
    // Used to keep track of tokens burned through transactions on the destination chain. The destination burned amount can
    // be reported to this contract to burn an equivalent number of tokens on this chain.
    bytes32 public immutable destinationBlockchainID;
    address public immutable nativeTokenDestinationAddress;

    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 destinationBlockchainID_,
        address nativeTokenDestinationAddress_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress, teleporterManager) {
        require(
            destinationBlockchainID_ != bytes32(0), "TokenSource: zero destination blockchain ID"
        );
        require(
            destinationBlockchainID_
                != IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID(),
            "TokenSource: cannot bridge with same blockchain"
        );
        destinationBlockchainID = destinationBlockchainID_;

        require(
            nativeTokenDestinationAddress_ != address(0),
            "TokenSource: zero destination contract address"
        );
        nativeTokenDestinationAddress = nativeTokenDestinationAddress_;
    }

    /**
     * @dev See {TeleporterUpgradeable-receiveTeleporterMessage}.
     *
     * Receives a Teleporter message and routes to the appropriate internal function call.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        // Only allow messages from the destination chain.
        require(
            sourceBlockchainID == destinationBlockchainID, "TokenSource: invalid destination chain"
        );

        // Only allow the partner contract to send messages.
        require(
            originSenderAddress == nativeTokenDestinationAddress, "TokenSource: unauthorized sender"
        );

        (address recipient, uint256 amount) = abi.decode(message, (address, uint256));
        _unlockTokens(recipient, amount);
    }

    /**
     * @dev Unlocks tokens to recipient.
     */
    function _unlockTokens(address recipient, uint256 amount) internal virtual;
}
