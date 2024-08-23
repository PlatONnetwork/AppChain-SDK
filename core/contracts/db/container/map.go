package container

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"reflect"
)

type Map[T any] struct {
	Store *db.Store
}

func NewMap[T any](name []byte, addr common.Address, statedb vm.StateDB) *Map[T] {
	store := db.NewStore(name, addr, statedb)
	return &Map[T]{
		Store: store,
	}
}

func (a *Map[T]) Set(key any, elem T) error {
	v := reflect.ValueOf(elem)
	ty := v.Type()
	if ty.Implements(initContainerType) {
		prefix := v.Interface().(Container).Prefix()
		return db.SetState(a.Store, key, prefix)
	}
	return db.SetState(a.Store, key, elem)
}

func (a *Map[T]) MustSet(key any, elem T) {
	err := a.Set(key, elem)
	if err != nil {
		contracts.Require(false, "Map: set value failed")
	}
}

func (a *Map[T]) Get(key any) (T, error) {
	var t T
	v := reflect.ValueOf(t)
	ty := v.Type()
	if ty.Implements(initContainerType) {
		prefix, err := db.GetState[[]byte](a.Store, key)
		if err != nil {
			return t, err
		}
		if len(prefix) == 0 {
			prefix = a.CreatePrefix(key)
		}
		m, _ := v.Interface().(Container)

		impl := m.InitContainer(prefix, a.Store.Address, a.Store.StateDB)
		return impl.(T), nil
	}

	return db.GetState[T](a.Store, key)
}

func (a *Map[T]) MustGet(key any) T {
	t, err := a.Get(key)
	if err != nil {
		contracts.Require(false, "Map: get value failed")
	}
	return t
}

func (a *Map[T]) CreatePrefix(key any) []byte {
	keyBuf, _ := a.Store.KeyEncoder.EncodeKey(a.Store.Prefix, key)
	return keyBuf
}

func (a *Map[T]) Prefix() []byte {
	return a.Store.Prefix
}
func (a *Map[T]) InitContainer(name []byte, addr common.Address, statedb vm.StateDB) Container {
	return NewMap[T](name, addr, statedb)
}
