package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	examplecrosschainmessenger "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/ExampleMessenger/ExampleCrossChainMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	localUtils "github.com/ava-labs/teleporter/tests/utils/local-network-utils"
	. "github.com/onsi/gomega"
)

func RetrySuccessfulExecution(network network.Network) {
	subnets := network.GetSubnetsInfo()
	Expect(len(subnets)).Should(BeNumerically(">=", 2))
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()

	_, subnetAExampleMessenger := localUtils.DeployExampleCrossChainMessenger(ctx, fundedKey, subnetAInfo)
	exampleMessengerContractAddressB, subnetBExampleMessenger := localUtils.DeployExampleCrossChainMessenger(
		ctx, fundedKey, subnetBInfo,
	)

	//
	// Call the example messenger contract on Subnet A
	//
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.ChainIDInt)
	Expect(err).Should(BeNil())
	tx, err := subnetAExampleMessenger.SendMessage(
		optsA,
		subnetBInfo.BlockchainID,
		exampleMessengerContractAddressB,
		fundedAddress,
		big.NewInt(0),
		examplecrosschainmessenger.SendMessageRequiredGas,
		message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	event, err := utils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID := event.Message.MessageID

	//
	// Relay the message to the destination
	//
	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)
	receiveEvent, err :=
		utils.GetEventFromLogs(receipt.Logs, subnetBInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	deliveredTeleporterMessage := receiveEvent.Message

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Verify we received the expected string
	//
	_, currMessage, err := subnetBExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, subnetAInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))

	//
	// Attempt to retry message execution, which should fail
	//
	optsB, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetBInfo.ChainIDInt)
	Expect(err).Should(BeNil())
	tx, err =
		subnetBInfo.TeleporterMessenger.RetryMessageExecution(optsB, subnetAInfo.BlockchainID, deliveredTeleporterMessage)
	Expect(err).Should(Not(BeNil()))
	Expect(tx).Should(BeNil())
}