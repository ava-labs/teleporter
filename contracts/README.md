# Contract Structure

The ERC20 and native token bridges built on top of Teleporter are composed of interfaces and abstract contracts that make them extendable to new implementations in the future.

### `ITokenBridge`
Interface that defines the events bridge contract implementations must emit. Also defines the message types and formats of messages between all implementations.

### `IERC20TokenBridge` and `INativeTokenBridge`
Interfaces that define the external functions for interacting with bridge contract implementations of each type. ERC20 and native token bridge interfaces vary from each other in that the native token bridge functions are `payable` and do not take an explicit amount parameter (it is implied by `msg.value`), while the ERC20 token bridge functions are not `payable` and require the explicit amount parameter. Otherwise, they include the same functions.

### `TokenHub`
An abstract implementation of `ITokenBridge` for a bridge contract on the "hub" chain with the asset to be bridged. Each `TokenHub` instance supports bridging exactly one token type (ERC20 or native) on its chain to arbitrarily many "spoke" chains. It handles locking tokens to be sent to spoke chains, as well as receiving bridge messages to either redeem tokens it holds as collateral (i.e unlock), or route them to other spoke chains (i.e. "multi-hop"). In the case of a multi-hop transfer, the `TokenHub` already has the collateral locked from when the tokens were originally bridged to the first destination chain, so it simply updates the accounting of the transferred balances to each respective spoke. Spoke contracts must first be registered with a `TokenHub` instance before the hub contract will allow for sending tokens to them. This is to prevent tokens from being bridged to invalid spoke addresses. Anyone is able to deploy and register spoke contracts, which may have been modified from this repository. It is the responsibility of the users of the hub contract to independently evaluate each spoke for its security and correctness.

### `ERC20TokenHub`
A concrete implementation of `TokenHub` and `IERC20TokenBridge` that handles the locking and releasing of an ERC20 token on the hub chain.

### `NativeTokenHub`
A concrete implementation of `TokenHub` and `INativeTokenBridge` that handles the locking and release of the native EVM asset on the hub chain.

### `TokenSpoke`
An abstract implementation of `ITokenBridge` for a bridge contract on a spoke chain that receives bridged assets from a specific `TokenHub` instance. Each `TokenSpoke` instance has a single `TokenHub` instance that it receives bridge transfers from to mint tokens. It also handles burning tokens and sending messages to route tokens back to other chains (either its `TokenHub`, or other `TokenSpoke` instances). Once deployed, a `TokenSpoke` instance must be registered with its specified `TokenHub` contract. This is done by calling `registerWithHub` on the spoke contract, which will send a Teleporter message to the hub contract with the information to register.

All messages sent by `TokenSpoke` instances are sent to the specified `TokenHub` contract, whether they are to redeem the collateral on the hub chain or route the tokens to another spoke chain. Routing tokens from one spoke chain to another is referred to as a "multi-hop", where the tokens are first sent back to their `TokenHub` contract to update its accounting, and then automatically routed on to their intended destination spoke.

`TokenSpoke` contracts allow for scaling token amounts, if the representative "wrapped" token is not a 1-to-1 equivalent of the backing asset. This token scaling can be used when the spoke asset has a higher or lower denomination than the hub asset, such as allowing for a ERC20 hub asset with a denomination of 6 to be used as the native EVM asset on a spoke chain (with a denomination of 18).

### `ERC20TokenSpoke`
A concrete implementation of `TokenSpoke`, `IERC20TokenBridge`, and `IERC20` that handles the minting and burning of an ERC20 spoke chain asset. Note that the `ERC20TokenSpoke` contract _is_ the "wrapped" ERC20 implementation itself, which is why it takes the `tokenName`, `tokenSymbol`, and `tokenDecimals` in its constructor. All of the ERC20 interface implementations are inherited from the standard OpenZeppelin ERC20 implementation, and can be overriden in other implementations if desired.

### `NativeTokenSpoke`
A concrete implementation of `TokenSpoke`, `INativeTokenBridge`, and `IWrappedNativeToken` that handles the minting and burning of the native EVM asset on the spoke chain using the native minter precompile. Deployments of this contract must be given the permission mint native coins in the chain's configuration. Note that the `NativeTokenSpoke` is also an implementation of `IWrappedNativeToken` itself, which is why the `nativeAssetSymbol` must be provided in its constructor.

The [native minter precompile](https://docs.avax.network/build/subnet/upgrade/customize-a-subnet#minting-native-coins) must be configured to allow the contract address of the `NativeTokenSpoke` instance to call `mintNativeCoin`. The correctness of a native token bridge implemented using `NativeTokenSpoke` relies on no other accounts being allowed to call `mintNativeCoin`, which could result in the bridge becoming undercollateralized. Example initialization steps for a `NativeTokenSpoke` instance are shown below.

Since the native minter precompile does not provide an interface for burning the native EVM asset, the "burn" functionality is implemented by transferring the native coins to an unowned address. The contract also provides a `reportBurnedTxFees` interface in order to burn the hub chain collateral that should also be made unredeemable to account for tokens burnt on the spoke chain to pay for transaction fees.

To account for the need to bootstrap the chain using a bridged asset as its native token, the `NativeTokenSpoke` takes the `intitialReserveImbalance` in its constructor. Once registered with its `TokenHub`, the `TokenHub` will require the `initialReserveImbalance` to be accounted for before sending token amounts to be minted on the given spoke. The following example demonstrates the intended initialization flow:

1. Create a new blockchain with 100 native tokens allocated in its genesis block, and the pre-derived `NativeTokenSpoke` contract address (based on the deployer nonce) set with the permission to mint native tokens using the native minter precompile. Note that the deployer account will need to be funded in order to deploy the `NativeTokenSpoke` contract, and an account used to relay messages into this chain must also be funded to relay the first messages.
2. Deploy the `NativeTokenSpoke` contract to the pre-derived address set in the blockchain configuration of step 1. The `initialReserveImbalance` should be 100, matching the number of tokens allocated in the genesis block that were not initially backed by collateral on the hub chain.
3. Call the `registerWithHub` function on the `NativeTokenSpoke` instance to send a Teleporter message registering this spoke with its `TokenHub`. This message should be relayed and delivered to the `TokenHub` instance.
3. Once registered on the `TokenHub` contract, add 100 tokens as collateral for the new `NativeTokenSpoke` instance by calling the `addCollateral` function on the `TokenHub` contract. A `CollateralAdded` event will be emitted by the `TokenHub` contract with a `remaining` amount of 0 once the `NativeTokenSpoke` is fully collateralized.
4. Now that the `NativeTokenSpoke` contract is fully collateralized, tokens can be moved normally in both directions across the bridge contracts by calling their `send` functions.

The `totalNativeAssetSupply` implementation of `NativeTokenSpoke` takes into account:
- the initial reserve imbalance
- the number of native tokens that it has minted
- the number of native tokens that have been burned to pay for transaction fees
- the number of native tokens "burned" to be bridged to other chains, which are sent to a pre-defined `BURNED_FOR_BRIDGE_ADDRESS`.

Note that the value returned by `totalNativeAssetSupply` is an upper bound on the circulating supply of the native asset on the spoke chain since tokens could be burned in other ways that it does not account for.

# Teleporter Message Fees

Fees can be optionally added to Teleporter messages in order to incentivize relayers to deliver them, as documented [here](https://github.com/ava-labs/teleporter/tree/main/contracts/src/Teleporter#fees). The bridge contracts in this repository allow for specifying any ERC20 token and amount to be used as the Teleporter message fee for single-hop transfers in eiter direction between `TokenHub` and `TokenSpoke` instances. Multi-hop transfers between two `TokenSpoke` instances involve two Teleporter messages: the first from the initiating spoke to its hub, and the second from the hub on to the destination spoke. In the multi-hop case, the first message fee can be paid in any ERC20 token and amount (similar to the single-hop case), but the second message fee must be paid in-kind of the asset being transferred and is deducted from the amount being bridged. This restriction on the secondary message fee is necessary because the transaction on the hub chain routing the funds to the destination spoke is not sent by the wallet performing the transfer. Because of this, it can not directly spend an arbitrary ERC20 token from that wallet. Using the asset being transferred for the optional secondary fee allows users to perform an incentivized multi-hop transfer without needing to make any interaction with the hub themselves. If there is a need for the second message from the hub to the destination spoke to pay a fee in another asset, it is recommended to perform two single-hop transfers, which allows for specifying an arbitrary ERC20 token to be used for the fee of each.
