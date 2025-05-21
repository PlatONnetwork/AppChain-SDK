package upgrade

import (
	"encoding/json"
	"fmt"
	"math/big"

	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/storage"
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
	_ module.InitModule         = (*Module)(nil)
)

type Module struct {
	kv     *storage.Storage
	logger log.Logger

	upgradeHandlers      map[string]map[uint64]module.UpgradeHandler
	localVersionMap      module.VersionMap
	moduleValidNumberMap module.ValidNumberMap

	isContractModule func(string) bool
}

func NewModule(store store.Store) *Module {
	return &Module{
		kv:     storage.NewStorage(store.GetKVStore(ModuleName)),
		logger: log.New("module", ModuleName),

		upgradeHandlers:      make(map[string]map[uint64]module.UpgradeHandler, 0),
		localVersionMap:      make(module.VersionMap, 0),
		moduleValidNumberMap: make(module.ValidNumberMap, 0),
	}
}

func (m *Module) SetIsContractModule(isContractModule func(string) bool) {
	m.isContractModule = isContractModule
}

func (m *Module) Init(ctx sdk.InitContext) error {
	localVm, err := m.kv.GetVersionMap()
	if err != nil {
		m.logger.Error("Failed to get local version map", "err", err)
		return err
	}
	m.localVersionMap = localVm
	return nil
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
	upgrade, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	upgrade.SetOwner(config.Owner)
	upgrade.InitGenesis(config.CreateBlock)
	m.logger.Info("Init genesis", "owner", config.Owner, "createBlock", config.CreateBlock)

	vn, vm := m.GetModuleInitValidNumberMap(chainConfig, db)

	buf, _ := json.Marshal(&vm)
	m.logger.Info("Set version map to local storage", "vm", string(buf))
	m.kv.SetVersionMap(vm)
	if err := upgrade.SetModuleValidNumberMap(vn); err != nil {
		return err
	}
	return upgrade.SetModuleVersionMap(vm)
}

func (m *Module) Address() common.Address {
	return constants.UpgradeAddress
}

func (m *Module) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	upgrade, _ := contracts.NewUpgrade(evm, contract, readOnly)
	return upgrade.Run(input)
}

func (m *Module) ContractCreateBlockNumber(db sdk.StateDBReader) uint64 {
	upgrade, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(coretypes.NewStateDBWrapper(db), big.NewInt(0)), sdkcontracts.NewContract(m, m), true)
	return upgrade.GetCreateBlock()
}

func (m *Module) BeginBlock(ctx sdk.WorkerContext) error {
	localVm := m.localVersionMap

	blockNumber := ctx.Header().Number
	upgrade, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(ctx.StateDB(), blockNumber), sdkcontracts.NewContract(m, m), true)
	stateVm, err := upgrade.GetModuleVersionMap()
	if err != nil {
		m.logger.Error("Failed to get state version map", "err", err)
		return err
	}

	for name, ver := range stateVm {
		if lver, ok := localVm[name]; ok && ver <= lver {
			continue
		}
		localVm[name] = ver

		// Contract module's states store in `StateDB`, so skip it after fast sync.
		if m.isContractModule(name) {
			continue
		}

		if m.upgradeHandlers[name] == nil {
			panic(fmt.Sprintf("Cannot found the module's upgrade handler(name: %s,version:%d,height:%d)", name, ver, blockNumber))
		}
		handlers := m.upgradeHandlers[name]

		for i := localVm[name]; i <= ver; i++ {
			err = handlers[i](ctx)
			if err != nil {
				panic(fmt.Sprintf("Failed to execute upgrade module handler(name:%s,version:%d,err:%v)", name, i, err))
			}
			m.logger.Info("Success invoke upgrade handler after fast sync", "upgradedModule", name, "version", i)
		}
	}

	plans, err := upgrade.GetUpgradePlan(ctx.Header().Number.Uint64())
	if err != nil {
		m.logger.Error("Failed to get plans", "height", ctx.Header().Number, "err", err)
		return err
	}
	m.logger.Info("Begin block", "plans", len(plans), "number", ctx.Header().Number)

	for _, plan := range plans {
		m.logger.Info("Begin block", "plan", plan.String())
		for _, mod := range plan.Modules {
			if ver, ok := localVm[mod.ModuleName]; ok && ver >= mod.Version {
				continue
			}

			handlerFound := false
			if handlers, ok := m.upgradeHandlers[mod.ModuleName]; ok {
				if handler, found := handlers[mod.Version]; found && handler != nil {
					handlerFound = true
					err = handler(ctx)
					if err != nil {
						panic(fmt.Sprintf("Failed to execute upgrade module handler(name:%s,version:%d,height:%d,err:%v)", mod.ModuleName, mod.Version, plan.Height, err))
					}
					localVm[mod.ModuleName] = mod.Version
					stateVm[mod.ModuleName] = mod.Version
					m.logger.Info("Module success upgraded", "upgradedModule", mod.ModuleName, "version", mod.Version, "height", plan.Height)
				}
			}
			if !handlerFound {
				panic(fmt.Sprintf("Cannot found the module's upgrade handler(name: %s,version:%d,height:%d)", mod.ModuleName, mod.Version, plan.Height))
			}
		}
	}
	if len(plans) > 0 {
		if err := upgrade.SetUpgradePlanDone(blockNumber.Uint64()); err != nil {
			panic(err)
		}

		m.localVersionMap = localVm
		if err := upgrade.SetModuleVersionMap(stateVm); err != nil {
			return err
		}

		vn, err := m.GetModuleValidNumberMap(ctx.StateDB())
		if err != nil || len(vn) == 0 {
			panic(fmt.Sprintf("empty module valid number map(err: %v)", err))
		}
		m.moduleValidNumberMap = vn
	}
	return nil
}

func (m *Module) EndBlock(ctx *sdk.WorkerContext) error {
	if err := m.kv.SetVersionMap(m.localVersionMap); err != nil {
		m.logger.Error("Failed to set local version map", "err", err)
		return err
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
	upgradeContract, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	return upgradeContract.SetModuleValidNumberMap(vn)
}

func (m *Module) GetModuleValidNumberMap(db sdk.StateDBReader) (module.ValidNumberMap, error) {
	upgradeContract, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(coretypes.NewStateDBWrapper(db), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	return upgradeContract.GetModuleValidNumberMap()
}

func (m *Module) SetModuleVersionMap(db sdk.StateDB, vm module.VersionMap) error {
	upgrade, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	return upgrade.SetModuleVersionMap(vm)
}

func (m *Module) GetModuleVersionMap(db sdk.StateDBReader) (module.VersionMap, error) {
	upgrade, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(coretypes.NewStateDBWrapper(db), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	return upgrade.GetModuleVersionMap()
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

func (m *Module) GetModuleInitValidNumberMap(chainConfig *params.ChainConfig, db sdk.StateDBReader) (module.ValidNumberMap, module.VersionMap) {
	vn := make(module.ValidNumberMap)
	vm := make(module.VersionMap)
	for name, data := range chainConfig.Modules {
		vn[name] = 0
		var gen types.GenesisConfig
		json.Unmarshal(data, &gen)
		if m.isContractModule(name) {
			vn[name] = gen.CreateBlock
		}
		vm[name] = gen.Version
	}
	return vn, vm
}
