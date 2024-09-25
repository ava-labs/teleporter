## Delegation

`PoSValidatorManager` supports Delegation to an active Validator as a way for users to earn staking rewards without having to validate the chain. Delegators pay a configurable percentage fee on any earned staking rewards to the host Validator. Delegations are reflected on the P-Chain by adjusting the Validator's registered weight via a [`SetSubnetValidatorWeightTx`](https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/77-reinventing-subnets#setsubnetvalidatorweighttx). The weight change acknowledgement is delivered to the `PoSValidatorManager` via a [`SubnetValidatorWeightUpdateMessage`](./README.md#subnetvalidatorweightupdatemessage). 

> [!NOTE]
> The P-Chain is only willing to sign a `SubnetValidatorWeightUpdateMessage` for an active Validator. Once Validator exit has been initiated (via a call to `initializeEndValidation`), the `PoSValidatorManager` must assume that the Validator has been deactivated on the P-Chain, and will therefore not sign any further weight updates. Therefore, it is invalid to _initiate_ adding or removing a Delegator when the Validator is in this state, though it _may be_ valid to _complete_ an already initiated Delegator action, depending on the order of delivery to the P-Chain. If the Delegator weight change was submitted (and a Warp signature on the acknowledgement retrieved) before the Validator was removed, then the Delegator action may be completed. Otherwise, the acknowledgement of the Validation end must first be delivered before completing the Delegator action. 

The following state transition diagram illustrates the relationship between Validator and Delegator state. `Validator` is abbreviated as `V`, `Delegator` is abbreviated as `D`, and function names are shortened to improve readability. `Delegator.Completed` is omitted to show equivalance between the Validator's initial state and the state after a Delegation completes.
```mermaid
stateDiagram-v2
    % Happy path
    [*] --> V.PendingAdded : initAddVdr
    V.PendingAdded --> V.Active : completeVdrReg
    V.Active --> V.Active,D.PendingAdded : initAddDel
    V.Active,D.PendingAdded --> V.Active,D.Active : completeAddDel
    V.Active,D.Active --> V.Active,D.PendingRemoved : initEndDel
    V.Active,D.PendingRemoved --> V.Active : completeEndDel
    V.Active --> V.PendingRemoved : initEndVdr
    V.PendingRemoved --> V.Completed : completeEndVdr

    % Validator/Delegator state changes do not affect the Delegator/Validator state
    V.Active,D.PendingRemoved --> V.PendingRemoved,D.PendingRemoved : initEndVdr
    V.Active,D.PendingAdded --> V.PendingRemoved,D.PendingAdded : initEndVdr
    V.Active,D.Active --> V.PendingRemoved,D.Active: initEndVdr

    % When the Validator is in PendingRemoved or Completed, in general Delegator actions
    % may be completed, but not initialized.
    V.PendingRemoved,D.PendingAdded --> V.Completed,D.PendingAdded : completeEndVdr
    V.PendingRemoved,D.PendingRemoved --> V.Completed,D.PendingRemoved : completeEndVdr
    V.PendingRemoved,D.PendingRemoved --> V.PendingRemoved : completeEndDel
    V.PendingRemoved,D.Active --> V.Completed,D.Active : completeEndVdr
    % This is a no-op
    V.PendingRemoved,D.PendingAdded --> V.PendingRemoved,D.Active : completeAddDel
    V.Completed,D.PendingRemoved --> V.Completed : endDelegationCompletedValidator
    V.Completed,D.Active --> V.Completed : endDelegationCompletedValidator
    V.Completed,D.PendingAdded --> V.Completed : endDelegationCompletedValidator
```