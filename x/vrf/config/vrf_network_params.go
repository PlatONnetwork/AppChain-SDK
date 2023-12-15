package config

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

type VRFNetworkParams struct {
	GenesisVRFNonce common.Hash `json:"genesisVRFNonce"`
}

func DefualtVRFNetworkParams() *VRFNetworkParams {
	return &VRFNetworkParams{
		GenesisVRFNonce: common.BytesToHash([]byte("genesisVRFNonce")),
	}
}

func (params *VRFNetworkParams) String() string {
	return fmt.Sprintf(`{"genesisVRFNonce": "%s"}`, params.GenesisVRFNonce.Hex())
}
