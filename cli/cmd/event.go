// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cmd

import (
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	topicArgs []string
	data      []byte
)

var eventCmd = &cobra.Command{
	Use:   "event",
	Short: "Parses a Teleporter log's topics and data",
	Long: `Given the topics and data of a Teleporter log, parses the log into
	the corresponding Teleporter event. Topics are represented by a hash, and data
	is the hex encoding of the bytes. For example:

	teleporter-cli event --topics topic1, topic2 --data 0x1234`,
	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		var topics []common.Hash
		for _, topic := range topicArgs {
			topics = append(topics, common.HexToHash(topic))
		}

		event, err := teleporterABI.EventByID(topics[0])
		cobra.CheckErr(err)

		out, err := teleportermessenger.FilterTeleporterEvents(topics, data, event.Name)
		cobra.CheckErr(err)
		logger.Info("Parsed Teleporter event", zap.String("name", event.Name), zap.Any("event", out))
	},
}

func init() {
	rootCmd.AddCommand(eventCmd)
	eventCmd.PersistentFlags().StringSliceVar(&topicArgs, "topics", []string{}, "The topics of the event")
	eventCmd.Flags().BytesHexVar(&data, "data", []byte{}, "The data of the event")

	err := eventCmd.MarkPersistentFlagRequired("topics")
	cobra.CheckErr(err)
}
