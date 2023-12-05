// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"encoding/hex"

	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var messageCmd = &cobra.Command{
	Use:   "message MESSAGE_BYTES",
	Short: "Decodes hex encoded Teleporter message bytes into a TeleporterMessage struct",
	Long: `Given the hex encoded bytes of a Teleporter message, this command will decode
the bytes into a TeleporterMessage struct and print the struct fields.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		encodedMsg := args[0]
		b, err := hex.DecodeString(encodedMsg)
		cobra.CheckErr(err)

		msg, err := teleportermessenger.UnpackTeleporterMessage(b)
		cobra.CheckErr(err)
		logger.Info("Teleporter Message unpacked", zap.Any("message", msg))
		cmd.Println("Message command ran successfully")
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
}
