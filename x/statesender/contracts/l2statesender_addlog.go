package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	baselog "github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/status-im/keycard-go/hexutils"
	"math/big"
)

func (c *L2StateSender) addLogL2StateSyncedEvent(id *big.Int, sender basecommon.Address, receiver basecommon.Address, callData []byte) error {
	log, err := c.EmitL2StateSyncedEvent(id, sender, receiver, callData)
	if nil != err {
		baselog.Error("Failed to emit L2StateSyncedEvent", "id", id, "sender", sender.Hex(), "receiver", receiver.Hex(), "callData", hexutils.BytesToHex(callData), "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("L2StateSender: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}
