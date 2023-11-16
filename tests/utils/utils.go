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
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/params"
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/subnet-evm/x/warp"
	examplecrosschainmessenger "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/ExampleMessenger/ExampleCrossChainMessenger"
	exampleerc20 "github.com/ava-labs/teleporter/abi-bindings/go/Mocks/ExampleERC20"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
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
	source network.SubnetTestInfo,
	destination network.SubnetTestInfo,
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
	receipt, err := bind.WaitMined(ctx, source.RPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	fmt.Println("sent tx", txn.Hash())

	log.Info("Sent SendCrossChainMessage transaction on source chain",
		"sourceChainID", source.BlockchainID,
		"destinationChainID", destination.BlockchainID,
		"txHash", txn.Hash())

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, transactor.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationChainID[:]).Should(Equal(destination.BlockchainID[:]))

	return receipt, event.Message.MessageID
}

func SendAddFeeAmountAndWaitForAcceptance(
	ctx context.Context,
	source network.SubnetTestInfo,
	destination network.SubnetTestInfo,
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
	receipt, err := bind.WaitMined(ctx, source.RPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	addFeeAmountEvent, err := GetEventFromLogs(receipt.Logs, transactor.ParseAddFeeAmount)
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
	subnet network.SubnetTestInfo,
	message teleportermessenger.TeleporterMessage,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) *types.Receipt {
	opts := CreateTransactorOpts(ctx, subnet, fundedAddress, fundedKey)

	txn, err := transactor.RetryMessageExecution(opts, originChainID, message)
	Expect(err).Should(BeNil())

	receipt, err := bind.WaitMined(ctx, subnet.RPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	return receipt
}

func SendSpecifiedReceiptsAndWaitForAcceptance(
	ctx context.Context,
	originChainID ids.ID,
	source network.SubnetTestInfo,
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

	receipt, err := bind.WaitMined(ctx, source.RPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, transactor.ParseSendCrossChainMessage)
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

func CreateTransactorOpts(ctx context.Context, network network.SubnetTestInfo, fundedAddress common.Address, fundedKey *ecdsa.PrivateKey) *bind.TransactOpts {
	// set up parameters
	transactor, err := bind.NewKeyedTransactorWithChainID(
		fundedKey, network.EVMChainID)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, network.RPCClient, fundedAddress)

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
	source network.SubnetTestInfo,
	input teleportermessenger.TeleporterMessageInput,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	data, err := teleportermessenger.PackSendCrossChainMessage(input)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, source.RPCClient, fundedAddress)

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

	return SignTransaction(tx, fundedKey, source.EVMChainID)
}

func CreateRetryMessageExecutionTransaction(
	ctx context.Context,
	subnetInfo network.SubnetTestInfo,
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

	// TODO: replace with actual number of signers
	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(10, message.RequiredGasLimit)
	Expect(err).Should(BeNil())

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo.RPCClient, fundedAddress)

	// Send a transaction to the Teleporter contract
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

	return SignTransaction(tx, fundedKey, subnetInfo.EVMChainID)
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
	subnetInfo network.SubnetTestInfo,
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

	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, subnetInfo.RPCClient, fundedAddress)

	if alterMessage {
		alterTeleporterMessage(signedMessage)
	}

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

	return SignTransaction(destinationTx, fundedKey, subnetInfo.EVMChainID)
}

func CreateNativeTransferTransaction(
	ctx context.Context,
	network network.SubnetTestInfo,
	fromAddress common.Address,
	fromKey *ecdsa.PrivateKey,
	recipient common.Address,
	amount *big.Int,
) *types.Transaction {
	gasFeeCap, gasTipCap, nonce := CalculateTxParams(ctx, network.RPCClient, fromAddress)

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   network.EVMChainID,
		Nonce:     nonce,
		To:        &recipient,
		Gas:       NativeTransferGas,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     amount,
	})

	return SignTransaction(tx, fromKey, network.EVMChainID)
}

func WaitForTransaction(ctx context.Context, txHash common.Hash, rpcClient ethclient.Client) *types.Receipt {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Loop until we find the transaction or time out
	for {
		receipt, err := rpcClient.TransactionReceipt(cctx, txHash)
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

func DeployContract(ctx context.Context, byteCodeFileName string, deployerPK *ecdsa.PrivateKey, subnetInfo network.SubnetTestInfo, abi *abi.ABI, constructorArgs ...interface{}) common.Address {
	// Deploy an example ERC20 contract to be used as the source token
	byteCode, err := deploymentUtils.ExtractByteCode(byteCodeFileName)
	Expect(err).Should(BeNil())
	Expect(len(byteCode) > 0).Should(BeTrue())
	transactor, err := bind.NewKeyedTransactorWithChainID(deployerPK, subnetInfo.EVMChainID)
	Expect(err).Should(BeNil())
	contractAddress, tx, _, err := bind.DeployContract(transactor, *abi, byteCode, subnetInfo.RPCClient, constructorArgs...)
	Expect(err).Should(BeNil())

	// Wait for transaction, then check code was deployed
	WaitForTransaction(ctx, tx.Hash(), subnetInfo.RPCClient)
	code, err := subnetInfo.RPCClient.CodeAt(ctx, contractAddress, nil)
	Expect(err).Should(BeNil())
	Expect(len(code)).Should(BeNumerically(">", 2)) // 0x is an EOA, contract returns the bytecode

	return contractAddress
}

func DeployMockToken(ctx context.Context, fundedAddress common.Address, fundedKey *ecdsa.PrivateKey, source network.SubnetTestInfo) (common.Address, *exampleerc20.ExampleERC20) {
	opts := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)

	// Deploy Mock ERC20 contract
	address, txn, mockToken, err := exampleerc20.DeployExampleERC20(opts, source.RPCClient)
	Expect(err).Should(BeNil())
	log.Info("Deployed Mock ERC20 contract", "address", address.Hex(), "txHash", txn.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, source.RPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	return address, mockToken
}

func DeployExampleCrossChainMessenger(ctx context.Context, fundedAddress common.Address, fundedKey *ecdsa.PrivateKey, source network.SubnetTestInfo) (common.Address, *examplecrosschainmessenger.ExampleCrossChainMessenger) {
	optsA := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)

	// Deploy the example messenger contract
	address, tx, exampleMessenger, err := examplecrosschainmessenger.DeployExampleCrossChainMessenger(optsA, source.RPCClient, source.TeleporterRegistryAddress)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(ctx, source.RPCClient, tx)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))

	return address, exampleMessenger
}

func ERC20Approve(ctx context.Context,
	fundedAddress common.Address,
	fundedKey *ecdsa.PrivateKey,
	token *exampleerc20.ExampleERC20,
	spender common.Address,
	amount *big.Int,
	source network.SubnetTestInfo) {
	opts := CreateTransactorOpts(ctx, source, fundedAddress, fundedKey)
	txn, err := token.Approve(opts, spender, amount)
	Expect(err).Should(BeNil())
	log.Info("Approved Mock ERC20", "spender", spender.Hex(), "txHash", txn.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, source.RPCClient, txn)
	Expect(err).Should(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
}

func alterTeleporterMessage(signedMessage *avalancheWarp.Message) {
	warpMsgPayload, err := warpPayload.ParseAddressedCall(signedMessage.UnsignedMessage.Payload)
	Expect(err).Should(BeNil())

	teleporterMessage, err := teleportermessenger.UnpackTeleporterMessage(warpMsgPayload.Payload)
	Expect(err).Should(BeNil())
	// Alter the message
	teleporterMessage.Message[0] = ^teleporterMessage.Message[0]

	// Pack the teleporter message
	teleporterMessageBytes, err := teleportermessenger.PackTeleporterMessage(*teleporterMessage)
	Expect(err).Should(BeNil())

	payload, err := warpPayload.NewAddressedCall(warpMsgPayload.SourceAddress, teleporterMessageBytes)
	Expect(err).Should(BeNil())

	signedMessage.UnsignedMessage.Payload = payload.Bytes()

	signedMessage.Initialize()
}
