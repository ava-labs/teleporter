package validator_manager_test

import (
	"context"
	"os"
	"testing"
	"time"

	validatorManagerFlows "github.com/ava-labs/teleporter/tests/flows/validator-manager"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"
	validatorManagerLabel   = "ValidatorManager"
)

var (
	LocalNetworkInstance *localnetwork.LocalNetwork
)

func TestValidatorManager(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Validator Manager e2e test")
}

// Define the before and after suite functions.
var _ = ginkgo.BeforeEach(func() {
	// Create the local network instance
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	LocalNetworkInstance = localnetwork.NewLocalNetwork(
		ctx,
		"validator-manager-test-local-network",
		warpGenesisTemplateFile,
		[]localnetwork.SubnetSpec{
			{
				Name:       "A",
				EVMChainID: 12345,
				NodeCount:  2,
			},
			{
				Name:       "B",
				EVMChainID: 54321,
				NodeCount:  2,
			},
		},
		2,
		2,
	)
	log.Info("Started local network")
})

var _ = ginkgo.AfterEach(func() {
	LocalNetworkInstance.TearDownNetwork()
	LocalNetworkInstance = nil
})

var _ = ginkgo.Describe("[Validator manager integration tests]", func() {
	// Validator Manager tests
	// ginkgo.It("Native token staking manager",
	// 	ginkgo.Label(validatorManagerLabel),
	// 	func() {
	// 		validatorManagerFlows.NativeTokenStakingManager(LocalNetworkInstance)
	// 	})
	// ginkgo.It("ERC20 token staking manager",
	// 	ginkgo.Label(validatorManagerLabel),
	// 	func() {
	// 		validatorManagerFlows.ERC20TokenStakingManager(LocalNetworkInstance)
	// 	})
	// ginkgo.It("PoA migration to PoS",
	// 	ginkgo.Label(validatorManagerLabel),
	// 	func() {
	// 		validatorManagerFlows.PoAMigrationToPoS(LocalNetworkInstance)
	// 	})
	ginkgo.It("Foo",
		func() {
			validatorManagerFlows.UpgradeFoo(LocalNetworkInstance)
		})
})
