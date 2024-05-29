// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package local

import (
	"os"
	"testing"

	"github.com/ava-labs/teleporter-token-bridge/tests/flows"
	"github.com/ava-labs/teleporter/tests/local"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

const (
	teleporterByteCodeFile = "./contracts/lib/teleporter/contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisFile        = "./tests/utils/warp-genesis.json"

	erc20TokenHubLabel    = "ERC20TokenHub"
	erc20TokenSpokeLabel  = "ERC20TokenSpoke"
	nativeTokenHubLabel   = "NativeTokenHub"
	nativeTokenSpokeLabel = "NativeTokenSpoke"
	multiHopLabel         = "MultiHop"
	sendAndCallLabel      = "SendAndCall"
	registrationLabel     = "Registration"
)

var LocalNetworkInstance *local.LocalNetwork

func TestE2E(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}
	format.MaxLength = 10000

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Teleporter e2e test")
}

// Define the Teleporter before and after suite functions.
var _ = ginkgo.BeforeSuite(func() {
	// Create the local network instance
	LocalNetworkInstance = local.NewLocalNetwork(warpGenesisFile)

	// Generate the Teleporter deployment values
	teleporterDeployerTransaction, teleporterDeployerAddress,
		teleporterContractAddress, err := deploymentUtils.ConstructKeylessTransaction(
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

var _ = ginkgo.Describe("[Teleporter Token Bridge integration tests]", func() {
	ginkgo.It("Bridge an ERC20 token between two Subnets",
		ginkgo.Label(erc20TokenHubLabel, erc20TokenSpokeLabel),
		func() {
			flows.ERC20TokenHubERC20TokenSpoke(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to an ERC20 token",
		ginkgo.Label(nativeTokenHubLabel, erc20TokenSpokeLabel),
		func() {
			flows.NativeTokenHubERC20TokenSpoke(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to a native token",
		ginkgo.Label(nativeTokenHubLabel, nativeTokenSpokeLabel),
		func() {
			flows.NativeTokenHubNativeDestination(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token with ERC20TokenHub multi-hop",
		ginkgo.Label(erc20TokenHubLabel, erc20TokenSpokeLabel, multiHopLabel),
		func() {
			flows.ERC20TokenHubERC20TokenSpokeMultiHop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token with NativeTokenHub multi-hop",
		ginkgo.Label(nativeTokenHubLabel, erc20TokenSpokeLabel, multiHopLabel),
		func() {
			flows.NativeTokenHubERC20TokenSpokeMultiHop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token to a native token",
		ginkgo.Label(erc20TokenHubLabel, nativeTokenSpokeLabel),
		func() {
			flows.ERC20TokenHubNativeTokenSpoke(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token with ERC20TokenHub multi-hop",
		ginkgo.Label(erc20TokenHubLabel, nativeTokenSpokeLabel, multiHopLabel),
		func() {
			flows.ERC20TokenHubNativeTokenSpokeMultiHop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to a native token multi-hop",
		ginkgo.Label(nativeTokenHubLabel, nativeTokenSpokeLabel, multiHopLabel),
		func() {
			flows.NativeTokenHubNativeTokenSpokeMultiHop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token using sendAndCall",
		ginkgo.Label(erc20TokenHubLabel, erc20TokenSpokeLabel, sendAndCallLabel),
		func() {
			flows.ERC20TokenHubERC20TokenSpokeSendAndCall(LocalNetworkInstance)
		})
	ginkgo.It("Registration and collateral checks",
		ginkgo.Label(erc20TokenHubLabel, nativeTokenSpokeLabel, registrationLabel),
		func() {
			flows.RegistrationAndCollateralCheck(LocalNetworkInstance)
		})
})
