package tests

import (
	"context"
	"math/big"
	"time"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	. "github.com/ava-labs/teleporter/tests/utils"

	"github.com/ava-labs/awm-relayer/messages/teleporter"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/params"
	"github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

// Ginkgo describe node that acts as a container for the teleporter e2e tests. This test suite
// will run through the following steps in order:
// 1. Send a transaction to the Teleporter contract on Subnet A
// 2. Aggregate signatures and send the Warp message to Subnet B
// 3. Verify receipt of the message on Subnet B
func BasicOneWaySend() {
	var (
		teleporterMessageID *big.Int
	)

	//
	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	//
	ctx := context.Background()

	nonceA, err := subnetAInfo.ChainRPCClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	data, err := teleporter.EVMTeleporterContractABI.Pack(
		"sendCrossChainMessage",
		SendCrossChainMessageInput{
			DestinationChainID: subnetBInfo.BlockchainID,
			DestinationAddress: fundedAddress,
			FeeInfo: FeeInfo{
				ContractAddress: fundedAddress,
				Amount:          big.NewInt(0),
			},
			RequiredGasLimit:        big.NewInt(1),
			AllowedRelayerAddresses: []common.Address{},
			Message:                 []byte{1, 2, 3, 4},
		},
	)
	Expect(err).Should(BeNil())

	// Send a transaction to the Teleporter contract
	tx := NewTestTeleporterTransaction(subnetAInfo.ChainIDInt, teleporterContractAddress, nonceA, data)

	txSigner := types.LatestSignerForChainID(subnetAInfo.ChainIDInt)
	signedTx, err := types.SignTx(tx, txSigner, fundedKey)
	Expect(err).Should(BeNil())

	subA, err := subnetAInfo.ChainWSClient.SubscribeNewHead(ctx, subnetAInfo.NewHeads)
	Expect(err).Should(BeNil())
	defer subA.Unsubscribe()

	log.Info("Sending Teleporter transaction on source chain", "destinationChainID", subnetBInfo.BlockchainID, "txHash", signedTx.Hash())
	err = subnetAInfo.ChainRPCClient.SendTransaction(ctx, signedTx)
	Expect(err).Should(BeNil())

	// Sleep to ensure the new block is published to the subscriber
	time.Sleep(5 * time.Second)
	receipt, err := subnetAInfo.ChainRPCClient.TransactionReceipt(ctx, signedTx.Hash())
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	//
	// Relay the message to the destination
	//

	// Get the latest block from Subnet A, and retrieve the warp message from the logs
	log.Info("Waiting for new block confirmation")
	newHeadA := <-subnetAInfo.NewHeads
	blockHashA := newHeadA.Hash()

	log.Info("Fetching relevant warp logs from the newly produced block")
	logs, err := subnetAInfo.ChainRPCClient.FilterLogs(ctx, interfaces.FilterQuery{
		BlockHash: &blockHashA,
		Addresses: []common.Address{warp.Module.Address},
	})
	Expect(err).Should(BeNil())
	Expect(len(logs)).Should(Equal(1))

	// Check for relevant warp log from subscription and ensure that it matches
	// the log extracted from the last block.
	txLog := logs[0]
	log.Info("Parsing logData as unsigned warp message")
	unsignedMsg, err := avalancheWarp.ParseUnsignedMessage(txLog.Data)
	Expect(err).Should(BeNil())

	// Set local variables for the duration of the test
	unsignedWarpMessageID := unsignedMsg.ID()
	unsignedWarpMsg := unsignedMsg
	log.Info("Parsed unsignedWarpMsg", "unsignedWarpMessageID", unsignedWarpMessageID, "unsignedWarpMessage", unsignedWarpMsg)

	// Loop over each client on chain A to ensure they all have time to accept the block.
	// Note: if we did not confirm this here, the next stage could be racy since it assumes every node
	// has accepted the block.
	WaitForAllValidatorsToAcceptBlock(ctx, subnetAInfo.ChainNodeURIs, subnetAInfo.BlockchainID, newHeadA.Number.Uint64())

	// Get the aggregate signature for the Warp message
	log.Info("Fetching aggregate signature from the source chain validators")
	warpClient, err := warpBackend.NewClient(subnetAInfo.ChainNodeURIs[0], subnetAInfo.BlockchainID.String())
	Expect(err).Should(BeNil())
	signedWarpMessageBytes, err := warpClient.GetAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
	Expect(err).Should(BeNil())

	// Construct the transaction to send the Warp message to the destination chain
	signedTxB := ConstructAndSendTransaction(
		ctx,
		signedWarpMessageBytes,
		big.NewInt(1),
		teleporterContractAddress,
		fundedAddress,
		fundedKey,
		subnetBInfo.ChainRPCClient,
		subnetBInfo.ChainIDInt,
	)

	// Sleep to ensure the new block is published to the subscriber
	time.Sleep(5 * time.Second)
	receipt, err = subnetBInfo.ChainRPCClient.TransactionReceipt(ctx, signedTxB.Hash())
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	sendCrossChainMessageLog := receipt.Logs[0]
	var event SendCrossChainMessageEvent
	err = teleporter.EVMTeleporterContractABI.UnpackIntoInterface(&event, "SendCrossChainMessage", sendCrossChainMessageLog.Data)
	Expect(err).Should(BeNil())
	teleporterMessageID = event.Message.MessageID

	//
	// Check Teleporter message received on the destination
	//
	time.Sleep(5 * time.Second) // Give the relayer a chance to deliver the message to the destination chain
	data, err = teleporter.PackMessageReceived(teleporter.MessageReceivedInput{
		OriginChainID: subnetAInfo.BlockchainID,
		MessageID:     teleporterMessageID,
	})
	Expect(err).Should(BeNil())
	callMessage := interfaces.CallMsg{
		To:   &teleporterContractAddress,
		Data: data,
	}
	result, err := subnetBInfo.ChainRPCClient.CallContract(context.Background(), callMessage, nil)
	Expect(err).Should(BeNil())

	// check the contract call result
	delivered, err := teleporter.UnpackMessageReceivedResult(result)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())
}
