// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package local

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/teleporter/tests/flows"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile  = "./out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"

	teleporterMessengerLabel = "TeleporterMessenger"
	upgradabilityLabel       = "upgradability"
	utilsLabel               = "utils"
	validatorSetSigLabel     = "ValidatorSetSig"
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
	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployedBytecode, teleporterDeployerAddress, teleporterContractAddress, err :=
		deploymentUtils.ConstructKeylessTransaction(
			teleporterByteCodeFile,
			false,
			deploymentUtils.GetDefaultContractCreationGasPrice(),
		)
	Expect(err).Should(BeNil())

	// Create the local network instance
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	LocalNetworkInstance = NewLocalNetwork(
		ctx,
		"teleporter-test-local-network",
		warpGenesisTemplateFile,
		[]SubnetSpec{
			{
				Name:                       "A",
				EVMChainID:                 12345,
				TeleporterContractAddress:  teleporterContractAddress,
				TeleporterDeployedBytecode: teleporterDeployedBytecode,
				TeleporterDeployerAddress:  teleporterDeployerAddress,
				NodeCount:                  2,
			},
			{
				Name:                       "B",
				EVMChainID:                 54321,
				TeleporterContractAddress:  teleporterContractAddress,
				TeleporterDeployedBytecode: teleporterDeployedBytecode,
				TeleporterDeployerAddress:  teleporterDeployerAddress,
				NodeCount:                  2,
			},
		},
		2,
	)
	log.Info("Started local network")

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the subnet chains.
	_, fundedKey := LocalNetworkInstance.GetFundedAccountInfo()
	LocalNetworkInstance.DeployTeleporterContractToCChain(
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)
	LocalNetworkInstance.SetTeleporterContractAddress(teleporterContractAddress)

	// Deploy the Teleporter registry contracts to all subnets and the C-Chain.
	LocalNetworkInstance.DeployTeleporterRegistryContracts(teleporterContractAddress, fundedKey)

	ginkgo.AddReportEntry(
		"network directory with node logs & configs; useful in the case of failures",
		LocalNetworkInstance.tmpnet.Dir,
		ginkgo.ReportEntryVisibilityFailureOrVerbose,
	)

	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
})

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B, and one from B to A",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.BasicSendReceive(LocalNetworkInstance)
		})
	ginkgo.It("Deliver to the wrong chain",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.DeliverToWrongChain(LocalNetworkInstance)
		})
	ginkgo.It("Deliver to non-existent contract",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.DeliverToNonExistentContract(LocalNetworkInstance)
		})
	ginkgo.It("Retry successful execution",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.RetrySuccessfulExecution(LocalNetworkInstance)
		})
	ginkgo.It("Unallowed relayer",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.UnallowedRelayer(LocalNetworkInstance)
		})
	ginkgo.It("Relay message twice",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.RelayMessageTwice(LocalNetworkInstance)
		})
	ginkgo.It("Add additional fee amount",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.AddFeeAmount(LocalNetworkInstance)
		})
	ginkgo.It("Send specific receipts",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.SendSpecificReceipts(LocalNetworkInstance)
		})
	ginkgo.It("Insufficient gas",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.InsufficientGas(LocalNetworkInstance)
		})
	ginkgo.It("Resubmit altered message",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.ResubmitAlteredMessage(LocalNetworkInstance)
		})
	ginkgo.It("Check upgrade access",
		ginkgo.Label(upgradabilityLabel),
		func() {
			flows.CheckUpgradeAccess(LocalNetworkInstance)
		})
	ginkgo.It("Pause and Unpause Teleporter",
		ginkgo.Label(upgradabilityLabel),
		func() {
			flows.PauseTeleporter(LocalNetworkInstance)
		})
	ginkgo.It("Calculate Teleporter message IDs",
		ginkgo.Label(utilsLabel),
		func() {
			flows.CalculateMessageID(LocalNetworkInstance)
		})

	// The following tests require special behavior by the relayer, so we only run them on a local network
	ginkgo.It("Relayer modifies message",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.RelayerModifiesMessage(LocalNetworkInstance)
		})
	ginkgo.It("Teleporter registry",
		ginkgo.Label(upgradabilityLabel),
		func() {
			flows.TeleporterRegistry(LocalNetworkInstance)
		})
	ginkgo.It("Deliver ValidatorSetSig signed message",
		ginkgo.Label(validatorSetSigLabel),
		func() {
			flows.ValidatorSetSig(LocalNetworkInstance)
		})
	ginkgo.It("Validator churn",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.ValidatorChurn(LocalNetworkInstance)
		})
	// Since the validator churn test modifies the network topology, we put it last for now.
	// It should not affect the other tests, but we get some errors if we run it before the other tests.
	// TODO: we should fix this so that the order of the tests does not matter.
})
