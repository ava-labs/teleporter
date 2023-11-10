package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ava-labs/subnet-evm/ethclient"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.String("rpcEndpoint", "", "RPC endpoint to connect to the node")
	pflag.String("contractAddress", "", "Teleporter contract address")
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}

	rpcEndpoint := viper.GetString("rpcEndpoint")
	contractAddress := viper.GetString("contractAddress")

	if rpcEndpoint == "" || contractAddress == "" {
		log.Fatal("Please provide both RPC endpoint and contract address")
	}
	fmt.Println("RPC endpoint:", rpcEndpoint)
	fmt.Println("Contract address:", contractAddress)

	ethClient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		panic(err)
	}
	_, err = teleportermessenger.NewTeleporterMessenger(common.HexToAddress(contractAddress), ethClient)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Teleporter event signature: ")
	eventSignature, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read event signature")
	}
	eventSignature = strings.TrimRight(eventSignature, "\n")

	fmt.Println("Event signature:", eventSignature)

	abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	event, err := abi.EventByID(common.HexToHash(eventSignature))
	if err != nil {
		panic(err)
	}

	fmt.Println("Event name:", event.Name)

}

// package main

// import (
// 	"encoding/hex"
// 	"fmt"
// 	"log"

// 	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
// 	"github.com/ava-labs/subnet-evm/ethclient"
// 	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
// 	"github.com/ethereum/go-ethereum/common"
// )

// func main() {
// 	client, err := ethclient.Dial("https://subnets.avax.network/firedrillt/testnet/rpc")
// 	if err != nil {
// 		panic(err)
// 	}

// 	contractAddress := common.HexToAddress("0xfEabB3b3F4eEaE6B5769507a5E6B808704E5c626")
// 	instance, err := teleportermessenger.NewTeleporterMessenger(contractAddress, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	tokenAddress := common.HexToAddress("0xdbfAb16E4eA8b872136275Dd6c06837c88EE1730")
// 	relayerAddress := common.HexToAddress("0xc71a61a815e49d16c425482a342a367cd42e38a6")
// 	myNumber, err := instance.RelayerRewardAmounts(&bind.CallOpts{}, relayerAddress, tokenAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("myNumber:", myNumber)

// 	// abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	// event, err := abi.EventByID(common.HexToHash("0x9cf033c0096a005bf84a177dba966f67ed266771072eb06b1280a9d394747b1d"))
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(event.Name)
// 	// fmt.Println("Arguments", event.Inputs)
// 	// fmt.Println(event.String())

// 	// fmt.Println(event.OriginChainID)
// 	// fmt.Println(event.MessageID)
// 	// fmt.Println(event.Deliverer)
// 	// fmt.Println("Hex version", hex.EncodeToString(myNumber[:]))

// 	// myString, err := instance.MyString(&bind.CallOpts{})
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println("myString:", myString)
// 	// hexStrings := []string{
// 	// 	"000000000000000000000000c71a61a815e49d16c425482a342a367cd42e38a6",
// 	// 	"0000000000000000000000000000000000000000000000000000000000000040",
// 	// 	"0000000000000000000000000000000000000000000000000000000000000251",
// 	// 	"00000000000000000000000063e0d024a6989c5c5335fd4ab77ff7db1aed56ef",
// 	// 	"000000000000000000000000f5a6083fa7e0494e757bf4a4a7f4961890bfed12",
// 	// 	"0000000000000000000000000000000000000000000000000000000000030d40",
// 	// 	"00000000000000000000000000000000000000000000000000000000000000e0",
// 	// 	"0000000000000000000000000000000000000000000000000000000000000100",
// 	// 	"00000000000000000000000000000000000000000000000000000000000001a0",
// 	// 	"0000000000000000000000000000000000000000000000000000000000000000",
// 	// 	"0000000000000000000000000000000000000000000000000000000000000002",
// 	// 	"000000000000000000000000000000000000000000000000000000000000024b",
// 	// 	"000000000000000000000000c71a61a815e49d16c425482a342a367cd42e38a6",
// 	// 	"000000000000000000000000000000000000000000000000000000000000024c",
// 	// 	"000000000000000000000000c71a61a815e49d16c425482a342a367cd42e38a6",
// 	// 	"00000000000000000000000000000000000000000000000000000000000000c0",
// 	// 	"0000000000000000000000000000000000000000000000000000000000000001",
// 	// 	"0000000000000000000000000000000000000000000000000000000000000040",
// 	// 	"0000000000000000000000000000000000000000000000000000000000000060",
// 	// 	"000000000000000000000000dbfab16e4ea8b872136275dd6c06837c88ee1730",
// 	// 	"00000000000000000000000013987c904c4ea4e2ee78823fd02ff14da19c3741",
// 	// 	"0000000000000000000000000000000000000000000000000000000000068910",
// 	// }

// 	// // Concatenate the byte arrays
// 	// var result []byte
// 	// for _, hexStr := range hexStrings {
// 	// 	bytes, err := hex.DecodeString(hexStr)
// 	// 	if err != nil {
// 	// 		fmt.Printf("Error decoding hex string: %v\n", err)
// 	// 		return
// 	// 	}
// 	// 	result = append(result, bytes...)
// 	// }
// 	// log := types.Log{
// 	// 	Address: common.HexToAddress("0xfeabb3b3f4eeae6b5769507a5e6b808704e5c626"),
// 	// 	Topics: []common.Hash{
// 	// 		common.HexToHash("0x9cf033c0096a005bf84a177dba966f67ed266771072eb06b1280a9d394747b1d"),
// 	// 		common.HexToHash("0x39bec71f2a981a5d5dc824471788df002194dee5db051c7902b51960c466ca81"),
// 	// 		common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000251"),
// 	// 		common.HexToHash("0x00000000000000000000000066fc895081cf815aa25bf4853cc377dfc8cc5871"),
// 	// 	},
// 	// 	Data: result,
// 	// }
// 	// receive, err := instance.ParseReceiveCrossChainMessage(log)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }

// 	selector, err := hex.DecodeString("ccb5f809")
// 	if err != nil {
// 		panic(err)
// 	}
// 	abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
// 	if err != nil {
// 		panic(err)
// 	}
// 	method, err := abi.MethodById(selector)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(method.Name)
// 	// .MethodById(selector)
// }
