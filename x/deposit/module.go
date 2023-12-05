package deposit

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/deposit/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
)

type DepositModule struct {
}

func NewDepositModule() *DepositModule {

	return &DepositModule{}
}

func (d *DepositModule) Name() string {
	return "deposit"
}

func (d *DepositModule) Address() basecommon.Address {
	return constants.StakeHandlerAddress
}

func (d *DepositModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	depositHandler, _ := contracts.NewDepositHandler(evm, contract, readOnly)
	return depositHandler.Run(input)
}
