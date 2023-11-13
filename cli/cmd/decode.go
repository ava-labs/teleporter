/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ava-labs/subnet-evm/accounts/abi"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/spf13/cobra"
)

var (
	teleporterABI *abi.ABI
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes Teleporter data into more readable formats",
	Long: `A decoder that can decode Teleporter data into more readable formats.
	Currently, the decoder can decode Teleporter messages and events.`,
}

func init() {
	rootCmd.AddCommand(decodeCmd)
	abi, err := teleportermessenger.TeleporterMessengerMetaData.GetAbi()
	cobra.CheckErr(err)
	teleporterABI = abi
}
