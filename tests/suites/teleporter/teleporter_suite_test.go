// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter_test

import (
	"context"
	goLog "log"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	warpMessage "github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
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
	ctx, cancel := context.WithTimeout(context.Background(), 120*2*time.Second)
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
	fundedAddress, fundedKey := LocalNetworkInstance.GetFundedAccountInfo()
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

	// Convert the subnets to be managed by a PoAValidatorManager
	cChainInfo := LocalNetworkInstance.GetPrimaryNetworkInfo()
	for _, subnet := range LocalNetworkInstance.GetSubnetsInfo() {
		signatureAggregator := utils.NewSignatureAggregator(
			cChainInfo.NodeURIs[0],
			[]ids.ID{
				subnet.SubnetID,
			},
		)
		vdrManagerAddress, _ := utils.DeployAndInitializePoAValidatorManager(
			ctx,
			fundedKey,
			subnet,
			fundedAddress,
		)

		tmpnetNodes := LocalNetworkInstance.GetExtraNodes(2)
		sort.Slice(tmpnetNodes, func(i, j int) bool {
			return string(tmpnetNodes[i].NodeID.Bytes()) < string(tmpnetNodes[j].NodeID.Bytes())
		})
		totalWeight := uint64(len(tmpnetNodes)-1) * units.Schmeckle
		var nodes []utils.Node
		// Construct the convert subnet info
		destAddr, err := address.ParseToID(utils.DefaultPChainAddress)
		Expect(err).Should(BeNil())
		vdrs := make([]*txs.ConvertSubnetValidator, len(tmpnetNodes))
		for i, node := range tmpnetNodes {
			weight := units.Schmeckle
			if i == len(tmpnetNodes)-1 {
				weight = 4 * totalWeight
			}

			signer, err := node.GetProofOfPossession()
			Expect(err).Should(BeNil())
			nodes = append(nodes, utils.Node{
				NodeID:  node.NodeID,
				NodePoP: signer,
				Weight:  weight,
			})
			vdrs[i] = &txs.ConvertSubnetValidator{
				NodeID:  node.NodeID.Bytes(),
				Weight:  weight,
				Balance: units.Avax * 100,
				Signer:  *signer,
				RemainingBalanceOwner: warpMessage.PChainOwner{
					Threshold: 1,
					Addresses: []ids.ShortID{destAddr},
				},
				DeactivationOwner: warpMessage.PChainOwner{
					Threshold: 1,
					Addresses: []ids.ShortID{destAddr},
				},
			}
		}
		pChainWallet := LocalNetworkInstance.GetPChainWallet()
		_, err = pChainWallet.IssueConvertSubnetTx(
			subnet.SubnetID,
			subnet.BlockchainID,
			vdrManagerAddress[:],
			vdrs,
		)
		Expect(err).Should(BeNil())

		// Modify the each node's config to track the subnet
		for _, node := range tmpnetNodes {
			existingTrackedSubnets, err := node.Flags.GetStringVal(config.TrackSubnetsKey)
			Expect(err).Should(BeNil())
			if existingTrackedSubnets == subnet.SubnetID.String() {
				goLog.Printf("Node %s @ %s already tracking subnet %s\n", node.NodeID, node.URI, subnet.SubnetID)
				continue
			}
			node.Flags[config.TrackSubnetsKey] = subnet.SubnetID.String()

			// Add the node to the network
			LocalNetworkInstance.Network.Nodes = append(LocalNetworkInstance.Network.Nodes, node)
		}
		LocalNetworkInstance.Network.StartNodes(context.Background(), os.Stdout, tmpnetNodes...)

		// Update the tmpnet Subnet struct
		for _, tmpnetSubnet := range LocalNetworkInstance.Network.Subnets {
			if tmpnetSubnet.SubnetID == subnet.SubnetID {
				for _, tmpnetNode := range tmpnetNodes {
					tmpnetSubnet.ValidatorIDs = append(tmpnetSubnet.ValidatorIDs, tmpnetNode.NodeID)
				}
			}
		}
		// Refresh the subnet info after restarting the nodes
		subnet = LocalNetworkInstance.GetSubnetInfo(subnet.SubnetID)

		utils.PChainProposerVMWorkaround(pChainWallet)
		utils.AdvanceProposerVM(ctx, subnet, fundedKey, 5)

		utils.InitializeValidatorSet(
			ctx,
			fundedKey,
			subnet,
			utils.GetPChainInfo(cChainInfo),
			vdrManagerAddress,
			LocalNetworkInstance.GetNetworkID(),
			signatureAggregator,
			nodes,
		)
		// TODO: Remove the bootstrap node as a subnet validator
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
