package deposit

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/deposit/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"gopkg.in/urfave/cli.v1"
)

type DepositModule struct {
	logger log.Logger
}

func NewDepositModule(ctx *cli.Context) *DepositModule {

	return &DepositModule{
		logger: log.New("module", "deposit"),
	}
}

func (d *DepositModule) Name() string {
	return "deposit"
}

func (d *DepositModule) Address() basecommon.Address {
	return constants.DepositHandlerAddress
}

func (d *DepositModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	depositHandler, _ := contracts.NewDepositHandler(evm, contract, readOnly)
	return depositHandler.Run(input)
}
