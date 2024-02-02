# Getting Started with an Example Teleporter Application

> Note: All example applications in the [examples](./examples/) directory are meant for education purposes only and are not audited. The example contracts are not intended for use in production environments.

This section walks through how to build an example cross-chain application on top of the Teleporter protocol, recreating the `ExampleCrossChainMessenger` [contract](./examples/ExampleMessenger/ExampleCrossChainMessenger.sol) that sends arbitrary string data from one chain to another. Note that this tutorial is meant for education purposes only. The resulting code is not intended for use in production environments.

## Step 1: Create Initial Contract

Create a new file called `MyExampleCrossChainMessenger.sol` in a new directory: 

```bash
mkdir teleporter/contracts/src/CrossChainApplications/MyExampleCrossChainMessenger/
touch teleporter/contracts/src/CrossChainApplications/MyExampleCrossChainMessenger/MyExampleCrossChainMessenger.sol
```

At the top of the file define the Solidity version to work with, and import the necessary types and interfaces.

```solidity
pragma solidity 0.8.18;

import {ITeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts@4.8.1/security/ReentrancyGuard.sol";
import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";
```

Next, define the initial empty contract. The contract inherits from `ReentrancyGuard` to prevent reentrancy attacks, and inherits from `ITeleporterReceiver` to allow the contract to receive messages from Teleporter.

```solidity
contract MyExampleCrossChainMessenger is
    ReentrancyGuard,
    ITeleporterReceiver
{}
```

Finally, add the following struct and event declarations to the contract, which will be integrated in later:

```solidity
// Messages sent to this contract.
struct Message {
    address sender;
    string message;
}

/**
 * @dev Emitted when a message is submited to be sent.
 */
event SendMessage(
    bytes32 indexed destinationBlockchainID,
    address indexed destinationAddress,
    address feeTokenAddress,
    uint256 feeAmount,
    uint256 requiredGasLimit,
    string message
);

/**
 * @dev Emitted when a new message is received from a given chain ID.
 */
event ReceiveMessage(
    bytes32 indexed sourceBlockchainID,
    address indexed originSenderAddress,
    string message
);

```

## Step 2: Integrating Teleporter Messenger

Now that the initial empty `MyExampleCrossChainMessenger` is defined, it's time to integrate the `ITeleporterMessenger` that will provide the functionality to deliver cross chain messages.

Create a state variable of `ITeleporterMessenger` type called `teleporterMessenger`. Then create a constructor that takes in an address where the Teleporter Messenger would be deployed on this chain, and set the corresponding state variable.

```solidity
contract ExampleCrossChainMessenger is
    ReentrancyGuard,
    ITeleporterReceiver
{
    ITeleporterMessenger public immutable teleporterMessenger;

    constructor(address teleporterMessengerAddress) {
        teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);
    }
}
```

## Step 3: Send and Receive

Now that `MyExampleCrossChainMessenger` has an instantiation of `ITeleporterMessenger`, the next step is to add in functionality of sending and receiving arbitrary string data between chains.

To start, create the function declarations for `sendMessage`, which will send string data cross-chain to the specified destination address' receiver. This function allows callers to specify the destination chain ID, destination address to send to, relayer fees, required gas limit for message execution at the destination address' `receiveTeleporterMessage` function, and the actual message data.

```solidity
// Send a new message to another chain.
function sendMessage(
    bytes32 destinationBlockchainID,
    address destinationAddress,
    address feeTokenAddress,
    uint256 feeAmount,
    uint256 requiredGasLimit,
    string calldata message
) external returns (uint256 messageID) {}
```

`MyExampleCrossChainMessenger` also needs to implement `ITeleporterReceiver` by adding the method `receiveTeleporterMessage` that receives the cross-chain messages from Teleporter.

```solidity
// Receive a new message from another chain.
function receiveTeleporterMessage(
    bytes32 sourceBlockchainID,
    address originSenderAddress,
    bytes calldata message
) external {}
```

Now it's time to implement the methods, starting with `sendMessage`. First, add the import for OpenZeppelin's `IERC20` contract to the top of the contract, as well as the import for the `SafeERC20` library.

```solidity
import {IERC20} from "@openzeppelin/contracts@4.8.1/token/ERC20/IERC20.sol";
import {SafeERC20TransferFrom, SafeERC20} from "@teleporter/SafeERC20TransferFrom.sol";
```

Next, add a `using` directive in the contract declaration to specify `SafeERC20` as the `IERC20` implementation to use:

```solidity
contract ExampleCrossChainMessenger is
    ReentrancyGuard,
    TeleporterOwnerUpgradeable
{
    using SafeERC20 for IERC20;
    ...
}
```

Then in `sendMessage` check whether `feeAmount` is greater than zero. If it is, transfer and approve the amount of IERC20 asset at `feeTokenAddress` to the Teleporter Messenger saved as a state variable.

```solidity
function sendMessage(
    bytes32 destinationBlockchainID,
    address destinationAddress,
    address feeTokenAddress,
    uint256 feeAmount,
    uint256 requiredGasLimit,
    string calldata message
) external returns (uint256 messageID) {
    // For non-zero fee amounts, first transfer the fee to this contract, and then
    // allow the Teleporter contract to spend it.
    uint256 adjustedFeeAmount;
    if (feeAmount > 0) {
        adjustedFeeAmount = SafeERC20TransferFrom.safeTransferFrom(
            IERC20(feeTokenAddress),
            feeAmount
        );
        IERC20(feeTokenAddress).safeIncreaseAllowance(
            address(teleporterMessenger),
            adjustedFeeAmount
        );
    }
}
```

Note: Relayer fees are an optional way to incentive relayers to deliver a Teleporter message to its destination. They are not strictly necessary, and may be omitted if a relayer is willing to relay messages with no fee, such as with a self-hosted relayer.

Next, add the event to emit, as well as the call to the `TeleporterMessenger` contract with the message data to be executed when delivered to the destination address. In `sendMessage`, form a `TeleporterMessageInput` and call `sendCrossChainMessage` on the `TeleporterMessenger` instance to start the cross chain messaging process.

> `allowedRelayerAddresses` is empty in this example, meaning any relayer can try to deliver this cross chain message. Specific relayer addresses can be specified to ensure only those relayers can deliver the message.
> The `message` must be ABI encoded so that it can be properly decoded on the receiving end.

```solidity
emit SendMessage({
    destinationBlockchainID: destinationBlockchainID,
    destinationAddress: destinationAddress,
    feeTokenAddress: feeTokenAddress,
    feeAmount: adjustedFeeAmount,
    requiredGasLimit: requiredGasLimit,
    message: message
});
return
    teleporterMessenger.sendCrossChainMessage(
        TeleporterMessageInput({
            destinationBlockchainID: destinationBlockchainID,
            destinationAddress: destinationAddress,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: feeTokenAddress,
                amount: adjustedFeeAmount
            }),
            requiredGasLimit: requiredGasLimit,
            allowedRelayerAddresses: new address[](0),
            message: abi.encode(message)
        })
    );
```

With the sending side complete, the next step is to implement `ITeleporterReceiver.receiveTeleporterMessage`. The receiver in this example will just receive the arbitrary string data, and check that the message is sent through Teleporter.

```solidity
// Receive a new message from another chain.
function receiveTeleporterMessage(
    bytes32 sourceBlockchainID,
    address originSenderAddress,
    bytes calldata message
) external {
    // Only the Teleporter receiver can deliver a message.
    require(msg.sender == address(teleporterMessenger), "Unauthorized.");

    // do something with message.
}
```

The base of sending and receiving messages cross chain is complete. `MyExampleCrossChainMessenger` can now be expanded with functionality that saves the received messages, and allows users to query for the latest message received from a specified chain.

## Step 4: Storing the Message

Start by adding a map where the key is the `sourceBlockchainID`, and the value is the latest `message` sent from that chain. The `message` is of type `Message`, which is already declared in the contract.

```solidity
mapping(bytes32 sourceBlockchainID => Message message) private _messages;
```

Next, update `receiveTeleporterMessage` to save the message into the mapping after it is received and verified that it's sent from Teleporter. ABI decode the `message` bytes into a string. Also, emit the `ReceiveMessage` event.

```solidity

```solidity
// Receive a new message from another chain.
function receiveTeleporterMessage(
    bytes32 sourceBlockchainID,
    address originSenderAddress,
    bytes calldata message
) external {
    // Only the Teleporter receiver can deliver a message.
    require(msg.sender == address(teleporterMessenger), "Unauthorized.");

    // Store the message.
    string memory messageString = abi.decode(message, (string));
    _messages[sourceBlockchainID] = Message(
        originSenderAddress,
        messageString
    );
    emit ReceiveMessage(
        sourceBlockchainID,
        originSenderAddress,
        messageString
    );
}
```

Next, add a function called `getCurrentMessage` that allows users or contracts to easily query the contract for the latest message sent by a specified chain.

```solidity
// Check the current message from another chain.
function getCurrentMessage(
    bytes32 sourceBlockchainID
) external view returns (address, string memory) {
    Message memory messageInfo = messages[sourceBlockchainID];
    return (messageInfo.sender, messageInfo.message);
}
```

## Step 5: Upgrade Support

At this point, the contract is now fully usable, and can be used to send arbitrary string data between chains. However, there are a few more modifications that need to be made to support upgrades to `TeleporterMessenger`. For a more in-depth explanation of how to support upgrades, see the Upgrades README [here](../Teleporter/Upgrades/README.md).

The first change to make is to inherit from `TeleporterOwnerUpgradeable` instead of `ITeleporterReceiver`. `TeleporterOwnerUpgradeable` integrates with `TeleporterRegistry` via `TeleporterUpgradeable` to easily utilize the latest `TeleporterMessenger` implementation. `TeleporterOwnerUpgradeable` also ensures that only the contract owner is able to upgrade the `TeleporterMessenger` implementation used by the contract.

To start, replace the import for `ITeleporterReceiver` with `TeleporterOwnerUpgradeable`:

```diff
- import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";
+ import {TeleporterOwnerUpgradeable} from "@teleporter/upgrades/TeleporterOwnerUpgradeable.sol";
```

Also, replace the contract declaration to inherit from `TeleporterOwnerUpgradeable` instead of `ITeleporterReceiver`:

```diff
contract ExampleCrossChainMessenger is
     ReentrancyGuard,
-    ITeleporterReceiver
+    TeleporterOwnerUpgradeable
{}
```

Next, update the constructor to invoke the `TeleporterOwnerUpgradeable` constructor.

```diff
- constructor(address teleporterMessengerAddress) {
-     teleporterMessenger = ITeleporterMessenger(teleporterMessengerAddress);
- }
+ constructor(
+     address teleporterRegistryAddress
+ ) TeleporterOwnerUpgradeable(teleporterRegistryAddress) {}
```

Then, remove the `teleporterMessenger` state variable, and add a call to get the latest `ITeleporterMessenger` implementation from `TeleporterRegistry` in `sendMessage`.

```diff
- ITeleporterMessenger public immutable teleporterMessenger;
```

And finally, change `receiveTeleporterMessage` to `_receiveTeleporterMessage`, and mark it as `internal override`. It's also safe to remove the check against `teleporterMessenger` in `_receiveTeleporterMessage`, since that same check is handled in `TeleporterOwnerUpgradeable`'s `receiveTeleporterMessage` function.

```diff
- function receiveTeleporterMessage(
+ function _receiveTeleporterMessage(
    bytes32 sourceBlockchainID,
    address originSenderAddress,
    bytes memory message
- external {
+ internal override {
-    // Only the Teleporter receiver can deliver a message.
-    require(msg.sender == address(teleporterMessenger), "Unauthorized.");
```
    

`MyExampleCrossChainMessenger` is now a working cross-chain dApp built on top of Teleporter! Full example [here](./ExampleMessenger/ExampleCrossChainMessenger.sol).

## Step 6: Testing

For testing, `scripts/local/e2e_test.sh` sets up a local test environment consisting of three subnets deployed with Teleporter, and a lightweight inline relayer implementation to facilitate cross chain message delivery. An end-to-end test for `ExampleCrossChainMessenger` is included in `tests/flows/example_messenger.go`, which performs the following:

1. Deploys the [ExampleERC20](@mocks/ExampleERC20.sol) token to subnet A.
2. Deploys `ExampleCrossChainMessenger` to both subnets A and B.
3. Approves the cross-chain messenger on subnet A to spend ERC20 tokens from the default address.
4. Sends `"Hello, world!"` from subnet A to subnet B's cross-chain messenger to receive.
5. Calls `getCurrentMessage` on subnet B to make sure the right message and sender are received.

To run this test against the newly created `MyExampleCrossChainMessenger`, first generate the ABI Go bindings by running `./scripts/abi_bindings.sh` from the root of this repository. Then, modify `example_messenger.go` to use the ABI bindings for `MyExampleCrossChainMessenger` instead of `ExampleCrossChainMessenger`.
