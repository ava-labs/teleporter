// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	gasUtils "github.com/ava-labs/teleporter/utils/gas-utils"

	"github.com/ava-labs/avalanchego/ids"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/eth/tracers"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/params"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

type SubnetTestInfo struct {
	SubnetID                  ids.ID
	BlockchainID              ids.ID
	ChainNodeURIs             []string
	ChainWSClient             ethclient.Client
	ChainRPCClient            ethclient.Client
	ChainIDInt                *big.Int
	TeleporterRegistryAddress common.Address
	TeleporterMessenger       *teleportermessenger.TeleporterMessenger
}

//
// Test utility functions
//

// Subscribes to new heads, sends a tx, and waits for the new head to appear
// Returns the new head
func SendTransactionAndWaitForAcceptance(
	ctx context.Context,
	subnetInfo SubnetTestInfo,
	tx *types.Transaction,
	expectSuccess bool) *types.Receipt {
	err := subnetInfo.ChainRPCClient.SendTransaction(ctx, tx)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be accepted
	receipt, err := bind.WaitMined(ctx, subnetInfo.ChainRPCClient, tx)
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
	fundedKey *ecdsa.PrivateKey,
	// transactor *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.ChainIDInt)
	Expect(err).Should(BeNil())

	// Send a transaction to the Teleporter contract
	txn, err := source.TeleporterMessenger.SendCrossChainMessage(opts, input)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be accepted
	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())

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
	fundedKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, source.ChainIDInt)
	Expect(err).Should(BeNil())

	txn, err := transactor.AddFeeAmount(opts, destination.BlockchainID, messageID, feeContractAddress, amount)
	Expect(err).Should(BeNil())
	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	addFeeAmountEvent, err := GetEventFromLogs(receipt.Logs, transactor.ParseAddFeeAmount)
	Expect(err).Should(BeNil())
	Expect(addFeeAmountEvent.MessageID).Should(Equal(messageID))
	Expect(addFeeAmountEvent.DestinationBlockchainID[:]).Should(Equal(destination.BlockchainID[:]))

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
	fundedKey *ecdsa.PrivateKey,
	// transactor *teleportermessenger.TeleporterMessenger,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, subnet.ChainIDInt)
	Expect(err).Should(BeNil())

	txn, err := subnet.TeleporterMessenger.RetryMessageExecution(opts, originChainID, message)
	Expect(err).Should(BeNil())

	receipt, err := bind.WaitMined(ctx, subnet.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	return receipt
}

func RedeemRelayerRewardsAndConfirm(
	ctx context.Context,
	subnet SubnetTestInfo,
	feeToken *exampleerc20.ExampleERC20,
	feeTokenAddress common.Address,
	relayerKey *ecdsa.PrivateKey,
	expectedAmount *big.Int,
) *types.Receipt {
	relayerAddress := crypto.PubkeyToAddress(relayerKey.PublicKey)

	balanceBeforeRedemption, err := feeToken.BalanceOf(
		&bind.CallOpts{}, relayerAddress,
	)
	Expect(err).Should(BeNil())

	tx_opts, err := bind.NewKeyedTransactorWithChainID(
		relayerKey, subnet.ChainIDInt,
	)
	Expect(err).Should(BeNil())
	transaction, err := subnet.TeleporterMessenger.RedeemRelayerRewards(
		tx_opts, feeTokenAddress,
	)
	Expect(err).Should(BeNil())
	receipt, err := bind.WaitMined(ctx, subnet.ChainRPCClient, transaction)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	balanceAfterRedemption, err := feeToken.BalanceOf(
		&bind.CallOpts{}, relayerAddress,
	)
	Expect(err).Should(BeNil())

	Expect(balanceAfterRedemption).Should(
		Equal(
			big.NewInt(0).Add(
				balanceBeforeRedemption, expectedAmount,
			),
		),
	)

	return receipt
}

func SendSpecifiedReceiptsAndWaitForAcceptance(
	ctx context.Context,
	originChainID ids.ID,
	source SubnetTestInfo,
	messageIDs []*big.Int,
	feeInfo teleportermessenger.TeleporterFeeInfo,
	allowedRelayerAddresses []common.Address,
	fundedKey *ecdsa.PrivateKey,
	// transactor *teleportermessenger.TeleporterMessenger,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, source.ChainIDInt)
	Expect(err).Should(BeNil())

	txn, err := source.TeleporterMessenger.SendSpecifiedReceipts(
		opts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
	Expect(err).Should(BeNil())

	receipt, err := bind.WaitMined(ctx, source.ChainRPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(originChainID[:]))

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

// Constructs a transaction to call sendCrossChainMessage
// Returns the signed transaction.
func CreateSendCrossChainMessageTransaction(
	ctx context.Context,
	source SubnetTestInfo,
	input teleportermessenger.TeleporterMessageInput,
	relayerKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	data, err := teleportermessenger.PackSendCrossChainMessage(input)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, source, PrivateKeyToAddress(relayerKey))

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

	return SignTransaction(tx, relayerKey, source.ChainIDInt)
}

func CreateRetryMessageExecutionTransaction(
	ctx context.Context,
	subnetInfo SubnetTestInfo,
	originChainID ids.ID,
	message teleportermessenger.TeleporterMessage,
	fundedKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	fundedAddress := crypto.PubkeyToAddress(fundedKey.PublicKey)

	data, err := teleportermessenger.PackRetryMessageExecution(originChainID, message)
	Expect(err).Should(BeNil())

	// TODO: replace with actual number of signers
	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(10, message.RequiredGasLimit)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo, fundedAddress)

	// Sign a transaction to the Teleporter contract
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
	relayerKey *ecdsa.PrivateKey,
	subnetInfo SubnetTestInfo,
) *types.Transaction {
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing transaction for the destination chain")
	signedMessage, err := avalancheWarp.ParseMessage(warpMessageBytes)
	Expect(err).Should(BeNil())

	numSigners, err := signedMessage.Signature.NumSigners()
	Expect(err).Should(BeNil())

	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(numSigners, requiredGasLimit)
	Expect(err).Should(BeNil())

	callData, err := teleportermessenger.PackReceiveCrossChainMessage(0, PrivateKeyToAddress(relayerKey))
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo, PrivateKeyToAddress(relayerKey))

	destinationTx := predicateutils.NewPredicateTx(
		subnetInfo.ChainIDInt,
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

	return SignTransaction(destinationTx, relayerKey, subnetInfo.ChainIDInt)
}

func CreateNativeTransferTransaction(
	ctx context.Context,
	subnetInfo SubnetTestInfo,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Transaction {
	fromAddress := crypto.PubkeyToAddress(fromKey.PublicKey)
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo, fromAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   subnetInfo.ChainIDInt,
		Nonce:     nonce,
		To:        &recipient,
		Gas:       NativeTransferGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     amount,
	})

	return SignTransaction(tx, fromKey, subnetInfo.ChainIDInt)
}

func SendAndWaitForTransaction(ctx context.Context, tx *types.Transaction, subnetInfo SubnetTestInfo) *types.Receipt {
	err := subnetInfo.ChainWSClient.SendTransaction(ctx, tx)
	Expect(err).Should(BeNil())

	return WaitForTransaction(ctx, tx.Hash(), subnetInfo)
}

func WaitForTransactionSuccess(ctx context.Context, txHash common.Hash, subnetInfo SubnetTestInfo) *types.Receipt {
	receipt := WaitForTransaction(ctx, txHash, subnetInfo)
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	return receipt
}

func WaitForTransaction(ctx context.Context, txHash common.Hash, subnetInfo SubnetTestInfo) *types.Receipt {
	cctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// Loop until we find the transaction or time out
	for {
		receipt, err := subnetInfo.ChainRPCClient.TransactionReceipt(cctx, txHash)
		if err == nil {
			return receipt
		} else {
			Expect(err).ShouldNot(Equal(context.DeadlineExceeded))
			log.Info("Waiting for transaction", "hash", txHash.Hex())
			time.Sleep(500 * time.Millisecond)
		}
	}
}

//
// Event getters
//

// Returns the first log in 'logs' that is successfully parsed by 'parser'
func GetEventFromLogs[T any](logs []*types.Log, parser func(log types.Log) (T, error)) (T, error) {
	for _, log := range logs {
		event, err := parser(*log)
		if err == nil {
			return event, nil
		}
	}
	return *new(T), fmt.Errorf("failed to find %T event in receipt logs", *new(T))
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
func CalculateTxParams(
	ctx context.Context,
	subnetInfo SubnetTestInfo,
	sender common.Address,
) (*big.Int, *big.Int, uint64) {
	baseFee, err := subnetInfo.ChainRPCClient.EstimateBaseFee(ctx)
	Expect(err).Should(BeNil())

	gasTipCap, err := subnetInfo.ChainRPCClient.SuggestGasTipCap(ctx)
	Expect(err).Should(BeNil())

	nonce, err := subnetInfo.ChainRPCClient.NonceAt(ctx, sender, nil)
	Expect(err).Should(BeNil())

	gasFeeCap := baseFee.Mul(baseFee, big.NewInt(gasUtils.BaseFeeFactor))
	gasFeeCap.Add(gasFeeCap, big.NewInt(gasUtils.MaxPriorityFeePerGas))

	return gasFeeCap, gasTipCap, nonce
}

func PrivateKeyToAddress(k *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(k.PublicKey)
}

// Throws a Gomega error if there is a mismatch
func CheckBalance(ctx context.Context, addr common.Address, expectedBalance *big.Int, wsClient ethclient.Client) {
	bal, err := wsClient.BalanceAt(ctx, addr, nil)
	Expect(err).Should(BeNil())
	ExpectBigEqual(bal, expectedBalance)
}

func TraceTransaction(ctx context.Context, txHash common.Hash, subnetInfo SubnetTestInfo) string {
	url := HttpToRPCURI(subnetInfo.ChainNodeURIs[0], subnetInfo.BlockchainID.String())
	rpcClient, err := rpc.DialContext(ctx, url)
	Expect(err).Should(BeNil())
	defer rpcClient.Close()

	var result interface{}
	ct := "callTracer"
	err = rpcClient.Call(&result, "debug_traceTransaction", txHash.String(), tracers.TraceConfig{Tracer: &ct})
	Expect(err).Should(BeNil())

	jsonStr, err := json.Marshal(result)
	Expect(err).Should(BeNil())

	return string(jsonStr)
}

func DeployContract(ctx context.Context, byteCodeFileName string, deployerPK *ecdsa.PrivateKey, subnetInfo SubnetTestInfo, abi *abi.ABI, constructorArgs ...interface{}) {
	// Deploy an example ERC20 contract to be used as the source token
	byteCode, err := deploymentUtils.ExtractByteCode(byteCodeFileName)
	Expect(err).Should(BeNil())
	Expect(len(byteCode) > 0).Should(BeTrue())
	transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetInfo.ChainIDInt)
	Expect(err).Should(BeNil())
	contractAddress, tx, _, err := bind.DeployContract(transactor, *abi, byteCode, subnetInfo.ChainWSClient, constructorArgs...)
	Expect(err).Should(BeNil())

	// Wait for transaction, then check code was deployed
	receipt := WaitForTransaction(ctx, tx.Hash(), subnetInfo)
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	code, err := subnetInfo.ChainWSClient.CodeAt(ctx, contractAddress, nil)
	Expect(err).Should(BeNil())
	Expect(len(code)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode
}

func ExpectBigEqual(v1 *big.Int, v2 *big.Int) {
	// Compare strings, so gomega will print the numbers if they differ
	Expect(v1.String()).Should(Equal(v2.String()))
}

func BigIntSub(v1 *big.Int, v2 *big.Int) *big.Int {
	return big.NewInt(0).Sub(v1, v2)
}

func BigIntMul(v1 *big.Int, v2 *big.Int) *big.Int {
	return big.NewInt(0).Mul(v1, v2)
}
