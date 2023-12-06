package withdraw

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/withdraw/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"gopkg.in/urfave/cli.v1"
)

type WithdrawModule struct {
	logger log.Logger
}

func NewWithdrawModule(ctx *cli.Context) *WithdrawModule {
	return &WithdrawModule{
		logger: log.New("module", "withdraw"),
	}
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
