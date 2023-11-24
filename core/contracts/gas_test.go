package contracts

import (
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestStateDB(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db), nil)
	evm := &vm.EVM{StateDB: statedb}
	gas := uint64(100000)
	sdb := NewStateDB(evm, &vm.Contract{Gas: gas})
	sdb.SetState(common.BigToAddress(big.NewInt(1)), []byte{1}, []byte{2})
	t.Log(sdb.burn.Gas())
	require.Equal(t, gas-SstoreSetGas, sdb.burn.Gas())
	sdb.contract.Gas = gas
	sdb.SetState(common.BigToAddress(big.NewInt(1)), []byte{1}, []byte{2})
	t.Log(sdb.burn.Gas())
	require.Equal(t, gas-SstoreResetGas, sdb.burn.Gas())

	sdb.contract.Gas = gas
	sdb.SetState(common.BigToAddress(big.NewInt(1)), []byte{1}, common.BigToHash(big.NewInt(1)).Bytes())
	t.Log(sdb.burn.Gas())
	require.Equal(t, gas-1*SstoreSetGas-1*SstoreResetGas, sdb.burn.Gas())

	sdb.contract.Gas = gas
	sdb.AddLog(&types.Log{
		Address:     common.Address{},
		Topics:      []common.Hash{common.BigToHash(big.NewInt(1)), common.BigToHash(big.NewInt(2))},
		Data:        []byte{1, 2, 3},
		BlockNumber: 0,
		TxHash:      common.Hash{},
		TxIndex:     0,
		BlockHash:   common.Hash{},
		Index:       0,
		Removed:     false,
	})
	t.Log(sdb.burn.Gas())
	require.Equal(t, gas-GasLog-GasLogTopicGas*2-GasLogData*3, sdb.burn.Gas())

	sdb.contract.Gas = gas
	sdb.SetCode(common.BigToAddress(big.NewInt(2)), []byte{1})
	require.Equal(t, gas-GasQuickStep-GasCopy*1, sdb.burn.Gas())

}

func TestGas(t *testing.T) {
	burn := NewBurner(&vm.Contract{Gas: 0})
	ret, err := testExec(func() {
		burn.UseGas(10)
	})
	require.NotNil(t, err)
	require.Nil(t, ret)
	ret, err = testExec(func() {
		panic(typesdk.NewRevertError("hello"))
	})
	require.NotNil(t, err)
	require.Equal(t, []byte("hello"), ret)

	ret, err = testExec(func() {
		panic("panic")
	})
	require.NotNil(t, err)
	require.NotNil(t, ret)
}

func testExec(exec func()) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				if r, ok := e.(*typesdk.RevertError); ok {
					ret, err = r.ReturnData, vm.ErrExecutionReverted
				} else {
					ret, err = nil, e
				}
			default:
				ret, err = typesdk.UndefinedError, vm.ErrExecutionReverted
			}
		}
	}()
	exec()
	return nil, nil
}
