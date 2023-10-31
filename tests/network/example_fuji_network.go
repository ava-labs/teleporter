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
	amplifyChainIDInt   *big.Int

	bulletinSubnetID     ids.ID
	bulletinBlockchainID ids.ID
	bulletinChainIDInt   *big.Int

	conduitSubnetID     ids.ID
	conduitBlockchainID ids.ID
	conduitChainIDInt   *big.Int

	userAddress = common.HexToAddress("0xF0Aa98fDE5f1d08F0CCEC68A7d7A7Eae31c5E9C9")
	skHex       = ""
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
	amplifyChainIDInt = big.NewInt(0).SetBytes(amplifyBlockchainID[:])

	bulletinSubnetID, err = ids.FromString("cbXsFGWSDWUYTmRXUoCirVDdQkZmUWrkQQYoVc2wUoDm8eFup")
	if err != nil {
		panic(err)
	}
	bulletinBlockchainID, err = ids.FromString("2e3RJ3ub9Pceh8fJ3HX3gZ6nSXJLvBJ9WoXLcU4nwdpZ8X2RLq")
	if err != nil {
		panic(err)
	}
	bulletinChainIDInt = big.NewInt(0).SetBytes(bulletinBlockchainID[:])

	conduitSubnetID, err = ids.FromString("wW7JVmjXp8SKrpacGzM81RBXdfcLDVY6M2DkFyArEXgtkyozK")
	if err != nil {
		panic(err)
	}
	conduitBlockchainID, err = ids.FromString("9asUA3QckLh7vGnFQiiUJGPTx8KE4nFtP8c1wTWJuP8XiWW75")
	if err != nil {
		panic(err)
	}
	conduitChainIDInt = big.NewInt(0).SetBytes(conduitBlockchainID[:])
}

// Implements Network, pointing to subnets deployed on Fuji
type FujiNetwork struct{}

func (g *FujiNetwork) GetSubnetsInfo() []utils.SubnetTestInfo {
	amplifyWSURI := "https://subnets.avax.network/amplify/testnet/rpc"
	amplifyClient, err := ethclient.Dial(amplifyWSURI)
	Expect(err).Should(BeNil())

	bulletinWSURI := "https://subnets.avax.network/bulletin/testnet/rpc"
	bulletinClient, err := ethclient.Dial(bulletinWSURI)
	Expect(err).Should(BeNil())

	conduitWSURI := "https://subnets.avax.network/conduit/testnet/rpc"
	conduitClient, err := ethclient.Dial(conduitWSURI)
	Expect(err).Should(BeNil())

	amplifyInfo := utils.SubnetTestInfo{
		SubnetID:      amplifySubnetID,
		BlockchainID:  amplifyBlockchainID,
		ChainIDInt:    amplifyChainIDInt,
		ChainWSURI:    amplifyWSURI,
		ChainWSClient: amplifyClient,
	}
	bulletinInfo := utils.SubnetTestInfo{
		SubnetID:      bulletinSubnetID,
		BlockchainID:  bulletinBlockchainID,
		ChainIDInt:    bulletinChainIDInt,
		ChainWSURI:    bulletinWSURI,
		ChainWSClient: bulletinClient,
	}
	conduitInfo := utils.SubnetTestInfo{
		SubnetID:      conduitSubnetID,
		BlockchainID:  conduitBlockchainID,
		ChainIDInt:    conduitChainIDInt,
		ChainWSURI:    conduitWSURI,
		ChainWSClient: conduitClient,
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
