// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	erc20bridge "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ERC20Bridge/ERC20Bridge"
	examplecrosschainmessenger "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ExampleMessenger/ExampleCrossChainMessenger"
	blockhashpublisher "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/VerifiedBlockHash/BlockHashPublisher"
	blockhashreceiver "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/VerifiedBlockHash/BlockHashReceiver"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	teleporterregistry "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/upgrades/TeleporterRegistry"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	gasUtils "github.com/ava-labs/teleporter/utils/gas-utils"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/eth/tracers"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetEvmInterfaces "github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

var (
	NativeTransferGas                   uint64 = 21_000
	DefaultTeleporterTransactionGas     uint64 = 300_000
	DefaultTeleporterTransactionValue          = common.Big0
	ExpectedExampleERC20DeployerBalance        = new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e10))
)

const (
	CChainPathSpecifier = "C"
)

//
// Test utility functions
//

func SendAddFeeAmountAndWaitForAcceptance(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	messageID ids.ID,
	amount *big.Int,
	feeContractAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.AddFeeAmount(opts, messageID, feeContractAddress, amount)
	Expect(err).Should(BeNil())

	receipt := WaitForTransactionSuccess(ctx, source, tx.Hash())

	addFeeAmountEvent, err := GetEventFromLogs(receipt.Logs, transactor.ParseAddFeeAmount)
	Expect(err).Should(BeNil())
	Expect(addFeeAmountEvent.MessageID[:]).Should(Equal(messageID[:]))

	log.Info("Send AddFeeAmount transaction on source chain",
		"messageID", messageID,
		"sourceChainID", source.BlockchainID,
		"destinationBlockchainID", destination.BlockchainID)

	return receipt
}

func RetryMessageExecutionAndWaitForAcceptance(
	ctx context.Context,
	sourceBlockchainID ids.ID,
	subnet interfaces.SubnetTestInfo,
	message teleportermessenger.TeleporterMessage,
	senderKey *ecdsa.PrivateKey,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := subnet.TeleporterMessenger.RetryMessageExecution(opts, sourceBlockchainID, message)
	Expect(err).Should(BeNil())

	return WaitForTransactionSuccess(ctx, subnet, tx.Hash())
}

func RedeemRelayerRewardsAndConfirm(
	ctx context.Context,
	subnet interfaces.SubnetTestInfo,
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
		relayerKey, subnet.EVMChainID,
	)
	Expect(err).Should(BeNil())
	tx, err := subnet.TeleporterMessenger.RedeemRelayerRewards(
		tx_opts, feeTokenAddress,
	)
	Expect(err).Should(BeNil())

	receipt := WaitForTransactionSuccess(ctx, subnet, tx.Hash())

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

	updatedRewardAmount, err := subnet.TeleporterMessenger.CheckRelayerRewardAmount(
		&bind.CallOpts{},
		relayerAddress,
		feeTokenAddress,
	)
	Expect(err).Should(BeNil())
	Expect(updatedRewardAmount.Cmp(big.NewInt(0))).Should(Equal(0))

	return receipt
}

func SendSpecifiedReceiptsAndWaitForAcceptance(
	ctx context.Context,
	sourceBlockchainID ids.ID,
	source interfaces.SubnetTestInfo,
	messageIDs [][32]byte,
	feeInfo teleportermessenger.TeleporterFeeInfo,
	allowedRelayerAddresses []common.Address,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := source.TeleporterMessenger.SendSpecifiedReceipts(
		opts, sourceBlockchainID, messageIDs, feeInfo, allowedRelayerAddresses)
	Expect(err).Should(BeNil())

	receipt := WaitForTransactionSuccess(ctx, source, tx.Hash())

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(sourceBlockchainID[:]))

	log.Info("Sending SendSpecifiedReceipts transaction",
		"sourceBlockchainID", sourceBlockchainID,
		"txHash", tx.Hash())

	return receipt, event.MessageID
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
	source interfaces.SubnetTestInfo,
	input teleportermessenger.TeleporterMessageInput,
	senderKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	data, err := teleportermessenger.PackSendCrossChainMessage(input)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, source, PrivateKeyToAddress(senderKey))

	// Send a transaction to the Teleporter contract
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   source.EVMChainID,
		Nonce:     nonce,
		To:        &teleporterContractAddress,
		Gas:       DefaultTeleporterTransactionGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     DefaultTeleporterTransactionValue,
		Data:      data,
	})

	return SignTransaction(tx, senderKey, source.EVMChainID)
}

func CreateRetryMessageExecutionTransaction(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	sourceBlockchainID ids.ID,
	message teleportermessenger.TeleporterMessage,
	senderKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	data, err := teleportermessenger.PackRetryMessageExecution(sourceBlockchainID, message)
	Expect(err).Should(BeNil())

	// TODO: replace with actual number of signers
	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(10, message.RequiredGasLimit)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo, PrivateKeyToAddress(senderKey))

	// Sign a transaction to the Teleporter contract
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   subnetInfo.EVMChainID,
		Nonce:     nonce,
		To:        &teleporterContractAddress,
		Gas:       gasLimit,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     DefaultTeleporterTransactionValue,
		Data:      data,
	})

	return SignTransaction(tx, senderKey, subnetInfo.EVMChainID)
}

// Constructs a transaction to call receiveCrossChainMessage
// Returns the signed transaction.
func CreateReceiveCrossChainMessageTransaction(
	ctx context.Context,
	signedMessage *avalancheWarp.Message,
	requiredGasLimit *big.Int,
	teleporterContractAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
) *types.Transaction {
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing receiveCrossChainMessage transaction for the destination chain")
	numSigners, err := signedMessage.Signature.NumSigners()
	Expect(err).Should(BeNil())

	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(numSigners, requiredGasLimit)
	Expect(err).Should(BeNil())

	callData, err := teleportermessenger.PackReceiveCrossChainMessage(0, PrivateKeyToAddress(senderKey))
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo, PrivateKeyToAddress(senderKey))

	destinationTx := predicateutils.NewPredicateTx(
		subnetInfo.EVMChainID,
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

	return SignTransaction(destinationTx, senderKey, subnetInfo.EVMChainID)
}

// Constructs a transaction to call addProtocolVersion
// Returns the signed transaction.
func CreateAddProtocolVersionTransaction(
	ctx context.Context,
	signedMessage *avalancheWarp.Message,
	teleporterRegistryAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
) *types.Transaction {
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing addProtocolVersion transaction for the destination chain")

	callData, err := teleporterregistry.PackAddProtocolVersion(0)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo, PrivateKeyToAddress(senderKey))

	destinationTx := predicateutils.NewPredicateTx(
		subnetInfo.EVMChainID,
		nonce,
		&teleporterRegistryAddress,
		500_000,
		gasFeeCap,
		gasTipCap,
		big.NewInt(0),
		callData,
		types.AccessList{},
		warp.ContractAddress,
		signedMessage.Bytes(),
	)

	return SignTransaction(destinationTx, senderKey, subnetInfo.EVMChainID)
}

func AddProtocolVersionAndWaitForAcceptance(
	ctx context.Context,
	network interfaces.Network,
	subnet interfaces.SubnetTestInfo,
	newTeleporterAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	unsignedMessage *avalancheWarp.UnsignedMessage,
) {
	signedWarpMsg := network.GetSignedMessage(ctx, subnet, subnet, unsignedMessage.ID())
	log.Info("Got signed warp message", "messageID", signedWarpMsg.ID())

	// Construct tx to add protocol version and send to destination chain
	signedTx := CreateAddProtocolVersionTransaction(
		ctx,
		signedWarpMsg,
		subnet.TeleporterRegistryAddress,
		senderKey,
		subnet,
	)

	curLatestVersion, err := subnet.TeleporterRegistry.LatestVersion(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	expectedLatestVersion := big.NewInt(curLatestVersion.Int64() + 1)

	// Wait for tx to be accepted, and verify events emitted
	receipt := SendTransactionAndWaitForSuccess(ctx, subnet, signedTx)
	addProtocolVersionEvent, err := GetEventFromLogs(receipt.Logs, subnet.TeleporterRegistry.ParseAddProtocolVersion)
	Expect(err).Should(BeNil())
	Expect(addProtocolVersionEvent.Version.Cmp(expectedLatestVersion)).Should(Equal(0))
	Expect(addProtocolVersionEvent.ProtocolAddress).Should(Equal(newTeleporterAddress))

	versionUpdatedEvent, err := GetEventFromLogs(receipt.Logs, subnet.TeleporterRegistry.ParseLatestVersionUpdated)
	Expect(err).Should(BeNil())
	Expect(versionUpdatedEvent.OldVersion.Cmp(curLatestVersion)).Should(Equal(0))
	Expect(versionUpdatedEvent.NewVersion.Cmp(expectedLatestVersion)).Should(Equal(0))
}

func CreateNativeTransferTransaction(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Transaction {
	fromAddress := crypto.PubkeyToAddress(fromKey.PublicKey)
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo, fromAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   subnetInfo.EVMChainID,
		Nonce:     nonce,
		To:        &recipient,
		Gas:       NativeTransferGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     amount,
	})

	return SignTransaction(tx, fromKey, subnetInfo.EVMChainID)
}

func SendNativeTransfer(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Receipt {
	tx := CreateNativeTransferTransaction(ctx, subnetInfo, fromKey, recipient, amount)
	return SendTransactionAndWaitForSuccess(ctx, subnetInfo, tx)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals success.
func sendAndWaitForTransaction(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
	success bool,
) *types.Receipt {
	err := subnetInfo.RPCClient.SendTransaction(ctx, tx)
	Expect(err).Should(BeNil())

	return waitForTransaction(ctx, subnetInfo, tx.Hash(), success)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals false.
func SendTransactionAndWaitForFailure(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
) *types.Receipt {
	return sendAndWaitForTransaction(ctx, subnetInfo, tx, false)
}

// Sends a tx, and waits for it to be mined.
// Asserts Receipt.status equals true.
func SendTransactionAndWaitForSuccess(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
) *types.Receipt {
	return sendAndWaitForTransaction(ctx, subnetInfo, tx, true)
}

func SendCrossChainMessageAndWaitForAcceptance(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	input teleportermessenger.TeleporterMessageInput,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, ids.ID) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	// Send a transaction to the Teleporter contract
	tx, err := source.TeleporterMessenger.SendCrossChainMessage(opts, input)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be accepted
	receipt := WaitForTransactionSuccess(ctx, source, tx.Hash())

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())

	log.Info("Sending SendCrossChainMessage transaction on source chain",
		"sourceChainID", source.BlockchainID,
		"destinationBlockchainID", destination.BlockchainID,
		"txHash", tx.Hash())

	return receipt, event.MessageID
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals true.
func WaitForTransactionSuccess(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	txHash common.Hash,
) *types.Receipt {
	return waitForTransaction(ctx, subnetInfo, txHash, true)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals false.
func WaitForTransactionFailure(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	txHash common.Hash,
) *types.Receipt {
	return waitForTransaction(ctx, subnetInfo, txHash, false)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals success.
func waitForTransaction(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	txHash common.Hash,
	success bool,
) *types.Receipt {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	receipt, err := WaitMined(cctx, subnetInfo.RPCClient, txHash)
	Expect(err).Should(BeNil())

	if success {
		if receipt.Status == types.ReceiptStatusFailed {
			TraceTransactionAndExit(ctx, subnetInfo, receipt.TxHash)
		}
	} else {
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusFailed))
	}
	return receipt
}

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
// Takes a tx hash instead of the full tx in the subnet-evm version of this function.
// Copied and modified from https://github.com/ava-labs/subnet-evm/blob/v0.6.0-fuji/accounts/abi/bind/util.go#L42
func WaitMined(ctx context.Context, b bind.DeployBackend, txHash common.Hash) (*types.Receipt, error) {
	queryTicker := time.NewTicker(200 * time.Millisecond)
	defer queryTicker.Stop()
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for {
		receipt, err := b.TransactionReceipt(cctx, txHash)
		if err == nil {
			return receipt, nil
		}

		if errors.Is(err, subnetEvmInterfaces.NotFound) {
			log.Debug("Transaction not yet mined")
		} else {
			log.Error("Receipt retrieval failed", "err", err)
			return nil, err
		}

		// Wait for the next round.
		select {
		case <-cctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

// Returns the first log in 'logs' that is successfully parsed by 'parser'
// Errors and prints a trace of the transaction if no log is found.
func GetEventFromLogsOrTrace[T any](
	ctx context.Context,
	receipt *types.Receipt,
	subnetInfo interfaces.SubnetTestInfo,
	parser func(log types.Log) (T, error),
) T {
	log, err := GetEventFromLogs(receipt.Logs, parser)
	if err != nil {
		TraceTransactionAndExit(ctx, subnetInfo, receipt.TxHash)
	}
	return log
}

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

// Returns true if the transaction receipt contains a ReceiptReceived log with the specified messageID
func CheckReceiptReceived(
	receipt *types.Receipt,
	messageID [32]byte,
	transactor *teleportermessenger.TeleporterMessenger,
) bool {
	for _, log := range receipt.Logs {
		event, err := transactor.ParseReceiptReceived(*log)
		if err == nil && bytes.Equal(event.MessageID[:], messageID[:]) {
			return true
		}
	}
	return false
}

func ClearReceiptQueue(
	ctx context.Context,
	network interfaces.Network,
	fundedKey *ecdsa.PrivateKey,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
) {
	Expect(network.IsExternalNetwork()).Should(BeFalse())
	outstandReceiptCount := GetOutstandingReceiptCount(source, destination.BlockchainID)
	for outstandReceiptCount.Cmp(big.NewInt(0)) != 0 {
		log.Info("Emptying receipt queue", "remainingReceipts", outstandReceiptCount.String())
		// Send message from Subnet B to Subnet A to trigger the "regular" method of delivering receipts.
		// The next message from B->A will contain the same receipts that were manually sent in the above steps,
		// but they should not be processed again on Subnet A.
		sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
			DestinationBlockchainID: destination.BlockchainID,
			DestinationAddress:      common.HexToAddress("0x1111111111111111111111111111111111111111"),
			RequiredGasLimit:        big.NewInt(1),
			FeeInfo: teleportermessenger.TeleporterFeeInfo{
				FeeTokenAddress: common.Address{},
				Amount:          big.NewInt(0),
			},
			AllowedRelayerAddresses: []common.Address{},
			Message:                 []byte{1, 2, 3, 4},
		}

		// This message will also have the same receipts as the previous message
		receipt, _ := SendCrossChainMessageAndWaitForAcceptance(
			ctx, source, destination, sendCrossChainMessageInput, fundedKey)

		// Relay message
		network.RelayMessage(ctx, receipt, source, destination, true)

		outstandReceiptCount = GetOutstandingReceiptCount(source, destination.BlockchainID)
	}
	log.Info("Receipt queue emptied")
}

func GetOutstandingReceiptCount(source interfaces.SubnetTestInfo, destinationBlockchainID ids.ID) *big.Int {
	size, err := source.TeleporterMessenger.GetReceiptQueueSize(&bind.CallOpts{}, destinationBlockchainID)
	Expect(err).Should(BeNil())
	return size
}

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
	subnetInfo interfaces.SubnetTestInfo,
	fundedAddress common.Address,
) (*big.Int, *big.Int, uint64) {
	baseFee, err := subnetInfo.RPCClient.EstimateBaseFee(ctx)
	Expect(err).Should(BeNil())

	gasTipCap, err := subnetInfo.RPCClient.SuggestGasTipCap(ctx)
	Expect(err).Should(BeNil())

	nonce, err := subnetInfo.RPCClient.NonceAt(ctx, fundedAddress, nil)
	Expect(err).Should(BeNil())

	gasFeeCap := baseFee.Mul(baseFee, big.NewInt(gasUtils.BaseFeeFactor))
	gasFeeCap.Add(gasFeeCap, big.NewInt(gasUtils.MaxPriorityFeePerGas))

	return gasFeeCap, gasTipCap, nonce
}

func PrivateKeyToAddress(k *ecdsa.PrivateKey) common.Address {
	return crypto.PubkeyToAddress(k.PublicKey)
}

// Throws a Gomega error if there is a mismatch
func CheckBalance(ctx context.Context, addr common.Address, expectedBalance *big.Int, rpcClient ethclient.Client) {
	bal, err := rpcClient.BalanceAt(ctx, addr, nil)
	Expect(err).Should(BeNil())
	ExpectBigEqual(bal, expectedBalance)
}

// Gomega will print the transaction trace and exit
func TraceTransactionAndExit(ctx context.Context, subnetInfo interfaces.SubnetTestInfo, txHash common.Hash) {
	Expect(TraceTransaction(ctx, subnetInfo, txHash)).Should(Equal(""))
}

func TraceTransaction(ctx context.Context, subnetInfo interfaces.SubnetTestInfo, txHash common.Hash) string {
	url := HttpToRPCURI(subnetInfo.NodeURIs[0], subnetInfo.BlockchainID.String())
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

func DeployContract(
	ctx context.Context,
	byteCodeFileName string,
	deployerPK *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
	abi *abi.ABI,
	constructorArgs ...interface{},
) {
	// Deploy an example ERC20 contract to be used as the source token
	byteCode, err := deploymentUtils.ExtractByteCode(byteCodeFileName)
	Expect(err).Should(BeNil())
	Expect(len(byteCode) > 0).Should(BeTrue())
	transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetInfo.EVMChainID)
	Expect(err).Should(BeNil())
	contractAddress, tx, _, err := bind.DeployContract(
		transactor,
		*abi,
		byteCode,
		subnetInfo.RPCClient,
		constructorArgs...,
	)
	Expect(err).Should(BeNil())

	// Wait for transaction, then check code was deployed
	WaitForTransactionSuccess(ctx, subnetInfo, tx.Hash())

	code, err := subnetInfo.RPCClient.CodeAt(ctx, contractAddress, nil)
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

func ERC20Approve(
	ctx context.Context,
	token *exampleerc20.ExampleERC20,
	spender common.Address,
	amount *big.Int,
	source interfaces.SubnetTestInfo,
	senderKey *ecdsa.PrivateKey,
) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := token.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())
	log.Info("Approved ERC20", "spender", spender.Hex(), "txHash", tx.Hash().Hex())

	WaitForTransactionSuccess(ctx, source, tx.Hash())
}

func DeployExampleERC20(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	source interfaces.SubnetTestInfo,
) (common.Address, *exampleerc20.ExampleERC20) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	// Deploy Mock ERC20 contract
	address, tx, token, err := exampleerc20.DeployExampleERC20(opts, source.RPCClient)
	Expect(err).Should(BeNil())
	log.Info("Deployed Mock ERC20 contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, source, tx.Hash())

	// Check that the deployer has the expected initial balance
	senderAddress := crypto.PubkeyToAddress(senderKey.PublicKey)
	balance, err := token.BalanceOf(&bind.CallOpts{}, senderAddress)
	Expect(err).Should(BeNil())
	Expect(balance).Should(Equal(ExpectedExampleERC20DeployerBalance))

	return address, token
}

func DeployExampleCrossChainMessenger(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	teleporterManager common.Address,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *examplecrosschainmessenger.ExampleCrossChainMessenger) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, exampleMessenger, err := examplecrosschainmessenger.DeployExampleCrossChainMessenger(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, exampleMessenger
}

func DeployERC20Bridge(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	teleporterManager common.Address,
	source interfaces.SubnetTestInfo,
) (common.Address, *erc20bridge.ERC20Bridge) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, erc20Bridge, err := erc20bridge.DeployERC20Bridge(
		opts,
		source.RPCClient,
		source.TeleporterRegistryAddress,
		teleporterManager,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, source, tx.Hash())

	log.Info("Deployed ERC20 Bridge contract", "address", address.Hex(), "txHash", tx.Hash().Hex())

	return address, erc20Bridge
}

func DeployBlockHashPublisher(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	subnet interfaces.SubnetTestInfo,
) (common.Address, *blockhashpublisher.BlockHashPublisher) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, publisher, err := blockhashpublisher.DeployBlockHashPublisher(
		opts, subnet.RPCClient, subnet.TeleporterRegistryAddress,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, publisher
}

func DeployBlockHashReceiver(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	teleporterManager common.Address,
	subnet interfaces.SubnetTestInfo,
	publisherAddress common.Address,
	publisherChainID [32]byte,
) (common.Address, *blockhashreceiver.BlockHashReceiver) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, receiver, err := blockhashreceiver.DeployBlockHashReceiver(
		opts,
		subnet.RPCClient,
		subnet.TeleporterRegistryAddress,
		teleporterManager,
		publisherChainID,
		publisherAddress,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx.Hash())

	return address, receiver
}

func GetTwoSubnets(network interfaces.Network) (
	interfaces.SubnetTestInfo,
	interfaces.SubnetTestInfo,
) {
	subnets := network.GetSubnetsInfo()
	Expect(len(subnets)).Should(BeNumerically(">=", 2))
	return subnets[0], subnets[1]
}

func SendExampleCrossChainMessageAndVerify(
	ctx context.Context,
	network interfaces.Network,
	source interfaces.SubnetTestInfo,
	sourceExampleMessenger *examplecrosschainmessenger.ExampleCrossChainMessenger,
	destination interfaces.SubnetTestInfo,
	destExampleMessengerAddress common.Address,
	destExampleMessenger *examplecrosschainmessenger.ExampleCrossChainMessenger,
	senderKey *ecdsa.PrivateKey,
	message string,
	expectSuccess bool,
) {
	// Call the example messenger contract on Subnet A
	optsA, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := sourceExampleMessenger.SendMessage(
		optsA,
		destination.BlockchainID,
		destExampleMessengerAddress,
		common.BigToAddress(common.Big0),
		big.NewInt(0),
		examplecrosschainmessenger.SendMessageRequiredGas,
		message,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt := WaitForTransactionSuccess(ctx, source, tx.Hash())

	event, err := GetEventFromLogs(receipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(destination.BlockchainID[:]))

	teleporterMessageID := event.MessageID

	//
	// Relay the message to the destination
	//
	receipt = network.RelayMessage(ctx, receipt, source, destination, true)

	//
	// Check Teleporter message received on the destination
	//
	delivered, err := destination.TeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeTrue())

	if expectSuccess {
		// Check that message execution was successful
		messageExecutedEvent, err := GetEventFromLogs(receipt.Logs, destination.TeleporterMessenger.ParseMessageExecuted)
		Expect(err).Should(BeNil())
		Expect(messageExecutedEvent.MessageID[:]).Should(Equal(teleporterMessageID[:]))
	} else {
		// Check that message execution failed
		messageExecutionFailedEvent, err := GetEventFromLogs(
			receipt.Logs,
			destination.TeleporterMessenger.ParseMessageExecutionFailed)
		Expect(err).Should(BeNil())
		Expect(messageExecutionFailedEvent.MessageID[:]).Should(Equal(teleporterMessageID[:]))
	}

	//
	// Verify we received the expected string
	//
	_, currMessage, err := destExampleMessenger.GetCurrentMessage(&bind.CallOpts{}, source.BlockchainID)
	Expect(err).Should(BeNil())
	if expectSuccess {
		Expect(currMessage).Should(Equal(message))
	} else {
		Expect(currMessage).ShouldNot(Equal(message))
	}
}

// Creates an Warp message that registers a Teleporter protocol version with TeleporterRegistry.
// Returns the Warp message, as well as the chain config adding the message to the list of approved
// off-chain Warp messages
func InitOffChainMessageChainConfig(
	networkID uint32,
	subnet interfaces.SubnetTestInfo,
	teleporterAddress common.Address,
	version uint64,
) (*avalancheWarp.UnsignedMessage, string) {
	unsignedMessage := CreateOffChainRegistryMessage(networkID, subnet, teleporterregistry.ProtocolRegistryEntry{
		Version:         big.NewInt(int64(version)),
		ProtocolAddress: teleporterAddress,
	})
	offChainMessage := hexutil.Encode(unsignedMessage.Bytes())
	log.Info("Adding off-chain message to Warp chain config",
		"messageID", unsignedMessage.ID(),
		"blockchainID", subnet.BlockchainID.String())

	return unsignedMessage, fmt.Sprintf(`{
    "warp-api-enabled": true, 
    "warp-off-chain-messages": ["%s"],
	"log-level": "debug",
    "eth-apis":["eth","eth-filter","net","admin","web3",
                "internal-eth","internal-blockchain","internal-transaction",
                "internal-debug","internal-account","internal-personal",
                "debug","debug-tracer","debug-file-tracer","debug-handler"]
	}`, offChainMessage)
}

// Creates an off-chain Warp message that registers a Teleporter protocol version with TeleporterRegistry
func CreateOffChainRegistryMessage(
	networkID uint32,
	subnet interfaces.SubnetTestInfo,
	entry teleporterregistry.ProtocolRegistryEntry,
) *avalancheWarp.UnsignedMessage {
	sourceAddress := []byte{}
	payloadBytes, err := teleporterregistry.PackTeleporterRegistryWarpPayload(entry, subnet.TeleporterRegistryAddress)
	Expect(err).Should(BeNil())

	addressedPayload, err := payload.NewAddressedCall(sourceAddress, payloadBytes)
	Expect(err).Should(BeNil())

	unsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		subnet.BlockchainID,
		addressedPayload.Bytes())
	Expect(err).Should(BeNil())

	return unsignedMessage
}

// Deploys a new version of Teleporter and returns its address
// Does NOT modify the global Teleporter contract address to provide greater testing flexibility.
func DeployNewTeleporterVersion(
	ctx context.Context,
	network interfaces.LocalNetwork,
	fundedKey *ecdsa.PrivateKey,
	teleporterByteCodeFile string,
) common.Address {
	contractCreationGasPrice := (&big.Int{}).Add(deploymentUtils.GetDefaultContractCreationGasPrice(), big.NewInt(1))
	teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		err := deploymentUtils.ConstructKeylessTransaction(
		teleporterByteCodeFile,
		false,
		contractCreationGasPrice,
	)
	Expect(err).Should(BeNil())

	network.DeployTeleporterContracts(
		teleporterDeployerTransaction,
		teleporterDeployerAddress,
		teleporterContractAddress,
		fundedKey,
		false)
	return teleporterContractAddress
}

// Sets the chain config in customChainConfigs for the specified subnet
func SetChainConfig(customChainConfigs map[string]string, subnet interfaces.SubnetTestInfo, chainConfig string) {
	if subnet.SubnetID == constants.PrimaryNetworkID {
		customChainConfigs[CChainPathSpecifier] = chainConfig
	} else {
		customChainConfigs[subnet.BlockchainID.String()] = chainConfig
	}
}
