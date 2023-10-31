package network

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/teleporter/tests/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	. "github.com/onsi/gomega"
)

// Amplify, Bulletin, Conduit subnet constants
var (
	teleporterContractAddress = common.HexToAddress("0x50A46AA7b2eCBe2B1AbB7df865B9A87f5eed8635")

	amplifySubnetID     ids.ID
	amplifyBlockchainID ids.ID

	bulletinSubnetID     ids.ID
	bulletinBlockchainID ids.ID

	conduitSubnetID     ids.ID
	conduitBlockchainID ids.ID

	userAddress = common.HexToAddress("0xF0Aa98fDE5f1d08F0CCEC68A7d7A7Eae31c5E9C9")
	skHex       = "" // To be supplied by user
)

func init() {
	var err error

	amplifySubnetID, err = ids.FromString("2PsShLjrFFwR51DMcAh8pyuwzLn1Ym3zRhuXLTmLCR1STk2mL6")
	if err != nil {
		panic(err)
	}
	amplifyBlockchainID, err = ids.FromString("2nFUad4Nw4pCgEF6MwYgGuKrzKbHJzM8wF29jeVUL41RWHgNRa")
	if err != nil {
		panic(err)
	}

	bulletinSubnetID, err = ids.FromString("cbXsFGWSDWUYTmRXUoCirVDdQkZmUWrkQQYoVc2wUoDm8eFup")
	if err != nil {
		panic(err)
	}
	bulletinBlockchainID, err = ids.FromString("2e3RJ3ub9Pceh8fJ3HX3gZ6nSXJLvBJ9WoXLcU4nwdpZ8X2RLq")
	if err != nil {
		panic(err)
	}

	conduitSubnetID, err = ids.FromString("wW7JVmjXp8SKrpacGzM81RBXdfcLDVY6M2DkFyArEXgtkyozK")
	if err != nil {
		panic(err)
	}
	conduitBlockchainID, err = ids.FromString("9asUA3QckLh7vGnFQiiUJGPTx8KE4nFtP8c1wTWJuP8XiWW75")
	if err != nil {
		panic(err)
	}
}

// Implements Network, pointing to subnets deployed on Fuji
type FujiNetwork struct{}

func (g *FujiNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	amplifyWSURI := "wss://subnets.avax.network/amplify/testnet/ws"
	amplifyWSClient, err := ethclient.Dial(amplifyWSURI)
	Expect(err).Should(BeNil())
	amplifyRPCURI := "https://subnets.avax.network/amplify/testnet/rpc"
	amplifyRPCClient, err := ethclient.Dial(amplifyRPCURI)
	Expect(err).Should(BeNil())
	amplifyChainIDInt, err := amplifyRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	bulletinWSURI := "wss://subnets.avax.network/bulletin/testnet/ws"
	bulletinWSClient, err := ethclient.Dial(bulletinWSURI)
	Expect(err).Should(BeNil())
	bulletinRPCURI := "https://subnets.avax.network/bulletin/testnet/rpc"
	bulletinRPCClient, err := ethclient.Dial(bulletinRPCURI)
	Expect(err).Should(BeNil())
	bulletinChainIDInt, err := bulletinRPCClient.ChainID(context.Background())
	Expect(err).Should(BeNil())

	conduitWSURI := "wss://subnets.avax.network/conduit/testnet/ws"
	conduitWSClient, err := ethclient.Dial(conduitWSURI)
	Expect(err).Should(BeNil())
	conduitRPCURI := "https://subnets.avax.network/conduit/testnet/rpc"
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

func (g *FujiNetwork) GetTeleporterContractAddress() common.Address {
	return teleporterContractAddress
}

func (g *FujiNetwork) GetFundedAccountInfo() (common.Address, *ecdsa.PrivateKey) {
	key, err := crypto.HexToECDSA(skHex)
	Expect(err).Should(BeNil())

	return userAddress, key
}

func (g *FujiNetwork) RelayMessage(ctx context.Context,
	sourceBlockHash common.Hash,
	sourceBlockNumber *big.Int,
	source utils.SubnetTestInfo,
	destination utils.SubnetTestInfo) *types.Receipt {

	// Rely on a separately deployed relayer to relay the message
	time.Sleep(20 * time.Second)
	return nil
}
