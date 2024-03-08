package storage

import (
	"encoding/binary"
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/log"
)

var (
	versionMapKeyPrefix = "version/map/"
)

type Storage struct {
	kv store.KVStore
}

func NewStorage(kv store.KVStore) *Storage {
	return &Storage{kv: kv}
}

func (s *Storage) GetVersionMap() (module.VersionMap, error) {
	vm := module.VersionMap{}
	it := s.kv.NewIterator([]byte(versionMapKeyPrefix), nil)
	for it.Next() {
		version := binary.BigEndian.Uint64(it.Value())
		key := string(it.Key()[len(versionMapKeyPrefix):])
		vm[key] = version
		log.Info("Get version", "module", string(it.Key()), "value", common.Bytes2Hex(it.Value()), "version", version)
	}
	it.Release()
	return vm, nil
}

func (s *Storage) SetVersionMap(vm module.VersionMap) error {
	for name, version := range vm {
		if err := s.setVersion(name, version); err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) setVersion(moduleName string, version uint64) error {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	log.Info("Set version", "module", moduleName, "version", version, "data", fmt.Sprintf("%x", data[:]))
	return s.kv.Set(encodeKey(moduleName), data[:])
}

func encodeKey(moduleName string) []byte {
	return []byte(fmt.Sprintf("%s%s", versionMapKeyPrefix, moduleName))
}
