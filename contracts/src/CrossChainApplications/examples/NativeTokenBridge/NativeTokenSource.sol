// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Address} from "@openzeppelin/contracts@4.8.1/utils/Address.sol";
import {INativeTokenSource} from "./INativeTokenSource.sol";
import {TokenSource} from "./TokenSource.sol";
import {
    ITeleporterMessenger,
    TeleporterFeeInfo,
    TeleporterMessageInput
} from "@teleporter/ITeleporterMessenger.sol";
import {SafeERC20TransferFrom} from "@teleporter/SafeERC20TransferFrom.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Implementation of the {TokenSource} abstract contract.
 *
 * This contracts implements {TokenSource} and uses native tokens as the currency.
 */
contract NativeTokenSource is INativeTokenSource, TokenSource {
    constructor(
        address teleporterRegistryAddress,
        bytes32 destinationBlockchainID_,
        address nativeTokenDestinationAddress_
    )
        TokenSource(teleporterRegistryAddress, destinationBlockchainID_, nativeTokenDestinationAddress_)
    {}

    /**
     * @dev See {INativeTokenSource-transferToDestination}.
     */
    function transferToDestination(
        address recipient,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external payable nonReentrant {
        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();

        // The recipient cannot be the zero address.
        require(recipient != address(0), "NativeTokenSource: zero recipient address");

        // Lock tokens in this bridge instance. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedFeeAmount;
        if (feeInfo.amount > 0) {
            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(feeInfo.feeTokenAddress), feeInfo.amount
            );
            SafeERC20.safeIncreaseAllowance(
                IERC20(feeInfo.feeTokenAddress), address(teleporterMessenger), adjustedFeeAmount
            );
        }

        bytes32 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: nativeTokenDestinationAddress,
                feeInfo: feeInfo,
                requiredGasLimit: MINT_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(recipient, msg.value)
            })
        );

        emit TransferToDestination({
            sender: msg.sender,
            recipient: recipient,
            amount: msg.value,
            teleporterMessageID: messageID
        });
    }

    /**
     * @dev See {TokenSource-_unlockTokens}
     */
    function _unlockTokens(address recipient, uint256 amount) internal override {
        require(recipient != address(0), "NativeTokenSource: zero recipient address");
        require(address(this).balance >= amount, "NativeTokenSource: insufficient collateral");

        // Transfer to recipient
        emit UnlockTokens(recipient, amount);
        Address.sendValue(payable(recipient), amount);
    }

    /**
     * @dev See {TokenSource-_handleBurnTokens}
     */
    function _handleBurnTokens(uint256 newBurnTotal) internal override {
        if (newBurnTotal > destinationBurnedTotal) {
            uint256 difference = newBurnTotal - destinationBurnedTotal;
            _burnTokens(difference);
            destinationBurnedTotal = newBurnTotal;
        }
    }

    /**
     * @dev See {TokenSource-_burnTokens}
     */
    function _burnTokens(uint256 amount) private {
        emit BurnTokens(amount);
        Address.sendValue(payable(BURN_ADDRESS), amount);
    }
}
