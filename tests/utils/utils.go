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
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/params"
	predicateutils "github.com/ava-labs/subnet-evm/utils/predicate"
	"github.com/ava-labs/subnet-evm/x/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	gasUtils "github.com/ava-labs/teleporter/utils/gas-utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

var (
	NativeTransferGas                     uint64 = 21_000
	DefaultTeleporterTransactionGas       uint64 = 300_000
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
	rpcClient ethclient.Client,
	tx *types.Transaction,
	expectSuccess bool) *types.Receipt {

	err := rpcClient.SendTransaction(ctx, tx)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be accepted
	receipt, err := bind.WaitMined(ctx, rpcClient, tx)
	Expect(err).Should(BeNil())
	if expectSuccess {
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	} else {
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusFailed))
	}

	return receipt
}

func SendCrossChainMessageAndWaitForAcceptance(
	ctx context.Context,
	source SubnetTestInfo,
	destination SubnetTestInfo,
	input teleportermessenger.TeleporterMessageInput,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, *big.Int) {
	opts := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)

	// Send a transaction to the Teleporter contract
	txn, err := transactor.SendCrossChainMessage(opts, input)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be accepted
	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetSendEventFromLogs(receipt.Logs, transactor)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(destination.BlockchainID[:]))

	log.Info("Sending SendCrossChainMessage transaction on source chain",
		"sourceChainID", source.BlockchainID,
		"destinationChainID", destination.BlockchainID,
		"txHash", txn.Hash())

	return receipt, event.Message.MessageID
}

func SendAddFeeAmountAndWaitForAcceptance(
	ctx context.Context,
	source SubnetTestInfo,
	destination SubnetTestInfo,
	messageID *big.Int,
	amount *big.Int,
	feeContractAddress common.Address,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) *types.Receipt {
	opts := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)

	txn, err := transactor.AddFeeAmount(opts, destination.BlockchainID, messageID, feeContractAddress, amount)
	Expect(err).Should(BeNil())
	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	addFeeAmountEvent, err := GetAddFeeAmountEventFromLogs(receipt.Logs, transactor)
	Expect(err).Should(BeNil())
	Expect(addFeeAmountEvent.MessageID).Should(Equal(messageID))
	Expect(addFeeAmountEvent.DestinationChainID[:]).Should(Equal(destination.BlockchainID[:]))

	log.Info("Send AddFeeAmount transaction on source chain",
		"messageID", messageID,
		"sourceChainID", source.BlockchainID,
		"destinationChainID", destination.BlockchainID)

	return receipt
}

func RetryMessageExecutionAndWaitForAcceptance(
	ctx context.Context,
	originChainID ids.ID,
	subnet SubnetTestInfo,
	message teleportermessenger.TeleporterMessage,
	funedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) *types.Receipt {
	opts := CreateTransactorOpts(ctx, subnet, fundedAddress, fundedKey)

	txn, err := transactor.RetryMessageExecution(opts, originChainID, message)
	Expect(err).Should(BeNil())

	receipt, err := bind.WaitMined(ctx, subnet.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	return receipt
}

func SendSpecifiedReceiptsAndWaitForAcceptance(
	ctx context.Context,
	originChainID ids.ID,
	source SubnetTestInfo,
	messageIDs []*big.Int,
	feeInfo teleportermessenger.TeleporterFeeInfo,
	allowedRelayerAddresses []common.Address,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, *big.Int) {
	opts := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)

	txn, err := transactor.SendSpecifiedReceipts(opts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
	Expect(err).Should(BeNil())

	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetSendEventFromLogs(receipt.Logs, transactor)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(originChainID[:]))

	log.Info("Sending SendSpecifiedReceipts transaction",
		"originChainID", originChainID,
		"txHash", txn.Hash())

	return receipt, event.Message.MessageID
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

//
// Transaction creation functions
//

func CreateTransactorOpts(ctx context.Context, subnet SubnetTestInfo, fundedAddress common.Address, fundedKey *ecdsa.PrivateKey) *bind.TransactOpts {
	// set up parameters
	transactor, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, subnet.ChainIDInt)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnet.ChainRPCClient, fundedAddress)

	transactor.From = fundedAddress
	transactor.Nonce = new(big.Int).SetUint64(nonce)
	transactor.GasTipCap = gasTipCap
	transactor.GasFeeCap = gasFeeCap

	return transactor
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

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, source.ChainRPCClient, fundedAddress)

	// Send a transaction to the Teleporter contract
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   source.ChainIDInt,
		Nonce:     nonce,
		To:        &teleporterContractAddress,
		Gas:       DefaultTeleporterTransactionGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     DefaultTeleporterTransactionValue,
		Data:      data,
	})

	return SignTransaction(tx, fundedKey, source.ChainIDInt)
}

func CreateRetryMessageExecutionTransaction(
	ctx context.Context,
	subnetInfo SubnetTestInfo,
	originChainID ids.ID,
	message teleportermessenger.TeleporterMessage,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	teleporterABI, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
	Expect(err).Should(BeNil())

	data, err := teleporterABI.Pack("retryMessageExecution", originChainID, message)
	Expect(err).Should(BeNil())

	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(10, message.RequiredGasLimit) // TODO: replace with actual number of signers
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo.ChainRPCClient, fundedAddress)

	// Send a transaction to the Teleporter contract
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   subnetInfo.ChainIDInt,
		Nonce:     nonce,
		To:        &teleporterContractAddress,
		Gas:       gasLimit,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     DefaultTeleporterTransactionValue,
		Data:      data,
	})

	return SignTransaction(tx, fundedKey, subnetInfo.ChainIDInt)
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
	rpcClient ethclient.Client,
	chainID *big.Int,
	alterMessage bool,
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

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, rpcClient, fundedAddress)

	if alterMessage {
		// Alter the message
		modifiedMessage := make([]byte, len(signedMessage.Payload))
		copy(modifiedMessage, signedMessage.Payload)
		modifiedMessage[0] = ^modifiedMessage[0]
		Expect(modifiedMessage[:]).ShouldNot(Equal(signedMessage.Payload[:]))
		signedMessage.Payload = modifiedMessage
	}

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

	return SignTransaction(destinationTx, fundedKey, chainID)
}

func CreateNativeTransferTransaction(
	ctx context.Context,
	network SubnetTestInfo,
	fromAddress common.Address,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Transaction {
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, network.ChainRPCClient, fromAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   network.ChainIDInt,
		Nonce:     nonce,
		To:        &recipient,
		Gas:       NativeTransferGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     amount,
	})

	return SignTransaction(tx, fromKey, network.ChainIDInt)
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

//
// Event getters
//

func GetReceiveEventFromLogs(logs []*types.Log, bind *teleportermessenger.TeleporterMessenger) (*teleportermessenger.TeleporterMessengerReceiveCrossChainMessage, error) {
	for _, log := range logs {
		event, err := bind.ParseReceiveCrossChainMessage(*log)
		if err == nil {
			return event, nil

		}
	}
	return nil, fmt.Errorf("failed to find ReceiveCrossChainMessage event in receipt logs")
}

func GetMessageExecutionFailedFromLogs(logs []*types.Log, bind *teleportermessenger.TeleporterMessenger) (*teleportermessenger.TeleporterMessengerMessageExecutionFailed, error) {
	for _, log := range logs {
		event, err := bind.ParseMessageExecutionFailed(*log)
		if err == nil {
			return event, nil

		}
	}
	return nil, fmt.Errorf("failed to find MessageExecutionFailed event in receipt logs")
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

func GetMessageExecutionFailedEventFromLogs(logs []*types.Log, bind *teleportermessenger.TeleporterMessenger) (*teleportermessenger.TeleporterMessengerMessageExecutionFailed, error) {
	for _, log := range logs {
		event, err := bind.ParseMessageExecutionFailed(*log)
		if err == nil {
			return event, nil

		}
	}
	return nil, fmt.Errorf("failed to find MessageExecutionFailed event in receipt logs")
}

func GetMessageExecutedEventFromLogs(logs []*types.Log, bind *teleportermessenger.TeleporterMessenger) (*teleportermessenger.TeleporterMessengerMessageExecuted, error) {
	for _, log := range logs {
		event, err := bind.ParseMessageExecuted(*log)
		if err == nil {
			return event, nil

		}
	}
	return nil, fmt.Errorf("failed to find MessageExecuted event in receipt logs")
}

func GetAddFeeAmountEventFromLogs(logs []*types.Log, bind *teleportermessenger.TeleporterMessenger) (*teleportermessenger.TeleporterMessengerAddFeeAmount, error) {
	for _, log := range logs {
		event, err := bind.ParseAddFeeAmount(*log)
		if err == nil {
			return event, nil

		}
	}
	return nil, fmt.Errorf("failed to find AddFeeAmount event in receipt logs")
}

//
// Unexported functions
//

// Signs a transaction using the provided key for the specified chainID
func SignTransaction(tx *types.Transaction, key *ecdsa.PrivateKey, chainID *big.Int) *types.Transaction {
	txSigner := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(tx, txSigner, key)
	Expect(err).Should(BeNil())

	return signedTx
}

// Returns the gasFeeCap, gasTipCap, and nonce the be used when constructing a transaction from fundedAddress
func CalculateTxParams(ctx context.Context, rpcClient ethclient.Client, fundedAddress common.Address) (*big.Int, *big.Int, uint64) {
	baseFee, err := rpcClient.EstimateBaseFee(ctx)
	Expect(err).Should(BeNil())

	gasTipCap, err := rpcClient.SuggestGasTipCap(ctx)
	Expect(err).Should(BeNil())

	nonce, err := rpcClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	gasFeeCap := baseFee.Mul(baseFee, big.NewInt(2))
	gasFeeCap.Add(gasFeeCap, big.NewInt(2500000000))

	return gasFeeCap, gasTipCap, nonce
}
