package tests

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

// Disallow this test from being run on anything but a local network, since it manipulates the validator set
func ValidatorChurnGinkgo() {
	network := &network.LocalNetwork{}
	var (
		teleporterMessageID *big.Int
	)

	subnets := network.GetSubnetsInfo()
	subnetAInfo := subnets[0]
	subnetBInfo := subnets[1]
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetAInfo.ChainRPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, subnetBInfo.ChainRPCClient)
	Expect(err).Should(BeNil())

	ctx := context.Background()

	//
	// Send a Teleporter message on Subnet A
	//
	log.Info("Sending Teleporter message on source chain", "destinationChainID", subnetBInfo.BlockchainID)
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationChainID: subnetBInfo.BlockchainID,
		DestinationAddress: fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			ContractAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	var receipt *types.Receipt
	receipt, teleporterMessageID = utils.SendCrossChainMessageAndWaitForAcceptance(ctx, subnetAInfo, subnetBInfo, sendCrossChainMessageInput, fundedAddress, fundedKey, subnetATeleporterMessenger)

	sendEvent, err := utils.GetSendEventFromLogs(receipt.Logs, subnetATeleporterMessenger)
	Expect(err).Should(BeNil())
	sentTeleporterMessage := sendEvent.Message

	// Construct the signed warp message
	signedWarpMessageBytes := utils.ConstructSignedWarpMessageBytes(ctx, receipt, subnetAInfo, subnetBInfo)

	//
	// Modify the validator set on Subnet A
	//

	// var nodesToRemove []string // All but one Subnet A validator nodes
	// for i := 2; i <= 5; i++ {
	// 	n := fmt.Sprintf("node%d-bls", i)
	// 	nodesToRemove = append(nodesToRemove, n)
	// }

	pChainClient := platformvm.NewClient(subnetAInfo.ChainNodeURIs[0])
	// currentVdrs, err := pChainClient.GetCurrentValidators(ctx, subnetAInfo.SubnetID, []ids.NodeID{})
	// Expect(err).Should(BeNil())
	// var vdrNodes []string
	// for _, vdr := range currentVdrs {
	// 	vdrNodes = append(vdrNodes, vdr.NodeID.String())
	// }
	// height, err := pChainClient.GetHeight(ctx)
	// log.Info("dbg original nodes", "height", height, "nodes", vdrNodes)

	// // Remove the nodes from the validator set
	// log.Info("Removing nodes from the validator set", "nodes", nodesToRemove)
	// utils.RemoveSubnetValidators(ctx, subnetAInfo.SubnetID, nodesToRemove)

	currentVdrs, err := pChainClient.GetCurrentValidators(ctx, subnetAInfo.SubnetID, []ids.NodeID{})
	Expect(err).Should(BeNil())
	var vdrNodes []string
	for _, vdr := range currentVdrs {
		vdrNodes = append(vdrNodes, vdr.NodeID.String())
	}
	height, err := pChainClient.GetHeight(ctx)
	log.Info("dbg after removing nodes", "height", height, "nodes", vdrNodes)

	// Add new nodes to the validator set
	log.Info("Adding nodes to the validator set")
	var nodesToAdd []string
	for i := 15; i <= 19; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		nodesToAdd = append(nodesToAdd, n)
	}
	utils.AddSubnetValidators(ctx, subnetAInfo.SubnetID, nodesToAdd)

	// var originalNodes []string
	// for i := 6; i <= 10; i++ {
	// 	n := fmt.Sprintf("node%d-bls", i)
	// 	originalNodes = append(originalNodes, n)
	// }
	// utils.RestartNodes(ctx, originalNodes)

	currentVdrs, err = pChainClient.GetCurrentValidators(ctx, subnetAInfo.SubnetID, []ids.NodeID{})
	Expect(err).Should(BeNil())
	vdrNodes = nil
	for _, vdr := range currentVdrs {
		vdrNodes = append(vdrNodes, vdr.NodeID.String())
	}
	height, err = pChainClient.GetHeight(ctx)
	log.Info("dbg after adding nodes", "height", height, "nodes", vdrNodes)

	// time.Sleep(3 * time.Minute) // Wait for the proposervm to catch up to the P-Chain
	// Refresh the subnet info
	subnets = network.GetSubnetsInfo()
	subnetAInfo = subnets[0]
	subnetBInfo = subnets[1]

	err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(ctx, subnetAInfo.ChainIDInt, fundedKey, subnetAInfo.ChainWSClient)
	Expect(err).Should(BeNil())
	err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(ctx, subnetBInfo.ChainIDInt, fundedKey, subnetBInfo.ChainWSClient)
	Expect(err).Should(BeNil())

	//
	// Attempt to deliver the warp message signed by the old validator set. This should fail.
	//
	// Construct the transaction to send the Warp message to the destination chain
	signedTx := utils.CreateReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessageBytes,
		sendEvent.Message.RequiredGasLimit,
		teleporterContractAddress,
		fundedAddress,
		fundedKey,
		subnetBInfo.ChainRPCClient,
		subnetBInfo.ChainIDInt,
	)

	log.Info("Sending transaction to destination chain")
	receipt = utils.SendTransactionAndWaitForAcceptance(ctx, subnetBInfo.ChainWSClient, subnetBInfo.ChainRPCClient, signedTx, false)

	// Verify the message was not delivered
	delivered, err := subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())

	//
	// Retry sending the message, and attempt to relay again. This should succeed.
	//
	log.Info("Retrying message sending on source chain")
	optsA := utils.CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)
	tx, err := subnetATeleporterMessenger.RetrySendCrossChainMessage(optsA, subnetBInfo.BlockchainID, sentTeleporterMessage)

	// Wait for the transaction to be mined
	receipt, err = bind.WaitMined(ctx, subnetAInfo.ChainRPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// Verify the message was delivered
	delivered, err = subnetBTeleporterMessenger.MessageReceived(&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// The test cases now do not require any specific nodes to be validators, so leave the validator set as is.
	// If this changes in the future, this test will need to perform cleanup by removing the nodes that were added
	// and re-adding the nodes that were removed.
}
