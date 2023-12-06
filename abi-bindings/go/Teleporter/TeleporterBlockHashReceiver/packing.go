package teleporterblockhashreceiver

// Pack packs a ReceiveCrossChainMessageInput to form a call to the receiveCrossChainMessage function
func PackReceiveBlockHash(messageIndex uint32) ([]byte, error) {
	abi, err := TeleporterBlockHashReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return abi.Pack("receiveBlockHash", messageIndex)
}
