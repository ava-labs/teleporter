package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Destination"
	erc20source "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Source"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy an ERC20 token source on the primary network
 * Deploys ERC20Destination to Subnet A
 * Bridges C-Chain example ERC20 tokens to Subnet A and calls contract on Subnet A using sendAndCall
 * Bridges C-Chain example ERC20 to EOA on Subnet A, and then bridge tokens from Subnet A back
 * C-Chain and calls contract on the C-Chain using sendAndCall
 */
func ERC20SourceERC20DestinationSendAndCall(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on the primary network as the source token to be bridged
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

	sourceMockERC20SACRAddress, sourceMockERC20SACR := utils.DeployMockERC20SendAndCallReceiver(
		ctx,
		fundedKey,
		cChainInfo,
	)

	destMockERC20SACRAddress, destMockERC20SACR := utils.DeployMockERC20SendAndCallReceiver(
		ctx,
		fundedKey,
		subnetAInfo,
	)

	// Token representation on subnet A will have same name, symbol, and decimals
	tokenName, err := sourceToken.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := sourceToken.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := sourceToken.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20Destination to Subnet A
	erc20DestinationAddress, erc20Destination := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20SourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Generate new recipient to receive bridged tokens
	fallbackKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	fallbackAddress := crypto.PubkeyToAddress(fallbackKey.PublicKey)

	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(13))
	primaryFee := big.NewInt(1e18)
	var bridgedAmount = teleporterUtils.BigIntSub(amount, primaryFee)

	// Send tokens from C-Chain to Mock contract on subnet A
	{
		input := erc20source.SendAndCallInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: erc20DestinationAddress,
			RecipientContract:        destMockERC20SACRAddress,
			RecipientPayload:         []byte{1},
			RequiredGasLimit:         teleporterUtils.BigIntMul(big.NewInt(10), utils.DefaultERC20RequiredGasLimit),
			RecipientGasLimit:        teleporterUtils.BigIntMul(big.NewInt(5), utils.DefaultERC20RequiredGasLimit),
			FallbackRecipient:        fallbackAddress,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
		}

		receipt, bridgedAmount := utils.SendAndCallERC20Source(
			ctx,
			cChainInfo,
			erc20Source,
			erc20SourceAddress,
			sourceToken,
			input,
			amount,
			fundedKey,
		)

		// Relay the message to Subnet A and check for message delivery
		receipt = network.RelayMessage(
			ctx,
			receipt,
			cChainInfo,
			subnetAInfo,
			true,
		)

		event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Destination.ParseCallSucceeded)
		Expect(err).Should(BeNil())
		Expect(event.RecipientContract).Should(Equal(input.RecipientContract))
		Expect(event.Amount).Should(Equal(bridgedAmount))

		receiverEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, destMockERC20SACR.ParseTokensReceived)
		Expect(err).Should(BeNil())
		Expect(receiverEvent.Amount).Should(Equal(bridgedAmount))
		Expect(receiverEvent.Payload).Should(Equal(input.RecipientPayload))

		// Check that the contract received the tokens
		balance, err := erc20Destination.BalanceOf(&bind.CallOpts{}, destMockERC20SACRAddress)
		Expect(err).Should(BeNil())
		Expect(balance).Should(Equal(bridgedAmount))
	}

	// Bridge ERC20 tokens to account on subnet A
	{
		// Send ERC20 tokens from C-Chain to recipient on subnet A
		input := erc20source.SendTokensInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: erc20DestinationAddress,
			Recipient:                recipientAddress,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultERC20RequiredGasLimit,
		}

		receipt, bridgedAmount, _ := utils.SendERC20Source(
			ctx,
			cChainInfo,
			erc20Source,
			erc20SourceAddress,
			sourceToken,
			input,
			amount,
			fundedKey,
		)

		// Relay the message to Subnet A and check for message delivery
		receipt = network.RelayMessage(
			ctx,
			receipt,
			cChainInfo,
			subnetAInfo,
			true,
		)

		utils.CheckERC20DestinationWithdrawal(
			ctx,
			erc20Destination,
			receipt,
			recipientAddress,
			bridgedAmount,
		)

		// Check that the recipient received the tokens
		balance, err := erc20Destination.BalanceOf(&bind.CallOpts{}, recipientAddress)
		Expect(err).Should(BeNil())
		Expect(balance).Should(Equal(bridgedAmount))
	}

	// Send tokens to Mock contract on C-Chain using Send and Call
	{
		// Fund recipient with gas tokens on subnet A
		teleporterUtils.SendNativeTransfer(
			ctx,
			subnetAInfo,
			fundedKey,
			recipientAddress,
			big.NewInt(1e18),
		)

		inputB := erc20destination.SendAndCallInput{
			DestinationBlockchainID:  cChainInfo.BlockchainID,
			DestinationBridgeAddress: erc20SourceAddress,
			RecipientContract:        sourceMockERC20SACRAddress,
			RecipientPayload:         []byte{1},
			RequiredGasLimit:         teleporterUtils.BigIntMul(big.NewInt(10), utils.DefaultERC20RequiredGasLimit),
			RecipientGasLimit:        teleporterUtils.BigIntMul(big.NewInt(5), utils.DefaultERC20RequiredGasLimit),
			FallbackRecipient:        fallbackAddress,
			PrimaryFee:               big.NewInt(0),
			SecondaryFee:             big.NewInt(0),
		}

		receipt, bridgedAmount := utils.SendAndCallERC20Destination(
			ctx,
			subnetAInfo,
			erc20Destination,
			erc20DestinationAddress,
			inputB,
			bridgedAmount,
			recipientKey,
		)

		receipt = network.RelayMessage(
			ctx,
			receipt,
			subnetAInfo,
			cChainInfo,
			true,
		)

		sourceEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20Source.ParseCallSucceeded)
		Expect(err).Should(BeNil())
		Expect(sourceEvent.RecipientContract).Should(Equal(inputB.RecipientContract))
		Expect(sourceEvent.Amount).Should(Equal(bridgedAmount))

		receiverEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, sourceMockERC20SACR.ParseTokensReceived)
		Expect(err).Should(BeNil())
		Expect(receiverEvent.Amount).Should(Equal(bridgedAmount))
		Expect(receiverEvent.Payload).Should(Equal(inputB.RecipientPayload))

		// Check that the recipient received the tokens
		balance, err := sourceToken.BalanceOf(&bind.CallOpts{}, sourceMockERC20SACRAddress)
		Expect(err).Should(BeNil())
		Expect(balance).Should(Equal(bridgedAmount))
	}
}
