package network

import (
	"context"
	"crypto/ecdsa"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

var _ Network = &FujiNetwork{}

// Amplify, Bulletin, Conduit subnet constants
var (
	teleporterContractAddress = common.HexToAddress("0x50A46AA7b2eCBe2B1AbB7df865B9A87f5eed8635")

	amplifySubnetIDStr     = "2PsShLjrFFwR51DMcAh8pyuwzLn1Ym3zRhuXLTmLCR1STk2mL6"
	amplifyBlockchainIDStr = "2nFUad4Nw4pCgEF6MwYgGuKrzKbHJzM8wF29jeVUL41RWHgNRa"
	amplifyWSURI           = "wss://subnets.avax.network/amplify/testnet/ws"
	amplifyRPCURI          = "https://subnets.avax.network/amplify/testnet/rpc"
	amplifySubnetID        ids.ID
	amplifyBlockchainID    ids.ID

	bulletinSubnetIDStr     = "cbXsFGWSDWUYTmRXUoCirVDdQkZmUWrkQQYoVc2wUoDm8eFup"
	bulletinBlockchainIDStr = "2e3RJ3ub9Pceh8fJ3HX3gZ6nSXJLvBJ9WoXLcU4nwdpZ8X2RLq"
	bulletinWSURI           = "wss://subnets.avax.network/bulletin/testnet/ws"
	bulletinRPCURI          = "https://subnets.avax.network/bulletin/testnet/rpc"
	bulletinSubnetID        ids.ID
	bulletinBlockchainID    ids.ID

	conduitSubnetIDStr     = "wW7JVmjXp8SKrpacGzM81RBXdfcLDVY6M2DkFyArEXgtkyozK"
	conduitBlockchainIDStr = "9asUA3QckLh7vGnFQiiUJGPTx8KE4nFtP8c1wTWJuP8XiWW75"
	conduitWSURI           = "wss://subnets.avax.network/conduit/testnet/ws"
	conduitRPCURI          = "https://subnets.avax.network/conduit/testnet/rpc"
	conduitSubnetID        ids.ID
	conduitBlockchainID    ids.ID

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
type FujiNetwork struct{}

func (n *FujiNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	amplifyWSClient, err := ethclient.Dial(amplifyWSURI)
	Expect(err).Should(BeNil())
	amplifyRPCClient, err := ethclient.Dial(amplifyRPCURI)
	Expect(err).Should(BeNil())
	amplifyChainIDInt, err := amplifyRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	bulletinWSClient, err := ethclient.Dial(bulletinWSURI)
	Expect(err).Should(BeNil())
	bulletinRPCClient, err := ethclient.Dial(bulletinRPCURI)
	Expect(err).Should(BeNil())
	bulletinChainIDInt, err := bulletinRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	conduitWSClient, err := ethclient.Dial(conduitWSURI)
	Expect(err).Should(BeNil())
	conduitRPCClient, err := ethclient.Dial(conduitRPCURI)
	Expect(err).Should(BeNil())
	conduitChainIDInt, err := conduitRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	amplifyInfo := utils.SubnetTestInfo{
		SubnetID:       amplifySubnetID,
		BlockchainID:   amplifyBlockchainID,
		ChainIDInt:     amplifyChainIDInt,
		ChainWSClient:  amplifyWSClient,
		ChainRPCClient: amplifyRPCClient,
	}
	bulletinInfo := utils.SubnetTestInfo{
		SubnetID:       bulletinSubnetID,
		BlockchainID:   bulletinBlockchainID,
		ChainIDInt:     bulletinChainIDInt,
		ChainWSClient:  bulletinWSClient,
		ChainRPCClient: bulletinRPCClient,
	}
	conduitInfo := utils.SubnetTestInfo{
		SubnetID:       conduitSubnetID,
		BlockchainID:   conduitBlockchainID,
		ChainIDInt:     conduitChainIDInt,
		ChainWSClient:  conduitWSClient,
		ChainRPCClient: conduitRPCClient,
	}

	return []utils.SubnetTestInfo{
		amplifyInfo,
		bulletinInfo,
		conduitInfo,
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
	expectSuccess bool) *types.Receipt {

	// Rely on a separately deployed relayer to relay the message
	time.Sleep(20 * time.Second)
	return nil
}
