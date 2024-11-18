package teleporter

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func InsufficientGas(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	L1AInfo := network.GetPrimaryNetworkInfo()
	L1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	ctx := context.Background()

	// Deploy TestMessenger to L1s A
	_, l1ATestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(L1AInfo),
		L1AInfo,
	)
	// Deploy TestMessenger to L1s B
	testMessengerContractB, l1BTestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(L1BInfo),
		L1BInfo,
	)

	// Send message from L1A to L1B with 0 execution gas, which should fail to execute
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, L1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := l1ATestMessenger.SendMessage(
		optsA, L1BInfo.BlockchainID, testMessengerContractB, fundedAddress, big.NewInt(0), big.NewInt(0), message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, L1AInfo, tx.Hash())

	event, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(L1AInfo).ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(L1BInfo.BlockchainID[:]))

	messageID := event.MessageID

	// Relay message from L1A to L1B
	receipt = teleporter.RelayTeleporterMessage(ctx, receipt, L1AInfo, L1BInfo, true, fundedKey)

	// Check Teleporter message received on the destination
	delivered, err :=
		teleporter.TeleporterMessenger(L1BInfo).MessageReceived(&bind.CallOpts{}, messageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// Check message execution failed event
	failedMessageExecutionEvent, err := utils.GetEventFromLogs(
		receipt.Logs, teleporter.TeleporterMessenger(L1BInfo).ParseMessageExecutionFailed,
	)
	Expect(err).Should(BeNil())
	Expect(failedMessageExecutionEvent.MessageID[:]).Should(Equal(messageID[:]))
	Expect(failedMessageExecutionEvent.SourceBlockchainID[:]).Should(Equal(L1AInfo.BlockchainID[:]))

	// Retry message execution. This will execute the message with as much gas as needed
	// (up to the transaction gas limit), rather than using the required gas specified in the message itself.
	receipt = utils.RetryMessageExecutionAndWaitForAcceptance(
		ctx,
		L1AInfo.BlockchainID,
		teleporter.TeleporterMessenger(L1BInfo),
		L1BInfo,
		failedMessageExecutionEvent.Message,
		fundedKey,
	)
	executedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(L1BInfo).ParseMessageExecuted,
	)
	Expect(err).Should(BeNil())
	Expect(executedEvent.MessageID[:]).Should(Equal(messageID[:]))
	Expect(executedEvent.SourceBlockchainID[:]).Should(Equal(L1AInfo.BlockchainID[:]))

	//
	// Verify we received the expected string
	//
	_, currMessage, err := l1BTestMessenger.GetCurrentMessage(&bind.CallOpts{}, L1AInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))
}
