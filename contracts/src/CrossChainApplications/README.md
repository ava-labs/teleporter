# Teleporter Cross Chain Applications

This directory includes cross-chain applications that are built on top of the [Teleporter protocol](../Teleporter/README.md).

## Example Applications

> Note: All example applications in the [examples](./examples/) directory are meant for education purposes only and are not audited. The example contracts are not intended for use in production environments.

- `ERC20Bridge` allows cross-chain transfers of existing ERC20 assets. More details found [here](./examples/ERC20Bridge/README.md)
- `ExampleMessenger` a simple cross chain messenger that demonstrates Teleporter application sending arbitrary string data. More details found [here](./examples/ExampleMessenger/README.md)
- `VerifiedBlockHash` publishes the latest block hash of one chain to a destination chain that receives the hash and verifies the sender. Includes `BlockHashPublisher` and `BlockHashReceiver`. More details found [here](./examples/VerifiedBlockHash/README.md)
- `NativeTokenBridge` demonstrates transferring both a native EVM token and ERC20 assets to be the native EVM token on another chain. More details found [here](./examples/NativeTokenBridge/README.md)