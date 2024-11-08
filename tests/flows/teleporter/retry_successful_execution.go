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
	subnetAInfo := network.GetPrimaryNetworkInfo()
	subnetBInfo, _ := network.GetTwoSubnets()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy TestMessenger to Subnets A and B
	//
	ctx := context.Background()

	_, subnetATestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(subnetAInfo),
		subnetAInfo,
	)
	testMessengerContractAddressB, subnetBTestMessenger := utils.DeployTestMessenger(
		ctx,
		fundedKey,
		fundedAddress,
		teleporter.TeleporterRegistryAddress(subnetBInfo),
		subnetBInfo,
	)

	//
	// Call the test messenger contract on Subnet A
	//
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := subnetATestMessenger.SendMessage(
		optsA,
		subnetBInfo.BlockchainID,
		testMessengerContractAddressB,
		fundedAddress,
		big.NewInt(0),
		testmessenger.SendMessageRequiredGas,
		message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	event, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(subnetAInfo).ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID := event.MessageID

	//
	// Relay the message to the destination
	//
	receipt = teleporter.RelayTeleporterMessage(ctx, receipt, subnetAInfo, subnetBInfo, true, fundedKey, nil, network.GetSignatureAggregator())
	receiveEvent, err :=
		utils.GetEventFromLogs(receipt.Logs, teleporter.TeleporterMessenger(subnetBInfo).ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	deliveredTeleporterMessage := receiveEvent.Message

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := teleporter.TeleporterMessenger(subnetBInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Verify we received the expected string
	//
	_, currMessage, err := subnetBTestMessenger.GetCurrentMessage(&bind.CallOpts{}, subnetAInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))

	//
	// Attempt to retry message execution, which should fail
	//
	optsB, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetBInfo.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err = teleporter.TeleporterMessenger(subnetBInfo).RetryMessageExecution(
		optsB,
		subnetAInfo.BlockchainID,
		deliveredTeleporterMessage,
	)
	Expect(err).Should(Not(BeNil()))
	Expect(tx).Should(BeNil())
}
