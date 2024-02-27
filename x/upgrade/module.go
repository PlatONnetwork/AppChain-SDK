package upgrade

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	ModuleName           = "upgrade"
	ModuleVersion uint64 = 1
)

var (
	_ module.Module             = (*Module)(nil)
	_ module.ContractModule     = (*Module)(nil)
	_ module.BeginBlockerModule = (*Module)(nil)
	_ module.GenesisModule      = (*Module)(nil)
)

type Module struct {
	logger log.Logger

	upgradeHandlers map[string]map[uint64]module.UpgradeHandler
	initVersionMap  module.VersionMap
}

func NewModule() *Module {
	return &Module{
		logger: log.New("module", ModuleName),

		upgradeHandlers: make(map[string]map[uint64]module.UpgradeHandler, 0),
		initVersionMap:  make(module.VersionMap, 0),
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}

func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	var config types.GenesisConfig
	raw, err := data.MarshalJSON()
	if err != nil {
		m.logger.Error("Failed to MarshalJSON bytes", "err", err)
		return err
	}
	if err := json.Unmarshal(raw, &config); err != nil {
		m.logger.Error("Failed to unmarshal upgrade.GenesisConfig", "err", err)
		return err
	}
	upgradeContract, _ := contracts.NewUpgrade(vm.NewEVM(vm.BlockContext{GasLimit: math.MaxUint64}, vm.TxContext{}, db, &params.ChainConfig{}, vm.Config{}, nil), vm.NewContract(m, m, big.NewInt(0), math.MaxUint64), false)
	upgradeContract.SetOwner(config.Owner)
	upgradeContract.SetCreateBlock(config.CreateBlock)
	m.logger.Info("Init genesis", "owner", config.Owner, "createBlock", config.CreateBlock)
	return nil
}

func (m *Module) Address() common.Address {
	return constants.UpgradeAddress
}

func (m *Module) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	upgradeContract, _ := contracts.NewUpgrade(evm, contract, readOnly)
	return upgradeContract.Run(input)
}

func (m *Module) ContractCreateBlockNumber(statedb sdk.StateDB) uint64 {
	upgradeContract, _ := contracts.NewUpgrade(vm.NewEVM(vm.BlockContext{GasLimit: math.MaxUint64}, vm.TxContext{}, statedb, &params.ChainConfig{}, vm.Config{}, nil), vm.NewContract(m, m, big.NewInt(0), math.MaxUint64), true)
	return upgradeContract.GetCreateBlock()
}

func (m *Module) BeginBlock(ctx sdk.WorkerContext) error {
	// TODO: executing upgrade plan when fast sync finish

	upgradeContract, _ := contracts.NewUpgrade(vm.NewEVM(vm.BlockContext{GasLimit: math.MaxUint64}, vm.TxContext{}, ctx.StateDB(), &params.ChainConfig{}, vm.Config{}, nil), vm.NewContract(m, m, big.NewInt(0), math.MaxUint64), true)
	plans, err := upgradeContract.GetUpgradePlan(ctx.Header().Number.Uint64())
	if err != nil {
		m.logger.Error("Failed to get plans", "height", ctx.Header().Number, "err", err)
		return err
	}

	for _, plan := range plans {
		for _, mod := range plan.Modules {
			handlerFound := false
			if handlers, ok := m.upgradeHandlers[mod.ModuleName]; ok {
				if handler, found := handlers[mod.Version]; found && handler != nil {
					handlerFound = true
					err = handler(ctx)
					if err != nil {
						panic(fmt.Sprintf("Failed to execute upgrade module handler(name:%s,version:%d,height:%d,err:%v)", mod.ModuleName, mod.Version, plan.Height, err))
					}
					m.logger.Info("Module success upgraded", "upgradedModule", mod.ModuleName, "version", mod.Version, "height", plan.Height)
				}
			}
			if !handlerFound {
				panic(fmt.Sprintf("Cannot found the module's upgrade handler(name: %s,version:%d,height:%d)", mod.ModuleName, mod.Version, plan.Height))
			}
		}
	}
	if len(plans) > 0 {
		upgradeContract.SetUpgradePlanDone(ctx.Header().Number.Uint64())
	}
	return nil
}

func (m *Module) RegisterUpgradeHandler(module string, version uint64, handle module.UpgradeHandler) error {
	m.logger.Info("Register upgrade handler", "registeredModule", module, "version", version)
	m.upgradeHandlers[module][version] = handle
	return nil
}
