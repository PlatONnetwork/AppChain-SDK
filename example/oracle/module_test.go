package oracle

import (
	"encoding/json"
	election "github.com/PlatONnetwork/AppChain-SDK/example/election"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/stretchr/testify/require"
	"math/big"
	"sync"
	"testing"
)

func TestModule(t *testing.T) {
	election := election.NewModule()
	store := memorydb.New()
	oracleModule := NewModule(store, testutil.DefaultAccount[0].NodePrivateKey(), FixedRate{Value: 8})
	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{oracleModule})

	manager := module.NewManager(oracleModule, election, extraVote)

	//manager.SetOrderGenesis(oracleModule.Name())
	manager.SetElection(election.Name())
	app := testutil.NewApp(manager)
	config := GenesisConfig{
		Decimals: 3,
	}
	s, _ := json.Marshal(config)
	stack, backend, err := testutil.CreateCluster(testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		oracleModule.Name(): s,
	}, []common.Address{common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2))})
	require.Nil(t, err)
	require.Nil(t, stack[0].Start())
	require.Nil(t, backend[0].Start())
	var sg sync.WaitGroup
	sg.Add(1)
	sg.Wait()
}
