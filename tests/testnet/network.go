package testnet

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	subnetevminterfaces "github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ava-labs/subnet-evm/precompile/contracts/warp"
	warpBackend "github.com/ava-labs/subnet-evm/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	teleporterregistry "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/upgrades/TeleporterRegistry"
	"github.com/ava-labs/teleporter/tests/interfaces"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"

	. "github.com/onsi/gomega"
)

const (
	subnetAPrefix                   = "subnet_a"
	subnetBPrefix                   = "subnet_b"
	subnetCPrefix                   = "subnet_c"
	cChainPrefix                    = "c_chain"
	teleporterContractAddress       = "teleporter_contract_address"
	teleporterRegistryAddressSuffix = "_teleporter_registry_address"
	subnetIDSuffix                  = "_subnet_id"
	blockchainIDSuffix              = "_chain_id"
	rpcURLSuffix                    = "_rpc_url"
	wsURLSuffix                     = "_ws_url"
	userAddress                     = "user_address"
	userPrivateKey                  = "user_private_key"

	receiveCrossChainMessageEventName      = "ReceiveCrossChainMessage"
	receiveCrossChainMessageLookBackBlocks = 500

	privateKeyHexLength = 64
)

var (
	errInvalidPrivateKeyString = errors.New("invalid private key string")
)

var _ interfaces.Network = &testNetwork{}

type testNetwork struct {
	teleporterContractAddress common.Address
	primaryNetwork            *interfaces.SubnetTestInfo
	subnetsInfo               map[ids.ID]*interfaces.SubnetTestInfo
	fundedAddress             common.Address
	fundedKey                 *ecdsa.PrivateKey
}

func initializeSubnetInfo(
	subnetPrefix string,
	teleporterContractAddress common.Address,
) (interfaces.SubnetTestInfo, error) {
	subnetIDStr := os.Getenv(subnetPrefix + subnetIDSuffix)
	subnetID, err := ids.FromString(subnetIDStr)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	blockchainIDStr := os.Getenv(subnetPrefix + blockchainIDSuffix)
	blockchainID, err := ids.FromString(blockchainIDStr)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	rpcURLStr := os.Getenv(subnetPrefix + rpcURLSuffix)
	rpcClient, err := ethclient.Dial(rpcURLStr)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	wsURLStr := os.Getenv(subnetPrefix + wsURLSuffix)
	wsClient, err := ethclient.Dial(wsURLStr)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	evmChainID, err := rpcClient.ChainID(context.Background())
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	teleporterRegistryAddress := common.HexToAddress(os.Getenv(subnetPrefix + teleporterRegistryAddressSuffix))

	teleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, rpcClient,
	)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	teleporterRegistry, err := teleporterregistry.NewTeleporterRegistry(teleporterRegistryAddress, rpcClient)
	if err != nil {
		return interfaces.SubnetTestInfo{}, err
	}

	return interfaces.SubnetTestInfo{
		SubnetID:                  subnetID,
		BlockchainID:              blockchainID,
		NodeURIs:                  []string{}, // no specific node URIs for a testnet subnet, only RPC endpoints.
		RPCClient:                 rpcClient,
		WSClient:                  wsClient,
		EVMChainID:                evmChainID,
		TeleporterRegistryAddress: teleporterRegistryAddress,
		TeleporterRegistry:        teleporterRegistry,
		TeleporterMessenger:       teleporterMessenger,
	}, nil
}

func NewTestNetwork() (*testNetwork, error) {
	teleporterContractAddressStr := os.Getenv(teleporterContractAddress)
	teleporterContractAddress := common.HexToAddress(teleporterContractAddressStr)
	log.Info("Set teleporter contract address", "teleporterContractAddress", teleporterContractAddressStr)

	subnetAInfo, err := initializeSubnetInfo(subnetAPrefix, teleporterContractAddress)
	if err != nil {
		return nil, err
	}

	subnetBInfo, err := initializeSubnetInfo(subnetBPrefix, teleporterContractAddress)
	if err != nil {
		return nil, err
	}

	subnetCInfo, err := initializeSubnetInfo(subnetCPrefix, teleporterContractAddress)
	if err != nil {
		return nil, err
	}

	cChainInfo, err := initializeSubnetInfo(cChainPrefix, teleporterContractAddress)
	if err != nil {
		return nil, err
	}
	log.Info("Set testnet subnet info", subnetAPrefix, subnetAInfo, subnetBPrefix, subnetBInfo, subnetCPrefix, subnetCInfo)

	fundedAddressStr := os.Getenv(userAddress)
	fundedKeyStr := os.Getenv(userPrivateKey)
	if len(fundedKeyStr) >= 2 && fundedKeyStr[0:2] == "0x" {
		fundedKeyStr = fundedKeyStr[2:]
	}
	if len(fundedKeyStr) != privateKeyHexLength {
		return nil, errInvalidPrivateKeyString
	}
	fundedKey, err := crypto.HexToECDSA(fundedKeyStr)
	if err != nil {
		return nil, err
	}
	log.Info("Set user funded address", "address", fundedAddressStr)

	subnetsInfo := make(map[ids.ID]*interfaces.SubnetTestInfo)
	subnetsInfo[subnetAInfo.SubnetID] = &subnetAInfo
	subnetsInfo[subnetBInfo.SubnetID] = &subnetBInfo
	subnetsInfo[subnetCInfo.SubnetID] = &subnetCInfo
	return &testNetwork{
		teleporterContractAddress: teleporterContractAddress,
		primaryNetwork:            &cChainInfo,
		subnetsInfo:               subnetsInfo,
		fundedAddress:             common.HexToAddress(fundedAddressStr),
		fundedKey:                 fundedKey,
	}, nil
}

func (n *testNetwork) GetPrimaryNetworkInfo() interfaces.SubnetTestInfo {
	return *n.primaryNetwork
}

func (n *testNetwork) GetSubnetsInfo() []interfaces.SubnetTestInfo {
	subnetsInfo := make([]interfaces.SubnetTestInfo, 0, len(n.subnetsInfo))
	for _, subnetInfo := range n.subnetsInfo {
		subnetsInfo = append(subnetsInfo, *subnetInfo)
	}
	return subnetsInfo
}

// Returns subnet info for all subnets, including the primary network
func (n *testNetwork) GetAllSubnetsInfo() []interfaces.SubnetTestInfo {
	subnets := n.GetSubnetsInfo()
	return append(subnets, n.GetPrimaryNetworkInfo())
}

func (n *testNetwork) GetTeleporterContractAddress() common.Address {
	return n.teleporterContractAddress
}

func (n *testNetwork) SetTeleporterContractAddress(newTeleporterAddress common.Address) {
	n.teleporterContractAddress = newTeleporterAddress
	subnets := n.GetAllSubnetsInfo()
	for _, subnetInfo := range subnets {
		teleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
			n.teleporterContractAddress, subnetInfo.RPCClient,
		)
		Expect(err).Should(BeNil())
		if subnetInfo.SubnetID == constants.PrimaryNetworkID {
			n.primaryNetwork.TeleporterMessenger = teleporterMessenger
		} else {
			n.subnetsInfo[subnetInfo.SubnetID].TeleporterMessenger = teleporterMessenger
		}
	}
}

func (n *testNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	return n.fundedAddress, n.fundedKey
}

func (n *testNetwork) IsExternalNetwork() bool {
	return true
}

func (n *testNetwork) SupportsIndependentRelaying() bool {
	// The test application cannot relay its own messages on testnets
	// because it can't query validators directly for their BLS signatures.
	return false
}

// For testnet messages, rely on a separately deployed relayer to relay the message.
// The implementation checks for the deliver of the given message on the destination
// within a time window of {relayWaitTime} seconds, and returns the receipt of the
// transaction that delivered the message.
func (n *testNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	expectSuccess bool) *types.Receipt {
	// Set the context to expire after 20 seconds
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

	teleporterMessageID := sendEvent.MessageID

	receipt, err := n.getMessageDeliveryTransactionReceipt(cctx, source.BlockchainID, destination, teleporterMessageID)
	Expect(err).Should(BeNil())
	Expect(receipt).ShouldNot(BeNil())
	Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
	return receipt
}

func (n *testNetwork) checkMessageDelivered(
	sourceBlockchainID ids.ID,
	destination interfaces.SubnetTestInfo,
	teleporterMessageID ids.ID) (bool, error) {
	destinationTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		n.teleporterContractAddress,
		destination.RPCClient)
	if err != nil {
		return false, err
	}

	return destinationTeleporterMessenger.MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
}

func (n *testNetwork) getMessageDeliveryTransactionReceipt(
	ctx context.Context,
	sourceBlockchainID ids.ID,
	destination interfaces.SubnetTestInfo,
	teleporterMessageID ids.ID) (*types.Receipt, error) {
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
	if currentBlockHeight > receiveCrossChainMessageLookBackBlocks {
		startBlock = currentBlockHeight - receiveCrossChainMessageLookBackBlocks
	} else {
		startBlock = 0
	}

	abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	// Get the log event of the delivery. The log must be in the last {receiveCrossChainMessageLookBackBlocks} blocks.
	logs, err := destination.RPCClient.FilterLogs(ctx, subnetevminterfaces.FilterQuery{
		FromBlock: big.NewInt(int64(startBlock)),
		Addresses: []common.Address{n.teleporterContractAddress},
		Topics: [][]common.Hash{
			{abi.Events[receiveCrossChainMessageEventName].ID},
			{common.BytesToHash(sourceBlockchainID[:])},
			{common.BytesToHash(teleporterMessageID[:])},
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

func (n *testNetwork) GetSignedMessage(
	ctx context.Context,
	source interfaces.SubnetTestInfo,
	destination interfaces.SubnetTestInfo,
	unsignedWarpMessageID ids.ID,
) *avalancheWarp.Message {
	Expect(len(source.NodeURIs)).Should(BeNumerically(">", 0))
	warpClient, err := warpBackend.NewClient(source.NodeURIs[0], source.BlockchainID.String())
	Expect(err).Should(BeNil())

	signingSubnetID := source.SubnetID
	if source.SubnetID == constants.PrimaryNetworkID {
		signingSubnetID = destination.SubnetID
	}

	signedWarpMessageBytes, err := warpClient.GetMessageAggregateSignature(
		ctx,
		unsignedWarpMessageID,
		warp.WarpDefaultQuorumNumerator,
		signingSubnetID.String(),
	)
	Expect(err).Should(BeNil())

	signedWarpMsg, err := avalancheWarp.ParseMessage(signedWarpMessageBytes)
	Expect(err).Should(BeNil())

	return signedWarpMsg
}
