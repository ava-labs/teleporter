package flows

import (
	"context"
	"math/big"

	nativetokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/NativeTokenHome"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys NativeTokenRemote to Subnet A and Subnet B
 * Bridges native tokens from the C-Chain to Subnet A as Subnet A's native token
 * Bridges native tokens from the C-Chain to Subnet B as Subnet B's native token to collateralize the Subnet B bridge
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func NativeTokenHomeNativeTokenRemoteMultiHop(network interfaces.Network) {
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

	// Create a NativeTokenHome on the primary network
	nativeTokenHomeAddress, nativeTokenHome := utils.DeployNativeTokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		wavaxAddress,
	)

	// Deploy a NativeTokenRemote to Subnet A
	nativeTokenRemoteAddressA, nativeTokenRemoteA := utils.DeployNativeTokenRemote(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHomeAddress,
		utils.NativeTokenDecimals,
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
		nativeTokenHomeAddress,
		utils.NativeTokenDecimals,
		initialReserveImbalance,
		multiplyOnRemote,
		burnedFeesReportingRewardPercentage,
	)

	// Register both NativeTokenDestinations on the NativeTokenHome
	collateralAmountA := utils.RegisterTokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		nativeTokenHomeAddress,
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
		nativeTokenHomeAddress,
		subnetBInfo,
		nativeTokenRemoteAddressB,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		multiplyOnRemote,
	)

	// Add collateral for both NativeTokenDestinations
	utils.AddCollateralToNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		subnetAInfo.BlockchainID,
		nativeTokenRemoteAddressA,
		collateralAmountA,
		fundedKey,
	)

	utils.AddCollateralToNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		subnetBInfo.BlockchainID,
		nativeTokenRemoteAddressB,
		collateralAmountB,
		fundedKey,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10))

	// Send tokens from C-Chain to Subnet A
	inputA := nativetokenhome.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenRemoteAddressA,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   wavaxAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}

	receipt, bridgedAmountA := utils.SendNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
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
	inputB := nativetokenhome.SendTokensInput{
		DestinationBlockchainID:  subnetBInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenRemoteAddressB,
		Recipient:                recipientAddress,
		PrimaryFeeTokenAddress:   wavaxAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
	}
	receipt, bridgedAmountB := utils.SendNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
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
		nativeTokenRemoteA,
		nativeTokenRemoteAddressA,
		subnetBInfo,
		nativeTokenRemoteB,
		nativeTokenRemoteAddressB,
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
		nativeTokenRemoteB,
		nativeTokenRemoteAddressB,
		subnetAInfo,
		nativeTokenRemoteA,
		nativeTokenRemoteAddressA,
		cChainInfo,
		amountToSendB,
		secondaryFeeAmount,
	)
}
