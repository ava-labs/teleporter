package flows

import (
	"context"
	"math/big"

	erc20tokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/ERC20TokenHub"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a ERC20TokenHub on the primary network
 * Deploys NativeTokenSpoke to Subnet A and Subnet B
 * Bridges C-Chain example ERC20 tokens to Subnet A as Subnet A's native token
 * Bridges C-Chain example ERC20 tokens to Subnet B as Subnet B's native token to collateralize the bridge on Subnet B
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func ERC20TokenHubNativeTokenSpokeMultiHop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the token to be bridged
	exampleERC20Address, exampleERC20 := teleporterUtils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
	)

	// Create an ERC20TokenHub for bridging the ERC20 token
	erc20TokenHubAddress, erc20TokenHub := utils.DeployERC20TokenHub(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
	)

	// Deploy a NativeTokenSpoke to Subnet A
	nativeTokenDestinationAddressA, nativeTokenDestinationA := utils.DeployNativeTokenSpoke(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHubAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnSpoke,
		burnedFeesReportingRewardPercentage,
	)

	// Deploy a NativeTokenSpoke to Subnet B
	nativeTokenDestinationAddressB, nativeTokenDestinationB := utils.DeployNativeTokenSpoke(
		ctx,
		subnetBInfo,
		"SUBB",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHubAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnSpoke,
		burnedFeesReportingRewardPercentage,
	)

	// Register both NativeTokenDestinations on the ERC20TokenHub
	collateralAmountA := utils.RegisterTokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		erc20TokenHubAddress,
		subnetAInfo,
		nativeTokenDestinationAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnSpoke,
	)

	collateralAmountB := utils.RegisterTokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		erc20TokenHubAddress,
		subnetBInfo,
		nativeTokenDestinationAddressB,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnSpoke,
	)

	// Add collateral for both NativeTokenDestinations
	utils.AddCollateralToERC20TokenHub(
		ctx,
		cChainInfo,
		erc20TokenHub,
		erc20TokenHubAddress,
		exampleERC20,
		subnetAInfo.BlockchainID,
		nativeTokenDestinationAddressA,
		collateralAmountA,
		fundedKey,
	)

	utils.AddCollateralToERC20TokenHub(
		ctx,
		cChainInfo,
		erc20TokenHub,
		erc20TokenHubAddress,
		exampleERC20,
		subnetBInfo.BlockchainID,
		nativeTokenDestinationAddressB,
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
	inputA := erc20tokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenDestinationAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	receipt, bridgedAmountA := utils.SendERC20TokenHub(
		ctx,
		cChainInfo,
		erc20TokenHub,
		erc20TokenHubAddress,
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
	inputB := erc20tokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetBInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenDestinationAddressB,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   exampleERC20Address,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	receipt, bridgedAmountB := utils.SendERC20TokenHub(
		ctx,
		cChainInfo,
		erc20TokenHub,
		erc20TokenHubAddress,
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
		nativeTokenDestinationA,
		nativeTokenDestinationAddressA,
		subnetBInfo,
		nativeTokenDestinationB,
		nativeTokenDestinationAddressB,
		cChainInfo,
		amountToSend,
	)

	// Multi-hop transfer back to Subnet A
	utils.SendNativeMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientAddress,
		subnetBInfo,
		nativeTokenDestinationB,
		nativeTokenDestinationAddressB,
		subnetAInfo,
		nativeTokenDestinationA,
		nativeTokenDestinationAddressA,
		cChainInfo,
		amountToSend,
	)
}
