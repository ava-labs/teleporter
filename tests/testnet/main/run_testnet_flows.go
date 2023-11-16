package main

import (
	"fmt"

	"github.com/ava-labs/teleporter/tests/flows"
	"github.com/ava-labs/teleporter/tests/testnet"
	"github.com/onsi/gomega"
)

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
	flows.BasicOneWaySend(network)
	fmt.Println("Finished running test BasicOneWaySend")

	flows.DeliverToWrongChain(network)
	fmt.Println("Finished running test DeliverToWrongChain")

	flows.DeliverToNonExistentContract(network)
	fmt.Println("Finished running test DeliverToNonExistentContract")

	flows.RetrySuccessfulExecution(network)
	fmt.Println("Finished running test RetrySuccessfulExecution")

	flows.UnallowedRelayer(network)
	fmt.Println("Finished running test UnallowedRelayer")

	flows.ReceiveMessageTwice(network)
	fmt.Println("Finished running test ReceiveMessageTwice")

	flows.AddFeeAmount(network)
	fmt.Println("Finished running test AddFeeAmount")

	flows.InsufficientGas(network)
	fmt.Println("Finished running test InsufficientGas")

	flows.ResubmitAlteredMessage(network)
	fmt.Println("Finished running test ResubmitAlteredMessage")

	flows.RelayerModifiesMessage(network)
	fmt.Println("Finished running test RelayerModifiesMessage")

	fmt.Println("Finished Teleporter test flows.")

	// Run the cross-chain application test flows.
	flows.ExampleMessenger(network)
	fmt.Println("Finished running test ExampleMessenger")

	fmt.Println("Finished cross-chain application test flows.")

}
