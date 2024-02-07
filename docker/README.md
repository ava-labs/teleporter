# Docker Development Environment

The Teleporter Docker development environment consists of one container running a local network with Teleporter deployed to all subnets, and one container running an instance of [AWM Relayer](https://github.com/ava-labs/awm-relayer) to facilitate message passing between the subnets.

The main entry point for the development environment is [run.sh](../scripts/local/run.sh), as described in the top-level [README](../README.md#run-a-local-testnet-in-docker).

This README highlights the commands used in this setup to initialize Teleporter. They should be easily adapted to other environments.

## Set up Teleporter in a local network using run_setup.sh

This script uses [avalanche-cli](https://github.com/ava-labs/avalanche-cli) to deploy and manage the subnets.

After populating the [template genesis file](./genesisTemplate.json), the script deploys the subnets. For example, to deploy subnet A:

```bash
avalanche subnet create subneta --force --genesis ./subnetGenesis_A.json --config ./docker/defaultNodeConfig.json --evm --vm-version $SUBNET_EVM_VERSION --log-level info --skip-update-check
avalanche subnet configure subneta --node-config ./docker/defaultNodeConfig.json --chain-config ./docker/defaultChainConfig.json --skip-update-check
avalanche subnet deploy subneta --local --avalanchego-version $AVALANCHEGO_VERSION --config ./docker/defaultNodeConfig.json --log-level info --skip-update-check
```

The script then derives the Teleporter deployment information using the [contract utility](../utils/contract-deployment/contractDeploymentTools.go) included in this repository:
```bash
go run utils/contract-deployment/contractDeploymentTools.go constructKeylessTx contracts/out/TeleporterMessenger.sol/TeleporterMessenger.json
teleporter_deployer_address=$(cat UniversalTeleporterDeployerAddress.txt)
teleporter_deploy_tx=$(cat UniversalTeleporterDeployerTransaction.txt)
teleporter_contract_address=$(cat UniversalTeleporterMessengerContractAddress.txt)
```

The deployer address is funded:
```bash
cast send --private-key $user_private_key --value 10ether $teleporter_deployer_address --rpc-url $subnet_a_rpc_url
```

And the Teleporter contract deployed and initialized:
```bash
cast publish --rpc-url $subnet_a_rpc_url $teleporter_deploy_tx
cast send --private-key $user_private_key $teleporter_contract_address "initializeBlockchainID()(bytes32)" --rpc-url $subnet_a_rpc_url
```

Finally, the Teleporter registry is deployed. It is optional but recommended for cross-chain dApps to interact with Teleporter via the registry.
```bash
forge create --private-key $user_private_key \
        --rpc-url $subnet_a_rpc_url src/Teleporter/upgrades/TeleporterRegistry.sol:TeleporterRegistry --constructor-args "[(1,$teleporter_contract_address)]")
```

## Set up the AWM Relayer using run_relayer.sh

This script waits until `run_setup.sh` has finished initializing the network:
```bash
until cat $dir_prefix/NETWORK_READY
```

It then reads the environment variables set by `run_setup.sh`, which are used to fund the relayer address and to populate the relayer configuration file.
```bash
cast send --private-key $user_private_key --value 500ether $relayer_address --rpc-url $subnet_a_rpc_url
```
