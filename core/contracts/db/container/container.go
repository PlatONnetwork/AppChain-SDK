package container

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"reflect"
)

var (
	initContainerType = reflect.TypeOf((*Container)(nil)).Elem()
)

type Container interface {
	Prefix() []byte
	InitContainer(name []byte, addr common.Address, statedb vm.StateDB) Container
}
