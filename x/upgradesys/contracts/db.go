package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
)

var (
	upgradePrefix   = "__upgrade_"
	initializingKey = []byte(fmt.Sprintf("%s%s", upgradePrefix, "initializing"))
	initializedKey  = []byte(fmt.Sprintf("%s%s", upgradePrefix, "initialized"))

	implementKey = []byte(fmt.Sprintf("%s%s", upgradePrefix, "implement"))
)

func Initializer(statedb vm.StateDB, address common.Address) error {
	if !GetInitializing(statedb, address) {
		return typesdk.NewRevertError("UPGRADE: contract is already initialized")
	}
	return nil
}

func OnlyInitialized(statedb vm.StateDB, address common.Address) error {
	if !GetInitialized(statedb, address) {
		return typesdk.NewRevertError("UPGRADE: contract is not initializing")
	}
	return nil
}

func GetInitialized(statedb vm.StateDB, address common.Address) bool {
	value := statedb.GetState(address, initializedKey)
	if len(value) == 0 {
		return false
	}
	return value[0] == 1
}

func GetInitializing(statedb vm.StateDB, address common.Address) bool {
	value := statedb.GetState(address, initializingKey)
	if len(value) == 0 {
		return false
	}
	return value[0] == 1
}

func (c *Upgrade) setInitialized() {
	c.evm.StateDB.SetState(c.contract.Address(), initializedKey, []byte{1})
}

func (c *Upgrade) setInitializing() {
	c.evm.StateDB.SetState(c.contract.Address(), initializingKey, []byte{1})
}

func (c *Upgrade) getInitialized() bool {
	return GetInitialized(c.evm.StateDB, c.contract.Address())
}

func (c *Upgrade) getInitializing() bool {
	return GetInitializing(c.evm.StateDB, c.contract.Address())
}

func (c *Upgrade) setImplement(addr common.Address) {
	c.evm.StateDB.SetState(c.contract.Address(), implementKey, addr.Bytes())
}

func (c *Upgrade) getImplement() common.Address {
	value := c.evm.StateDB.GetState(c.contract.Address(), implementKey)
	return common.BytesToAddress(value)
}
