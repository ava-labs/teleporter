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
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

func DeliverToNonExistentContractGinkgo() {
	DeliverToNonExistentContract(&network.LocalNetwork{})
}

func DeliverToNonExistentContract(network network.Network) {
	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	deployerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	deployerAddress := crypto.PubkeyToAddress(deployerKey.PublicKey)

	//
	// Fund the deployer address on Subnet B
	//
	ctx := context.Background()

	fundAmount := big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(10)) // 10eth
	fundDeployerTx := utils.CreateNativeTransferTransaction(
		ctx, subnetBInfo, fundedAddress, fundedKey, deployerAddress, fundAmount,
	)
	utils.SendTransactionAndWaitForAcceptance(ctx, subnetBInfo, fundDeployerTx, true)

	//
	// Deploy ExampleMessenger to Subnet A, but not to Subnet B
	// Send a message that should fail to be executed on Subnet B
	//
	_, subnetAExampleMessenger := localUtils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetAInfo)

	// Derive the eventual address of the destination contract on Subnet B
	nonce, err := subnetBInfo.ChainRPCClient.NonceAt(ctx, deployerAddress, nil)
	Expect(err).Should(BeNil())
	destinationContractAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, nonce)
	Expect(err).Should(BeNil())

	//
	// Call the example messenger contract on Subnet A
	//
	message := "Hello, world!"
	optsA, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnetAInfo.ChainIDInt)
	Expect(err).Should(BeNil())
	tx, err := subnetAExampleMessenger.SendMessage(
		optsA,
		subnetBInfo.BlockchainID,
		destinationContractAddress,
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

	sendEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(sendEvent.DestinationChainID[:]).Should(Equal(subnetBInfo.BlockchainID[:]))

	teleporterMessageID := sendEvent.Message.MessageID

	//
	// Relay the message to the destination
	//

	receipt = network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, true)
	receiveEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseReceiveCrossChainMessage)
	Expect(err).Should(BeNil())
	deliveredTeleporterMessage := receiveEvent.Message

	//
	// Check that the message was successfully relayed
	//
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{},
		subnetAInfo.BlockchainID,
		teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	//
	// Check that the message was not successfully executed
	//
	executionFailedEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		subnetBInfo.TeleporterMessenger.ParseMessageExecutionFailed,
	)
	Expect(err).Should(BeNil())
	Expect(executionFailedEvent.Message.MessageID).Should(Equal(deliveredTeleporterMessage.MessageID))

	//
	// Deploy the contract on Subnet B
	//
	exampleMessengerContractB, subnetBExampleMessenger :=
		localUtils.DeployExampleCrossChainMessenger(ctx, fundedAddress, fundedKey, subnetBInfo)

	// Wait for the transaction to be mined
	_, err = bind.WaitMined(ctx, subnetBInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	// Confirm that it was deployed at the expected address
	Expect(exampleMessengerContractB).Should(Equal(destinationContractAddress))

	//
	// Call retryMessageExecution on Subnet B
	//

	receipt = utils.RetryMessageExecutionAndWaitForAcceptance(
		ctx,
		subnetAInfo.BlockchainID,
		subnetBInfo,
		deliveredTeleporterMessage,
		fundedAddress,
		fundedKey,
		subnetBInfo.TeleporterMessenger,
	)

	//
	// Verify we received the expected string
	//
	_, currMessage, err := subnetBExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, subnetAInfo.BlockchainID)
	Expect(err).Should(BeNil())
	Expect(currMessage).Should(Equal(message))
}
