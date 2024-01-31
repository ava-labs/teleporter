// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package flows

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/teleporter/tests/interfaces"
	teleporterutils "github.com/ava-labs/teleporter/utils/teleporter-utils"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

// Tests Teleporter message ID calculation
func CalculateMessageID(network interfaces.Network) {
	subnetInfo := network.GetPrimaryNetworkInfo()
	teleporterContractAddress := network.GetTeleporterContractAddress()

	sourceBlockchainID := common.HexToHash("0xabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcdefabcd")
	destinationBlockchainID := common.HexToHash("0x1234567812345678123456781234567812345678123456781234567812345678")
	nonce := big.NewInt(42)

	expectedMessageID, err := subnetInfo.TeleporterMessenger.CalculateMessageID(
		&bind.CallOpts{},
		sourceBlockchainID,
		destinationBlockchainID,
		nonce,
	)
	Expect(err).Should(BeNil())

	calculatedMessageID, err := teleporterutils.CalculateMessageID(
		teleporterContractAddress,
		ids.ID(sourceBlockchainID),
		ids.ID(destinationBlockchainID),
		nonce,
	)
	Expect(err).Should(BeNil())
	Expect(ids.ID(expectedMessageID)).Should(Equal(calculatedMessageID))
}
