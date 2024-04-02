package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	erc20destination "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/ERC20Destination"
	nativetokensource "github.com/ava-labs/teleporter-token-bridge/abi-bindings/go/NativeTokenSource"
	"github.com/ava-labs/teleporter-token-bridge/tests/utils"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

// NativeSourceNativeDestination deploys NativeTokenSource to subnet A and NativeTokenDestination to subnet B.
// Then it sends native tokens from subnet A for a native token on subnet B, and transfers back to
// subnet A for its original native tokens.
func NativeSourceNativeDestination(network interfaces.Network) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := teleporterUtils.GetTwoSubnets(network)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	// Deploy a native token source on subnet A
	wavaxAddressA, wavaxA := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		subnetAInfo,
	)

	// Deploy a native token source on subnet A
	wavaxAddressB, wavaxB := utils.DeployExampleWAVAX(
		ctx,
		fundedKey,
		subnetBInfo,
	)

	nativeTokenSourceAddress, nativeTokenSource := utils.DeployNativeTokenSource(
		ctx,
		fundedKey,
		subnetAInfo,
		fundedAddress,
		wavaxAddressA,
	)

	// Token representation on subnet B will have same name, symbol, and decimals
	tokenName, err := wavaxA.Name(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenSymbol, err := wavaxA.Symbol(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	tokenDecimals, err := wavaxA.Decimals(&bind.CallOpts{})
	Expect(err).Should(BeNil())

	// Deploy an ERC20Destination on subnet B for the token source on subnet A
	erc20DestinationAddress, erc20Destination := utils.DeployERC20Destination(
		ctx,
		fundedKey,
		subnetBInfo,
		fundedAddress,
		subnetAInfo.BlockchainID,
		nativeTokenSourceAddress,
		tokenName,
		tokenSymbol,
		tokenDecimals,
	)

	// Generate new recipient to receive bridged tokens
	recipientKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	recipientAddress := crypto.PubkeyToAddress(recipientKey.PublicKey)

	// Send tokens from subnet A to recipient on subnet B
	input := nativetokensource.SendTokensInput{
		DestinationBlockchainID:  subnetBInfo.BlockchainID,
		DestinationBridgeAddress: erc20DestinationAddress,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}

	// Send the tokens and verify expected events
	amount := big.NewInt(2e18)
	receipt, teleporterMessageID, bridgedAmount := utils.SendNativeTokenSource(
		ctx,
		subnetAInfo,
		nativeTokenSource,
		nativeTokenSourceAddress,
		input,
		amount,
		fundedKey,
	)

	// Relay the message to subnet B and check for message delivery
	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetAInfo,
		subnetBInfo,
		true,
	)

	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	messageExecutedEvent, err := teleporterUtils.GetEventFromLogs(receipt.Logs, subnetBInfo.TeleporterMessenger.ParseMessageExecuted)
	Expect(err).Should(BeNil())
	Expect(messageExecutedEvent.MessageID).Should(Equal(teleporterMessageID))

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

	// Fund recipient with gas tokens on subnet B
	teleporterUtils.SendNativeTransfer(
		ctx,
		subnetBInfo,
		fundedKey,
		recipientAddress,
		big.NewInt(1e18),
	)
	inputB := erc20destination.SendTokensInput{
		DestinationBlockchainID:  subnetAInfo.BlockchainID,
		DestinationBridgeAddress: nativeTokenSourceAddress,
		Recipient:                recipientAddress,
		PrimaryFee:               big.NewInt(0),
		SecondaryFee:             big.NewInt(0),
		AllowedRelayerAddresses:  []common.Address{},
	}

	// Send tokens on subnet B back for native tokens on subnet A
	receipt, teleporterMessageID, bridgedAmount = utils.SendERC20Destination(
		ctx,
		subnetBInfo,
		erc20Destination,
		erc20DestinationAddress,
		inputB,
		bridgedAmount,
		recipientKey,
	)

	receipt = network.RelayMessage(
		ctx,
		receipt,
		subnetBInfo,
		subnetAInfo,
		true,
	)

	messageExecutedEvent, err = teleporterUtils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseMessageExecuted)
	Expect(err).Should(BeNil())
	Expect(messageExecutedEvent.MessageID).Should(Equal(teleporterMessageID))

	// Check that the recipient received the tokens
	utils.CheckNativeTokenSourceWithdrawal(
		ctx,
		nativeTokenSourceAddress,
		wavaxA,
		receipt,
		bridgedAmount,
	)

	teleporterUtils.CheckBalance(ctx, recipientAddress, bridgedAmount, subnetAInfo.RPCClient)
}
