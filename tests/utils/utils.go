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
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Destination"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	nativetokendestination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenDestination"
	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenSource"
	teleportertokendestination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TeleporterTokenDestination"
	teleportertokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TeleporterTokenSource"
	examplewavax "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/mocks/ExampleWAVAX"
	mockERC20SACR "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/mocks/MockERC20SendAndCallReceiver"
	mockNSACR "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/mocks/MockNativeSendAndCallReceiver"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

// Deployer keys set in the genesis file in order to determine the deployed address in advance.
// The deployed address is set as an admin for the Native Minter precompile.

var nativeTokenDestinationDeployerKeys = []string{
	// Deployer address: 			   0x1337cfd2dCff6270615B90938aCB1efE79801704
	// NativeTokenDestination address: 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B
	"aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985",

	// Deployer address: 			   0xFcec6c0674037f99fa473de09609B4b6D8158863
	// NativeTokenDestination address: 0x962c62B01529ecc0561D85d3fe395921ddC3665B
	"81e5e98c89023dabbe43e1081314eaae174330aae6b44c9d1371b6c0bb7ae74a",

	// Deployer address:			   0x2e1533d976A675bCD6306deC3B05e9f73e6722Fb
	// NativeTokenDestination address: 0x1549B96D9D97F435CA9b25000FEDE3A7e54C0bb9
	"5ded9cacaca7b88d6a3dc24641cfe41ef00186f98e7fa65135eac50fd5977f7a",

	// Deployer address:			   0xA638b0a597dc0520e2f20E83cFbeBBCd45a79990
	// NativeTokenDestination address: 0x190110D1228EB2cDd36559b2215A572Dc8592C3d
	"a6c530cb407778d10e1f70be6624aa57d0c724f6f9cb585e9744052d7f48ba19",

	// Deployer address:			   0x787C079cB0d5A7AA1Cae95d991F76Dce771A432D
	// NativeTokenDestination address: 0xf9EF017A764F265A1fD0975bfc200725E41d860E
	"e95fa6fd1d2a6b02890b75062bed583ce6256c5b473b3323b93ac4cbf20dbe7a",

	// Deployer address:			   0x741D536f5B07bcD43727CD8435389CA36aE5A4Ae
	// NativeTokenDestination address: 0x4f3663be6d22B0F19F8617f1A9E9485aB0144Bff
	"8a92f3f468ce5b0d99f9aaa55695f93e03dbbb6d5e3faba80f92a7876be740d6",
	// Deployer address:			   0xd466f12795BA59d0fef389c21fA63c287956fb18
	// NativeTokenDestination address: 0x463a6bE7a5098A5f06435c6c468adD338F15B93A
	"ebb7f0cf71e0b6fd880326e5f5061b8456b0aef81901566cbe578b5024852ec9",
}
var nativeTokenDestinationDeployerKeyIndex = 0

func DeployERC20Source(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenSourceAddress common.Address,
) (common.Address, *erc20source.ERC20Source) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20Source, err := erc20source.DeployERC20Source(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		tokenSourceAddress,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, erc20Source
}

func DeployERC20Destination(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	sourceBlockchainID ids.ID,
	tokenSourceAddress common.Address,
	tokenName string,
	tokenSymbol string,
	tokenDecimals uint8,
) (common.Address, *erc20destination.ERC20Destination) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20Destination, err := erc20destination.DeployERC20Destination(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		sourceBlockchainID,
		tokenSourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, erc20Destination
}

func DeployNativeTokenDestination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	symbol string,
	teleporterManager common.Address,
	sourceBlockchainID ids.ID,
	tokenSourceAddress common.Address,
	initialReserveImbalance *big.Int,
	decimalsShift uint8,
	multiplyOnDestination bool,
	burnedFeesReportingRewardPercentage *big.Int,
) (common.Address, *nativetokendestination.NativeTokenDestination) {
	// The Native Token Destination needs a unique deployer key, whose nonce 0 is used to deploy the contract.
	// The resulting contract address has been added to the genesis file as an admin for the Native Minter precompile.
	Expect(nativeTokenDestinationDeployerKeyIndex).Should(BeNumerically("<", len(nativeTokenDestinationDeployerKeys)))
	deployerKeyStr := nativeTokenDestinationDeployerKeys[nativeTokenDestinationDeployerKeyIndex]
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(
		deployerPK,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())

	address, tx, nativeTokenDestination, err := nativetokendestination.DeployNativeTokenDestination(
		opts,
		subnet.RPCClient,
		nativetokendestination.NativeTokenDestinationSettings{
			NativeAssetSymbol:                   symbol,
			TeleporterRegistryAddress:           subnet.TeleporterRegistryAddress,
			TeleporterManager:                   teleporterManager,
			SourceBlockchainID:                  sourceBlockchainID,
			TokenSourceAddress:                  tokenSourceAddress,
			InitialReserveImbalance:             initialReserveImbalance,
			DecimalsShift:                       decimalsShift,
			MultiplyOnDestination:               multiplyOnDestination,
			BurnedFeesReportingRewardPercentage: burnedFeesReportingRewardPercentage,
		},
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Increment to the next deployer key so that the next contract deployment succeeds
	nativeTokenDestinationDeployerKeyIndex++

	return address, nativeTokenDestination
}

func DeployNativeTokenSource(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenAddress common.Address,
) (common.Address, *nativetokensource.NativeTokenSource) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, nativeTokenSource, err := nativetokensource.DeployNativeTokenSource(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		tokenAddress,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, nativeTokenSource
}

func DeployExampleWAVAX(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *examplewavax.ExampleWAVAX) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	// Deploy mock WAVAX contract
	address, tx, token, err := examplewavax.DeployExampleWAVAX(opts, subnet.RPCClient)
	Expect(err).Should(BeNil())
	log.Info("Deployed ExampleWAVAX contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

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

func RegisterERC20DestinationOnSource(
	ctx context.Context,
	network interfaces.Network,
	sourceSubnet interfaces.SubnetTestInfo,
	sourceBridgeAddress common.Address,
	destinationSubnet interfaces.SubnetTestInfo,
	destinationBridgeAddress common.Address,
) *big.Int {
	return RegisterTokenDestinationOnSource(
		ctx,
		network,
		sourceSubnet,
		sourceBridgeAddress,
		destinationSubnet,
		destinationBridgeAddress,
		big.NewInt(0),
		big.NewInt(1),
		false,
	)
}

func RegisterTokenDestinationOnSource(
	ctx context.Context,
	network interfaces.Network,
	sourceSubnet interfaces.SubnetTestInfo,
	sourceBridgeAddress common.Address,
	destinationSubnet interfaces.SubnetTestInfo,
	destinationBridgeAddress common.Address,
	expectedInitialReserveBalance *big.Int,
	expectedTokenMultiplier *big.Int,
	expectedMultiplyOnDestination bool,
) *big.Int {
	// Call the destination to send a register message to the source
	tokenDestination, err := teleportertokendestination.NewTeleporterTokenDestination(
		destinationBridgeAddress,
		destinationSubnet.RPCClient,
	)
	Expect(err).Should(BeNil())
	_, fundedKey := network.GetFundedAccountInfo()
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, destinationSubnet.EVMChainID)
	Expect(err).Should(BeNil())
	sendRegisterTx, err := tokenDestination.RegisterWithSource(
		opts,
		teleportertokendestination.TeleporterFeeInfo{FeeTokenAddress: common.Address{}, Amount: big.NewInt(0)},
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, destinationSubnet, sendRegisterTx.Hash())

	// Relay the register message to the source
	receipt = network.RelayMessage(ctx, receipt, destinationSubnet, sourceSubnet, true)

	// Check that the destination registered event was emitted
	tokenSource, err := teleportertokensource.NewTeleporterTokenSource(sourceBridgeAddress, sourceSubnet.RPCClient)
	Expect(err).Should(BeNil())
	registerEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, tokenSource.ParseDestinationRegistered)
	Expect(err).Should(BeNil())
	Expect(registerEvent.DestinationBlockchainID[:]).Should(Equal(destinationSubnet.BlockchainID[:]))
	Expect(registerEvent.DestinationBridgeAddress).Should(Equal(destinationBridgeAddress))

	// Based on the initial reserve balance of the destination bridge,
	// calculate the collateral amount of source tokens needed to collateralize the destination.
	collateralNeeded := calculateCollateralNeeded(
		expectedInitialReserveBalance,
		expectedTokenMultiplier,
		expectedMultiplyOnDestination,
	)
	teleporterUtils.ExpectBigEqual(registerEvent.InitialCollateralNeeded, collateralNeeded)
	teleporterUtils.ExpectBigEqual(registerEvent.TokenMultiplier, expectedTokenMultiplier)
	Expect(registerEvent.MultiplyOnDestination).Should(Equal(expectedMultiplyOnDestination))

	return collateralNeeded
}

// AddCollateralToERC20Source adds collateral to the ERC20Source contract
// and verifies the collateral was added successfully. Any excess amount
// is returned to the caller.
func AddCollateralToERC20Source(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Source *erc20source.ERC20Source,
	erc20SourceAddress common.Address,
	exampleERC20 *exampleerc20.ExampleERC20,
	destinationBlockchainID ids.ID,
	destinationBridgeAddress common.Address,
	collateralAmount *big.Int,
	senderKey *ecdsa.PrivateKey,
) {
	// Approve the ERC20Source to spend the collateral
	teleporterUtils.ERC20Approve(
		ctx,
		exampleERC20,
		erc20SourceAddress,
		collateralAmount,
		subnet,
		senderKey,
	)

	// Add collateral to the ERC20Source
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Source.AddCollateral(
		opts,
		destinationBlockchainID,
		destinationBridgeAddress,
		collateralAmount,
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Source.ParseCollateralAdded)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationBlockchainID[:]))
	Expect(event.DestinationBridgeAddress).Should(Equal(destinationBridgeAddress))

	destinationSettings, err := erc20Source.RegisteredDestinations(
		&bind.CallOpts{},
		destinationBlockchainID,
		destinationBridgeAddress)
	Expect(err).Should(BeNil())
	if collateralAmount.Cmp(destinationSettings.CollateralNeeded) > 0 {
		collateralAmount.Sub(collateralAmount, destinationSettings.CollateralNeeded)
	}
	teleporterUtils.ExpectBigEqual(event.Amount, collateralAmount)
	teleporterUtils.ExpectBigEqual(event.Remaining, big.NewInt(0))
}

// AddCollateralToNativeTokenSource adds collateral to the NativeTokenSource contract
// and verifies the collateral was added successfully. Any excess amount
// is returned to the caller.
func AddCollateralToNativeTokenSource(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenSource *nativetokensource.NativeTokenSource,
	nativeTokenSourceAddress common.Address,
	destinationBlockchainID ids.ID,
	destinationBridgeAddress common.Address,
	collateralAmount *big.Int,
	senderKey *ecdsa.PrivateKey,
) {
	// Add collateral to the ERC20Source
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = collateralAmount

	tx, err := nativeTokenSource.AddCollateral(
		opts,
		destinationBlockchainID,
		destinationBridgeAddress,
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenSource.ParseCollateralAdded)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destinationBlockchainID[:]))
	Expect(event.DestinationBridgeAddress).Should(Equal(destinationBridgeAddress))
	destinationSettings, err := nativeTokenSource.RegisteredDestinations(
		&bind.CallOpts{},
		destinationBlockchainID,
		destinationBridgeAddress)
	Expect(err).Should(BeNil())
	if collateralAmount.Cmp(destinationSettings.CollateralNeeded) > 0 {
		collateralAmount.Sub(collateralAmount, destinationSettings.CollateralNeeded)
	}
	teleporterUtils.ExpectBigEqual(event.Amount, collateralAmount)
	teleporterUtils.ExpectBigEqual(event.Remaining, big.NewInt(0))
}

func SendERC20Source(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Source *erc20source.ERC20Source,
	erc20SourceAddress common.Address,
	sourceToken *exampleerc20.ExampleERC20,
	input erc20source.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	// Approve the ERC20Source to spend the tokens
	teleporterUtils.ERC20Approve(
		ctx,
		sourceToken,
		erc20SourceAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		subnet,
		senderKey,
	)

	// Send the tokens and verify expected events
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Source.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Source.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	// Compute the scaled amount
	scaledAmount := GetScaledAmountFromERC20Source(
		erc20Source,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendNativeTokenSource(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenSource *nativetokensource.NativeTokenSource,
	nativeTokenSourceAddress common.Address,
	wrappedToken *examplewavax.ExampleWAVAX,
	input nativetokensource.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	DepositAndApproveWavaxForFees(
		ctx,
		subnet,
		wrappedToken,
		input.PrimaryFee,
		nativeTokenSourceAddress,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenSource.Send(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenSource.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	// Compute the scaled amount
	scaledAmount := GetScaledAmountFromNativeTokenSource(
		nativeTokenSource,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func DepositAndApproveWavaxForFees(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	wrappedToken *examplewavax.ExampleWAVAX,
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

	depositReceipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	depositEvent, err := teleporterUtils.GetEventFromLogs(depositReceipt.Logs, wrappedToken.ParseDeposit)
	Expect(err).Should(BeNil())
	Expect(depositEvent.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(depositEvent.Amount, amount)

	opts, err = bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = wrappedToken.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())

	approveReceipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	approvalEvent, err := teleporterUtils.GetEventFromLogs(approveReceipt.Logs, wrappedToken.ParseApproval)
	Expect(err).Should(BeNil())
	Expect(approvalEvent.Spender).Should(Equal(spender))
}

func DepositAndApproveNativeForFees(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	amount *big.Int,
	spender common.Address,
	senderKey *ecdsa.PrivateKey,
) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		return
	}

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount
	tx, err := nativeTokenDestination.Deposit(opts)
	Expect(err).Should(BeNil())

	depositReceipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	depositEvent, err := teleporterUtils.GetEventFromLogs(depositReceipt.Logs, nativeTokenDestination.ParseDeposit)
	Expect(err).Should(BeNil())
	Expect(depositEvent.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(depositEvent.Amount, amount)

	opts, err = bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err = nativeTokenDestination.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())

	approveReceipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	approvalEvent, err := teleporterUtils.GetEventFromLogs(approveReceipt.Logs, nativeTokenDestination.ParseApproval)
	Expect(err).Should(BeNil())
	Expect(approvalEvent.Spender).Should(Equal(spender))
}

func SendNativeTokenDestination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	nativeTokenDestinationAddress common.Address,
	input nativetokendestination.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	DepositAndApproveNativeForFees(
		ctx,
		subnet,
		nativeTokenDestination,
		input.PrimaryFee,
		nativeTokenDestinationAddress,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenDestination.Send(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	bridgedAmount := new(big.Int).Sub(amount, input.PrimaryFee)

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenDestination.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, bridgedAmount)

	return receipt, event.Amount
}

func SendERC20Destination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Destination *erc20destination.ERC20Destination,
	erc20DestinationAddress common.Address,
	input erc20destination.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Destination.Approve(
		opts,
		erc20DestinationAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Bridge the tokens back to subnet A
	tx, err = erc20Destination.Send(
		opts,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Destination.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	return receipt, event.Amount
}

func SendAndCallERC20Source(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Source *erc20source.ERC20Source,
	erc20SourceAddress common.Address,
	sourceToken *exampleerc20.ExampleERC20,
	input erc20source.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	// Approve the ERC20Source to spend the tokens
	teleporterUtils.ERC20Approve(
		ctx,
		sourceToken,
		erc20SourceAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		subnet,
		senderKey,
	)

	// Send the tokens and verify expected events
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Source.SendAndCall(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Source.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	// Computer the scaled amount
	destinationSettings, err := erc20Source.RegisteredDestinations(
		&bind.CallOpts{},
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress)
	Expect(err).Should(BeNil())

	scaledAmount := ApplyTokenScaling(
		destinationSettings.TokenMultiplier,
		destinationSettings.MultiplyOnDestination,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendAndCallNativeTokenSource(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenSource *nativetokensource.NativeTokenSource,
	input nativetokensource.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenSource.SendAndCall(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenSource.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	// Computer the scaled amount
	destinationSettings, err := nativeTokenSource.RegisteredDestinations(
		&bind.CallOpts{},
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress)
	Expect(err).Should(BeNil())

	scaledAmount := ApplyTokenScaling(
		destinationSettings.TokenMultiplier,
		destinationSettings.MultiplyOnDestination,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendAndCallNativeTokenDestination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenDestination *nativetokendestination.NativeTokenDestination,
	input nativetokendestination.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
	tokenMultiplier *big.Int,
	multiplyOnDestination bool,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenDestination.SendAndCall(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenDestination.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))
	teleporterUtils.ExpectBigEqual(event.Amount, bridgedAmount)

	return receipt, event.Amount
}

func SendAndCallERC20Destination(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20Destination *erc20destination.ERC20Destination,
	erc20DestinationAddress common.Address,
	input erc20destination.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20Destination.Approve(
		opts,
		erc20DestinationAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Bridge the tokens back to subnet A
	tx, err = erc20Destination.SendAndCall(
		opts,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Destination.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	return receipt, event.Amount
}

// Send a native token from fromBridge to toBridge via multi-hop through the C-Chain
// Requires that both fromBridge and toBridge are fully collateralized
// Requires that both fromBridge and toBridge have the same tokenMultiplier and multiplyOnDestination
// with respect to the original asset on the C-Chain
func SendNativeMultihopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	sendingKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *nativetokendestination.NativeTokenDestination,
	fromBridgeAddress common.Address,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *nativetokendestination.NativeTokenDestination,
	toBridgeAddress common.Address,
	cChainInfo interfaces.SubnetTestInfo,
	bridgedAmount *big.Int,
) {
	input := nativetokendestination.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		FeeTokenAddress:          fromBridgeAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         DefaultNativeTokenRequiredGasLimit,
		FallbackRecipient:        recipientAddress,
	}

	// Find the amount sent by fromBridge. This is before any scaling/unscaling is applied.
	bridgedAmount = teleporterUtils.BigIntSub(bridgedAmount, input.PrimaryFee)

	// Send tokens through a multi-hop transfer
	originReceipt, bridgedAmount := SendNativeTokenDestination(
		ctx,
		fromSubnet,
		fromBridge,
		fromBridgeAddress,
		input,
		bridgedAmount,
		sendingKey,
	)

	// Relay the first message back to the home-chain, in this case C-Chain,
	// which then performs the multi-hop transfer to the destination chain
	intermediateReceipt := network.RelayMessage(
		ctx,
		originReceipt,
		fromSubnet,
		cChainInfo,
		true,
	)

	initialBalance, err := toSubnet.RPCClient.BalanceAt(ctx, recipientAddress, nil)
	Expect(err).Should(BeNil())

	// When we relay the above message to the home-chain, a multi-hop transfer
	// is performed to the destination chain. Parse for the send tokens event
	// and relay to final destination.
	network.RelayMessage(
		ctx,
		intermediateReceipt,
		cChainInfo,
		toSubnet,
		true,
	)

	teleporterUtils.CheckBalance(
		ctx,
		recipientAddress,
		big.NewInt(0).Add(initialBalance, bridgedAmount),
		toSubnet.RPCClient,
	)
}

func SendERC20MultihopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	fundedKey *ecdsa.PrivateKey,
	sendingKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *erc20destination.ERC20Destination,
	fromBridgeAddress common.Address,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *erc20destination.ERC20Destination,
	toBridgeAddress common.Address,
	cChainInfo interfaces.SubnetTestInfo,
	bridgedAmount *big.Int,
) {
	// Send tokens to the sender address to have gas for submitting the send tokens transaction
	teleporterUtils.SendNativeTransfer(
		ctx,
		fromSubnet,
		fundedKey,
		crypto.PubkeyToAddress(sendingKey.PublicKey),
		big.NewInt(1e18),
	)
	input := erc20destination.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		FeeTokenAddress:          common.Address{},
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         DefaultERC20RequiredGasLimit,
		FallbackRecipient:        recipientAddress,
	}

	// Send tokens through a multi-hop transfer
	originReceipt, bridgedAmount := SendERC20Destination(
		ctx,
		fromSubnet,
		fromBridge,
		fromBridgeAddress,
		input,
		bridgedAmount,
		sendingKey,
	)

	// Relay the first message back to the home-chain, in this case C-Chain,
	// which then performs the multi-hop transfer to the destination chain
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

	// When we relay the above message to the home-chain, a multi-hop transfer
	// is performed to the destination chain. Parse for the send tokens event
	// and relay to final destination.
	destinationReceipt := network.RelayMessage(
		ctx,
		intermediateReceipt,
		cChainInfo,
		toSubnet,
		true,
	)
	_, err = teleporterUtils.GetEventFromLogs(destinationReceipt.Logs, toSubnet.TeleporterMessenger.ParseMessageExecuted)
	if err != nil {
		teleporterUtils.TraceTransactionAndExit(ctx, toSubnet, destinationReceipt.TxHash)
	}

	CheckERC20DestinationWithdrawal(
		ctx,
		toBridge,
		destinationReceipt,
		recipientAddress,
		bridgedAmount,
	)

	balance, err := toBridge.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, big.NewInt(0).Add(initialBalance, bridgedAmount))
}

func CheckERC20SourceWithdrawal(
	ctx context.Context,
	erc20SourceAddress common.Address,
	sourceToken *exampleerc20.ExampleERC20,
	receipt *types.Receipt,
	expectedRecipientAddress common.Address,
	expectedAmount *big.Int,
) {
	sourceTransferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, sourceToken.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(sourceTransferEvent.From).Should(Equal(erc20SourceAddress))
	Expect(sourceTransferEvent.To).Should(Equal(expectedRecipientAddress))
	teleporterUtils.ExpectBigEqual(sourceTransferEvent.Value, expectedAmount)
}

func CheckERC20DestinationWithdrawal(
	ctx context.Context,
	erc20Destination *erc20destination.ERC20Destination,
	receipt *types.Receipt,
	expectedRecipientAddress common.Address,
	expectedAmount *big.Int,
) {
	transferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Destination.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(transferEvent.From).Should(Equal(common.Address{}))
	Expect(transferEvent.To).Should(Equal(expectedRecipientAddress))
	teleporterUtils.ExpectBigEqual(transferEvent.Value, expectedAmount)
}

func CheckNativeTokenSourceWithdrawal(
	ctx context.Context,
	nativeTokenSourceAddress common.Address,
	sourceToken *examplewavax.ExampleWAVAX,
	receipt *types.Receipt,
	expectedAmount *big.Int,
) {
	withdrawalEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, sourceToken.ParseWithdrawal)
	Expect(err).Should(BeNil())
	Expect(withdrawalEvent.Sender).Should(Equal(nativeTokenSourceAddress))
	teleporterUtils.ExpectBigEqual(withdrawalEvent.Amount, expectedAmount)
}

func GetTokenMultiplier(
	decimalsShift uint8,
) *big.Int {
	return big.NewInt(int64(math.Pow10(int(decimalsShift))))
}
