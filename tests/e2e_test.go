// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"os"
	"testing"

	localUtils "github.com/ava-labs/teleporter/tests/utils/local-network-utils"
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
	localUtils.SetupNetwork(warpGenesisFile)
	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err :=
		deploymentUtils.ConstructKeylessTransaction(teleporterByteCodeFile, false)
	Expect(err).Should(BeNil())

	_, fundedKey := localUtils.GetFundedAccountInfo()
	localUtils.DeployTeleporterContracts(
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)
	localUtils.DeployTeleporterRegistryContracts(teleporterContractAddress, fundedKey)
	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(localUtils.TearDownNetwork)

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Cross-chain application tests
	ginkgo.It("Example cross chain messenger", ExampleMessengerGinkgo)
	ginkgo.It("ERC20 bridge multihop", ERC20BridgeMultihopGinkgo)

	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B, and one from B to A", BasicSendReceiveGinkgo)
	ginkgo.It("Deliver to the wrong chain", DeliverToWrongChainGinkgo)
	ginkgo.It("Deliver to non-existent contract", DeliverToNonExistentContractGinkgo) // TODO: fix
	ginkgo.It("Retry successful execution", RetrySuccessfulExecutionGinkgo)
	ginkgo.It("Unallowed relayer", UnallowedRelayerGinkgo)
	ginkgo.It("Receive message twice", ReceiveMessageTwiceGinkgo)
	ginkgo.It("Add additional fee amount", AddFeeAmountGinkgo)
	ginkgo.It("Send specific receipts", SendSpecificReceiptsGinkgo)
	ginkgo.It("Insufficient gas", InsufficientGasGinkgo)
	ginkgo.It("Resubmit altered message", ResubmitAlteredMessageGinkgo)
	ginkgo.It("Relayer modifies message", RelayerModifiesMessageGinkgo)
	ginkgo.It("Validator churn", ValidatorChurnGinkgo)
	// Since the validator churn test modifies the network topology, we put it last for now.
	// It should not affect the other tests, but we get some errors if we run it before the other tests.
	// TODO: we should fix this so that the order of the tests does not matter.
})
