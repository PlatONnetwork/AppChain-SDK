package sync

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"math/big"
)

var (
	L1SyncDBName = "l1sync"
	lastBlockKey = []byte("lastBlock")
	maxSyncKey   = []byte("maxSyncId")
	eventKey     = []byte("event")
	batchSize    = 1024 * 1024
)

func encodeEventKey(id *big.Int) []byte {
	return append(eventKey, math.PaddedBigBytes(id, 32)...)
}

func decodeEventKey(key []byte) *big.Int {
	return new(big.Int).SetBytes(key[len(eventKey):])
}

type L1SyncDB struct {
	db store.KVStore
}

func NewL1SyncDB(db store.Store) *L1SyncDB {
	return &L1SyncDB{
		db: db.GetKVStore(L1SyncDBName),
	}
}

func (l *L1SyncDB) SetLastBlockNumber(start *big.Int) error {
	return l.db.Set(lastBlockKey, start.Bytes())
}

func (l *L1SyncDB) SetMaxSyncId(id *big.Int) error {
	return l.db.Set(maxSyncKey, id.Bytes())
}

func (l *L1SyncDB) GetMaxSyncId() (*big.Int, error) {
	value, err := l.db.Get(maxSyncKey)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(value), nil
}

func (l *L1SyncDB) LastBlockNumber() (*big.Int, error) {
	value, err := l.db.Get(lastBlockKey)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(value), nil
}

func (l *L1SyncDB) WriteStateSenderEvent(events []*StateSender) error {
	batch := l.db.NewBatch()
	for _, event := range events {
		value, err := rlp.EncodeToBytes(event)
		if err != nil {
			return err
		}
		batch.Put(encodeEventKey(event.Id), value)
		if batch.ValueSize() > batchSize {
			batch.Write()
		}
	}
	batch.Write()
	return nil
}

func (l *L1SyncDB) ClearStateSenderHistory(target *big.Int) error {
	batch := l.db.NewBatch()
	it := l.db.NewIterator(nil, nil)
	for it.Next() {
		id := new(big.Int).SetBytes(it.Key())
		if id.Cmp(target) >= 0 {
			break
		}
		batch.Delete(it.Key())
		if batch.ValueSize() > batchSize {
			batch.Write()
		}
	}
	return batch.Write()
}

func (l *L1SyncDB) FindStateSenderEvent(start *big.Int, end *big.Int) ([]*StateSender, error) {
	var events []*StateSender
	it := l.db.NewIterator(eventKey, math.PaddedBigBytes(start, 32))
	next := new(big.Int).SetBytes(start.Bytes())
	for it.Next() {
		if len(events) == 0 {
			if decodeEventKey(it.Key()).Cmp(next) != 0 {
				return nil, errors.New(fmt.Sprintf("database loss of data, expect:%s acutal:%s cmp:%d", hex.EncodeToString(next.Bytes()), hex.EncodeToString(decodeEventKey(it.Key()).Bytes()), new(big.Int).SetBytes(it.Key()).Cmp(next)))
			}
		}
		if decodeEventKey(it.Key()).Cmp(end) > 0 {
			break
		}
		var event StateSender
		if err := rlp.DecodeBytes(it.Value(), &event); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func (l *L1SyncDB) GetStateSenderEvent(start *big.Int) (*StateSender, error) {
	var event *StateSender
	value, err := l.db.Get(encodeEventKey(start))
	if err != nil {
		return nil, err
	}
	err = rlp.DecodeBytes(value, &event)
	if err != nil {
		return nil, err
	}
	return event, nil
}
