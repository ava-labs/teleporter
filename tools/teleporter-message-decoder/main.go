package main

import (
	"context"
	"flag"
	"fmt"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/subnet-evm/ethclient"
	"github.com/ava-labs/subnet-evm/x/warp"
	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/go/Teleporter/TeleporterMessenger"
	"github.com/ethereum/go-ethereum/common"
)

const (
	warpPrecompileAddress = "0x0200000000000000000000000000000000000005"
)

var cfg settings

func main() {
	cfg = loadSetting()

	client, err := ethclient.Dial(cfg.Rpc)
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")
	recipet, err := client.TransactionReceipt(context.Background(),
		common.HexToHash(cfg.TxnHash))
	if err != nil {
		panic(err)
	}

	for _, log := range recipet.Logs {
		if log.Address == common.HexToAddress(warpPrecompileAddress) {
			msg, err := UnpackWarpMessage(log.Data)
			if err != nil {
				panic(err)
			}

			fmt.Println("neworkID:", msg.WarpUnsignedMessage.NetworkID)
			fmt.Println("souceChainID:", msg.WarpUnsignedMessage.SourceChainID)

			teleporterMessage, err := teleportermessenger.UnpackTeleporterMessage(msg.WarpPayload)
			if err != nil {
				panic(err)
			}
			fmt.Printf("messageID: %v\n", teleporterMessage.MessageID)
			fmt.Println("senderAddress:", teleporterMessage.SenderAddress)
			fmt.Println("destinationAddress:", teleporterMessage.DestinationAddress)
			fmt.Println("msg", teleporterMessage.Message)
		}
	}
}

type WarpMessageInfo struct {
	WarpUnsignedMessage *avalancheWarp.UnsignedMessage
	WarpPayload         []byte
}

func UnpackWarpMessage(unsignedMsgBytes []byte) (*WarpMessageInfo, error) {
	unsignedMsg, err := warp.UnpackSendWarpEventDataToMessage(unsignedMsgBytes)
	if err != nil {
		return nil, err
	}

	warpPayload, err := warpPayload.ParseAddressedCall(unsignedMsg.Payload)
	if err != nil {
		return nil, err
	}

	messageInfo := WarpMessageInfo{
		WarpUnsignedMessage: unsignedMsg,
		WarpPayload:         warpPayload.Payload,
	}
	return &messageInfo, nil
}

type settings struct {
	Rpc     string
	TxnHash string
}

func loadSetting() settings {
	url := flag.String("rpc", "", "rpc url")
	txnHash := flag.String("txnHash", "", "transaction hash")

	flag.Parse()

	cfg := settings{
		Rpc:     *url,
		TxnHash: *txnHash,
	}

	if cfg.Rpc == "" {
		panic("rpc url is required")
	}

	if cfg.TxnHash == "" {
		panic("transaction hash is required")
	}

	return cfg
}
