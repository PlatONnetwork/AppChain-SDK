package storage

import (
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/types"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
)

type Storage struct {
	kv store.KVStore
}

func NewStorage(kv store.KVStore) *Storage {
	return &Storage{
		kv: kv,
	}
}

func (s *Storage) SetCheckpoint(blockNumber uint64, checkpoint *types.StorageCheckpointData) error {
	encoded, err := rlp.EncodeToBytes(checkpoint)
	if err != nil {
		return err
	}
	return s.kv.Set(checkpointKey(blockNumber), encoded)
}

func (s *Storage) GetCheckpoint(blockNumber uint64) (*types.StorageCheckpointData, error) {
	var checkpoint types.StorageCheckpointData
	encoded, err := s.kv.Get(checkpointKey(blockNumber))
	if err != nil {
		return nil, err
	}
	if err := rlp.DecodeBytes(encoded, &checkpoint); err != nil {
		return nil, err
	}
	return &checkpoint, nil
}

func checkpointKey(blockNumber uint64) []byte {
	return []byte(fmt.Sprintf(types.CheckpointKeyTpl, blockNumber))
}
