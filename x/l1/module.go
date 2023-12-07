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

type L1ConfigParams struct {
	ChainID        *big.Int       `json:"chainId"`
	State          common.Address `json:"state"`
	Checkpoint     common.Address `json:"checkpoint"`
	StakeManager   common.Address `json:"stakeManager"`
	DepositManager common.Address `json:"depositManager"`
}

type L1Module struct {
	db *l1GenesisDB
}

func NewL1Module(db store.Store) *L1Module {
	return &L1Module{
		db: newL1GenesisDB(db),
	}
}

func (l *L1Module) Name() string {
	return "l1"
}

func (l *L1Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) {
	var g L1ConfigParams
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON l1 L1ConfigParams bytes", "error", err)
	}

	if err := json.Unmarshal(raw, &g); nil != err {
		log.Error("Failed UnmarshalJSON l1 L1ConfigParams", "error", err)
	}

	l.db.setChainID(g.ChainID)
	l.db.setStateAddress(g.State)
	l.db.setCheckpointAddress(g.Checkpoint)
	l.db.setStakeManagerAddress(g.StakeManager)
	l.db.setDepositManagerAddress(g.DepositManager)

	log.Info("Succeed init genesis", "module", l.Name(), "chainId", g.ChainID, "state", g.State.Hex(), "checkpoint", g.Checkpoint.Hex(), "stakeManager", g.StakeManager.Hex(), "depositManager", g.DepositManager.Hex())
}

// extern

func (l *L1Module) GetChainID() (*big.Int, error) {
	return l.db.getChainID()
}
func (l *L1Module) GetStateAddress() (common.Address, error) {
	return l.db.getStateAddress()
}
func (l *L1Module) GetCheckpointAddress() (common.Address, error) {
	return l.db.getCheckpointAddress()
}
func (l *L1Module) GetStakeManagerAddress() (common.Address, error) {
	return l.db.getStakeManagerAddress()
}
func (l *L1Module) GetDepositManagerAddress() (common.Address, error) {
	return l.db.getDepositManagerAddress()
}
