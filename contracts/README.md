# Contract Structure

The ERC20 and native token bridges built on top of Teleporter are composed of interfaces and abstract contracts that make them extendable to new implementations in the future.

### `ITokenBridge`

Interface that defines the events bridge contract implementations must emit. Also defines the message types and formats of messages between all implementations.

### `IERC20TokenBridge` and `INativeTokenBridge`

Interfaces that define the external functions for interacting with bridge contract implementations of each type. ERC20 and native token bridge interfaces vary from each other in that the native token bridge functions are `payable` and do not take an explicit amount parameter (it is implied by `msg.value`), while the ERC20 token bridge functions are not `payable` and require the explicit amount parameter. Otherwise, they include the same functions.

### `TokenHome`

An abstract implementation of `ITokenBridge` for a bridge contract on the "home" chain with the asset to be bridged. Each `TokenHome` instance supports bridging exactly one token type (ERC20 or native) on its chain to arbitrarily many "remote" instances on other chains. It handles locking tokens to be sent to `TokenRemote` instances, as well as receiving bridge messages to either redeem tokens it holds as collateral (i.e. unlock), or route them to other `TokenRemote` instances (i.e. "multi-hop"). In the case of a multi-hop transfer, the `TokenHome` already has the collateral locked from when the tokens were originally bridged to the first `TokenRemote` instance, so it simply updates the accounting of the transferred balances to each respective `TokenRemote` instance. Remote contracts must first be registered with a `TokenHome` instance before the home contract will allow for sending tokens to them. This is to prevent tokens from being bridged to invalid remote addresses. Anyone is able to deploy and register remote contracts, which may have been modified from this repository. It is the responsibility of the users of the home contract to independently evaluate each remote for its security and correctness.

### `ERC20TokenHome`

A concrete implementation of `TokenHome` and `IERC20TokenBridge` that handles the locking and releasing of an ERC20 token.

### `NativeTokenHome`

A concrete implementation of `TokenHome` and `INativeTokenBridge` that handles the locking and release of the native EVM asset.

### `TokenRemote`

An abstract implementation of `ITokenBridge` for a bridge contract on a "remote" chain that receives bridged assets from a specific `TokenHome` instance. Each `TokenRemote` instance has a single `TokenHome` instance that it receives bridge transfers from to mint tokens. It also handles sending messages (and correspondingly burning tokens) to route tokens back to other chains (either its `TokenHome`, or other `TokenRemote` instances). Once deployed, a `TokenRemote` instance must be registered with its specified `TokenHome` contract. This is done by calling `registerWithHome` on the remote contract, which will send a Teleporter message to the home contract with the information to register.

All messages sent by `TokenRemote` instances are sent to the specified `TokenHome` contract, whether they are to redeem the collateral from the `TokenHome` instance or route the tokens to another `TokenRemote` instance. Routing tokens from one `TokenRemote` instance to another is referred to as a "multi-hop", where the tokens are first sent back to their `TokenHome` contract to update its accounting, and then automatically routed on to their intended destination `TokenRemote` instance.

TokenRemote contracts allow for scaling token amounts, which should be used when the remote asset has a higher or lower denomination than the home asset, such as allowing for a ERC20 home asset with a denomination of 6 to be used as the native EVM asset on a remote chain (with a denomination of 18).

### `ERC20TokenRemote`

A concrete implementation of `TokenRemote`, `IERC20TokenBridge`, and `IERC20` that handles the minting and burning of an ERC20 asset. Note that the `ERC20TokenRemote` contract is an ERC20 implementation itself, which is why it takes the `tokenName`, `tokenSymbol`, and `tokenDecimals` in its constructor. All of the ERC20 interface implementations are inherited from the standard OpenZeppelin ERC20 implementation, and can be overriden in other implementations if desired.

### `NativeTokenRemote`

A concrete implementation of `TokenRemote`, `INativeTokenBridge`, and `IWrappedNativeToken` that handles the minting and burning of the native EVM asset on its chain using the native minter precompile. Deployments of this contract must be given the permission to mint native coins in the chain's configuration. Note that the `NativeTokenRemote` is also an implementation of `IWrappedNativeToken` itself, which is why the `nativeAssetSymbol` must be provided in its constructor. `NativeTokenRemote` instances always have a denomination of 18, which is the denomination of the native asset of EVM chains.

The [native minter precompile](https://docs.avax.network/build/subnet/upgrade/customize-a-subnet#minting-native-coins) must be configured to allow the contract address of the `NativeTokenRemote` instance to call `mintNativeCoin`. The correctness of a native token bridge implemented using `NativeTokenRemote` relies on no other accounts being allowed to call `mintNativeCoin`, which could result in the bridge becoming undercollateralized. Example initialization steps for a `NativeTokenRemote` instance are shown below.

Since the native minter precompile does not provide an interface for burning the native EVM asset, the "burn" functionality is implemented by transferring the native coins to an unowned address. The contract also provides a `reportBurnedTxFees` interface in order to burn the collateral in the `TokenHome` instance that should be made unredeemable to account for native tokens burnt on the chain with the `NativeTokenRemote` instance to pay for transaction fees.

To account for the need to bootstrap the chain using a bridged asset as its native token, the `NativeTokenRemote` takes the `initialReserveImbalance` in its constructor. Once registered with its `TokenHome`, the `TokenHome` will require the `initialReserveImbalance` to be accounted for before sending token amounts to be minted on the given remote chain. The following example demonstrates the intended initialization flow:

1. Create a new blockchain with 100 native tokens allocated in its genesis block, and set the pre-derived `NativeTokenRemote` contract address (based on the deployer nonce) to have the permission to mint native tokens using the native minter precompile. Note that the deployer account will need to be funded in order to deploy the `NativeTokenRemote` contract, and an account used to relay messages into this chain must also be funded to relay the first messages.
2. Deploy the `NativeTokenRemote` contract to the pre-derived address set in the blockchain configuration of step 1. The `initialReserveImbalance` should be 100, matching the number of tokens allocated in the genesis block that were not initially backed by collateral in the `TokenHome` instance.
3. Call the `registerWithHome` function on the `NativeTokenRemote` instance to send a Teleporter message registering this remote with its `TokenHome`. This message should be relayed and delivered to the `TokenHome` instance.
4. Once registered on the `TokenHome` contract, add 100 tokens as collateral for the new `NativeTokenRemote` instance by calling the `addCollateral` function on the `TokenHome` contract. A `CollateralAdded` event will be emitted by the `TokenHome` contract with a `remaining` amount of 0 once the `NativeTokenRemote` is fully collateralized.
5. Now that the `NativeTokenRemote` contract is fully collateralized, tokens can be moved normally in both directions across the bridge contracts by calling their `send` functions.

The `totalNativeAssetSupply` implementation of `NativeTokenRemote` takes into account:

- the initial reserve imbalance
- the number of native tokens that it has minted
- the number of native tokens that have been burned to pay for transaction fees
- the number of native tokens "burned" to be bridged to other chains, which are sent to a pre-defined `BURNED_FOR_BRIDGE_ADDRESS`.

Note that the value returned by `totalNativeAssetSupply` is an upper bound on the circulating supply of the native asset on the chain using the `NativeTokenRemote` instance since tokens could be burned in other ways that it does not account for.

# Teleporter Message Fees

Fees can be optionally added to Teleporter messages in order to incentivize relayers to deliver them, as documented [here](https://github.com/ava-labs/teleporter/tree/main/contracts/src/Teleporter#fees). The bridge contracts in this repository allow for specifying any ERC20 token and amount to be used as the Teleporter message fee for single-hop transfers in either direction between `TokenHome` and `TokenRemote` instances. Fee amounts must be pre-approved to be spent by the bridge contract before initiating a transfer.

Multi-hop transfers between two `TokenRemote` instances involve two Teleporter messages: the first from the initiating `TokenRemote` instance to its home, and the second from its home to the destination `TokenRemote` instance. In the multi-hop case, the first message fee can be paid in any ERC20 token and amount (similar to the single-hop case), but the second message fee must be paid in-kind of the asset being transferred and is deducted from the amount being bridged. This restriction on the secondary message fee is necessary because the transaction on the intermediate chain routing the funds to the destination `TokenRemote` instance is not sent by the wallet performing the transfer. Because of this, it can not directly spend an arbitrary ERC20 token from that wallet. Using the asset being transferred for the optional secondary fee allows users to perform an incentivized multi-hop transfer without needing to make any interaction with the home themselves. If there is a need for the second message from the home to the destination `TokenRemote` instance to pay a fee in another asset, it is recommended to perform two single-hop transfers, which allows for specifying an arbitrary ERC20 token to be used for the fee of each.
