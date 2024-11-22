package teleporter

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/TeleporterMessenger"
	poavalidatormanager "github.com/ava-labs/teleporter/abi-bindings/go/validator-manager/PoAValidatorManager"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const (
	newNodeCount       = 2
	sleepPeriodSeconds = 5
)

func ValidatorChurn(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	subnetAInfo, subnetBInfo := network.GetTwoSubnets()
	teleporterContractAddress := teleporter.TeleporterMessengerAddress(subnetAInfo)
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
		teleporter.TeleporterMessenger(subnetAInfo),
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	sendEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(subnetAInfo).ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	sentTeleporterMessage := sendEvent.Message

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	// Construct the signed warp message
	signedWarpMessage := utils.ConstructSignedWarpMessage(
		ctx,
		receipt,
		subnetAInfo,
		subnetBInfo,
		nil,
		aggregator,
	)

	//
	// Modify the validator set on Subnet A
	//

	// Add new nodes to the validator set
	addValidatorsCtx, cancel := context.WithTimeout(ctx, (90+sleepPeriodSeconds)*newNodeCount*time.Second)
	defer cancel()
	newNodes := network.GetExtraNodes(newNodeCount)
	validatorManagerAddress := network.GetValidatorManager(subnetAInfo.SubnetID)
	validatorManager, err := poavalidatormanager.NewPoAValidatorManager(validatorManagerAddress, subnetAInfo.RPCClient)
	pChainInfo := utils.GetPChainInfo(network.GetPrimaryNetworkInfo())
	Expect(err).Should(BeNil())

	subnetAInfo = network.AddSubnetValidators(newNodes, subnetAInfo)

	for i := 0; i < newNodeCount; i++ {
		expiry := uint64(time.Now().Add(24 * time.Hour).Unix())
		pop, err := newNodes[i].GetProofOfPossession()
		Expect(err).Should(BeNil())
		node := utils.Node{
			NodeID:  newNodes[i].NodeID,
			NodePoP: pop,
			Weight:  units.Schmeckle,
		}
		utils.InitializeAndCompletePoAValidatorRegistration(
			addValidatorsCtx,
			aggregator,
			fundedKey,
			fundedKey,
			subnetAInfo,
			pChainInfo,
			validatorManager,
			validatorManagerAddress,
			expiry,
			node,
			network.GetPChainWallet(),
			network.GetNetworkID(),
		)
		// Sleep to ensure the validator manager uses a new churn tracking period
		time.Sleep(sleepPeriodSeconds * time.Second)
	}

	// Refresh the subnet info
	subnetAInfo, subnetBInfo = network.GetTwoSubnets()

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
	delivered, err := teleporter.TeleporterMessenger(subnetBInfo).MessageReceived(
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
	tx, err := teleporter.TeleporterMessenger(subnetAInfo).RetrySendCrossChainMessage(
		optsA, sentTeleporterMessage,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt = utils.WaitForTransactionSuccess(ctx, subnetAInfo, tx.Hash())

	teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		subnetAInfo,
		subnetBInfo,
		true,
		fundedKey,
		nil,
		aggregator,
	)

	// Verify the message was delivered
	delivered, err = teleporter.TeleporterMessenger(subnetBInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// The test cases now do not require any specific nodes to be validators, so leave the validator set as is.
	// If this changes in the future, this test will need to perform cleanup by removing the nodes that were added
	// and re-adding the nodes that were removed.
}
