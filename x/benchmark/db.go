package benchmark

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/store"
)

var (
	namespace = "benchmark"
)

type DB struct {
	kv store.KVStore
}

func NewDB(db store.Store) *DB {
	return &DB{
		kv: db.GetKVStore(namespace),
	}
}

func (db *DB) Set(info *BlockInfo) error {
	k := db.key(info.Number)
	v, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return db.kv.Set(k, v)
}

func (db *DB) Get(num uint64) (*BlockInfo, error) {
	k := db.key(num)
	v, err := db.kv.Get(k)
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New("not found")
	}
	var info BlockInfo
	err = json.Unmarshal(v, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (db *DB) GetRaw(num uint64) ([]byte, error) {
	k := db.key(num)
	v, err := db.kv.Get(k)
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New("not found")
	}
	return v, nil
}

func (db *DB) key(num uint64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], num)
	return buf[:]
}
