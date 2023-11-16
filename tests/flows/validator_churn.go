package flows

import (
	"context"
	"fmt"
	"math/big"

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

type constructSignedWarpMessageFunc func(
	context.Context,
	*types.Receipt,
	network.SubnetTestInfo,
) []byte

type addSubnetAValidatorsFunc func(context.Context, []string)

// TODO: Disallow this test from being run on anything but a local network, since it manipulates the validator set
func ValidatorChurn(
	network network.Network,
	constructSignedMessageFunc constructSignedWarpMessageFunc,
	addSubnetAValidatorsFunc addSubnetAValidatorsFunc) {
	var (
		teleporterMessageID *big.Int
	)

	subnetAInfo, subnetBInfo := network.GetSubnetInfo()
	teleporterContractAddress := network.GetTeleporterContractAddress()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	subnetATeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		subnetAInfo.RPCClient)
	Expect(err).Should(BeNil())
	subnetBTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress,
		subnetBInfo.RPCClient)
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
	receipt, teleporterMessageID = utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx,
		subnetAInfo,
		subnetBInfo,
		sendCrossChainMessageInput,
		fundedAddress,
		fundedKey,
		subnetATeleporterMessenger,
	)

	sendEvent, err := utils.GetEventFromLogs(receipt.Logs, subnetATeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	sentTeleporterMessage := sendEvent.Message

	// Construct the signed warp message
	signedWarpMessageBytes := constructSignedMessageFunc(ctx, receipt, subnetAInfo)

	//
	// Modify the validator set on Subnet A
	//

	// Add new nodes to the validator set
	log.Info("Adding nodes to the validator set")
	var nodesToAdd []string
	for i := 15; i <= 19; i++ {
		n := fmt.Sprintf("node%d-bls", i)
		nodesToAdd = append(nodesToAdd, n)
	}
	addSubnetAValidatorsFunc(ctx, nodesToAdd)

	// Refresh the subnet info
	subnetAInfo, subnetBInfo = network.GetSubnetInfo()

	// Trigger the proposer VM to update its height so that the inner VM can see the new validator set
	err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(ctx, subnetAInfo.EVMChainID, fundedKey, subnetAInfo.WSClient)
	Expect(err).Should(BeNil())
	err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(ctx, subnetBInfo.EVMChainID, fundedKey, subnetBInfo.WSClient)
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
		subnetBInfo,
		false,
	)

	log.Info("Sending transaction to destination chain")
	receipt = utils.SendTransactionAndWaitForAcceptance(ctx, subnetBInfo.RPCClient, signedTx, false)

	// Verify the message was not delivered
	delivered, err := subnetBTeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())

	//
	// Retry sending the message, and attempt to relay again. This should succeed.
	//
	log.Info("Retrying message sending on source chain")
	optsA := utils.CreateTransactorOpts(ctx, subnetAInfo, fundedAddress, fundedKey)
	tx, err := subnetATeleporterMessenger.RetrySendCrossChainMessage(
		optsA, subnetBInfo.BlockchainID, sentTeleporterMessage,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err = bind.WaitMined(ctx, subnetAInfo.RPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	network.RelayMessage(ctx, receipt, subnetAInfo, subnetBInfo, false, true)

	// Verify the message was delivered
	delivered, err = subnetBTeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, subnetAInfo.BlockchainID, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	// The test cases now do not require any specific nodes to be validators, so leave the validator set as is.
	// If this changes in the future, this test will need to perform cleanup by removing the nodes that were added
	// and re-adding the nodes that were removed.
}
