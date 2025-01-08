package statesender

import (
	"encoding/json"
	"math/big"

	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesender/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	ModuleName           = "l2StateSender"
	ModuleVersion uint64 = 0
)

var _ module.ContractModule = (*StateSenderModule)(nil)

type StateSenderModule struct {
	logger log.Logger
}

func NewModule(ctx *cli.Context) *StateSenderModule {
	return &StateSenderModule{
		logger: log.New("module", ModuleName),
	}
}

func (s *StateSenderModule) Name() string {
	return ModuleName
}

func (s *StateSenderModule) Version() uint64 {
	return ModuleVersion
}

func (s *StateSenderModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	raw, err := data.MarshalJSON()
	if err != nil {
		return err
	}
	var config module.ModuleGenesisConfig
	if err := json.Unmarshal(raw, &config); err != nil {
		return err
	}

	stateSender, _ := contracts.NewL2StateSender(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(s, s), false)
	stateSender.InitGenesis(config.CreateBlock)

	return nil
}

func (s *StateSenderModule) Address() basecommon.Address {
	return constants.StateSenderAddress
}

func (s *StateSenderModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	l2StateSender, _ := contracts.NewL2StateSender(evm, contract, readOnly)
	return l2StateSender.Run(input)
}

func (s *StateSenderModule) ContractCreateBlockNumber(statedb sdk.StateDBReader) uint64 {
	stateSender, _ := contracts.NewL2StateSender(sdkcontracts.NewEVM(types.NewStateDBWrapper(statedb), big.NewInt(0)), sdkcontracts.NewContract(s, s), false)
	return stateSender.GetCreateBlock()
}
