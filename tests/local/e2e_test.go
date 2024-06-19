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

	erc20TokenHomeLabel    = "ERC20TokenHome"
	erc20TokenRemoteLabel  = "ERC20TokenRemote"
	nativeTokenHomeLabel   = "NativeTokenHome"
	nativeTokenRemoteLabel = "NativeTokenRemote"
	multiHopLabel          = "MultiHop"
	sendAndCallLabel       = "SendAndCall"
	registrationLabel      = "Registration"
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
		ginkgo.Label(erc20TokenHomeLabel, erc20TokenRemoteLabel),
		func() {
			flows.ERC20TokenHomeERC20TokenRemote(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to an ERC20 token",
		ginkgo.Label(nativeTokenHomeLabel, erc20TokenRemoteLabel),
		func() {
			flows.NativeTokenHomeERC20TokenRemote(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to a native token",
		ginkgo.Label(nativeTokenHomeLabel, nativeTokenRemoteLabel),
		func() {
			flows.NativeTokenHomeNativeDestination(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token with ERC20TokenHome multi-hop",
		ginkgo.Label(erc20TokenHomeLabel, erc20TokenRemoteLabel, multiHopLabel),
		func() {
			flows.ERC20TokenHomeERC20TokenRemoteMultiHop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token with NativeTokenHome multi-hop",
		ginkgo.Label(nativeTokenHomeLabel, erc20TokenRemoteLabel, multiHopLabel),
		func() {
			flows.NativeTokenHomeERC20TokenRemoteMultiHop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token to a native token",
		ginkgo.Label(erc20TokenHomeLabel, nativeTokenRemoteLabel),
		func() {
			flows.ERC20TokenHomeNativeTokenRemote(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token with ERC20TokenHome multi-hop",
		ginkgo.Label(erc20TokenHomeLabel, nativeTokenRemoteLabel, multiHopLabel),
		func() {
			flows.ERC20TokenHomeNativeTokenRemoteMultiHop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to a native token multi-hop",
		ginkgo.Label(nativeTokenHomeLabel, nativeTokenRemoteLabel, multiHopLabel),
		func() {
			flows.NativeTokenHomeNativeTokenRemoteMultiHop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token using sendAndCall",
		ginkgo.Label(erc20TokenHomeLabel, erc20TokenRemoteLabel, sendAndCallLabel),
		func() {
			flows.ERC20TokenHomeERC20TokenRemoteSendAndCall(LocalNetworkInstance)
		})
	ginkgo.It("Registration and collateral checks",
		ginkgo.Label(erc20TokenHomeLabel, nativeTokenRemoteLabel, registrationLabel),
		func() {
			flows.RegistrationAndCollateralCheck(LocalNetworkInstance)
		})
})
