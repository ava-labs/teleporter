package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/ERC20TokenHome"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a ERC20TokenHome on the primary network
 * Deploys NativeTokenRemote to Subnet A and Subnet B
 * Bridges C-Chain example ERC20 tokens to Subnet A as Subnet A's native token
 * Bridges C-Chain example ERC20 tokens to Subnet B as Subnet B's native token to collateralize the bridge on Subnet B
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func ERC20TokenHomeNativeTokenRemoteMultiHop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the token to be bridged
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHomeDecimals,
	)

	exampleERC20Decimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	erc20TokenHomeAddress, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		exampleERC20Decimals,
	)

	// Deploy a NativeTokenRemote to Subnet A
	nativeTokenRemoteAddressA, nativeTokenRemoteA := utils.DeployNativeTokenRemote(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		exampleERC20Decimals,
		initialReserveImbalance,
		multiplyOnRemote,
		burnedFeesReportingRewardPercentage,
	)

	// Deploy a NativeTokenRemote to Subnet B
	nativeTokenRemoteAddressB, nativeTokenRemoteB := utils.DeployNativeTokenRemote(
		ctx,
		subnetBInfo,
		"SUBB",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		exampleERC20Decimals,
		initialReserveImbalance,
		multiplyOnRemote,
		burnedFeesReportingRewardPercentage,
	)

	// Register both NativeTokenDestinations on the ERC20TokenHome
	collateralAmountA := utils.RegisterTokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		erc20TokenHomeAddress,
		subnetAInfo,
		nativeTokenRemoteAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnRemote,
	)

	collateralAmountB := utils.RegisterTokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		erc20TokenHomeAddress,
		subnetBInfo,
		nativeTokenRemoteAddressB,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnRemote,
	)

	// Add collateral for both NativeTokenDestinations
	utils.AddCollateralToERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		subnetAInfo.BlockchainID,
		nativeTokenRemoteAddressA,
		collateralAmountA,
		fundedKey,
	)

	utils.AddCollateralToERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		subnetBInfo.BlockchainID,
		nativeTokenRemoteAddressB,
		collateralAmountB,
		fundedKey,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// These are set during the initial bridging, and used in the multi-hop transfers
	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))

	// Send tokens from C-Chain to Subnet A
	inputA := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenRemoteAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	receipt, bridgedAmountA := utils.SendERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		inputA,
		amount,
		fundedKey,
	)

	// Relay the message to subnet A and check for a native token mint withdrawal
	network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	// Verify the recipient received the tokens
	teleporterUtils.CheckBalance(ctx, recipientAddress, bridgedAmountA, subnetAInfo.RPCClient)

	// Send tokens from C-Chain to Subnet B
	inputB := erc20tokenhome.SendTokensInput{
		DestinationBlockchainID:  subnetBInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenRemoteAddressB,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	receipt, bridgedAmountB := utils.SendERC20TokenHome(
		ctx,
		cChainInfo,
		erc20TokenHome,
		erc20TokenHomeAddress,
		exampleERC20,
		inputB,
		amount,
		fundedKey,
	)

	// Relay the message to subnet B and check for a native token mint withdrawal
	network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetBInfo,
		true,
	)

	// Verify the recipient received the tokens
	teleporterUtils.CheckBalance(ctx, recipientAddress, bridgedAmountB, subnetBInfo.RPCClient)

	// Multi-hop transfer to Subnet B
	// Send half of the received amount to account for gas expenses
	amountToSend := new(big.Int).Div(bridgedAmountA, big.NewInt(2))

	utils.SendNativeMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientAddress,
		subnetAInfo,
		nativeTokenRemoteA,
		nativeTokenRemoteAddressA,
		subnetBInfo,
		nativeTokenRemoteB,
		nativeTokenRemoteAddressB,
		cChainInfo,
		amountToSend,
		big.NewInt(0),
	)

	// Multi-hop transfer back to Subnet A
	secondaryFeeAmount := new(big.Int).Div(amountToSend, big.NewInt(4))
	utils.SendNativeMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientAddress,
		subnetBInfo,
		nativeTokenRemoteB,
		nativeTokenRemoteAddressB,
		subnetAInfo,
		nativeTokenRemoteA,
		nativeTokenRemoteAddressA,
		cChainInfo,
		amountToSend,
		secondaryFeeAmount,
	)
}
