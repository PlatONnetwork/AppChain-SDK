package oracle

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
)

var (
	name  = "oracle"
	qcKey = []byte("qc")
)

type QCRateDB struct {
	db store.KVStore
}

func NewQCRateDB(store store.Store) *QCRateDB {
	return &QCRateDB{
		db: store.GetKVStore(name),
	}
}

func (q *QCRateDB) InsertRate(blockHash common.Hash, rate uint64) error {
	data, _ := rlp.EncodeToBytes(rate)
	return q.db.Set(blockHash.Bytes(), data)
}

func (q *QCRateDB) Rate(blockHash common.Hash) (uint64, error) {
	data, err := q.db.Get(blockHash.Bytes())
	if err != nil {
		return 0, err
	}
	var rate uint64
	err = rlp.DecodeBytes(data, &rate)
	if err != nil {
		return 0, err
	}
	return rate, nil
}

func (q *QCRateDB) InsertQCRate(blockHash common.Hash) error {
	_, err := q.Rate(blockHash)
	if err != nil {
		return err
	}
	return q.db.Set(qcKey, blockHash.Bytes())
}

func (q *QCRateDB) QCRate() (common.Hash, uint64, error) {
	data, err := q.db.Get(qcKey)
	if err != nil {
		return common.Hash{}, 0, err
	}
	if len(data) == 0 {
		return common.Hash{}, 0, errors.New("invalid qc data")
	}
	hash := common.BytesToHash(data)
	rate, err := q.Rate(hash)
	if err != nil {
		return common.Hash{}, 0, err
	}
	return hash, rate, nil
}
