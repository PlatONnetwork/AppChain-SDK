package l1

import (
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

var (
	L1GenesisDBName      = "l1sync"
	chainIdKey           = []byte("chainId")
	stateAddressKey      = []byte("stateAddress")
	checkpointAddressKey = []byte("checkpointAddress")
)

type L1GenesisDB struct {
	db store.KVStore
}

func NewL1GenesisDB(db store.Store) *L1GenesisDB {
	return &L1GenesisDB{
		db: db.GetKVStore(L1GenesisDBName),
	}
}
func (l *L1GenesisDB) SetChainID(chainId *big.Int) error {
	return l.db.Set(chainIdKey, chainId.Bytes())
}

func (l *L1GenesisDB) GetChainID() (*big.Int, error) {
	value, err := l.db.Get(chainIdKey)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(value), nil
}

func (l *L1GenesisDB) SetStateAddress(addr common.Address) error {
	return l.db.Set(stateAddressKey, addr.Bytes())
}

func (l *L1GenesisDB) GetStateAddress() (common.Address, error) {
	value, err := l.db.Get(stateAddressKey)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(value), nil
}

func (l *L1GenesisDB) SetCheckpointAddress(addr common.Address) error {
	return l.db.Set(checkpointAddressKey, addr.Bytes())
}

func (l *L1GenesisDB) GetCheckpointAddress() (common.Address, error) {
	value, err := l.db.Get(checkpointAddressKey)
	if err != nil {
		return common.Address{}, err
	}
	return common.BytesToAddress(value), nil
}
