// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleportermessenger

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func createTestTeleporterMessage(messageID int64) TeleporterMessage {
	m := TeleporterMessage{
		MessageID:          big.NewInt(messageID),
		SenderAddress:      common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		DestinationAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		RequiredGasLimit:   big.NewInt(2),
		AllowedRelayerAddresses: []common.Address{
			common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		},
		Receipts: []TeleporterMessageReceipt{
			{
				ReceivedMessageID:    big.NewInt(1),
				RelayerRewardAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
			},
		},
		Message: []byte{1, 2, 3, 4},
	}
	return m
}

func TestPackUnpackTeleporterMessage(t *testing.T) {
	var (
		messageID int64 = 4
	)
	message := createTestTeleporterMessage(messageID)

	b, err := PackSendCrossChainMessageEvent(common.HexToHash("0x03"), message, TeleporterFeeInfo{
		FeeTokenAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		Amount:          big.NewInt(1),
	})
	if err != nil {
		t.Errorf("failed to pack teleporter message: %v", err)
		t.FailNow()
	}

	unpacked, err := UnpackTeleporterMessage(b)
	if err != nil {
		t.Errorf("failed to unpack teleporter message: %v", err)
		t.FailNow()
	}

	for i := 0; i < len(message.AllowedRelayerAddresses); i++ {
		require.Equal(t, unpacked.AllowedRelayerAddresses[i], message.AllowedRelayerAddresses[i])
	}

	for i := 0; i < len(message.Receipts); i++ {
		require.Equal(t, message.Receipts[i].ReceivedMessageID, unpacked.Receipts[i].ReceivedMessageID)
		require.Equal(t, message.Receipts[i].RelayerRewardAddress, unpacked.Receipts[i].RelayerRewardAddress)
	}

	require.True(t, bytes.Equal(message.Message, unpacked.Message))
}
