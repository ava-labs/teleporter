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
 * Bridge tokens from Subnet A to Subnet B through multihop
 * Bridge back tokens from Subnet B to Subnet A through multihop
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

	// Deploy an example WAVAX on Subnet A
	wavaxAddressA, _ := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		subnetAInfo,
	)

	// Deploy an example WAVAX on Subnet B
	wavaxAddressB, _ := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		subnetBInfo,
	)

	// Deploy a NativeTokenDestination to Subnet A
	nativeTokenDestinationAddressA, nativeTokenDestinationA := utils.DeployNativeTokenDestination(
		ctx,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		wavaxAddressA,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnReceive,
	)

	// Deploy a NativeTokenDestination to Subnet B
	nativeTokenDestinationAddressB, nativeTokenDestinationB := utils.DeployNativeTokenDestination(
		ctx,
		subnetBInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		wavaxAddressB,
		initialReserveImbalance,
		decimalsShift,
		multiplyOnReceive,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// These are set during the initial bridging, and used in the multihop transfers
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
		receivedAmountA = big.NewInt(0).Sub(big.NewInt(0).Mul(bridgedAmountA, big.NewInt(10)), initialReserveImbalance)

		// Relay the message to subnet A and check for a native token mint withdrawal
		receipt = network.RelayMessage(
			ctx,
			receipt,
			cChainInfo,
			subnetAInfo,
			true,
		)

		utils.CheckNativeTokenDestinationMint(
			ctx,
			nativeTokenDestinationA,
			recipientAddress,
			receipt,
			receivedAmountA,
		)
		teleporterUtils.CheckBalance(
			ctx,
			recipientAddress,
			receivedAmountA,
			subnetAInfo.RPCClient,
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
		receivedAmountB = big.NewInt(0).Sub(big.NewInt(0).Mul(bridgedAmountB, big.NewInt(10)), initialReserveImbalance)

		// Relay the message to subnet B and check for a native token mint withdrawal
		receipt = network.RelayMessage(
			ctx,
			receipt,
			cChainInfo,
			subnetBInfo,
			true,
		)

		utils.CheckNativeTokenDestinationMint(
			ctx,
			nativeTokenDestinationB,
			recipientAddress,
			receipt,
			receivedAmountB,
		)
		teleporterUtils.CheckBalance(
			ctx,
			recipientAddress,
			receivedAmountB,
			subnetBInfo.RPCClient,
		)

		// Verify the recipient received the tokens
		teleporterUtils.CheckBalance(ctx, recipientAddress, receivedAmountB, subnetBInfo.RPCClient)
	}

	// Multihop transfer to Subnet B
	// Send half of the received amount to account for gas expenses
	amountToSendA := big.NewInt(0).Div(receivedAmountA, big.NewInt(2))

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
	amountToSendB := big.NewInt(0).Div(amountToSendA, big.NewInt(2))

	// Multihop transfer back to Subnet A
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
