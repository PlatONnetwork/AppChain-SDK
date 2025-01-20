package testutil

import (
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/stretchr/testify/require"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"testing"
)

func TestRunNode(t *testing.T) {
	go http.ListenAndServe("0.0.0.0:6060", nil)
	vals, err := NewValidator(DefaultAccount[0:1])
	require.Nil(t, err)
	manager := module.NewManager(vals)
	manager.SetElection(vals.Name())
	app := NewApp(manager)
	require.Nil(t, err)
	stack, backend, err := CreateCluster(DefaultAccount[0:1], DefaultAccount[0:1], []sdk.App{app}, nil, nil)
	require.Nil(t, err)
	require.Nil(t, stack[0].Start())
	require.Nil(t, backend[0].Start())
	var sg sync.WaitGroup
	sg.Add(1)
	sg.Wait()
}
