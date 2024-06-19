package flows

import (
	"context"
	"math/big"

	nativetokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/NativeTokenHome"
	nativetokenremote "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenRemote/NativeTokenRemote"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys a NativeTokenRemote to Subnet A
 * Bridges C-Chain native tokens to Subnet A
 * Bridge back tokens from Subnet A to C-Chain
 */
func NativeTokenHomeNativeDestination(network interfaces.Network) {
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

	// Create a NativeTokenHome on the primary network
	nativeTokenHomeAddress, nativeTokenHome := utils.DeployNativeTokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		cChainWAVAXAddress,
	)

	// Deploy a NativeTokenRemote to Subnet A
	nativeTokenRemoteAddress, nativeTokenRemote := utils.DeployNativeTokenRemote(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHomeAddress,
		18,
		initialReserveImbalance,
		multiplyOnRemote,
		burnedFeesReportingRewardPercentage,
	)

	// Register the NativeTokenRemote on the NativeTokenHome
	collateralAmount := utils.RegisterTokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		nativeTokenHomeAddress,
		subnetAInfo,
		nativeTokenRemoteAddress,
		initialReserveImbalance,
		big.NewInt(1),
		multiplyOnRemote,
	)

	utils.AddCollateralToNativeTokenHome(
		ctx,
		cChainInfo,
		nativeTokenHome,
		nativeTokenHomeAddress,
		subnetAInfo.BlockchainID,
		nativeTokenRemoteAddress,
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
		input := nativetokenhome.SendTokensInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenRemoteAddress,
			Recipient:                recipientAddress,
			PrimaryFeeTokenAddress:   cChainWAVAXAddress,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
		}

		// Send initialReserveImbalance tokens to fully collateralize bridge and mint the remainder.
		receipt, _ := utils.SendNativeTokenHome(
			ctx,
			cChainInfo,
			nativeTokenHome,
			nativeTokenHomeAddress,
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
		input_A := nativetokenremote.SendTokensInput{
			DestinationBlockchainID:  cChainInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenHomeAddress,
			Recipient:                recipientAddress,
			PrimaryFeeTokenAddress:   nativeTokenRemoteAddress,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGas,
		}

		// Send half of the tokens back to C-Chain
		amount := big.NewInt(0).Div(amount, big.NewInt(2))
		receipt, bridgedAmount := utils.SendNativeTokenRemote(
			ctx,
			subnetAInfo,
			nativeTokenRemote,
			nativeTokenRemoteAddress,
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
		homeAmount := bridgedAmount
		utils.CheckNativeTokenHomeWithdrawal(
			ctx,
			nativeTokenHomeAddress,
			wavax,
			receipt,
			homeAmount,
		)

		teleporterUtils.CheckBalance(ctx, recipientAddress, homeAmount, cChainInfo.RPCClient)
	}
}
