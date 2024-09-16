// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poavalidatormanager

import "encoding/binary"

func (s *SubnetConversionData) Pack() ([]byte, error) {
	packedLen := 72 + len(s.ValidatorManagerAddress) + 88*len(s.InitialValidators)

	b := make([]byte, packedLen)
	copy(b[0:32], s.ConvertSubnetTxID[:])
	copy(b[32:64], s.BlockchainID[:])
	// These are evm addresses and have lengths of 20 so hardcoding here
	binary.BigEndian.PutUint32(b[64:68], uint32(20))
	copy(b[68:88], s.ValidatorManagerAddress.Bytes())
	binary.BigEndian.PutUint32(b[88:92], uint32(len(s.InitialValidators)))
	for i, iv := range s.InitialValidators {
		packedIv, err := iv.Pack()
		if err != nil {
			return nil, err
		}
		copy(b[92+i*88:92+(i+1)*88], packedIv)
	}
	return b, nil
}

func (iv *InitialValidator) Pack() ([]byte, error) {
	b := make([]byte, 88)
	copy(b[0:32], iv.NodeID[:])
	binary.BigEndian.PutUint64(b[32:40], uint64(iv.Weight))
	copy(b[40:88], iv.BlsPublicKey)
	return b, nil
}
