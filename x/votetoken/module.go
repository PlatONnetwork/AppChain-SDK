package votetoken

import (
	"encoding/json"
	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/votetoken/contracts/erc20vote"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
	"sync"
)

const (
	ModuleName           = "votetoken"
	ModuleVersion uint64 = 0
)

type GenesisParams struct {
	Name    string         `json:"name"`
	Symbol  string         `json:"symbol"`
	Version string         `json:"version"`
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

func (g *Module) Init(ctx sdk.InitContext) error {
	return nil
}

func (g *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, raw json.RawMessage) error {
	var params GenesisParams
	if err := json.Unmarshal(raw, &params); nil != err {
		log.Error("Failed UnmarshalJSON params", "error", err)
		return err
	}
	vote, _ := erc20vote.NewERC20Vote(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(g, g), false)
	vote.Init(params.Name, params.Symbol, params.Version, params.Owner)
	return nil
}
