package contracts

import (
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

func (c *L2StateSender) addLogL2StateSyncedEvent(id *big.Int, sender basecommon.Address, receiver basecommon.Address, callData []byte) error {
	c.EmitL2StateSyncedEvent(id, sender, receiver, callData)
	return nil
}
