# Native Token Bridge

Smart contracts built on top of Teleporter to support using an ERC20 token or the native token of any EVM-based subnet as the native token for another subnet.

## Design
The Native Token Bridge is implemented using two primary contracts. 

1. A `TokenSource` contract that lives on the source chain, which can be any chain on Avalanche. 
2. A `NativeTokenDestination` contract that lives on the destination chain, which should be a newly created subnet.

### `TokenSource`
Can either be a `NativeTokenSource` contract or an `ERC20TokenSource` contract. It lives on the source chain. Pairs with exactly one `NativeTokenDestination` contract on the destination chain. It locks and unlocks native tokens on the Source chain corresponding to mints and burns on the destination chain.
- `transferToDestination`: Transfers all tokens paid to this function call to `recipient` on the destination chain by locking them and instructing the destination chain to mint. Optionally takes the address of an ERC20 contract `feeContractAddress` as well as an amount `feeAmount` that will be used as the relayer-incentivization for the Teleporter cross-chain call. Also allows for the caller to specify `allowedRelayerAddresses`.
- `receiveTeleporterMessage`: Only accepts Teleporter messages from the `NativeTokenDestination` contract on the destination chain. Can receive two types of messages:
  - `Unlock`: Unlocks tokens on the source chain when instructed to by the `NativeTokenDestination` contract.
  - `Burn`: Burns tokens on the source chain when instructed to by a call to `reportTotalBurnedTxFees` on the `NativeTokenDestination` contract. This function will burn tokens equal to the increase from the highest previously reported total from the destination chain.

### `NativeTokenDestination`
A contract that lives on the destination chain. Pairs with exactly one `TokenSource` contract on the source chain. It mints and burns native tokens on the Destination chain corresponding to locks and unlocks on the source chain.
- `transferToSource`: Transfers all tokens paid to this function call to `recipient` on the source chain by burning the tokens and instructing the source chain to unlock. Optionally takes the address of an ERC20 contract `feeContractAddress` as well as an amount `feeAmount` that will be used as the relayer-incentivisation for the Teleporter cross-chain call. Also allows for the caller to specify `allowedRelayerAddresses`.
- `receiveTeleporterMessage`: Only accepts Teleporter messages from the `TokenSource` contract on the source chain. Mints tokens on the destination chain when instructed to by the `NativeTokenDestination` contract.
- `isCollateralized`: Returns true if `currentReserveImbalance == 0`, meaning that enough tokens have been sent from the source chain to offset the `initialReserveImbalance`. If true, all tokens sent to the destination chain will be minted, and burning/unlocking tokens will be enabled.
- `totalSupply`: Returns the best estimate of available native tokens on this chain. Equal to the `initialReserveImbalance` + `tokens minted` - `tokens burned`. The amount of tokens burned is calculated as the combined balance of the burn address where transaction fees are sent and the burn address where token are sent when being transferred back to the source chain through this contract.
- `reportTotalBurnedTxFees`: Sends a Teleporter message to the source chain containing the total number of tokens at the burn address where transaction fees are sent (`0x0100000000000000000000000000000000000000`).

### Collateralizing the bridge
On initialization, the bridge will be undercollateralized by exactly the number of tokens minted in the genesis block on the destination chain. These tokens could theoretically be sent through the bridge, with no corresponding tokens able to be unlocked on the source chain. In order to avoid this problem, the `NativeTokenDestination` contract is initialized with the value for `initialReserveImbalance`.

`initialReserveImbalance` should correspond to the number of tokens allocated in the genesis block of the destination chain. If `initialReserveImbalance` is not properly set, behavior of this contract is undefined. The `NativeTokenDestination` contract will not mint tokens until it at least `initialReserveImbalance` tokens have been locked on the source chain. 

It is left to the contract deployer to ensure that the bridge is properly collateralized. Burning/unlocking is disabled until the bridge is properly collateralized.

### Reporting tokens burned for transaction fees
When tokens are burned as a part of transaction fees within the subnet, they are no longer redeemable to be able to be transferred back to the source chain. However, the collateral backing those tokens is still held within the `TokenSource` contract. In order to prevent the amount of unredeemable collateral held in the `TokenSource` contract from continually growing over time, the cumulative amount burned as transactions fees within the subnet can be reported back the `TokenSource` contract in order for it to burn the corresponding collateral. This amount is tracked as the balance of the address `0x0100...00`, where transaction fees are credited to. In order to differentiate the amount of the tokens burned by transaction fee from the amount of tokens burned by the `NativeTokenDestination` contract when transferring them back to the source chain, `transferToSource` burns tokens by sending them to the address `0x0100...01`.

Periodically reporting the amount of burned transaction fees back to the source chain can help mitigate the risk of theoretical subnet validator exploits by reducing the amount of collateral held on the source chain that should not be redeemable under any circumstances. 

## Setup
- `TeleporterMessenger` must be deployed on both chains, and the address must be passed to the constructor of the `TokenSource` and `NativeTokenDestination` contracts.
- `NativeTokenDestination` is meant to be deployed on a new subnet, and should be the only method for minting tokens on that subnet. The address of `NativeTokenDestination` must be included as the only entry for `adminAddresses` under `contractNativeMinterConfig` in the genesis config for the newly created subnet. See `warp-genesis.json` for an example.
- Both the chosen `TokenSource` contract and `NativeTokenDestination` need to be deployed to addresses known beforehand. Each address must be passed to the constructor of the other contract. To do this, you will need a known EOA, and preferably use the first transaction from the EOA (nonce 0) to deploy the contract on each chain. It is advised to allocate tokens to the EOA in the genesis of the new subnet such that it can easily deploy the contract.
- Both contracts need to be initialized with `teleporterMessengerAddress`, which is the only address from which cross-chain messages can be delivered.
- `NativeTokenDestination` needs to be initialized with `initialReserveImbalance`, which should equal the number of tokens allocated in the genesis of the new subnet. If `initialReserveImbalance` is not properly set, behavior of these contracts in undefined.
- On the source chain, at least `initialReserveImbalance` tokens need to be transferred to the `TokenSource` contract using `transferToSource` in order to properly collateralize the bridge and allow regular functionality in both directions. The first `initialReserveImbalance` tokens will not be delivered to the recipient. Burning/unlocking is disabled until the contracts are fully collateralized.
