package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func InsufficientGas(network *localnetwork.LocalNetwork) {
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := network.GetTwoSubnets()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Deploy TestMessenger to Subnets A
	_, subnetATestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		subnetAInfo,
	)
	// Deploy TestMessenger to Subnets B
	testMessengerContractB, subnetBTestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		subnetBInfo,
	)

	// Send message from SubnetA to SubnetB with 0 execution gas, which should fail to execute
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := subnetATestMessenger.SendMessage(
		optsA, subnetBInfo.BlockchainID, testMessengerContractB, fundedAddress, big.NewInt(0), big.NewInt(0), message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	event, err := utils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	messageID := event.MessageID

	// Relay message from SubnetA to SubnetB
	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// Check Teleporter message received on the destination
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check message execution failed event
	failedMessageExecutionEvent, err := utils.GetEventFromLogs(
		receipt.Logs, subnetBInfo.TeleporterMessenger.ParseMessageExecutionFailed,
	)
	Expect(err).Should(BeNil())
	Expect(failedMessageExecutionEvent.MessageID[:]).Should(Equal(messageID[:]))
	Expect(failedMessageExecutionEvent.SourceBlockchainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))

	// Retry message execution. This will execute the message with as much gas as needed
	// (up to the transaction gas limit), rather than using the required gas specified in the message itself.
	receipt = utils.RetryMessageExecutionAndWaitForAcceptance(
		ctx,
		subnetAInfo.BlockchainID,
		subnetBInfo,
		failedMessageExecutionEvent.Message,
		fundedKey,
	)
	executedEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetBInfo.TeleporterMessenger.ParseMessageExecuted)
	Expect(err).Should(BeNil())
	Expect(executedEvent.MessageID[:]).Should(Equal(messageID[:]))
	Expect(executedEvent.SourceBlockchainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))

	//
	// Verify we received the expected string
	//
	_, currMessage, err := subnetBTestMessenger.GetCurrentMessage(&bind.CallOpts{}, subnetAInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))
}
