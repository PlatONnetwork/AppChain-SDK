package l1

import (
	"encoding/json"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestModule(t *testing.T) {
	store := memorydb.New()
	l1 := NewL1Module(store)
	g := L1ConfigParams{
		ChainID:    big.NewInt(11),
		State:      common.BigToAddress(big.NewInt(1)),
		Checkpoint: common.BigToAddress(big.NewInt(2)),
	}
	raw, _ := json.Marshal(g)
	l1.InitGenesis(nil, nil, nil, raw)
	chainId, err := l1.db.getChainID()
	require.Nil(t, err)
	require.Equal(t, g.ChainID, chainId)
	addr, err := l1.db.getStateAddress()
	require.Nil(t, err)
	require.Equal(t, g.State, addr)
	addr, err = l1.db.getCheckpointAddress()
	require.Nil(t, err)
	require.Equal(t, g.Checkpoint, addr)
}
