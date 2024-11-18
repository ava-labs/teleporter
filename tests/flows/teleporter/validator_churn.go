package teleporter

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/teleporter/TeleporterMessenger"
	localnetwork "github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

const newNodeCount = 2

func ValidatorChurn(network *localnetwork.LocalNetwork, teleporter utils.TeleporterTestInfo) {
	L1AInfo, L1BInfo := network.GetTwoL1s()
	teleporterContractAddress := teleporter.TeleporterMessengerAddress(L1AInfo)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	ctx := context.Background()

	//
	// Send a Teleporter message on L1 A
	//
	log.Info("Sending Teleporter message on source chain", "destinationBlockchainID", L1BInfo.BlockchainID)
	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: L1BInfo.BlockchainID,
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
		teleporter.TeleporterMessenger(L1AInfo),
		L1AInfo,
		L1BInfo,
		sendCrossChainMessageInput,
		fundedKey,
	)

	sendEvent, err := utils.GetEventFromLogs(
		receipt.Logs,
		teleporter.TeleporterMessenger(L1AInfo).ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	sentTeleporterMessage := sendEvent.Message

	// Construct the signed warp message
	signedWarpMessage := utils.ConstructSignedWarpMessage(ctx, receipt, L1AInfo, L1BInfo)

	//
	// Modify the validator set on L1 A
	//

	// Add new nodes to the validator set
	addValidatorsCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	network.AddL1Validators(addValidatorsCtx, L1AInfo.L1ID, newNodeCount)

	// Refresh the L1 info
	L1AInfo, L1BInfo = network.GetTwoL1s()

	// Trigger the proposer VM to update its height so that the inner VM can see the new validator set
	// We have to update all L1s, not just the ones directly involved in this test to ensure that the
	// proposer VM is updated on all L1s.
	for _, l1Info := range network.GetL1Infos() {
		err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(
			ctx, l1Info.EVMChainID, fundedKey, l1Info.WSClient,
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
		L1BInfo,
	)

	log.Info("Sending transaction to destination chain")
	utils.SendTransactionAndWaitForFailure(ctx, L1BInfo, signedTx)

	// Verify the message was not delivered
	delivered, err := teleporter.TeleporterMessenger(L1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())

	//
	// Retry sending the message, and attempt to relay again. This should succeed.
	//
	log.Info("Retrying message sending on source chain")
	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, L1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := teleporter.TeleporterMessenger(L1AInfo).RetrySendCrossChainMessage(
		optsA, sentTeleporterMessage,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt = utils.WaitForTransactionSuccess(ctx, L1AInfo, tx.Hash())

	teleporter.RelayTeleporterMessage(ctx, receipt, L1AInfo, L1BInfo, true, fundedKey)

	// Verify the message was delivered
	delivered, err = teleporter.TeleporterMessenger(L1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// The test cases now do not require any specific nodes to be validators, so leave the validator set as is.
	// If this changes in the future, this test will need to perform cleanup by removing the nodes that were added
	// and re-adding the nodes that were removed.
}
