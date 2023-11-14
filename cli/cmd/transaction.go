/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/hex"
	"fmt"

	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/x/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	warpPrecompileAddress = "0x0200000000000000000000000000000000000005"
)

var (
	rpcEndpoint     string
	contractAddress common.Address
	client          ethclient.Client
	instance        *teleportermessenger.TeleporterMessenger
)

// transactionCmd represents the transaction command
var transactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("transaction called")
		fmt.Println("RPC endpoint:", rpcEndpoint)
		fmt.Println("Contract address:", contractAddress)
		b, err := instance.BlockchainID(&bind.CallOpts{})
		cobra.CheckErr(err)
		fmt.Println("Instance:", hex.EncodeToString(b[:]))
		fmt.Println("TxHash", common.HexToHash(args[0]))

		receipt, err := client.TransactionReceipt(context.Background(),
			common.HexToHash(args[0]))
		cobra.CheckErr(err)
		for _, log := range receipt.Logs {
			fmt.Println("Processing log", zap.String("address", log.Address.Hex()))
			if log.Address == contractAddress {
				logger.Debug("Teleporter log", zap.Any("log", log))

				_ = FilterTeleporterEvents(log.Topics, log.Data)
			}

			if log.Address == common.HexToAddress(warpPrecompileAddress) {
				unsignedMsg, err := warp.UnpackSendWarpEventDataToMessage(log.Data)
				cobra.CheckErr(err)

				warpPayload, err := warpPayload.ParseAddressedCall(unsignedMsg.Payload)
				cobra.CheckErr(err)

				_, err = teleportermessenger.UnpackTeleporterMessage(warpPayload.Payload)
				cobra.CheckErr(err)
				// logger.Info("Parsed Teleporter message", zap.Any("message", teleporterMessage))
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(transactionCmd)
	transactionCmd.PersistentFlags().StringVar(&rpcEndpoint, "rpc", "", "RPC endpoint to connect to the node")
	address := transactionCmd.PersistentFlags().StringP("contract-address", "c", "", "Teleporter contract address")
	err := transactionCmd.MarkPersistentFlagRequired("rpc")
	cobra.CheckErr(err)
	err = transactionCmd.MarkPersistentFlagRequired("contract-address")
	cobra.CheckErr(err)
	transactionCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if rootCmd.PersistentPreRunE != nil {
			if err := rootCmd.PersistentPreRunE(cmd, args); err != nil {
				return err
			}
		}
		contractAddress = common.HexToAddress(*address)
		c, err := ethclient.Dial(rpcEndpoint)
		if err != nil {
			return err
		}

		client = c
		instance, err = teleportermessenger.NewTeleporterMessenger(contractAddress, client)
		return err
	}
}
