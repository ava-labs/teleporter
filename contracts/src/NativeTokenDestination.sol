// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {
    TeleporterTokenDestination,
    TeleporterFeeInfo,
    TeleporterMessageInput
} from "./TeleporterTokenDestination.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {INativeTokenDestination} from "./interfaces/INativeTokenDestination.sol";
import {INativeSendAndCallReceiver} from "./interfaces/INativeSendAndCallReceiver.sol";
import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
// We need IAllowList as an indirect dependency in order to compile.
// solhint-disable-next-line no-unused-import
import {IAllowList} from "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IAllowList.sol";
import {IWrappedNativeToken} from "./interfaces/IWrappedNativeToken.sol";
import {
    SendTokensInput,
    SendAndCallInput,
    BridgeMessageType,
    BridgeMessage,
    SingleHopSendMessage,
    SingleHopCallMessage
} from "./interfaces/ITeleporterTokenBridge.sol";
import {SafeWrappedNativeTokenDeposit} from "./utils/SafeWrappedNativeTokenDeposit.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {SendReentrancyGuard} from "./utils/SendReentrancyGuard.sol";
import {CallUtils} from "./utils/CallUtils.sol";

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
    SendReentrancyGuard,
    TeleporterTokenDestination
{
    using SafeERC20 for IWrappedNativeToken;

    /**
     * @notice The address where the burned transaction fees are credited.
     *
     * @dev Defined as BLACKHOLE_ADDRESS at
     * https://github.com/ava-labs/subnet-evm/blob/v0.6.0/constants/constants.go
     * C-Chain value found at https://github.com/ava-labs/coreth/blob/v0.13.2/constants/constants.go
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
     * @notice Initial reserve imbalance that must be collateralized on the source before minting.
     *
     * @dev The first `initialReserveImbalance` tokens sent to this subnet will not be minted.
     * `initialReserveImbalance` should be constructed to match the initial token supply of this subnet.
     * This means tokens will not be minted until the source contact is collateralized.
     */
    uint256 public immutable initialReserveImbalance;

    /**
     * @notice Percentage of burned transaction fees that will be rewarded to a relayer delivering
     * the message created by calling calling reportBurnedTxFees().
     */
    uint256 public immutable burnedFeesReportingRewardPercentage;

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
    uint256 public lastestBurnedFeesReported;

    /**
     * @notice The wrapped native token contract that represents the native tokens on this chain.
     */
    IWrappedNativeToken public immutable token;

    modifier onlyWhenCollateralized() {
        require(_isCollateralized(), "NativeTokenDestination: contract undercollateralized");
        _;
    }

    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID_,
        address tokenSourceAddress_,
        address feeTokenAddress_,
        uint256 initialReserveImbalance_,
        uint8 decimalsShift,
        bool multiplyOnReceive_,
        uint256 burnedFeesReportingRewardPercentage_
    )
        TeleporterTokenDestination(
            teleporterRegistryAddress,
            teleporterManager,
            sourceBlockchainID_,
            tokenSourceAddress_,
            feeTokenAddress_,
            decimalsShift,
            multiplyOnReceive_
        )
    {
        token = IWrappedNativeToken(feeTokenAddress);

        require(
            initialReserveImbalance_ != 0, "NativeTokenDestination: zero initial reserve imbalance"
        );

        initialReserveImbalance = initialReserveImbalance_;
        currentReserveImbalance = initialReserveImbalance_;

        require(
            burnedFeesReportingRewardPercentage_ < 100, "NativeTokenDestination: invalid percentage"
        );
        burnedFeesReportingRewardPercentage = burnedFeesReportingRewardPercentage_;
    }

    /**
     * @notice Receives native tokens transferred to this contract without calldata.
     * @dev This function is called when the token bridge is withdrawing native tokens to
     * transfer to the recipient. Only the wrapped native token contract is allowed to call
     * this function to prevent accidental loss of funds from other accounts.
     */
    receive() external payable {
        require(
            msg.sender == feeTokenAddress, "NativeTokenDestination: invalid receive payable sender"
        );
    }

    /**
     * @dev See {INativeTokenBridge-send}.
     */
    function send(SendTokensInput calldata input) external payable onlyWhenCollateralized {
        _send(input, msg.value);
    }

    /**
     * @dev See {INativeTokenBridge-sendAndCall}
     */
    function sendAndCall(SendAndCallInput calldata input) external payable onlyWhenCollateralized {
        _sendAndCall(input, msg.value);
    }

    /**
     * @dev See {INativeTokenDestination-reportTotalBurnedTxFees}.
     */
    function reportBurnedTxFees(uint256 requiredGasLimit) external sendNonReentrant {
        uint256 burnAddressBalance = BURNED_TX_FEES_ADDRESS.balance;
        require(
            burnAddressBalance > lastestBurnedFeesReported,
            "NativeTokenDestination: burn address balance not greater than last report"
        );

        uint256 burnedDifference = burnAddressBalance - lastestBurnedFeesReported;
        uint256 reward = (burnedDifference * burnedFeesReportingRewardPercentage) / 100;
        uint256 burnedTxFees = burnedDifference - reward;
        lastestBurnedFeesReported = burnAddressBalance;

        if (reward > 0) {
            _mint(address(this), reward);
            _deposit(reward);
        }

        uint256 scaledAmount = scaleTokens(burnedTxFees, false);
        require(scaledAmount > 0, "NativeTokenDestination: zero scaled amount to report burn");

        BridgeMessage memory message = BridgeMessage({
            messageType: BridgeMessageType.SINGLE_HOP_SEND,
            amount: scaledAmount,
            payload: abi.encode(SingleHopSendMessage({recipient: SOURCE_CHAIN_BURN_ADDRESS}))
        });

        bytes32 messageID = _sendTeleporterMessage(
            TeleporterMessageInput({
                destinationBlockchainID: sourceBlockchainID,
                destinationAddress: tokenSourceAddress,
                feeInfo: TeleporterFeeInfo({feeTokenAddress: feeTokenAddress, amount: reward}),
                requiredGasLimit: requiredGasLimit,
                allowedRelayerAddresses: new address[](0),
                message: abi.encode(message)
            })
        );

        emit ReportBurnedTxFees({teleporterMessageID: messageID, feesBurned: burnedTxFees});
    }

    /**
     * @dev See {INativeTokenDestination-isCollateralized}.
     */
    function isCollateralized() external view returns (bool) {
        return _isCollateralized();
    }

    /**
     * @dev See {INativeTokenDestination-totalSupply}.
     */
    function totalSupply() external view returns (uint256) {
        uint256 burned = BURNED_TX_FEES_ADDRESS.balance + token.balanceOf(BURN_FOR_TRANSFER_ADDRESS);
        uint256 created = totalMinted + initialReserveImbalance;

        return created - burned;
    }

    /**
     * @dev See {TeleporterTokenDestination-_deposit}
     */
    function _deposit(uint256 amount) internal override returns (uint256) {
        return SafeWrappedNativeTokenDeposit.safeDeposit(token, amount);
    }

    /**
     * @dev See {TeleporterTokenDestination-_withdraw}
     */
    function _withdraw(address recipient, uint256 amount) internal override {
        emit TokensWithdrawn(recipient, amount);

        // If the contract has not yet been collateralized, we will deduct as many tokens
        // as needed from the transfer as needed. If there are any excess tokens, they will
        // be minted and sent to the recipient.
        uint256 adjustedAmount = amount;
        uint256 reserveImbalance = currentReserveImbalance;
        if (reserveImbalance > 0) {
            if (amount > reserveImbalance) {
                emit CollateralAdded({amount: reserveImbalance, remaining: 0});
                adjustedAmount -= reserveImbalance;
                currentReserveImbalance = 0;
            } else {
                uint256 updatedReserveImbalance = reserveImbalance - amount;
                emit CollateralAdded({amount: amount, remaining: updatedReserveImbalance});
                adjustedAmount = 0;
                currentReserveImbalance = updatedReserveImbalance;
            }
        }

        // Call {_mint} even if the adjustedAmount is 0 to improve traceability.
        _mint(recipient, adjustedAmount);
    }

    /**
     * @dev See {TeleporterTokenDestination-_burn}
     */
    function _burn(uint256 amount) internal override {
        // Burn native token by transferring to BURN_FOR_TRANSFER_ADDRESS
        token.safeTransfer(BURN_FOR_TRANSFER_ADDRESS, amount);
    }

    /**
     * @dev See {TeleporterTokenDestination-_handleSendAndCall}
     *
     * Mints the tokens to this contract, and send them to the recipient contract as a
     * part of the call to {INativeSendAndCallReceiver-receiveTokens} on the recipient contract.
     * If the call fails, the amount is sent to the fallback recipient.
     *
     * Note: If the recipient contract does not properly handle the full msg.value sent,
     * the balance can be locked in the recipient contract. Receiving contracts must make
     * sure to properly handle the balance to ensure it does not get locked improperly.
     */
    function _handleSendAndCall(
        SingleHopCallMessage memory message,
        uint256 amount
    ) internal override {
        // If the contract is not yet fully collateralized, the use of send and call is not allowed
        // because it could result in unexpected behavior given that the amount of tokens used to make the
        // call to "receiveTokens" is less than expected. Instead, the amount is handled as a normal bridge
        // event to fallback recipient.
        if (!_isCollateralized()) {
            _withdraw(message.fallbackRecipient, amount);
            return;
        }

        // Mint the tokens to this contract address.
        NATIVE_MINTER.mintNativeCoin(address(this), amount);

        // Encode the call to {INativeSendAndCallReceiver-receiveTokens}
        bytes memory payload =
            abi.encodeCall(INativeSendAndCallReceiver.receiveTokens, (message.recipientPayload));

        // Call the destination contract with the given payload, gas amount, and value.
        bool success = CallUtils._callWithExactGasAndValue(
            message.recipientGasLimit, amount, message.recipientContract, payload
        );

        // If the call failed, send the funds to the fallback recipient.
        if (success) {
            emit CallSucceeded(message.recipientContract, amount);
        } else {
            emit CallFailed(message.recipientContract, amount);
            payable(message.fallbackRecipient).transfer(amount);
        }
    }

    /**
     * @dev Returns whether or not the NativeTokenDestination contract has been collateralized
     * after being initialized with its given initialReserveImbalance.
     */
    function _isCollateralized() internal view returns (bool) {
        return currentReserveImbalance == 0;
    }

    /**
     * @dev Mints coins to the recipient through the NativeMinter precompile.
     */
    function _mint(address recipient, uint256 amount) private {
        totalMinted += amount;
        // Calls NativeMinter precompile through INativeMinter interface.
        NATIVE_MINTER.mintNativeCoin(recipient, amount);
    }
}
