package container

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/db"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"reflect"
)

var (
	lengthKey = []byte("length")
)

type Array[T any] struct {
	Store  *db.Store
	length *uint32
}

func NewArray[T any](name []byte, addr common.Address, statedb vm.StateDB) *Array[T] {
	store := db.NewStore(name, addr, statedb)
	return &Array[T]{
		Store: store,
	}
}

func (a *Array[T]) Length() uint32 {
	if a.length == nil {
		len, err := db.GetState[uint32](a.Store, lengthKey)
		if err != nil {
			return 0
		}
		a.length = &len
	}
	return *a.length
}
func (a *Array[T]) SetLength(i uint32) error {
	a.length = &i
	return db.SetState(a.Store, lengthKey, i)
}

func (a *Array[T]) Push(elem T) error {
	index := a.Length()

	len := index + 1
	if err := a.SetLength(len); err != nil {
		return err
	}
	return a.set(index, elem)
}

func (a *Array[T]) MustPush(elem T) {
	err := a.Push(elem)
	if err != nil {
		contracts.Require(false, "Array: push value failed")
	}
}

func (a *Array[T]) set(i uint32, elem T) error {
	v := reflect.ValueOf(elem)
	ty := v.Type()
	if ty.Implements(initContainerType) {
		prefix := v.Interface().(Container).Prefix()
		return db.SetState(a.Store, i, prefix)
	}
	return db.SetState(a.Store, i, elem)
}

func (a *Array[T]) Replace(i uint32, elem T) error {
	return a.set(i, elem)
}

func (a *Array[T]) MustReplace(i uint32, elem T) {
	err := a.set(i, elem)
	if err != nil {
		contracts.Require(false, "Array: replace value failed")
	}
}

func (a *Array[T]) Index(i uint32) (T, error) {
	return a.get(i)
}

func (a *Array[T]) MustIndex(i uint32) T {
	t, err := a.get(i)
	if err != nil {
		contracts.Require(false, "Array: index value failed")
	}
	return t
}

func (a *Array[T]) get(i uint32) (T, error) {
	var t T
	v := reflect.ValueOf(t)
	ty := v.Type()
	if ty.Implements(initContainerType) {
		prefix, err := db.GetState[[]byte](a.Store, i)
		if err != nil || len(prefix) == 0 {
			return t, err
		}
		m, _ := v.Interface().(Container)

		impl := m.InitContainer(prefix, a.Store.Address, a.Store.StateDB)
		return impl.(T), nil
	}
	return db.GetState[T](a.Store, i)
}

func (a *Array[T]) CreatePrefix(index uint32) []byte {
	keyBuf, _ := a.Store.KeyEncoder.EncodeKey(a.Store.Prefix, index)
	return keyBuf
}

func (a *Array[T]) Prefix() []byte {
	return a.Store.Prefix
}
func (a *Array[T]) InitContainer(name []byte, addr common.Address, statedb vm.StateDB) Container {
	return NewArray[T](name, addr, statedb)
}
