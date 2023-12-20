# Getting Started with an Example Teleporter Application

> Note: All example applications in the [examples](./examples/) directory are meant for education purposes only and are not audited. The example contracts are not intended for use in production environments.

This section walks through how to build an example cross-chain application on top of the Teleporter protocol, recreating the `ExampleCrossChainMessenger` [contract](./examples/ExampleMessenger/ExampleCrossChainMessenger.sol) that sends arbitrary string data from one chain to another. Note that this tutorial is meant for education purposes only. The resulting code is not intended for use in production environments.

## Step 1: Create Initial Contract

Create a new file called `MyExampleCrossChainMessenger.sol` in the directory that will hold the application.

At the top of the file define the Solidity version to work with, and import the necessary types and interfaces.

```solidity
pragma solidity 0.8.18;

import {ITeleporterMessenger, TeleporterMessageInput, TeleporterFeeInfo} from "@teleporter/ITeleporterMessenger.sol";
import {ITeleporterReceiver} from "@teleporter/ITeleporterReceiver.sol";
```

Next, define the initial empty contract.

```solidity
contract MyExampleCrossChainMessenger {}
```

## Step 2: Integrating Teleporter Messenger

Now that the initial empty `MyExampleCrossChainMessenger` is defined, it's time to integrate the `ITeleporterMessenger` that will provide the functionality to deliver cross chain messages.

Create a state variable of `ITeleporterMessenger` type called `teleporterMessenger`. Then create a constructor for our contract that takes in an address where the Teleporter Messenger would be deployed on this chain, and set our state variable with it.

```solidity
contract ExampleCrossChainMessenger {
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
    bytes32 originBlockchainID,
    address originSenderAddress,
    bytes calldata message
) external {}
```

Now it's time to implement the methods, starting with `sendMessage`. First, add the import for OpenZeppelin's `IERC20` contract to the top of your contract.

```solidity
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
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
    if (feeAmount > 0) {
        IERC20 feeToken = IERC20(feeTokenAddress);
        require(
            feeToken.transferFrom(msg.sender, address(this), feeAmount),
            "Failed to transfer fee amount"
        );
        require(
            feeToken.approve(address(teleporterMessenger), feeAmount),
            "Failed to approve fee amount"
        );
    }
}
```

Note: Relayer fees are an optional way to incentive relayers to deliver a Teleporter message to its destination. They are not strictly necessary, and may be omitted if a relayer is willing to relay messages with no fee, such as with a self-hosted relayer.

Next, add the call to the `TeleporterMessenger` contract with the message data to be executed when delivered to the destination address. In `sendMessage`, form a `TeleporterMessageInput` and call `sendCrossChainMessage` on the `TeleporterMessenger` instance to start the cross chain messaging process.

> `allowedRelayerAddresses` is empty in this example, meaning any relayer can try to deliver this cross chain message. Specific relayer addresses can be specified to ensure only those relayers can deliver the message.
> The `message` must be ABI encoded so that it can be properly decoded on the receiving end.

```solidity
return
    teleporterMessenger.sendCrossChainMessage(
        TeleporterMessageInput({
            destinationBlockchainID: destinationBlockchainID,
            destinationAddress: destinationAddress,
            feeInfo: TeleporterFeeInfo({
                feeTokenAddress: feeTokenAddress,
                amount: feeAmount
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
    bytes32 originBlockchainID,
    address originSenderAddress,
    bytes calldata message
) external {
    // Only the Teleporter receiver can deliver a message.
    require(msg.sender == address(teleporterMessenger), "Unauthorized.");

    // do something with message.
    return true;
}
```

The base of sending and receiving messages cross chain is complete. `MyExampleCrossChainMessenger` can now be expanded with functionality that saves the received messages, and allows users to query for the latest message received from a specified chain.

## Step 4: Storing the Message

Start by defining the `struct` for how to save our messages. It saves the string message itself and the address of the sender.

A map will also be added where the key is the `originBlockchainID`, and the value is the latest `message` sent from that chain.

```solidity
// Messages sent to this contract.
struct Message {
    address sender;
    string message;
}

mapping(bytes32 originBlockchainID => Message message) private _messages;
```

Next, update `receiveTeleporterMessage` to save the message into our mapping after we receive and verify that it's sent from Teleporter. ABI decode the `message` bytes into a string.

```solidity
// Receive a new message from another chain.
function receiveTeleporterMessage(
    bytes32 originBlockchainID,
    address originSenderAddress,
    bytes calldata message
) external {
    // Only the Teleporter receiver can deliver a message.
    require(msg.sender == address(teleporterMessenger), "Unauthorized.");

    // Store the message.
    messages[originBlockchainID] = Message(originSenderAddress, abi.decode(message, (string)));
}
```

Next, add a function called `getCurrentMessage` that allows users or contracts to easily query our contract for the latest message sent by a specified chain.

```solidity
// Check the current message from another chain.
function getCurrentMessage(
    bytes32 originBlockchainID
) external view returns (address, string memory) {
    Message memory messageInfo = messages[originBlockchainID];
    return (messageInfo.sender, messageInfo.message);
}
```

There we have it, a simple cross chain messenger built on top of Teleporter! Full example [here](./ExampleMessenger/ExampleCrossChainMessenger.sol).

## Step 5: Testing

For testing, `scripts/local/e2e_test.sh` sets up a local test environment consisting of three subnets deployed with Teleporter, and a lightweight inline relayer implementation to facilitate cross chain message delivery. An end-to-end test for `ExampleCrossChainMessenger` is included in `tests/example_messenger.go`, which performs the following:

1. Deploys the [ExampleERC20](@mocks/ExampleERC20.sol) token to subnet A.
2. Deploys `ExampleCrossChainMessenger` to both subnets A and B.
3. Approves the cross-chain messenger on subnet A to spend ERC20 tokens from the default address.
4. Sends `"Hello, world!"` from subnet A to subnet B's cross-chain messenger to receive.
5. Calls `getCurrentMessage` on subnet B to make sure the right message and sender are received.
