package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	. "github.com/onsi/gomega"
)

func ExampleMessenger(network interfaces.Network) {
	var (
		teleporterMessageID *big.Int
	)

	subnetAInfo, subnetBInfo := network.GetSubnetInfo()
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		subnetBInfo.RPCClient)
	Expect(err).Should(BeNil())

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()

	_, subnetAExampleMessenger := utils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetAInfo)
	exampleMessengerContractB, subnetBExampleMessenger := utils.DeployExampleCrossChainMessenger(
		ctx,
		fundedAddress,
		fundedKey,
		subnetBInfo)

	//
	// Call the example messenger contract on Subnet A
	//
	message := "Hello, world!"
	optsA := utils.CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)
	tx, err := subnetAExampleMessenger.SendMessage(
		optsA,
		subnetBInfo.BlockchainID,
		exampleMessengerContractB,
		fundedAddress,
		big.NewInt(0),
		big.NewInt(300000),
		message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, subnetAInfo.RPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	event, err := utils.GetEventFromLogs(receipt.Logs, subnetATeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID = event.Message.MessageID

	//
	// Relay the message to the destination
	//
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, true)

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := subnetBTeleporterMessenger.MessageReceived(
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
}
