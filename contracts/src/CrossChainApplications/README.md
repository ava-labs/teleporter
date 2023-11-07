# Teleporter Cross Chain Applications

This directory includes cross-chain applications that are built on top of the [Teleporter protocol](../Teleporter/README.md).

- [Teleporter Cross Chain Applications](#teleporter-cross-chain-applications)
  - [Applications](#applications)
  - [Get Started](#get-started)
    - [Step 1: Create Initial Contract](#step-1-create-initial-contract)
    - [Step 2: Integrating Teleporter Messenger](#step-2-integrating-teleporter-messenger)
    - [Step 3: Send and Receive](#step-3-send-and-receive)
    - [Step 4: Storing the Message](#step-4-storing-the-message)
    - [Step 5: Testing](#step-5-testing)


## Applications

- `ERC20Bridge` allows cross chain bridging of erc20 assets. More details found [here](./ERC20Bridge/README.md)
- `ExampleMessenger` a simple cross chain messenger that demonstrates Teleporter application sending arbitrary string data.
- `VerifiedBlockHash` publishes the latest block hash of one chain to a destination chain that receives the hash and verifies the sender. Includes `BlockHashPublisher` and `BlockHashReceiver`. More details found [here](./VerifiedBlockHash/README.md)

## Getting started with an example application

This section walks through how to build a cross-chain application on top of the Teleporter protocol, recreating the `ExampleCrossChainMessenger` contract that sends arbitrary string data from one chain to another.

### Step 1: Create Initial Contract

Create a new file called `MyExampleCrossChainMessenger.sol` in the directory that will hold the application.

At the top of the file define the Solidity version to work with, and import the `ITeleporterMessenger` and `ITeleporterReceiver` interfaces.

```solidity
pragma solidity 0.8.18;

import "../../Teleporter/ITeleporterMessenger.sol";
```

Next, define the initial empty contract.

```solidity
contract MyExampleCrossChainMessenger {}
```

### Step 2: Integrating Teleporter Messenger

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

### Step 3: Send and Receive

Now that `MyExampleCrossChainMessenger` has an instantiation of `ITeleporterMessenger`, the next step is to add in functionality of sending and receiving arbitrary string data between chains.

To start, create the function declarations for `sendMessage`, which will send string data cross-chain to the specified destination address' receiver. This function allows callers to specify the destination chain ID, destination address to send to, relayer fees, required gas limit for message execution on destination address, and the actual message data.

```solidity
// Send a new message to another chain.
function sendMessage(
    bytes32 destinationChainID,
    address destinationAddress,
    address feeContractAddress,
    uint256 feeAmount,
    uint256 requiredGasLimit,
    string calldata message
) external returns (uint256 messageID) {}
```

`MyExampleCrossChainMessenger` also needs to implement `ITeleporterReceiver` by adding the method `receiveTeleporterMessage` that receives the cross-chain messages from Teleporter.

```solidity
// Receive a new message from another chain.
function receiveTeleporterMessage(
    bytes32 originChainID,
    address originSenderAddress,
    bytes calldata message
) external {
```

Now it's time to implement the methods, starting with `sendMessage`. First, import OpenZeppelin's `IERC20` contract, then in `sendMessage` check whether `feeAmount` is greater than zero. If it is, transfer and approve the amount of IERC20 asset at `feeContractAddress` to the Teleporter Messenger saved as a state variable. Relayer fees are an optional way to incentive relayers to deliver a Teleporter message to its destination. They are not strictly necessary, and may be omitted if relaying is guaranteed, such as with a self-hosted relayer.

```solidity
// For non-zero fee amounts, transfer the fee into the control of this contract first, and then
// allow the Teleporter contract to spend it.
if (feeAmount > 0) {
    IERC20 feeAsset = IERC20(feeContractAddress);
    require(
        feeAsset.transferFrom(msg.sender, address(this), feeAmount),
        "Failed to transfer fee amount"
    );
    require(
        feeAsset.approve(address(teleporterMessenger), feeAmount),
        "Failed to approve fee amount"
    );
}
```

Next, add the call to the `TeleporterMessenger` contract with the message data to be executed when delivered to the destination address. In `sendMessage`, form a `TeleporterMessageInput` and call `sendCrossChainMessage` on the `TeleporterMessenger` instance to start the cross chain messaging process.

> `allowedRelayerAddresses` is empty in this example, meaning any relayer can try to deliver this cross chain message. Specific relayer addresses can be specified to ensure only those relayers can deliver the message.
> The `message` must be ABI encoded so that it can be properly decoded on the receiving end.

```solidity
return
    teleporterMessenger.sendCrossChainMessage(
        TeleporterMessageInput({
            destinationChainID: destinationChainID,
            destinationAddress: destinationAddress,
            feeInfo: TeleporterFeeInfo({
                contractAddress: feeContractAddress,
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
    bytes32 originChainID,
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

### Step 4: Storing the Message

Start by defining the `struct` for how to save our messages. It saves the string message itself and the address of the sender.

A map will also be added where the key is the `originChainID`, and the value is the latest `message` sent from that chain.

```solidity
// Messages sent to this contract.
struct Message {
    address sender;
    string message;
}

mapping(bytes32 => Message) messages;
```

Next, update `receiveTeleporterMessage` to save the message into our mapping after we receive and verify that it's sent from Teleporter. ABI decode the `message` bytes into a string.

```solidity
// Receive a new message from another chain.
function receiveTeleporterMessage(
    bytes32 originChainID,
    address originSenderAddress,
    bytes calldata message
) external {
    // Only the Teleporter receiver can deliver a message.
    require(msg.sender == address(teleporterMessenger), "Unauthorized.");

    // Store the message.
    messages[originChainID] = Message(originSenderAddress, abi.decode(message, (string)));
    return true;
}
```

Next, add a function called `getCurrentMessage` that allows users or contracts to easily query our contract for the latest message sent by a specified chain.

```solidity
// Check the current message from another chain.
function getCurrentMessage(
    bytes32 originChainID
) external view returns (address, string memory) {
    Message memory messageInfo = messages[originChainID];
    return (messageInfo.sender, messageInfo.message);
}
```

There we have it, a simple cross chain messenger built on top of Teleporter! Full example [here](./ExampleMessenger/ExampleCrossChainMessenger.sol).

### Step 5: Testing

For testing, `scripts/local/test.sh` sets up a local Avalanche network with three subnets deployed with Teleporter, and a relayer to deliver Teleporter messages. To add an integration test simply add a new test script under `integration-tests`. An integration test for `ExampleCrossChainMessenger` is already included (`scripts/local/integration_tests/example_messenger.sh`), which performs the following steps:

1. Deploys the [ExampleERC20](../Mocks/ExampleERC20.sol) token to subnet A.
2. Deploys `ExampleCrossChainMessenger` to both subnets A and B.
3. Approves the cross-chain messenger on subnet A to spend ERC20 tokens from the default address.
4. Sends `"hello world"` from subnet A to subnet B's cross-chain messenger to receive.
5. Calls `getCurrentMessage` on subnet B to make sure the right message and sender are received.

Running `./test.sh example_messenger` at the root of the repo should yield at the end of the test:

```bash
test_runner         | Raw result is: "0x5DB9A7629912EBF95876228C24A848de0bfB43A9
test_runner         | hello world!"
test_runner         | The latest message from chain ID GoTmTVw77eGauaL17e1xPrtWEa72SQnvf9G8ackU6KZVGVYpz was:
test_runner         |     Message Sender: 0x5DB9A7629912EBF95876228C24A848de0bfB43A9
test_runner         |     Message: "hello world!"
test_runner         | Successfully called getCurrentMessage.
test_runner         | Done running test example_messenger.
test_runner         |
test_runner exited with code 0
```
