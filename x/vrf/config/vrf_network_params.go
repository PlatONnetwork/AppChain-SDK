package config

import "github.com/PlatONnetwork/PlatON-Go/common"

type VRFNetworkParams struct {
	GenesisVRFNonce common.Hash `json:"genesisVRFNonce"`
}

func DefualtVRFNetworkParams() *VRFNetworkParams {
	return &VRFNetworkParams{
		GenesisVRFNonce: common.BytesToHash([]byte("genesisVRFNonce")),
	}
}
