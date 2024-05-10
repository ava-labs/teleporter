package flows

import (
	"context"
	"math/big"

	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenSource"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a native token source on the primary network
 * Deploys NativeDestination to Subnet A and Subnet B
 * Bridges native tokens from the C-Chain to Subnet A as Subnet A's native token
 * Bridges native tokens from the C-Chain to Subnet B as Subnet B's native token to collateralize the Subnet B bridge
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func NativeSourceNativeDestinationMultihop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example WAVAX on the primary network
	wavaxAddress, wavax := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		cChainInfo,
	)

	// Create a NativeTokenSource on the primary network
	nativeTokenSourceAddress, nativeTokenSource := utils.DeployNativeTokenSource(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Deploy a NativeTokenDestination to Subnet A
	nativeTokenDestinationAddressA, nativeTokenDestinationA := utils.DeployNativeTokenDestination(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenSourceAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnDestination,
		burnedFeesReportingRewardPercentage,
	)

	// Deploy a NativeTokenDestination to Subnet B
	nativeTokenDestinationAddressB, nativeTokenDestinationB := utils.DeployNativeTokenDestination(
		ctx,
		subnetBInfo,
		"SUBB",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenSourceAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnDestination,
		burnedFeesReportingRewardPercentage,
	)

	// Register both NativeTokenDestinations on the NativeTokenSource
	collateralAmountA := utils.RegisterTokenDestinationOnSource(
		ctx,
		network,
		cChainInfo,
		nativeTokenSourceAddress,
		subnetAInfo,
		nativeTokenDestinationAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnDestination,
	)

	collateralAmountB := utils.RegisterTokenDestinationOnSource(
		ctx,
		network,
		cChainInfo,
		nativeTokenSourceAddress,
		subnetBInfo,
		nativeTokenDestinationAddressB,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnDestination,
	)

	// Add collateral for both NativeTokenDestinations
	utils.AddCollateralToNativeTokenSource(
		ctx,
		cChainInfo,
		nativeTokenSource,
		nativeTokenSourceAddress,
		subnetAInfo.BlockchainID,
		nativeTokenDestinationAddressA,
		collateralAmountA,
		fundedKey,
	)

	utils.AddCollateralToNativeTokenSource(
		ctx,
		cChainInfo,
		nativeTokenSource,
		nativeTokenSourceAddress,
		subnetBInfo.BlockchainID,
		nativeTokenDestinationAddressB,
		collateralAmountB,
		fundedKey,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))

	// Send tokens from C-Chain to Subnet A
	inputA := nativetokensource.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenDestinationAddressA,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
	}

	receipt, bridgedAmountA := utils.SendNativeTokenSource(
		ctx,
		cChainInfo,
		nativeTokenSource,
		nativeTokenSourceAddress,
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
	inputB := nativetokensource.SendTokensInput{
		DestinationBlockchainID:  subnetBInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenDestinationAddressB,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
	}
	receipt, bridgedAmountB := utils.SendNativeTokenSource(
		ctx,
		cChainInfo,
		nativeTokenSource,
		nativeTokenSourceAddress,
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

	utils.SendNativeMultihopAndVerify(
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
		amountToSendA,
	)

	// Again, send half of the received amount to account for gas expenses
	amountToSendB := new(big.Int).Div(amountToSendA, big.NewInt(2))

	// Multi-hop transfer back to Subnet A
	utils.SendNativeMultihopAndVerify(
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
		amountToSendB,
	)
}
