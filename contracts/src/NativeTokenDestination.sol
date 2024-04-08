// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    TeleporterTokenDestination,
    TeleporterFeeInfo,
    TeleporterMessageInput
} from "./TeleporterTokenDestination.sol";
import {Address} from "@openzeppelin/contracts@4.8.1/utils/Address.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {INativeTokenDestination} from "./interfaces/INativeTokenDestination.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
// We need IAllowList as an indirect dependency in order to compile.
// solhint-disable-next-line no-unused-import
import {IAllowList} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IAllowList.sol";
import {IWrappedNativeToken} from "./interfaces/IWrappedNativeToken.sol";
import {SendTokensInput} from "./interfaces/ITeleporterTokenBridge.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @notice Implementation of the {INativeTokenDestination} interface.
 *
 * @dev This contract pairs with exactly one `TokenSource` contract on the source chain and one wrapped native
 * token on the chain this contract is deployed to.
 * It mints and burns native tokens on the destination chain corresponding to locks and unlocks on the source chain.
 */
contract NativeTokenDestination is
    TeleporterOwnerUpgradeable,
    INativeTokenDestination,
    TeleporterTokenDestination
{
    /**
     * @notice The address where the burned transaction fees are credited.
     *
     * @dev Defined as BLACKHOLE_ADDRESS at
     * https://github.com/ava-labs/subnet-evm/blob/e23ab058d039ff9c8469c89b139d21d52c4bd283/constants/constants.go
     * It is a system-reserved address by default in subnet-evm and coreth, and transfers cannot be sent here manually.
     */
    address public constant BURNED_TX_FEES_ADDRESS = 0x0100000000000000000000000000000000000000;

    /**
     * @notice Address used by this contract to blackhole funds, effectively burning them.
     *
     * @dev Native tokens are burned by this contract by sending them to this address when transferring tokens back to
     * the source chain. Different from BURNED_TX_FEES_ADDRESS so that the total amount burned in transaction fees and
     * the total amount burned to be sent back to the source chain can be tracked separately.
     * This address was chosen arbitrarily.
     */
    address public constant BURN_FOR_TRANSFER_ADDRESS = 0x0100000000000000000000000000000000010203;

    /**
     * @notice Address used by the source chain to blackhole funds, effectively burning them.
     *
     * @dev When reporting burned transaction fee amounts, this address is used as the recipient
     * address for the funds to be sent to be burned on the source chain.
     * This address was chosen arbitrarily.
     */
    address public constant SOURCE_CHAIN_BURN_ADDRESS = 0x0100000000000000000000000000000000010203;

    /**
     * @notice Address of the Native Minter precompile contract.
     */
    INativeMinter public constant NATIVE_MINTER =
        INativeMinter(0x0200000000000000000000000000000000000001);

    /**
     * @notice Estimated gas needed for a transfer call to execute successfully on the source chain.
     */
    uint256 public constant TRANSFER_NATIVE_TOKENS_REQUIRED_GAS = 100_000;

    /**
     * @notice Initial reserve imbalance that must be collateralized on the source before minting.
     *
     * @dev The first `initialReserveImbalance` tokens sent to this subnet will not be minted.
     * `initialReserveImbalance` should be constructed to match the initial token supply of this subnet.
     * This means tokens will not be minted until the source contact is collateralized.
     */
    uint256 public immutable initialReserveImbalance;

    /**
     * @notice Tokens will not be minted until this value is 0 meaning the source contact is collateralized.
     */
    uint256 public currentReserveImbalance;

    /**
     * @notice Total number of tokens minted by this contract through the native minter precompile.
     */
    uint256 public totalMinted;

    /**
     * @notice The balance of BURNED_TX_FEES_ADDRESS the last time burned fees were reported to the source chain.
     */
    uint256 public latestBurnAddressBalance;

    /**
     * @notice tokenMultiplier allows this contract to scale the number of tokens it sends/receives to/from
     * the source chain.
     *
     * @dev This can be used to normalize the number of decimals places between the tokens on
     * the two subnets. Is calculated as 10^d, where d is decimalsShift specified in the constructor.
     */
    uint256 public immutable tokenMultiplier;

    /**
     * @notice If multiplyOnReceive is true, the raw token amount value will be multiplied by `tokenMultiplier` when tokens
     * are transferred from the source chain into this destination chain, and divided by `tokenMultiplier` when
     * tokens are transferred from this destination chain back to the source chain. This is intended
     * when the "decimals" value on the source chain is less than the native EVM denomination of 18.
     * If multiplyOnReceive is false, the raw token amount value will be divided by `tokenMultiplier` when tokens
     * are transferred from the source chain into this destination chain, and multiplied by `tokenMultiplier` when
     * tokens are transferred from this destination chain back to the source chain.
     */
    bool public immutable multiplyOnReceive;

    /**
     * @notice The wrapped native token contract that represents the native tokens on this chain.
     */
    IWrappedNativeToken public immutable token;

    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID_,
        address tokenSourceAddress_,
        address feeTokenAddress_,
        uint256 initialReserveImbalance_,
        uint256 decimalsShift,
        bool multiplyOnReceive_
    )
        TeleporterTokenDestination(
            teleporterRegistryAddress,
            teleporterManager,
            sourceBlockchainID_,
            tokenSourceAddress_,
            feeTokenAddress_
        )
    {
        token = IWrappedNativeToken(feeTokenAddress);

        require(
            initialReserveImbalance_ != 0, "NativeTokenDestination: zero initial reserve imbalance"
        );

        initialReserveImbalance = initialReserveImbalance_;
        currentReserveImbalance = initialReserveImbalance_;

        require(decimalsShift <= 18, "NativeTokenDestination: invalid decimalsShift");
        tokenMultiplier = 10 ** decimalsShift;
        multiplyOnReceive = multiplyOnReceive_;
    }

    /**
     * @notice Receives native tokens transferred to this contract.
     * @dev This function is called when the token bridge is withdrawing native tokens to
     * transfer to the recipient. The caller must be the wrapped native token contract.
     */
    receive() external payable {
        require(
            msg.sender == feeTokenAddress, "NativeTokenDestination: invalid receive payable sender"
        );
    }

    /**
     * @dev See {INativeTokenBridge-send}.
     */
    function send(SendTokensInput calldata input) external payable {
        require(
            currentReserveImbalance == 0, "NativeTokenDestination: contract undercollateralized"
        );
        uint256 amount = msg.value;
        require(
            amount > input.primaryFee + input.secondaryFee,
            "NativeTokenDestination: insufficient amount to cover fees"
        );

        // TODO we need to guarantee that this function deposits the whole amount, or find a workaround.
        if (input.primaryFee > 0) {
            _deposit(input.primaryFee);
            amount -= input.primaryFee;
        }

        _burn(amount);

        _send(input, _scaleTokens(amount, false));
    }

    /**
     * @dev See {INativeTokenDestination-reportTotalBurnedTxFees}.
     */
    function reportBurnedTxFees(uint256 requiredGasLimit) external payable {
        uint256 adjustedFeeAmount;
        if (msg.value > 0) {
            adjustedFeeAmount = _deposit(msg.value);
        }

        uint256 burnAddressBalance = BURNED_TX_FEES_ADDRESS.balance;
        require(
            burnAddressBalance > latestBurnAddressBalance,
            "NativeTokenDestination: burn address balance not greater than last report"
        );

        uint256 burnedDifference = burnAddressBalance - latestBurnAddressBalance;
        latestBurnAddressBalance = burnAddressBalance;

        uint256 scaledAmount = _scaleTokens(burnedDifference, false);
        require(scaledAmount > 0, "NativeTokenDestination: zero scaled amount to report burn");

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: tokenSourceAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeTokenAddress, amount: adjustedFeeAmount}),
                requiredGasLimit: requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(
                    SendTokensInput({
                        destinationBlockchainID: sourceBlockchainID,
                        destinationBridgeAddress: tokenSourceAddress,
                        recipient: SOURCE_CHAIN_BURN_ADDRESS,
                        primaryFee: 0,
                        secondaryFee: 0,
                        requiredGasLimit: 0
                    }),
                    scaledAmount
                    )
            })
        );

        emit ReportBurnedTxFees({teleporterMessageID: messageID, feesBurned: burnedDifference});
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
     * @dev See {TeleportTokenDestination-_deposit}
     */
    function _deposit(uint256 amount) internal virtual override returns (uint256) {
        token.deposit{value: amount}();
        return amount;
    }

    /**
     * @dev See {TeleportTokenDestination-_withdraw}
     */
    function _withdraw(address recipient, uint256 amount) internal virtual override {
        uint256 scaledAmount = _scaleTokens(amount, true);

        // If the contract has not yet been collateralized, we will deduct as many tokens
        // as needed from the transfer as needed. If there are any excess tokens, they will
        // be minted and sent to the recipient.
        uint256 adjustedAmount = scaledAmount;
        uint256 reserveImbalance = currentReserveImbalance;
        if (reserveImbalance > 0) {
            if (scaledAmount > reserveImbalance) {
                emit CollateralAdded({amount: reserveImbalance, remaining: 0});
                adjustedAmount -= reserveImbalance;
                currentReserveImbalance = 0;
            } else {
                uint256 updatedReserveImbalance = reserveImbalance - scaledAmount;
                emit CollateralAdded({amount: scaledAmount, remaining: updatedReserveImbalance});
                adjustedAmount = 0;
                currentReserveImbalance = updatedReserveImbalance;
            }
        }

        // Emit an event even if the amount is zero to improve traceability.
        emit NativeTokensMinted(recipient, adjustedAmount);

        // Only call the native minter precompile if we are minting any coins.
        if (adjustedAmount > 0) {
            totalMinted += adjustedAmount;
            // Calls NativeMinter precompile through INativeMinter interface.
            NATIVE_MINTER.mintNativeCoin(recipient, adjustedAmount);
        }
    }

    /**
     * @dev See {TeleportTokenDestination-_burn}
     */
    function _burn(uint256 amount) internal virtual override {
        // Burn native token by sending to BURN_FOR_TRANSFER_ADDRESS
        Address.sendValue(payable(BURN_FOR_TRANSFER_ADDRESS), amount);
    }

    /**
     * @dev Scales `value` based on `tokenMultiplier` and the direction of the transfer.
     * Should be used for all tokens being transferred to/from other subnets.
     */
    function _scaleTokens(uint256 value, bool isReceive) private view returns (uint256) {
        // Multiply when multiplyOnReceive and isReceive are both true or both false.
        if (multiplyOnReceive == isReceive) {
            return value * tokenMultiplier;
        }
        // Otherwise divide.
        return value / tokenMultiplier;
    }
}
