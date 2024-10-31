// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter_test

import (
	"context"
	"os"
	"testing"
	"time"

	teleporterFlows "github.com/ava-labs/teleporter/tests/flows/teleporter"
	registryFlows "github.com/ava-labs/teleporter/tests/flows/teleporter/registry"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
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
	validatorManagerLabel    = "ValidatorManager"
)

var (
	LocalNetworkInstance *localnetwork.LocalNetwork
	TeleporterInfo       utils.TeleporterTestInfo
)

func TestTeleporter(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
	// Generate the Teleporter deployment values
	teleporterDeployerTransaction,
		teleporterDeployedBytecode,
		teleporterDeployerAddress,
		teleporterContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		teleporterByteCodeFile,
		false,
		deploymentUtils.GetDefaultContractCreationGasPrice(),
	)
	Expect(err).Should(BeNil())

	// Create the local network instance
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	LocalNetworkInstance = localnetwork.NewLocalNetwork(
		ctx,
		"teleporter-test-local-network",
		warpGenesisTemplateFile,
		[]localnetwork.SubnetSpec{
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
	TeleporterInfo = utils.NewTeleporterTestInfo(LocalNetworkInstance.GetAllSubnetsInfo())
	log.Info("Started local network")

	// Only need to deploy Teleporter on the C-Chain since it is included in the genesis of the subnet chains.
	_, fundedKey := LocalNetworkInstance.GetFundedAccountInfo()
	TeleporterInfo.DeployTeleporterMessenger(
		ctx,
		LocalNetworkInstance.GetPrimaryNetworkInfo(),
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
	)

	for _, subnet := range LocalNetworkInstance.GetAllSubnetsInfo() {
		TeleporterInfo.SetTeleporter(teleporterContractAddress, subnet)
		TeleporterInfo.InitializeBlockchainID(subnet, fundedKey)
		TeleporterInfo.DeployTeleporterRegistry(subnet, fundedKey)
	}

	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
	LocalNetworkInstance = nil
})

var _ = ginkgo.Describe("[Teleporter integration tests]", func() {
	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B, and one from B to A",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.BasicSendReceive(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Deliver to the wrong chain",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.DeliverToWrongChain(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Deliver to non-existent contract",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.DeliverToNonExistentContract(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Retry successful execution",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.RetrySuccessfulExecution(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Unallowed relayer",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.UnallowedRelayer(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Relay message twice",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.RelayMessageTwice(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Add additional fee amount",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.AddFeeAmount(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Send specific receipts",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.SendSpecificReceipts(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Insufficient gas",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.InsufficientGas(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Resubmit altered message",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.ResubmitAlteredMessage(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Calculate Teleporter message IDs",
		ginkgo.Label(utilsLabel),
		func() {
			teleporterFlows.CalculateMessageID(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Relayer modifies message",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.RelayerModifiesMessage(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Validator churn",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			teleporterFlows.ValidatorChurn(LocalNetworkInstance, TeleporterInfo)
		})

	// Teleporter Registry tests
	ginkgo.It("Teleporter registry",
		ginkgo.Label(upgradabilityLabel),
		func() {
			registryFlows.TeleporterRegistry(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Check upgrade access",
		ginkgo.Label(upgradabilityLabel),
		func() {
			registryFlows.CheckUpgradeAccess(LocalNetworkInstance, TeleporterInfo)
		})
	ginkgo.It("Pause and Unpause Teleporter",
		ginkgo.Label(upgradabilityLabel),
		func() {
			registryFlows.PauseTeleporter(LocalNetworkInstance, TeleporterInfo)
		})
})
