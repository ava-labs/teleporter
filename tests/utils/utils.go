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

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/awm-relayer/messages/teleporter"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/params"
	"github.com/ava-labs/subnet-evm/tests/utils"
	"github.com/ava-labs/subnet-evm/tests/utils/runner"
	predicateutils "github.com/ava-labs/subnet-evm/utils/predicate"
	"github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

var (
	DefaultTeleporterTransactionGas       uint64 = 200_000
	DefaultTeleporterTransactionGasFeeCap        = big.NewInt(225 * params.GWei)
	DefaultTeleporterTransactionGasTipCap        = big.NewInt(params.GWei)
	DefaultTeleporterTransactionValue            = common.Big0
)

type SendCrossChainMessageEvent struct {
	DestinationChainID ids.ID
	MessageID          *big.Int
	Message            teleporter.TeleporterMessage
}

type ReceiveCrossChainMessageEvent struct {
	OriginChainID ids.ID
	MessageID     *big.Int
	Message       teleporter.TeleporterMessage
}

// Teleporter contract sendCrossChainMessage input type
type SendCrossChainMessageInput struct {
	DestinationChainID      ids.ID
	DestinationAddress      common.Address
	FeeInfo                 FeeInfo
	RequiredGasLimit        *big.Int
	Message                 []byte
	AllowedRelayerAddresses []common.Address
}

type FeeInfo struct {
	ContractAddress common.Address
	Amount          *big.Int
}

//
// Test utility functions
//

// Subscribes to new heads, sends a tx, and waits for the new head to appear
// Returns the new head
func SendAndWaitForTransaction(
	ctx context.Context,
	wsClient ethclient.Client,
	tx *types.Transaction) *types.Header {

	newHeads := make(chan *types.Header, 1)
	subA, err := wsClient.SubscribeNewHead(ctx, newHeads)
	Expect(err).Should(BeNil())
	defer subA.Unsubscribe()

	err = wsClient.SendTransaction(ctx, tx)
	Expect(err).Should(BeNil())
	newHead := <-newHeads

	return newHead
}

func HttpToWebsocketURI(uri string, blockchainID string) string {
	return fmt.Sprintf("ws://%s/ext/bc/%s/ws", strings.TrimPrefix(uri, "http://"), blockchainID)
}

func HttpToRPCURI(uri string, blockchainID string) string {
	return fmt.Sprintf("http://%s/ext/bc/%s/rpc", strings.TrimPrefix(uri, "http://"), blockchainID)
}

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

func NewTestTeleporterTransaction(chainIDInt *big.Int, teleporterAddress common.Address, nonce uint64, data []byte) *types.Transaction {
	return types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainIDInt,
		Nonce:     nonce,
		To:        &teleporterAddress,
		Gas:       DefaultTeleporterTransactionGas,
		GasFeeCap: DefaultTeleporterTransactionGasFeeCap,
		GasTipCap: DefaultTeleporterTransactionGasTipCap,
		Value:     DefaultTeleporterTransactionValue,
		Data:      data,
	})
}

func SetUpProposerVM(ctx context.Context, fundedKey *ecdsa.PrivateKey, manager *runner.NetworkManager, index int) {
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
	for i, uri := range nodeURIs {
		chainAWSURI := HttpToWebsocketURI(uri, blockchainID.String())
		log.Info("Creating ethclient for blockchainA", "wsURI", chainAWSURI)
		client, err := ethclient.Dial(chainAWSURI)
		Expect(err).Should(BeNil())

		// Loop until each node has advanced to >= the height of the block that emitted the warp log
		for {
			block, err := client.BlockByNumber(ctx, nil)
			Expect(err).Should(BeNil())
			if block.NumberU64() >= height {
				log.Info("client accepted the block containing SendWarpMessage", "client", i, "height", block.NumberU64())
				break
			}
		}
	}
}

// Constructs and sends a transaction containing a warp message for the destination chain.
// Returns the signed transaction.
func ConstructAndSendTransaction(
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

	gasLimit, err := teleporter.CalculateReceiveMessageGasLimit(numSigners, requiredGasLimit)
	Expect(err).Should(BeNil())

	callData, err := teleporter.PackReceiveCrossChainMessage(teleporter.ReceiveCrossChainMessageInput{
		MessageIndex:         0,
		RelayerRewardAddress: fundedAddress,
	})
	Expect(err).Should(BeNil())

	baseFee, err := wsClient.EstimateBaseFee(ctx)
	Expect(err).Should(BeNil())

	gasTipCap, err := wsClient.SuggestGasTipCap(ctx)
	Expect(err).Should(BeNil())

	nonce, err := wsClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	gasFeeCap := baseFee.Mul(baseFee, big.NewInt(2))
	gasFeeCap.Add(gasFeeCap, big.NewInt(2500000000))
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

	// Sign and send the transaction on the destination chain
	signer := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(destinationTx, signer, fundedKey)
	Expect(err).Should(BeNil())

	log.Info("Sending transaction to destination chain")
	SendAndWaitForTransaction(ctx, wsClient, signedTx)

	return signedTx
}
