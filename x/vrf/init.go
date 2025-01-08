package vrf

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf/config"
	vrfdb "github.com/PlatONnetwork/AppChain-SDK/x/vrf/db"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initGenesisVRFNonce(statedb sdk.StateDB, addr common.Address, chainConfig *params.ChainConfig, configParams *config.VRFNetworkParams) {
	// set genesis vrf nonce (32 byte)
	vrfdb.SetNonceAndProof(statedb, addr, 0, configParams.GenesisVRFNonce.Bytes())
}
