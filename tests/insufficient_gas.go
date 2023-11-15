package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	localUtils "github.com/ava-labs/teleporter/tests/utils/local-network-utils"
	. "github.com/onsi/gomega"
)

func InsufficientGasGinkgo() {
	InsufficientGas(&network.LocalNetwork{})
}

func InsufficientGas(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())

	// Deploy ExampleMessenger to Subnets A
	_, subnetAExampleMessenger := localUtils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetAInfo)

	// Deploy ExampleMessenger to Subnets B
	exampleMessengerContractB, _ := localUtils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetBInfo)

	// Send message from SubnetA to SubnetB with 0 execution gas, which should fail to execute
	message := "Hello, world!"
	optsA := utils.CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)
	tx, err := subnetAExampleMessenger.SendMessage(optsA, subnetBInfo.BlockchainID, exampleMessengerContractB, fundedAddress, big.NewInt(0), big.NewInt(0), message)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	event, err := utils.GetEventFromLogs(receipt.Logs, subnetATeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	messageID := event.Message.MessageID

	// Relay message from SubnetA to SubnetB
	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// Check Teleporter message received on the destination
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check message execution failed event
	failedMessageExecutionEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetBTeleporterMessenger.ParseMessageExecutionFailed)
	Expect(err).Should(BeNil())
	Expect(failedMessageExecutionEvent.MessageID).Should(Equal(messageID))
	Expect(failedMessageExecutionEvent.OriginChainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))

	// Retry message execution. This will execute the message with as much gas as needed
	// (up to the transaction gas limit), rather than using the required gas specified in the message itself.s
	receipt = utils.RetryMessageExecutionAndWaitForAcceptance(
		ctx, subnetAInfo.BlockchainID, subnetBInfo, failedMessageExecutionEvent.Message, fundedAddress, fundedKey, subnetBTeleporterMessenger)
	executedEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetBTeleporterMessenger.ParseMessageExecuted)
	Expect(err).Should(BeNil())
	Expect(executedEvent.MessageID).Should(Equal(messageID))
	Expect(executedEvent.OriginChainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))
}
