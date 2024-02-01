# Generic ERC20 Token Bridge

An ERC20 token bridge contract built on top of Teleporter to support bridging arbitrary ERC20 tokens between any `subnet-evm` instance.

## Design
The generic ERC20 bridge is implemented using two primary contracts.
- `BridgeToken`
    - A simple extension of OpenZeppelin's `ERC20Burnable` contract.
    - Representation of tokens bridged from another chain.
    - Identifies the specific native token and bridge it represents.
    - Automatically deployed by `ERC20Bridge` when new tokens are added to be supported by a bridge instance.
- `ERC20Bridge`
    - Bridge contract that uses Teleporter to facilitate the bridging operations of adding new supported tokens and moving tokens between chains.
    - The bridge contract tracks the balances of each token sent to other bridge instances, and only allows those bridge instances to redeem up to the amount of tokens that has been previously sent to them.
    - Primary functions include:
        - `submitCreateBridgeToken`: Called on the origin chain to add support for an ERC20 on a different chain's bridge instance. Submits a Teleporter message that invokes `createBridgeToken` on the destination chain.
        - `createBridgeToken`: Called by cross-chain messages to add support for a new bridge token from another chain. Assuming the token is not already supported, deploys a new `BridgeToken` contract instance to represent the tokens on this chain.
        - `bridgeTokens`: Called to move supported tokens from one chain to another. This includes both "wrapping" of tokens native to the chain and also "unwrapping" of bridge tokens on the chain to other chains. In the case of unwrapping or moving of a bridge token, if the destination chain ID is the native chain for the given token, the original token is sent to the recipient on the destination. If the destination chain ID is another non-native chain for the underlying token, a Teleporter message is first sent to the native chain, which then subsequently updates its accounting of balances and sends a Teleporter message on the destination chain to re-wrap the given tokens. This multi-hop flow is illustrated below.
        - `mintBridgeTokens`: Called by cross-chain messages in the event of new bridge transfers of tokens from their native chains.
        - `transferBridgeTokens`: Called by cross-chain messages being sent back from other chains to unwrap tokens previously bridged to them. Optionally submits a secondary Teleporter message to re-wrap the tokens on to the destination chain, as described above.

## Multiple Hop Flow
<div align="center">
  <img src="../../../../../resources/ERC20BridgeMultiHopDiagram.png?raw=true">
</div>

## End-to-end test
An end-to-end test demonstrating the use of the generic ERC20 token bridge contracts can be found in `tests/flows/erc20_bridge_multihop.go`. This test implements the following flow and checks that each step is successful:
1. An example "native" ERC20 token is deployed on subnet A
1. The `ERC20Bridge` contract is deployed on subnets A, B, and C.
1. Teleporter messages are sent from subnet A to B and subnet A to C to add the example ERC20 deployed in step 1 to the bridge instances on the other chains.
1. Tokens are bridged from subnet A to B.
1. Tokens are bridged from subnet B to C using a multi-hop through subnet A.
1. Tokens are unwrapped from subnet C back to subnet A, redeeming the original underlying tokens.
