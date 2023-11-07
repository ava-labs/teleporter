package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	examplecrosschainmessenger "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/ExampleMessenger/ExampleCrossChainMessenger"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
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
	testAddress, testKey := network.GetTestAccountInfo()
	ctx := context.Background()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())

	// Deploy ExampleMessenger to Subnets A
	optsA := utils.CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)
	_, tx, subnetAExampleMessenger, err := examplecrosschainmessenger.DeployExampleCrossChainMessenger(optsA, subnetAInfo.ChainRPCClient, subnetAInfo.TeleporterRegistryAddress)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	// Deploy ExampleMessenger to Subnets B
	optsB := utils.CreateTransactorOpts(ctx, subnetBInfo, fundedAddress, fundedKey)
	exampleMessengerContractB, tx, _, err := examplecrosschainmessenger.DeployExampleCrossChainMessenger(optsB, subnetBInfo.ChainRPCClient, subnetBInfo.TeleporterRegistryAddress)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err = bind.WaitMined(ctx, subnetBInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	message := "Hello, world!"
	optsA = utils.CreateTransactorOpts(ctx, subnetAInfo, testAddress, testKey)
	tx, err = subnetAExampleMessenger.SendMessage(optsA, subnetBInfo.BlockchainID, exampleMessengerContractB, fundedAddress, big.NewInt(0), big.NewInt(17000000), message)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err = bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	event, err := utils.GetSendEventFromLogs(receipt.Logs, subnetATeleporterMessenger)
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
	failedMessageExecutionEvent, err := utils.GetMessageExecutionFailedEventFromLogs(receipt.Logs, subnetATeleporterMessenger)
	Expect(err).Should(BeNil())
	Expect(failedMessageExecutionEvent.MessageID).Should(Equal(messageID))
	Expect(failedMessageExecutionEvent.OriginChainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))

	// // Add funds to test address and retry message execution
	// fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
	// transferTxn := utils.CreateNativeTransferTransaction(ctx, subnetBInfo, fundedAddress, fundedKey, testAddress, fundAmount)
	// utils.SendTransactionAndWaitForAcceptance(ctx, subnetBInfo.ChainWSClient, subnetBInfo.ChainRPCClient, transferTxn, true)

	// // Retry message execution
	// receipt = utils.RetryMessageExecution(
	// 	ctx, subnetAInfo.BlockchainID, subnetBInfo, failedMessageExecutionEvent.Message, fundedAddress, fundedKey, subnetBTeleporterMessenger)
	// executedEvent, err := utils.GetMessageExecutedEventFromLogs(receipt.Logs, subnetBTeleporterMessenger)
	// Expect(err).Should(BeNil())
	// Expect(executedEvent.MessageID).Should(Equal(messageID))
	// Expect(executedEvent.OriginChainID[:]).Should(Equal(subnetAInfo.BlockchainID[:]))
}
