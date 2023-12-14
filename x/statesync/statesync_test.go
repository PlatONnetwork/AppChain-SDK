package statesync

import (
	"context"
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync/sync"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	types2 "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/utils"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/PlatONnetwork/PlatON-Go/trie"
	"github.com/stretchr/testify/require"
	"github.com/test-go/testify/assert"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"testing"
)

var (
	privateKey, _ = crypto.GenerateKey()
)

func TestFlow(t *testing.T) {
	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db), nil)
	assert.NotNil(t, statedb)
	storedb := memorydb.New()
	extradb := extravote.NewExtraVoteDB(storedb)
	simSync, err := sync.MakeL1Sync(constants.StateSyncAddress, sync.NewSimClient(), big.NewInt(0), storedb, nil)
	require.Nil(t, err)
	stateSync, err := newStateSync(cli.NewContext(cli.NewApp(), nil, nil), storedb, extradb)
	require.Nil(t, err)
	stateSync.l1Sync = simSync.L1Sync
	ctx := &ConsensusContext{
		backend: &testBackend{statedb: statedb},
	}
	simSync.Scan(context.Background(), 1, 10)
	extraData := stateSync.ExtendData(ctx)
	assert.NotNil(t, extraData)

	err = stateSync.VerifyExtendData(ctx, extraData)
	assert.Nil(t, err)

	stateSync.PrepareQC(ctx, &protocols.PrepareBlock{
		Epoch:         0,
		ViewNumber:    0,
		Block:         types.NewBlock(&types.Header{}, nil, nil, new(trie.Trie)),
		BlockIndex:    0,
		ProposalIndex: 0,
		PrepareQC:     nil,
		ViewChangeQC:  nil,
		Signature:     types2.Signature{},
		ExtendData:    nil,
	}, nil)
	leaves, tree := GenTree(extraData)
	stateSync.extraDb.InsertProof(0, 0, 0, leaves, tree)
	wctx := &WorkContext{
		backend: ctx.backend,
		header:  &types.Header{},
	}
	local, remote := stateSync.AddTxs(wctx, make(map[common.Address]types.Transactions), make(map[common.Address]types.Transactions))
	assert.Equal(t, 1, len(local))
	assert.Equal(t, 0, len(remote))
}

func GenTree(extraData []byte) ([][]byte, *merkle.MerkleTree) {
	leaves := [][]byte{extraData}
	tree, _ := merkle.NewMerkleTree(leaves)
	return leaves, tree
}

func newStateSync(ctx *cli.Context, store store.Store, extraDb *extravote.ExtraVoteDB) (*StateSync, error) {
	eventProofDb := NewEventProofDB(store)
	return &StateSync{
		logger:       log.New("module", "statesync"),
		startBlock:   big.NewInt(0),
		privateKey:   privateKey,
		store:        store,
		eventProofDb: eventProofDb,
		extraDb:      extraDb,
		p2p:          NewSyncP2P(),
		syncUpdateCh: make(chan struct{}),
	}, nil
}

type WorkContext struct {
	backend sdk.Backend
	header  *types.Header
}

func (w WorkContext) Context() context.Context {
	//TODO implement me
	panic("implement me")
}

func (w WorkContext) Backend() sdk.Backend {
	return w.backend
}

func (w WorkContext) StateDB() sdk.StateDB {
	//TODO implement me
	panic("implement me")
}

func (w WorkContext) Header() *types.Header {
	return w.header
}

func (w WorkContext) IsWorker() bool { return false }

type ConsensusContext struct {
	EpochNumber      uint64
	ViewNumber       uint64
	BlockIndexNumber uint32
	backend          sdk.Backend
}

func (c ConsensusContext) Context() context.Context {
	//TODO implement me
	panic("implement me")
}

func (c ConsensusContext) Backend() sdk.Backend {
	return c.backend
}

func (c ConsensusContext) Epoch() uint64 {
	return c.EpochNumber
}

func (c ConsensusContext) View() uint64 {
	return c.ViewNumber
}

func (c ConsensusContext) BlockIndex() uint32 {
	return c.BlockIndexNumber
}

func (c ConsensusContext) Header() *types.Header {
	return &types.Header{}
}

func (c ConsensusContext) StateDB() sdk.StateDBReader {
	//TODO implement me
	panic("implement me")
}

func (c ConsensusContext) ParentStateDB() sdk.StateDBReader {
	//TODO implement me
	panic("implement me")
}

func (c ConsensusContext) IsConsensusNode() bool {
	//TODO implement me
	panic("implement me")
}

func (c ConsensusContext) IsProposer() bool {
	//TODO implement me
	panic("implement me")
}

func (c ConsensusContext) NumberValidators(epoch uint64) int {
	//TODO implement me
	panic("implement me")
}

type testBackend struct {
	statedb *state.StateDB
}

func (t testBackend) ChainId() (*big.Int, error) {
	return big.NewInt(101), nil
}

func (t testBackend) GetAddressHrp() string {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetPoolNonce(addr common.Address) (uint64, error) {
	return 0, nil
}

func (t testBackend) CurrentHeader() *types.Header {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetHeader(hash common.Hash, number uint64) *types.Header {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetHeaderByHash(hash common.Hash) *types.Header {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetHeaderByNumber(number uint64) *types.Header {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetBlock(hash common.Hash, number uint64) *types.Block {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetBlockByHash(hash common.Hash) *types.Block {
	block := types.NewBlock(&types.Header{}, nil, nil, new(trie.Trie))
	data, _ := types2.EncodeExtra(1, &types2.QuorumCert{
		Epoch:        0,
		ViewNumber:   0,
		BlockHash:    common.Hash{},
		BlockNumber:  0,
		BlockIndex:   0,
		ExtendHash:   common.Hash{},
		Signature:    types2.Signature{},
		ValidatorSet: utils.NewBitArray(1),
	})
	block.SetExtraData(data)
	return block
}

func (t testBackend) GetBlockByNumber(number uint64) *types.Block {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetReceiptsByHash(hash common.Hash) types.Receipts {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) ReadReceipts(sealHash common.Hash) types.Receipts {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) ContractCode(hash common.Hash) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) State() (sdk.StateDBReader, error) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) StateAt(root common.Hash) (sdk.StateDBReader, error) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetEVMCommit(msg sdk.Message, blockHash common.Hash) (*vm.EVM, func() error, error) {
	//TODO implement me
	panic("implement me")
}

func (t testBackend) GetEVM(msg sdk.Message, header *types.Header) (*vm.EVM, func() error, error) {
	return newEVM(t.statedb, vm.Config{}, vm.TxContext{}, nil), nil, nil
}

func (t testBackend) GetTransaction(txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error) {
	//TODO implement me
	panic("implement me")
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
