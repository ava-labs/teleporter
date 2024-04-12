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
)

const (
	teleporterByteCodeFile = "./contracts/lib/teleporter/contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json"
	warpGenesisFile        = "./tests/utils/warp-genesis.json"

	erc20SourceLabel            = "ERC20Source"
	erc20DestinationLabel       = "ERC20Destination"
	nativeTokenSourceLabel      = "NativeTokenSource"
	nativeTokenDestinationLabel = "NativeTokenDestination"
	multiHopLabel               = "MultiHop"
	sendAndCallLabel            = "SendAndCall"
)

var LocalNetworkInstance *local.LocalNetwork

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
		ginkgo.Label(erc20SourceLabel, erc20DestinationLabel),
		func() {
			flows.ERC20SourceERC20Destination(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to an ERC20 token",
		ginkgo.Label(nativeTokenSourceLabel, erc20DestinationLabel),
		func() {
			flows.NativeSourceERC20Destination(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to a native token",
		ginkgo.Label(nativeTokenSourceLabel, nativeTokenDestinationLabel),
		func() {
			flows.NativeSourceNativeDestination(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token with ERC20Source multihop",
		ginkgo.Label(erc20SourceLabel, erc20DestinationLabel, multiHopLabel),
		func() {
			flows.ERC20SourceERC20DestinationMultihop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token with NativeTokenSource multihop",
		ginkgo.Label(nativeTokenSourceLabel, erc20DestinationLabel, multiHopLabel),
		func() {
			flows.NativeSourceERC20DestinationMultihop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a Native token with ERC20Source",
		ginkgo.Label(erc20SourceLabel, nativeTokenDestinationLabel),
		func() {
			flows.ERC20SourceNativeDestination(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a Native token with ERC20Source multihop",
		ginkgo.Label(erc20SourceLabel, nativeTokenDestinationLabel, multiHopLabel),
		func() {
			flows.ERC20SourceNativeDestinationMultihop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge a native token to a native token multihop",
		ginkgo.Label(nativeTokenSourceLabel, nativeTokenDestinationLabel, multiHopLabel),
		func() {
			flows.NativeSourceNativeDestinationMultihop(LocalNetworkInstance)
		})
	ginkgo.It("Bridge an ERC20 token with ERC20TokenSource Send and Call",
		ginkgo.Label(erc20SourceLabel, erc20DestinationLabel, sendAndCallLabel),
		func() {
			flows.ERC20SourceERC20DestinationSendAndCall(LocalNetworkInstance)
		})
})
