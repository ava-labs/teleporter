package flows

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	nodesPerSubnet = 5
	newNodeCount   = 5
)

func ValidatorChurn(network interfaces.LocalNetwork) {
	subnetAInfo, subnetBInfo := utils.GetTwoSubnets(network)
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	//
	// Send a Teleporter message on Subnet A
	//
	log.Info("Sending Teleporter message on source chain", "destinationBlockchainID", subnetBInfo.BlockchainID)
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: subnetBInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 []byte{1, 2, 3, 4},
	}

	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	sendEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetAInfo.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	sentTeleporterMessage := sendEvent.Message

	// Construct the signed warp message
	signedWarpMessage := network.ConstructSignedWarpMessage(ctx, receipt, subnetAInfo, subnetBInfo)

	//
	// Modify the validator set on Subnet A
	//

	// Add new nodes to the validator set
	network.AddSubnetValidators(ctx, subnetAInfo.SubnetID, constructNodesToAddNames(network))

	// Refresh the subnet info
	subnetAInfo, subnetBInfo = utils.GetTwoSubnets(network)

	// Trigger the proposer VM to update its height so that the inner VM can see the new validator set
	// We have to update all subnets, not just the ones directly involved in this test to ensure that the
	// proposer VM is updated on all subnets.
	for _, subnetInfo := range network.GetSubnetsInfo() {
		err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(
			ctx, subnetInfo.EVMChainID, fundedKey, subnetInfo.WSClient,
		)
		Expect(err).Should(BeNil())
	}

	//
	// Attempt to deliver the warp message signed by the old validator set. This should fail.
	//
	// Construct the transaction to send the Warp message to the destination chain
	signedTx := utils.CreateReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessage,
		sendEvent.Message.RequiredGasLimit,
		teleporterContractAddress,
		fundedKey,
		subnetBInfo,
	)

	log.Info("Sending transaction to destination chain")
	utils.SendTransactionAndWaitForFailure(ctx, subnetBInfo, signedTx)

	// Verify the message was not delivered
	delivered, err := subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())

	//
	// Retry sending the message, and attempt to relay again. This should succeed.
	//
	log.Info("Retrying message sending on source chain")
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnetAInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := subnetAInfo.TeleporterMessenger.RetrySendCrossChainMessage(
		optsA, sentTeleporterMessage,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt = utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx)

	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, true)

	// Verify the message was delivered
	delivered, err = subnetBInfo.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// The test cases now do not require any specific nodes to be validators, so leave the validator set as is.
	// If this changes in the future, this test will need to perform cleanup by removing the nodes that were added
	// and re-adding the nodes that were removed.
}

// Each subnet is assumed to have {nodesPerSubnet} nodes named nodeN-bls, where
// N is unique across each subnet. Nodes to be added should thus be named nodeN-bls
// where N starts one greater than the current total number of nodes.
func constructNodesToAddNames(network interfaces.Network) []string {
	startingNodeId := len(network.GetSubnetsInfo())*nodesPerSubnet + 1
	var nodesToAdd []string
	for i := startingNodeId; i < startingNodeId+newNodeCount; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		nodesToAdd = append(nodesToAdd, n)
	}
	return nodesToAdd
}
