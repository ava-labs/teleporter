# Contract Structure

The ERC20 and native token bridges built on top of Teleporter are composed of interfaces and abstract contracts that make them extendable to new implementations in the future.

### `ITokenBridge`
Interface that defines the events bridge contract implementations must emit. Also defines the message types and formats of messages between all implementations.

### `IERC20TokenBridge` and `INativeTokenBridge`
Interfaces that define the external functions for interacting with bridge contract implementations of each type. ERC20 and native token bridge interfaces vary from each other in that the native token bridge functions are `payable` and do not take an explicit amount parameter (it is implied by `msg.value`), while the ERC20 token bridge functions are not `payable` and require the explicit amount parameter. Otherwise, they include the same functions.

### `TeleporterTokenSource`
An abstract implementation of `ITokenBridge` for a bridge contract on the home chain with the asset to be bridged. Handles locking tokens to be sent to destination chains, as well as receiving bridge messages to either redeem tokens it holds as collateral (i.e unlock), or route them to another chain (i.e. "multi-hop"). In the case of a multi-hop transfer, the `TeleporterTokenSource` already has the collateral locked from when the tokens were originally bridged to the first destination chain, so it simply updates the accounting of the transferred balances to each respective destination. Destination contracts must first be registered with a `TeleporterTokenSource` instance before the source contract will allow for sending tokens to them. This is to prevent tokens from being bridged to invalid destination addresses. Anyone is able to deploy and register destination contracts, which may have been modified from this repository. It is the responsibility of the users of the source contract to independently evaluate each destination for their security and correctness.

It is intended for there to be a single `TeleporterTokenSource` implementation and instance per asset issued on one blockchain to be bridge to other chains. The single source contract instance supports arbitrarily many destinations on other blockchains.

Note that the term "source" in reference to bridge contracts refers to the home chain of a given asset, where the collateral is locked to be bridged. It is unrelated to the source and destination chain of an individual Teleporter message. `TeleporterTokenSource` instances both send and receive Teleporter messages.

### `ERC20Source`
A concrete implementation of `TeleporterTokenSource` and `IERC20TokenBridge` that handles the locking and releasing of an ERC20 asset on the home chain.

### `NativeTokenSource`
A concrete implementation of `TeleporterTokenSource` and `INativeTokenBridge` that handles the locking and release of the native EVM asset on the home chain.

### `TeleporterTokenDestination`
An abstract implementation of `ITokenBridge` for a bridge contract on a destination chain that receives bridged assets from a `TeleporterTokenSource` instance. Handles receiving bridge messages from the specified token source contract to process token imports from the home chain (i.e mints), as well as burning tokens and sending messages to route them back to other chains. Once deployed, a `TeleporterTokenDestination` instance must be registered with its specified `TeleporterTokenSource` contract. This is done by calling `registerWithHub` on the destination contract, which will send a Teleporter message to the source contract with the message to register.

All messages sent by `TeleporterTokenDestination` instances are sent to the specified token source contract, whether they are to redeem the collateral on the home or chain or route the tokens to another destination chain. Routing tokens from one destination chain to another is referred to as a "multi-hop", where the tokens are first sent back to their token source contract to update its accounting, and then automatically routed on to their intended destination.

`TeleporterTokenDestination` contracts allow for scaling token amounts, if the representative "wrapped" token is not a 1-to-1 equivalent of the backing asset. This token scaling can be used when the destination chain asset has a higher or lower denomination than the source asset, such as allowing for a ERC20 source asset with a denomination of 6 to be used as the native EVM asset on a destination chain (with a denomination of 18).

Note that the term "destination" in reference to bridge contracts refers to a non-home chain of a given asset where a representative asset is minted, backed by the collateral locked on the home (source) chain. It is unrelated to the destination of an individual Teleporter message. `TeleporterTokenDestination` instances both send and receive Teleporter messages.

### `ERC20Destination`
A concrete implementation of `TeleporterTokenDestination`, `IERC20TokenBridge`, and `IERC20` that handles the minting and burning of a destination chain asset. Note that the `ERC20Destination` contract _is_ the "wrapped" ERC20 implementation itself, which is why it takes the `tokenName`, `tokenSymbol`, and `tokenDecimals` in its constructor. All of the ERC20 interface implementations are inherited from the standard OpenZeppelin ERC20 implementation, and can be overriden in other implementations if desired.

### `NativeTokenDestination`
A concrete implementation of `TeleporterTokenDestination`, `INativeTokenBridge`, and `IWrappedNativeToken` that handles the minting and burning the native EVM asset on the destination chain using the native minter precompile. Deployments of this contract must be given the permission mint native coins in the chains configuration. Note that the `NativeTokenDestination` is also an implementation of `IWrappedNativeToken` itself, which is why the `nativeAssetSymbol` must be provided in its constructor.

The [native minter precompile](https://docs.avax.network/build/subnet/upgrade/customize-a-subnet#minting-native-coins) must be configured to allow the contract address of the `NativeTokenDestination` instance to call `mintNativeCoin`. The correctness of a native token bridge implemented using `NativeTokenDestination` relies on no other accounts being allowed to call `mintNativeCoin`, which could result in the bridge becoming undercollateralized. Example initialization steps for a `NativeTokenDestination` instance are shown below.

Since the native minter precompile does not provide an interface for burning the native EVM asset, the "burn" functionality is implemented by transferring the native coins to an unowned address. The contract also provides a `reportBurnedTxFees` interface in order to burn the home chain collateral that should also be made un-redeemable to account for tokens burnt on the destination chain to pay for transaction fees.

To account for the need to bootstrap the chain using a bridged asset as its native token, the `NativeTokenDestination` takes the `intitialReserveImbalance` in its constructor, and does not mint that first amount of tokens that it receives. The following example demonstrates the intended initialization flow:

1. Create a new blockchain with 100 native tokens allocated in its genesis block, and the pre-derived `NativeTokenDestination` contract address (based on the deployer nonce) set with the permission to mint native tokens using the native minter precompile. Note that the deployer account will need to be funded in order to deploy the `NativeTokenDestination` contract, and an account used to relay messages into this chain must also be funded to relay the first messages.
2. Deploy the `NativeTokenDestination` contract to the pre-derived address set in the blockchain configuration of step 1. The `initialReserveImbalance` should be 100, matching the number of tokens allocated in the genesis block that were not initially backed by collateral on the home chain.
3. Bridge at least 100 source tokens from the home chain to the new blockchain. The first 100 tokens bridged, which are possibly moved in multiple independent transfers, will not be minted on the destination since they were pre-allocated in the genesis file. Instead, a `CollateralAdded` event will be emitted by the `NativeTokenDestination` contract.
4. Now that the `NativeTokenDestination` contract is fully collateralized, tokens can be moved normally in both directions across the bridge contracts. 

The `totalNativeAssetSupply` implementation of `NativeTokenDestination` takes into account:
- the initial reserve imbalance
- the number of native tokens that it has minted
- the number of native tokens that have been burned to pay for transaction fees
- the number of native tokens "burned" to be bridged back to the home chain, which are sent to a pre-defined `BURNED_FOR_BRIDGE_ADDRESS`.

# Teleporter Message Fees

Fees can be optionally added to Teleporter messages in order to incentivize relayers to deliver them, as documented [here](https://github.com/ava-labs/teleporter/tree/main/contracts/src/Teleporter#fees). The bridge contracts in this repository allow for specifying any ERC20 token and amount to be used as the Teleporter message fee for single-hop transfers between `TeleporterTokenSource` and `TeleporterTokenDestination` instances. Multi-hop transfers between two `TeleporterTokenDestination` instances involve two Teleporter messages: the first from the initiating chain to source chain, and the second from the source chain on to the final destination. In the multi-hop case, the first message fee can be paid in any ERC20 token and amount (similar to the single-hop case), but the second message fee must be paid in-kind of the asset being transferred and is deducted from the amount being bridged. This restriction on the secondary message fee is necessary because the transaction on the source chain routing the funds to the destination chain is not sent by the wallet performing the transfer. Because of this, it can not directly spend an arbitrary ERC20 token from that wallet. Using the asset being transferred for the optional secondary fee allows users to perform an incentivized multi-hop transfer without needing to make any interaction with the source chain themselves. If there is a need for the second message from the source chain to the final destination chain to pay a fee in another asset, it is recommended to perform two single-hop transfers, which allows for specifying an arbitrary ERC20 token to be used for the fee of each.
