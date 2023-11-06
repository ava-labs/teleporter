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
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

func DeliverToNonExistentContractGinkgo() {
	DeliverToNonExistentContract(&network.LocalNetwork{})
}

func DeliverToNonExistentContract(network network.Network) {
	var (
		teleporterMessageID *big.Int
	)

	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	deployerKey, _ := crypto.GenerateKey()
	deployerAddress := crypto.PubkeyToAddress(deployerKey.PublicKey)

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())

	//
	// Fund the deployer address on Subnet B
	//
	ctx := context.Background()

	fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
	fundDeployerTx := utils.CreateNativeTransferTransaction(ctx, subnetBInfo, fundedAddress, fundedKey, deployerAddress, fundAmount)
	utils.SendTransactionAndWaitForAcceptance(ctx, subnetBInfo.ChainWSClient, subnetBInfo.ChainRPCClient, fundDeployerTx, true)

	//
	// Deploy ExampleMessenger to Subnet A, but not to Subnet B
	// Send a message that should fail to be executed on Subnet B
	//

	exampleMessengerABI, err := examplecrosschainmessenger.ExampleCrossChainMessengerMetaData.GetAbi()
	Expect(err).Should(BeNil())

	exampleMessengerContractA := utils.DeployContract(ctx, ExampleMessengerByteCodeFile, fundedKey, subnetAInfo, exampleMessengerABI, subnetAInfo.TeleporterRegistryAddress)

	// Derive the eventual address of the destination contract on Subnet B
	nonce, err := subnetBInfo.ChainRPCClient.NonceAt(ctx, deployerAddress, nil)
	Expect(err).Should(BeNil())
	destinationContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, nonce)
	Expect(err).Should(BeNil())

	//
	// Call the example messenger contract on Subnet A
	//
	message := "Hello, world!"
	data, err := exampleMessengerABI.Pack("sendMessage", subnetBInfo.BlockchainID, destinationContractAddress, fundedAddress, big.NewInt(0), big.NewInt(300000), message)
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

	sendEvent, err := utils.GetSendEventFromLogs(receipt.Logs, subnetATeleporterMessenger)
	Expect(err).Should(BeNil())
	Expect(sendEvent.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID = sendEvent.Message.MessageID

	//
	// Relay the message to the destination
	//

	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)
	receiveEvent, err := utils.GetReceiveEventFromLogs(receipt.Logs, subnetATeleporterMessenger)
	Expect(err).Should(BeNil())
	deliveredTeleporterMessage := receiveEvent.Message

	//
	// Check that the message was successfully relayed
	//
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Check that the message was not successfully executed
	//
	executionFailedEvent, err := utils.GetMessageExecutionFailedFromLogs(receipt.Logs, subnetBTeleporterMessenger)
	Expect(err).Should(BeNil())
	Expect(executionFailedEvent.Message.MessageID).Should(Equal(deliveredTeleporterMessage.MessageID))

	//
	// Deploy the contract on Subnet B
	//
	exampleMessengerContractB := utils.DeployContract(ctx, ExampleMessengerByteCodeFile, deployerKey, subnetBInfo, exampleMessengerABI, subnetBInfo.TeleporterRegistryAddress)

	// Confirm that it was deployed at the expected address
	Expect(exampleMessengerContractB).Should(Equal(destinationContractAddress))

	//
	// Call retryMessageExecution on Subnet B
	//
	signedTx = utils.CreateRetryMessageExecutionTransaction(
		ctx,
		subnetBInfo,
		subnetAInfo.BlockchainID,
		deliveredTeleporterMessage,
		fundedAddress,
		fundedKey,
		teleporterContractAddress,
	)
	utils.SendTransactionAndWaitForAcceptance(ctx, subnetBInfo.ChainWSClient, subnetBInfo.ChainRPCClient, signedTx, true)

	//
	// Verify we received the expected string
	//
	subnetBExampleMessenger, err := examplecrosschainmessenger.NewExampleCrossChainMessenger(exampleMessengerContractB, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	res, err := subnetBExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, subnetAInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(res.Message).Should(Equal(message))
}
