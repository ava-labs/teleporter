package ictt

import (
	"context"
	"math/big"

	nativetokenhome "github.com/ava-labs/teleporter/abi-bindings/go/ictt/TokenHome/NativeTokenHome"
	nativetokenremote "github.com/ava-labs/teleporter/abi-bindings/go/ictt/TokenRemote/NativeTokenRemote"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a NativeTokenHome on the primary network
 * Deploys a NativeTokenRemote to Subnet A
 * Transfers C-Chain native tokens to Subnet A
 * Transfer back tokens from Subnet A to C-Chain
 */
func NativeTokenHomeNativeDestination(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := network.GetTwoSubnets()
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
		teleporter,
		fundedKey,
		cChainInfo,
		fundedAddress,
		cChainWAVAXAddress,
	)

	// Deploy a NativeTokenRemote to Subnet A
	nativeTokenRemoteAddress, nativeTokenRemote := utils.DeployNativeTokenRemote(
		ctx,
		teleporter,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		nativeTokenHomeAddress,
		18,
		initialReserveImbalance,
		burnedFeesReportingRewardPercentage,
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Register the NativeTokenRemote on the NativeTokenHome
	collateralAmount := utils.RegisterTokenRemoteOnHome(
		ctx,
		teleporter,
		cChainInfo,
		nativeTokenHomeAddress,
		subnetAInfo,
		nativeTokenRemoteAddress,
		initialReserveImbalance,
		big.NewInt(1),
		multiplyOnRemote,
		fundedKey,
		aggregator,
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

	// Generate new recipient to receive transferred tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from C-Chain to recipient on subnet A that fully collateralize token transferrer with leftover tokens.
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(13))
	{
		input := nativetokenhome.SendTokensInput{
			DestinationBlockchainID:            subnetAInfo.BlockchainID,
			DestinationTokenTransferrerAddress: nativeTokenRemoteAddress,
			Recipient:                          recipientAddress,
			PrimaryFeeTokenAddress:             cChainWAVAXAddress,
			PrimaryFee:                         big.NewInt(1e18),
			SecondaryFee:                       big.NewInt(0),
			RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
		}

		// Send initialReserveImbalance tokens to fully collateralize token transferrer and mint the remainder.
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

		teleporter.RelayTeleporterMessage(
			ctx,
			receipt,
			cChainInfo,
			subnetAInfo,
			true,
			fundedKey,
			nil,
			aggregator,
		)

		utils.CheckBalance(
			ctx,
			recipientAddress,
			amount,
			subnetAInfo.RPCClient,
		)
	}

	// Send tokens on Subnet A back for native tokens on C-Chain
	{
		input_A := nativetokenremote.SendTokensInput{
			DestinationBlockchainID:            cChainInfo.BlockchainID,
			DestinationTokenTransferrerAddress: nativeTokenHomeAddress,
			Recipient:                          recipientAddress,
			PrimaryFeeTokenAddress:             nativeTokenRemoteAddress,
			PrimaryFee:                         big.NewInt(1e18),
			SecondaryFee:                       big.NewInt(0),
			RequiredGasLimit:                   utils.DefaultNativeTokenRequiredGas,
		}

		// Send half of the tokens back to C-Chain
		amount := big.NewInt(0).Div(amount, big.NewInt(2))
		receipt, transferredAmount := utils.SendNativeTokenRemote(
			ctx,
			subnetAInfo,
			nativeTokenRemote,
			nativeTokenRemoteAddress,
			input_A,
			amount,
			recipientKey,
		)

		receipt = teleporter.RelayTeleporterMessage(
			ctx,
			receipt,
			subnetAInfo,
			cChainInfo,
			true,
			fundedKey,
			nil,
			aggregator,
		)

		// Check that the recipient received the tokens
		homeAmount := transferredAmount
		utils.CheckNativeTokenHomeWithdrawal(
			ctx,
			nativeTokenHomeAddress,
			wavax,
			receipt,
			homeAmount,
		)

		utils.CheckBalance(ctx, recipientAddress, homeAmount, cChainInfo.RPCClient)
	}
}
