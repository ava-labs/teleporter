package main

import (
	"github.com/ethereum/go-ethereum/log"

	"github.com/ava-labs/teleporter/tests/flows"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/testnet"
	"github.com/onsi/gomega"
)

func runFlow(flowName string, flow func(interfaces.Network), network interfaces.Network) {
	log.Info("Running test", "flowName", flowName)
	flow(network)
	log.Info("Finished running test", "flowName", flowName)
}

func main() {
	// Register a failure handler that panics
	gomega.RegisterFailHandler(func(message string, callerSkip ...int) {
		panic(message)
	})

	// Create the new network instances
	network, err := testnet.NewTestNetwork()
	if err != nil {
		panic(err)
	}

	// Run the Teleporter test flows.
	// Note that the following flows are not able to run against testnet because they
	// require the tests to be able to relay messages independently or modifying validator sets.
	//   - RelayerModifiesMessage
	//   - UnallowedRelayer
	//   - ValidatorSetChrun
	//   - TeleporterRegistry
	runFlow("AddFeeAmount", flows.AddFeeAmount, network)
	runFlow("BasicSendRecevie", flows.BasicSendReceive, network)
	runFlow("DeliverToNonExistentContract", flows.DeliverToNonExistentContract, network)
	runFlow("DeliverToWrongChain", flows.DeliverToWrongChain, network)
	runFlow("InsufficientGas", flows.InsufficientGas, network)
	runFlow("RelayMessageTwice", flows.RelayMessageTwice, network)
	runFlow("ResubmitAlteredMessage", flows.ResubmitAlteredMessage, network)
	runFlow("RetrySuccessfulExecution", flows.RetrySuccessfulExecution, network)
	runFlow("SendSpecificReceipts", flows.SendSpecificReceipts, network)
	log.Info("Finished Teleporter test flows")

	// Run the cross-chain application test flows.
	runFlow("ExampleMessenger", flows.ExampleMessenger, network)
	runFlow("ERC20BridgeMutlihop", flows.ERC20BridgeMultihop, network)
	log.Info("Finished cross-chain application test flows")

	// Run the upgradeability test flows
	runFlow("CheckUpgradeAccess", flows.CheckUpgradeAccess, network)
	runFlow("PauseTeleporter", flows.PauseTeleporter, network)
	log.Info("Finished upgradeability test flows")
}
