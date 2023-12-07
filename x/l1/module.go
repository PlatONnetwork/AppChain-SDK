package l1

import (
	"encoding/json"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type Genesis struct {
	ChainID    *big.Int       `json:"chainId"`
	State      common.Address `json:"state"`
	Checkpoint common.Address `json:"checkpoint"`
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

func (l *L1) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) {
	var g Genesis
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON l1 Genesis bytes", "error", err)
	}

	if err := json.Unmarshal(raw, &g); nil != err {
		log.Error("Failed UnmarshalJSON l1 Genesis", "error", err)
	}

	l.db.SetChainID(g.ChainID)
	l.db.SetStateAddress(g.State)
	l.db.SetCheckpointAddress(g.Checkpoint)

	log.Info("Succeed init genesis", "module", l.Name(), "chainId", g.ChainID, "state", g.State.Hex(), "checkpoint", g.Checkpoint.Hex())
}
