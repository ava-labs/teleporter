// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {ITeleporterTokenBridge, SendTokensInput} from "./interfaces/ITeleporterTokenBridge.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @title TeleporterTokenDestination
 * @dev Abstract contract for a Teleporter token bridge that receives tokens from a {TeleporterTokenSource} in exchange for the tokens of this token bridge instance.
 *
 * @custom:security-contact https://github.com/ava-labs/teleporter-token-bridge/blob/main/SECURITY.md
 */
abstract contract TeleporterTokenDestination is
    ITeleporterTokenBridge,
    TeleporterOwnerUpgradeable
{
    /// @notice The blockchain ID of the chain this contract is deployed on.
    bytes32 public immutable blockchainID;

    /// @notice The blockchain ID of the source chain this contract receives tokens from.
    bytes32 public immutable sourceBlockchainID;
    /// @notice The address of the source token bridge instance this contract receives tokens from.
    address public immutable tokenSourceAddress;
    /// @notice The ERC20 token this contract uses to pay for Teleporter fees.
    address public immutable feeTokenAddress;

    /**
     * @notice Fixed gas cost for performing a multihop transfer on the `sourceBlockchainID`
     * , before forwarding to the final destination bridge instance.
     */
    uint256 public constant MULTIHOP_REQUIRED_GAS = 220_000;

    /**
     * @notice Initializes this destination token bridge instance to receive
     * tokens from the specified source blockchain and token bridge instance.
     */
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID_,
        address tokenSourceAddress_,
        address feeTokenAddress_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress, teleporterManager) {
        blockchainID = IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID();
        require(
            sourceBlockchainID_ != bytes32(0),
            "TeleporterTokenDestination: zero source blockchain ID"
        );
        require(
            sourceBlockchainID_ != blockchainID,
            "TeleporterTokenDestination: cannot deploy to same blockchain as source"
        );
        require(
            tokenSourceAddress_ != address(0),
            "TeleporterTokenDestination: zero token source address"
        );
        require(
            feeTokenAddress_ != address(0), "TeleporterTokenDestination: zero fee token address"
        );
        sourceBlockchainID = sourceBlockchainID_;
        tokenSourceAddress = tokenSourceAddress_;
        feeTokenAddress = feeTokenAddress_;
    }

    /**
     * @notice Sends tokens to the specified destination token bridge instance.
     *
     * @dev Burns the bridged amount, and uses Teleporter to send a cross chain message.
     * Tokens can be sent to the same blockchain this bridge instance is deployed on,
     * to another destination bridge instance.
     * Requirements:
     *
     * - `input.destinationBridgeAddress` cannot be the zero address
     * - `input.recipient` cannot be the zero address
     * - `amount` must be greater than 0
     * - `amount` must be greater than `input.primaryFee`
     */
    function _send(SendTokensInput calldata input, uint256 amount) internal virtual {
        require(
            input.destinationBlockchainID != bytes32(0),
            "TeleporterTokenDestination: zero destination blockchain ID"
        );
        require(
            input.destinationBridgeAddress != address(0),
            "TeleporterTokenDestination: zero destination bridge address"
        );
        require(input.recipient != address(0), "TeleporterTokenDestination: zero recipient address");

        amount = _deposit(amount);
        require(
            amount > input.primaryFee + input.secondaryFee,
            "TeleporterTokenDestination: insufficient amount to cover fees"
        );

        amount -= input.primaryFee;
        _burn(amount);

        uint256 scaledAmount = _scaleTokens(amount, false);
        require(scaledAmount > 0, "NativeTokenDestination: insufficient tokens to transfer");

        // If the destination blockchain is the source blockchain,
        // no multihop is needed. Only the required gas limit for the Teleporter message back to
        // `sourceBlockchainID` is needed, which is provided by `input.requiredGasLimit`.
        // Else, there will be a multihop transfer to the final destination.
        // The first hop back to `sourceBlockchainID` requires `MULTIHOP_REQUIRED_GAS`,
        // and the second hop to the final destination requires `input.requiredGasLimit`.
        uint256 firstHopRequiredGas = input.requiredGasLimit;
        uint256 secondHopRequiredGas;
        if (input.destinationBlockchainID == sourceBlockchainID) {
            // If the destination blockchain is the source bridge instance's blockchain,
            // the destination bridge address must match the token source address,
            // and no secondary fee is needed.
            require(
                input.destinationBridgeAddress == tokenSourceAddress,
                "TeleporterTokenDestination: invalid destination bridge address"
            );
            require(input.secondaryFee == 0, "TeleporterTokenDestination: non-zero secondary fee");
        } else {
            // Do not allow bridging to the same token bridge instance.
            if (input.destinationBlockchainID == blockchainID) {
                require(
                    input.destinationBridgeAddress != address(this),
                    "TeleporterTokenDestination: invalid destination bridge address"
                );
            }
            firstHopRequiredGas = MULTIHOP_REQUIRED_GAS;
            secondHopRequiredGas = input.requiredGasLimit;
        }

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: tokenSourceAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeTokenAddress, amount: input.primaryFee}),
                requiredGasLimit: firstHopRequiredGas,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(
                    SendTokensInput({
                        destinationBlockchainID: input.destinationBlockchainID,
                        destinationBridgeAddress: input.destinationBridgeAddress,
                        recipient: input.recipient,
                        primaryFee: input.secondaryFee,
                        secondaryFee: 0,
                        requiredGasLimit: secondHopRequiredGas
                    }),
                    scaledAmount
                    )
            })
        );

        emit SendTokens(messageID, msg.sender, input, amount);
    }

    /**
     * @notice Verifies the source token bridge instance, and withdraws the amount to the recipient address.
     *
     * @dev See {ITeleporterUpgradeable-_receiveTeleporterMessage}
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID_,
        address originSenderAddress,
        bytes memory message
    ) internal virtual override {
        require(
            sourceBlockchainID_ == sourceBlockchainID,
            "TeleporterTokenDestination: invalid source chain"
        );
        require(
            originSenderAddress == tokenSourceAddress,
            "TeleporterTokenDestination: invalid token source address"
        );
        (address recipient, uint256 amount) = abi.decode(message, (address, uint256));

        emit WithdrawTokens(recipient, amount);
        _withdraw(recipient, amount);
    }

    /**
     * @notice Deposits tokens from the sender to this contract,
     * and returns the adjusted amount of tokens deposited.
     * @param amount is initial amount sent to this contract.
     * @return The actual amount deposited to this contract.
     */
    function _deposit(uint256 amount) internal virtual returns (uint256);

    /**
     * @notice Withdraws tokens to the recipient address.
     * @param recipient The address to withdraw tokens to
     * @param amount The amount of tokens to withdraw
     */
    function _withdraw(address recipient, uint256 amount) internal virtual;

    /**
     * @notice Burns a fee adjusted amount of tokens that the user
     * has deposited to this token bridge instance.
     * @param amount The amount of tokens to burn
     */
    function _burn(uint256 amount) internal virtual;

    /**
     * @dev Scales `value` based on `tokenMultiplier` and the direction of the transfer.
     * Should be used for all tokens being transferred to/from other subnets.
     */
    function _scaleTokens(uint256 value, bool isReceive) internal view virtual returns (uint256);
}
