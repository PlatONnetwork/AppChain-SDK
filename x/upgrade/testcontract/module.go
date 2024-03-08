package testcontract

import (
	"encoding/json"
	"math/big"

	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/testcontract/contracts"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	ModuleName           = "testcontract"
	ModuleVersion uint64 = 1
)

var (
	addr = common.BigToAddress(big.NewInt(11111111))

	_ module.Module         = (*Module)(nil)
	_ module.GenesisModule  = (*Module)(nil)
	_ module.ContractModule = (*Module)(nil)
	_ module.RegistryModule = (*Module)(nil)
)

type Module struct {
	logger log.Logger
}

func NewModule() *Module {
	return &Module{
		logger: log.New("module", ModuleName),
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
	db.SetNonce(addr, 1)
	return nil
}

func (m *Module) Address() common.Address {
	return addr
}

func (m *Module) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	c, _ := contracts.NewCounter(evm, contract, readOnly)
	return c.Run(input)
}

func (m *Module) ContractCreateBlockNumber(db sdk.StateDBReader) uint64 {
	c, _ := contracts.NewCounter(sdkcontracts.NewEVM(types.NewStateDBWrapper(db), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	return c.GetCreateBlock()
}

func (m *Module) RegistryUpgradeHandler(registrar module.UpgradeRegistrar) error {
	registrar.RegisterUpgradeHandler(ModuleName, 0, func(ctx sdk.WorkerContext) error {
		if err := m.InitGenesis(ctx, ctx.StateDB(), ctx.Backend().ChainConfig(), nil); err != nil {
			return err
		}

		c, _ := contracts.NewCounter(sdkcontracts.NewEVM(ctx.StateDB(), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
		c.SetCreateBlock(ctx.Header().Number.Uint64())
		return nil
	})
	registrar.RegisterUpgradeHandler(ModuleName, ModuleVersion, func(ctx sdk.WorkerContext) error {
		m.logger.Info("Update contract version", "version", ModuleVersion)
		c, _ := contracts.NewCounter(sdkcontracts.NewEVM(ctx.StateDB(), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
		c.SetVersion(ModuleVersion)
		return nil
	})
	return nil
}
