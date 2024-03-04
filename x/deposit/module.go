package deposit

import (
	"encoding/json"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/deposit/contracts"
	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	deposittypes "github.com/PlatONnetwork/AppChain-SDK/x/deposit/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	ModuleName           = "deposit"
	ModuleVersion uint64 = 0
)

var _ module.ContractModule = (*DepositModule)(nil)

type DepositModule struct {
	logger   log.Logger
	l1Module deposittypes.L1Moduler
}

func NewModule(ctx *cli.Context, l1Module deposittypes.L1Moduler) *DepositModule {
	return &DepositModule{
		logger:   log.New("module", ModuleName),
		l1Module: l1Module,
	}
}

func (d *DepositModule) Name() string {
	return ModuleName
}

func (d *DepositModule) Version() uint64 {
	return ModuleVersion
}

func (d *DepositModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	// init deposit handler  account nonce
	initAccountNonce(db, d.Address())

	raw, err := data.MarshalJSON()
	if err != nil {
		return err
	}
	var config module.ModuleGenesisConfig
	if err := json.Unmarshal(raw, &config); err != nil {
		return err
	}
	depositContract, _ := contracts.NewDepositHandler(sdkcontracts.NewEVM(db), sdkcontracts.NewContract(d, d), false)
	depositContract.SetCreateBlock(config.CreateBlock)
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

func (d *DepositModule) ContractCreateBlockNumber(statedb sdk.StateDB) uint64 {
	depositContract, _ := contracts.NewDepositHandler(sdkcontracts.NewEVM(statedb), sdkcontracts.NewContract(d, d), false)
	return depositContract.GetCreateBlock()
}
