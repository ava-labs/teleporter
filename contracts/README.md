# Contract Structure

The ERC20 and native token bridges built on top of Teleporter are composed of interfaces and abstract contracts that make them extendable to new implementations in the future.

### `ITeleporterTokenBridge`
Interface that defines the events bridge contract implementations must emit. Also defines the message types and formats of messages between all implementations.

### `IERC20Bridge` and `INativeTokenBridge`
Interfaces that define the external functions for interacting with bridge contract implementations of each type. ERC20 and native token bridge interfaces vary from each other in that the native token bridge functions are `payable` and do not take an explicit amount parameter (it is implied by `msg.value`), while the ERC20 token bridge functions are not `payable` and require the explicit amount parameter. Otherwise, they include the same functions.

### `TeleporterTokenSource`
An abstract implementation of `ITeleporterTokenBridge` for a bridge contract on the home chain with the asset to be bridged. Handles locking tokens to be sent to destination chains, as well as receiving bridge messages to either redeem tokens it holds as collateral (i.e unlock), or route them to another chain (i.e. "multi-hop"). In the case of a multi-hop transfer, the `TeleporterTokenSource` already has the collateral locked from when the tokens were originally bridged to the first destination chain, so it simply updates the accounting of the transferred balances to each respective destination.

### `ERC20Source`
A concrete implementation of `TeleporterTokenSource` and `IERC20Bridge` that handles the locking and releasing of an ERC20 asset on the home chain. The optional Teleporter message fees used to incentivize a relayer for messages sent by this contract are all paid in the source ERC20 asset that it facilitates the bridging of.

### `NativeTokenSource`
A concrete implementation of `TeleporterTokenSource` and `INativeTokenBridge` that handles the locking and release of the native EVM asset on the home chain. The optional Teleporter message fees used to incentivize a relayer for messages sent by this contract are all paid in a "wrapped native token" asset (e.g. WAVAX for the case in which AVAX is the native token). The contract address of the wrapped native token contract to be used must be provided in its constructor.

### `TeleporterTokenDestination`
An abstract implementation of `ITeleporterTokenBridge` for a bridge contract on a destination chain that receives bridged assets from a `TeleporterTokenSource` instance. Handles receiving bridge messages from the specified token source contract to process token imports from the home chain (i.e mints), as well as burning tokens and sending messages to route them back to other chains. 

All messages sent by `TeleporterTokenDestination` instances are sent to the specified token source contract, whether they are to redeem the collateral on the home or chain or route the tokens to another destination chain. Routing tokens from one destination chain to another is referred to as a "multi-hop", where the tokens are first sent back to their token source contract to update its accounting, and then automatically routed on to their intended destination.

`TeleporterTokenDestination` contracts allow for scaling token amounts, if the representative "wrapped" token is not a 1-to-1 equivalent of the backing asset. This token scaling can be used when the destination chain asset has a higher or lower denomincation than the source asset, such as allowing for a ERC20 source asset with a denomination of 6 to be used as the native EVM asset on a destination chain (with a denomination of 18).

### `ERC20Destination`
A concrete implementation of `TeleporterTokenDestination`, `IERC20Bridge`, and `IERC20` that handles the minting and burning of a destination chain asset. Note that the `ERC20Destination` contract _is_ the "wrapped" ERC20 implementation itself, which is why it takes the `tokenName`, `tokenSymbol`, and `tokenDecimals` in its constructor. All of the ERC20 interface implementations are inherited from the standard OpenZeppelin ERC20 implementation, and can be overriden in other implementations if desired. The optional Teleporter message fees used to incentivize a relayer for messages sent by this contract are all paid in kind, since the contract is an ERC20 token itself.

### `NativeTokenDestination`
A concrete implementation of `TeleporterTokenDestination`, `INativeTokenBridge`, and `IWrappedNativeToken` that handles the minting and burning the native EVM on the destination chain using the native minter precompile. Deployments of this contract must be given the permission mint native coins in the chains configuration. Note that the `NativeTokenDestination` is also an implementation of `IWrappedNativeToken` itself, which is why the `symbol` must be provided in its constructor. The optional Teleporter message fees used to incentivize a relayer for messages sent by this contract are all paid using the wrapped native token that the contract implements.

Since the native minter precompile does not provide an interface for burning the native EVM asset, the "burn" functionality is implemented by locking the native asset in the contract itself in an un-redeemable manner. The contract also provides a `reportBurnedTxFees` interface in order to burn the home chain collateral that should also be made un-redeemable to account for tokens burnt on the destination chain to pay for transaction fees.

To account for the need to bootstrap the chain using a bridged asset as its native token, the `NativeTokenDestination` takes the `intitialReserveImbalance` in its constructor, and does not mint that first amount of tokens that it recevies. The following example demonstrates the intended initialization flow:

1. Create a new blockchain with 100 native tokens allocated in its genesis block, and the pre-derived `NativeTokenDestination` contract address (based on the deployer nonce) set with the permission to mint native tokens using the native minter precompile. Note that the deployer account will need to be funded in order to deploy the `NativeTokenDestination` contract, and an account used to relay messages into this chain must also be funded to relay the first messages.
2. Deploy the `NativeTokenDestination` contract to the pre-derived address set in the blockchain configuration of step 1. The `initialReserveImbalance` should be 100, matching the number of tokens allocated in the genesis block that were not initially backed by collateral on the home chain.
3. Bridge at least 100 source tokens from the home chain to the new blockchain. The first 100 tokens bridged, which are possibly moved in multiple independent transfers, will not be minted on the destination since they were pre-allocated in the genesis file. Instead, a `CollateralAdded` event will be emitted by the `NativeTokenDestination` contract.
4. Now that the `NativeTokenDestination` contract is fully collateralized, tokens can be moved normally in both directions across the bridge contracts. 

The `totalSupply` implementation of `NativeTokenDestination` takes into account:
- the initial reserve imbalance
- the number of native tokens that it has minted
- the number of native tokens that have been burned to pay for transaction fees
- the number of native tokens "burned" to be bridged back to the home chain, which are locked in the `NativeTokenDestination` contract in perpetuity 

