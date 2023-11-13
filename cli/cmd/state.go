/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var (
	contractAddressStr string
	rpcEndpoint        string
	contractAddress    common.Address
)

// stateCmd represents the state command
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgs: []string{"blockchainID", "latestMessageID"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("state called", args)
		contractAddress = common.HexToAddress(contractAddressStr)
	},
}

func init() {
	rootCmd.AddCommand(stateCmd)
	stateCmd.PersistentFlags().StringVarP(&contractAddressStr, "contract-address", "c", "", "The address of the Teleporter contract")
	stateCmd.PersistentFlags().StringVar(&rpcEndpoint, "rpc", "", "A RPC endpoint to query for on chain state")
	stateCmd.MarkPersistentFlagRequired("contract-address")
	stateCmd.MarkPersistentFlagRequired("rpc")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
