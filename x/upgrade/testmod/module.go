package testmod

import (
	"encoding/json"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	ModuleName           = "testmod"
	ModuleVersion uint64 = 0
)

var (
	_ module.Module             = (*Module)(nil)
	_ module.BeginBlockerModule = (*Module)(nil)
	_ module.GenesisModule      = (*Module)(nil)
	_ module.RegistryModule     = (*Module)(nil)
)

type Module struct{
	logger log.Logger
}

func NewModule() *Module {
	return &Module{
		logger: log.New("module", ModuleName, "version", ModuleVersion),
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}

func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	m.logger.Info("Init genesis")
	return nil
}

func (m *Module) BeginBlock(ctx sdk.WorkerContext) error {
	m.logger.Info("Begin block", "number", ctx.Header().Number)
	return nil
}

func (m *Module) RegistryUpgradeHandler(registrar module.UpgradeRegistrar) error {
	return registrar.RegisterUpgradeHandler(ModuleName, ModuleVersion, m.UpgradeHandle)
}

func (m *Module) UpgradeHandle(ctx sdk.WorkerContext) error {
	m.logger.Info("Invoked UpgradeHandle", "number", ctx.Header().Number)
	return m.InitGenesis(ctx, ctx.StateDB(), nil, nil)
}
