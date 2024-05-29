package iteleporterreceiver

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

// PackReceiveTeleporterMessage packs input to form a call to the receiveTeleporterMessage function
func PackReceiveTeleporterMessage(
	sourceBlockchainID [32]byte,
	originSenderAddress common.Address,
	message []byte) ([]byte, error) {
	abi, err := ITeleporterReceiverMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get abi")
	}

	return abi.Pack("receiveTeleporterMessage", sourceBlockchainID, originSenderAddress, message)
}
