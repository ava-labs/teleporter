package teleporter

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	testmessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/tests/TestMessenger"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func RetrySuccessfulExecution(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	L1AInfo := network.GetPrimaryNetworkInfo()
	L1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy TestMessenger to L1s A and B
	//
	ctx := context.Background()

	_, l1ATestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(L1AInfo),
		L1AInfo,
	)
	testMessengerContractAddressB, l1BTestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(L1BInfo),
		L1BInfo,
	)

	//
	// Call the test messenger contract on L1 A
	//
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, L1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := l1ATestMessenger.SendMessage(
		optsA,
		L1BInfo.BlockchainID,
		testMessengerContractAddressB,
		fundedAddress,
		big.NewInt(0),
		testmessenger.SendMessageRequiredGas,
		message,
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

	teleporterMessageID := event.MessageID

	//
	// Relay the message to the destination
	//
	receipt = teleporter.RelayTeleporterMessage(ctx, receipt, L1AInfo, L1BInfo, true, fundedKey)
	receiveEvent, err :=
		utils.GetEventFromLogs(receipt.Logs, teleporter.TeleporterMessenger(L1BInfo).ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	deliveredTeleporterMessage := receiveEvent.Message

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := teleporter.TeleporterMessenger(L1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Verify we received the expected string
	//
	_, currMessage, err := l1BTestMessenger.GetCurrentMessage(&bind.CallOpts{}, L1AInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))

	//
	// Attempt to retry message execution, which should fail
	//
	optsB, err := bind.NewKeyedTransactorWithChainID(fundedKey, L1BInfo.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err = teleporter.TeleporterMessenger(L1BInfo).RetryMessageExecution(
		optsB,
		L1AInfo.BlockchainID,
		deliveredTeleporterMessage,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(tx).Should(BeNil())
}
