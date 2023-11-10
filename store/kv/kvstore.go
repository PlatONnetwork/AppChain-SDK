package kv

import (
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/store"
)

const StorePrefixTpl = "s/k:%s/"

var _ store.KVStore = (*KVStore)(nil)

type KVStore struct {
	db       store.Database
	storeKey string
}

func NewKVStore(db store.Database, storeKey string) *KVStore {
	return &KVStore{
		db:       db,
		storeKey: storeKey,
	}
}

func (kv *KVStore) GetStoreKey() string {
	return kv.storeKey
}

func (kv *KVStore) Has(key []byte) (bool, error) {
	return kv.db.Has(prependStoreKey(kv.storeKey, key))
}

func (kv *KVStore) Get(key []byte) ([]byte, error) {
	return kv.db.Get(prependStoreKey(kv.storeKey, key))
}

func (kv *KVStore) Set(key, value []byte) error {
	return kv.db.Put(prependStoreKey(kv.storeKey, key), value)
}

func (kv *KVStore) Delete(key []byte) error {
	return kv.db.Delete(prependStoreKey(kv.storeKey, key))
}

func (kv *KVStore) NewIterator(prefix, start []byte) store.Iterator {
	return kv.db.NewIterator(prefix, start)
}
func (kv *KVStore) NewBatch() store.Batch {
	return &batch{
		batch:    kv.db.NewBatch(),
		storeKey: kv.storeKey,
	}
}

type batch struct {
	batch    store.Batch
	storeKey string
}

func (b *batch) Put(key []byte, value []byte) error {
	return b.batch.Put(prependStoreKey(b.storeKey, key), value)
}

func (b *batch) Delete(key []byte) error {
	return b.batch.Delete(prependStoreKey(b.storeKey, key))
}

func (b *batch) ValueSize() int {
	return b.batch.ValueSize()
}

func (b *batch) Write() error {
	return b.batch.Write()
}

func (b *batch) Reset() {
	b.batch.Reset()
}

func storePrefix(storeKey string) []byte {
	return []byte(fmt.Sprintf(StorePrefixTpl, storeKey))
}

func prependStoreKey(storeKey string, key []byte) []byte {
	return append(storePrefix(storeKey), key...)
}
