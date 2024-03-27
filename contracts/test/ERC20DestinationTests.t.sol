// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ERC20BridgeTest} from "./ERC20BridgeTests.t.sol";
import {TeleporterTokenDestinationTest} from "./TeleporterTokenDestinationTests.t.sol";
import {ERC20Destination} from "../src/ERC20Destination.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";

contract ERC20DestinationTest is ERC20BridgeTest, TeleporterTokenDestinationTest {
    using SafeERC20 for IERC20;

    ERC20Destination public app;

    string public constant MOCK_TOKEN_NAME = "Test Token";
    string public constant MOCK_TOKEN_SYMBOL = "TST";
    uint8 public constant MOCK_TOKEN_DECIMALS = 18;

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
            DEFAULT_SOURCE_BLOCKCHAIN_ID, TOKEN_SOURCE_ADDRESS, abi.encode(address(this), 10e18)
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

    function testZeroBlockchainID() public {
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

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), recipient, amount);
    }
}
