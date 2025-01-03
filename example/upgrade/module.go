package upgrade

import (
	"encoding/json"
	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/contracts"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

const ModuleVersion uint64 = 0
const ModuleName = "upgrade-contract"

type Module struct {
}

func NewModule() *Module {
	return &Module{}
}
func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	upgrade, _ := contracts.NewUpgrade(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(vm.AccountRef(constants.UpgradeAddress), vm.AccountRef(constants.UpgradeAddress)), false)
	if err := upgrade.SetModuleValidNumberMap(module.ValidNumberMap{
		"election": 0,
		"upgrade":  0,
	}); err != nil {
		return err
	}

	err := upgrade.SetModuleVersionMap(module.VersionMap{"election": 0})
	if err != nil {
		return err
	}
	err = upgrade.AddUpgradePlan(contracts.IUpgradePlan{
		Name: "oracle-upgrade",
		Modules: []contracts.IUpgradeModule{
			{
				ModuleName: "oracle",
				Version:    0,
			},
		},
		Info:   "init oracle module",
		Height: 5,
		Status: 0,
	})
	if err != nil {
		return err
	}
	return nil
}

//新部署合约
//增加新接口
//增加新模块
