package withdraw

import (
	"encoding/json"
	"math/big"

	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/withdraw/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	ModuleName           = "withdraw"
	ModuleVersion uint64 = 0
)

var _ module.ContractModule = (*WithdrawModule)(nil)

type WithdrawModule struct {
	logger log.Logger
}

func NewModule(ctx *cli.Context) *WithdrawModule {
	return &WithdrawModule{
		logger: log.New("module", ModuleName),
	}
}

func (w *WithdrawModule) Name() string {
	return ModuleName
}

func (w *WithdrawModule) Version() uint64 {
	return ModuleVersion
}

func (w *WithdrawModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	raw, err := data.MarshalJSON()
	if err != nil {
		return err
	}
	var config module.ModuleGenesisConfig
	if err := json.Unmarshal(raw, &config); err != nil {
		return err
	}

	withdraw, _ := contracts.NewWithdrawManager(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(w, w), false)
	withdraw.InitGenesis(config.CreateBlock)

	return nil
}

func (w *WithdrawModule) Address() basecommon.Address {
	return constants.WithdrawManagerAddress
}

func (w *WithdrawModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	withdrawManaher, _ := contracts.NewWithdrawManager(evm, contract, readOnly)
	return withdrawManaher.Run(input)
}

func (w *WithdrawModule) ContractCreateBlockNumber(statedb sdk.StateDBReader) uint64 {
	withdraw, _ := contracts.NewWithdrawManager(sdkcontracts.NewEVM(types.NewStateDBWrapper(statedb), big.NewInt(0)), sdkcontracts.NewContract(w, w), false)
	return withdraw.GetCreateBlock()
}
