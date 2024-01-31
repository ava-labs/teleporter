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
    // Designated Blackhole Address for this contract. Tokens are sent here to be "burned" when
    // a SourceAction.Burn message is received from the destination chain.
    address public constant BURN_ADDRESS = 0x0100000000000000000000000000000000010203;
    uint256 public constant MINT_NATIVE_TOKENS_REQUIRED_GAS = 100_000;
    // Used to keep track of tokens burned through transactions on the destination chain. The destination burned amount can
    // be reported to this contract to burn an equivalent number of tokens on this chain.
    uint256 public destinationBurnedTotal;
    bytes32 public immutable destinationBlockchainID;
    address public immutable nativeTokenDestinationAddress;

    constructor(
        address teleporterRegistryAddress,
        bytes32 destinationBlockchainID_,
        address nativeTokenDestinationAddress_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {
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

        // Decode the payload to recover the action and corresponding function parameters
        (SourceAction action, bytes memory actionData) = abi.decode(message, (SourceAction, bytes));

        // Route to the appropriate function.
        if (action == SourceAction.Unlock) {
            (address recipient, uint256 amount) = abi.decode(actionData, (address, uint256));
            _unlockTokens(recipient, amount);
        } else if (action == SourceAction.Burn) {
            uint256 newBurnTotal = abi.decode(actionData, (uint256));
            _handleBurnTokens(newBurnTotal);
        } else {
            revert("TokenSource: invalid action");
        }
    }

    /**
     * @dev Unlocks tokens to recipient.
     */
    function _unlockTokens(address recipient, uint256 amount) internal virtual;

    /**
     * @dev Update destinationBurnedTotal sent from destination chain
     * If the new burned total is less than the highest known burned total, this transaction is a no-op.
     * The burned total on the destination will only ever increase, but new totals may be relayed to this
     * chain out of order.
     */
    function _handleBurnTokens(uint256 newBurnTotal) internal virtual;
}
