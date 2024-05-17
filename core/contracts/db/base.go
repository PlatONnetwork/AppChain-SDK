package db

import "github.com/PlatONnetwork/AppChain-SDK/core/contracts"

type Base[T any] struct {
	key   any
	store *Store
}

func NewBase[T any](store *Store, key any) *Base[T] {
	return &Base[T]{
		key:   key,
		store: store,
	}
}

func (b *Base[T]) Get() (T, error) {
	return GetState[T](b.store, b.key)
}

func (b *Base[T]) MustGet() T {
	t, err := GetState[T](b.store, b.key)
	contracts.Require(err == nil, "Base: get value failed")
	return t
}

func (b *Base[T]) Set(t T) error {
	return SetState(b.store, b.key, t)
}

func (b *Base[T]) MustSet(t T) {
	err := SetState(b.store, b.key, t)
	contracts.Require(err == nil, "Base: set value failed")
}
