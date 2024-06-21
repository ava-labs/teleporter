// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {ERC20TokenBridgeTest} from "./ERC20TokenBridgeTests.t.sol";
import {TokenHubTest} from "./TokenHubTests.t.sol";
import {IERC20SendAndCallReceiver} from "../src/interfaces/IERC20SendAndCallReceiver.sol";
import {SendTokensInput} from "../src/interfaces/ITokenBridge.sol";
import {ERC20TokenHub} from "../src/TokenHub/ERC20TokenHub.sol";
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {ExampleERC20} from "../lib/teleporter/contracts/src/Mocks/ExampleERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/utils/SafeERC20.sol";
import {TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {TokenScalingUtils} from "../src/utils/TokenScalingUtils.sol";

contract ERC20TokenHubTest is ERC20TokenBridgeTest, TokenHubTest {
    using SafeERC20 for IERC20;

    ERC20TokenHub public app;
    IERC20 public mockERC20;

    function setUp() public override {
        TokenHubTest.setUp();

        mockERC20 = new ExampleERC20();
        tokenHubDecimals = 6;
        app = deployERC20TokenHub(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            address(mockERC20),
            tokenHubDecimals
        );
        erc20Bridge = app;
        tokenHub = app;
        tokenBridge = app;

        bridgedToken = mockERC20;
    }

    function deployERC20TokenHub(
        address teleporterRegistryAddress,
        address teleporterManager,
        address tokenAddress,
        uint8 tokenDecimals
    ) internal returns (ERC20TokenHub) {
        ERC20TokenHub instance = new ERC20TokenHub();
        instance.initialize(
            teleporterRegistryAddress, teleporterManager, tokenAddress, tokenDecimals
        );
        return instance;
    }

    /**
     * Initialization unit tests
     */
    function testZeroTeleporterRegistryAddress() public {
        deployExpectError(
            address(0),
            address(this),
            address(mockERC20),
            tokenHubDecimals,
            "TeleporterUpgradeable: zero teleporter registry address"
        );
    }

    function testZeroTeleporterManagerAddress() public {
        deployExpectError(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(0),
            address(mockERC20),
            tokenHubDecimals,
            "Ownable: new owner is the zero address"
        );
    }

    function testZeroFeeTokenAddress() public {
        deployExpectError(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(0),
            tokenHubDecimals,
            _formatErrorMessage("zero token address")
        );
    }

    function testTokenDecimalsTooHigh() public {
        deployExpectError(
            MOCK_TELEPORTER_REGISTRY_ADDRESS,
            address(this),
            address(mockERC20),
            uint8(TokenScalingUtils.MAX_TOKEN_DECIMALS) + 1,
            _formatErrorMessage("token decimals too high")
        );
    }

    function testReceiveZeroHubTokenAmount() public {
        // Set up a registered spoke that will scale down the received amount
        // to zero hub tokens.
        uint256 tokenMultiplier = 100_000;
        _setUpRegisteredSpoke(
            DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID, DEFAULT_TOKEN_SPOKE_ADDRESS, 0, tokenMultiplier, true
        );

        // Send over hub token to the spoke
        // and check for expected calls for scaled amount of tokens sent.
        SendTokensInput memory input = _createDefaultSendTokensInput();
        uint256 amount = 1;
        _setUpExpectedDeposit(amount, input.primaryFee);

        uint256 scaledAmount = tokenMultiplier * amount;
        _checkExpectedTeleporterCallsForSend(
            _createSingleHopTeleporterMessageInput(input, scaledAmount)
        );
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, scaledAmount);
        _send(input, amount);

        // Receive an amount from spoke less than `scaledAmount`
        // which will be scaled down to zero hub tokens.
        vm.expectRevert(_formatErrorMessage("zero token amount"));
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        tokenHub.receiveTeleporterMessage(
            DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID,
            DEFAULT_TOKEN_SPOKE_ADDRESS,
            _encodeSingleHopSendMessage(scaledAmount - 1, DEFAULT_RECIPIENT_ADDRESS)
        );
    }

    function testRegisterDestinationRoundUpCollateralNeeded() public {
        _setUpRegisteredSpoke(
            DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID, DEFAULT_TOKEN_SPOKE_ADDRESS, 11, 10, true
        );
        (, uint256 collateralNeeded,,) = tokenHub.registeredSpokes(
            DEFAULT_TOKEN_SPOKE_BLOCKCHAIN_ID, DEFAULT_TOKEN_SPOKE_ADDRESS
        );
        assertEq(collateralNeeded, 2);
    }

    function testSendScaledUpAmount() public {
        uint256 amount = 100;
        uint256 feeAmount = 1;

        SendTokensInput memory input = _createDefaultSendTokensInput();
        input.primaryFee = feeAmount;

        // Raw amount sent over wire should be multipled by 1e2.
        uint256 tokenMultiplier = 1e2;
        _setUpRegisteredSpoke(
            input.destinationBlockchainID, input.destinationBridgeAddress, 0, tokenMultiplier, true
        );
        _setUpExpectedDeposit(amount, input.primaryFee);
        TeleporterMessageInput memory expectedMessage = TeleporterMessageInput({
            destinationBlockchainID: input.destinationBlockchainID,
            destinationAddress: input.destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: address(bridgedToken),
                amount: input.primaryFee
            }),
            requiredGasLimit: input.requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: _encodeSingleHopSendMessage(amount * tokenMultiplier, DEFAULT_RECIPIENT_ADDRESS)
        });
        _checkExpectedTeleporterCallsForSend(expectedMessage);
        vm.expectEmit(true, true, true, true, address(tokenBridge));
        emit TokensSent(_MOCK_MESSAGE_ID, address(this), input, amount * tokenMultiplier);
        _send(input, amount);
    }

    function _checkExpectedWithdrawal(address recipient, uint256 amount) internal override {
        vm.expectEmit(true, true, true, true, address(tokenHub));
        emit TokensWithdrawn(recipient, amount);
        vm.expectCall(
            address(mockERC20), abi.encodeCall(IERC20.transfer, (address(recipient), amount))
        );
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Transfer(address(app), recipient, amount);
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
        // The recipient contract will be approved to spend the tokens from the token hub contract.
        vm.expectEmit(true, true, true, true, address(mockERC20));
        emit Approval(address(app), DEFAULT_RECIPIENT_CONTRACT_ADDRESS, amount);

        if (targetHasCode) {
            vm.etch(recipient, new bytes(1));

            bytes memory expectedCalldata = abi.encodeCall(
                IERC20SendAndCallReceiver.receiveTokens,
                (
                    sourceBlockchainID,
                    originInfo.bridgeAddress,
                    originInfo.senderAddress,
                    address(mockERC20),
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
        vm.expectEmit(true, true, true, true, address(mockERC20));
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
            vm.expectEmit(true, true, true, true, address(mockERC20));
            emit Transfer(address(app), address(DEFAULT_FALLBACK_RECIPIENT_ADDRESS), amount);
        }
    }

    function _addCollateral(
        bytes32 spokeBlockchainID,
        address spokeBridgeAddress,
        uint256 amount
    ) internal override {
        app.addCollateral(spokeBlockchainID, spokeBridgeAddress, amount);
    }

    function _setUpDeposit(uint256 amount) internal virtual override {
        // Increase the allowance of the bridge to transfer the funds from the user
        bridgedToken.safeIncreaseAllowance(address(tokenBridge), amount);
    }

    function _setUpExpectedZeroAmountRevert() internal override {
        vm.expectRevert("SafeERC20TransferFrom: balance not increased");
    }

    function _setUpExpectedDeposit(uint256 amount, uint256 feeAmount) internal virtual override {
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
        // Increase the allowance of the bridge to transfer the funds from the user
        bridgedToken.safeIncreaseAllowance(address(tokenBridge), amount);

        // Check that transferFrom is called to deposit the funds sent from the user to the bridge
        // This is the case because for the {ERC20TokenHub) is not the fee token itself
        vm.expectCall(
            address(bridgedToken),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(tokenBridge), amount))
        );
        vm.expectEmit(true, true, true, true, address(bridgedToken));
        emit Transfer(address(this), address(tokenBridge), amount);
    }

    function deployExpectError(
        address teleporterRegistryAddress,
        address teleporterManager,
        address tokenAddress,
        uint8 tokenDecimals,
        bytes memory expectedError
    ) private {
        ERC20TokenHub instance = new ERC20TokenHub();
        vm.expectRevert(expectedError);
        instance.initialize(
            teleporterRegistryAddress, teleporterManager, tokenAddress, tokenDecimals
        );
    }
}
