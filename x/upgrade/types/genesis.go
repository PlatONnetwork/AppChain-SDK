package types

import (
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

type GenesisConfig struct {
	module.ModuleGenesisConfig
	Owner       common.Address `json:"owner"`
}
