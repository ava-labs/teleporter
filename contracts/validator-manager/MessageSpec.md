## Warp Message Format Reference

### `SubnetConversionMessage`

- Description: Confirms conversion to a Permissionless Subnet on the P-Chain

- Signed by: P-Chain

- Consumed by: Validator Manager contract

Specification:

```
SubnetConversionMessage
+--------------------+----------+----------+
|            codecID :   uint16 |  2 bytes |
+--------------------+----------+----------+
|             typeID :   uint32 |  4 bytes |
+--------------------+----------+----------+
| subnetConversionID : [32]byte | 32 bytes |
+--------------------+----------+----------+
                                | 38 bytes |
                                +----------+
```

where `subnetConversionID` is defined as the `sha256` hash of the serialization of the `SubnetConversionData`, defined as follows:

```
ValidatorData
+--------------+----------+------------------------+
|       nodeID :   []byte |  4 + len(nodeID) bytes |
+--------------+----------+------------------------+
| blsPublicKey : [48]byte |               48 bytes |
+--------------+----------+------------------------+
|       weight :   uint64 |                8 bytes |
+--------------+----------+------------------------+
                          | 60 + len(nodeID) bytes |
                          +------------------------+

SubnetConversionData
+----------------+-----------------+--------------------------------------------------------+
|       codecID  :          uint16 |                                                2 bytes |
+----------------+-----------------+--------------------------------------------------------+
|       subnetID :        [32]byte |                                               32 bytes |
+----------------+-----------------+--------------------------------------------------------+
| managerChainID :        [32]byte |                                               32 bytes |
+----------------+-----------------+--------------------------------------------------------+
| managerAddress :          []byte |                          4 + len(managerAddress) bytes |
+----------------+-----------------+--------------------------------------------------------+
|     validators : []ValidatorData |                        4 + sum(validatorLengths) bytes |
+----------------+-----------------+--------------------------------------------------------+
                                   | 74 + len(managerAddress) + len(validatorLengths) bytes |
                                   +--------------------------------------------------------+                          
```

### `RegisterSubnetValidatorMessage`

- Description: Registers a Subnet Validator on the P-Chain

- Signed by: Subnet

- Consumed by: P-Chain

Specification:

```
PChainOwner
+-----------+------------+-------------------------------+
| threshold :     uint32 |                       4 bytes |
+-----------+------------+-------------------------------+
| addresses : [][20]byte | 4 + len(addresses) * 20 bytes |
+-----------+------------+-------------------------------+
                         | 8 + len(addresses) * 20 bytes |
                         +-------------------------------+

RegisterSubnetValidatorMessage
+-----------------------+-------------+--------------------------------------------------------------------+
|               codecID :      uint16 |                                                            2 bytes |
+-----------------------+-------------+--------------------------------------------------------------------+
|                typeID :      uint32 |                                                            4 bytes |
+-----------------------+-------------+-------------------------------------------------------------------+
|              subnetID :    [32]byte |                                                           32 bytes |
+-----------------------+-------------+--------------------------------------------------------------------+
|                nodeID :      []byte |                                              4 + len(nodeID) bytes |
+-----------------------+-------------+--------------------------------------------------------------------+
|          blsPublicKey :    [48]byte |                                                           48 bytes |
+-----------------------+-------------+--------------------------------------------------------------------+
|                expiry :      uint64 |                                                            8 bytes |
+-----------------------+-------------+--------------------------------------------------------------------+
| remainingBalanceOwner : PChainOwner |                                      4 + len(addresses) * 20 bytes |
+-----------------------+-------------+--------------------------------------------------------------------+
|          disableOwner : PChainOwner |                                      4 + len(addresses) * 20 bytes |
+-----------------------+-------------+--------------------------------------------------------------------+
|                weight :      uint64 |                                                            8 bytes |
+-----------------------+-------------+--------------------------------------------------------------------+
                                      | 114 + len(nodeID) + (len(addresses1) + len(addresses2)) * 20 bytes |
                                      +--------------------------------------------------------------------+

```

### `SubnetValidatorRegistrationMessage`

- Description: Confirms a Subnet Validator's registration validity on the P-Chain

- Signed by: P-Chain

- Consumed by: Validator Manager contract

Specification:

```
SubnetValidatorRegistrationMessage
+--------------+----------+----------+
|      codecID :   uint16 |  2 bytes |
+--------------+----------+----------+
|       typeID :   uint32 |  4 bytes |
+--------------+----------+----------+
| validationID : [32]byte | 32 bytes |
+--------------+----------+----------+
|        valid :     bool |  1 byte  |
+--------------+----------+----------+
                          | 39 bytes |
                          +----------+
```

### `SubnetValidatorWeightMessage`

- Description: Used to set and acknowledge a Validator's stake weight on another chain

- To effect a weight change:

    - Signed by: Subnet

    - Consumed by: P-Chain

- To acknowledge the weight change:

    - Signed by P-Chain


    - Consumed by Subnet

Specification:

```
SubnetValidatorWeightMessage
+--------------+----------+----------+
|      codecID :   uint16 |  2 bytes |
+--------------+----------+----------+
|       typeID :   uint32 |  4 bytes |
+--------------+----------+----------+
| validationID : [32]byte | 32 bytes |
+--------------+----------+----------+
|        nonce :   uint64 |  8 bytes |
+--------------+----------+----------+
|       weight :   uint64 |  8 bytes |
+--------------+----------+----------+
                          | 54 bytes |
                          +----------+
```

### `ValidationUptimeMessage`

- Description: Provides a Validator's uptime for calculating staking rewards

- Signed by: Subnet

- Consumed by: Validator Manager contract

Specification:

```
ValidationUptimeMessage
+--------------+----------+----------+
|      codecID :   uint16 |  2 bytes |
+--------------+----------+----------+
|       typeID :   uint32 |  4 bytes |
+--------------+----------+----------+
| validationID : [32]byte | 32 bytes |
+--------------+----------+----------+
|       uptime :   uint64 |  8 bytes |
+--------------+----------+----------+
                          | 46 bytes |
                          +----------+
```