## Utility Contracts

### ValidatorSetSig Contract

This contract provides an alternative to MultiSig contracts for upgradeability where instead of requiring k of n signatures on an upgrade transaction we can require a quorum of validators from a given blockchain.

The contract accepts off-chain [Avalanche Warp Messages](https://docs.avax.network/build/cross-chain/awm/overview) signed via a manual approve by a chain's validators. It requires these messages to have source address set to the zero address to enforce this. 

Note:
1. The blockchain validating the message may or may not be the same chain where the target contract and the ValidatorSetSig contract are deployed.
2. The quorum of signers required for a Warp Message to be considered valid is configured by the validators themselves via their [Warp configs](https://github.com/ava-labs/subnet-evm/issues/729)

A public `validateMessage` view method is provided as a way to verify that the message has correct fields before being passed on to validators for signing.