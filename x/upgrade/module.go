package upgrade

import (
	"encoding/json"
	"fmt"
	"math/big"

	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	ModuleName           = "upgrade"
	ModuleVersion uint64 = 0
)

var (
	_ module.Module             = (*Module)(nil)
	_ module.ContractModule     = (*Module)(nil)
	_ module.BeginBlockerModule = (*Module)(nil)
	_ module.GenesisModule      = (*Module)(nil)
)

type Module struct {
	logger log.Logger

	upgradeHandlers      map[string]map[uint64]module.UpgradeHandler
	moduleValidNumberMap map[string]uint64
}

func NewModule() *Module {
	return &Module{
		logger: log.New("module", ModuleName),

		upgradeHandlers:      make(map[string]map[uint64]module.UpgradeHandler, 0),
		moduleValidNumberMap: make(map[string]uint64, 0),
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}

func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	db.SetNonce(m.Address(), 1)

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
	upgradeContract, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
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

func (m *Module) ContractCreateBlockNumber(db sdk.StateDBReader) uint64 {
	upgradeContract, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(coretypes.NewStateDBWrapper(db), big.NewInt(0)), sdkcontracts.NewContract(m, m), true)
	return upgradeContract.GetCreateBlock()
}

func (m *Module) BeginBlock(ctx sdk.WorkerContext) error {
	// TODO: executing upgrade plan when fast sync finish

	blockNumber := ctx.Header().Number
	upgradeContract, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(ctx.StateDB(), blockNumber), sdkcontracts.NewContract(m, m), true)
	plans, err := upgradeContract.GetUpgradePlan(ctx.Header().Number.Uint64())
	if err != nil {
		m.logger.Error("Failed to get plans", "height", ctx.Header().Number, "err", err)
		return err
	}
	m.logger.Info("Begin block", "plans", len(plans), "number", ctx.Header().Number)

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
		if err := upgradeContract.SetUpgradePlanDone(blockNumber.Uint64()); err != nil {
			panic(err)
		}

		vn, err := m.GetModuleValidNumberMap(ctx.StateDB())
		if err != nil || len(vn) == 0 {
			panic(fmt.Sprintf("empty module valid number map(err: %v)", err))
		}
		m.moduleValidNumberMap = vn
	}
	return nil
}

func (m *Module) RegisterUpgradeHandler(moduleName string, version uint64, handle module.UpgradeHandler) error {
	m.logger.Info("Register upgrade handler", "registeredModule", moduleName, "version", version)
	if m.upgradeHandlers[moduleName] == nil {
		m.upgradeHandlers[moduleName] = make(map[uint64]module.UpgradeHandler, 0)
	}
	m.upgradeHandlers[moduleName][version] = handle
	return nil
}

func (m *Module) SetModuleValidNumberMap(db sdk.StateDB, vn module.ValidNumberMap) error {
	buf, _ := json.Marshal(&vn)
	m.logger.Info("Set module valid number map", "vn", string(buf))
	upgradeContract, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	return upgradeContract.SetModuleValidNumberMap(vn)
}

func (m *Module) GetModuleValidNumberMap(db sdk.StateDBReader) (module.ValidNumberMap, error) {
	upgradeContract, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(coretypes.NewStateDBWrapper(db), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	return upgradeContract.GetModuleValidNumberMap()
}

func (m *Module) IsModuleValid(db sdk.StateDBReader, moduleName string, blockNumber uint64) bool {
	if len(m.moduleValidNumberMap) == 0 {
		vn, err := m.GetModuleValidNumberMap(db)
		if err != nil || len(vn) == 0 {
			panic(fmt.Sprintf("empty module valid number map(err: %v)", err))
		}
		m.moduleValidNumberMap = vn
	}
	if validNumber, ok := m.moduleValidNumberMap[moduleName]; ok {
		return blockNumber >= validNumber
	}
	return false
}
