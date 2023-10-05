package teleporterblockhashreceiver

import "github.com/ava-labs/avalanchego/ids"

// ReceiveBlockHashInput is the input to receiveBlockHash call
// in the contract deployed on the destination chain
// - messageIndex: specifies the warp message in the transaction's storage slots.
// -
type ReceiveBlockHashInput struct {
	MessageIndex  uint32 `json:"messageIndex"`
	SourceChainID ids.ID `json:"sourceChainID"`
}

// Pack packs a ReceiveCrossChainMessageInput to form a call to the receiveCrossChainMessage function
func PackReceiveBlockHash(inputStruct ReceiveBlockHashInput) ([]byte, error) {
	abi, err := TeleporterblockhashreceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return abi.Pack("receiveBlockHash", inputStruct.MessageIndex, inputStruct.SourceChainID)
}
