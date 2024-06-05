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
	proxyadmin "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ProxyAdmin"
	erc20tokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/ERC20TokenHub"
	nativetokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/NativeTokenHub"
	tokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/TokenHub"
	erc20tokenspoke "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenSpoke/ERC20TokenSpoke"
	nativetokenspoke "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenSpoke/NativeTokenSpoke"
	tokenspoke "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenSpoke/TokenSpoke"
	transparentupgradeableproxy "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TransparentUpgradeableProxy"
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

var nativeTokenSpokeDeployerKeys = []string{
	// Deployer address: 			   0x1337cfd2dCff6270615B90938aCB1efE79801704
	// NativeTokenSpoke address: 0xAcB633F5B00099c7ec187eB00156c5cd9D854b5B
	"aad7440febfc8f9d73a58c3cb1f1754779a566978f9ebffcd4f4698e9b043985",

	// Deployer address: 			   0xFcec6c0674037f99fa473de09609B4b6D8158863
	// NativeTokenSpoke address: 0x962c62B01529ecc0561D85d3fe395921ddC3665B
	"81e5e98c89023dabbe43e1081314eaae174330aae6b44c9d1371b6c0bb7ae74a",

	// Deployer address:			   0x2e1533d976A675bCD6306deC3B05e9f73e6722Fb
	// NativeTokenSpoke address: 0x1549B96D9D97F435CA9b25000FEDE3A7e54C0bb9
	"5ded9cacaca7b88d6a3dc24641cfe41ef00186f98e7fa65135eac50fd5977f7a",

	// Deployer address:			   0xA638b0a597dc0520e2f20E83cFbeBBCd45a79990
	// NativeTokenSpoke address: 0x190110D1228EB2cDd36559b2215A572Dc8592C3d
	"a6c530cb407778d10e1f70be6624aa57d0c724f6f9cb585e9744052d7f48ba19",

	// Deployer address:			   0x787C079cB0d5A7AA1Cae95d991F76Dce771A432D
	// NativeTokenSpoke address: 0xf9EF017A764F265A1fD0975bfc200725E41d860E
	"e95fa6fd1d2a6b02890b75062bed583ce6256c5b473b3323b93ac4cbf20dbe7a",

	// Deployer address:			   0x741D536f5B07bcD43727CD8435389CA36aE5A4Ae
	// NativeTokenSpoke address: 0x4f3663be6d22B0F19F8617f1A9E9485aB0144Bff
	"8a92f3f468ce5b0d99f9aaa55695f93e03dbbb6d5e3faba80f92a7876be740d6",
	// Deployer address:			   0xd466f12795BA59d0fef389c21fA63c287956fb18
	// NativeTokenSpoke address: 0x463a6bE7a5098A5f06435c6c468adD338F15B93A
	"ebb7f0cf71e0b6fd880326e5f5061b8456b0aef81901566cbe578b5024852ec9",
}

var (
	nativeTokenSpokeDeployerKeyIndex    = 0
	ExpectedExampleERC20DeployerBalance = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e10))
)

const NativeTokenDecimals = 18

func DeployERC20TokenHub(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenAddress common.Address,
	tokenHubDecimals uint8,
) (common.Address, *erc20tokenhub.ERC20TokenHub) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	implAddress, tx, erc20TokenHub, err := erc20tokenhub.DeployERC20TokenHub(
		opts,
		subnet.RPCClient,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	proxyAdminAddress, tx, _, err := proxyadmin.DeployProxyAdmin(
		opts,
		subnet.RPCClient,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	proxyAddress, tx, _, err := transparentupgradeableproxy.DeployTransparentUpgradeableProxy(
		opts,
		subnet.RPCClient,
		implAddress,
		proxyAdminAddress,
		[]byte{},
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	erc20TokenHub, err = erc20tokenhub.NewERC20TokenHub(proxyAddress, subnet.RPCClient)
	Expect(err).Should(BeNil())

	tx, err = erc20TokenHub.Initialize(
		opts,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		tokenAddress,
		tokenHubDecimals,
	)
	Expect(err).Should(BeNil())
	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return proxyAddress, erc20TokenHub
}

func DeployERC20TokenSpoke(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenHubBlockchainID ids.ID,
	tokenHubAddress common.Address,
	tokenHubDecimals uint8,
	tokenName string,
	tokenSymbol string,
	tokenDecimals uint8,
) (common.Address, *erc20tokenspoke.ERC20TokenSpoke) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, erc20TokenSpoke, err := erc20tokenspoke.DeployERC20TokenSpoke(
		opts,
		subnet.RPCClient,
		erc20tokenspoke.TokenSpokeSettings{
			TeleporterRegistryAddress: subnet.TeleporterRegistryAddress,
			TeleporterManager:         teleporterManager,
			TokenHubBlockchainID:      tokenHubBlockchainID,
			TokenHubAddress:           tokenHubAddress,
			TokenHubDecimals:          tokenHubDecimals,
		},
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return implAddress, erc20TokenSpoke
}

func DeployNativeTokenSpoke(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	symbol string,
	teleporterManager common.Address,
	tokenHubBlockchainID ids.ID,
	tokenHubAddress common.Address,
	tokenHubDecimals uint8,
	initialReserveImbalance *big.Int,
	multiplyOnSpoke bool,
	burnedFeesReportingRewardPercentage *big.Int,
) (common.Address, *nativetokenspoke.NativeTokenSpoke) {
	// The NativeTokenSpoke needs a unique deployer key, whose nonce 0 is used to deploy the contract.
	// The resulting contract address has been added to the genesis file as an admin for the Native Minter precompile.
	Expect(nativeTokenSpokeDeployerKeyIndex).Should(BeNumerically("<", len(nativeTokenSpokeDeployerKeys)))
	deployerKeyStr := nativeTokenSpokeDeployerKeys[nativeTokenSpokeDeployerKeyIndex]
	deployerPK, err := crypto.HexToECDSA(deployerKeyStr)
	Expect(err).Should(BeNil())

	opts, err := bind.NewKeyedTransactorWithChainID(
		deployerPK,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())

	address, tx, nativeTokenSpoke, err := nativetokenspoke.DeployNativeTokenSpoke(
		opts,
		subnet.RPCClient,
		nativetokenspoke.TokenSpokeSettings{
			TeleporterRegistryAddress: subnet.TeleporterRegistryAddress,
			TeleporterManager:         teleporterManager,
			TokenHubBlockchainID:      tokenHubBlockchainID,
			TokenHubAddress:           tokenHubAddress,
			TokenHubDecimals:          tokenHubDecimals,
		},
		symbol,
		initialReserveImbalance,
		burnedFeesReportingRewardPercentage,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Increment to the next deployer key so that the next contract deployment succeeds
	nativeTokenSpokeDeployerKeyIndex++

	return implAddress, nativeTokenSpoke
}

func DeployNativeTokenHub(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
	teleporterManager common.Address,
	tokenAddress common.Address,
) (common.Address, *nativetokenhub.NativeTokenHub) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey,
		subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	address, tx, nativeTokenHub, err := nativetokenhub.DeployNativeTokenHub(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		tokenAddress,
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return implAddress, nativeTokenHub
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

	return implAddress, token
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

	return implAddress, contract
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

	return implAddress, contract
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

	return implAddress, token
}

func RegisterERC20TokenSpokeOnHub(
	ctx context.Context,
	network interfaces.Network,
	hubSubnet interfaces.SubnetTestInfo,
	hubAddress common.Address,
	spokeSubnet interfaces.SubnetTestInfo,
	spokeAddress common.Address,
) *big.Int {
	return RegisterTokenSpokeOnHub(
		ctx,
		network,
		hubSubnet,
		hubAddress,
		spokeSubnet,
		spokeAddress,
		big.NewInt(0),
		big.NewInt(1),
		false,
	)
}

func RegisterTokenSpokeOnHub(
	ctx context.Context,
	network interfaces.Network,
	hubSubnet interfaces.SubnetTestInfo,
	hubAddress common.Address,
	spokeSubnet interfaces.SubnetTestInfo,
	spokeAddress common.Address,
	expectedInitialReserveBalance *big.Int,
	expectedTokenMultiplier *big.Int,
	expectedmultiplyOnSpoke bool,
) *big.Int {
	// Call the spoke to send a register message to the hub
	tokenSpoke, err := tokenspoke.NewTokenSpoke(
		spokeAddress,
		spokeSubnet.RPCClient,
	)
	Expect(err).Should(BeNil())
	_, fundedKey := network.GetFundedAccountInfo()

	// Deploy a new ERC20 token for testing registering with fees
	feeTokenAddress, feeToken := DeployExampleERC20(
		ctx,
		fundedKey,
		spokeSubnet,
		18,
	)

	// Approve the ERC20TokenHub to spend the tokens
	feeAmount := big.NewInt(1e18)
	ERC20Approve(
		ctx,
		feeToken,
		spokeAddress,
		feeAmount,
		spokeSubnet,
		fundedKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, spokeSubnet.EVMChainID)
	Expect(err).Should(BeNil())

	sendRegisterTx, err := tokenSpoke.RegisterWithHub(
		opts,
		tokenspoke.TeleporterFeeInfo{
			FeeTokenAddress: feeTokenAddress,
			Amount:          feeAmount,
		},
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, spokeSubnet, sendRegisterTx.Hash())

	// Relay the register message to the hub
	receipt = network.RelayMessage(ctx, receipt, spokeSubnet, hubSubnet, true)

	// Check that the spoke registered event was emitted
	tokenHub, err := tokenhub.NewTokenHub(hubAddress, hubSubnet.RPCClient)
	Expect(err).Should(BeNil())
	registerEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, tokenHub.ParseSpokeRegistered)
	Expect(err).Should(BeNil())
	Expect(registerEvent.SpokeBlockchainID[:]).Should(Equal(spokeSubnet.BlockchainID[:]))
	Expect(registerEvent.SpokeBridgeAddress).Should(Equal(spokeAddress))

	// Based on the initial reserve balance of the spoke instance,
	// calculate the collateral amount of hub tokens needed to collateralize the spoke.
	collateralNeeded := calculateCollateralNeeded(
		expectedInitialReserveBalance,
		expectedTokenMultiplier,
		expectedmultiplyOnSpoke,
	)
	teleporterUtils.ExpectBigEqual(registerEvent.InitialCollateralNeeded, collateralNeeded)

	return collateralNeeded
}

// AddCollateralToERC20TokenHub adds collateral to the ERC20TokenHub contract
// and verifies the collateral was added successfully. Any excess amount
// is returned to the caller.
func AddCollateralToERC20TokenHub(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenHub *erc20tokenhub.ERC20TokenHub,
	erc20TokenHubAddress common.Address,
	exampleERC20 *exampleerc20.ExampleERC20Decimals,
	spokeBlockchainID ids.ID,
	spokeAddress common.Address,
	collateralAmount *big.Int,
	senderKey *ecdsa.PrivateKey,
) {
	// Approve the ERC20TokenHub to spend the collateral
	ERC20Approve(
		ctx,
		exampleERC20,
		erc20TokenHubAddress,
		collateralAmount,
		subnet,
		senderKey,
	)

	// Add collateral to the ERC20TokenHub
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenHub.AddCollateral(
		opts,
		spokeBlockchainID,
		spokeAddress,
		collateralAmount,
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHub.ParseCollateralAdded)
	Expect(err).Should(BeNil())
	Expect(event.SpokeBlockchainID[:]).Should(Equal(spokeBlockchainID[:]))
	Expect(event.SpokeBridgeAddress).Should(Equal(spokeAddress))

	spokeSettings, err := erc20TokenHub.RegisteredSpokes(
		&bind.CallOpts{},
		spokeBlockchainID,
		spokeAddress)
	Expect(err).Should(BeNil())
	if collateralAmount.Cmp(spokeSettings.CollateralNeeded) > 0 {
		collateralAmount.Sub(collateralAmount, spokeSettings.CollateralNeeded)
	}
	teleporterUtils.ExpectBigEqual(event.Amount, collateralAmount)
	teleporterUtils.ExpectBigEqual(event.Remaining, big.NewInt(0))
}

// AddCollateralToNativeTokenHub adds collateral to the NativeTokenHub contract
// and verifies the collateral was added successfully. Any excess amount
// is returned to the caller.
func AddCollateralToNativeTokenHub(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenHub *nativetokenhub.NativeTokenHub,
	nativeTokenHubAddress common.Address,
	spokeBlockchainID ids.ID,
	spokeAddress common.Address,
	collateralAmount *big.Int,
	senderKey *ecdsa.PrivateKey,
) {
	// Add collateral to the ERC20TokenHub
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = collateralAmount

	tx, err := nativeTokenHub.AddCollateral(
		opts,
		spokeBlockchainID,
		spokeAddress,
	)
	Expect(err).Should(BeNil())
	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenHub.ParseCollateralAdded)
	Expect(err).Should(BeNil())
	Expect(event.SpokeBlockchainID[:]).Should(Equal(spokeBlockchainID[:]))
	Expect(event.SpokeBridgeAddress).Should(Equal(spokeAddress))
	spokeSettings, err := nativeTokenHub.RegisteredSpokes(
		&bind.CallOpts{},
		spokeBlockchainID,
		spokeAddress)
	Expect(err).Should(BeNil())
	if collateralAmount.Cmp(spokeSettings.CollateralNeeded) > 0 {
		collateralAmount.Sub(collateralAmount, spokeSettings.CollateralNeeded)
	}
	teleporterUtils.ExpectBigEqual(event.Amount, collateralAmount)
	teleporterUtils.ExpectBigEqual(event.Remaining, big.NewInt(0))
}

func SendERC20TokenHub(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenHub *erc20tokenhub.ERC20TokenHub,
	erc20TokenHubAddress common.Address,
	token *exampleerc20.ExampleERC20Decimals,
	input erc20tokenhub.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	// Approve the ERC20TokenHub to spend the tokens
	ERC20Approve(
		ctx,
		token,
		erc20TokenHubAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		subnet,
		senderKey,
	)

	// Send the tokens and verify expected events
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenHub.Send(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHub.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))

	// Compute the scaled amount
	scaledAmount := GetScaledAmountFromERC20TokenHub(
		erc20TokenHub,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendNativeTokenHub(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenHub *nativetokenhub.NativeTokenHub,
	nativeTokenHubAddress common.Address,
	wrappedToken *wrappednativetoken.WrappedNativeToken,
	input nativetokenhub.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	DepositAndApproveWrappedTokenForFees(
		ctx,
		subnet,
		wrappedToken,
		input.PrimaryFee,
		nativeTokenHubAddress,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenHub.Send(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenHub.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))

	// Compute the scaled amount
	scaledAmount := GetScaledAmountFromNativeTokenHub(
		nativeTokenHub,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendNativeTokenSpoke(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenSpoke *nativetokenspoke.NativeTokenSpoke,
	nativeTokenSpokeAddress common.Address,
	input nativetokenspoke.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	DepositAndApproveWrappedTokenForFees(
		ctx,
		subnet,
		nativeTokenSpoke,
		input.PrimaryFee,
		nativeTokenSpokeAddress,
		senderKey,
	)

	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenSpoke.Send(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenSpoke.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	return receipt, event.Amount
}

func SendERC20TokenSpoke(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenSpoke *erc20tokenspoke.ERC20TokenSpoke,
	erc20TokenSpokeAddress common.Address,
	input erc20tokenspoke.SendTokensInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenSpoke.Approve(
		opts,
		erc20TokenSpokeAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Bridge the tokens back to subnet A
	tx, err = erc20TokenSpoke.Send(
		opts,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenSpoke.ParseTokensSent)
	Expect(err).Should(BeNil())
	Expect(event.Sender).Should(Equal(crypto.PubkeyToAddress(senderKey.PublicKey)))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	return receipt, event.Amount
}

func SendAndCallERC20TokenHub(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenHub *erc20tokenhub.ERC20TokenHub,
	erc20TokenHubAddress common.Address,
	exampleToken *exampleerc20.ExampleERC20Decimals,
	input erc20tokenhub.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	// Approve the ERC20TokenHub to spend the tokens
	ERC20Approve(
		ctx,
		exampleToken,
		erc20TokenHubAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
		subnet,
		senderKey,
	)

	// Send the tokens and verify expected events
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenHub.SendAndCall(
		optsA,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHub.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))

	// Computer the scaled amount
	scaledAmount := GetScaledAmountFromERC20TokenHub(
		erc20TokenHub,
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendAndCallNativeTokenHub(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenHub *nativetokenhub.NativeTokenHub,
	input nativetokenhub.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenHub.SendAndCall(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenHub.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))

	// Compute the scaled amount
	spokeSettings, err := nativeTokenHub.RegisteredSpokes(
		&bind.CallOpts{},
		input.DestinationBlockchainID,
		input.DestinationBridgeAddress)
	Expect(err).Should(BeNil())

	scaledAmount := ApplyTokenScaling(
		spokeSettings.TokenMultiplier,
		spokeSettings.MultiplyOnSpoke,
		amount,
	)
	teleporterUtils.ExpectBigEqual(event.Amount, scaledAmount)

	return receipt, event.Amount
}

func SendAndCallNativeTokenSpoke(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	nativeTokenSpoke *nativetokenspoke.NativeTokenSpoke,
	input nativetokenspoke.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
	tokenMultiplier *big.Int,
	multiplyOnSpoke bool,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	opts.Value = amount

	tx, err := nativeTokenSpoke.SendAndCall(
		opts,
		input,
	)
	Expect(err).Should(BeNil())

	bridgedAmount := big.NewInt(0).Sub(amount, input.PrimaryFee)

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, nativeTokenSpoke.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))
	teleporterUtils.ExpectBigEqual(event.Amount, bridgedAmount)

	return receipt, event.Amount
}

func SendAndCallERC20TokenSpoke(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
	erc20TokenSpoke *erc20tokenspoke.ERC20TokenSpoke,
	erc20TokenSpokeAddress common.Address,
	input erc20tokenspoke.SendAndCallInput,
	amount *big.Int,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := erc20TokenSpoke.Approve(
		opts,
		erc20TokenSpokeAddress,
		big.NewInt(0).Add(amount, input.PrimaryFee),
	)
	Expect(err).Should(BeNil())

	teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	// Bridge the tokens back to subnet A
	tx, err = erc20TokenSpoke.SendAndCall(
		opts,
		input,
		amount,
	)
	Expect(err).Should(BeNil())

	receipt := teleporterUtils.WaitForTransactionSuccess(ctx, subnet, tx.Hash())
	event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenSpoke.ParseTokensAndCallSent)
	Expect(err).Should(BeNil())
	Expect(event.Input.RecipientContract).Should(Equal(input.RecipientContract))
	teleporterUtils.ExpectBigEqual(event.Amount, amount)

	return receipt, event.Amount
}

// Send a native token from fromBridge to toBridge via multi-hop through the C-Chain
// Requires that both fromBridge and toBridge are fully collateralized
// Requires that both fromBridge and toBridge have the same tokenMultiplier and multiplyOnSpoke
// with respect to the original asset on the C-Chain
func SendNativeMultiHopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	sendingKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *nativetokenspoke.NativeTokenSpoke,
	fromBridgeAddress common.Address,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *nativetokenspoke.NativeTokenSpoke,
	toBridgeAddress common.Address,
	cChainInfo interfaces.SubnetTestInfo,
	bridgedAmount *big.Int,
) {
	input := nativetokenspoke.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   fromBridgeAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         DefaultNativeTokenRequiredGas,
		MultiHopFallback:         recipientAddress,
	}

	// Find the amount sent by fromBridge. This is before any scaling/unscaling is applied.
	bridgedAmount = teleporterUtils.BigIntSub(bridgedAmount, input.PrimaryFee)

	// Send tokens through a multi-hop transfer
	originReceipt, bridgedAmount := SendNativeTokenSpoke(
		ctx,
		fromSubnet,
		fromBridge,
		fromBridgeAddress,
		input,
		bridgedAmount,
		sendingKey,
	)

	// Relay the first message back to the hub chain, in this case C-Chain,
	// which then performs the multi-hop transfer to the destination spoke
	intermediateReceipt := network.RelayMessage(
		ctx,
		originReceipt,
		fromSubnet,
		cChainInfo,
		true,
	)

	initialBalance, err := toSubnet.RPCClient.BalanceAt(ctx, recipientAddress, nil)
	Expect(err).Should(BeNil())

	// When we relay the above message to the hub chain, a multi-hop transfer
	// is performed to the destination spoke. Parse for the send tokens event
	// and relay to destination spoke.
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

func SendERC20TokenMultiHopAndVerify(
	ctx context.Context,
	network interfaces.Network,
	fundedKey *ecdsa.PrivateKey,
	sendingKey *ecdsa.PrivateKey,
	recipientAddress common.Address,
	fromSubnet interfaces.SubnetTestInfo,
	fromBridge *erc20tokenspoke.ERC20TokenSpoke,
	fromBridgeAddress common.Address,
	toSubnet interfaces.SubnetTestInfo,
	toBridge *erc20tokenspoke.ERC20TokenSpoke,
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
	input := erc20tokenspoke.SendTokensInput{
		DestinationBlockchainID:  toSubnet.BlockchainID,
		DestinationBridgeAddress: toBridgeAddress,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   common.Address{},
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         DefaultERC20RequiredGas,
		MultiHopFallback:         recipientAddress,
	}

	// Send tokens through a multi-hop transfer
	originReceipt, bridgedAmount := SendERC20TokenSpoke(
		ctx,
		fromSubnet,
		fromBridge,
		fromBridgeAddress,
		input,
		bridgedAmount,
		sendingKey,
	)

	// Relay the first message back to the hub chain, in this case C-Chain,
	// which then performs the multi-hop transfer to the destination spoke
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

	// When we relay the above message to the hub chain, a multi-hop transfer
	// is performed to the destination spoke. Parse for the send tokens event
	// and relay to destination spoke.
	spokeReceipt := network.RelayMessage(
		ctx,
		intermediateReceipt,
		cChainInfo,
		toSubnet,
		true,
	)
	_, err = teleporterUtils.GetEventFromLogs(spokeReceipt.Logs, toSubnet.TeleporterMessenger.ParseMessageExecuted)
	if err != nil {
		teleporterUtils.TraceTransactionAndExit(ctx, toSubnet, spokeReceipt.TxHash)
	}

	CheckERC20TokenSpokeWithdrawal(
		ctx,
		toBridge,
		spokeReceipt,
		recipientAddress,
		bridgedAmount,
	)

	balance, err := toBridge.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	teleporterUtils.ExpectBigEqual(balance, big.NewInt(0).Add(initialBalance, bridgedAmount))
}

func CheckERC20TokenHubWithdrawal(
	ctx context.Context,
	erc20TokenHubAddress common.Address,
	exampleERC20 *exampleerc20.ExampleERC20Decimals,
	receipt *types.Receipt,
	expectedRecipientAddress common.Address,
	expectedAmount *big.Int,
) {
	hubTransferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, exampleERC20.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(hubTransferEvent.From).Should(Equal(erc20TokenHubAddress))
	Expect(hubTransferEvent.To).Should(Equal(expectedRecipientAddress))
	teleporterUtils.ExpectBigEqual(hubTransferEvent.Value, expectedAmount)
}

func CheckERC20TokenSpokeWithdrawal(
	ctx context.Context,
	erc20TokenSpoke *erc20tokenspoke.ERC20TokenSpoke,
	receipt *types.Receipt,
	expectedRecipientAddress common.Address,
	expectedAmount *big.Int,
) {
	transferEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenSpoke.ParseTransfer)
	Expect(err).Should(BeNil())
	Expect(transferEvent.From).Should(Equal(common.Address{}))
	Expect(transferEvent.To).Should(Equal(expectedRecipientAddress))
	teleporterUtils.ExpectBigEqual(transferEvent.Value, expectedAmount)
}

func CheckNativeTokenHubWithdrawal(
	ctx context.Context,
	nativeTokenHubAddress common.Address,
	wrappedNativeToken *wrappednativetoken.WrappedNativeToken,
	receipt *types.Receipt,
	expectedAmount *big.Int,
) {
	withdrawalEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, wrappedNativeToken.ParseWithdrawal)
	Expect(err).Should(BeNil())
	Expect(withdrawalEvent.Sender).Should(Equal(nativeTokenHubAddress))
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
