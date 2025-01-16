package node

import (
	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/example/upgrade/node/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

const ModuleVersion uint64 = 0
const ModuleName = "node"

var (
	NodeAddress = common.BigToAddress(big.NewInt(2223))
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

//func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
//	return nil
//}

func (m *Module) Address() common.Address {
	return NodeAddress
}

func (m *Module) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	upgrade, _ := contracts.NewNode(evm, contract, readOnly)
	return upgrade.Run(input)
}

func (m *Module) ContractCreateBlockNumber(db sdk.StateDBReader) uint64 {
	upgrade, _ := contracts.NewNode(sdkcontracts.NewEVM(types.NewStateDBWrapper(db), big.NewInt(0)), sdkcontracts.NewContract(m, m), true)
	return upgrade.GetCreateBlock()
}
func (m *Module) RegistryUpgradeHandler(registrar module.UpgradeRegistrar) error {
	registrar.RegisterUpgradeHandler(ModuleName, 0, func(ctx sdk.WorkerContext) error {
		c, _ := contracts.NewNode(sdkcontracts.NewEVM(ctx.StateDB(), ctx.Header().Number), sdkcontracts.NewContract(m, m), false)
		c.InitGenesis(ctx.Header().Number.Uint64())
		m.logger.Info("Run upgrade handler success")
		return nil
	})
	registrar.RegisterUpgradeHandler(ModuleName, 1, func(ctx sdk.WorkerContext) error {
		m.logger.Info("Update contract version", "version", 1)
		c, _ := contracts.NewNode(sdkcontracts.NewEVM(ctx.StateDB(), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
		c.SetVersion(1)
		return nil
	})
	registrar.RegisterUpgradeHandler(ModuleName, 2, func(ctx sdk.WorkerContext) error {
		m.logger.Info("Update contract version", "version", 2)
		c, _ := contracts.NewNode(sdkcontracts.NewEVM(ctx.StateDB(), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
		c.SetVersion(2)
		return nil
	})
	registrar.RegisterUpgradeHandler(ModuleName, 3, func(ctx sdk.WorkerContext) error {
		m.logger.Info("Update contract version", "version", 3)
		c, _ := contracts.NewNode(sdkcontracts.NewEVM(ctx.StateDB(), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
		c.SetVersion(3)
		return nil
	})
	return nil
}

//version 1, 添加节点
//version 2, 查询节点
//version 3, 修改节点
