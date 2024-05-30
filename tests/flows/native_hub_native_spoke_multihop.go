package flows

import (
	"context"
	"math/big"

	nativetokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/NativeTokenHub"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHub on the primary network
 * Deploys NativeTokenSpoke to Subnet A and Subnet B
 * Bridges native tokens from the C-Chain to Subnet A as Subnet A's native token
 * Bridges native tokens from the C-Chain to Subnet B as Subnet B's native token to collateralize the Subnet B bridge
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func NativeTokenHubNativeTokenSpokeMultiHop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// decimalsShift is always 0 for native to native
	decimalsShift := uint8(0)

	// Deploy an example WAVAX on the primary network
	wavaxAddress, wavax := utils.DeployWrappedNativeToken(
		ctx,
		fundedKey,
		cChainInfo,
		"AVAX",
	)

	// Create a NativeTokenHub on the primary network
	nativeTokenHubAddress, nativeTokenHub := utils.DeployNativeTokenHub(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Deploy a NativeTokenSpoke to Subnet A
	nativeTokenSpokeAddressA, nativeTokenSpokeA := utils.DeployNativeTokenSpoke(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHubAddress,
		utils.NativeTokenDecimals,
		initialReserveImbalance,
		multiplyOnSpoke,
		burnedFeesReportingRewardPercentage,
	)

	// Deploy a NativeTokenSpoke to Subnet B
	nativeTokenSpokeAddressB, nativeTokenSpokeB := utils.DeployNativeTokenSpoke(
		ctx,
		subnetBInfo,
		"SUBB",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHubAddress,
		utils.NativeTokenDecimals,
		initialReserveImbalance,
		multiplyOnSpoke,
		burnedFeesReportingRewardPercentage,
	)

	// Register both NativeTokenDestinations on the NativeTokenHub
	collateralAmountA := utils.RegisterTokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		nativeTokenHubAddress,
		subnetAInfo,
		nativeTokenSpokeAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnSpoke,
	)

	collateralAmountB := utils.RegisterTokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		nativeTokenHubAddress,
		subnetBInfo,
		nativeTokenSpokeAddressB,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnSpoke,
	)

	// Add collateral for both NativeTokenDestinations
	utils.AddCollateralToNativeTokenHub(
		ctx,
		cChainInfo,
		nativeTokenHub,
		nativeTokenHubAddress,
		subnetAInfo.BlockchainID,
		nativeTokenSpokeAddressA,
		collateralAmountA,
		fundedKey,
	)

	utils.AddCollateralToNativeTokenHub(
		ctx,
		cChainInfo,
		nativeTokenHub,
		nativeTokenHubAddress,
		subnetBInfo.BlockchainID,
		nativeTokenSpokeAddressB,
		collateralAmountB,
		fundedKey,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))

	// Send tokens from C-Chain to Subnet A
	inputA := nativetokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenSpokeAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   wavaxAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	receipt, bridgedAmountA := utils.SendNativeTokenHub(
		ctx,
		cChainInfo,
		nativeTokenHub,
		nativeTokenHubAddress,
		wavax,
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
	inputB := nativetokenhub.SendTokensInput{
		DestinationBlockchainID:  subnetBInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenSpokeAddressB,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   wavaxAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}
	receipt, bridgedAmountB := utils.SendNativeTokenHub(
		ctx,
		cChainInfo,
		nativeTokenHub,
		nativeTokenHubAddress,
		wavax,
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
	amountToSendA := new(big.Int).Div(bridgedAmountA, big.NewInt(2))

	utils.SendNativeMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientAddress,
		subnetAInfo,
		nativeTokenSpokeA,
		nativeTokenSpokeAddressA,
		subnetBInfo,
		nativeTokenSpokeB,
		nativeTokenSpokeAddressB,
		cChainInfo,
		amountToSendA,
		big.NewInt(0),
	)

	// Again, send half of the received amount to account for gas expenses
	amountToSendB := new(big.Int).Div(amountToSendA, big.NewInt(2))
	secondaryFeeAmount := new(big.Int).Div(amountToSendB, big.NewInt(4))

	// Multi-hop transfer back to Subnet A
	utils.SendNativeMultiHopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientAddress,
		subnetBInfo,
		nativeTokenSpokeB,
		nativeTokenSpokeAddressB,
		subnetAInfo,
		nativeTokenSpokeA,
		nativeTokenSpokeAddressA,
		cChainInfo,
		amountToSendB,
		secondaryFeeAmount,
	)
}
