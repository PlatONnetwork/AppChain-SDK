package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestDB(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db), nil)
	statedb.CreateAccount(from)
	statedb.CreateAccount(to)
	evm := newEVM(statedb, vm.Config{}, vm.TxContext{}, new(contractsApp))
	contract := vm.NewContract(vm.AccountRef(from), vm.AccountRef(to), nil, 10000000)
	stateReceiver, _ := NewStateReceiver(evm, contract, false)
	stateReceiver.verifyQCFunc = func(qc *QuorumCert) error {
		return nil
	}
	stateReceiver.SetLastCommittedId(big.NewInt(10))
	require.Equal(t, big.NewInt(10), stateReceiver.GetLastCommittedId())
	stateReceiver.SetExecutedId(big.NewInt(11))
	require.Equal(t, big.NewInt(11), stateReceiver.getExecutedId())
	stateReceiver.SetCommitment(&StateSyncCommitment{
		StartId: big.NewInt(1),
		EndId:   big.NewInt(11),
		Root:    common.Hash{},
	})
	stateReceiver.SetLastCommittedId(big.NewInt(11))
	cm := stateReceiver.GetCommitment(big.NewInt(11))
	require.Equal(t, cm.StartId, big.NewInt(1))
	require.Equal(t, cm.EndId, big.NewInt(11))
	cm = stateReceiver.GetCommitment(big.NewInt(111))
	require.Nil(t, cm)
	cm = stateReceiver.FindCommitment(big.NewInt(5))
	require.Equal(t, cm.StartId, big.NewInt(1))
	require.Equal(t, cm.EndId, big.NewInt(11))
	cm = stateReceiver.FindCommitment(big.NewInt(3))
	require.Equal(t, cm.StartId, big.NewInt(1))
	require.Equal(t, cm.EndId, big.NewInt(11))
	cm = stateReceiver.FindCommitment(big.NewInt(111))
	require.Nil(t, cm)
}
