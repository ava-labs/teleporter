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

var (
	LocalNetworkInstance *LocalNetwork
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
	// Create the local network instance
	LocalNetworkInstance = NewLocalNetwork(warpGenesisFile)

	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err :=
		deploymentUtils.ConstructKeylessTransaction(
			teleporterByteCodeFile,
			false,
			deploymentUtils.GetDefaultContractCreationGasPrice(),
		)
	Expect(err).Should(BeNil())

	_, fundedKey := LocalNetworkInstance.GetFundedAccountInfo()
	LocalNetworkInstance.DeployTeleporterContracts(
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
		true,
	)

	LocalNetworkInstance.DeployTeleporterRegistryContracts(teleporterContractAddress, fundedKey)
	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
})

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Cross-chain application tests
	ginkgo.It("Send native tokens from subnet A to B and back",
		ginkgo.Label("cross chain apps"),
		func() {
			flows.NativeTokenBridge(LocalNetworkInstance)
		})
	ginkgo.It("Send ERC20 tokens from subnet A to Native tokens on subnet B and back",
		ginkgo.Label("cross chain apps"),
		func() {
			flows.ERC20ToNativeTokenBridge(LocalNetworkInstance)
		})
	ginkgo.It("Example cross chain messenger",
		ginkgo.Label("cross chain apps"),
		func() {
			flows.ExampleMessenger(LocalNetworkInstance)
		})
	ginkgo.It("ERC20 bridge multihop",
		ginkgo.Label("cross chain apps"),
		func() {
			flows.ERC20BridgeMultihop(LocalNetworkInstance)
		})
	ginkgo.It("Block hash publish and receive",
		ginkgo.Label("cross chain apps"),
		func() {
			flows.BlockHashPublishReceive(LocalNetworkInstance)
		})

	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B, and one from B to A",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.BasicSendReceive(LocalNetworkInstance)
		})
	ginkgo.It("Deliver to the wrong chain",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.DeliverToWrongChain(LocalNetworkInstance)
		})
	ginkgo.It("Deliver to non-existent contract",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.DeliverToNonExistentContract(LocalNetworkInstance)
		})
	ginkgo.It("Retry successful execution",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.RetrySuccessfulExecution(LocalNetworkInstance)
		})
	ginkgo.It("Unallowed relayer",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.UnallowedRelayer(LocalNetworkInstance)
		})
	ginkgo.It("Relay message twice",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.RelayMessageTwice(LocalNetworkInstance)
		})
	ginkgo.It("Add additional fee amount",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.AddFeeAmount(LocalNetworkInstance)
		})
	ginkgo.It("Send specific receipts",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.SendSpecificReceipts(LocalNetworkInstance)
		})
	ginkgo.It("Insufficient gas",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.InsufficientGas(LocalNetworkInstance)
		})
	ginkgo.It("Resubmit altered message",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.ResubmitAlteredMessage(LocalNetworkInstance)
		})
	ginkgo.It("Check upgrade access",
		ginkgo.Label("upgradeability"),
		func() {
			flows.CheckUpgradeAccess(LocalNetworkInstance)
		})
	ginkgo.It("Pause and Unpause Teleporter",
		ginkgo.Label("upgradeability"),
		func() {
			flows.PauseTeleporter(LocalNetworkInstance)
		})
	ginkgo.It("Calculate Teleporter message IDs",
		ginkgo.Label("utils"),
		func() {
			flows.CalculateMessageID(LocalNetworkInstance)
		})

	// The following tests require special behavior by the relayer, so we only run them on a local network
	ginkgo.It("Relayer modifies message",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.RelayerModifiesMessage(LocalNetworkInstance)
		})
	ginkgo.It("Teleporter registry",
		ginkgo.Label("upgradeability"),
		func() {
			flows.TeleporterRegistry(LocalNetworkInstance)
		})
	ginkgo.It("Validator churn",
		ginkgo.Label("TeleporterMessenger"),
		func() {
			flows.ValidatorChurn(LocalNetworkInstance)
		})
	// Since the validator churn test modifies the network topology, we put it last for now.
	// It should not affect the other tests, but we get some errors if we run it before the other tests.
	// TODO: we should fix this so that the order of the tests does not matter.
})
