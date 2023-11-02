package tests

import (
	"context"
	"math/big"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	examplecrosschainmessenger "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/ExampleMessenger/ExampleCrossChainMessenger"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

const (
	ExampleMessengerByteCodeFile = "./contracts/out/ExampleCrossChainMessenger.sol/ExampleCrossChainMessenger.json"
)

func ExampleMessengerGinkgo() {
	ExampleMessenger(&network.LocalNetwork{})
}

func ExampleMessenger(network network.Network) {
	var (
		teleporterMessageID *big.Int
	)

	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Deploy ExampleMessenger to Subnets A and B
	//
	ctx := context.Background()
	exampleMessengerABI, err := examplecrosschainmessenger.ExampleCrossChainMessengerMetaData.GetAbi()
	Expect(err).Should(BeNil())

	exampleMessengerContractA := utils.DeployContract(ctx, ExampleMessengerByteCodeFile, fundedKey, subnetAInfo, exampleMessengerABI, subnetAInfo.TeleporterRegistryAddress)
	exampleMessengerContractB := utils.DeployContract(ctx, ExampleMessengerByteCodeFile, fundedKey, subnetBInfo, exampleMessengerABI, subnetBInfo.TeleporterRegistryAddress)

	//
	// Call the example messenger contract on Subnet A
	//
	message := "Hello, world!"
	data, err := exampleMessengerABI.Pack("sendMessage", subnetBInfo.BlockchainID, exampleMessengerContractB, fundedAddress, big.NewInt(0), big.NewInt(300000), message)
	gasFeeCap, gasTipCap, nonce := utils.CalculateTxParams(ctx, subnetAInfo.ChainRPCClient, fundedAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   subnetAInfo.ChainIDInt,
		Nonce:     nonce,
		To:        &exampleMessengerContractA,
		Gas:       utils.DefaultTeleporterTransactionGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Data:      data,
	})

	signedTx := utils.SignTransaction(tx, fundedKey, subnetAInfo.ChainIDInt)
	receipt := utils.SendTransactionAndWaitForAcceptance(ctx, subnetAInfo.ChainWSClient, subnetAInfo.ChainRPCClient, signedTx, true)

	bind, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	event, err := utils.GetSendEventFromLogs(receipt.Logs, bind)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID = event.Message.MessageID

	//
	// Relay the message to the destination
	//

	network.RelayMessage(ctx, receipt.BlockHash, receipt.BlockNumber, subnetAInfo, subnetBInfo, true)

	//
	// Check Teleporter message received on the destination
	//
	data, err = teleportermessenger.PackMessageReceived(subnetAInfo.BlockchainID, teleporterMessageID)
	Expect(err).Should(BeNil())
	callMessage := interfaces.CallMsg{
		To:   &teleporterContractAddress,
		Data: data,
	}
	result, err := subnetBInfo.ChainRPCClient.CallContract(context.Background(), callMessage, nil)
	Expect(err).Should(BeNil())

	// check the contract call result
	delivered, err := teleportermessenger.UnpackMessageReceivedResult(result)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Verify we received the expected string
	//
	data, err = exampleMessengerABI.Pack("getCurrentMessage", subnetAInfo.BlockchainID)
	Expect(err).Should(BeNil())
	callMessage = interfaces.CallMsg{
		To:   &exampleMessengerContractB,
		Data: data,
	}
	result, err = subnetBInfo.ChainRPCClient.CallContract(context.Background(), callMessage, nil)
	Expect(err).Should(BeNil())

	var messageInfo struct {
		Sender  common.Address
		Message string
	}
	exampleMessengerABI.UnpackIntoInterface(&messageInfo, "getCurrentMessage", result)
	Expect(messageInfo.Message).Should(Equal(message))
}
