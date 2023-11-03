package storage

import (
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/PlatON-Go/ethdb/leveldb"
)

var (
	_ store.Store    = (*Storage)(nil)
	_ store.Database = (*Storage)(nil)
)

type Storage struct {
	db *leveldb.Database
}

func NewStorage(file string, cache, handles int, namespace string) (*Storage, error) {
	db, err := leveldb.New(file, cache, handles, namespace, false)
	if err != nil {
		return nil, err
	}
	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) GetKVStore(storeKey string) store.KVStore {
	return kv.NewKVStore(s, storeKey)
}

func (s *Storage) Has(key []byte) (bool, error) {
	return s.db.Has(key)
}

func (s *Storage) Get(key []byte) ([]byte, error) {
	return s.db.Get(key)
}

func (s *Storage) Put(key, value []byte) error {
	return s.db.Put(key, value)
}

func (s *Storage) Delete(key []byte) error {
	return s.db.Delete(key)
}

func (s *Storage) NewIterator(prefix, start []byte) store.Iterator {
	return s.db.NewIterator(prefix, start)
}

func (s *Storage) Close() error {
	return s.db.Close()
}
