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

	erc20bridge "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ERC20Bridge/ERC20Bridge"
	examplecrosschainmessenger "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/ExampleMessenger/ExampleCrossChainMessenger"
	blockhashpublisher "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/VerifiedBlockHash/BlockHashPublisher"
	blockhashreceiver "github.com/ava-labs/teleporter/abi-bindings/go/CrossChainApplications/examples/VerifiedBlockHash/BlockHashReceiver"
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
	predicateutils "github.com/ava-labs/subnet-evm/predicate"
	"github.com/ava-labs/subnet-evm/rpc"
	"github.com/ava-labs/subnet-evm/x/warp"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ethereum/go-ethereum/common"
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

//
// Test utility functions
//

func SendAddFeeAmountAndWaitForAcceptance(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	messageID *big.Int,
	amount *big.Int,
	feeContractAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	transactor *teleportermessenger.TeleporterMessenger,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := transactor.AddFeeAmount(opts, destination.BlockchainID, messageID, feeContractAddress, amount)
	Expect(err).Should(BeNil())

	receipt := WaitForTransactionSuccess(ctx, source, tx)

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
	subnet interfaces.SubnetTestInfo,
	message teleportermessenger.TeleporterMessage,
	senderKey *ecdsa.PrivateKey,
) *types.Receipt {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := subnet.TeleporterMessenger.RetryMessageExecution(opts, originChainID, message)
	Expect(err).Should(BeNil())

	return WaitForTransactionSuccess(ctx, subnet, tx)
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

	receipt := WaitForTransactionSuccess(ctx, subnet, tx)

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
	originChainID ids.ID,
	source interfaces.SubnetTestInfo,
	messageIDs []*big.Int,
	feeInfo teleportermessenger.TeleporterFeeInfo,
	allowedRelayerAddresses []common.Address,
	senderKey *ecdsa.PrivateKey,
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	tx, err := source.TeleporterMessenger.SendSpecifiedReceipts(
		opts, originChainID, messageIDs, feeInfo, allowedRelayerAddresses)
	Expect(err).Should(BeNil())

	receipt := WaitForTransactionSuccess(ctx, source, tx)

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(event.DestinationBlockchainID[:]).Should(Equal(originChainID[:]))

	log.Info("Sending SendSpecifiedReceipts transaction",
		"originChainID", originChainID,
		"txHash", tx.Hash())

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
	originChainID ids.ID,
	message teleportermessenger.TeleporterMessage,
	senderKey *ecdsa.PrivateKey,
	teleporterContractAddress common.Address,
) *types.Transaction {
	data, err := teleportermessenger.PackRetryMessageExecution(originChainID, message)
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
	warpMessageBytes []byte,
	requiredGasLimit *big.Int,
	teleporterContractAddress common.Address,
	senderKey *ecdsa.PrivateKey,
	subnetInfo interfaces.SubnetTestInfo,
) *types.Transaction {
	// Construct the transaction to send the Warp message to the destination chain
	log.Info("Constructing transaction for the destination chain")
	signedMessage, err := avalancheWarp.ParseMessage(warpMessageBytes)
	Expect(err).Should(BeNil())

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

	return waitForTransaction(ctx, subnetInfo, tx, success)
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
) (*types.Receipt, *big.Int) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())

	// Send a transaction to the Teleporter contract
	tx, err := source.TeleporterMessenger.SendCrossChainMessage(opts, input)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be accepted
	receipt := WaitForTransactionSuccess(ctx, source, tx)

	// Check the transaction logs for the SendCrossChainMessage event emitted by the Teleporter contract
	event, err := GetEventFromLogs(receipt.Logs, source.TeleporterMessenger.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())

	log.Info("Sending SendCrossChainMessage transaction on source chain",
		"sourceChainID", source.BlockchainID,
		"destinationChainID", destination.BlockchainID,
		"txHash", tx.Hash())

	return receipt, event.Message.MessageID
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals true.
func WaitForTransactionSuccess(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
) *types.Receipt {
	return waitForTransaction(ctx, subnetInfo, tx, true)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals false.
func WaitForTransactionFailure(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
) *types.Receipt {
	return waitForTransaction(ctx, subnetInfo, tx, false)
}

// Waits for a transaction to be mined.
// Asserts Receipt.status equals success.
func waitForTransaction(
	ctx context.Context,
	subnetInfo interfaces.SubnetTestInfo,
	tx *types.Transaction,
	success bool,
) *types.Receipt {
	cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	receipt, err := bind.WaitMined(cctx, subnetInfo.RPCClient, tx)
	Expect(err).Should(BeNil())

	if success {
		if receipt.Status == types.ReceiptStatusFailed {
			fmt.Println(TraceTransaction(ctx, subnetInfo, tx))
		}
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	} else {
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusFailed))
	}
	return receipt
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

func TraceTransaction(ctx context.Context, subnetInfo interfaces.SubnetTestInfo, tx *types.Transaction) string {
	url := HttpToRPCURI(subnetInfo.NodeURIs[0], subnetInfo.BlockchainID.String())
	rpcClient, err := rpc.DialContext(ctx, url)
	Expect(err).Should(BeNil())
	defer rpcClient.Close()

	var result interface{}
	ct := "callTracer"
	err = rpcClient.Call(&result, "debug_traceTransaction", tx.Hash().String(), tracers.TraceConfig{Tracer: &ct})
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
	WaitForTransactionSuccess(ctx, subnetInfo, tx)

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

	WaitForTransactionSuccess(ctx, source, tx)
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
	WaitForTransactionSuccess(ctx, source, tx)

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
	subnet interfaces.SubnetTestInfo,
) (common.Address, *examplecrosschainmessenger.ExampleCrossChainMessenger) {
	opts, err := bind.NewKeyedTransactorWithChainID(
		senderKey, subnet.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, exampleMessenger, err := examplecrosschainmessenger.DeployExampleCrossChainMessenger(
		opts, subnet.RPCClient, subnet.TeleporterRegistryAddress,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx)

	return address, exampleMessenger
}

func DeployERC20Bridge(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
	source interfaces.SubnetTestInfo,
) (common.Address, *erc20bridge.ERC20Bridge) {
	opts, err := bind.NewKeyedTransactorWithChainID(senderKey, source.EVMChainID)
	Expect(err).Should(BeNil())
	address, tx, erc20Bridge, err := erc20bridge.DeployERC20Bridge(
		opts, source.RPCClient, source.TeleporterRegistryAddress,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, source, tx)

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
	WaitForTransactionSuccess(ctx, subnet, tx)

	return address, publisher
}

func DeployBlockHashReceiver(
	ctx context.Context,
	senderKey *ecdsa.PrivateKey,
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
		publisherChainID,
		publisherAddress,
	)
	Expect(err).Should(BeNil())

	// Wait for the transaction to be mined
	WaitForTransactionSuccess(ctx, subnet, tx)

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
