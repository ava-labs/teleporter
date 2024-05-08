package flows

import (
	"context"
	"math/big"

	nativetokendestination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenDestination"
	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenSource"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

var (
	decimalsShift           = uint8(1)
	tokenMultiplier         = utils.GetTokenMultiplier(decimalsShift)
	initialReserveImbalance = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e6))
	multiplyOnReceive       = true

	burnedFeesReportingRewardPercentage = big.NewInt(1)
)

/**
 * Deploy a native token source on the primary network
 * Deploys a native token destination to Subnet A
 * Bridges C-Chain native tokens to Subnet A
 * Bridge back tokens from Subnet A to C-Chain
 */
func NativeSourceNativeDestination(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example WAVAX on the primary network
	cChainWAVAXAddress, wavaxA := utils.DeployExampleWAVAX(
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
		cChainWAVAXAddress,
	)

	// Deploy an NativeTokenDestination to Subnet A
	nativeTokenDestinationAddress, nativeTokenDestination, deployReceipt := utils.DeployNativeTokenDestination(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenSourceAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnReceive,
		burnedFeesReportingRewardPercentage,
	)

	// Register the NativeTokenDestination on the NativeTokenSource
	collateralAmount := utils.RegisterNativeTokenDestinationOnNativeTokenSource(
		ctx,
		network,
		cChainInfo,
		nativeTokenSource,
		subnetAInfo,
		nativeTokenDestinationAddress,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		!multiplyOnReceive,
		deployReceipt,
	)

	utils.AddCollateralToNativeTokenSource(
		ctx,
		cChainInfo,
		nativeTokenSource,
		nativeTokenSourceAddress,
		subnetAInfo.BlockchainID,
		nativeTokenDestinationAddress,
		collateralAmount,
		fundedKey,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A that fully collateralize bridge with leftover tokens.
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))
	{
		input := nativetokensource.SendTokensInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenDestinationAddress,
			Recipient:                recipientAddress,
			PrimaryFee:               big.NewInt(0),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
		}

		// Send initialReserveImbalance tokens to fully collateralize bridge and mint the remainder.
		receipt, _ := utils.SendNativeTokenSource(
			ctx,
			cChainInfo,
			nativeTokenSource,
			input,
			utils.RemoveTokenScaling(tokenMultiplier, !multiplyOnReceive, amount),
			fundedKey,
		)

		network.RelayMessage(
			ctx,
			receipt,
			cChainInfo,
			subnetAInfo,
			true,
		)

		teleporterUtils.CheckBalance(
			ctx,
			recipientAddress,
			amount,
			subnetAInfo.RPCClient,
		)
	}

	// Send tokens on Subnet A back for native tokens on C-Chain
	{
		input_A := nativetokendestination.SendTokensInput{
			DestinationBlockchainID:  cChainInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenSourceAddress,
			Recipient:                recipientAddress,
			PrimaryFee:               big.NewInt(0),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
		}

		// Send half of the tokens back to C-Chain
		amount := big.NewInt(0).Div(amount, big.NewInt(2))
		receipt, bridgedAmount := utils.SendNativeTokenDestination(
			ctx,
			subnetAInfo,
			nativeTokenDestination,
			input_A,
			amount,
			recipientKey,
		)

		receipt = network.RelayMessage(
			ctx,
			receipt,
			subnetAInfo,
			cChainInfo,
			true,
		)

		// Check that the recipient received the tokens
		sourceAmount := utils.RemoveTokenScaling(tokenMultiplier, !multiplyOnReceive, bridgedAmount)
		utils.CheckNativeTokenSourceWithdrawal(
			ctx,
			nativeTokenSourceAddress,
			wavaxA,
			receipt,
			sourceAmount,
		)

		teleporterUtils.CheckBalance(ctx, recipientAddress, sourceAmount, cChainInfo.RPCClient)
	}
}
