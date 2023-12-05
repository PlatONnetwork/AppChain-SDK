package withdraw

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/withdraw/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
)

type WithdrawModule struct {
}

func NewWithdrawModule() *WithdrawModule {
	return &WithdrawModule{}
}

func (w *WithdrawModule) Name() string {
	return "withdraw"
}

func (w *WithdrawModule) Address() basecommon.Address {
	return constants.WithdrawManagerAddress
}

func (w *WithdrawModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	withdrawManaher, _ := contracts.NewWithdrawManager(evm, contract, readOnly)
	return withdrawManaher.Run(input)
}
