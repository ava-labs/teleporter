package teleportermessenger

import (
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var teleporterMessageType abi.Type

func init() {
	// Create an ABI binding for TeleporterMessage, defined in ITeleporterMessenger.sol
	// abigen does not support ABI bindings for standalone structs, only methods and events,
	// so we must manually keep this up-to-date with the struct defined in the contract.
	var err error
	teleporterMessageType, err = abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "messageID", Type: "uint256"},
		{Name: "senderAddress", Type: "address"},
		{Name: "destinationChainID", Type: "bytes32"},
		{Name: "destinationAddress", Type: "address"},
		{Name: "requiredGasLimit", Type: "uint256"},
		{Name: "allowedRelayerAddresses", Type: "address[]"},
		{Name: "receipts", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "receivedMessageID", Type: "uint256"},
			{Name: "relayerRewardAddress", Type: "address"},
		}},
		{Name: "message", Type: "bytes"},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create TeleporterMessage ABI type: %v", err))
	}
}

func PackTeleporterMessage(message TeleporterMessage) ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "teleporterMessage",
			Type: teleporterMessageType,
		},
	}
	return args.Pack(message)
}

func UnpackTeleporterMessage(messageBytes []byte) (*TeleporterMessage, error) {
	args := abi.Arguments{
		{
			Name: "teleporterMessage",
			Type: teleporterMessageType,
		},
	}
	unpacked, err := args.Unpack(messageBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack to teleporter message with err: %v", err)
	}
	type teleporterMessageArg struct {
		TeleporterMessage TeleporterMessage `json:"teleporterMessage"`
	}
	var teleporterMessage teleporterMessageArg
	err = args.Copy(&teleporterMessage, unpacked)
	if err != nil {
		return nil, err
	}
	return &teleporterMessage.TeleporterMessage, nil
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
func PackSendCrossChainMessageEvent(destinationChainID common.Hash, message TeleporterMessage, feeInfo TeleporterFeeInfo) ([]byte, error) {
	abi, err := TeleporterMessengerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	_, hashes, err := abi.PackEvent("SendCrossChainMessage", destinationChainID, message.MessageID, message, feeInfo)
	return hashes, err
}
