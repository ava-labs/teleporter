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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func ERC20SourceMultihop(network interfaces.Network) {
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

	// Token representation on subnet B will have same name, symbol, and decimals
	tokenName, err := sourceToken.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := sourceToken.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := sourceToken.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20Destination for Subnet A
	erc20DestinationAddress_A, erc20Destination_A := utils.DeployERC20Destination(
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

	// Deploy an ERC20Destination for Subnet B
	erc20DestinationAddress_B, erc20Destination_B := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetBInfo,
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

	// Send tokens from subnet A to recipient on subnet B
	input := erc20source.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress_A,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(1e18),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}
	amount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(13))

	receipt, teleporterMessageID, bridgedAmount := utils.SendERC20Source(
		ctx,
		cChainInfo,
		erc20Source,
		erc20SourceAddress,
		sourceToken,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to subnet B and check for message delivery
	receipt = network.RelayMessage(
		ctx,
		receipt,
		cChainInfo,
		subnetAInfo,
		true,
	)

	messageExecutedEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseMessageExecuted)
	Expect(err).Should(BeNil())
	Expect(messageExecutedEvent.MessageID).Should(Equal(teleporterMessageID))

	utils.CheckERC20DestinationWithdrawal(
		ctx,
		erc20Destination_A,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err := erc20Destination_A.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))

	// Multihop transfer to Subnet B
	// Fund recipient with gas tokens on subnet A
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetAInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	input_A := erc20destination.SendTokensInput{
		DestinationBlockchainID:  subnetBInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress_B,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}

	receipt, teleporterMessageID, bridgedAmount = utils.SendERC20Destination(
		ctx,
		subnetAInfo,
		erc20Destination_A,
		erc20DestinationAddress_A,
		input_A,
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

	messageExecutedEvent, err = teleporterUtils.GetEventFromLogs(receipt.Logs, cChainInfo.TeleporterMessenger.ParseMessageExecuted)
	Expect(err).Should(BeNil())
	Expect(messageExecutedEvent.MessageID).Should(Equal(teleporterMessageID))

	utils.CheckERC20SourceWithdrawal(
		ctx,
		erc20SourceAddress,
		sourceToken,
		receipt,
		recipientAddress,
		bridgedAmount,
	)

	// Check that the recipient received the tokens
	balance, err = sourceToken.BalanceOf(&bind.CallOpts{}, recipientAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(bridgedAmount))
}
