package network

import (
	"context"
	"crypto/ecdsa"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/interfaces"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	. "github.com/onsi/gomega"
)

var _ Network = &FujiNetwork{}

// Amplify, Bulletin, Conduit subnet constants
var (
	teleporterContractAddress = common.HexToAddress("0x50A46AA7b2eCBe2B1AbB7df865B9A87f5eed8635")

	amplifySubnetIDStr               = "2PsShLjrFFwR51DMcAh8pyuwzLn1Ym3zRhuXLTmLCR1STk2mL6"
	amplifyBlockchainIDStr           = "2nFUad4Nw4pCgEF6MwYgGuKrzKbHJzM8wF29jeVUL41RWHgNRa"
	amplifyWSURI                     = "wss://subnets.avax.network/amplify/testnet/ws"
	amplifyRPCURI                    = "https://subnets.avax.network/amplify/testnet/rpc"
	amplifySubnetID                  ids.ID
	amplifyBlockchainID              ids.ID
	amplifyTeleporterRegistryAddress common.Address // Empty for now
	amplifyTeleporterMessenger       *teleportermessenger.TeleporterMessenger

	bulletinSubnetIDStr               = "cbXsFGWSDWUYTmRXUoCirVDdQkZmUWrkQQYoVc2wUoDm8eFup"
	bulletinBlockchainIDStr           = "2e3RJ3ub9Pceh8fJ3HX3gZ6nSXJLvBJ9WoXLcU4nwdpZ8X2RLq"
	bulletinWSURI                     = "wss://subnets.avax.network/bulletin/testnet/ws"
	bulletinRPCURI                    = "https://subnets.avax.network/bulletin/testnet/rpc"
	bulletinSubnetID                  ids.ID
	bulletinBlockchainID              ids.ID
	bulletinTeleporterRegistryAddress common.Address // Empty for now
	bulletinTeleporterMessenger       *teleportermessenger.TeleporterMessenger

	conduitSubnetIDStr               = "wW7JVmjXp8SKrpacGzM81RBXdfcLDVY6M2DkFyArEXgtkyozK"
	conduitBlockchainIDStr           = "9asUA3QckLh7vGnFQiiUJGPTx8KE4nFtP8c1wTWJuP8XiWW75"
	conduitWSURI                     = "wss://subnets.avax.network/conduit/testnet/ws"
	conduitRPCURI                    = "https://subnets.avax.network/conduit/testnet/rpc"
	conduitSubnetID                  ids.ID
	conduitBlockchainID              ids.ID
	conduitTeleporterRegistryAddress common.Address // Empty for now
	conduitTeleporterMessenger       *teleportermessenger.TeleporterMessenger

	userAddress = common.HexToAddress("")      // To be supplied by user
	skHex       = strings.TrimPrefix("", "0x") // To be supplied by user
)

func init() {
	var err error

	amplifySubnetID, err = ids.FromString(amplifySubnetIDStr)
	if err != nil {
		panic(err)
	}
	amplifyBlockchainID, err = ids.FromString(amplifyBlockchainIDStr)
	if err != nil {
		panic(err)
	}

	bulletinSubnetID, err = ids.FromString(bulletinSubnetIDStr)
	if err != nil {
		panic(err)
	}
	bulletinBlockchainID, err = ids.FromString(bulletinBlockchainIDStr)
	if err != nil {
		panic(err)
	}

	conduitSubnetID, err = ids.FromString(conduitSubnetIDStr)
	if err != nil {
		panic(err)
	}
	conduitBlockchainID, err = ids.FromString(conduitBlockchainIDStr)
	if err != nil {
		panic(err)
	}
}

// Implements Network, pointing to subnets deployed on Fuji
type FujiNetwork struct {
	amplifyInfo  utils.SubnetTestInfo
	bulletinInfo utils.SubnetTestInfo
	conduitInfo  utils.SubnetTestInfo
}

func NewFujiNetwork() *FujiNetwork {
	amplifyWSClient, err := ethclient.Dial(amplifyWSURI)
	Expect(err).Should(BeNil())
	amplifyRPCClient, err := ethclient.Dial(amplifyRPCURI)
	Expect(err).Should(BeNil())
	amplifyChainIDInt, err := amplifyRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())
	amplifyTeleporterMessenger, err = teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, amplifyRPCClient,
	)
	Expect(err).Should(BeNil())

	bulletinWSClient, err := ethclient.Dial(bulletinWSURI)
	Expect(err).Should(BeNil())
	bulletinRPCClient, err := ethclient.Dial(bulletinRPCURI)
	Expect(err).Should(BeNil())
	bulletinChainIDInt, err := bulletinRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())
	bulletinTeleporterMessenger, err = teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, bulletinRPCClient,
	)
	Expect(err).Should(BeNil())

	conduitWSClient, err := ethclient.Dial(conduitWSURI)
	Expect(err).Should(BeNil())
	conduitRPCClient, err := ethclient.Dial(conduitRPCURI)
	Expect(err).Should(BeNil())
	conduitChainIDInt, err := conduitRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())
	conduitTeleporterMessenger, err = teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, conduitRPCClient,
	)

	return &FujiNetwork{
		amplifyInfo: utils.SubnetTestInfo{
			SubnetID:                  amplifySubnetID,
			BlockchainID:              amplifyBlockchainID,
			ChainIDInt:                amplifyChainIDInt,
			ChainWSClient:             amplifyWSClient,
			ChainRPCClient:            amplifyRPCClient,
			TeleporterRegistryAddress: amplifyTeleporterRegistryAddress,
			TeleporterMessenger:       amplifyTeleporterMessenger,
		},
		bulletinInfo: utils.SubnetTestInfo{
			SubnetID:                  bulletinSubnetID,
			BlockchainID:              bulletinBlockchainID,
			ChainIDInt:                bulletinChainIDInt,
			ChainWSClient:             bulletinWSClient,
			ChainRPCClient:            bulletinRPCClient,
			TeleporterRegistryAddress: bulletinTeleporterRegistryAddress,
			TeleporterMessenger:       bulletinTeleporterMessenger,
		},
		conduitInfo: utils.SubnetTestInfo{
			SubnetID:                  conduitSubnetID,
			BlockchainID:              conduitBlockchainID,
			ChainIDInt:                conduitChainIDInt,
			ChainWSClient:             conduitWSClient,
			ChainRPCClient:            conduitRPCClient,
			TeleporterRegistryAddress: conduitTeleporterRegistryAddress,
			TeleporterMessenger:       conduitTeleporterMessenger,
		},
	}
}

func (n *FujiNetwork) CloseNetworkConnections() {
	n.amplifyInfo.ChainWSClient.Close()
	n.amplifyInfo.ChainRPCClient.Close()
	n.bulletinInfo.ChainWSClient.Close()
	n.bulletinInfo.ChainRPCClient.Close()
	n.conduitInfo.ChainWSClient.Close()
	n.conduitInfo.ChainRPCClient.Close()
}

func (n *FujiNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	return []utils.SubnetTestInfo{
		n.amplifyInfo,
		n.bulletinInfo,
		n.conduitInfo,
	}
}

func (n *FujiNetwork) GetTeleporterContractAddress() common.Address {
	return teleporterContractAddress
}

func (n *FujiNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	key, err := crypto.HexToECDSA(skHex)
	Expect(err).Should(BeNil())

	return userAddress, key
}

func (n *FujiNetwork) RelayMessage(ctx context.Context,
	sourceReceipt *types.Receipt,
	source utils.SubnetTestInfo,
	destination utils.SubnetTestInfo,
	alterMessage bool,
	expectSuccess bool) *types.Receipt {
	// Set the context to expire after 20 seconds
	var cancel context.CancelFunc
	cctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	sourceSubnetTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, source.ChainRPCClient,
	)
	Expect(err).Should(BeNil())
	destinationSubnetTeleporterMessenger, err := teleportermessenger.NewTeleporterMessenger(
		teleporterContractAddress, destination.ChainRPCClient,
	)
	Expect(err).Should(BeNil())

	// Get the Teleporter message ID from the receipt
	sendEvent, err := utils.GetEventFromLogs(
		sourceReceipt.Logs, sourceSubnetTeleporterMessenger.ParseSendCrossChainMessage,
	)
	Expect(err).Should(BeNil())
	Expect(sendEvent.DestinationChainID[:]).Should(Equal(destination.BlockchainID[:]))

	teleporterMessageID := sendEvent.Message.MessageID

	// Block until the message with the corresponding Teleporter message ID is received on the destination chain
	newHeads := make(chan *types.Header, 10)
	sub, err := destination.ChainWSClient.SubscribeNewHead(cctx, newHeads)
	Expect(err).Should(BeNil())
	defer sub.Unsubscribe()

	select {
	case <-cctx.Done():
		log.Error("Message was not relayed in time")
		Expect(true).Should(BeFalse())
	case head := <-newHeads:
		hash := head.Hash()
		logs, err := destination.ChainRPCClient.FilterLogs(cctx, interfaces.FilterQuery{
			BlockHash: &hash,
			Addresses: []common.Address{teleporterContractAddress},
		})
		Expect(err).Should(BeNil())
		if len(logs) > 0 {
			var l []*types.Log
			for _, log := range logs {
				l = append(l, &log)
			}

			receiveEvent, err := utils.GetEventFromLogs(l, destinationSubnetTeleporterMessenger.ParseReceiveCrossChainMessage)
			Expect(err).Should(BeNil())
			if receiveEvent.MessageID.Cmp(teleporterMessageID) == 0 {
				receipt, err := destination.ChainRPCClient.TransactionReceipt(cctx, receiveEvent.Raw.TxHash)
				Expect(err).Should(BeNil())
				Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful))
				return receipt
			}
		}
	}

	return nil
}
