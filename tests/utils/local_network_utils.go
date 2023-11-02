package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/params"
	"github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	"github.com/ava-labs/subnet-evm/x/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

// Issues txs to activate the proposer VM fork on the specified subnet index in the manager
func SetupProposerVM(ctx context.Context, fundedKey *ecdsa.PrivateKey, manager *runner.NetworkManager, index int) {
	subnet := manager.GetSubnets()[index]
	subnetDetails, ok := manager.GetSubnet(subnet)
	Expect(ok).Should(BeTrue())

	chainID := subnetDetails.BlockchainID
	uri := HttpToWebsocketURI(subnetDetails.ValidatorURIs[0], chainID.String())

	client, err := ethclient.Dial(uri)
	Expect(err).Should(BeNil())
	chainIDInt, err := client.ChainID(ctx)
	Expect(err).Should(BeNil())

	err = utils.IssueTxsToActivateProposerVMFork(ctx, chainIDInt, fundedKey, client)
	Expect(err).Should(BeNil())
}

// Blocks until all validators specified in nodeURIs have reached the specified block height
func WaitForAllValidatorsToAcceptBlock(ctx context.Context, nodeURIs []string, blockchainID ids.ID, height uint64) {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for i, uri := range nodeURIs {
		chainAWSURI := HttpToWebsocketURI(uri, blockchainID.String())
		log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI)
		client, err := ethclient.Dial(chainAWSURI)
		Expect(err).Should(BeNil())
		defer client.Close()

		// Loop until each node has advanced to >= the height of the block that emitted the warp log
		for {
			block, err := client.BlockByNumber(cctx, nil)
			Expect(err).Should(BeNil())
			if block.NumberU64() >= height {
				log.Info("client accepted the block containing SendWarpMessage", "client", i, "height", block.NumberU64())
				break
			}
		}
	}
}

// Constructs the aggregate signature, packs the Teleporter message, and relays to the destination
// Returns the receipt on the destination chain
func RelayMessage(
	ctx context.Context,
	sourceBlockHash common.Hash,
	sourceBlockNumber *big.Int,
	source SubnetTestInfo,
	destination SubnetTestInfo,
) *types.Receipt {
	log.Info("Fetching relevant warp logs from the newly produced block")
	logs, err := source.ChainRPCClient.FilterLogs(ctx, interfaces.FilterQuery{
		BlockHash: &sourceBlockHash,
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
	WaitForAllValidatorsToAcceptBlock(ctx, source.ChainNodeURIs, source.BlockchainID, sourceBlockNumber.Uint64())

	// Get the aggregate signature for the Warp message
	log.Info("Fetching aggregate signature from the source chain validators")
	warpClient, err := warpBackend.NewClient(source.ChainNodeURIs[0], source.BlockchainID.String())
	Expect(err).Should(BeNil())
	signedWarpMessageBytes, err := warpClient.GetAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
	Expect(err).Should(BeNil())

	// Construct the transaction to send the Warp message to the destination chain
	signedTx := CreateReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessageBytes,
		big.NewInt(1),
		teleporterContractAddress,
		fundedAddress,
		fundedKey,
		destination.ChainRPCClient,
		destination.ChainIDInt,
	)

	log.Info("Sending transaction to destination chain")
	receipt := SendTransactionAndWaitForAcceptance(ctx, destination.ChainWSClient, destination.ChainRPCClient, signedTx)

	bind, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, source.ChainRPCClient)
	Expect(err).Should(BeNil())
	// Check the transaction logs for the ReceiveCrossChainMessage event emitted by the Teleporter contract
	event, err := GetReceiveEventFromLogs(receipt.Logs, bind)
	Expect(err).Should(BeNil())
	Expect(event.OriginChainID[:]).Should(Equal(source.BlockchainID[:]))
	return receipt
}
