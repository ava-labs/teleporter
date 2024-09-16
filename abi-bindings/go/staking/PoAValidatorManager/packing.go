// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poavalidatormanager

import (
	"fmt"

	"github.com/ava-labs/subnet-evm/accounts/abi"
)

var subnetConversionDataType abi.Type

func init() {
	var err error
	subnetConversionDataType, err = abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "convertSubnetTxID", Type: "bytes32"},
		{Name: "blockchainID", Type: "bytes32"},
		{Name: "validatorManagerAddress", Type: "bytes"},
		{Name: "initialValidators", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "nodeID", Type: "bytes32"},
			{Name: "weight", Type: "uint64"},
			{Name: "blsPublicKey", Type: "bytes"},
		}},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create subnetConversionData ABI type: %v", err))
	}
}

func (s *SubnetConversionData) Pack() ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "subnetConversionData",
			Type: subnetConversionDataType,
		},
	}
	return args.Pack(s)
}

func (s *SubnetConversionData) Unpack(b []byte) error {
	args := abi.Arguments{
		{
			Name: "subnetConversionData",
			Type: subnetConversionDataType,
		},
	}
	unpacked, err := args.Unpack(b)
	if err != nil {
		return fmt.Errorf("failed to unpack to subnetConversionData with err: %v", err)
	}
	return args.Copy(&s, unpacked)
}
