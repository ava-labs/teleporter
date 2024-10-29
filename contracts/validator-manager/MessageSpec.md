## Warp Message Format Reference

### `SubnetToL1ConversionMessage`

Description: Confirms conversion to L1 on the P-Chain

Signed by: P-Chain
Consumed by: L1

Specification:

The `SubnetToL1ConversionMessage` is specified as an `AddressedCall` with `sourceChainID` set to the P-Chain ID, the `sourceAddress` set to an empty byte array, and a payload of:

|          Field |       Type |     Size |
| -------------: | ---------: | -------: |
|      `codecID` |   `uint16` |  2 bytes |
|       `typeID` |   `uint32` |  4 bytes |
| `conversionID` | `[32]byte` | 32 bytes |
|                |            | 38 bytes |

- `codecID` is the codec version used to serialize the payload, and is hardcoded to `0x0000`
- `typeID` is the payload type identifier and is `0x00000000` for this message
- `conversionID` is the SHA256 hash of the `ConversionData` from a given `ConvertSubnetToL1Tx`

The following serialization is defined as the `ValidatorData`:

|          Field |       Type |                     Size |
| -------------: | ---------: | -----------------------: |
|       `nodeID` |   `[]byte` |  4 + len(`nodeID`) bytes |
| `blsPublicKey` | `[48]byte` |                 48 bytes |
|       `weight` |   `uint64` |                  8 bytes |
|                |            | 60 + len(`nodeID`) bytes |

The following serialization is defined as the `ConversionData`:

|            Field |              Type |                                                       Size |
| ---------------: | ----------------: | ---------------------------------------------------------: |
|        `codecID` |          `uint16` |                                                    2 bytes |
|       `subnetID` |        `[32]byte` |                                                   32 bytes |
| `managerChainID` |        `[32]byte` |                                                   32 bytes |
| `managerAddress` |          `[]byte` |                            4 + len(`managerAddress`) bytes |
|     `validators` | `[]ValidatorData` |                          4 + sum(`validatorLengths`) bytes |
|                  |                   | 74 + len(`managerAddress`) + sum(`validatorLengths`) bytes |

- `codecID` is the codec version used to serialize the payload, and is hardcoded to `0x0000`
- `sum(validatorLengths)` is the sum of the lengths of `ValidatorData` serializations included in `validators`.
- `subnetID` identifies the Subnet that is being converted to an L1 (described further below).
- `managerChainID` and `managerAddress` identify the validator manager for the newly created L1. This is the (blockchain ID, address) tuple allowed to send Warp messages to modify the L1's validator set.
- `validators` are the initial continuous-fee-paying validators for the given L1.

### `RegisterL1ValidatorMessage`

Description: Register an L1 Validator on the P-Chain

Signed by: L1
Consumed by: P-Chain

Specification:

|                   Field |          Type |                                                                      Size |
| ----------------------: | ------------: | ------------------------------------------------------------------------: |
|               `codecID` |      `uint16` |                                                                   2 bytes |
|                `typeID` |      `uint32` |                                                                   4 bytes |
|              `subnetID` |    `[32]byte` |                                                                  32 bytes |
|                `nodeID` |      `[]byte` |                                                   4 + len(`nodeID`) bytes |
|          `blsPublicKey` |    `[48]byte` |                                                                  48 bytes |
|                `expiry` |      `uint64` |                                                                   8 bytes |
| `remainingBalanceOwner` | `PChainOwner` |                                          8 + len(`addresses`) \* 20 bytes |
|          `disableOwner` | `PChainOwner` |                                          8 + len(`addresses`) \* 20 bytes |
|                `weight` |      `uint64` |                                                                   8 bytes |
|                         |               | 122 + len(`nodeID`) + (len(`addresses1`) + len(`addresses2`)) \* 20 bytes |

- `codecID` is the codec version used to serialize the payload, and is hardcoded to `0x0000`
- `typeID` is the payload type identifier and is `0x00000001` for this payload
- `subnetID`, `nodeID`, `weight`, and `blsPublicKey` are for the validator being added
- `expiry` is the time at which this message becomes invalid. As of a P-Chain timestamp `>= expiry`, this Avalanche Warp Message can no longer be used to add the `nodeID` to the validator set of `subnetID`
- `remainingBalanceOwner` is the P-Chain owner where leftover $AVAX from the validator's Balance will be issued to when this validator it is removed from the validator set.
- `disableOwner` is the only P-Chain owner allowed to disable the validator using `DisableL1ValidatorTx`, specified below.

The following is the serialization of a `PChainOwner`:

|       Field |         Type |                             Size |
| ----------: | -----------: | -------------------------------: |
| `threshold` |     `uint32` |                          4 bytes |
| `addresses` | `[][20]byte` | 4 + len(`addresses`) \* 20 bytes |
|             |              | 8 + len(`addresses`) \* 20 bytes |

- `threshold` is the number of `addresses` that must provide a signature for the `PChainOwner` to authorize an action.
- Validation criteria:
  - If `threshold` is `0`, `addresses` must be empty
  - `threshold` <= len(`addresses`)
  - Entries of `addresses` must be unique and sorted in ascending order

### `L1ValidatorRegistrationMessage`

Description: Confirms an L1 validator's registration validity on the P-Chain

Signed by: P-Chain
Consumed by: Validator Manager contract

Specification:

|          Field |       Type |     Size |
| -------------: | ---------: | -------: |
|      `codecID` |   `uint16` |  2 bytes |
|       `typeID` |   `uint32` |  4 bytes |
| `validationID` | `[32]byte` | 32 bytes |
|   `registered` |     `bool` |   1 byte |
|                |            | 39 bytes |

- `codecID` is the codec version used to serialize the payload, and is hardcoded to `0x0000`
- `typeID` is the payload type identifier and is `0x00000002` for this message
- `validationID` identifies the validator for the message
- `registered` is a boolean representing the status of the `validationID`. If true, the `validationID` corresponds to a validator in the current validator set. If false, the `validationID` does not correspond to a validator in the current validator set, and never will in the future.

### `ValidationUptimeMessage`

Description: Provides a Validator's uptime for calculating staking rewards

Signed by: L1
Consumed by: Validator Manager Contract

Specification:

|          Field |       Type |     Size |
| -------------: | ---------: | -------: |
|      `codecID` |   `uint16` |  2 bytes |
|       `typeID` |   `uint32` |  4 bytes |
| `validationID` | `[32]byte` | 32 bytes |
|       `uptime` |   `uint64` |   8 byte |
|                |            | 46 bytes |

### `L1ValidatorWeightMessage`

Description: Used to set a Validator's stake weight on another chain

Signed by: P-Chain or L1
Consumed by: P-Chain or L1

Specification:

|          Field |       Type |     Size |
| -------------: | ---------: | -------: |
|      `codecID` |   `uint16` |  2 bytes |
|       `typeID` |   `uint32` |  4 bytes |
| `validationID` | `[32]byte` | 32 bytes |
|        `nonce` |   `uint64` |  8 bytes |
|       `weight` |   `uint64` |  8 bytes |
|                |            | 54 bytes |

- `codecID` is the codec version used to serialize the payload, and is hardcoded to `0x0000`
- `typeID` is the payload type identifier and is `0x00000003` for this message
- `validationID` identifies the validator for the message
- `nonce` is a strictly increasing number that denotes the latest validator weight update and provides replay protection for this transaction
- `weight` is the new `weight` of the validator
