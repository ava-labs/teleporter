package main

// import (
// 	"github.com/ethereum/go-ethereum/log"

// 	"github.com/ava-labs/teleporter/tests/flows/teleporter"
// 	"github.com/ava-labs/teleporter/tests/flows/teleporter/registry"

// 	"github.com/ava-labs/teleporter/tests/interfaces"
// 	"github.com/ava-labs/teleporter/tests/testnet"
// 	"github.com/onsi/gomega"
// )

// func runFlow(flowName string, flow func(interfaces.Network), network interfaces.Network) {
// 	log.Info("Running test", "flowName", flowName)
// 	flow(network)
// 	log.Info("Finished running test", "flowName", flowName)
// }

// func main() {
// 	// Register a failure handler that panics
// 	gomega.RegisterFailHandler(func(message string, callerSkip ...int) {
// 		panic(message)
// 	})

// 	// Create the new network instances
// 	network, err := testnet.NewTestNetwork()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Run the Teleporter test flows.
// 	// Note that the following flows are not able to run against testnet because they
// 	// require the tests to be able to relay messages independently or modifying validator sets.
// 	//   - RelayerModifiesMessage
// 	//   - UnallowedRelayer
// 	//   - ValidatorSetChrun
// 	//   - TeleporterRegistry
// 	runFlow("AddFeeAmount", teleporter.AddFeeAmount, network)
// 	runFlow("BasicSendRecevie", teleporter.BasicSendReceive, network)
// 	runFlow("DeliverToNonExistentContract", teleporter.DeliverToNonExistentContract, network)
// 	runFlow("DeliverToWrongChain", teleporter.DeliverToWrongChain, network)
// 	runFlow("InsufficientGas", teleporter.InsufficientGas, network)
// 	runFlow("RelayMessageTwice", teleporter.RelayMessageTwice, network)
// 	runFlow("ResubmitAlteredMessage", teleporter.ResubmitAlteredMessage, network)
// 	runFlow("RetrySuccessfulExecution", teleporter.RetrySuccessfulExecution, network)
// 	runFlow("SendSpecificReceipts", teleporter.SendSpecificReceipts, network)
// 	log.Info("Finished Teleporter test flows")

// 	// Run the upgradability test flows
// 	runFlow("CheckUpgradeAccess", registry.CheckUpgradeAccess, network)
// 	runFlow("PauseTeleporter", registry.PauseTeleporter, network)
// 	log.Info("Finished upgradability test flows")
// }
