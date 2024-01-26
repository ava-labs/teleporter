// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	deploymentUtils "github.com/ava-labs/teleporter/utils/deployment-utils"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	if len(os.Args) < 2 {
		log.Panic("Invalid argument count. Must provide at least one argument to specify command type.")
	}
	commandType := os.Args[1]

	switch commandType {
	case "constructKeylessTx":
		// Get the byte code of the teleporter contract to be deployed.
		if len(os.Args) != 3 {
			log.Panic("Invalid argument count. Must provide JSON file containing contract bytecode.")
		}
		_, _, _, err := deploymentUtils.ConstructKeylessTransaction(
			os.Args[2],
			true,
			deploymentUtils.GetDefaultContractCreationGasPrice(),
		)
		if err != nil {
			log.Panic("Failed to construct keyless transaction.", err)
		}
	case "deriveContractAddress":
		// Get the byte code of the teleporter contract to be deployed.
		if len(os.Args) != 4 {
			log.Panic("Invalid argument count. Must provide address and nonce.")
		}

		deployerAddress := common.HexToAddress(os.Args[2])
		nonce, err := strconv.ParseUint(os.Args[3], 10, 64)
		if err != nil {
			log.Panic("Failed to parse nonce as uint", err)
		}

		resultAddress, err := deploymentUtils.DeriveEVMContractAddress(deployerAddress, nonce)
		if err != nil {
			log.Panic("Failed to derive contract address.", err)
		}
		fmt.Println(resultAddress.Hex())
	default:
		log.Panic("Invalid command type. Supported options are \"constructKeylessTx\" and \"deriveContractAddress\".")
	}
}
