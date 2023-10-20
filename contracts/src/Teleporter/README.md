# Teleporter Protocol

- [Teleporter Protocol](#teleporter-protocol)
  - [Overview](#overview)
  - [Data Flow](#data-flow)
  - [Properties](#properties)
  - [Fees](#fees)
  - [Required Interface](#required-interface)
  - [Message Delivery and Execution](#message-delivery-and-execution)


## Overview
Teleporter is a messaging protocol built on top of [Avalanche Warp Messaging (AWM)](https://docs.avax.network/learn/avalanche/awm) that provides a develeoper-friendly interface for sending and receiving cross-chain messages from within the EVM.

The `ITeleporterMessenger` interface provides two primary methods:
- `sendCrossChainMessage`: called by contracts on the origin chain to initiate the sending of a message to a contract on another EVM instance.
- `receiveCrossChainMessage`: called by cross-chain relayers on the destination chain to deliver signed messages to the destination EVM instance.

The `ITeleporterReceiver` interface provides a single method. All contracts that wish to receive Teleporter messages on the destination chain must implement this interface:
- `receiveTeleporterMessage`: called by the Teleporter contract on the destination chain to deliver a message to the destination contract.

## Data Flow
<div align="center">
  <img src="../../../resources/TeleporterDataFlowDiagram.png?raw=true">
</div>

## Properties
Teleporter provides a handful of useful properties to cross-chain applications that Avalanche Warp Messages do not provide by default. These include:

1. Replay protection: Teleporter ensures that a cross-chain message is not delivered multiple times.
2. Retries: In certain edge cases when there is significant validator churn, it is possible for an Avalanche Warp Message to be dropped before a valid aggregate signature is created for it. Teleporter ensures that messages can still be delivered even in this event by allowing for retries of previously submitted messages.
3. Relay incentivization: Teleporter provides a mechanism for messages to optionally incentivize relayers to perform the necessary signature aggregation and pay the transaction fee to broadcast the signed message on the destination chain.
4. Allowed relayers: Teleporter allows users to specify a list of `allowedRelayerAddresses`, where only the specified addresses can relay and deliver the Teleporter message. Leaving this list empty allows all relayers to deliver.
5. Message execution: Teleporter enables cross-chain messages to have direct effect on their destination chain by using `evm.Call()` to invoke the `receiveTeleporterMessage` function of destination contracts that implement the `ITeleporterReceiver` interface.

## Fees
Fees can be paid on a per message basis by specifing the ERC20 asset and amount to be used to incentivize a relayer to deliver the message in the call to `sendCrossChainMessage`. The fee amount is transferred into the control of the Teleporter contract (i.e. locked) before the Warp message is sent. The Teleporter contract tracks the fee amount for each message ID it creates. When it subsquently receives a message back from the destination chain of the original message, the new message will have a list of receipts identifying the relayer that delivered the given message ID. At this point, the fee amount originally locked by Teleporter for the given message will be redeemable by the relayer identified in the receipt. If the initial fee amount was not sufficient to incentivize a relayer, it can be added to by using `AddFeeAmount`.

## Required Interface
Teleporter messages are delivered by calling the `receiveTeleporterMessage` function defined by the `ITeleporterReceiver` interface. Contracts must implement this interface in order to be able to receive messages. The first two paramaters of `receiveTeleporterMessage` identify the original sender of the given message on the origin chain and are set by the `TeleporterMessenger`. The third parameter to `receiveTeleporterMessage`, is the raw message payload. Applications using Teleporter are responsible for defining the exact format of this payload in a way that can be decoded on the receiving end. For example, applications may encode and action enum value along with the target method parameters on the sending side, then decode this data and route to the target method within `receiveTeleporterMessage`. See `ERC20Bridge.sol` for an example of this approach.

## Teleporter Contract Deployment
The `TeleporterMessenger` contract must be deployed to the same contract address on every chain. This is acheived using Nick's keyless transaction method as described [here](../../../utils/contract-deployment/README.md). As a result, Warp messages sent from the resulting contract address are ensured to have the same payload format as defined by the contract itself.

## Message Delivery and Execution
Teleporter is able to ensure that messages are considered delivered even if their execution fails (i.e. reverts) by using `evm.Call()` with a pre-defined gas limit to execute the message payload. This gas limit is specified by each message in the call to `sendCrossChainMessage`. Relayers must provide at least enough gas for the sub-call in addition to the standard gas used by a call to `receiveCrossChainMessage`. In the event that a message exeuction runs out of gas or reverts for any other reason, the hash of the message payload is stored by the receiving Teleporter contract instance. This allows for the message execution to be retried in the future, with possibly a higher gas limit. Importantly, a message is still considered delivered on its destination chain even if its execution fails. This allows the relayer of the message to redeem their reward for deliverying the message, because they have no control on whether or not its execution will succeed or not so long as they provide sufficient gas to meet the specified `requiredGasLimit`.
