package main

import (
	"log"

	"github.com/ava-labs/teleporter/tests/flows"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/testnet"
	"github.com/onsi/gomega"
)

func runFlow(flowName string, flow func(interfaces.Network), network interfaces.Network) {
	log.Println("Running test", flowName)
	flow(network)
	log.Println("Finished running test", flowName)
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
	runFlow("BasicOneWaySend", flows.BasicSendReceive, network)
	// runFlow("DeliverToWrongChain", flows.DeliverToWrongChain, network)
	runFlow("DeliverToNonExistentContract", flows.DeliverToNonExistentContract, network)
	runFlow("RetrySuccessfulExecution", flows.RetrySuccessfulExecution, network)
	// runFlow("UnallowedRelayer", flows.UnallowedRelayer, network)
	runFlow("ReceiveMessageTwice", flows.ReceiveMessageTwice, network)
	runFlow("AddFeeAmount", flows.AddFeeAmount, network)
	runFlow("InsufficientGas", flows.InsufficientGas, network)
	runFlow("ResubmitAlteredMessage", flows.ResubmitAlteredMessage, network)
	// runFlow("RelayerModifiesMessage", flows.RelayerModifiesMessage, network)
	log.Println("Finished Teleporter test flows")

	// Run the cross-chain application test flows.
	runFlow("ExampleMessenger", flows.ExampleMessenger, network)
	log.Println("Finished cross-chain application test flows")
}
