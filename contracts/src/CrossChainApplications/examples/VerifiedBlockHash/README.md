# Verified Block Hashes

A sample cross-chain application built on top of Teleporter that publishes the block hash of one EVM chain to another. These block hashes could then be used to prove additional claims about the state of the origin chain.

## Contracts
- `BlockHashPublisher`:
    - Must be deployed first on the origin EVM chain.
    - Provides `publishLatestBlockHash` function to publish the last block hash to any other EVM instance running in an Avalanche subnet.
    - Does not currently include fees to incentivize a relayer to deliver the cross-chain message. It is assumed the relayers will want to deliver the message even without a monetary incentive, but the contract could be easily extended to support providing message fees.
- `BlockHashReceiver`:
    - Must be deployed on the destination chain after `BlockHashPublisher` is deployed on the origin such that the `BlockHashPublisher` contract address can be provided in the contract constructor.
    - `receiveTeleporterMessage` is the method invoked by cross-chain messages sent by the publisher contract.
    - Anyone can view the latest block hashes received from other chains using `getLatestBlockInfo`.

## End-to-end test
An end-to-end test demonstrating the use of these contracts can be found in `tests/flows/block_hash_publish_receive.go`. This test deploys the contracts on two different chains in the required order, and then checks that it can publish the block hash from one chain to another.