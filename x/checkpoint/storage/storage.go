package storage

import (
	"encoding/binary"
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi"
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

func (s *Storage) InsertCheckpoint(blockNumber uint64, checkpoint *types.StorageCheckpointData) error {
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

func (s *Storage) InsertExitEvent(exitEvent *contractsapi.ExitEvent) error {
	encoded, err := rlp.EncodeToBytes(exitEvent)
	if err != nil {
		return err
	}
	if err := s.kv.Set(exitEventKey(exitEvent.Epoch, exitEvent.L2StateSyncedEvent.Id.Uint64()), encoded); err != nil {
		return err
	}

	epochBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(epochBytes, exitEvent.Epoch)
	return s.kv.Set(exitEventLookupKey(exitEvent.L2StateSyncedEvent.Id.Uint64()), epochBytes)
}

func (s *Storage) GetExitEventsByEpoch(epoch uint64) ([]*contractsapi.ExitEvent, error) {
	exitEvents := make([]*contractsapi.ExitEvent, 0)
	it := s.kv.NewIterator([]byte{}, exitEventPrefix(epoch))
	for it.Next() {
		var exitEvent contractsapi.ExitEvent
		if err := rlp.DecodeBytes(it.Value(), &exitEvent); err != nil {
			return exitEvents,err
		}
		exitEvents = append(exitEvents, &exitEvent)
	}
	it.Release()
	return exitEvents, nil
}

func checkpointKey(blockNumber uint64) []byte {
	return []byte(fmt.Sprintf(types.CheckpointKeyTpl, blockNumber))
}

func exitEventKey(epoch, evID uint64) []byte {
	return []byte(fmt.Sprintf(types.ExitEventKeyTpl, epoch, evID))
}

func exitEventLookupKey(evID uint64) []byte {
	return []byte(fmt.Sprintf(types.ExitEventLookupKey, evID))
}

func exitEventPrefix(epoch uint64) []byte {
	return []byte(fmt.Sprintf(types.ExitEventPrefixKeyTpl, epoch))
}
