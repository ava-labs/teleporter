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

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
	setUpNetwork(warpGenesisFile)
	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err := deploymentUtils.ConstructKeylessTransaction(teleporterByteCodeFile, false)
	Expect(err).Should(BeNil())

	deployTeleporterContracts(teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress)
	deployTeleporterRegistryContracts(teleporterContractAddress)
	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(tearDownNetwork)

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B", func() { flows.BasicOneWaySend(&LocalNetwork{}) })
	ginkgo.It("Deliver to the wrong chain", func() { flows.DeliverToWrongChain(&LocalNetwork{}) })
	ginkgo.It("Deliver to non-existent contract", func() { flows.DeliverToNonExistentContract(&LocalNetwork{}) })
	ginkgo.It("Retry successful execution", func() { flows.RetrySuccessfulExecution(&LocalNetwork{}) })
	ginkgo.It("Unallowed relayer", func() { flows.UnallowedRelayer(&LocalNetwork{}) })
	ginkgo.It("Receive message twice", func() { flows.ReceiveMessageTwice(&LocalNetwork{}) })
	ginkgo.It("Add additional fee amount", func() { flows.AddFeeAmount(&LocalNetwork{}) })
	ginkgo.It("Send specific receipts", func() { flows.SendSpecificReceipts(&LocalNetwork{}) })
	ginkgo.It("Insufficient gas", func() { flows.InsufficientGas(&LocalNetwork{}) })

	// Cross-chain application tests
	ginkgo.It("Example cross chain messenger", func() { flows.ExampleMessenger(&LocalNetwork{}) })
})
