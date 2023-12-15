package deposit

import (
	"encoding/json"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/deposit/contracts"
	deposittypes "github.com/PlatONnetwork/AppChain-SDK/x/deposit/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	MODULE_NAME_DEPOSIT = "deposit"
)

type DepositModule struct {
	logger   log.Logger
	l1Module deposittypes.L1Moduler
}

func NewDepositModule(ctx *cli.Context, l1Module deposittypes.L1Moduler) *DepositModule {

	return &DepositModule{
		logger:   log.New("module", MODULE_NAME_DEPOSIT),
		l1Module: l1Module,
	}
}

func (d *DepositModule) Name() string {
	return MODULE_NAME_DEPOSIT
}

func (d *DepositModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	// init deposit handler  account nonce
	initAccountNonce(db, d.Address())
	return nil
}

func (d *DepositModule) Address() basecommon.Address {
	return constants.DepositHandlerAddress
}

func (d *DepositModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	depositHandler, _ := contracts.NewDepositHandler(evm, contract, readOnly)
	depositHandler.SetL1Module(d.l1Module)
	return depositHandler.Run(input)
}
