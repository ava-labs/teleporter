// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	gasUtils "github.com/ava-labs/teleporter/utils/gas-utils"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/params"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	"github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

var (
	NativeTransferGas                     uint64 = 21_000
	DefaultTeleporterTransactionGas       uint64 = 200_000
	DefaultTeleporterTransactionGasFeeCap        = big.NewInt(225 * params.GWei)
	DefaultTeleporterTransactionGasTipCap        = big.NewInt(params.GWei)
	DefaultTeleporterTransactionValue            = common.Big0
)

//
// Test utility functions
//

// Subscribes to new heads, sends a tx, and waits for the new head to appear
// Returns the new head
func SendTransactionAndWaitForAcceptance(
	ctx context.Context,
	wsClient ethclient.Client,
	tx *types.Transaction) *types.Receipt {

	newHeads := make(chan *types.Header, 1)
	subA, err := wsClient.SubscribeNewHead(ctx, newHeads)
	Expect(err).Should(BeNil())
	defer subA.Unsubscribe()

	err = wsClient.SendTransaction(ctx, tx)
	Expect(err).Should(BeNil())
	// Wait for the transaction to be accepted
	<-newHeads

	receipt, err := wsClient.TransactionReceipt(ctx, tx.Hash())
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	return receipt
}

func HttpToWebsocketURI(uri string, blockchainID string) string {
	return fmt.Sprintf("ws://%s/ext/bc/%s/ws", strings.TrimPrefix(uri, "http://"), blockchainID)
}

func HttpToRPCURI(uri string, blockchainID string) string {
	return fmt.Sprintf("http://%s/ext/bc/%s/rpc", strings.TrimPrefix(uri, "http://"), blockchainID)
}

// Get the host and port from a URI. The URI should be in the format http://host:port or https://host:port
func GetURIHostAndPort(uri string) (string, uint32, error) {
	// At a minimum uri should have http:// of 7 characters
	Expect(len(uri)).Should(BeNumerically(">", 7))
	if uri[:7] == "http://" {
		uri = uri[7:]
	} else if uri[:8] == "https://" {
		uri = uri[8:]
	} else {
		return "", 0, fmt.Errorf("invalid uri: %s", uri)
	}

	// Split the uri into host and port
	hostAndPort := strings.Split(uri, ":")
	Expect(len(hostAndPort)).Should(Equal(2))

	// Parse the port
	port, err := strconv.ParseUint(hostAndPort[1], 10, 32)
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse port: %w", err)
	}

	return hostAndPort[0], uint32(port), nil
}

// Constructs a transaction to call sendCrossChainMessage
// Returns the signed transaction.
func CreateSendCrossChainMessageTransaction(
	ctx context.Context,
	source SubnetTestInfo,
	input teleportermessenger.TeleporterMessageInput,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	data, err := teleportermessenger.PackSendCrossChainMessage(input)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := calculateTxParams(ctx, source.WSClient, fundedAddress)

	// Send a transaction to the Teleporter contract
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   source.ChainID,
		Nonce:     nonce,
		To:        &teleporterContractAddress,
		Gas:       DefaultTeleporterTransactionGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     DefaultTeleporterTransactionValue,
		Data:      data,
	})

	return signTransaction(tx, fundedKey, source.ChainID)
}

// Constructs a transaction to call receiveCrossChainMessage
// Returns the signed transaction.
func CreateReceiveCrossChainMessageTransaction(
	ctx context.Context,
	warpMessageBytes []byte,
	requiredGasLimit *big.Int,
	teleporterContractAddress common.Address,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	wsClient ethclient.Client,
	chainID *big.Int,
) *types.Transaction {
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing transaction for the destination chain")
	signedMessage, err := avalancheWarp.ParseMessage(warpMessageBytes)
	Expect(err).Should(BeNil())

	numSigners, err := signedMessage.Signature.NumSigners()
	Expect(err).Should(BeNil())

	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(numSigners, requiredGasLimit)
	Expect(err).Should(BeNil())

	callData, err := teleportermessenger.PackReceiveCrossChainMessage(0, fundedAddress)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := calculateTxParams(ctx, wsClient, fundedAddress)

	destinationTx := predicateutils.NewPredicateTx(
		chainID,
		nonce,
		&teleporterContractAddress,
		gasLimit,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessage.Bytes(),
	)

	return signTransaction(destinationTx, fundedKey, chainID)
}

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

func WaitForTransaction(ctx context.Context, txHash common.Hash, client ethclient.Client) *types.Receipt {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Loop until we find the transaction or time out
	for {
		receipt, err := client.TransactionReceipt(cctx, txHash)
		if err == nil {
			return receipt
		} else {
			log.Info("Waiting for transaction", "hash", txHash.Hex())
			time.Sleep(200 * time.Millisecond)
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
	logs, err := source.WSClient.FilterLogs(ctx, interfaces.FilterQuery{
		BlockHash: &sourceBlockHash,
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
	WaitForAllValidatorsToAcceptBlock(ctx, source.NodeURIs, source.BlockchainID, sourceBlockNumber.Uint64())

	// Get the aggregate signature for the Warp message
	log.Info("Fetching aggregate signature from the source chain validators")
	warpClient, err := warpBackend.NewClient(source.NodeURIs[0], source.BlockchainID.String())
	Expect(err).Should(BeNil())
	signedWarpMessageBytes, err := warpClient.GetMessageAggregateSignature(ctx, unsignedWarpMessageID, params.WarpQuorumDenominator)
	Expect(err).Should(BeNil())

	// Construct the transaction to send the Warp message to the destination chain
	signedTx := CreateReceiveCrossChainMessageTransaction(
		ctx,
		signedWarpMessageBytes,
		big.NewInt(1),
		teleporterContractAddress,
		fundedAddress,
		fundedKey,
		destination.WSClient,
		destination.ChainID,
	)

	log.Info("Sending transaction to destination chain")
	receipt := SendTransactionAndWaitForAcceptance(ctx, destination.WSClient, signedTx)

	bind, err := teleportermessenger.NewTeleporterMessenger(teleporterContractAddress, source.WSClient)
	Expect(err).Should(BeNil())
	// Check the transaction logs for the ReceiveCrossChainMessage event emitted by the Teleporter contract
	event, err := GetReceiveEventFromLogs(receipt.Logs, bind)
	Expect(err).Should(BeNil())
	Expect(event.OriginChainID[:]).Should(Equal(source.BlockchainID[:]))
	return nil
}

func GetReceiveEventFromLogs(logs []*types.Log, bind *teleportermessenger.TeleporterMessenger) (*teleportermessenger.TeleporterMessengerReceiveCrossChainMessage, error) {
	for _, log := range logs {
		event, err := bind.ParseReceiveCrossChainMessage(*log)
		if err == nil {
			return event, nil

		}
	}
	return nil, fmt.Errorf("failed to find ReceiveCrossChainMessage event in receipt logs")
}

func GetSendEventFromLogs(logs []*types.Log, bind *teleportermessenger.TeleporterMessenger) (*teleportermessenger.TeleporterMessengerSendCrossChainMessage, error) {
	for _, log := range logs {
		event, err := bind.ParseSendCrossChainMessage(*log)
		if err == nil {
			return event, nil

		}
	}
	return nil, fmt.Errorf("failed to find SendCrossChainMessage event in receipt logs")
}

//
// Unexported functions
//

// Signs a transaction using the provided key for the specified chainID
func signTransaction(tx *types.Transaction, key *ecdsa.PrivateKey, chainID *big.Int) *types.Transaction {
	txSigner := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(tx, txSigner, key)
	Expect(err).Should(BeNil())

	return signedTx
}

// Returns the gasFeeCap, gasTipCap, and nonce the be used when constructing a transaction from fundedAddress
func calculateTxParams(ctx context.Context, wsClient ethclient.Client, fundedAddress common.Address) (*big.Int, *big.Int, uint64) {
	baseFee, err := wsClient.EstimateBaseFee(ctx)
	Expect(err).Should(BeNil())

	gasTipCap, err := wsClient.SuggestGasTipCap(ctx)
	Expect(err).Should(BeNil())

	nonce, err := wsClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	gasFeeCap := baseFee.Mul(baseFee, big.NewInt(2))
	gasFeeCap.Add(gasFeeCap, big.NewInt(2500000000))

	return gasFeeCap, gasTipCap, nonce
}

func createNativeTransferTransaction(
	ctx context.Context,
	network SubnetTestInfo,
	fromAddress common.Address,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Transaction {
	gasFeeCap, gasTipCap, nonce := calculateTxParams(ctx, network.WSClient, fundedAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   network.ChainID,
		Nonce:     nonce,
		To:        &recipient,
		Gas:       NativeTransferGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     amount,
	})

	return signTransaction(tx, fundedKey, network.ChainID)
}
