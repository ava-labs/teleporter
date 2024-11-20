# TeleporterMessenger Contract Deployment

The `TeleporterMessenger` contract is designed to only send and receive Avalanche ICM messages to and from its own address on different chains. We ensure that the contract can be deployed to the same address on every EVM based chain by using [Nick's Method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c). Only allowing messages to be sent and received by the same address guarantees that all messages use the same TeleporterMessenger message format because only the same exact contract bytecode could have been deployed to the same address.

This directory contains scripts written in Golang to construct a raw transaction using Nick's method that deploys the TeleporterMessenger contract, and determine the keyless address that must be prefunded in order for the transaction to be sent.

## Running

There are two supporting subcommands: `constructKeylessTx` and `deriveContractAddress`.

`go run utils/contract-deployment/contractDeploymentTools.go constructKeylessTx <PATH_TO_CONTRACT_JSON_FILE>`
OR
`go run utils/contract-deployment/contractDeploymentTools.go deriveContractAddress <DEPLOYER_ADDRESS> <NONCE>`

For example:
`go run utils/contract-deployment/contractDeploymentTools.go constructKeylessTx out/TeleporterMessenger.sol/TeleporterMessenger.json`
OR
`go run utils/contract-deployment/contractDeploymentTools.go deriveContractAddress 0x38545c4b331D8BFb3bee94C62D77a6735b5eF8c0 1`

## Results

The resulting raw transaction, `TeleporterMessenger` contract address, and universal deployer address are written to standard output, as well as to `UniversalTeleporterDeployerTransaction.txt`, `UniversalTeleporterMessengerContractAddress.txt`, and `UniversalTeleporterDeployerAddress.txt` respectively.

## Deploy the contract

Now that the keyless transaction is constructed, fund the deployer address. For example, using `cast`:

```bash
teleporter_deployer_address=$(cat UniversalTeleporterDeployerAddress.txt)
cast send --private-key $my_private_key --value 10ether $teleporter_deployer_address --rpc-url $my_rpc_url
```

Then, deploy TeleporterMessenger by sending the keyless transaction:

```bash
teleporter_deploy_tx=$(cat UniversalTeleporterDeployerTransaction.txt)
cast publish --rpc-url $my_rpc_url $teleporter_deploy_tx
```

Once you've verified that TeleporterMessenger was deployed to the address in `UniversalTeleporterMessengerContractAddress.txt`, TeleporterMessenger and ICM is ready to use.
