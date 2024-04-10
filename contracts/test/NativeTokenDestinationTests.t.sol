// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {TeleporterTokenDestinationTest} from "./TeleporterTokenDestinationTests.t.sol";
import {NativeTokenBridgeTest} from "./NativeTokenBridgeTests.t.sol";
import {INativeSendAndCallReceiver} from "../src/interfaces/INativeSendAndCallReceiver.sol";
import {
    NativeTokenDestination,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "../src/NativeTokenDestination.sol";
import {IWrappedNativeToken} from "../src/interfaces/IWrappedNativeToken.sol";
import {ExampleWAVAX} from "../src/mocks/ExampleWAVAX.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/INativeMinter.sol";
import {ITeleporterMessenger} from "@teleporter/ITeleporterMessenger.sol";
import {SendTokensInput} from "../src/interfaces/ITeleporterTokenBridge.sol";

contract NativeTokenDestinationTest is NativeTokenBridgeTest, TeleporterTokenDestinationTest {
    NativeTokenDestination public app;
    IWrappedNativeToken public mockWrappedToken;

    event CollateralAdded(uint256 amount, uint256 remaining);
    event ReportBurnedTxFees(bytes32 indexed teleporterMessageID, uint256 feesBurned);

    function setUp() public override {
        TeleporterTokenDestinationTest.setUp();

        mockWrappedToken = new ExampleWAVAX();
        app = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID_: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress_: TOKEN_SOURCE_ADDRESS,
            feeTokenAddress_: address(mockWrappedToken),
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: _DEFAULT_DECIMALS_SHIFT,
            multiplyOnReceive_: true,
            burnedFeesReportingRewardPercentage_: _DEFAULT_BURN_FEE_REWARDS_PERCENTAGE
        });
        tokenDestination = app;
        nativeTokenBridge = app;
        tokenBridge = app;
        feeToken = mockWrappedToken;
        collateralizeBridge();
    }

    function collateralizeBridge() public {
        vm.expectEmit(true, true, true, true, address(app));
        emit CollateralAdded({amount: _DEFAULT_INITIAL_RESERVE_IMBALANCE, remaining: 0});

        bool isCollateralized = app.isCollateralized();
        assertEq(isCollateralized, false);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(
                _DEFAULT_INITIAL_RESERVE_IMBALANCE / _DEFAULT_TOKEN_MULTIPLIER,
                DEFAULT_RECIPIENT_ADDRESS
            )
        );

        isCollateralized = app.isCollateralized();
        assertEq(isCollateralized, true);
    }

    function testZeroSourceBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("zero source blockchain ID"));
        new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID_: bytes32(0),
            tokenSourceAddress_: TOKEN_SOURCE_ADDRESS,
            feeTokenAddress_: address(mockWrappedToken),
            initialReserveImbalance_: 1_000,
            decimalsShift: 0,
            multiplyOnReceive_: false,
            burnedFeesReportingRewardPercentage_: 1
        });
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
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: app.feeTokenAddress(),
                amount: input.primaryFee
            }),
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
        app = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID_: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress_: TOKEN_SOURCE_ADDRESS,
            feeTokenAddress_: address(mockWrappedToken),
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: decimalShift,
            multiplyOnReceive_: false,
            burnedFeesReportingRewardPercentage_: 1
        });
        assertEq(app.scaleTokens(100, false), 100_000);
        assertEq(app.scaleTokens(100_000, true), 100);
    }

    function testScaleTokensMultiplyOnReceive() public {
        uint8 decimalShift = 3;
        app = new NativeTokenDestination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: MOCK_TELEPORTER_MESSENGER_ADDRESS,
            sourceBlockchainID_: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress_: TOKEN_SOURCE_ADDRESS,
            feeTokenAddress_: address(mockWrappedToken),
            initialReserveImbalance_: _DEFAULT_INITIAL_RESERVE_IMBALANCE,
            decimalsShift: decimalShift,
            multiplyOnReceive_: true,
            burnedFeesReportingRewardPercentage_: 1
        });
        assertEq(app.scaleTokens(100, true), 100_000);
        assertEq(app.scaleTokens(100_000, false), 100);
    }

    function _checkExpectedWithdrawal(address addr, uint256 amount) internal override {
        uint256 scaledAmount = _scaleTokens(amount, true);
        vm.mockCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (addr, scaledAmount)),
            new bytes(0)
        );
        vm.expectCall(
            NATIVE_MINTER_PRECOMPILE_ADDRESS,
            abi.encodeCall(INativeMinter.mintNativeCoin, (addr, scaledAmount))
        );
        vm.deal(addr, scaledAmount);
    }

    function _setUpExpectedSendAndCall(
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
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

        bytes memory expectedCalldata =
            abi.encodeCall(INativeSendAndCallReceiver.receiveTokens, (payload));
        vm.etch(recipient, new bytes(1));
        if (expectSuccess) {
            vm.mockCall(recipient, scaledAmount, expectedCalldata, new bytes(0));
        } else {
            vm.mockCallRevert(recipient, scaledAmount, expectedCalldata, new bytes(0));
        }
        vm.expectCall(recipient, scaledAmount, uint64(gasLimit), expectedCalldata);
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
