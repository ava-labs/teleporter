// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ERC20TokenBridgeTest} from "./ERC20TokenBridgeTests.t.sol";
import {TokenSpokeTest} from "./TokenSpokeTests.t.sol";
import {IERC20SendAndCallReceiver} from "../src/interfaces/IERC20SendAndCallReceiver.sol";
import {TokenSpoke} from "../src/TokenSpoke/TokenSpoke.sol";
import {TokenSpokeSettings} from "../src/TokenSpoke/interfaces/ITokenSpoke.sol";
import {ERC20TokenSpoke} from "../src/TokenSpoke/ERC20TokenSpoke.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";
import {SendTokensInput} from "../src/interfaces/ITokenBridge.sol";

contract ERC20TokenSpokeTest is ERC20TokenBridgeTest, TokenSpokeTest {
    using SafeERC20 for IERC20;

    string public constant MOCK_TOKEN_NAME = "Test Token";
    string public constant MOCK_TOKEN_SYMBOL = "TST";

    ERC20TokenSpoke public app;

    function setUp() public virtual override {
        TokenSpokeTest.setUp();

        tokenDecimals = 14;
        tokenHubDecimals = 18;
        app = ERC20TokenSpoke(address(_createNewSpokeInstance()));

        erc20Bridge = app;
        tokenSpoke = app;
        tokenBridge = app;
        bridgedToken = IERC20(app);

        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(this), 10e18);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        app.receiveTeleporterMessage(
            DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_HUB_ADDRESS,
            _encodeSingleHopSendMessage(10e18, address(this))
        );
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20TokenSpoke(
            TokenSpokeSettings({
                teleporterRegistryAddress: address(0),
                teleporterManager: address(this),
                tokenHubBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
                tokenHubAddress: DEFAULT_TOKEN_HUB_ADDRESS,
                tokenHubDecimals: tokenHubDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
    }

    function testZeroTeleporterManagerAddress() public {
        vm.expectRevert("Ownable: new owner is the zero address");
        new ERC20TokenSpoke(
            TokenSpokeSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(0),
                tokenHubBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
                tokenHubAddress: DEFAULT_TOKEN_HUB_ADDRESS,
                tokenHubDecimals: tokenHubDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
    }

    function testZeroTokenHubBlockchainID() public {
        vm.expectRevert(_formatErrorMessage("zero token hub blockchain ID"));
        new ERC20TokenSpoke(
            TokenSpokeSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                tokenHubBlockchainID: bytes32(0),
                tokenHubAddress: DEFAULT_TOKEN_HUB_ADDRESS,
                tokenHubDecimals: tokenHubDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
    }

    function testDecimals() public {
        uint8 res = app.decimals();
        assertEq(tokenDecimals, res);
    }

    function testDeployToSameBlockchain() public {
        vm.expectRevert(_formatErrorMessage("cannot deploy to same blockchain as token hub"));
        new ERC20TokenSpoke(
            TokenSpokeSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                tokenHubBlockchainID: DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID,
                tokenHubAddress: DEFAULT_TOKEN_HUB_ADDRESS,
                tokenHubDecimals: tokenHubDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
    }

    function testZeroTokenHubAddress() public {
        vm.expectRevert(_formatErrorMessage("zero token hub address"));
        new ERC20TokenSpoke(
            TokenSpokeSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                tokenHubBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
                tokenHubAddress: address(0),
                tokenHubDecimals: 18
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            18
        );
    }

    function testSendWithSeparateFeeAsset() public {
        uint256 amount = 200_000;
        uint256 feeAmount = 100;
        ExampleERC20 separateFeeAsset = new ExampleERC20();
        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFeeTokenAddress = address(separateFeeAsset);
        input.primaryFee = feeAmount;

        IERC20(separateFeeAsset).safeIncreaseAllowance(address(tokenBridge), feeAmount);
        vm.expectCall(
            address(separateFeeAsset),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenBridge), feeAmount))
        );
        // Increase the allowance of the bridge to transfer the funds from the user
        bridgedToken.safeIncreaseAllowance(address(tokenBridge), amount);

        vm.expectEmit(true, true, true, true, address(bridgedToken));
        emit Transfer(address(this), address(0), amount);
        _checkExpectedTeleporterCallsForSend(_createSingleHopTeleporterMessageInput(input, amount));
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount);
        _send(input, amount);
    }

    function _createNewSpokeInstance() internal override returns (TokenSpoke) {
        return new ERC20TokenSpoke(
            TokenSpokeSettings({
                teleporterRegistryAddress: MOCK_TELEPORTER_REGISTRY_ADDRESS,
                teleporterManager: address(this),
                tokenHubBlockchainID: DEFAULT_TOKEN_HUB_BLOCKCHAIN_ID,
                tokenHubAddress: DEFAULT_TOKEN_HUB_ADDRESS,
                tokenHubDecimals: tokenHubDecimals
            }),
            MOCK_TOKEN_NAME,
            MOCK_TOKEN_SYMBOL,
            tokenDecimals
        );
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenSpoke));
        emit TokensWithdrawn(recipient, amount);
        vm.expectEmit(true, true, true, true, address(tokenSpoke));
        emit Transfer(address(0), recipient, amount);
    }

    function _setUpExpectedSendAndCall(
        bytes32 sourceBlockchainID,
        OriginSenderInfo memory originInfo,
        address recipient,
        uint256 amount,
        bytes memory payload,
        uint256 gasLimit,
        bool targetHasCode,
        bool expectSuccess
    ) internal override {
        // The bridge tokens will be minted to the contract itself
        vm.expectEmit(true, true, true, true, address(app));
        emit Transfer(address(0), address(tokenSpoke), amount);

        // Then recipient contract is then approved to spend them
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        if (targetHasCode) {
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                IERC20SendAndCallReceiver.receiveTokens,
                (
                    sourceBlockchainID,
                    originInfo.bridgeAddress,
                    originInfo.senderAddress,
                    address(app),
                    amount,
                    payload
                )
            );
            if (expectSuccess) {
                vm.mockCall(recipient, expectedCalldata, new bytes(0));
            } else {
                vm.mockCallRevert(recipient, expectedCalldata, new bytes(0));
            }
            vm.expectCall(recipient, 0, uint64(gasLimit), expectedCalldata);
        } else {
            vm.etch(recipient, new bytes(0));
        }

        // Then recipient contract approval is reset
        vm.expectEmit(true, true, true, true, address(app));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, 0);

        if (targetHasCode && expectSuccess) {
            // The call should have succeeded.
            vm.expectEmit(true, true, true, true, address(app));
            emit CallSucceeded(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);
        } else {
            // The call should have failed.
            vm.expectEmit(true, true, true, true, address(app));
            emit CallFailed(DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

            // Then the amount should be sent to the fallback recipient.
            vm.expectEmit(true, true, true, true, address(app));
            emit Transfer(address(app), address(DEFAULT_FALLBACK_RECIPIENT_ADDRESS), amount);
        }
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert(_formatErrorMessage("insufficient tokens to transfer"));
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal virtual override {
        // Transfer the fee to the bridge if it is greater than 0
        if (feeAmount > 0) {
            bridgedToken.safeIncreaseAllowance(address(tokenBridge), feeAmount);
        }

        // Increase the allowance of the bridge to transfer the funds from the user
        bridgedToken.safeIncreaseAllowance(address(tokenBridge), amount);

        uint256 currentAllowance = bridgedToken.allowance(address(this), address(tokenBridge));
        if (feeAmount > 0) {
            vm.expectEmit(true, true, true, true, address(bridgedToken));
            emit Approval(address(this), address(tokenBridge), currentAllowance - feeAmount);
            vm.expectEmit(true, true, true, true, address(bridgedToken));
            emit Transfer(address(this), address(tokenBridge), feeAmount);
        }
        vm.expectEmit(true, true, true, true, address(bridgedToken));
        emit Approval(address(this), address(tokenBridge), currentAllowance - feeAmount - amount);
        vm.expectEmit(true, true, true, true, address(bridgedToken));
        emit Transfer(address(this), address(0), amount);
    }

    function _getTotalSupply() internal view override returns (uint256) {
        return app.totalSupply();
    }

    function _setUpMockMint(address, uint256) internal pure override {
        // Don't need to mock the minting of an ERC20TokenSpoke since it is an internal call
        // on the spoke contract.
        return;
    }
}
