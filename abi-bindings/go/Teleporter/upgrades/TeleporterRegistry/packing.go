// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterregistry

import (
	"fmt"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var teleporterRegistryEntryType abi.Type

type TeleporterRegistryEntry struct {
	Entry              ProtocolRegistryEntry `json:"protocolRegistryEntry"`
	DestinationAddress common.Address        `json:"destinationAddress"`
}

func init() {
	// Create an ABI binding for TeleporterRegistryEntry, defined in TeleporterRegistry.sol
	// abigen does not support ABI bindings for standalone structs, only methods and events,
	// so we must manually keep this up-to-date with the struct defined in the contract.
	var err error
	teleporterRegistryEntryType, err = abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "entry", Type: "tuple", Components: []abi.ArgumentMarshaling{
			{Name: "version", Type: "uint256"},
			{Name: "protocolAddress", Type: "address"},
		}},
		{Name: "destinationAddress", Type: "address"},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create TeleporterRegistryEntry ABI type: %v", err))
	}
}

func PackTeleporterRegistryEntry(entry ProtocolRegistryEntry, destinationAddress common.Address) ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "teleporterRegistryEntry",
			Type: teleporterRegistryEntryType,
		},
	}
	return args.Pack(TeleporterRegistryEntry{
		Entry:              entry,
		DestinationAddress: destinationAddress,
	})
}

func UnpackTeleporterRegistryEntry(entryBytes []byte) (*TeleporterRegistryEntry, error) {
	args := abi.Arguments{
		{
			Name: "teleporterRegistryEntry",
			Type: teleporterRegistryEntryType,
		},
	}
	unpacked, err := args.Unpack(entryBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack to Teleporter registry entry with err: %v", err)
	}
	type teleporterRegistryEntryArg struct {
		Entry TeleporterRegistryEntry `json:"teleporterRegistryEntry"`
	}
	var teleporterRegistryEntry teleporterRegistryEntryArg
	err = args.Copy(&teleporterRegistryEntry, unpacked)
	if err != nil {
		return nil, err
	}
	return &teleporterRegistryEntry.Entry, nil
}

// PackAddProtocolVersion packs input to form a call to the addProtocolVersion function
func PackAddProtocolVersion(messageIndex uint32) ([]byte, error) {
	abi, err := TeleporterRegistryMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("addProtocolVersion", messageIndex)
}
