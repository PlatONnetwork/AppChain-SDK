package votetoken

import (
	"encoding/json"
	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/votetoken/contracts/erc20vote"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math"
	"math/big"
	"sync"
)

const (
	ModuleName           = "votetoken"
	ModuleVersion uint64 = 0
)

type GenesisParams struct {
	module.ModuleGenesisConfig
	Name    string         `json:"name"`
	Symbol  string         `json:"symbol"`
	Version string         `json:"tokenVersion"`
	Owner   common.Address `json:"owner"`
}
type Module struct {
	sync.Mutex
	logger log.Logger
}

func NewModule() (*Module, error) {
	return &Module{
		logger: log.New("module", "votetoken"),
	}, nil
}

func (g *Module) Name() string {
	return ModuleName
}

func (g *Module) Version() uint64 {
	return ModuleVersion
}

func (g *Module) Address() common.Address {
	return constants.VoteTokenAddress
}
func (g *Module) ContractCreateBlockNumber(statedb vm.StateDBReader) uint64 {
	vote, _ := erc20vote.NewERC20Vote(sdkcontracts.NewEVM(types.NewStateDBWrapper(statedb), big.NewInt(0)), sdkcontracts.NewContract(g, g), false)
	return vote.GetCreateBlock()
}
func (m *Module) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	vote, _ := erc20vote.NewERC20Vote(evm, contract, readOnly)
	return vote.Run(input)
}

func (g *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, raw json.RawMessage) error {
	db.SetNonce(constants.VoteTokenAddress, 1)
	var params GenesisParams
	if err := json.Unmarshal(raw, &params); nil != err {
		log.Error("Failed UnmarshalJSON params", "error", err)
		return err
	}
	evm := vm.NewEVM(vm.BlockContext{GasLimit: math.MaxUint64, BlockNumber: big.NewInt(0)}, vm.TxContext{}, db, chainConfig, vm.Config{}, nil)
	vote, _ := erc20vote.NewERC20Vote(evm, sdkcontracts.NewContract(g, g), false)
	vote.Init(params.Name, params.Symbol, params.Version, params.Owner)
	vote.SetCreateBlock(params.CreateBlock)
	return nil
}
