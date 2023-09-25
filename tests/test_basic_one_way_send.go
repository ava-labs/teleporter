package tests

import (
	"context"
	"math/big"
	"time"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	predicateutils "github.com/ava-labs/subnet-evm/utils/predicate"
	warpBackend "github.com/ava-labs/subnet-evm/warp"

	"github.com/ava-labs/awm-relayer/messages/teleporter"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/params"
	"github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// Ginkgo describe node that acts as a container for the teleporter e2e tests. This test suite
// will run through the following steps in order:
// 1. Send a transaction to the Teleporter contract on Subnet A
// 2. Aggregate signatures and send the Warp message to Subnet B
// 3. Verify receipt of the message on Subnet B
var _ = ginkgo.Describe("[Teleporter one way send]", ginkgo.Ordered, func() {
	var (
		teleporterMessageID *big.Int
	)

	// Send a transaction to Subnet A to issue a Warp Message from the Teleporter contract to Subnet B
	ginkgo.It("Send Message from A to B", ginkgo.Label("Teleporter", "SendTeleporter"), func() {
		ctx := context.Background()

		nonceA, err := ChainARPCClient.NonceAt(ctx, FundedAddress, nil)
		Expect(err).Should(BeNil())

		data, err := teleporter.EVMTeleporterContractABI.Pack(
			"sendCrossChainMessage",
			TeleporterMessageInput{
				DestinationChainID: BlockchainIDB,
				DestinationAddress: FundedAddress,
				FeeInfo: FeeInfo{
					ContractAddress: FundedAddress,
					Amount:          big.NewInt(0),
				},
				RequiredGasLimit:        big.NewInt(1),
				AllowedRelayerAddresses: []common.Address{},
				Message:                 []byte{1, 2, 3, 4},
			},
		)
		Expect(err).Should(BeNil())

		// Send a transaction to the Teleporter contract
		tx := NewTestTeleporterTransaction(ChainAIDInt, TeleporterContractAddress, nonceA, data)

		txSigner := types.LatestSignerForChainID(ChainAIDInt)
		signedTx, err := types.SignTx(tx, txSigner, FundedKey)
		Expect(err).Should(BeNil())

		subA, err := ChainAWSClient.SubscribeNewHead(ctx, NewHeadsA)
		Expect(err).Should(BeNil())
		defer subA.Unsubscribe()

		log.Info("Sending Teleporter transaction on source chain", "destinationChainID", BlockchainIDB, "txHash", signedTx.Hash())
		err = ChainARPCClient.SendTransaction(ctx, signedTx)
		Expect(err).Should(BeNil())

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := ChainARPCClient.TransactionReceipt(ctx, signedTx.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	})

	ginkgo.It("Relay message to destination", ginkgo.Label("Teleporter", "RelayMessage"), func() {
		ctx := context.Background()

		// Get the latest block from Subnet A, and retrieve the warp message from the logs
		log.Info("Waiting for new block confirmation")
		newHeadA := <-NewHeadsA
		blockHashA := newHeadA.Hash()

		log.Info("Fetching relevant warp logs from the newly produced block")
		logs, err := ChainARPCClient.FilterLogs(ctx, interfaces.FilterQuery{
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
		for i, uri := range ChainANodeURIs {
			chainAWSURI := HttpToWebsocketURI(uri, BlockchainIDA.String())
			log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI)
			client, err := ethclient.Dial(chainAWSURI)
			Expect(err).Should(BeNil())

			// Loop until each node has advanced to >= the height of the block that emitted the warp log
			for {
				block, err := client.BlockByNumber(ctx, nil)
				Expect(err).Should(BeNil())
				if block.NumberU64() >= newHeadA.Number.Uint64() {
					log.Info("client accepted the block containing SendWarpMessage", "client", i, "height", block.NumberU64())
					break
				}
			}
		}

		// Get the aggregate signature for the Warp message
		log.Info("Fetching aggregate signature from the source chain validators")
		warpClient, err := warpBackend.NewWarpClient(ChainANodeURIs[0], BlockchainIDA.String())
		Expect(err).Should(BeNil())
		signedWarpMessageBytes, err := warpClient.GetAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
		Expect(err).Should(BeNil())

		// Construct the transaction to send the Warp message to the destination chain
		log.Info("Constructing transaction for the destination chain")
		signedMessage, err := avalancheWarp.ParseMessage(signedWarpMessageBytes)
		Expect(err).Should(BeNil())

		numSigners, err := signedMessage.Signature.NumSigners()
		Expect(err).Should(BeNil())

		gasLimit, err := teleporter.CalculateReceiveMessageGasLimit(numSigners, teleporterMessage.RequiredGasLimit)
		Expect(err).Should(BeNil())

		callData, err := teleporter.EVMTeleporterContractABI.Pack("receiveCrossChainMessage", FundedAddress)
		Expect(err).Should(BeNil())

		baseFee, err := ChainBRPCClient.EstimateBaseFee(ctx)
		Expect(err).Should(BeNil())

		gasTipCap, err := ChainBRPCClient.SuggestGasTipCap(ctx)
		Expect(err).Should(BeNil())

		nonce, err := ChainBRPCClient.NonceAt(ctx, FundedAddress, nil)
		Expect(err).Should(BeNil())

		gasFeeCap := baseFee.Mul(baseFee, big.NewInt(2))
		gasFeeCap.Add(gasFeeCap, big.NewInt(2500000000))
		destinationTx := predicateutils.NewPredicateTx(
			ChainBIDInt,
			nonce,
			&TeleporterContractAddress,
			gasLimit,
			gasFeeCap,
			gasTipCap,
			big.NewInt(0),
			callData,
			types.AccessList{},
			warp.ContractAddress,
			signedMessage.Bytes(),
		)

		// Sign and send the transaction on the destination chain
		signer := types.LatestSignerForChainID(ChainBIDInt)
		signedTxB, err := types.SignTx(destinationTx, signer, FundedKey)
		Expect(err).Should(BeNil())

		log.Info("Sending transaction to destination chain")
		err = ChainBRPCClient.SendTransaction(context.Background(), signedTxB)
		Expect(err).Should(BeNil())

		// Sleep to ensure the new block is published to the subscriber
		time.Sleep(5 * time.Second)
		receipt, err := ChainBRPCClient.TransactionReceipt(ctx, signedTxB.Hash())
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

		sendCrossChainMessageLog := receipt.Logs[0]
		var event SendCrossChainMessageEvent
		err = teleporter.EVMTeleporterContractABI.UnpackIntoInterface(&event, "SendCrossChainMessage", sendCrossChainMessageLog.Data)
		Expect(err).Should(BeNil())
		teleporterMessageID = event.Message.MessageID
	})

	ginkgo.It("Check Teleporter Message Received", ginkgo.Label("Teleporter", "TeleporterMessageReceived"), func() {
		time.Sleep(5 * time.Second) // Give the relayer a chance to deliver the message to the destination chain
		data, err := teleporter.PackMessageReceivedMessage(teleporter.MessageReceivedInput{
			OriginChainID: BlockchainIDA,
			MessageID:     teleporterMessageID,
		})
		Expect(err).Should(BeNil())
		callMessage := interfaces.CallMsg{
			To:   &TeleporterContractAddress,
			Data: data,
		}
		result, err := ChainBRPCClient.CallContract(context.Background(), callMessage, nil)
		Expect(err).Should(BeNil())

		// check the contract call result
		delivered, err := teleporter.UnpackMessageReceivedResult(result)
		Expect(err).Should(BeNil())
		Expect(delivered).Should(BeTrue())
	})

})
