// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {ERC20Bridge, BridgeToken, IERC20, ERC20, IWarpMessenger} from "../ERC20Bridge.sol";
import {
    ITeleporterMessenger,
    TeleporterMessageInput,
    TeleporterFeeInfo
} from "@teleporter/ITeleporterMessenger.sol";
import {TeleporterRegistry} from "@teleporter/upgrades/TeleporterRegistry.sol";
import {UnitTestMockERC20} from "@mocks/UnitTestMockERC20.sol";

contract ERC20BridgeTest is Test {
    address public constant MOCK_TELEPORTER_MESSENGER_ADDRESS =
        0x644E5b7c5D4Bc8073732CEa72c66e0BB90dFC00f;
    address public constant MOCK_TELEPORTER_REGISTRY_ADDRESS =
        0xf9FA4a0c696b659328DDaaBCB46Ae4eBFC9e68e4;
    address public constant WARP_PRECOMPILE_ADDRESS =
        address(0x0200000000000000000000000000000000000005);
    bytes32 internal constant _MOCK_MESSAGE_ID =
        bytes32(hex"1111111111111111111111111111111111111111111111111111111111111111");
    bytes32 private constant _MOCK_BLOCKCHAIN_ID = bytes32(uint256(123456));
    bytes32 private constant _DEFAULT_OTHER_CHAIN_ID =
        bytes32(hex"abcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd");
    address private constant _DEFAULT_OTHER_BRIDGE_ADDRESS =
        0xd54e3E251b9b0EEd3ed70A858e927bbC2659587d;
    address private constant _DEFAULT_OWNER_ADDRESS = 0x1234512345123451234512345123451234512345;
    string private constant _DEFAULT_TOKEN_NAME = "Test Token";
    string private constant _DEFAULT_SYMBOL = "TSTTK";
    uint8 private constant _DEFAULT_DECIMALS = 18;
    address private constant _DEFAULT_RECIPIENT = 0xa4CEE7d1aF6aDdDD33E3b1cC680AB84fdf1b6d1d;

    ERC20Bridge public erc20Bridge;
    UnitTestMockERC20 public mockERC20;

    event BridgeTokens(
        address indexed tokenContractAddress,
        bytes32 indexed destinationBlockchainID,
        bytes32 indexed teleporterMessageID,
        address destinationBridgeAddress,
        address recipient,
        uint256 amount
    );

    event SubmitCreateBridgeToken(
        bytes32 indexed destinationBlockchainID,
        address indexed destinationBridgeAddress,
        address indexed nativeContractAddress,
        bytes32 teleporterMessageID
    );

    event CreateBridgeToken(
        bytes32 indexed nativeBlockchainID,
        address indexed nativeBridgeAddress,
        address indexed nativeContractAddress,
        address bridgeTokenAddress
    );

    event MintBridgeTokens(address indexed contractAddress, address recipient, uint256 amount);

    function setUp() public virtual {
        vm.mockCall(
            WARP_PRECOMPILE_ADDRESS,
            abi.encodeWithSelector(IWarpMessenger.getBlockchainID.selector),
            abi.encode(_MOCK_BLOCKCHAIN_ID)
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

        erc20Bridge = new ERC20Bridge(MOCK_TELEPORTER_REGISTRY_ADDRESS, _DEFAULT_OWNER_ADDRESS);
        mockERC20 = new UnitTestMockERC20();
    }

    function testSameBlockchainID() public {
        vm.expectRevert(_formatERC20BridgeErrorMessage("cannot bridge to same chain"));
        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _MOCK_BLOCKCHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: address(mockERC20),
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: 130,
            primaryFeeAmount: 0,
            secondaryFeeAmount: 0
        });
    }

    function testInvalidFeeAmountsNativeTransfer() public {
        _setUpMockERC20ContractValues(address(mockERC20));
        _submitCreateBridgeToken(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, address(mockERC20)
        );
        vm.expectRevert(_formatERC20BridgeErrorMessage("insufficient adjusted amount"));
        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: address(mockERC20),
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: 130,
            primaryFeeAmount: 130,
            secondaryFeeAmount: 0
        });
    }

    function testInvalidFeeAmountsWrappedTransfer() public {
        address bridgeTokenAddress = _setUpBridgeToken({
            nativeBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            nativeBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nativeContractAddress: address(mockERC20),
            nativeName: _DEFAULT_TOKEN_NAME,
            nativeSymbol: _DEFAULT_SYMBOL,
            nativeDecimals: _DEFAULT_DECIMALS,
            contractNonce: 1
        });

        vm.expectRevert(_formatERC20BridgeErrorMessage("insufficient total amount"));
        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: bridgeTokenAddress,
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: 130,
            primaryFeeAmount: 127,
            secondaryFeeAmount: 3
        });
    }

    function testNativeTokenTransferFailure() public {
        _setUpMockERC20ContractValues(address(mockERC20));
        _submitCreateBridgeToken(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, address(mockERC20)
        );

        uint256 totalAmount = 123456;
        vm.mockCall(
            address(mockERC20),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(erc20Bridge), totalAmount)),
            abi.encode(false)
        );
        vm.expectCall(
            address(mockERC20),
            abi.encodeCall(IERC20.transferFrom, (address(this), address(erc20Bridge), totalAmount))
        );

        vm.expectRevert("SafeERC20: ERC20 operation did not succeed");
        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: address(mockERC20),
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: totalAmount,
            primaryFeeAmount: 0,
            secondaryFeeAmount: 0
        });
    }

    function testBridgeNativeTokensNoFee() public {
        uint256 totalAmount = 131313131313;
        _setUpExpectedTransferFromCall(address(mockERC20), totalAmount);
        _setUpMockERC20ContractValues(address(mockERC20));
        _submitCreateBridgeToken(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, address(mockERC20)
        );

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: 0}),
            requiredGasLimit: erc20Bridge.MINT_BRIDGE_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: erc20Bridge.encodeMintBridgeTokensData(
                address(mockERC20), _DEFAULT_RECIPIENT, totalAmount
                )
        });

        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit BridgeTokens({
            tokenContractAddress: address(mockERC20),
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            teleporterMessageID: _MOCK_MESSAGE_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            recipient: _DEFAULT_RECIPIENT,
            amount: totalAmount
        });

        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: address(mockERC20),
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: totalAmount,
            primaryFeeAmount: 0,
            secondaryFeeAmount: 0
        });
    }

    function testFeeApprovalFails() public {
        uint256 totalAmount = 131313131313;
        uint256 feeAmount = 131313;
        _setUpExpectedTransferFromCall(address(mockERC20), totalAmount);
        _setUpMockERC20ContractValues(address(mockERC20));
        _submitCreateBridgeToken(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, address(mockERC20)
        );

        vm.mockCall(
            address(mockERC20),
            abi.encodeCall(
                IERC20.allowance, (address(erc20Bridge), MOCK_TELEPORTER_MESSENGER_ADDRESS)
            ),
            abi.encode(0)
        );
        vm.mockCall(
            address(mockERC20),
            abi.encodeCall(IERC20.approve, (MOCK_TELEPORTER_MESSENGER_ADDRESS, feeAmount)),
            abi.encode(false)
        );
        vm.expectCall(
            address(mockERC20),
            abi.encodeCall(IERC20.approve, (MOCK_TELEPORTER_MESSENGER_ADDRESS, feeAmount))
        );

        vm.expectRevert("SafeERC20: ERC20 operation did not succeed");

        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: address(mockERC20),
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: totalAmount,
            primaryFeeAmount: feeAmount,
            secondaryFeeAmount: 0
        });
    }

    function testBridgeNativeTokensWithFee() public {
        uint256 totalAmount = 131313131313;
        uint256 feeAmount = 131313;
        uint256 bridgeAmount = totalAmount - feeAmount;
        _setUpExpectedTransferFromCall(address(mockERC20), totalAmount);
        _setUpMockERC20ContractValues(address(mockERC20));
        _submitCreateBridgeToken(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, address(mockERC20)
        );

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: feeAmount}),
            requiredGasLimit: erc20Bridge.MINT_BRIDGE_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: erc20Bridge.encodeMintBridgeTokensData(
                address(mockERC20), _DEFAULT_RECIPIENT, bridgeAmount
                )
        });

        vm.mockCall(
            address(mockERC20),
            abi.encodeCall(
                IERC20.allowance, (address(erc20Bridge), MOCK_TELEPORTER_MESSENGER_ADDRESS)
            ),
            abi.encode(0)
        );
        vm.mockCall(
            address(mockERC20),
            abi.encodeCall(IERC20.approve, (MOCK_TELEPORTER_MESSENGER_ADDRESS, feeAmount)),
            abi.encode(true)
        );
        vm.expectCall(
            address(mockERC20),
            abi.encodeCall(IERC20.approve, (MOCK_TELEPORTER_MESSENGER_ADDRESS, feeAmount))
        );

        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit BridgeTokens({
            tokenContractAddress: address(mockERC20),
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            teleporterMessageID: _MOCK_MESSAGE_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            recipient: _DEFAULT_RECIPIENT,
            amount: bridgeAmount
        });

        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: address(mockERC20),
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: totalAmount,
            primaryFeeAmount: feeAmount,
            secondaryFeeAmount: 0
        });
    }

    // Tests that "fee on transfer" ERC20 implementations are handled properly when
    // bridged using the ERC20Bridge contract.
    function testBridgeNativeFeeOnTransferTokens() public {
        uint256 totalAmount = 131313131313;
        uint256 tokenFeeOnTransferAmount = 654321;
        uint256 bridgeFeeAmount = 131313;
        uint256 expectedBridgeAmount = totalAmount - tokenFeeOnTransferAmount - bridgeFeeAmount;

        // Mock the ERC20 being a "fee on transfer" implementation;
        mockERC20.setFeeOnTransferSender(address(this), tokenFeeOnTransferAmount);

        _setUpExpectedTransferFromCall(address(mockERC20), totalAmount);
        _setUpMockERC20ContractValues(address(mockERC20));
        _submitCreateBridgeToken(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, address(mockERC20)
        );

        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(mockERC20), amount: bridgeFeeAmount}),
            requiredGasLimit: erc20Bridge.MINT_BRIDGE_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: erc20Bridge.encodeMintBridgeTokensData(
                address(mockERC20), _DEFAULT_RECIPIENT, expectedBridgeAmount
                )
        });

        vm.mockCall(
            address(mockERC20),
            abi.encodeCall(
                IERC20.allowance, (address(erc20Bridge), MOCK_TELEPORTER_MESSENGER_ADDRESS)
            ),
            abi.encode(0)
        );
        vm.mockCall(
            address(mockERC20),
            abi.encodeCall(IERC20.approve, (MOCK_TELEPORTER_MESSENGER_ADDRESS, bridgeFeeAmount)),
            abi.encode(true)
        );
        vm.expectCall(
            address(mockERC20),
            abi.encodeCall(IERC20.approve, (MOCK_TELEPORTER_MESSENGER_ADDRESS, bridgeFeeAmount))
        );

        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit BridgeTokens({
            tokenContractAddress: address(mockERC20),
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            teleporterMessageID: _MOCK_MESSAGE_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            recipient: _DEFAULT_RECIPIENT,
            amount: expectedBridgeAmount
        });

        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: address(mockERC20),
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: totalAmount,
            primaryFeeAmount: bridgeFeeAmount,
            secondaryFeeAmount: 0
        });
    }

    // Checks for the expected failure in the event that a "Fee On Transfer" token has a
    // fee so high that the remaining amount is less than the bridge fee amount.
    function testBridgeNativeFeeOnTransferAmountTooHigh() public {
        uint256 totalAmount = 717171;
        uint256 tokenFeeOnTransferAmount = 654321;
        uint256 bridgeFeeAmount = 131313;

        // Mock the ERC20 being a "fee on transfer" implementation;
        mockERC20.setFeeOnTransferSender(address(this), tokenFeeOnTransferAmount);

        _setUpExpectedTransferFromCall(address(mockERC20), totalAmount);
        _setUpMockERC20ContractValues(address(mockERC20));
        _submitCreateBridgeToken(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, address(mockERC20)
        );

        vm.expectRevert(_formatERC20BridgeErrorMessage("insufficient adjusted amount"));
        erc20Bridge.bridgeTokens({
            destinationBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            destinationBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            tokenContractAddress: address(mockERC20),
            recipient: _DEFAULT_RECIPIENT,
            totalAmount: totalAmount,
            primaryFeeAmount: bridgeFeeAmount,
            secondaryFeeAmount: 0
        });
    }

    function testNewBridgeTokenMint() public {
        uint256 amount = 654321;
        address bridgeTokenAddress = _setUpBridgeToken({
            nativeBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            nativeBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nativeContractAddress: address(mockERC20),
            nativeName: _DEFAULT_TOKEN_NAME,
            nativeSymbol: _DEFAULT_SYMBOL,
            nativeDecimals: _DEFAULT_DECIMALS,
            contractNonce: 1
        });

        bytes memory message =
            erc20Bridge.encodeMintBridgeTokensData(address(mockERC20), _DEFAULT_RECIPIENT, amount);

        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit MintBridgeTokens(bridgeTokenAddress, _DEFAULT_RECIPIENT, amount);
        erc20Bridge.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, message
        );

        // Check the values and balance of the newly created ERC20.
        BridgeToken newBridgeToken = BridgeToken(bridgeTokenAddress);
        assertEq(amount, newBridgeToken.totalSupply());
        assertEq(amount, newBridgeToken.balanceOf(_DEFAULT_RECIPIENT));
        assertEq(_DEFAULT_TOKEN_NAME, newBridgeToken.name());
        assertEq(_DEFAULT_SYMBOL, newBridgeToken.symbol());
        assertEq(_DEFAULT_DECIMALS, newBridgeToken.decimals());
        assertEq(_DEFAULT_OTHER_CHAIN_ID, newBridgeToken.nativeBlockchainID());
        assertEq(_DEFAULT_OTHER_BRIDGE_ADDRESS, newBridgeToken.nativeBridge());
        assertEq(address(mockERC20), newBridgeToken.nativeAsset());
    }

    function testMintExistingBridgeToken() public {
        address recipient1 = address(56);
        uint256 amount1 = 654321;
        address recipient2 = address(57);
        uint256 amount2 = 123456;
        address bridgeTokenAddress = _setUpBridgeToken({
            nativeBlockchainID: _DEFAULT_OTHER_CHAIN_ID,
            nativeBridgeAddress: _DEFAULT_OTHER_BRIDGE_ADDRESS,
            nativeContractAddress: address(mockERC20),
            nativeName: _DEFAULT_TOKEN_NAME,
            nativeSymbol: _DEFAULT_SYMBOL,
            nativeDecimals: _DEFAULT_DECIMALS,
            contractNonce: 1
        });

        bytes memory message1 =
            erc20Bridge.encodeMintBridgeTokensData(address(mockERC20), recipient1, amount1);

        // Call mintBridgeTokens the first time, which should create the new bridge token.
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit MintBridgeTokens(bridgeTokenAddress, recipient1, amount1);
        erc20Bridge.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, message1
        );

        bytes memory message2 =
            erc20Bridge.encodeMintBridgeTokensData(address(mockERC20), recipient2, amount2);

        // Call mintBridgeTokens the second time, which should mint additional of the existing token.
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit MintBridgeTokens(bridgeTokenAddress, recipient2, amount2);
        erc20Bridge.receiveTeleporterMessage(
            _DEFAULT_OTHER_CHAIN_ID, _DEFAULT_OTHER_BRIDGE_ADDRESS, message2
        );

        // Check the values and balance of the newly created ERC20.
        BridgeToken newBridgeToken = BridgeToken(bridgeTokenAddress);
        assertEq(amount1 + amount2, newBridgeToken.totalSupply());
        assertEq(amount1, newBridgeToken.balanceOf(recipient1));
        assertEq(amount2, newBridgeToken.balanceOf(recipient2));
        assertEq(_DEFAULT_TOKEN_NAME, newBridgeToken.name());
        assertEq(_DEFAULT_SYMBOL, newBridgeToken.symbol());
        assertEq(_DEFAULT_DECIMALS, newBridgeToken.decimals());
        assertEq(_DEFAULT_OTHER_CHAIN_ID, newBridgeToken.nativeBlockchainID());
        assertEq(_DEFAULT_OTHER_BRIDGE_ADDRESS, newBridgeToken.nativeBridge());
        assertEq(address(mockERC20), newBridgeToken.nativeAsset());
    }

    function testZeroTeleporterRegistryAddress() public {
        vm.expectRevert("TeleporterUpgradeable: zero teleporter registry address");
        new ERC20Bridge(address(0), _DEFAULT_OWNER_ADDRESS);
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

    function _setUpExpectedTransferFromCall(address tokenContract, uint256 amount) private {
        vm.expectCall(
            tokenContract,
            abi.encodeCall(IERC20.transferFrom, (address(this), address(erc20Bridge), amount))
        );
    }

    // Calls submitCreateBridgeToken of the test's ERC20Bridge instance to add the specified
    // token to be able to be sent to the specified destination bridge. Checks that the expected
    // call to the Teleporter contract is made and that the expected event is emitted. This is
    // required before attempting to call bridgeTokens for the given token and bridge.
    function _submitCreateBridgeToken(
        bytes32 destinationBlockchainID,
        address destinationBridgeAddress,
        address nativeContractAddress
    ) private {
        ERC20 nativeToken = ERC20(nativeContractAddress);
        TeleporterMessageInput memory expectedMessageInput = TeleporterMessageInput({
            destinationBlockchainID: destinationBlockchainID,
            destinationAddress: destinationBridgeAddress,
            feeInfo: TeleporterFeeInfo({feeTokenAddress: address(0), amount: 0}),
            requiredGasLimit: erc20Bridge.CREATE_BRIDGE_TOKENS_REQUIRED_GAS(),
            allowedRelayerAddresses: new address[](0),
            message: erc20Bridge.encodeCreateBridgeTokenData(
                nativeContractAddress, nativeToken.name(), nativeToken.symbol(), nativeToken.decimals()
                )
        });

        vm.mockCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput)),
            abi.encode(_MOCK_MESSAGE_ID)
        );
        vm.expectCall(
            MOCK_TELEPORTER_MESSENGER_ADDRESS,
            abi.encodeCall(ITeleporterMessenger.sendCrossChainMessage, (expectedMessageInput))
        );

        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit SubmitCreateBridgeToken(
            destinationBlockchainID,
            destinationBridgeAddress,
            nativeContractAddress,
            _MOCK_MESSAGE_ID
        );

        erc20Bridge.submitCreateBridgeToken(
            destinationBlockchainID, destinationBridgeAddress, nativeToken, address(0), 0
        );
    }

    function _setUpBridgeToken(
        bytes32 nativeBlockchainID,
        address nativeBridgeAddress,
        address nativeContractAddress,
        string memory nativeName,
        string memory nativeSymbol,
        uint8 nativeDecimals,
        uint8 contractNonce
    ) private returns (address) {
        address expectedBridgeTokenAddress =
            _deriveExpectedContractAddress(address(erc20Bridge), contractNonce);
        bytes memory message = erc20Bridge.encodeCreateBridgeTokenData(
            nativeContractAddress, nativeName, nativeSymbol, nativeDecimals
        );
        vm.prank(MOCK_TELEPORTER_MESSENGER_ADDRESS);
        vm.expectEmit(true, true, true, true, address(erc20Bridge));
        emit CreateBridgeToken(
            nativeBlockchainID,
            nativeBridgeAddress,
            nativeContractAddress,
            expectedBridgeTokenAddress
        );
        erc20Bridge.receiveTeleporterMessage(nativeBlockchainID, nativeBridgeAddress, message);
        return expectedBridgeTokenAddress;
    }

    function _setUpMockERC20ContractValues(address tokenContract) private {
        ERC20 token = ERC20(tokenContract);
        vm.mockCall(tokenContract, abi.encodeCall(token.name, ()), abi.encode(_DEFAULT_TOKEN_NAME));
        vm.expectCall(tokenContract, abi.encodeCall(token.name, ()));
        vm.mockCall(tokenContract, abi.encodeCall(token.symbol, ()), abi.encode(_DEFAULT_SYMBOL));
        vm.expectCall(tokenContract, abi.encodeCall(token.symbol, ()));
        vm.mockCall(
            tokenContract, abi.encodeCall(token.decimals, ()), abi.encode(_DEFAULT_DECIMALS)
        );
        vm.expectCall(tokenContract, abi.encodeCall(token.decimals, ()));
    }

    function _deriveExpectedContractAddress(
        address creator,
        uint8 nonce
    ) private pure returns (address) {
        // Taken from https://ethereum.stackexchange.com/questions/760/how-is-the-address-of-an-ethereum-contract-computed
        // We must use encodePacked rather than encode so that the parameters are not padded with extra zeros.
        return address(
            uint160(
                uint256(
                    keccak256(abi.encodePacked(bytes1(0xd6), bytes1(0x94), creator, bytes1(nonce)))
                )
            )
        );
    }

    function _formatERC20BridgeErrorMessage(string memory errorMessage)
        private
        pure
        returns (bytes memory)
    {
        return bytes(string.concat("ERC20Bridge: ", errorMessage));
    }
}
