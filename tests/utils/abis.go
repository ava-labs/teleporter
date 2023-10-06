package utils

import (
	"math/big"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"

	teleportermessenger "github.com/ava-labs/teleporter/abi-bindings/Teleporter/TeleporterMessenger"
)

type teleporterABI struct {
	teleporterMessengerABI abi.ABI
}

func NewTeleporterABI() *teleporterABI {
	parsed, err := abi.JSON(strings.NewReader(teleportermessenger.TeleporterMessengerABI))
	if err != nil {
		panic(err)
	}

	return &teleporterABI{
		teleporterMessengerABI: parsed,
	}
}

// UnpackTeleporterMessage unpacks message bytes according to EVM ABI encoding rules into a TeleporterMessage
func (t *teleporterABI) UnpackTeleporterMessage(messageBytes []byte) (*teleportermessenger.TeleporterMessage, error) {
	var teleporterMessage teleportermessenger.TeleporterMessage
	err := t.teleporterMessengerABI.UnpackIntoInterface(&teleporterMessage, "teleporterMessage", messageBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack to teleporter message")
	}

	return &teleporterMessage, nil
}

func (t *teleporterABI) PackSendCrossChainMessage(input teleportermessenger.TeleporterMessageInput) ([]byte, error) {
	return t.teleporterMessengerABI.Pack("sendCrossChainMessage", input)
}

// PackReceiveCrossChainMessage packs a ReceiveCrossChainMessageInput to form a call to the receiveCrossChainMessage function
func (t *teleporterABI) PackReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) ([]byte, error) {
	return t.teleporterMessengerABI.Pack("receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// PackMessageReceived packs a MessageReceivedInput to form a call to the messageReceived function
func (t *teleporterABI) PackMessageReceived(originChainID ids.ID, messageID *big.Int) ([]byte, error) {
	return t.teleporterMessengerABI.Pack("messageReceived", originChainID, messageID)
}

// UnpackMessageReceivedResult attempts to unpack result bytes to a bool indicating whether the message was received
func (t *teleporterABI) UnpackMessageReceivedResult(result []byte) (bool, error) {
	var success bool
	err := t.teleporterMessengerABI.UnpackIntoInterface(&success, "messageReceived", result)
	return success, err
}

func (t *teleporterABI) PackMessageReceivedOutput(success bool) ([]byte, error) {
	return t.teleporterMessengerABI.PackOutput("messageReceived", success)
}

// CAUTION: PackEvent is documented as not supporting struct types, so this should only be used for testing puposes.
// In a real setting, the Teleporter contract should pack the event.
// PackSendCrossChainMessageEvent packs the SendCrossChainMessage event type.
func (t *teleporterABI) PackSendCrossChainMessageEvent(destinationChainID common.Hash, message teleportermessenger.TeleporterMessage) ([]byte, error) {
	_, hashes, err := t.teleporterMessengerABI.PackEvent("SendCrossChainMessage", destinationChainID, message.MessageID, message)
	return hashes, err
}
