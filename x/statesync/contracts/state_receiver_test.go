package contracts

import (
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var (
	from   = common.BigToAddress(big.NewInt(101))
	to     = common.BigToAddress(big.NewInt(102))
	events []*StateSync
	tree   *merkle.MerkleTree
)

type stateReceiverContract struct {
}

func (c stateReceiverContract) Address() common.Address {
	return common.BigToAddress(big.NewInt(102))
}

func (c stateReceiverContract) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	stateReceiver, _ := NewStateReceiver(evm, contract, readOnly)
	stateReceiver.verifyQCFunc = func(qc *QuorumCert) error {
		return nil
	}
	return stateReceiver.Run(input)
}

func (c stateReceiverContract) ContractCreateBlockNumber(statedb sdk.StateDBReader) uint64 {
	return 0
}

type contractsApp struct {
}

func (c contractsApp) Contracts(statedb sdk.StateDB, blockNumber uint64) []sdk.SDKContract {
	return []sdk.SDKContract{
		stateReceiverContract{},
	}
}

func TestStateReceiver(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db), nil)
	statedb.CreateAccount(from)
	statedb.CreateAccount(to)
	evm := newEVM(statedb, vm.Config{}, vm.TxContext{}, new(contractsApp))
	testCommit(t, evm)
	testExecutedId(t, evm)
	testStateSyncId(t, evm)
	testRootByStateSyncId(t, evm)
	testCommitmentByStateSyncId(t, evm)
	testExecute(t, evm)
}

func testCommit(t *testing.T, evm *vm.EVM) {
	for i := 1; i <= 10; i++ {
		events = append(events, &StateSync{
			Id:       big.NewInt(int64(i)),
			Sender:   common.BigToAddress(big.NewInt(int64(i))),
			Receiver: common.BigToAddress(big.NewInt(int64(i))),
			Data:     []byte{byte(i)},
		})
	}
	var trieNodes [][]byte

	leaves := make(map[*big.Int]common.Hash)
	for i := 0; i < 10; i++ {
		raw, _ := rlp.EncodeToBytes(events[i])
		hash := crypto.Keccak256Hash(raw)
		leaves[big.NewInt(int64(i))] = hash
		trieNodes = append(trieNodes, hash.Bytes())
	}
	var err error
	tree, err = merkle.NewMerkleTree(trieNodes)
	require.Nil(t, err)

	commitment := StateSyncCommitment{StartId: big.NewInt(1), EndId: big.NewInt(10), Root: tree.Hash()}
	cr, _ := rlp.EncodeToBytes(commitment)
	extraTree, _ := merkle.NewMerkleTree([][]byte{crypto.Keccak256(cr)})
	proof, err := extraTree.GenerateProof(crypto.Keccak256(cr))
	require.Nil(t, err)
	input, err := Abi.Pack("commit", commitment, uint64(0), proof, QuorumCert{
		Epoch:        0,
		ViewNumber:   0,
		BlockHash:    [32]byte{},
		BlockNumber:  0,
		BlockIndex:   0,
		ExtendHash:   extraTree.Hash(),
		Signature:    nil,
		ValidatorSet: BitArray{},
	})
	require.Nil(t, err)
	_, _, err = evm.Call(vm.AccountRef(from), to, input, 1000, big.NewInt(0))
	require.Nil(t, err)
}

func testExecute(t *testing.T, evm *vm.EVM) {
	for _, event := range events {
		raw, _ := rlp.EncodeToBytes(event)
		hash := crypto.Keccak256Hash(raw)
		proof, err := tree.GenerateProof(hash.Bytes())
		require.Nil(t, err)
		input, err := Abi.Pack("execute", proof, event)
		require.Nil(t, err)
		_, _, err = evm.Call(vm.AccountRef(from), to, input, 1000, big.NewInt(0))
		require.Nil(t, err)

		input, err = Abi.Pack("getExecutedId")
		require.Nil(t, err)
		ret, _, err := evm.Call(vm.AccountRef(from), to, input, 1000, big.NewInt(0))
		res, err := Abi.Methods["getExecutedId"].Outputs.UnpackValues(ret)
		require.True(t, res[0].(*big.Int).Cmp(event.Id) == 0)
		require.Nil(t, err)
	}
}

func testExecutedId(t *testing.T, evm *vm.EVM) {
	input, err := Abi.Pack("getExecutedId")
	require.Nil(t, err)
	ret, _, err := evm.Call(vm.AccountRef(from), to, input, 1000, big.NewInt(0))
	res, err := Abi.Methods["getExecutedId"].Outputs.UnpackValues(ret)
	require.True(t, res[0].(*big.Int).Cmp(big.NewInt(0)) == 0)
	require.Nil(t, err)

}
func testStateSyncId(t *testing.T, evm *vm.EVM) {
	input, err := Abi.Pack("getStateSyncId")
	require.Nil(t, err)
	ret, _, err := evm.Call(vm.AccountRef(from), to, input, 1000, big.NewInt(0))
	res, err := Abi.Methods["getStateSyncId"].Outputs.UnpackValues(ret)
	require.True(t, res[0].(*big.Int).Cmp(big.NewInt(10)) == 0)
	require.Nil(t, err)

}
func testRootByStateSyncId(t *testing.T, evm *vm.EVM) {
	input, err := Abi.Pack("getRootByStateSyncId", big.NewInt(5))
	require.Nil(t, err)
	ret, _, err := evm.Call(vm.AccountRef(from), to, input, 1000, big.NewInt(0))
	res, err := Abi.Methods["getRootByStateSyncId"].Outputs.UnpackValues(ret)
	require.Nil(t, err)
	require.Len(t, res, 1)
}
func testCommitmentByStateSyncId(t *testing.T, evm *vm.EVM) {
	input, err := Abi.Pack("getCommitmentByStateSyncId", big.NewInt(5))
	require.Nil(t, err)
	ret, _, err := evm.Call(vm.AccountRef(from), to, input, 1000, big.NewInt(0))
	//var cm StateSyncCommitment
	res, err := Abi.Methods["getCommitmentByStateSyncId"].Outputs.UnpackValues(ret)

	require.Nil(t, err)
	require.Len(t, res, 1)

}

func newEVM(statedb *state.StateDB, vmconfig vm.Config, txContext vm.TxContext, app sdk.ContractsApp) *vm.EVM {
	initialCall := true
	canTransfer := func(db vm.StateDB, address common.Address, amount *big.Int) bool {
		if initialCall {
			initialCall = false
			return true
		}
		return core.CanTransfer(db, address, amount)
	}
	transfer := func(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {}

	context := vm.BlockContext{
		CanTransfer: canTransfer,
		Transfer:    transfer,
		GetHash:     vmTestBlockHash,
		BlockNumber: new(big.Int).SetUint64(0),
		Time:        new(big.Int).SetUint64(1),
		GasLimit:    1000,
	}
	return vm.NewEVM(context, txContext, statedb, params.MainnetChainConfig, vmconfig, app)
}
func vmTestBlockHash(n uint64) common.Hash {
	return common.BytesToHash(crypto.Keccak256([]byte(big.NewInt(int64(n)).String())))
}
