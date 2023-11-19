package testnet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	. "github.com/onsi/gomega"
)

const (
	subnetAPrefix                   = "subnet_a"
	subnetBPrefix                   = "subnet_b"
	teleporterContractAddress       = "teleporter_contract_address"
	teleporterRegistryAddressSuffix = "_teleporter_registry_address"
	subnetIDSuffix                  = "_subnet_id"
	blockchainIDSuffix              = "_chain_id"
	rpcURLSuffix                    = "_rpc"
	wsURLSuffix                     = "_ws"
	userAddress                     = "user_address"
	userPrivateKey                  = "user_private_key"
)

var _ network.Network = &testNetwork{}

type testNetwork struct {
	teleporterContractAddress common.Address
	subnetAInfo, subnetBInfo  network.SubnetTestInfo
	fundedAddress             common.Address
	fundedKey                 *ecdsa.PrivateKey
}

func initializeSubnetInfo(subnetPrefix string) (network.SubnetTestInfo, error) {
	subnetIDStr := os.Getenv(subnetPrefix + subnetIDSuffix)
	subnetID, err := ids.FromString(subnetIDStr)
	if err != nil {
		return network.SubnetTestInfo{}, err
	}

	blockchainIDStr := os.Getenv(subnetPrefix + blockchainIDSuffix)
	blockchainID, err := ids.FromString(blockchainIDStr)
	if err != nil {
		return network.SubnetTestInfo{}, err
	}

	rpcURLStr := os.Getenv(subnetPrefix + rpcURLSuffix)
	rpcClient, err := ethclient.Dial(rpcURLStr)
	if err != nil {
		return network.SubnetTestInfo{}, err
	}

	wsURLStr := os.Getenv(subnetPrefix + wsURLSuffix)
	wsClient, err := ethclient.Dial(wsURLStr)
	if err != nil {
		return network.SubnetTestInfo{}, err
	}

	evmChainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return network.SubnetTestInfo{}, err
	}

	teleporterRegistryAddress := os.Getenv(subnetPrefix + teleporterRegistryAddressSuffix)

	return network.SubnetTestInfo{
		SubnetID:                  subnetID,
		BlockchainID:              blockchainID,
		NodeURIs:                  []string{}, // no specific node URIs for a testnet subnet, only RPC endpoints.
		RPCClient:                 rpcClient,
		WSClient:                  wsClient,
		EVMChainID:                evmChainID,
		TeleporterRegistryAddress: common.HexToAddress(teleporterRegistryAddress),
	}, nil
}

func NewTestNetwork() (*testNetwork, error) {
	teleporterContractAddressStr := os.Getenv(teleporterContractAddress)
	fmt.Println("Using Teleporter contract address:", teleporterContractAddressStr)

	subnetAInfo, err := initializeSubnetInfo(subnetAPrefix)
	if err != nil {
		return nil, err
	}
	fmt.Println("Using subnet A info:", subnetAInfo)

	subnetBInfo, err := initializeSubnetInfo(subnetBPrefix)
	if err != nil {
		return nil, err
	}
	fmt.Println("Using subnet B info:", subnetBInfo)

	fundedAddressStr := os.Getenv(userAddress)
	fmt.Println("Using user funded address:", fundedAddressStr)
	fundedKeyStr := os.Getenv(userPrivateKey)
	fundedKey, err := crypto.HexToECDSA(fundedKeyStr)
	if err != nil {
		return nil, err
	}

	return &testNetwork{
		teleporterContractAddress: common.HexToAddress(teleporterContractAddressStr),
		subnetAInfo:               subnetAInfo,
		subnetBInfo:               subnetBInfo,
		fundedAddress:             common.HexToAddress(fundedAddressStr),
		fundedKey:                 fundedKey,
	}, nil
}

func (n *testNetwork) GetSubnetInfo() (network.SubnetTestInfo, network.SubnetTestInfo) {
	return n.subnetAInfo, n.subnetBInfo
}

func (n *testNetwork) GetTeleporterContractAddress() common.Address {
	return n.teleporterContractAddress
}

func (n *testNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return n.fundedAddress, n.fundedKey
}

func (n *testNetwork) checkMessageDelivered(
	sourceBlockchainID ids.ID,
	destination network.SubnetTestInfo,
	teleporterMessageID *big.Int) (bool, error) {
	destinationTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		n.teleporterContractAddress,
		destination.RPCClient)
	if err != nil {
		return false, err
	}

	return destinationTeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, sourceBlockchainID, teleporterMessageID,
	)
}

func (n *testNetwork) getMessageDeliveryTransactionReceipt(
	ctx context.Context,
	sourceBlockchainID ids.ID,
	destination network.SubnetTestInfo,
	teleporterMessageID *big.Int) (*types.Receipt, error) {
	// Wait until the message is delivered.
	delivered := false
	var err error
	for !delivered || err != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		delivered, err = n.checkMessageDelivered(sourceBlockchainID, destination, teleporterMessageID)
		time.Sleep(time.Second)
	}

	// Get the latest block height
	currentBlockHeight, err := destination.RPCClient.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}

	var startBlock uint64
	if currentBlockHeight > 500 {
		startBlock = currentBlockHeight - 500
	} else {
		startBlock = 0
	}

	abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	// Get the log event of the delivery. The log must be in the last 500 blocks.
	logs, err := destination.RPCClient.FilterLogs(ctx, interfaces.FilterQuery{
		FromBlock: big.NewInt(int64(startBlock)),
		Addresses: []common.Address{n.teleporterContractAddress},
		Topics: [][]common.Hash{
			{abi.Events["ReceiveCrossChainMessage"].ID},
			{common.BytesToHash(sourceBlockchainID[:])},
			{common.BigToHash(teleporterMessageID)},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(logs) == 0 {
		return nil, errors.New("Failed to find ReceiveCrossChainMessage log for relayed message")
	} else if len(logs) > 1 {
		return nil, errors.New("Found multiple ReceiveCrossChainMessage logs for relayed message")
	}

	return destination.RPCClient.TransactionReceipt(ctx, logs[0].TxHash)
}

// For testnet messages, rely on a separately deployed relayer to relay the message.
// The implementation checks for the deliver of the given message on the destination
// within a time window of {relayWaitTime} seconds, and returns the receipt of the
// transaction that delivered the message.
func (n *testNetwork) RelayMessage(
	ctx context.Context,
	sourceReceipt *types.Receipt,
	source network.SubnetTestInfo,
	destination network.SubnetTestInfo,
	alterMessage bool,
	expectSuccess bool) *types.Receipt {
	// Set the context to expire after 20 seconds
	var cancel context.CancelFunc
	cctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	sourceSubnetTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		n.teleporterContractAddress, source.RPCClient,
	)
	Expect(err).Should(BeNil())

	// Get the Teleporter message ID from the receipt
	sendEvent, err := utils.GetEventFromLogs(
		sourceReceipt.Logs, sourceSubnetTeleporterMessenger.ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(sendEvent.DestinationChainID[:]).Should(Equal(destination.BlockchainID[:]))

	teleporterMessageID := sendEvent.Message.MessageID

	receipt, err := n.getMessageDeliveryTransactionReceipt(cctx, source.BlockchainID, destination, teleporterMessageID)
	Expect(err).Should(BeNil())
	return receipt
}
