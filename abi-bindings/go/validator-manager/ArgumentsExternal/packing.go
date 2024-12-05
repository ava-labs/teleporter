package argumentsexternal

import "github.com/ava-labs/subnet-evm/accounts/abi"

var (
	initializeValidatorRegistrationArgsType abi.Type
	initializeEndValidationArgsType         abi.Type
)

func init() {
	var err error
	initializeValidatorRegistrationArgsType, err = abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "delegationFeeBips", Type: "uint16"},
		{Name: "minStakeDuration", Type: "uint64"},
	})
	if err != nil {
		panic("failed to create InitializeValidatorRegistrationArgs ABI type")
	}

	initializeEndValidationArgsType, err = abi.NewType("tuple", "struct Overloader.F", []abi.ArgumentMarshaling{
		{Name: "force", Type: "bool"},
		{Name: "includeUptimeProof", Type: "bool"},
		{Name: "messageIndex", Type: "uint32"},
		{Name: "rewardRecipient", Type: "address"},
	})
	if err != nil {
		panic("failed to create InitializeEndValidationArgs ABI type")
	}
}

func (a *InitializeValidatorRegistrationArgs) Pack() ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "initializeValidatorRegistrationArgs",
			Type: initializeValidatorRegistrationArgsType,
		},
	}
	return args.Pack(a)
}

func (a *InitializeEndValidationArgs) Pack() ([]byte, error) {
	args := abi.Arguments{
		{
			Name: "initializeEndValidationArgs",
			Type: initializeEndValidationArgsType,
		},
	}
	return args.Pack(a)
}
