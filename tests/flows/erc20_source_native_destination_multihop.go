package flows

import (
	"context"
	"math/big"

	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy a ERC20 token source on the primary network
 * Deploys NativeDestination to Subnet A and Subnet B
 * Bridges C-Chain example ERC20 tokens to Subnet A as Subnet A's native token
 * Bridges C-Chain example ERC20 tokens to Subnet B as Subnet B's native token to collateralize the bridge on Subnet B
 * Bridge tokens from Subnet A to Subnet B through multi-hop
 * Bridge back tokens from Subnet B to Subnet A through multi-hop
 */
func ERC20SourceNativeDestinationMultihop(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, subnetBInfo := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on subnet A as the source token to be bridged
	sourceTokenAddress, sourceToken := teleporterUtils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
	)

	// Create an ERC20Source for bridging the source token
	erc20SourceAddress, erc20Source := utils.DeployERC20Source(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		sourceTokenAddress,
	)

	// Deploy a NativeTokenDestination to Subnet A
	nativeTokenDestinationAddressA, nativeTokenDestinationA, deployReceipt_A := utils.DeployNativeTokenDestination(
		ctx,
		subnetAInfo,
		"SUBA",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnReceive,
		burnedFeesReportingRewardPercentage,
	)

	// Deploy a NativeTokenDestination to Subnet B
	nativeTokenDestinationAddressB, nativeTokenDestinationB, deployReceipt_B := utils.DeployNativeTokenDestination(
		ctx,
		subnetBInfo,
		"SUBB",
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnReceive,
		burnedFeesReportingRewardPercentage,
	)

	// Register both NativeTokenDestinations on the ERC20Source
	utils.RegisterNativeTokenDestinationOnERC20Source(
		ctx,
		network,
		cChainInfo,
		erc20Source,
		subnetAInfo,
		nativeTokenDestinationAddressA,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		!multiplyOnReceive,
		deployReceipt_A,
	)

	utils.RegisterNativeTokenDestinationOnERC20Source(
		ctx,
		network,
		cChainInfo,
		erc20Source,
		subnetBInfo,
		nativeTokenDestinationAddressB,
		initialReserveImbalance,
		utils.GetTokenMultiplier(decimalsShift),
		!multiplyOnReceive,
		deployReceipt_B,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// These are set during the initial bridging, and used in the multi-hop transfers
	var receivedAmountA, receivedAmountB *big.Int

	// Send tokens from C-Chain to Subnet A
	{
		inputA := erc20source.SendTokensInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenDestinationAddressA,
			Recipient:                recipientAddress,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
		}
		// Bridge the initial imbalance, which is scaled up on the destination by 1 decimal place
		// This will mint 9/10 the initial imbalance amount on the destination (after fees)
		receipt, bridgedAmountA := utils.SendERC20Source(
			ctx,
			cChainInfo,
			erc20Source,
			erc20SourceAddress,
			sourceToken,
			inputA,
			initialReserveImbalance,
			fundedKey,
		)

		// Amount received by the destination is bridgedAmount * 10 - initialReserveImbalance
		receivedAmountA = new(big.Int).Sub(new(big.Int).Mul(bridgedAmountA, big.NewInt(10)), initialReserveImbalance)

		// Relay the message to subnet A and check for a native token mint withdrawal
		network.RelayMessage(
			ctx,
			receipt,
			cChainInfo,
			subnetAInfo,
			true,
		)

		// Verify the recipient received the tokens
		teleporterUtils.CheckBalance(ctx, recipientAddress, receivedAmountA, subnetAInfo.RPCClient)
	}

	// Send tokens from C-Chain to Subnet B
	{
		inputB := erc20source.SendTokensInput{
			DestinationBlockchainID:  subnetBInfo.BlockchainID,
			DestinationBridgeAddress: nativeTokenDestinationAddressB,
			Recipient:                recipientAddress,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultNativeTokenRequiredGasLimit,
		}
		// Bridge the initial imbalance, which is scaled up on the destination by 1 decimal place
		// This will mint 9/10 the initial imbalance amount on the destination (after fees)
		receipt, bridgedAmountB := utils.SendERC20Source(
			ctx,
			cChainInfo,
			erc20Source,
			erc20SourceAddress,
			sourceToken,
			inputB,
			initialReserveImbalance,
			fundedKey,
		)

		// Amount received by the destination is bridgedAmount * 10 - initialReserveImbalance
		receivedAmountB = new(big.Int).Sub(new(big.Int).Mul(bridgedAmountB, big.NewInt(10)), initialReserveImbalance)

		// Relay the message to subnet B and check for a native token mint withdrawal
		network.RelayMessage(
			ctx,
			receipt,
			cChainInfo,
			subnetBInfo,
			true,
		)

		// Verify the recipient received the tokens
		teleporterUtils.CheckBalance(ctx, recipientAddress, receivedAmountB, subnetBInfo.RPCClient)
	}

	// Multi-hop transfer to Subnet B
	// Send half of the received amount to account for gas expenses
	amountToSendA := new(big.Int).Div(receivedAmountA, big.NewInt(2))

	utils.SendNativeMultihopAndVerify(
		ctx,
		network,
		fundedKey,
		recipientAddress,
		subnetAInfo,
		nativeTokenDestinationA,
		subnetBInfo,
		nativeTokenDestinationB,
		nativeTokenDestinationAddressB,
		cChainInfo,
		amountToSendA,
		tokenMultiplier,
		multiplyOnReceive,
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
		subnetAInfo,
		nativeTokenDestinationA,
		nativeTokenDestinationAddressA,
		cChainInfo,
		amountToSendB,
		tokenMultiplier,
		multiplyOnReceive,
	)
}
