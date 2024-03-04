// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package local

import (
	"encoding/hex"
	"os"
	"testing"

	testUtils "github.com/ava-labs/teleporter-token-bridge/tests"
	"github.com/ava-labs/teleporter/tests/flows"
	"github.com/ava-labs/teleporter/tests/local"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile = "./contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisFile        = "./tests/utils/warp-genesis.json"

	teleporterMessengerLabel = "TeleporterMessenger"
)

var (
	LocalNetworkInstance *local.LocalNetwork
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
	LocalNetworkInstance = local.NewLocalNetwork(warpGenesisFile)
	// Generate the Teleporter deployment values
	teleporterContractAddress := common.HexToAddress(testUtils.ReadHexTextFile("./tests/utils/UniversalTeleporterMessengerContractAddress.txt"))
	teleporterDeployerAddress := common.HexToAddress(testUtils.ReadHexTextFile("./tests/utils/UniversalTeleporterDeployerAddress.txt"))
	teleporterDeployerTransactionStr := testUtils.ReadHexTextFile("./tests/utils/UniversalTeleporterDeployerTransaction.txt")
	teleporterDeployerTransaction, err := hex.DecodeString(testUtils.SanitizeHexString(teleporterDeployerTransactionStr))
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

var _ = ginkgo.Describe("[Teleporter Token Bridge integration tests]", func() {
	// Teleporter tests
	ginkgo.It("Send a message from Subnet A to Subnet B, and one from B to A",
		ginkgo.Label(teleporterMessengerLabel),
		func() {
			flows.BasicSendReceive(LocalNetworkInstance)
		})
})
