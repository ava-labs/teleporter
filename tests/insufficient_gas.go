package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
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
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Deploy ExampleMessenger to Subnets A
	_, subnetAExampleMessenger := localUtils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetAInfo)

	// Deploy ExampleMessenger to Subnets B
	exampleMessengerContractB, subnetBExampleMessenger :=
		localUtils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetBInfo)

	// Send message from SubnetA to SubnetB with 0 execution gas, which should fail to execute
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetAInfo.ChainIDInt)
	Expect(err).Should(BeNil())
	tx, err := subnetAExampleMessenger.SendMessage(
		optsA, subnetBInfo.BlockchainID, exampleMessengerContractB, fundedAddress, big.NewInt(0), big.NewInt(0), message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	event, err := utils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	messageID := event.Message.MessageID

	// Relay message from SubnetA to SubnetB
	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, true)

	// Check Teleporter message received on the destination
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check message execution failed event
	failedMessageExecutionEvent, err := utils.GetEventFromLogs(
		receipt.Logs, subnetBInfo.TeleporterMessenger.ParseMessageExecutionFailed,
	)
	Expect(err).Should(BeNil())
	Expect(failedMessageExecutionEvent.MessageID).Should(Equal(messageID))
	Expect(failedMessageExecutionEvent.OriginChainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))

	// Retry message execution. This will execute the message with as much gas as needed
	// (up to the transaction gas limit), rather than using the required gas specified in the message itself.
	receipt = utils.RetryMessageExecutionAndWaitForAcceptance(
		ctx,
		subnetAInfo.BlockchainID,
		subnetBInfo,
		failedMessageExecutionEvent.Message,
		fundedAddress,
		fundedKey,
		subnetBInfo.TeleporterMessenger,
	)
	executedEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetBInfo.TeleporterMessenger.ParseMessageExecuted)
	Expect(err).Should(BeNil())
	Expect(executedEvent.MessageID).Should(Equal(messageID))
	Expect(executedEvent.OriginChainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))

	//
	// Verify we received the expected string
	//
	_, currMessage, err := subnetBExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, subnetAInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))
}
