package l1

import (
	"encoding/json"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

type Genesis struct {
	ChainID    *big.Int       `json:"chainId"`
	State      common.Address `json:"state"`
	CheckPoint common.Address `json:"checkpoint"`
}

type L1 struct {
	db *L1GenesisDB
}

func NewL1(db store.Store) *L1 {
	return &L1{
		db: NewL1GenesisDB(db),
	}
}

func (l *L1) Name() string {
	return "l1"
}

func (l *L1) InitGenesis(ctx sdk.Context, db sdk.StateDB, genesis *core.Genesis, data json.RawMessage) {
	var g Genesis
	raw, _ := data.MarshalJSON()
	json.Unmarshal(raw, &g)
	l.db.SetChainID(g.ChainID)
	l.db.SetStateAddress(g.State)
	l.db.SetCheckpointAddress(g.CheckPoint)
}
