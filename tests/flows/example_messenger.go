package flows

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	examplecrosschainmessenger "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ExampleMessenger/ExampleCrossChainMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func ExampleMessenger(network interfaces.Network) {
	subnetAInfo, subnetBInfo, _ := utils.GetThreeSubnets(network)
	_, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()

	_, subnetAExampleMessenger := utils.DeployExampleCrossChainMessenger(ctx, fundedKey, subnetAInfo)
	exampleMessengerContractB, subnetBExampleMessenger := utils.DeployExampleCrossChainMessenger(
		ctx, fundedKey, subnetBInfo,
	)

	//
	// Call the example messenger contract on Subnet A
	//
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := subnetAExampleMessenger.SendMessage(
		optsA,
		subnetBInfo.BlockchainID,
		exampleMessengerContractB,
		common.BigToAddress(common.Big0),
		big.NewInt(0),
		examplecrosschainmessenger.SendMessageRequiredGas,
		message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx)

	event, err := utils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID := event.MessageID

	//
	// Relay the message to the destination
	//
	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
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
