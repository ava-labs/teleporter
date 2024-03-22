// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {IERC20BridgeTest} from "./IERC20BridgeTests.t.sol";
import {ERC20Destination} from "../src/ERC20Destination.sol";
import {
    ITeleporterTokenBridge, SendTokensInput
} from "../src/interfaces/ITeleporterTokenBridge.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {IWarpMessenger} from
    "@avalabs/subnet-evm-contracts@1.2.0/contracts/interfaces/IWarpMessenger.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";

contract ERC20DestinationTest is IERC20BridgeTest {
    using SafeERC20 for IERC20;

    event Transfer(address indexed from, address indexed to, uint256 value);

    ERC20Destination public app;

    string public constant MOCK_TOKEN_NAME = "Test Token";
    string public constant MOCK_TOKEN_SYMBOL = "TST";
    uint8 public constant MOCK_TOKEN_DECIMALS = 18;

    function setUp() public virtual {
        // super.setUp();
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

        app = new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS
        );
        erc20Bridge = app;
        tokenBridge = app;
        feeToken = IERC20(app);

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(this), 10e18);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, abi.encode(address(this), 10e18)
        );
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20Destination(
            address(0),
            address(this),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS
        );
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS
        );
    }

    function testZeroBlockchainID() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero source blockchain ID"));
        new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            bytes32(0),
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS
        );
    }

    function testDeployToSameBlockchain() public {
        vm.expectRevert(
            _formatTokenDestinationErrorMessage("cannot deploy to same blockchain as source")
        );
        new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS
        );
    }

    function testZeroTokenSourceAddress() public {
        vm.expectRevert(_formatTokenDestinationErrorMessage("zero token source address"));
        new ERC20Destination(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            address(0),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            MOCK_TOKEN_DECIMALS
        );
    }

    /**
     * Send tokens unit tests
     */
    function testInvalidSendingBackToSourceBlockchain() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBridgeAddress = address(this);
        vm.expectRevert(_formatTokenDestinationErrorMessage("invalid destination bridge address"));
        app.send(input, 0);
    }

    function testSendingToSameInstance() public {
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = app.blockchainID();
        input.destinationBridgeAddress = address(app);
        vm.expectRevert(_formatTokenDestinationErrorMessage("invalid destination bridge address"));
        app.send(input, 0);
    }

    function testSendToSameBlockchainDifferentDestination() public {
        // Send a transfer to the same app itself
        uint256 amount = 2;
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.destinationBlockchainID = app.blockchainID();
        input.destinationBridgeAddress = address(this);

        _sendSuccess(amount, input.primaryFee);
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

    function _createDefaultSendTokensInput()
        internal
        pure
        override
        returns (SendTokensInput memory)
    {
        return SendTokensInput({
            destinationBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            destinationBridgeAddress: TOKEN_SOURCE_ADDRESS,
            recipient: DEFAULT_RECIPIENT_ADDRESS,
            primaryFee: 0,
            secondaryFee: 0,
            allowedRelayerAddresses: new address[](0)
        });
    }

    function _formatErrorMessage(string memory message)
        internal
        pure
        override
        returns (bytes memory)
    {
        return _formatTokenDestinationErrorMessage(message);
    }

    function _requiredGasLimit() internal view virtual override returns (uint256) {
        return app.SEND_TOKENS_REQUIRED_GAS();
    }

    function _encodeMessage(
        SendTokensInput memory input,
        uint256 amount
    ) internal pure virtual override returns (bytes memory) {
        return abi.encode(
            SendTokensInput({
                destinationBlockchainID: input.destinationBlockchainID,
                destinationBridgeAddress: input.destinationBridgeAddress,
                recipient: input.recipient,
                primaryFee: input.secondaryFee,
                secondaryFee: 0,
                allowedRelayerAddresses: input.allowedRelayerAddresses
            }),
            amount
        );
    }

    function _formatTokenDestinationErrorMessage(string memory errorMessage)
        internal
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("TeleporterTokenDestination: ", errorMessage));
    }
}
