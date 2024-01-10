// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterregistry

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestPackUnpackProtocolRegistryEntry(t *testing.T) {
	entry := ProtocolRegistryEntry{
		Version:         big.NewInt(1),
		ProtocolAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
	}
	destinationAddress := common.HexToAddress("0x0123456789abcdef0123456789abcdef01234568")

	b, err := PackTeleporterRegistryEntry(entry, destinationAddress)
	require.NoError(t, err)

	unpackedEntry, err := UnpackTeleporterRegistryEntry(b)
	require.NoError(t, err)

	require.Equal(t, entry.Version, unpackedEntry.Entry.Version)
	require.Equal(t, entry.ProtocolAddress, unpackedEntry.Entry.ProtocolAddress)
	require.Equal(t, destinationAddress, unpackedEntry.DestinationAddress)
}
