package main

import (
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

	// Run the test, composing it with the Network implementation
	flows.BasicOneWaySend(network)
}
