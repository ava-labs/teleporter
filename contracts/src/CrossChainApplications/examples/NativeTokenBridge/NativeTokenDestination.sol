// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Address} from "@openzeppelin/contracts@4.8.1/utils/Address.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {INativeTokenDestination} from "./INativeTokenDestination.sol";
import {ITokenSource} from "./ITokenSource.sol";
import {
    ITeleporterMessenger,
    TeleporterFeeInfo,
    TeleporterMessageInput
} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {SafeERC20TransferFrom} from "@teleporter/SafeERC20TransferFrom.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
// We need IAllowList as an indirect dependency in order to compile.
// solhint-disable-next-line no-unused-import
import {IAllowList} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IAllowList.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev Implementation of the {INativeTokenDestination} interface.
 *
 * This contract pairs with exactly one `TokenSource` contract on the source chain.
 * It mints and burns native tokens on the destination chain corresponding to locks and unlocks on the source chain.
 */
contract NativeTokenDestination is TeleporterOwnerUpgradeable, INativeTokenDestination {
    // The address where the burned transaction fees are credited.
    // Defined as BLACKHOLE_ADDRESS at
    // https://github.com/ava-labs/subnet-evm/blob/e23ab058d039ff9c8469c89b139d21d52c4bd283/constants/constants.go
    address public constant BURNED_TX_FEES_ADDRESS = 0x0100000000000000000000000000000000000000;
    // Designated Blackhole Address for this contract. Tokens are sent here to be "burned" before
    // sending an unlock message to the source chain. Different from the burned tx fee address so
    // they can be tracked separately.
    address public constant BURN_FOR_TRANSFER_ADDRESS = 0x0100000000000000000000000000000000010203;

    INativeMinter public constant NATIVE_MINTER =
        INativeMinter(0x0200000000000000000000000000000000000001);

    uint256 public constant TRANSFER_NATIVE_TOKENS_REQUIRED_GAS = 100_000;
    uint256 public constant REPORT_BURNED_TOKENS_REQUIRED_GAS = 100_000;
    bytes32 public immutable sourceBlockchainID;
    address public immutable nativeTokenSourceAddress;
    // The first `initialReserveImbalance` tokens sent to this subnet will not be minted.
    // `initialReserveImbalance` should be constructed to match the initial token supply of this subnet.
    // This means tokens will not be minted until the source contact is collateralized.
    uint256 public immutable initialReserveImbalance;
    uint256 public currentReserveImbalance;
    uint256 public totalMinted;

    constructor(
        address teleporterRegistryAddress,
        bytes32 sourceBlockchainID_,
        address nativeTokenSourceAddress_,
        uint256 initialReserveImbalance_
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {
        require(
            sourceBlockchainID_ != bytes32(0), "NativeTokenDestination: zero source blockchain ID"
        );
        require(
            sourceBlockchainID_
                != IWarpMessenger(0x0200000000000000000000000000000000000005).getBlockchainID(),
            "NativeTokenDestination: cannot bridge with same blockchain"
        );
        sourceBlockchainID = sourceBlockchainID_;

        require(
            nativeTokenSourceAddress_ != address(0),
            "NativeTokenDestination: zero source contract address"
        );
        nativeTokenSourceAddress = nativeTokenSourceAddress_;

        require(
            initialReserveImbalance_ != 0, "NativeTokenDestination: zero initial reserve imbalance"
        );

        initialReserveImbalance = initialReserveImbalance_;
        currentReserveImbalance = initialReserveImbalance_;
    }

    /**
     * @dev See {INativeTokenDestination-transferToSource}.
     */
    function transferToSource(
        address recipient,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external payable nonReentrant {
        ITeleporterMessenger teleporterMessenger = teleporterRegistry.getLatestTeleporter();
        // The recipient cannot be the zero address.
        require(recipient != address(0), "NativeTokenDestination: zero recipient address");

        require(
            currentReserveImbalance == 0, "NativeTokenDestination: contract undercollateralized"
        );

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

        // Burn native token by sending to BURN_FOR_TRANSFER_ADDRESS
        Address.sendValue(payable(BURN_FOR_TRANSFER_ADDRESS), msg.value);

        bytes32 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: nativeTokenSourceAddress,
                feeInfo: feeInfo,
                requiredGasLimit: TRANSFER_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(ITokenSource.SourceAction.Unlock, abi.encode(recipient, msg.value))
            })
        );

        emit TransferToSource({
            sender: msg.sender,
            recipient: recipient,
            amount: msg.value,
            teleporterMessageID: messageID
        });
    }

    /**
     * @dev See {INativeTokenDestination-reportTotalBurnedTxFees}.
     */
    function reportTotalBurnedTxFees(
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external {
        ITeleporterMessenger teleporterMessenger = _getTeleporterMessenger();

        uint256 totalBurnedTxFees = BURNED_TX_FEES_ADDRESS.balance;
        bytes32 messageID = teleporterMessenger.sendCrossChainMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: nativeTokenSourceAddress,
                feeInfo: feeInfo,
                requiredGasLimit: REPORT_BURNED_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(ITokenSource.SourceAction.Burn, abi.encode(totalBurnedTxFees))
            })
        );

        emit ReportTotalBurnedTxFees({
            teleporterMessageID: messageID,
            burnAddressBalance: totalBurnedTxFees
        });
    }

    /**
     * @dev See {INativeTokenDestination-isCollateralized}.
     */
    function isCollateralized() external view returns (bool) {
        return currentReserveImbalance == 0;
    }

    /**
     * @dev See {INativeTokenDestination-totalSupply}.
     */
    function totalSupply() external view returns (uint256) {
        uint256 burned = BURNED_TX_FEES_ADDRESS.balance + BURN_FOR_TRANSFER_ADDRESS.balance;
        uint256 created = totalMinted + initialReserveImbalance;

        return created - burned;
    }

    /**
     * @dev See {TeleporterUpgradeable-receiveTeleporterMessage}.
     *
     * Receives a Teleporter message.
     */
    function _receiveTeleporterMessage(
        bytes32 sourceBlockchainID_,
        address originSenderAddress,
        bytes memory message
    ) internal override {
        // Only allow messages from the source chain.
        require(
            sourceBlockchainID_ == sourceBlockchainID,
            "NativeTokenDestination: invalid source chain"
        );

        // Only allow the partner contract to send messages.
        require(
            originSenderAddress == nativeTokenSourceAddress,
            "NativeTokenDestination: unauthorized sender"
        );

        (address recipient, uint256 amount) = abi.decode(message, (address, uint256));
        require(recipient != address(0), "NativeTokenDestination: zero recipient address");
        require(amount != 0, "NativeTokenDestination: zero transfer value");

        // If the contract has not yet been collateralized, we will deduct as many tokens
        // as needed from the transfer as needed. If there are any excess tokens, they will
        // be minted and sent to the recipient.
        uint256 adjustedAmount = amount;
        if (currentReserveImbalance > 0) {
            if (amount > currentReserveImbalance) {
                emit CollateralAdded({amount: currentReserveImbalance, remaining: 0});
                adjustedAmount = amount - currentReserveImbalance;
                currentReserveImbalance = 0;
            } else {
                currentReserveImbalance -= amount;
                emit CollateralAdded({amount: amount, remaining: currentReserveImbalance});
                return;
            }
        }

        totalMinted += adjustedAmount;
        emit NativeTokensMinted(recipient, adjustedAmount);
        // Calls NativeMinter precompile through INativeMinter interface.
        NATIVE_MINTER.mintNativeCoin(recipient, adjustedAmount);
    }
}
