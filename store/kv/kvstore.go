package kv

import (
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/syndtr/goleveldb/leveldb/errors"
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
	val, err := kv.db.Get(prependStoreKey(kv.storeKey, key))
	if err != nil && err != errors.ErrNotFound {
		return []byte{}, err
	}
	return val, nil
}

func (kv *KVStore) Set(key, value []byte) error {
	return kv.db.Put(prependStoreKey(kv.storeKey, key), value)
}

func (kv *KVStore) Delete(key []byte) error {
	return kv.db.Delete(prependStoreKey(kv.storeKey, key))
}

func (kv *KVStore) NewIterator(prefix, start []byte) store.Iterator {
	return &kvIterator{
		storeKey: kv.storeKey,
		iter:     kv.db.NewIterator(prependStoreKey(kv.storeKey, prefix), start),
	}
}

func (kv *KVStore) NewBatch() store.Batch {
	return &batch{
		batch:    kv.db.NewBatch(),
		storeKey: kv.storeKey,
	}
}

type kvIterator struct {
	storeKey string
	iter     store.Iterator
}

func (k *kvIterator) Next() bool {
	return k.iter.Next()
}

func (k *kvIterator) Error() error {
	return k.iter.Error()
}

func (k *kvIterator) Key() []byte {
	return removeStoreKey(k.iter.Key(), k.storeKey)
}

func (k *kvIterator) Value() []byte {
	return k.iter.Value()
}

func (k *kvIterator) Release() {
	k.iter.Release()
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
func removeStoreKey(key []byte, storeKey string) []byte {
	return key[len(storePrefix(storeKey)):]
}
