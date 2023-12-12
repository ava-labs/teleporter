# Teleporter Cross Chain Applications

This directory includes cross-chain applications that are built on top of the [Teleporter protocol](../Teleporter/README.md).

## Example Applications

- `ERC20Bridge` allows cross chain bridging of erc20 assets. More details found [here](./examples/ERC20Bridge/README.md)
- `ExampleMessenger` a simple cross chain messenger that demonstrates Teleporter application sending arbitrary string data.
- `VerifiedBlockHash` publishes the latest block hash of one chain to a destination chain that receives the hash and verifies the sender. Includes `BlockHashPublisher` and `BlockHashReceiver`. More details found [here](./examples/VerifiedBlockHash/README.md)
