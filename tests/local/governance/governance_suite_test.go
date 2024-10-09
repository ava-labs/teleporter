package governance_test

import (
	"context"
	"os"
	"testing"
	"time"

	governanceFlows "github.com/ava-labs/teleporter/tests/flows/governance"
	"github.com/ava-labs/teleporter/tests/local"
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
	LocalNetworkInstance *local.LocalNetwork
)

func TestGovernance(t *testing.T) {
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
	LocalNetworkInstance = local.NewLocalNetwork(
		ctx,
		"teleporter-test-local-network",
		warpGenesisTemplateFile,
		[]local.SubnetSpec{
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
	LocalNetworkInstance.InitializeBlockchainIDOnAllChains(fundedKey)

	// Deploy the Teleporter registry contracts to all subnets and the C-Chain.
	LocalNetworkInstance.DeployTeleporterRegistryContracts(teleporterContractAddress, fundedKey)

	ginkgo.AddReportEntry(
		"network directory with node logs & configs; useful in the case of failures",
		LocalNetworkInstance.TmpNet().Dir,
		ginkgo.ReportEntryVisibilityFailureOrVerbose,
	)

	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	LocalNetworkInstance.TearDownNetwork()
})

var _ = ginkgo.Describe("[Governance integration tests]", func() {
	// Governance tests
	ginkgo.It("Deliver ValidatorSetSig signed message",
		ginkgo.Label(validatorSetSigLabel),
		func() {
			governanceFlows.ValidatorSetSig(LocalNetworkInstance)
		})
})
