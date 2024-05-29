package flows

import (
	"context"
	"math/big"

	nativetokenhub "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHub/NativeTokenHub"
	nativetokenspoke "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenSpoke/NativeTokenSpoke"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHub on the primary network
 * Deploys a NativeTokenSpoke to Subnet A
 * Bridges C-Chain native tokens to Subnet A
 * Bridge back tokens from Subnet A to C-Chain
 */
func NativeTokenHubNativeDestination(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an example WAVAX on the primary network
	cChainWAVAXAddress, wavax := utils.DeployWrappedNativeToken(
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
		cChainWAVAXAddress,
	)

	// Deploy a NativeTokenSpoke to Subnet A
	nativeTokenSpokeAddress, nativeTokenSpoke := utils.DeployNativeTokenSpoke(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHubAddress,
		18,
		initialReserveImbalance,
		multiplyOnSpoke,
		burnedFeesReportingRewardPercentage,
	)

	// Register the NativeTokenSpoke on the NativeTokenHub
	collateralAmount := utils.RegisterTokenSpokeOnHub(
		ctx,
		network,
		cChainInfo,
		nativeTokenHubAddress,
		subnetAInfo,
		nativeTokenSpokeAddress,
		initialReserveImbalance,
		big.NewInt(1),
		multiplyOnSpoke,
	)

	utils.AddCollateralToNativeTokenHub(
		ctx,
		cChainInfo,
		nativeTokenHub,
		nativeTokenHubAddress,
		subnetAInfo.BlockchainID,
		nativeTokenSpokeAddress,
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
		input := nativetokenhub.SendTokensInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenSpokeAddress,
			Recipient:                recipientAddress,
			PrimaryFeeTokenAddress:   cChainWAVAXAddress,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
		}

		// Send initialReserveImbalance tokens to fully collateralize bridge and mint the remainder.
		receipt, _ := utils.SendNativeTokenHub(
			ctx,
			cChainInfo,
			nativeTokenHub,
			nativeTokenHubAddress,
			wavax,
			input,
			amount,
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
		input_A := nativetokenspoke.SendTokensInput{
			DestinationBlockchainID:  cChainInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenHubAddress,
			Recipient:                recipientAddress,
			PrimaryFeeTokenAddress:   nativeTokenSpokeAddress,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
		}

		// Send half of the tokens back to C-Chain
		amount := big.NewInt(0).Div(amount, big.NewInt(2))
		receipt, bridgedAmount := utils.SendNativeTokenSpoke(
			ctx,
			subnetAInfo,
			nativeTokenSpoke,
			nativeTokenSpokeAddress,
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
		hubAmount := bridgedAmount
		utils.CheckNativeTokenHubWithdrawal(
			ctx,
			nativeTokenHubAddress,
			wavax,
			receipt,
			hubAmount,
		)

		teleporterUtils.CheckBalance(ctx, recipientAddress, hubAmount, cChainInfo.RPCClient)
	}
}
