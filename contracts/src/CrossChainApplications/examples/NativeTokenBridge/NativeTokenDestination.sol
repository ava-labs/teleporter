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
import {WAVAX} from "./WAVAX.sol";
import {
    ITeleporterMessenger,
    TeleporterFeeInfo,
    TeleporterMessageInput
} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
import {SafeERC20TransferFrom} from "@teleporter/SafeERC20TransferFrom.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
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
    // sending an unlock message to the source chain. Reporting burned gas fees to the source chain
    // will send a transfer to this address. Different from the burned tx fee address so they can be
    // tracked separately.
    address public constant GENERAL_BURN_ADDRESS = 0x0100000000000000000000000000000000010203;

    INativeMinter public constant NATIVE_MINTER =
        INativeMinter(0x0200000000000000000000000000000000000001);
    WAVAX public immutable wavaxContract;

    uint256 public constant TRANSFER_NATIVE_TOKENS_REQUIRED_GAS = 100_000;
    bytes32 public immutable sourceBlockchainID;
    address public immutable nativeTokenSourceAddress;
    uint256 public immutable burnedFeesReportingRewardPercentage;
    // The first `initialReserveImbalance` tokens sent to this subnet will not be minted.
    // `initialReserveImbalance` should be constructed to match the initial token supply of this subnet.
    // This means tokens will not be minted until the source contact is collateralized.
    uint256 public immutable initialReserveImbalance;
    uint256 public currentReserveImbalance;
    uint256 public totalMinted;
    uint256 public lastestBurnedFeesReported;

    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID_,
        address nativeTokenSourceAddress_,
        uint256 initialReserveImbalance_,
        uint256 burnedFeesReportingRewardPercentage_,
        address wavaxContractAddress
    ) TeleporterOwnerUpgradeable(teleporterRegistryAddress, teleporterManager) {
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

        require(
            burnedFeesReportingRewardPercentage_ <= 100, "NativeTokenDestination: invalid percentage"
        );
        burnedFeesReportingRewardPercentage = burnedFeesReportingRewardPercentage_;

        require(
            wavaxContractAddress != address(0),
            "NativeTokenDestination: zero wavax contract address"
        );
        wavaxContract = WAVAX(payable(wavaxContractAddress));
    }

    /**
     * @dev See {INativeTokenDestination-transferToSource}.
     */
    function transferToSource(
        address recipient,
        TeleporterFeeInfo calldata feeInfo,
        address[] calldata allowedRelayerAddresses
    ) external payable nonReentrant {
        // The recipient cannot be the zero address.
        require(recipient != address(0), "NativeTokenDestination: zero recipient address");

        require(
            currentReserveImbalance == 0, "NativeTokenDestination: contract undercollateralized"
        );
        uint256 value = msg.value;
        require(value > 0, "NativeTokenDestination: zero transfer value");

        // Lock tokens in this bridge instance. Supports "fee/burn on transfer" ERC20 token
        // implementations by only transferring the actual balance increase reflected by the call
        // to transferFrom.
        uint256 adjustedFeeAmount;
        if (feeInfo.amount > 0) {
            adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
                IERC20(feeInfo.feeTokenAddress), feeInfo.amount
            );
        }

        // Burn native token by sending to GENERAL_BURN_ADDRESS
        Address.sendValue(payable(GENERAL_BURN_ADDRESS), value);
        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: nativeTokenSourceAddress,
                feeInfo: TeleporterFeeInfo({
                    feeTokenAddress: feeInfo.feeTokenAddress,
                    amount: adjustedFeeAmount
                }),
                requiredGasLimit: TRANSFER_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(recipient, value)
            })
        );

        emit TransferToSource({
            sender: msg.sender,
            recipient: recipient,
            teleporterMessageID: messageID,
            amount: value
        });
    }

    /**
     * @dev See {INativeTokenDestination-reportBurnedTxFees}.
     */
    function reportBurnedTxFees(address[] calldata allowedRelayerAddresses) external {
        uint256 burnedDifference = BURNED_TX_FEES_ADDRESS.balance - lastestBurnedFeesReported;
        uint256 reward = burnedDifference * burnedFeesReportingRewardPercentage / 100;
        uint256 burnedTxFees = burnedDifference - reward;

        totalMinted += reward;
        emit NativeTokensMinted(address(this), reward);
        // Calls NativeMinter precompile through INativeMinter interface.
        NATIVE_MINTER.mintNativeCoin(address(this), reward);

        wavaxContract.deposit{value: reward}();

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: nativeTokenSourceAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: address(wavaxContract), amount: reward}),
                requiredGasLimit: TRANSFER_NATIVE_TOKENS_REQUIRED_GAS,
                allowedRelayerAddresses: allowedRelayerAddresses,
                message: abi.encode(GENERAL_BURN_ADDRESS, burnedTxFees)
            })
        );

        emit ReportBurnedTxFees({teleporterMessageID: messageID, feesBurned: burnedTxFees});
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
        uint256 burned = BURNED_TX_FEES_ADDRESS.balance + GENERAL_BURN_ADDRESS.balance;
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
