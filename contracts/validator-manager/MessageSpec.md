## Warp Message Format Reference

### `RegisterSubnetValidatorMessage`

Description: Register a Subnet Validator on the P-Chain

Signed by: Subnet
Consumed by: P-Chain

Specification:

```
+--------------+----------+-----------+
|      codecID :   uint16 |   2 bytes |
+--------------+----------+-----------+
|       typeID :   uint32 |   4 bytes |
+--------------+----------+-----------+
|     subnetID : [32]byte |  32 bytes |
+--------------+----------+-----------+
|       nodeID : [32]byte |  32 bytes |
+--------------+----------+-----------+
|       weight :   uint64 |   8 bytes |
+--------------+----------+-----------+
| blsPublicKey : [48]byte |  48 bytes |
+--------------+----------+-----------+
|       expiry :   uint64 |   8 bytes |
+--------------+----------+-----------+
                          | 134 bytes |
                          +-----------+

```

### `SubnetValidatorRegistrationMessage`

Description: Confirms a Subnet Validator's registration validity on the P-Chain

Signed by: P-Chain
Consumed by: Validator Manager contract

Specification:

```
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

### `ValidationUptimeMessage`

Description: Provides a Validator's uptime for calculating staking rewards

Signed by: Subnet
Consumed by: Validator Manager contract

Specification:

```
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

### `SetSubnetValidatorWeightMessage`

Description: Used to set a Validator's stake weight on another chain

Signed by: Subnet
Consumed by: P-Chain

Specification:

```
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

### `SubnetValidatorWeightUpdateMessage`

Description: Acknowledges a Validator weight update

Signed by: P-Chain
Consumed by: Validator Manager contract

Specification:

```
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