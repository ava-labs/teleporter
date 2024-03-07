// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {TeleporterTokenDestination, IWarpMessenger} from "../src/TeleporterTokenDestination.sol";
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

contract ExampleDestinationApp is TeleporterTokenDestination {
    constructor(
        address teleporterRegistryAddress,
        address teleporterManager,
        bytes32 sourceBlockchainID,
        address tokenSourceAddress,
        address feeTokenAddress
    )
        TeleporterTokenDestination(
            teleporterRegistryAddress,
            teleporterManager,
            sourceBlockchainID,
            tokenSourceAddress,
            feeTokenAddress
        )
    {}

    function send(SendTokensInput calldata input, uint256 amount) external {
        _send(input, amount);
    }

    function _withdraw(address recipient, uint256 amount) internal virtual override {}

    function _burn(uint256 amount) internal virtual override {}
}

contract TeleporterTokenDestinationTest is Test {
    bytes32 public constant DEFAULT_SOURCE_BLOCKCHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    bytes32 public constant DEFAULT_DESTINATION_BLOCKCHAIN_ID =
        bytes32(hex"1234567812345678123456781234567812345678123456781234567812345678");
    address public constant TOKEN_SOURCE_ADDRESS = 0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address public constant DEFAULT_RECIPIENT_ADDRESS = 0xABCDabcdABcDabcDaBCDAbcdABcdAbCdABcDABCd;
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;

    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");

    ExampleDestinationApp public app;
    UnitTestMockERC20 public mockERC20;

    event SendTokens(bytes32 indexed teleporterMessageID, address indexed sender, uint256 amount);

    function setUp() public virtual {
        mockERC20 = new UnitTestMockERC20();
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(DEFAULT_DESTINATION_BLOCKCHAIN_ID)
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

        app = new ExampleDestinationApp(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            address(mockERC20));
    }

    /**
     * Send tokens unit tests
     */

    function testSendToSameChain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = DEFAULT_DESTINATION_BLOCKCHAIN_ID;
        vm.expectRevert(_formatTokenDestinationErrorMessage("cannot bridge to same chain"));
        app.send(input, 0);
    }

    function testZeroDestinationBridge() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(0);
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero destination bridge address"));
        app.send(input, 0);
    }

    function testZeroRecipient() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.recipient = address(0);
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero recipient address"));
        app.send(input, 0);
    }

    function testZeroSendAmount() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero send amount"));
        app.send(_createDefaultSendTokensInput(), 0);
    }

    function testInsufficientAmountToCoverFees() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = 1;
        vm.expectRevert(_formatTokenDestinationErrorMessage("insufficient amount to cover fees"));
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

    function testInvalidSourceChain() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("invalid source chain"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_DESTINATION_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, new bytes(0)
        );
    }

    function testInvalidTokenSourceAddress() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("invalid token source address"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, address(0), new bytes(0));
    }

    function testReceivedInvalidMessage() public {
        vm.expectRevert();
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, bytes("test")
        );
    }

    function _sendSuccess(uint256 amount, uint256 feeAmount) internal {
        uint256 bridgedAmount = amount - feeAmount;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

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
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationAddress: TOKEN_SOURCE_ADDRESS,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: input.primaryFee}),
            requiredGasLimit: 0,
            allowedRelayerAddresses: input.allowedRelayerAddresses,
            message: abi.encode(
                SendTokensInput({
                    destinationBlockchainID: input.destinationBlockchainID,
                    destinationBridgeAddress: input.destinationBridgeAddress,
                    recipient: input.recipient,
                    primaryFee: input.secondaryFee,
                    secondaryFee: 0,
                    allowedRelayerAddresses: input.allowedRelayerAddresses
                }),
                bridgeAmount
                )
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
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: TOKEN_SOURCE_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });
    }

    function _formatTokenDestinationErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("TeleporterTokenDestination: ", errorMessage));
    }
}
