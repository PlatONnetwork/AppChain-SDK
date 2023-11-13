package contracts

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

var (
	upgradePrefix = "__upgrade_"
	initKey       = []byte(fmt.Sprintf("%s%s", upgradePrefix, "init"))
	implementKey  = []byte(fmt.Sprintf("%s%s", upgradePrefix, "implement"))
)

func (c *Upgrade) setInit() {
	c.evm.StateDB.SetState(c.contract.Address(), initKey, []byte{1})
}

func (c *Upgrade) getInit() bool {
	value := c.evm.StateDB.GetState(c.contract.Address(), initKey)
	if len(value) == 0 {
		return false
	}
	return value[0] == 1
}

func (c *Upgrade) setImplement(addr common.Address) {
	c.evm.StateDB.SetState(c.contract.Address(), implementKey, addr.Bytes())
}

func (c *Upgrade) getImplement() common.Address {
	value := c.evm.StateDB.GetState(c.contract.Address(), implementKey)
	return common.BytesToAddress(value)
}
