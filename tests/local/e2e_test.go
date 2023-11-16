// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package local

import (
	"os"
	"testing"

	"github.com/ava-labs/teleporter/tests/flows"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile = "./contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisFile        = "./tests/utils/warp-genesis.json"
)

func TestE2E(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

var localNetworkInstance *localNetwork

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
	localNetworkInstance = newLocalNetwork(warpGenesisFile)
	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err := deploymentUtils.ConstructKeylessTransaction(teleporterByteCodeFile, false)
	Expect(err).Should(BeNil())

	localNetworkInstance.deployTeleporterContracts(teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress)
	localNetworkInstance.deployTeleporterRegistryContracts(teleporterContractAddress)
	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(localNetworkInstance.tearDownNetwork)

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B", func() { flows.BasicOneWaySend(localNetworkInstance) })
	ginkgo.It("Deliver to the wrong chain", func() { flows.DeliverToWrongChain(localNetworkInstance) })
	ginkgo.It("Deliver to non-existent contract", func() { flows.DeliverToNonExistentContract(localNetworkInstance) })
	ginkgo.It("Retry successful execution", func() { flows.RetrySuccessfulExecution(localNetworkInstance) })
	ginkgo.It("Unallowed relayer", func() { flows.UnallowedRelayer(localNetworkInstance) })
	ginkgo.It("Receive message twice", func() { flows.ReceiveMessageTwice(localNetworkInstance) })
	ginkgo.It("Add additional fee amount", func() { flows.AddFeeAmount(localNetworkInstance) })
	ginkgo.It("Send specific receipts", func() { flows.SendSpecificReceipts(localNetworkInstance) })
	ginkgo.It("Insufficient gas", func() { flows.InsufficientGas(localNetworkInstance) })
	ginkgo.It("Resubmit altered message", func() { flows.ResubmitAlteredMessage(localNetworkInstance) })
	ginkgo.It("Relayer modifies message", func() { flows.RelayerModifiesMessage(localNetworkInstance) })
	ginkgo.It("Validator churn", func() {
		flows.ValidatorChurn(localNetworkInstance, constructSignedWarpMessageBytes, localNetworkInstance.addSubnetAValidators)
	})

	// Cross-chain application tests
	ginkgo.It("Example cross chain messenger", func() { flows.ExampleMessenger(localNetworkInstance) })
})
