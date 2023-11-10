package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
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
	teleporter, err := teleportermessenger.NewTeleporterMessenger(common.HexToAddress(contractAddress), ethClient)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Teleporter event signature: ")
	_, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal("Failed to read event signature")
	}

	chainID, err := teleporter.BlockchainID(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(chainID[:]))

}
