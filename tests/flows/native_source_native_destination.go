package flows

import (
	"context"
	"math"
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
	tokenMultiplier         = big.NewInt(int64(math.Pow10(int(decimalsShift))))
	initialReserveImbalance = new(big.Int).Mul(big.NewInt(1e15), big.NewInt(1e9))
	valueToReceive          = new(big.Int).Div(initialReserveImbalance, big.NewInt(4))
	valueToSend             = new(big.Int).Div(valueToReceive, tokenMultiplier)
	valueToReturn           = new(big.Int).Div(valueToReceive, big.NewInt(4))
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
	utils.RegisterNativeTokenDestinationOnNativeTokenSource(
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

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A that don't fully collateralize bridge
	{
		input := nativetokensource.SendTokensInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenDestinationAddress,
			Recipient:                recipientAddress,
			PrimaryFee:               big.NewInt(0),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
		}

		receipt, _ := utils.SendNativeTokenSource(
			ctx,
			cChainInfo,
			nativeTokenSource,
			input,
			valueToSend,
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
			big.NewInt(0),
			subnetAInfo.RPCClient,
		)
	}

	// Send tokens from C-Chain to recipient on subnet A that fully collateralize bridge with leftover tokens.
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
			new(big.Int).Div(initialReserveImbalance, tokenMultiplier),
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
			valueToReceive,
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

		receipt, bridgedAmount := utils.SendNativeTokenDestination(
			ctx,
			subnetAInfo,
			nativeTokenDestination,
			input_A,
			valueToReturn,
			recipientKey,
			tokenMultiplier,
			multiplyOnReceive,
		)

		receipt = network.RelayMessage(
			ctx,
			receipt,
			subnetAInfo,
			cChainInfo,
			true,
		)

		// Check that the recipient received the tokens
		utils.CheckNativeTokenSourceWithdrawal(
			ctx,
			nativeTokenSourceAddress,
			wavaxA,
			receipt,
			bridgedAmount,
		)

		teleporterUtils.CheckBalance(ctx, recipientAddress, bridgedAmount, cChainInfo.RPCClient)
	}
}
