// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"os"
	"testing"

	"github.com/ava-labs/teleporter/tests/network"
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
	ginkgo.It("Send native tokens from subnet A to B and back", func() {
		NativeTokenBridge(&network.LocalNetwork{})
	})
	ginkgo.It("Send ERC20 tokens from subnet A to Native tokens on subnet B and back", func() {
		ERC20ToNativeTokenBridge(&network.LocalNetwork{})
	})

	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B, and one from B to A", func() {
		BasicSendReceive(&network.LocalNetwork{})
	})
	ginkgo.It("Deliver to the wrong chain", func() {
		DeliverToWrongChain(&network.LocalNetwork{})
	})
	ginkgo.It("Deliver to non-existent contract", func() {
		DeliverToNonExistentContract(&network.LocalNetwork{})
	})
	ginkgo.It("Retry successful execution", func() {
		RetrySuccessfulExecution(&network.LocalNetwork{})
	})
	ginkgo.It("Unallowed relayer", func() {
		UnallowedRelayer(&network.LocalNetwork{})
	})
	ginkgo.It("Receive message twice", func() {
		ReceiveMessageTwice(&network.LocalNetwork{})
	})
	ginkgo.It("Add additional fee amount", func() {
		AddFeeAmount(&network.LocalNetwork{})
	})
	ginkgo.It("Send specific receipts", func() {
		SendSpecificReceipts(&network.LocalNetwork{})
	})
	ginkgo.It("Insufficient gas", func() {
		InsufficientGas(&network.LocalNetwork{})
	})
	ginkgo.It("Resubmit altered message", func() {
		ResubmitAlteredMessage(&network.LocalNetwork{})
	})

	// The following tests require special behavior by the relayer, so we only run them on a local network
	ginkgo.It("Relayer modifies message", RelayerModifiesMessage)
	ginkgo.It("Validator churn", ValidatorChurn)
	// Since the validator churn test modifies the network topology, we put it last for now.
	// It should not affect the other tests, but we get some errors if we run it before the other tests.
	// TODO: we should fix this so that the order of the tests does not matter.
})
