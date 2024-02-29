// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {TeleporterTokenSource, IWarpMessenger} from "../src/TeleporterTokenSource.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {UnitTestMockERC20} from "../lib/teleporter/contracts/src/Mocks/UnitTestMockERC20.sol";
import {
    ITeleporterTokenBridge, SendTokensInput
} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";

contract ExampleSourceApp is TeleporterTokenSource {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        address feeTokenAddress
    ) TeleporterTokenSource(teleporterRegistryAddress, teleporterManager, feeTokenAddress) {}

    function send(SendTokensInput calldata input, uint256 amount) external {
        _send(input, amount);
    }

    function _withdraw(address recipient, uint256 amount) internal virtual override {}
}

contract TeleporterTokenSourceTest is Test {
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_DESTINATION_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    address public constant DEFAULT_DESTINATION_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_RECIPIENT_ADDRESS = 0xABCDabcdABcDabcDaBCDAbcdABcdAbCdABcDABCd;
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");

    ExampleSourceApp public app;
    UnitTestMockERC20 public mockERC20;

    event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, uint256 amount);

    function setUp() public virtual {
        mockERC20 = new UnitTestMockERC20();
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_SOURCE_BLOCKCHAIN_ID)
        );

        vm.expectCall(
            WARP_PRECOMPILE_ADDRESS, abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector)
        );

        _initMockTeleporterRegistry();

        vm.expectCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion.selector
            )
        );

        app = new ExampleSourceApp(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
             address(this),
              address(mockERC20));
    }

    /**
     * Send tokens unit tests
     */

    function testSendToSameChain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = DEFAULT_SOURCE_BLOCKCHAIN_ID;
        vm.expectRevert(_formatTokenSourceErrorMessage("cannot bridge to same chain"));
        app.send(input, 0);
    }

    function testZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatTokenSourceErrorMessage("zero destination bridge address"));
        app.send(input, 0);
    }

    function testZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        vm.expectRevert(_formatTokenSourceErrorMessage("zero recipient address"));
        app.send(input, 0);
    }

    function testZeroSendAmount() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("zero send amount"));
        app.send(_createDefaultSendTokensInput(), 0);
    }

    function testInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = 1;
        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient amount to cover fees"));
        app.send(input, input.primaryFee);
    }

    function testSendWithFees() public {
        uint256 amount = 2;
        uint256 primaryFee = 1;
        _sendSuccess(amount, primaryFee);
    }

    function testSendNoFees() public {
        uint256 amount = 2;
        uint256 primaryFee = 0;
        _sendSuccess(amount, primaryFee);
    }

    /**
     * Receive tokens unit tests
     */

    function testInsufficientBridgeBalance() public {
        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient bridge balance"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(
                SendTokensInput({
                    destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    destinationBridgeAddress: address(this),
                    recipient: DEFAULT_RECIPIENT_ADDRESS,
                    primaryFee: 0,
                    secondaryFee: 0,
                    allowedRelayerAddresses: new address[](0)
                }),
                1
            )
        );
    }

    function testInvalidDestinationBridgeAddress() public {
        // First send to destination blockchain to increase the bridge balance
        _sendSuccess(2, 0);
        vm.expectRevert(_formatTokenSourceErrorMessage("invalid bridge address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(
                SendTokensInput({
                    destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
                    destinationBridgeAddress: address(0),
                    recipient: DEFAULT_RECIPIENT_ADDRESS,
                    primaryFee: 0,
                    secondaryFee: 0,
                    allowedRelayerAddresses: new address[](0)
                }),
                1
            )
        );
    }

    function testMultiHopTransfer() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSuccess(amount, 0);

        uint256 feeAmount = 1;
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });

        _checkExpectedTeleporterCalls(input, bridgedAmount);

        vm.expectEmit(true, true, true, true, address(app));
        emit SendTokens(_MOCK_MESSAGE_ID, address(MOCK_TELEPORTER_MESSENGER_ADDRESS), bridgedAmount);

        // Expect to call {TeleporterTokenSource-_send} for a multihop transfer
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(input, amount)
        );
    }

    function testMultiHopTransferFails() public {
        // First send to destination blockchain to increase the bridge balance
        uint256 amount = 2;
        _sendSuccess(amount, 0);
        uint256 balanceBefore =
            app.bridgedBalances(DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS);
        assertEq(balanceBefore, amount);

        // Fail due to insufficient amount to cover fees
        uint256 feeAmount = amount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });

        vm.expectRevert(_formatTokenSourceErrorMessage("insufficient amount to cover fees"));
        // Expect to call {TeleporterTokenSource-_send} for a multihop transfer
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            DEFAULT_DESTINATION_ADDRESS,
            abi.encode(input, amount)
        );

        // Make sure the bridge balance is still the same
        assertEq(
            app.bridgedBalances(DEFAULT_DESTINATION_BLOCKCHAIN_ID, DEFAULT_DESTINATION_ADDRESS),
            balanceBefore
        );
    }

    function _sendSuccess(uint256 amount, uint256 feeAmount) internal {
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: feeAmount,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });

        _checkExpectedTeleporterCalls(input, bridgedAmount);

        vm.expectEmit(true, true, true, true, address(app));
        emit SendTokens(_MOCK_MESSAGE_ID, address(this), bridgedAmount);
        app.send(input, amount);
    }

    function _checkExpectedTeleporterCalls(
        SendTokensInput memory input,
        uint256 bridgeAmount
    ) internal {
        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: input.primaryFee}),
            requiredGasLimit: 0,
            allowedRelayerAddresses: input.allowedRelayerAddresses,
            message: abi.encode(DEFAULT_RECIPIENT_ADDRESS, bridgeAmount)
        });

        if (input.primaryFee > 0) {
            vm.expectCall(
                address(mockERC20),
                abi.encodeCall(
                    IERC20.allowance, (address(app), address(MOCK_TELEPORTER_MESSENGER_ADDRESS))
                )
            );
        }
        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );
    }

    function _initMockTeleporterRegistry() internal {
        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry(MOCK_TELEPORTER_REGISTRY_ADDRESS).latestVersion.selector
            ),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(
                TeleporterRegistry.getVersionFromAddress.selector,
                (MOCK_TELEPORTER_MESSENGER_ADDRESS)
            ),
            abi.encode(1)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getAddressFromVersion.selector, (1)),
            abi.encode(MOCK_TELEPORTER_MESSENGER_ADDRESS)
        );

        vm.mockCall(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            abi.encodeWithSelector(TeleporterRegistry.getLatestTeleporter.selector),
            abi.encode(ITeleporterMessenger(MOCK_TELEPORTER_MESSENGER_ADDRESS))
        );
    }

    function _createDefaultSendTokensInput() internal pure returns (SendTokensInput memory) {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            destinationBridgeAddress: DEFAULT_DESTINATION_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });
    }

    function _formatTokenSourceErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("TeleporterTokenSource: ", errorMessage));
    }
}
