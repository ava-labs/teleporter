package teleportermessenger

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// UnpackTeleporterMessage unpacks message bytes according to EVM ABI encoding rules into a TeleporterMessage
func UnpackTeleporterMessage(messageBytes []byte) (*TeleporterMessage, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	var teleporterMessage TeleporterMessage
	err = abi.UnpackIntoInterface(&teleporterMessage, "teleporterMessage", messageBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack to teleporter message")
	}

	return &teleporterMessage, nil
}

func PackSendCrossChainMessage(input TeleporterMessageInput) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("sendCrossChainMessage", input)
}

// PackReceiveCrossChainMessage packs a ReceiveCrossChainMessageInput to form a call to the receiveCrossChainMessage function
func PackReceiveCrossChainMessage(messageIndex uint32, relayerRewardAddress common.Address) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("receiveCrossChainMessage", messageIndex, relayerRewardAddress)
}

// PackMessageReceived packs a MessageReceivedInput to form a call to the messageReceived function
func PackMessageReceived(originChainID ids.ID, messageID *big.Int) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}
	return abi.Pack("messageReceived", originChainID, messageID)
}

// UnpackMessageReceivedResult attempts to unpack result bytes to a bool indicating whether the message was received
func UnpackMessageReceivedResult(result []byte) (bool, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return false, errors.Wrap(err, "failed to get abi")
	}

	var success bool
	err = abi.UnpackIntoInterface(&success, "messageReceived", result)
	return success, err
}

func PackMessageReceivedOutput(success bool) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.PackOutput("messageReceived", success)
}

// CAUTION: PackEvent is documented as not supporting struct types, so this should only be used for testing purposes.
// In a real setting, the Teleporter contract should pack the event.
// PackSendCrossChainMessageEvent packs the SendCrossChainMessage event type.
func PackSendCrossChainMessageEvent(destinationChainID common.Hash, message TeleporterMessage) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	_, hashes, err := abi.PackEvent("SendCrossChainMessage", destinationChainID, message.MessageID, message)
	return hashes, err
}
