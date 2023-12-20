# Example Cross-chain Messenger

An example cross-chain messenger built on top of Teleporter that sends and receives messages between chains.

## Design

- `ExampleCrossChainMessenger`:
  - `sendMessage` utilizes the Teleporter protocol to start the cross-chain messaging process
  - `_receiveTeleporterMessage` is the handler for receiving Teleporter messages, and saves the most recently sent message from each source chain.
  - `getCurrentMessage` returns the most recent message from a source chain.

## End-to-end test
An end-to-end test demonstrating the use of this contract can be found in `tests/flows/example_messenger.go`. This test deploys the contracts on two different chains, and sends a cross-chain message between the two messenger instances.