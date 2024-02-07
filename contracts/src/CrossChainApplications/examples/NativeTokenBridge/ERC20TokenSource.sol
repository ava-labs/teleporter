// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20TokenSource} from "./IERC20TokenSource.sol";
import {TokenSource} from "./TokenSource.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
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
 * This contracts implements {TokenSource} and uses a specified ERC20 token as the currency.
 */
contract ERC20TokenSource is IERC20TokenSource, TokenSource {
    address public immutable erc20ContractAddress;
    // tokenMultiplier allows this contract to scale the number of tokens it sends/receives to/from
    // the destination chain. This can be used to normalize the number of decimals places between
    // the two subnets.
    uint256 public immutable tokenMultiplier;
    // If multiplyOnSend is true, the number of tokens sent to the destination chain will be multiplied
    // by `tokenMultiplier`, and tokens received from the destination chain will be divided by `tokenMultiplier`.
    // If multiplyOnSend is false, the number of tokens sent to the destination chain will be divided
    // by `tokenMultiplier`, and tokens received from the destination chain will be divided by `tokenMultiplier`.
    bool public immutable multiplyOnSend;

    constructor(
        address teleporterRegistryAddress,
        bytes32 destinationBlockchainID_,
        address nativeTokenDestinationAddress_,
        address erc20ContractAddress_,
        uint256 tokenMultiplier_,
        bool multiplyOnSend_
    )
        TokenSource(teleporterRegistryAddress, destinationBlockchainID_, nativeTokenDestinationAddress_)
    {
        require(
            erc20ContractAddress_ != address(0), "ERC20TokenSource: zero ERC20 contract address"
        );
        erc20ContractAddress = erc20ContractAddress_;

        require(tokenMultiplier_ != 0, "ERC20TokenSource: zero tokenMultiplier");
        tokenMultiplier = tokenMultiplier_;
        multiplyOnSend = multiplyOnSend_;
    }

    /**
     * @dev See {IERC20TokenSource-transferToDestination}.
     */
    function transferToDestination(
        address recipient,
        uint256 totalAmount,
        uint256 feeAmount,
        address[] calldata allowedRelayerAddresses
    ) external nonReentrant {
        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();

        // The recipient cannot be the zero address.
        require(recipient != address(0), "ERC20TokenSource: zero recipient address");

        // Lock tokens in this contract. Supports "fee/burn on transfer" ERC20 token
        // implementations by only bridging the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedAmount =
            SafeERC20TransferFrom.safeTransferFrom(IERC20(erc20ContractAddress), totalAmount);

        // Ensure that the adjusted amount is greater than the fee to be paid.
        require(adjustedAmount > feeAmount, "ERC20TokenSource: insufficient adjusted amount");

        // Allow the Teleporter messenger to spend the fee amount.
        if (feeAmount > 0) {
            SafeERC20.safeIncreaseAllowance(
                IERC20(erc20ContractAddress), address(teleporterMessenger), feeAmount
            );
        }

        uint256 transferAmount = adjustedAmount - feeAmount;

        if (multiplyOnSend) {
            transferAmount *= tokenMultiplier;
        } else {
            transferAmount /= tokenMultiplier;
        }

        bytes32 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: destinationBlockchainID,
                destinationAddress: nativeTokenDestinationAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: erc20ContractAddress, amount: feeAmount}),
                requiredGasLimit: MINT_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(recipient, transferAmount)
            })
        );

        emit TransferToDestination({
            sender: msg.sender,
            recipient: recipient,
            amount: transferAmount,
            teleporterMessageID: messageID
        });
    }

    /**
     * @dev See {TokenSource-_unlockTokens}
     */
    function _unlockTokens(address recipient, uint256 amount) internal override {
        require(recipient != address(0), "ERC20TokenSource: zero recipient address");

        uint256 scaledAmount;
        if (multiplyOnSend) {
            scaledAmount = amount / tokenMultiplier;
        } else {
            scaledAmount = amount * tokenMultiplier;
        }

        // Transfer to recipient
        emit UnlockTokens(recipient, scaledAmount);
        SafeERC20.safeTransfer(IERC20(erc20ContractAddress), recipient, scaledAmount);
    }

    /**
     * @dev See {TokenSource-_handleBurnTokens}
     */
    function _handleBurnTokens(uint256 newBurnTotal) internal override {
        uint256 scaledNewBurnTotal;
        if (multiplyOnSend) {
            scaledNewBurnTotal = newBurnTotal / tokenMultiplier;
        } else {
            scaledNewBurnTotal = newBurnTotal * tokenMultiplier;
        }

        if (scaledNewBurnTotal > destinationBurnedTotal) {
            uint256 difference = scaledNewBurnTotal - destinationBurnedTotal;
            _burnTokens(difference);
            destinationBurnedTotal = scaledNewBurnTotal;
        }
    }

    /**
     * @dev See {TokenSource-_burnTokens}
     */
    function _burnTokens(uint256 amount) private {
        emit BurnTokens(amount);
        SafeERC20.safeTransfer(IERC20(erc20ContractAddress), BURN_ADDRESS, amount);
    }
}
