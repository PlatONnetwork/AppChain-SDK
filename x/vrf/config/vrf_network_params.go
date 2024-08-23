package config

import (
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

type VRFNetworkParams struct {
	module.ModuleGenesisConfig
	GenesisVRFNonce common.Hash `json:"genesisVRFNonce"`
}

func DefualtVRFNetworkParams() *VRFNetworkParams {
	return &VRFNetworkParams{
		ModuleGenesisConfig: module.ModuleGenesisConfig{CreateBlock: 0},
		GenesisVRFNonce:     common.BytesToHash([]byte("genesisVRFNonce")),
	}
}

func (params *VRFNetworkParams) String() string {
	return fmt.Sprintf(`{"createBlock": %d, "genesisVRFNonce": "%s"}`, params.CreateBlock, params.GenesisVRFNonce.Hex())
}
