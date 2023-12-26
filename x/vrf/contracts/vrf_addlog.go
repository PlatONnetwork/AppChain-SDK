package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	baselog "github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
)

func (c *VRFManager) addLogVRFNonceAddedEvent(block *big.Int, nonce []byte) error {
	log, err := c.EmitVRFNonceAddedEvent(block, nonce)
	if nil != err {
		baselog.Error("Failed to emit VRFNonceAddedEvent", "block", block.Uint64(), "nonce", hexutils.BytesToHex(nonce), "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("VRFManager: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}
