// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"os"
	"testing"

	deploymentUtils "github.com/ava-labs/teleporter/contract-deployment/utils"
	testUtils "github.com/ava-labs/teleporter/tests/utils"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	teleporterByteCodeFile = "./contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json"
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
	testUtils.SetupNetwork()
	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress, err := deploymentUtils.ConstructKeylessTransaction(teleporterByteCodeFile, false)
	Expect(err).Should(BeNil())

	testUtils.DeployTeleporterContract(teleporterDeployerTransaction, teleporterDeployerAddress, teleporterContractAddress)
})
var _ = ginkgo.AfterSuite(testUtils.TearDownNetwork)
