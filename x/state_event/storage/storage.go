package storage

import (
	"encoding/binary"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x/state_event/types"
)

type Storage struct {
	store store.KVStore
}

func NewStorage(store store.Store) *Storage {
	return &Storage{
		store: store.GetKVStore(types.ModuleName),
	}
}

func (s *Storage) InsertLastProcessedEventBlock(blockNumber uint64) error {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, blockNumber)
	return s.store.Set(types.LastProcessedEventBlockKey, result)
}

func (s *Storage) GetLastProcessedEventsBlock() (uint64, error) {
	result, err := s.store.Get(types.LastProcessedEventBlockKey)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(result), nil
}
