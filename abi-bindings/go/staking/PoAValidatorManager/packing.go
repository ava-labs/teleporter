// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poavalidatormanager

func (s *SubnetConversionData) Pack() ([]byte, error) {
	packedLen := 72 + len(s.ValidatorManagerAddress) + 88*len(s.InitialValidators)

	b := make([]byte, packedLen)
	copy(b[0:32], s.ConvertSubnetTxID[:])
	copy(b[32:64], s.BlockchainID[:])
	lenAddress := len(s.ValidatorManagerAddress)
	for i := 0; i < 4; i++ {
		b[i+64] = byte(lenAddress >> uint(8*(1-i)))
	}
	copy(b[68:68+lenAddress], s.ValidatorManagerAddress)
	lenValidators := len(s.InitialValidators)
	for i := 0; i < 4; i++ {
		b[i+68+lenAddress] = byte(lenValidators >> uint8(8*(3-i)))
	}
	for i, iv := range s.InitialValidators {
		packedIv, err := iv.Pack()
		if err != nil {
			return nil, err
		}
		copy(b[72+lenAddress+i*88:72+lenAddress+(i+1)*88], packedIv)
	}
	return b, nil
}

func (iv *InitialValidator) Pack() ([]byte, error) {
	b := make([]byte, 88)
	copy(b[0:32], iv.NodeID[:])
	for i := 0; i < 8; i++ {
		b[i+32] = byte(iv.Weight >> uint8(8*(7-i)))
	}
	copy(b[40:88], iv.BlsPublicKey)
	return b, nil
}
