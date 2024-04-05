// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ERC20BridgeTest} from "./ERC20BridgeTests.t.sol";
import {TeleporterTokenDestinationTest} from "./TeleporterTokenDestinationTests.t.sol";
import {IERC20SendAndCallReceiver} from "../src/interfaces/IERC20SendAndCallReceiver.sol";
import {ERC20Destination} from "../src/ERC20Destination.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

contract ERC20DestinationTest is ERC20BridgeTest, TeleporterTokenDestinationTest {
    using SafeERC20 for IERC20;

    string public constant MOCK_TOKEN_NAME = "Test Token";
    string public constant MOCK_TOKEN_SYMBOL = "TST";
    uint8 public constant MOCK_TOKEN_DECIMALS = 18;

    ERC20Destination public app;

    function setUp() public virtual override {
        TeleporterTokenDestinationTest.setUp();

        app = new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });

        erc20Bridge = app;
        tokenDestination = app;
        tokenBridge = app;
        feeToken = IERC20(app);

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(this), 10e18);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID,
            TOKEN_SOURCE_ADDRESS,
            _encodeSingleHopSendMessage(10e18, address(this))
        );
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20Destination({
            teleporterRegistryAddress: address(0),
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(0),
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testZeroSourceBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("zero source blockchain ID"));
        new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: bytes32(0),
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testDecimals() public {
        uint8 res = app.decimals();
        assertEq(MOCK_TOKEN_DECIMALS, res);
    }

    function testDeployToSameBlockchain() public {
        vm.expectRevert(_formatErrorMessage("cannot deploy to same blockchain as source"));
        new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_DESTINATION_BLOCKCHAIN_ID,
            tokenSourceAddress: TOKEN_SOURCE_ADDRESS,
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testZeroTokenSourceAddress() public {
        vm.expectRevert(_formatErrorMessage("zero token source address"));
        new ERC20Destination({
            teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
            teleporterManager: address(this),
            sourceBlockchainID: DEFAULT_SOURCE_BLOCKCHAIN_ID,
            tokenSourceAddress: address(0),
            tokenName: MOCK_TOKEN_NAME,
            tokenSymbol: MOCK_TOKEN_SYMBOL,
            tokenDecimals: MOCK_TOKEN_DECIMALS
        });
    }

    function testReceiveSendAndCallSuccess() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        // The bridge tokens will be minted to the contract itself
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(tokenDestination), amount);

        // Then recipient contract is then approved to spend them
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        // Then a call to the recipient is made.
        bytes memory expectedCalldata =
            abi.encodeCall(IERC20SendAndCallReceiver.receiveTokens, (address(app), amount, payload));
        vm.etch(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, new bytes(1));
        vm.mockCall(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, expectedCalldata, new bytes(0));
        vm.expectCall(
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            0,
            uint64(DEFAULT_RECIPIENT_GAS_LIMIT),
            expectedCalldata
        );

        // Then recipient contract approval is reset
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, 0);

        // The call should have succeeded.
        vm.expectEmit(true, true, true, true, address(app));
        emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        bytes memory message = _encodeSingleHopCallMessage(
            amount,
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            payload,
            DEFAULT_RECIPIENT_GAS_LIMIT,
            DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        );
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function testReceiveSendAndCallFailure() public {
        uint256 amount = 2;
        bytes memory payload = hex"DEADBEEF";

        // The bridge tokens will be minted to the contract itself
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(tokenDestination), amount);

        // Then recipient contract is then approved to spend them
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        // Then a call to the recipient is made - mock it reverting
        bytes memory expectedCalldata =
            abi.encodeCall(IERC20SendAndCallReceiver.receiveTokens, (address(app), amount, payload));
        vm.etch(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, new bytes(1));
        vm.mockCallRevert(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, expectedCalldata, new bytes(0));
        vm.expectCall(
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            0,
            uint64(DEFAULT_RECIPIENT_GAS_LIMIT),
            expectedCalldata
        );

        // Then recipient contract approval is reset
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, 0);

        // The call should have succeeded.
        vm.expectEmit(true, true, true, true, address(app));
        emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        // Then the amount should be sent to the fallback recipient.
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(app), address(DEFAULT_FALLBACK_RECIPIENT_ADDRESS), amount);

        bytes memory message = _encodeSingleHopCallMessage(
            amount,
            DEFAULT_RECIPIENT_CONTRACT_ADDRESS,
            payload,
            DEFAULT_RECIPIENT_GAS_LIMIT,
            DEFAULT_FALLBACK_RECIPIENT_ADDRESS
        );
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenDestination.receiveTeleporterMessage(
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, message
        );
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenDestination));
        emit Transfer(address(0), recipient, amount);
    }
}
