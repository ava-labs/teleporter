package testnet

import (
	"context"
	"crypto/ecdsa"
	"log"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/teleporter/tests/network"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	log.Println("Using Teleporter contract address:", teleporterContractAddressStr)

	subnetAInfo, err := initializeSubnetInfo(subnetAPrefix)
	if err != nil {
		return nil, err
	}
	log.Println("Using subnet A info:", subnetAInfo)

	subnetBInfo, err := initializeSubnetInfo(subnetBPrefix)
	if err != nil {
		return nil, err
	}
	log.Println("Using subnet B info:", subnetBInfo)

	fundedAddressStr := os.Getenv(userAddress)
	log.Println("Using user funded address:", fundedAddressStr)
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

func (n *testNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source network.SubnetTestInfo,
	destination network.SubnetTestInfo,
	alterMessage bool,
	expectSuccess bool) *types.Receipt {
	// Rely on a separately deployed relayer to relay the message
	time.Sleep(10 * time.Second)
	return nil
}
