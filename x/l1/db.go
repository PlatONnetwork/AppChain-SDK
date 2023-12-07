package l1

import (
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

var (
	L1ConfigParamsDBName     = "l1config"
	chainIdKey               = []byte("chainId")
	stateAddressKey          = []byte("stateAddress")
	checkpointAddressKey     = []byte("checkpointAddress")
	stakeManagerAddressKey   = []byte("stakeManagerAddress")
	depositManagerAddressKey = []byte("depositManagerAddress")
)

type l1GenesisDB struct {
	db store.KVStore
}

func newL1GenesisDB(db store.Store) *l1GenesisDB {
	return &l1GenesisDB{
		db: db.GetKVStore(L1ConfigParamsDBName),
	}
}
func (l *l1GenesisDB) setChainID(chainId *big.Int) error {
	return l.db.Set(chainIdKey, chainId.Bytes())
}

func (l *l1GenesisDB) getChainID() (*big.Int, error) {
	value, err := l.db.Get(chainIdKey)
	if err != nil {
		return nil, err
	}
	return new(big.Int).SetBytes(value), nil
}

func (l *l1GenesisDB) setStateAddress(addr common.Address) error {
	return l.db.Set(stateAddressKey, addr.Bytes())
}

func (l *l1GenesisDB) getStateAddress() (common.Address, error) {
	value, err := l.db.Get(stateAddressKey)
	if err != nil {
		return common.ZeroAddr, err
	}
	return common.BytesToAddress(value), nil
}

func (l *l1GenesisDB) setCheckpointAddress(addr common.Address) error {
	return l.db.Set(checkpointAddressKey, addr.Bytes())
}

func (l *l1GenesisDB) getCheckpointAddress() (common.Address, error) {
	value, err := l.db.Get(checkpointAddressKey)
	if err != nil {
		return common.ZeroAddr, err
	}
	return common.BytesToAddress(value), nil
}

func (l *l1GenesisDB) setStakeManagerAddress(addr common.Address) error {
	return l.db.Set(stakeManagerAddressKey, addr.Bytes())
}

func (l *l1GenesisDB) getStakeManagerAddress() (common.Address, error) {
	value, err := l.db.Get(stakeManagerAddressKey)
	if err != nil {
		return common.ZeroAddr, err
	}
	return common.BytesToAddress(value), nil
}

func (l *l1GenesisDB) setDepositManagerAddress(addr common.Address) error {
	return l.db.Set(depositManagerAddressKey, addr.Bytes())
}

func (l *l1GenesisDB) getDepositManagerAddress() (common.Address, error) {
	value, err := l.db.Get(depositManagerAddressKey)
	if err != nil {
		return common.ZeroAddr, err
	}
	return common.BytesToAddress(value), nil
}
