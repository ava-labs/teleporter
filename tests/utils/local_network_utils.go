package utils

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/params"
	subnetEvmUtils "github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	"github.com/ava-labs/subnet-evm/x/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
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

	err = subnetEvmUtils.IssueTxsToActivateProposerVMFork(ctx, chainIDInt, fundedKey, client)
	Expect(err).Should(BeNil())
}

// Blocks until all validators specified in nodeURIs have reached the specified block height
func WaitForAllValidatorsToAcceptBlock(ctx context.Context, nodeURIs []string, blockchainID ids.ID, height uint64) {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	for i, uri := range nodeURIs {
		chainAWSURI := HttpToWebsocketURI(uri, blockchainID.String())
		log.Info("Creating ethclient for blockchain", "blockchainID", blockchainID.String(), "wsURI", chainAWSURI)
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

func ConstructSignedWarpMessageBytes(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source SubnetTestInfo,
	destination SubnetTestInfo,
) []byte {
	log.Info("Fetching relevant warp logs from the newly produced block")
	logs, err := source.ChainRPCClient.FilterLogs(ctx, interfaces.FilterQuery{
		BlockHash: &sourceReceipt.BlockHash,
		Addresses: []common.Address{warp.Module.Address},
	})
	Expect(err).Should(BeNil())
	Expect(len(logs)).Should(Equal(1))

	// Check for relevant warp log from subscription and ensure that it matches
	// the log extracted from the last block.
	txLog := logs[0]
	log.Info("Parsing logData as unsigned warp message")
	unsignedMsg, err := warp.UnpackSendWarpEventDataToMessage(txLog.Data)
	Expect(err).Should(BeNil())

	// Set local variables for the duration of the test
	unsignedWarpMessageID := unsignedMsg.ID()
	unsignedWarpMsg := unsignedMsg
	log.Info("Parsed unsignedWarpMsg", "unsignedWarpMessageID", unsignedWarpMessageID, "unsignedWarpMessage", unsignedWarpMsg)

	// Loop over each client on chain A to ensure they all have time to accept the block.
	// Note: if we did not confirm this here, the next stage could be racy since it assumes every node
	// has accepted the block.
	WaitForAllValidatorsToAcceptBlock(ctx, source.ChainNodeURIs, source.BlockchainID, sourceReceipt.BlockNumber.Uint64())

	// Get the aggregate signature for the Warp message
	log.Info("Fetching aggregate signature from the source chain validators")
	warpClient, err := warpBackend.NewClient(source.ChainNodeURIs[0], source.BlockchainID.String())
	Expect(err).Should(BeNil())
	signedWarpMessageBytes, err := warpClient.GetMessageAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
	Expect(err).Should(BeNil())

	return signedWarpMessageBytes
}

// Constructs the aggregate signature, packs the Teleporter message, and relays to the destination
// Returns the receipt on the destination chain
func RelayMessage(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source SubnetTestInfo,
	destination SubnetTestInfo,
	expectSuccess bool,
) *types.Receipt {

	// Fetch the Teleporter message from the logs
	bind, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, source.ChainRPCClient)
	Expect(err).Should(BeNil())
	sendEvent, err := GetSendEventFromLogs(sourceReceipt.Logs, bind)
	Expect(err).Should(BeNil())

	signedWarpMessageBytes := ConstructSignedWarpMessageBytes(ctx, sourceReceipt, source, destination)

	// Construct the transaction to send the Warp message to the destination chain
	signedTx := CreateReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessageBytes,
		sendEvent.Message.RequiredGasLimit,
		teleporterContractAddress,
		fundedAddress,
		fundedKey,
		destination.ChainRPCClient,
		destination.ChainIDInt,
	)

	log.Info("Sending transaction to destination chain")
	receipt := SendTransactionAndWaitForAcceptance(ctx, destination.ChainWSClient, destination.ChainRPCClient, signedTx, expectSuccess)

	if !expectSuccess {
		return nil
	}

	bind, err = teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, source.ChainRPCClient)
	Expect(err).Should(BeNil())
	// Check the transaction logs for the ReceiveCrossChainMessage event emitted by the Teleporter contract
	receiveEvent, err := GetReceiveEventFromLogs(receipt.Logs, bind)
	Expect(err).Should(BeNil())
	Expect(receiveEvent.OriginChainID[:]).Should(Equal(source.BlockchainID[:]))
	return receipt
}

func DeployContract(ctx context.Context, byteCodeFileName string, deployerPK *ecdsa.PrivateKey, subnetInfo SubnetTestInfo, abi *abi.ABI, constructorArgs ...interface{}) common.Address {
	// Deploy an example ERC20 contract to be used as the source token
	byteCode, err := deploymentUtils.ExtractByteCode(byteCodeFileName)
	Expect(err).Should(BeNil())
	Expect(len(byteCode) > 0).Should(BeTrue())
	transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetInfo.ChainIDInt)
	Expect(err).Should(BeNil())
	contractAddress, tx, _, err := bind.DeployContract(transactor, *abi, byteCode, subnetInfo.ChainRPCClient, constructorArgs...)
	Expect(err).Should(BeNil())

	// Wait for transaction, then check code was deployed
	WaitForTransaction(ctx, tx.Hash(), subnetInfo.ChainRPCClient)
	code, err := subnetInfo.ChainRPCClient.CodeAt(ctx, contractAddress, nil)
	Expect(err).Should(BeNil())
	Expect(len(code)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

	return contractAddress
}
