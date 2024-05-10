// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenDestinationTest} from "./TeleporterTokenDestinationTests.t.sol";
import {NativeTokenBridgeTest} from "./NativeTokenBridgeTests.t.sol";
import {INativeSendAndCallReceiver} from "../src/interfaces/INativeSendAndCallReceiver.sol";
import {TeleporterTokenDestination} from "../src/TeleporterTokenDestination.sol";
import {
    NativeTokenDestinationSettings,
    NativeTokenDestination,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "../src/NativeTokenDestination.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {ITeleporterMessenger, TeleporterMessageInput} from "@teleporter/ITeleporterMessenger.sol";
import {SendTokensInput} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

contract NativeTokenDestinationTest is NativeTokenBridgeTest, TeleporterTokenDestinationTest {
    using SafeERC20 for IERC20;

    address public constant TEST_ACCOUNT = 0xd4E96eF8eee8678dBFf4d535E033Ed1a4F7605b7;
    string public constant DEFAULT_SYMBOL = "XYZ";
    NativeTokenDestination public app;

    event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned);

    function setUp() public override {
        TeleporterTokenDestinationTest.setUp();

        app = new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnDestination: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));
        tokenDestination = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        bridgedToken = app;
        assertEq(app.totalNativeAssetSupply(), _DEFAULT_INITIAL_RESERVE_IMBALANCE);
        _collateralizeBridge();
    }

    function testZeroInitialReserveImbalance() public {
        vm.expectRevert("NativeTokenDestination: zero initial reserve imbalance");
        new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: 0,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnDestination: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));
    }

    function testInvalidBurnedRewardPercentage() public {
        vm.expectRevert("NativeTokenDestination: invalid percentage");
        new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnDestination: true,
            burnedFeesReportingRewardPercentage: 100
        }));
    }

    function testZeroSourceBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("zero source blockchain ID"));
        new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: bytes32(0),
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: 1_000,
            decimalsShift: 0,
            multiplyOnDestination: false,
            burnedFeesReportingRewardPercentage: 1
        }));
    }

    function testDeployToSameBlockchain() public {
        vm.expectRevert(_formatErrorMessage("cannot deploy to same blockchain as source"));
        new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: 1_000,
            decimalsShift: 0,
            multiplyOnDestination: false,
            burnedFeesReportingRewardPercentage: 1
        }));
    }

    function testSendBeforeCollateralized() public {
        // Need a new instance since the default set up pre-collateralizes the contract.
        app = new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: 0,
            multiplyOnDestination: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));
        tokenDestination = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        bridgedToken = app;

        vm.expectRevert("NativeTokenDestination: contract undercollateralized");
        app.send{value: 100_000}(_createDefaultSendTokensInput());

        // Now mark the contract as collateralized and confirm sending is enabled.
        _collateralizeBridge();
        _sendSingleHopSendSuccess(100_000, 0);
    }

    function testSendAndCallBeforeCollateralized() public {
        // Need a new instance since the default set up pre-collateralizes the contract.
        app = new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: 0,
            multiplyOnDestination: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));
        tokenDestination = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        bridgedToken = app;

        vm.expectRevert("NativeTokenDestination: contract undercollateralized");
        app.sendAndCall{value: 100_000}(_createDefaultSendAndCallInput());

        // Now mark the contract as collateralized and confirm sending is enabled.
        _collateralizeBridge();
        _sendSingleHopSendSuccess(100_000, 0);
    }

    function testTotalNativeAssetSupply() public {
        uint256 initialTotalMinted = app.totalMinted();
        uint256 initialExpectedBalance = _DEFAULT_INITIAL_RESERVE_IMBALANCE + initialTotalMinted;
        assertEq(app.totalNativeAssetSupply(), initialExpectedBalance);

        // Mock tokens being burned as tx fees.
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), initialExpectedBalance - 1);
        assertEq(app.totalNativeAssetSupply(), 1);

        // Reset the burned tx fee amount.
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), 0);
        assertEq(app.totalNativeAssetSupply(), initialExpectedBalance);

        // Mock tokens being bridged out by crediting them to the native token destination contract
        vm.deal(app.BURNED_FOR_BRIDGE_ADDRESS(), initialExpectedBalance - 1);
        assertEq(app.totalNativeAssetSupply(), 1);

        // Depositing native tokens into the contract to be wrapped native tokens shouldn't affect the supply
        // of the native asset, but should be reflected in the total supply of the ERC20 representation.
        app.deposit{value: 2}();
        assertEq(app.totalNativeAssetSupply(), 1);
        assertEq(app.totalSupply(), 2);
    }

    function testTransferToSource() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = _DEFAULT_TRANSFER_AMOUNT;
        vm.expectEmit(true, true, true, true, address(app));
        emit TokensSent({
            teleporterMessageID: _MOCK_MESSAGE_ID,
            sender: address(this),
            input: input,
            amount: amount
        });

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(app), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(amount, input.recipient)
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        app.send{value: amount}(input);
    }

    function testReceiveSendAndCallFailureInsufficientValue() public {
        uint256 amount = 200;
        bytes memory payload = hex"DEADBEEF";
        bytes memory message = _encodeSingleHopCallMessage({
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            originSenderAddress: address(this),
            amount: amount,
            recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            recipientPayload: payload,
            recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
            fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        });

        _setUpMockMint(address(app), amount);
        vm.deal(address(app), amount - 1);
        vm.expectRevert("CallUtils: insufficient value");
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function testReportBurnFeesNoNewAmount() public {
        vm.expectRevert("NativeTokenDestination: burn address balance not greater than last report");
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testReportBurnFeesScaledToZero() public {
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), 1);
        vm.expectRevert("NativeTokenDestination: zero scaled amount to report burn");
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testReportBurnFeesSuccess() public {
        // First difference is 100,000
        uint256 initialBurnedTxFeeAmount = 100_003;
        uint256 expectedReward = 1_000; // 1%, rounded down due to integer division.
        uint256 expectedReportedAmount = initialBurnedTxFeeAmount - expectedReward;
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), initialBurnedTxFeeAmount);

        _setUpMockMint(address(app), expectedReward);
        TeleporterMessageInput memory expectedMessageInput = _createSingleHopTeleporterMessageInput(
            SendTokensInput({
                destinationBlockchainID: app.sourceBlockchainID(),
                destinationBridgeAddress: app.tokenSourceAddress(),
                recipient: app.SOURCE_CHAIN_BURN_ADDRESS(),
                feeTokenAddress: address(bridgedToken),
                primaryFee: expectedReward,
                secondaryFee: 0,
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            }),
            expectedReportedAmount
        );
        _checkExpectedTeleporterCallsForSend(expectedMessageInput);
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);

        // Calling it again should revert since no additional amount as been burned.
        vm.expectRevert("NativeTokenDestination: burn address balance not greater than last report");
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);

        // Mock more transaction fees being burned.
        uint256 additionalBurnTxFeeAmount = 50_007;
        expectedReward = 500; // 1%, rounded down due to integer division.
        expectedReportedAmount = additionalBurnTxFeeAmount - expectedReward;
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), initialBurnedTxFeeAmount + additionalBurnTxFeeAmount);

        _setUpMockMint(address(app), expectedReward);
        expectedMessageInput = _createSingleHopTeleporterMessageInput(
            SendTokensInput({
                destinationBlockchainID: app.sourceBlockchainID(),
                destinationBridgeAddress: app.tokenSourceAddress(),
                recipient: app.SOURCE_CHAIN_BURN_ADDRESS(),
                feeTokenAddress: address(bridgedToken),
                primaryFee: expectedReward,
                secondaryFee: 0,
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            }),
            expectedReportedAmount
        );
        _checkExpectedTeleporterCallsForSend(expectedMessageInput);
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testReportBurnFeesNoRewardSuccess() public {
        // Create a new destination instance with no rewards for reporting burned fees.
        app = new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: 0,
            multiplyOnDestination: false,
            burnedFeesReportingRewardPercentage: 0
        }));
        tokenDestination = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        bridgedToken = app;

        uint256 burnedTxFeeAmount = 100_000;
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), burnedTxFeeAmount);
        TeleporterMessageInput memory expectedMessageInput = _createSingleHopTeleporterMessageInput(
            SendTokensInput({
                destinationBlockchainID: app.sourceBlockchainID(),
                destinationBridgeAddress: app.tokenSourceAddress(),
                recipient: app.SOURCE_CHAIN_BURN_ADDRESS(),
                feeTokenAddress: address(bridgedToken),
                primaryFee: 0,
                secondaryFee: 0,
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            }),
            burnedTxFeeAmount
        );
        _checkExpectedTeleporterCallsForSend(expectedMessageInput);
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testFallback() public {
        (bool success,) = address(app).call{value: 1}("1234567812345678");
        assertTrue(success);
        assertEq(app.balanceOf(address(this)), 1);
    }

    function testDepositWithdrawWrappedNativeToken() public {
        uint256 depositAmount = 500;
        uint256 withdrawAmount = 100;
        vm.deal(TEST_ACCOUNT, depositAmount);
        vm.startPrank(TEST_ACCOUNT);
        app.deposit{value: depositAmount}();
        assertEq(app.balanceOf(TEST_ACCOUNT), depositAmount);
        app.withdraw(withdrawAmount);
        assertEq(app.balanceOf(TEST_ACCOUNT), depositAmount - withdrawAmount);
        assertEq(TEST_ACCOUNT.balance, withdrawAmount);
    }

    function _createNewDestinationInstance()
        internal
        override
        returns (TeleporterTokenDestination)
    {
        return new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnDestination: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenDestination));
        emit TokensWithdrawn(recipient, amount);
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, amount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, amount))
        );
        vm.deal(recipient, amount);
    }

    function _setUpMockMint(address recipient, uint256 amount) internal override {
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, amount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, amount))
        );
        vm.deal(recipient, amount);
    }

    function _setUpExpectedSendAndCall(
        bytes32 sourceBlockchainID,
        address originSenderAddress,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal override {
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (address(app), amount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (address(app), amount))
        );

        // Mock the native minter precompile crediting native balance to the contract.
        vm.deal(address(app), amount);

        if (targetHasCode) {
            // Non-zero code length
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                INativeSendAndCallReceiver.receiveTokens,
                (sourceBlockchainID, originSenderAddress, payload)
            );
            if (expectSuccess) {
                vm.mockCall(recipient, amount, expectedCalldata, new bytes(0));
            } else {
                vm.mockCallRevert(recipient, amount, expectedCalldata, new bytes(0));
            }
            vm.expectCall(recipient, amount, uint64(gasLimit), expectedCalldata);
        } else {
            // No code at target
            vm.etch(recipient, new bytes(0));
        }

        if (targetHasCode && expectSuccess) {
            vm.expectEmit(true, true, true, true, address(app));
            emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        } else {
            vm.expectEmit(true, true, true, true, address(app));
            emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        }
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal override {
        app.deposit{value: feeAmount}();
        // Transfer the fee to the bridge if it is greater than 0
        if (feeAmount > 0) {
            bridgedToken.safeIncreaseAllowance(address(tokenBridge), feeAmount);
            vm.expectCall(
                address(bridgedToken),
                abi.encodeCall(
                    IERC20.transferFrom, (address(this), address(tokenBridge), feeAmount)
                )
            );
        }
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(app), amount);
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert(_formatErrorMessage("insufficient amount to cover fees"));
    }

    function _getTotalSupply() internal view override returns (uint256) {
        return app.totalNativeAssetSupply();
    }

    // The native token destination contract is considered collateralized once it has received
    // a message from its configured source to mint tokens. Until then, the source contract is
    // still assumed to have insufficient collateral.
    function _collateralizeBridge() private {
        assertFalse(app.isCollateralized());
        uint256 amount = 10e18;
        _setUpMockMint(DEFAULT_RECIPIENT_ADDRESS, amount);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            app.sourceBlockchainID(),
            app.tokenSourceAddress(),
            _encodeSingleHopSendMessage(amount, DEFAULT_RECIPIENT_ADDRESS)
        );
        assertTrue(app.isCollateralized());
    }
}
