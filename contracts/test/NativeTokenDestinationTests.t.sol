// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenBridgeTest} from "./TeleporterTokenBridgeTests.t.sol";
import {TeleporterTokenDestinationTest} from "./TeleporterTokenDestinationTests.t.sol";
import {NativeTokenBridgeTest} from "./NativeTokenBridgeTests.t.sol";
import {INativeSendAndCallReceiver} from "../src/interfaces/INativeSendAndCallReceiver.sol";
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

contract NativeTokenDestinationTest is NativeTokenBridgeTest, TeleporterTokenDestinationTest {
    address public constant TEST_ACCOUNT = 0xd4E96eF8eee8678dBFf4d535E033Ed1a4F7605b7;
    string public constant DEFAULT_SYMBOL = "XYZ";
    NativeTokenDestination public app;

    event CollateralAdded(uint256 amount, uint256 remaining);
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
            multiplyOnReceive: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));
        tokenDestination = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        bridgedToken = app;
        assertEq(app.totalNativeAssetSupply(), _DEFAULT_INITIAL_RESERVE_IMBALANCE);
        _collateralizeBridge();
    }

    function testIsCollateralized() public {
        uint256 initialImbalance = 100;

        // Need a new instance since the default set up pre-collateralizes the contract.
        app = new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: initialImbalance,
            decimalsShift: 0,
            multiplyOnReceive: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));

        // Starts off under collateralized.
        assertFalse(app.isCollateralized());

        // Add less than the full amount of required collateral.
        vm.expectEmit(true, true, true, true, address(app));
        emit CollateralAdded({amount: initialImbalance / 2, remaining: initialImbalance / 2});

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(initialImbalance / 2, DEFAULT_RECIPIENT_ADDRESS)
        );
        assertEq(app.currentReserveImbalance(), initialImbalance / 2);
        assertFalse(app.isCollateralized());

        // Add more than the remaining amount of imbalance.
        vm.expectEmit(true, true, true, true, address(app));
        emit CollateralAdded({amount: initialImbalance / 2, remaining: 0});
        _setUpMockMint(DEFAULT_RECIPIENT_ADDRESS, initialImbalance / 2);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(initialImbalance, DEFAULT_RECIPIENT_ADDRESS)
        );

        // It should now be collateralized.
        assertEq(app.currentReserveImbalance(), 0);
        assertTrue(app.isCollateralized());
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
            multiplyOnReceive: true,
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
            multiplyOnReceive: true,
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
            multiplyOnReceive: false,
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
            multiplyOnReceive: false,
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
            multiplyOnReceive: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));

        vm.expectRevert("NativeTokenDestination: contract undercollateralized");
        app.send{value: 100_000}(_createDefaultSendTokensInput());
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
            multiplyOnReceive: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));

        vm.expectRevert("NativeTokenDestination: contract undercollateralized");
        app.sendAndCall{value: 100_000}(_createDefaultSendAndCallInput());
    }

    function testTotalNativeAssetSupply() public {
        assertEq(app.totalNativeAssetSupply(), _DEFAULT_INITIAL_RESERVE_IMBALANCE);

        // Mock tokens being burned as tx fees.
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), _DEFAULT_INITIAL_RESERVE_IMBALANCE - 1);
        assertEq(app.totalNativeAssetSupply(), 1);

        // Reset the burned tx fee amount.
        vm.deal(app.BURNED_TX_FEES_ADDRESS(), 0);
        assertEq(app.totalNativeAssetSupply(), _DEFAULT_INITIAL_RESERVE_IMBALANCE);

        // Mock tokens being bridged out by crediting them to the native token destination contract
        vm.deal(app.BURNED_FOR_BRIDGE_ADDRESS(), _DEFAULT_INITIAL_RESERVE_IMBALANCE - 1);
        assertEq(app.totalNativeAssetSupply(), 1);

        // Depositing native tokens into the contract to be wrapped native tokens shouldn't affect the supply
        // of the native asset, but should be reflected in the total supply of the ERC20 representation.
        app.deposit{value: 2}();
        assertEq(app.totalNativeAssetSupply(), 1);
        assertEq(app.totalSupply(), 2);
    }

    function testTransferToSource() public {
        vm.expectEmit(true, true, true, true, address(app));
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = _DEFAULT_TRANSFER_AMOUNT;
        uint256 scaledAmount = _scaleTokens(amount, false);
        emit TokensSent({
            teleporterMessageID: _MOCK_MESSAGE_ID,
            sender: address(this),
            input: input,
            amount: scaledAmount
        });

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(app), amount: input.primaryFee}),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(scaledAmount, input.recipient)
        });

        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        app.send{value: amount}(input);
    }

    function testScaleTokensMultiplyOnSend() public {
        uint8 decimalShift = 3;
        app = new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: decimalShift,
            multiplyOnReceive: false,
            burnedFeesReportingRewardPercentage: 1
        }));
        assertEq(app.scaleTokens(100, false), 100_000);
        assertEq(app.scaleTokens(100_000, true), 100);
    }

    function testScaleTokensMultiplyOnReceive() public {
        uint8 decimalShift = 3;
        app = new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: decimalShift,
            multiplyOnReceive: true,
            burnedFeesReportingRewardPercentage: 1
        }));
        assertEq(app.scaleTokens(100, true), 100_000);
        assertEq(app.scaleTokens(100_000, false), 100);
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

        uint256 scaledAmount = _scaleTokens(amount, true);
        _setUpMockMint(address(app), scaledAmount);
        vm.deal(address(app), scaledAmount - 1);
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
        uint256 expectedReportedAmount =
            _scaleTokens(initialBurnedTxFeeAmount - expectedReward, false);
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
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT
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
        expectedReportedAmount = _scaleTokens(additionalBurnTxFeeAmount - expectedReward, false);
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
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT
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
            multiplyOnReceive: false,
            burnedFeesReportingRewardPercentage: 0
        }));
        tokenDestination = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        bridgedToken = app;
        _collateralizeBridge();

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
                requiredGasLimit: DEFAULT_REQUIRED_GAS_LIMIT
            }),
            burnedTxFeeAmount
        );
        _checkExpectedTeleporterCallsForSend(expectedMessageInput);
        app.reportBurnedTxFees(DEFAULT_REQUIRED_GAS_LIMIT);
    }

    function testReceiveSendAndCallBeforeCollateralized() public {
        // Need a new instance since the default set up pre-collateralizes the contract.
        app = new NativeTokenDestination(NativeTokenDestinationSettings({
            nativeAssetSymbol: DEFAULT_SYMBOL,
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            initialReserveImbalance: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: 0,
            multiplyOnReceive: true,
            burnedFeesReportingRewardPercentage: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        }));

        // Add more than the full amount of required collateral.
        vm.expectEmit(true, true, true, true, address(app));
        emit CollateralAdded({amount: _DEFAULT_INITIAL_RESERVE_IMBALANCE, remaining: 0});
        _setUpMockMint(DEFAULT_FALLBACK_RECIPIENT_ADDRESS, _DEFAULT_INITIAL_RESERVE_IMBALANCE);

        bytes32 sourceBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        address originSenderAddress = address(this);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            sourceBlockchainID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopCallMessage({
                sourceBlockchainID: sourceBlockchainID,
                originSenderAddress: originSenderAddress,
                amount: _DEFAULT_INITIAL_RESERVE_IMBALANCE * 2,
                recipientContract: DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
                recipientPayload: new bytes(0),
                recipientGasLimit: DEFAULT_RECIPIENT_GAS_LIMIT,
                fallbackRecipient: DEFAULT_FALLBACK_RECIPIENT_ADDRESS
            })
        );
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

    function _collateralizeBridge() internal {
        vm.expectEmit(true, true, true, true, address(app));
        emit CollateralAdded({amount: app.currentReserveImbalance(), remaining: 0});

        assertFalse(app.isCollateralized());

        _setUpMockMint(DEFAULT_RECIPIENT_ADDRESS, 0);
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(
                app.currentReserveImbalance() / app.tokenMultiplier(), DEFAULT_RECIPIENT_ADDRESS
            )
        );

        assertTrue(app.isCollateralized());
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        uint256 scaledAmount = _scaleTokens(amount, true);
        vm.expectEmit(true, true, true, true, address(tokenDestination));
        emit TokensWithdrawn(recipient, scaledAmount);
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, scaledAmount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (recipient, scaledAmount))
        );
        vm.deal(recipient, scaledAmount);
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
        uint256 scaledAmount = _scaleTokens(amount, true);
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (address(app), scaledAmount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (address(app), scaledAmount))
        );

        // Mock the native minter precompile crediting native balance to the contract.
        vm.deal(address(app), scaledAmount);

        if (targetHasCode) {
            // Non-zero code length
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                INativeSendAndCallReceiver.receiveTokens,
                (sourceBlockchainID, originSenderAddress, payload)
            );
            if (expectSuccess) {
                vm.mockCall(recipient, scaledAmount, expectedCalldata, new bytes(0));
            } else {
                vm.mockCallRevert(recipient, scaledAmount, expectedCalldata, new bytes(0));
            }
            vm.expectCall(recipient, scaledAmount, uint64(gasLimit), expectedCalldata);
        } else {
            // No code at target
            vm.etch(recipient, new bytes(0));
        }

        if (targetHasCode && expectSuccess) {
            vm.expectEmit(true, true, true, true, address(app));
            emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, scaledAmount);
        } else {
            vm.expectEmit(true, true, true, true, address(app));
            emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, scaledAmount);
        }
    }

    function _setUpExpectedDeposit(uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(app), amount);
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert(_formatErrorMessage("insufficient amount to cover fees"));
    }

    function _getTotalSupply() internal view override returns (uint256) {
        return app.totalNativeAssetSupply();
    }

    function _scaleTokens(
        uint256 amount,
        bool isReceive
    ) internal view override returns (uint256) {
        if (app.multiplyOnReceive() == isReceive) {
            return amount * app.tokenMultiplier();
        }
        return amount / app.tokenMultiplier();
    }
}
