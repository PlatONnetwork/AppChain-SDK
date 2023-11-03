package store

import "io"

type Store interface {
	// GetKVStore returns the KVStore for the given store key.
	GetKVStore(storeKey string) KVStore

	io.Closer
}

// KVStore defines the core storage primitive for modules to read and write state.
type KVStore interface {
	GetStoreKey() string

	// Has retrieves if a key is present in the key-value data store.
	Has(key []byte) (bool, error)

	// Get returns a value for a given key from the store.
	Get(key []byte) ([]byte, error)

	// Set sets a key/value entry to the store.
	Set(key, value []byte) error

	// Delete deletes the key from the store.
	Delete(key []byte) error

	Iteratee
}
