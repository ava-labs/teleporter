// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"crypto/ecdsa"
	"math"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	erc20tokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/ERC20TokenHome"
	nativetokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/NativeTokenHome"
	tokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/TokenHome"
	erc20tokenremote "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenRemote/ERC20TokenRemote"
	nativetokenremote "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenRemote/NativeTokenRemote"
	tokenremote "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenRemote/TokenRemote"
	wrappednativetoken "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/WrappedNativeToken"
	exampleerc20 "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/mocks/ExampleERC20Decimals"
	mockERC20SACR "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/mocks/MockERC20SendAndCallReceiver"
	mockNSACR "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/mocks/MockNativeSendAndCallReceiver"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

// Deployer keys set in the genesis file in order to determine the deployed address in advance.
// The deployed address is set as an admin for the Native Minter precompile.

var nativeTokenRemoteDeployerKeys = []string{
	// Deployer address: 			   0x1337cfd2dCff6270615B90938aCB1efE79801704
	// NativeTokenRemote address: 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B
	"aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985",

	// Deployer address: 			   0xFcec6c0674037f99fa473de09609B4b6D8158863
	// NativeTokenRemote address: 0x962c62B01529ecc0561D85d3fe395921ddC3665B
	"81e5e98c89023dabbe43e1081314eaae174330aae6b44c9d1371b6c0bb7ae74a",

	// Deployer address:			   0x2e1533d976A675bCD6306deC3B05e9f73e6722Fb
	// NativeTokenRemote address: 0x1549B96D9D97F435CA9b25000FEDE3A7e54C0bb9
	"5ded9cacaca7b88d6a3dc24641cfe41ef00186f98e7fa65135eac50fd5977f7a",

	// Deployer address:			   0xA638b0a597dc0520e2f20E83cFbeBBCd45a79990
	// NativeTokenRemote address: 0x190110D1228EB2cDd36559b2215A572Dc8592C3d
	"a6c530cb407778d10e1f70be6624aa57d0c724f6f9cb585e9744052d7f48ba19",

	// Deployer address:			   0x787C079cB0d5A7AA1Cae95d991F76Dce771A432D
	// NativeTokenRemote address: 0xf9EF017A764F265A1fD0975bfc200725E41d860E
	"e95fa6fd1d2a6b02890b75062bed583ce6256c5b473b3323b93ac4cbf20dbe7a",

	// Deployer address:			   0x741D536f5B07bcD43727CD8435389CA36aE5A4Ae
	// NativeTokenRemote address: 0x4f3663be6d22B0F19F8617f1A9E9485aB0144Bff
	"8a92f3f468ce5b0d99f9aaa55695f93e03dbbb6d5e3faba80f92a7876be740d6",
	// Deployer address:			   0xd466f12795BA59d0fef389c21fA63c287956fb18
	// NativeTokenRemote address: 0x463a6bE7a5098A5f06435c6c468adD338F15B93A
	"ebb7f0cf71e0b6fd880326e5f5061b8456b0aef81901566cbe578b5024852ec9",
}

var (
	nativeTokenRemoteDeployerKeyIndex   = 0
	ExpectedExampleERC20DeployerBalance = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e10))
)

const NativeTokenDecimals = 18

func DeployERC20TokenHome(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenAddress common.Address,
	tokenHomeDecimals uint8,
) (common.Address, *erc20tokenhome.ERC20TokenHome) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20TokenHome, err := erc20tokenhome.DeployERC20TokenHome(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		tokenAddress,
		tokenHomeDecimals,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, erc20TokenHome
}

func DeployERC20TokenRemote(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenHomeBlockchainID ids.ID,
	tokenHomeAddress common.Address,
	tokenHomeDecimals uint8,
	tokenName string,
	tokenSymbol string,
	tokenDecimals uint8,
) (common.Address, *erc20tokenremote.ERC20TokenRemote) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20TokenRemote, err := erc20tokenremote.DeployERC20TokenRemote(
		opts,
		subnet.RPCClient,
		erc20tokenremote.TokenRemoteSettings{
			TeleporterRegistryAddress: subnet.TeleporterRegistryAddress,
			TeleporterManager:         teleporterManager,
			TokenHomeBlockchainID:     tokenHomeBlockchainID,
			TokenHomeAddress:          tokenHomeAddress,
			TokenHomeDecimals:         tokenHomeDecimals,
		},
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, erc20TokenRemote
}

func DeployNativeTokenRemote(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	symbol string,
	teleporterManager common.Address,
	tokenHomeBlockchainID ids.ID,
	tokenHomeAddress common.Address,
	tokenHomeDecimals uint8,
	initialReserveImbalance *big.Int,
	multiplyOnRemote bool,
	burnedFeesReportingRewardPercentage *big.Int,
) (common.Address, *nativetokenremote.NativeTokenRemote) {
	// The NativeTokenRemote needs a unique deployer key, whose nonce 0 is used to deploy the contract.
	// The resulting contract address has been added to the genesis file as an admin for the Native Minter precompile.
	Expect(nativeTokenRemoteDeployerKeyIndex).Should(BeNumerically("<", len(nativeTokenRemoteDeployerKeys)))
	deployerKeyStr := nativeTokenRemoteDeployerKeys[nativeTokenRemoteDeployerKeyIndex]
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(
		deployerPK,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())

	address, tx, nativeTokenRemote, err := nativetokenremote.DeployNativeTokenRemote(
		opts,
		subnet.RPCClient,
		nativetokenremote.TokenRemoteSettings{
			TeleporterRegistryAddress: subnet.TeleporterRegistryAddress,
			TeleporterManager:         teleporterManager,
			TokenHomeBlockchainID:     tokenHomeBlockchainID,
			TokenHomeAddress:          tokenHomeAddress,
			TokenHomeDecimals:         tokenHomeDecimals,
		},
		symbol,
		initialReserveImbalance,
		burnedFeesReportingRewardPercentage,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Increment to the next deployer key so that the next contract deployment succeeds
	nativeTokenRemoteDeployerKeyIndex++

	return address, nativeTokenRemote
}

func DeployNativeTokenHome(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenAddress common.Address,
) (common.Address, *nativetokenhome.NativeTokenHome) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, nativeTokenHome, err := nativetokenhome.DeployNativeTokenHome(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		tokenAddress,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, nativeTokenHome
}

func DeployWrappedNativeToken(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	tokenSymbol string,
) (common.Address, *wrappednativetoken.WrappedNativeToken) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	// Deploy mock WAVAX contract
	address, tx, token, err := wrappednativetoken.DeployWrappedNativeToken(opts, subnet.RPCClient, tokenSymbol)
	Expect(err).Should(BeNil())
	log.Info("Deployed WrappedNativeToken contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	// Wait for the transaction to be mined
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, token
}

func DeployMockNativeSendAndCallReceiver(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *mockNSACR.MockNativeSendAndCallReceiver) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	// Deploy MockNativeSendAndCallReceiver contract
	address, tx, contract, err := mockNSACR.DeployMockNativeSendAndCallReceiver(opts, subnet.RPCClient)
	Expect(err).Should(BeNil())
	log.Info("Deployed MockNativeSendAndCallReceiver contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	// Wait for the transaction to be mined
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, contract
}

func DeployMockERC20SendAndCallReceiver(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *mockERC20SACR.MockERC20SendAndCallReceiver) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	// Deploy MockERC20SendAndCallReceiver contract
	address, tx, contract, err := mockERC20SACR.DeployMockERC20SendAndCallReceiver(opts, subnet.RPCClient)
	Expect(err).Should(BeNil())
	log.Info("Deployed MockERC20SendAndCallReceiver contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	// Wait for the transaction to be mined
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, contract
}

func DeployExampleERC20(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	tokenDecimals uint8,
) (common.Address, *exampleerc20.ExampleERC20Decimals) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	// Deploy Mock ERC20 contract
	address, tx, token, err := exampleerc20.DeployExampleERC20Decimals(opts, subnet.RPCClient, tokenDecimals)
	Expect(err).Should(BeNil())
	log.Info("Deployed Mock ERC20 contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	// Wait for the transaction to be mined
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Check that the deployer has the expected initial balance
	senderAddress := crypto.PubkeyToAddress(senderKey.PublicKey)
	balance, err := token.BalanceOf(&bind.CallOpts{}, senderAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(ExpectedExampleERC20DeployerBalance))

	return address, token
}

func RegisterERC20TokenRemoteOnHome(
	ctx context.Context,
	network interfaces.Network,
	homeSubnet interfaces.SubnetTestInfo,
	homeAddress common.Address,
	remoteSubnet interfaces.SubnetTestInfo,
	remoteAddress common.Address,
) *big.Int {
	return RegisterTokenRemoteOnHome(
		ctx,
		network,
		homeSubnet,
		homeAddress,
		remoteSubnet,
		remoteAddress,
		big.NewInt(0),
		big.NewInt(1),
		false,
	)
}

func RegisterTokenRemoteOnHome(
	ctx context.Context,
	network interfaces.Network,
	homeSubnet interfaces.SubnetTestInfo,
	homeAddress common.Address,
	remoteSubnet interfaces.SubnetTestInfo,
	remoteAddress common.Address,
	expectedInitialReserveBalance *big.Int,
	expectedTokenMultiplier *big.Int,
	expectedmultiplyOnRemote bool,
) *big.Int {
	// Call the remote to send a register message to the home
	tokenRemote, err := tokenremote.NewTokenRemote(
		remoteAddress,
		remoteSubnet.RPCClient,
	)
	Expect(err).Should(BeNil())
	_, fundedKey := network.GetFundedAccountInfo()

	// Deploy a new ERC20 token for testing registering with fees
	feeTokenAddress, feeToken := DeployExampleERC20(
		ctx,
		fundedKey,
		remoteSubnet,
		18,
	)

	// Approve the ERC20TokenHome to spend the tokens
	feeAmount := big.NewInt(1e18)
	ERC20Approve(
		ctx,
		feeToken,
		remoteAddress,
		feeAmount,
		remoteSubnet,
		fundedKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, remoteSubnet.EVMChainID)
	Expect(err).Should(BeNil())

	sendRegisterTx, err := tokenRemote.RegisterWithHome(
		opts,
		tokenremote.TeleporterFeeInfo{
			FeeTokenAddress: feeTokenAddress,
			Amount:          feeAmount,
		},
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, remoteSubnet, sendRegisterTx.Hash())

	// Relay the register message to the home
	receipt = network.RelayMessage(ctx, receipt, remoteSubnet, homeSubnet, true)

	// Check that the remote registered event was emitted
	tokenHome, err := tokenhome.NewTokenHome(homeAddress, homeSubnet.RPCClient)
	Expect(err).Should(BeNil())
	registerEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, tokenHome.ParseRemoteRegistered)
	Expect(err).Should(BeNil())
	Expect(registerEvent.RemoteBlockchainID[:]).Should(Equal(remoteSubnet.BlockchainID[:]))
	Expect(registerEvent.RemoteBridgeAddress).Should(Equal(remoteAddress))

	// Based on the initial reserve balance of the TokenRemote instance,
	// calculate the collateral amount of home tokens needed to collateralize the remote.
	collateralNeeded := calculateCollateralNeeded(
		expectedInitialReserveBalance,
		expectedTokenMultiplier,
		expectedmultiplyOnRemote,
	)
	teleporterUtils.ExpectBigEqual(registerEvent.InitialCollateralNeeded, collateralNeeded)

	return collateralNeeded
}

// AddCollateralToERC20TokenHome adds collateral to the ERC20TokenHome contract
// and verifies the collateral was added successfully. Any excess amount
// is returned to the caller.
func AddCollateralToERC20TokenHome(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenHome *erc20tokenhome.ERC20TokenHome,
	erc20TokenHomeAddress common.Address,
	exampleERC20 *exampleerc20.ExampleERC20Decimals,
	remoteBlockchainID ids.ID,
	remoteAddress common.Address,
	collateralAmount *big.Int,
	senderKey *ecdsa.PrivateKey,
) {
	// Approve the ERC20TokenHome to spend the collateral
	ERC20Approve(
		ctx,
		exampleERC20,
		erc20TokenHomeAddress,
		collateralAmount,
		subnet,
		senderKey,
	)

	// Add collateral to the ERC20TokenHome
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenHome.AddCollateral(
		opts,
		remoteBlockchainID,
		remoteAddress,
		collateralAmount,
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHome.ParseCollateralAdded)
	Expect(err).Should(BeNil())
	Expect(event.RemoteBlockchainID[:]).Should(Equal(remoteBlockchainID[:]))
	Expect(event.RemoteBridgeAddress).Should(Equal(remoteAddress))

	remoteSettings, err := erc20TokenHome.RegisteredRemotes(
		&bind.CallOpts{},
		remoteBlockchainID,
		remoteAddress)
	Expect(err).Should(BeNil())
	if collateralAmount.Cmp(remoteSettings.CollateralNeeded) > 0 {
		collateralAmount.Sub(collateralAmount, remoteSettings.CollateralNeeded)
	}
	teleporterUtils.ExpectBigEqual(event.Amount, collateralAmount)
	teleporterUtils.ExpectBigEqual(event.Remaining, big.NewInt(0))
}

// AddCollateralToNativeTokenHome adds collateral to the NativeTokenHome contract
// and verifies the collateral was added successfully. Any excess amount
// is returned to the caller.
func AddCollateralToNativeTokenHome(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenHome *nativetokenhome.NativeTokenHome,
	nativeTokenHomeAddress common.Address,
	remoteBlockchainID ids.ID,
	remoteAddress common.Address,
	collateralAmount *big.Int,
	senderKey *ecdsa.PrivateKey,
) {
	// Add collateral to the ERC20TokenHome
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = collateralAmount

	tx, err := nativeTokenHome.AddCollateral(
		opts,
		remoteBlockchainID,
		remoteAddress,
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenHome.ParseCollateralAdded)
	Expect(err).Should(BeNil())
	Expect(event.RemoteBlockchainID[:]).Should(Equal(remoteBlockchainID[:]))
	Expect(event.RemoteBridgeAddress).Should(Equal(remoteAddress))
	remoteSettings, err := nativeTokenHome.RegisteredRemotes(
		&bind.CallOpts{},
		remoteBlockchainID,
		remoteAddress)
	Expect(err).Should(BeNil())
	if collateralAmount.Cmp(remoteSettings.CollateralNeeded) > 0 {
		collateralAmount.Sub(collateralAmount, remoteSettings.CollateralNeeded)
	}
	teleporterUtils.ExpectBigEqual(event.Amount, collateralAmount)
	teleporterUtils.ExpectBigEqual(event.Remaining, big.NewInt(0))
}

func SendERC20TokenHome(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenHome *erc20tokenhome.ERC20TokenHome,
	erc20TokenHomeAddress common.Address,
	token *exampleerc20.ExampleERC20Decimals,
	input erc20tokenhome.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	// Approve the ERC20TokenHome to spend the tokens
	ERC20Approve(
		ctx,
		token,
		erc20TokenHomeAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		subnet,
		senderKey,
	)

	// Send the tokens and verify expected events
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenHome.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHome.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))

	// Compute the scaled amount
	scaledAmount := GetScaledAmountFromERC20TokenHome(
		erc20TokenHome,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendNativeTokenHome(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenHome *nativetokenhome.NativeTokenHome,
	nativeTokenHomeAddress common.Address,
	wrappedToken *wrappednativetoken.WrappedNativeToken,
	input nativetokenhome.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	DepositAndApproveWrappedTokenForFees(
		ctx,
		subnet,
		wrappedToken,
		input.PrimaryFee,
		nativeTokenHomeAddress,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenHome.Send(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenHome.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))

	// Compute the scaled amount
	scaledAmount := GetScaledAmountFromNativeTokenHome(
		nativeTokenHome,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendNativeTokenRemote(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenRemote *nativetokenremote.NativeTokenRemote,
	nativeTokenRemoteAddress common.Address,
	input nativetokenremote.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	DepositAndApproveWrappedTokenForFees(
		ctx,
		subnet,
		nativeTokenRemote,
		input.PrimaryFee,
		nativeTokenRemoteAddress,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenRemote.Send(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenRemote.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	return receipt, event.Amount
}

func SendERC20TokenRemote(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenRemote *erc20tokenremote.ERC20TokenRemote,
	erc20TokenRemoteAddress common.Address,
	input erc20tokenremote.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenRemote.Approve(
		opts,
		erc20TokenRemoteAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Bridge the tokens back to subnet A
	tx, err = erc20TokenRemote.Send(
		opts,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenRemote.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	return receipt, event.Amount
}

func SendAndCallERC20TokenHome(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenHome *erc20tokenhome.ERC20TokenHome,
	erc20TokenHomeAddress common.Address,
	exampleToken *exampleerc20.ExampleERC20Decimals,
	input erc20tokenhome.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	// Approve the ERC20TokenHome to spend the tokens
	ERC20Approve(
		ctx,
		exampleToken,
		erc20TokenHomeAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		subnet,
		senderKey,
	)

	// Send the tokens and verify expected events
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenHome.SendAndCall(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHome.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))

	// Computer the scaled amount
	scaledAmount := GetScaledAmountFromERC20TokenHome(
		erc20TokenHome,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendAndCallNativeTokenHome(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenHome *nativetokenhome.NativeTokenHome,
	input nativetokenhome.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenHome.SendAndCall(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenHome.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))

	// Compute the scaled amount
	remoteSettings, err := nativeTokenHome.RegisteredRemotes(
		&bind.CallOpts{},
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress)
	Expect(err).Should(BeNil())

	scaledAmount := ApplyTokenScaling(
		remoteSettings.TokenMultiplier,
		remoteSettings.MultiplyOnRemote,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendAndCallNativeTokenRemote(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenRemote *nativetokenremote.NativeTokenRemote,
	input nativetokenremote.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
	tokenMultiplier *big.Int,
	multiplyOnRemote bool,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenRemote.SendAndCall(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenRemote.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))
	teleporterUtils.ExpectBigEqual(event.Amount, bridgedAmount)

	return receipt, event.Amount
}

func SendAndCallERC20TokenRemote(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenRemote *erc20tokenremote.ERC20TokenRemote,
	erc20TokenRemoteAddress common.Address,
	input erc20tokenremote.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenRemote.Approve(
		opts,
		erc20TokenRemoteAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Bridge the tokens back to subnet A
	tx, err = erc20TokenRemote.SendAndCall(
		opts,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenRemote.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	return receipt, event.Amount
}

// Send a native token from fromBridge to toBridge via multi-hop through the C-Chain
// Requires that both fromBridge and toBridge are fully collateralized
// Requires that both fromBridge and toBridge have the same tokenMultiplier and multiplyOnRemote
// with respect to the original asset on the C-Chain
func SendNativeMultiHopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	sendingKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *nativetokenremote.NativeTokenRemote,
	fromBridgeAddress common.Address,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *nativetokenremote.NativeTokenRemote,
	toBridgeAddress common.Address,
	cChainInfo interfaces.SubnetTestInfo,
	amount *big.Int,
	secondaryFeeAmount *big.Int,
) {
	input := nativetokenremote.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   fromBridgeAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             secondaryFeeAmount,
		RequiredGasLimit:         DefaultNativeTokenRequiredGas,
		MultiHopFallback:         recipientAddress,
	}

	// Send tokens through a multi-hop transfer
	originReceipt, amount := SendNativeTokenRemote(
		ctx,
		fromSubnet,
		fromBridge,
		fromBridgeAddress,
		input,
		amount,
		sendingKey,
	)

	// Relay the first message back to the home chain, in this case C-Chain,
	// which then performs the multi-hop transfer to the destination TokenRemote instance.
	intermediateReceipt := network.RelayMessage(
		ctx,
		originReceipt,
		fromSubnet,
		cChainInfo,
		true,
	)

	initialBalance, err := toSubnet.RPCClient.BalanceAt(ctx, recipientAddress, nil)
	Expect(err).Should(BeNil())

	// When we relay the above message to the home chain, a multi-hop transfer
	// is performed to the destination TokenRemote instance. Parse for the send tokens event
	// and relay to the destination TokenRemote instance.
	network.RelayMessage(
		ctx,
		intermediateReceipt,
		cChainInfo,
		toSubnet,
		true,
	)

	bridgedAmount := big.NewInt(0).Sub(amount, input.SecondaryFee)
	teleporterUtils.CheckBalance(
		ctx,
		recipientAddress,
		big.NewInt(0).Add(initialBalance, bridgedAmount),
		toSubnet.RPCClient,
	)
}

func SendERC20TokenMultiHopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	fundedKey *ecdsa.PrivateKey,
	sendingKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *erc20tokenremote.ERC20TokenRemote,
	fromBridgeAddress common.Address,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *erc20tokenremote.ERC20TokenRemote,
	toBridgeAddress common.Address,
	cChainInfo interfaces.SubnetTestInfo,
	amount *big.Int,
	secondaryFeeAmount *big.Int,
) {
	// Send tokens to the sender address to have gas for submitting the send tokens transaction
	teleporterUtils.SendNativeTransfer(
		ctx,
		fromSubnet,
		fundedKey,
		crypto.PubkeyToAddress(sendingKey.PublicKey),
		big.NewInt(1e18),
	)
	input := erc20tokenremote.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   common.Address{},
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             secondaryFeeAmount,
		RequiredGasLimit:         DefaultERC20RequiredGas,
		MultiHopFallback:         recipientAddress,
	}

	// Send tokens through a multi-hop transfer
	originReceipt, amount := SendERC20TokenRemote(
		ctx,
		fromSubnet,
		fromBridge,
		fromBridgeAddress,
		input,
		amount,
		sendingKey,
	)

	// Relay the first message back to the home chain, in this case C-Chain,
	// which then performs the multi-hop transfer to the destination TokenRemote instance.
	intermediateReceipt := network.RelayMessage(
		ctx,
		originReceipt,
		fromSubnet,
		cChainInfo,
		true,
	)
	_, err := teleporterUtils.GetEventFromLogs(
		intermediateReceipt.Logs,
		cChainInfo.TeleporterMessenger.ParseMessageExecuted,
	)
	if err != nil {
		teleporterUtils.TraceTransactionAndExit(ctx, cChainInfo, intermediateReceipt.TxHash)
	}

	initialBalance, err := toBridge.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())

	// When we relay the above message to the home chain, a multi-hop transfer
	// is performed to the destination TokenRemote instance. Parse for the send tokens event
	// and relay to the destination TokenRemote instance.
	remoteReceipt := network.RelayMessage(
		ctx,
		intermediateReceipt,
		cChainInfo,
		toSubnet,
		true,
	)
	_, err = teleporterUtils.GetEventFromLogs(remoteReceipt.Logs, toSubnet.TeleporterMessenger.ParseMessageExecuted)
	if err != nil {
		teleporterUtils.TraceTransactionAndExit(ctx, toSubnet, remoteReceipt.TxHash)
	}

	bridgedAmount := big.NewInt(0).Sub(amount, input.SecondaryFee)
	CheckERC20TokenRemoteWithdrawal(
		ctx,
		toBridge,
		remoteReceipt,
		recipientAddress,
		bridgedAmount,
	)

	balance, err := toBridge.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, big.NewInt(0).Add(initialBalance, bridgedAmount))
}

func CheckERC20TokenHomeWithdrawal(
	ctx context.Context,
	erc20TokenHomeAddress common.Address,
	exampleERC20 *exampleerc20.ExampleERC20Decimals,
	receipt *types.Receipt,
	expectedRecipientAddress common.Address,
	expectedAmount *big.Int,
) {
	homeTransferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, exampleERC20.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(homeTransferEvent.From).Should(Equal(erc20TokenHomeAddress))
	Expect(homeTransferEvent.To).Should(Equal(expectedRecipientAddress))
	teleporterUtils.ExpectBigEqual(homeTransferEvent.Value, expectedAmount)
}

func CheckERC20TokenRemoteWithdrawal(
	ctx context.Context,
	erc20TokenRemote *erc20tokenremote.ERC20TokenRemote,
	receipt *types.Receipt,
	expectedRecipientAddress common.Address,
	expectedAmount *big.Int,
) {
	transferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenRemote.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(transferEvent.From).Should(Equal(common.Address{}))
	Expect(transferEvent.To).Should(Equal(expectedRecipientAddress))
	teleporterUtils.ExpectBigEqual(transferEvent.Value, expectedAmount)
}

func CheckNativeTokenHomeWithdrawal(
	ctx context.Context,
	nativeTokenHomeAddress common.Address,
	wrappedNativeToken *wrappednativetoken.WrappedNativeToken,
	receipt *types.Receipt,
	expectedAmount *big.Int,
) {
	withdrawalEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, wrappedNativeToken.ParseWithdrawal)
	Expect(err).Should(BeNil())
	Expect(withdrawalEvent.Sender).Should(Equal(nativeTokenHomeAddress))
	teleporterUtils.ExpectBigEqual(withdrawalEvent.Amount, expectedAmount)
}

func GetTokenMultiplier(
	decimalsShift uint8,
) *big.Int {
	return big.NewInt(int64(math.Pow10(int(decimalsShift))))
}

type WrappedToken interface {
	Deposit(opts *bind.TransactOpts) (*types.Transaction, error)
	Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error)
}

func DepositAndApproveWrappedTokenForFees(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	wrappedToken WrappedToken,
	amount *big.Int,
	spender common.Address,
	senderKey *ecdsa.PrivateKey,
) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		return
	}

	// Deposit the native tokens for paying the fee
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount
	tx, err := wrappedToken.Deposit(opts)
	Expect(err).Should(BeNil())

	_ = teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	opts, err = bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = wrappedToken.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())

	_ = teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func ERC20Approve(
	ctx context.Context,
	token *exampleerc20.ExampleERC20Decimals,
	spender common.Address,
	amount *big.Int,
	subnet interfaces.SubnetTestInfo,
	senderKey *ecdsa.PrivateKey,
) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := token.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())
	log.Info("Approved ERC20", "spender", spender.Hex(), "txHash", tx.Hash().Hex())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}
