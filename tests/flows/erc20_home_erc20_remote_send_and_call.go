package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20tokenhome "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenHome/ERC20TokenHome"
	erc20tokenremote "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/TokenRemote/ERC20TokenRemote"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

/**
 * Deploy an ERC20TokenHome on the primary network
 * Deploys ERC20TokenRemote to Subnet A
 * Bridges C-Chain example ERC20 tokens to Subnet A and calls contract on Subnet A using sendAndCall
 * Bridges C-Chain example ERC20 to EOA on Subnet A, and then bridge tokens from Subnet A back
 * C-Chain and calls contract on the C-Chain using sendAndCall
 */
func ERC20TokenHomeERC20TokenRemoteSendAndCall(network interfaces.Network) {
	cChainInfo := network.GetPrimaryNetworkInfo()
	subnetAInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy an ExampleERC20 on the primary network as the token to be bridged
	exampleERC20Address, exampleERC20 := utils.DeployExampleERC20(
		ctx,
		fundedKey,
		cChainInfo,
		erc20TokenHomeDecimals,
	)

	exampleERC20Decimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Create an ERC20TokenHome for bridging the ERC20 token
	erc20TokenHomeAddress, erc20TokenHome := utils.DeployERC20TokenHome(
		ctx,
		fundedKey,
		cChainInfo,
		fundedAddress,
		exampleERC20Address,
		exampleERC20Decimals,
	)

	homeMockERC20SACRAddress, homeMockERC20SACR := utils.DeployMockERC20SendAndCallReceiver(
		ctx,
		fundedKey,
		cChainInfo,
	)

	remoteMockERC20SACRAddress, remoteMockERC20SACR := utils.DeployMockERC20SendAndCallReceiver(
		ctx,
		fundedKey,
		subnetAInfo,
	)

	// Token representation on subnet A will have same name, symbol, and decimals
	tokenName, err := exampleERC20.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := exampleERC20.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := exampleERC20.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20TokenRemote to Subnet A
	erc20TokenRemoteAddress, erc20TokenRemote := utils.DeployERC20TokenRemote(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		cChainInfo.BlockchainID,
		erc20TokenHomeAddress,
		exampleERC20Decimals,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	utils.RegisterERC20TokenRemoteOnHome(
		ctx,
		network,
		cChainInfo,
		erc20TokenHomeAddress,
		subnetAInfo,
		erc20TokenRemoteAddress,
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
	bridgedAmount := teleporterUtils.BigIntSub(amount, primaryFee)

	// Send tokens from C-Chain to Mock contract on subnet A
	{
		input := erc20tokenhome.SendAndCallInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: erc20TokenRemoteAddress,
			RecipientContract:        remoteMockERC20SACRAddress,
			RecipientPayload:         []byte{1},
			RequiredGasLimit:         teleporterUtils.BigIntMul(big.NewInt(10), utils.DefaultERC20RequiredGas),
			RecipientGasLimit:        teleporterUtils.BigIntMul(big.NewInt(5), utils.DefaultERC20RequiredGas),
			FallbackRecipient:        fallbackAddress,
			PrimaryFeeTokenAddress:   exampleERC20Address,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
		}

		receipt, bridgedAmount := utils.SendAndCallERC20TokenHome(
			ctx,
			cChainInfo,
			erc20TokenHome,
			erc20TokenHomeAddress,
			exampleERC20,
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

		event, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenRemote.ParseCallSucceeded)
		Expect(err).Should(BeNil())
		Expect(event.RecipientContract).Should(Equal(input.RecipientContract))
		Expect(event.Amount).Should(Equal(bridgedAmount))

		receiverEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, remoteMockERC20SACR.ParseTokensReceived)
		Expect(err).Should(BeNil())
		Expect(receiverEvent.Amount).Should(Equal(bridgedAmount))
		Expect(receiverEvent.Payload).Should(Equal(input.RecipientPayload))

		// Check that the contract received the tokens
		balance, err := erc20TokenRemote.BalanceOf(&bind.CallOpts{}, remoteMockERC20SACRAddress)
		Expect(err).Should(BeNil())
		Expect(balance).Should(Equal(bridgedAmount))
	}

	// Bridge ERC20 tokens to account on subnet A
	{
		// Send ERC20 tokens from C-Chain to recipient on subnet A
		input := erc20tokenhome.SendTokensInput{
			DestinationBlockchainID:  subnetAInfo.BlockchainID,
			DestinationBridgeAddress: erc20TokenRemoteAddress,
			Recipient:                recipientAddress,
			PrimaryFeeTokenAddress:   exampleERC20Address,
			PrimaryFee:               big.NewInt(1e18),
			SecondaryFee:             big.NewInt(0),
			RequiredGasLimit:         utils.DefaultERC20RequiredGas,
		}

		receipt, bridgedAmount := utils.SendERC20TokenHome(
			ctx,
			cChainInfo,
			erc20TokenHome,
			erc20TokenHomeAddress,
			exampleERC20,
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

		utils.CheckERC20TokenRemoteWithdrawal(
			ctx,
			erc20TokenRemote,
			receipt,
			recipientAddress,
			bridgedAmount,
		)

		// Check that the recipient received the tokens
		balance, err := erc20TokenRemote.BalanceOf(&bind.CallOpts{}, recipientAddress)
		Expect(err).Should(BeNil())
		Expect(balance).Should(Equal(bridgedAmount))
	}

	// Send tokens to mock contract on C-Chain using sendAndCall
	{
		// Fund recipient with gas tokens on subnet A
		teleporterUtils.SendNativeTransfer(
			ctx,
			subnetAInfo,
			fundedKey,
			recipientAddress,
			big.NewInt(1e18),
		)

		inputB := erc20tokenremote.SendAndCallInput{
			DestinationBlockchainID:  cChainInfo.BlockchainID,
			DestinationBridgeAddress: erc20TokenHomeAddress,
			RecipientContract:        homeMockERC20SACRAddress,
			RecipientPayload:         []byte{1},
			RequiredGasLimit:         teleporterUtils.BigIntMul(big.NewInt(10), utils.DefaultERC20RequiredGas),
			RecipientGasLimit:        teleporterUtils.BigIntMul(big.NewInt(5), utils.DefaultERC20RequiredGas),
			FallbackRecipient:        fallbackAddress,
			PrimaryFeeTokenAddress:   erc20TokenRemoteAddress,
			PrimaryFee:               big.NewInt(1e10),
			SecondaryFee:             big.NewInt(0),
		}

		receipt, bridgedAmount := utils.SendAndCallERC20TokenRemote(
			ctx,
			subnetAInfo,
			erc20TokenRemote,
			erc20TokenRemoteAddress,
			inputB,
			teleporterUtils.BigIntSub(bridgedAmount, inputB.PrimaryFee),
			recipientKey,
		)

		receipt = network.RelayMessage(
			ctx,
			receipt,
			subnetAInfo,
			cChainInfo,
			true,
		)

		homeEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, erc20TokenHome.ParseCallSucceeded)
		Expect(err).Should(BeNil())
		Expect(homeEvent.RecipientContract).Should(Equal(inputB.RecipientContract))
		Expect(homeEvent.Amount).Should(Equal(bridgedAmount))

		receiverEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, homeMockERC20SACR.ParseTokensReceived)
		Expect(err).Should(BeNil())
		Expect(receiverEvent.Amount).Should(Equal(bridgedAmount))
		Expect(receiverEvent.Payload).Should(Equal(inputB.RecipientPayload))

		// Check that the recipient received the tokens
		balance, err := exampleERC20.BalanceOf(&bind.CallOpts{}, homeMockERC20SACRAddress)
		Expect(err).Should(BeNil())
		Expect(balance).Should(Equal(bridgedAmount))
	}
}
