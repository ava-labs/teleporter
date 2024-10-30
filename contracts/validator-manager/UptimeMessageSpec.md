# `ValidationUptimeMessage` Warp Message Format Reference

Description: Provides a validator's uptime for calculating staking rewards

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

This is defined within `Subnet-EVM` [here](https://github.com/ava-labs/subnet-evm/blob/323eb0c7dd7204521e662a3a355fe78a0e19c0be/warp/messages/validator_uptime.go#L14-L19).
