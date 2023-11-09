package module

import (
	"encoding/json"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type Module interface {
	Name() string
}

type ContractModule interface {
	Module
	ContractAddress() common.Address
	Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error)
}

type TxPoolModule interface {
	Module
	CheckTx(ctx sdk.Context, tx *types.Transaction) error
	FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions
}

type RpcModule interface {
	Module
	Apis() []rpc.API
}

type P2PModule interface {
	Module
	Protocols() []p2p.Protocol
}

type ConsensusExtendModule interface {
	Module
	ExtendData(ctx sdk.Context) []byte
	VerifyExtendData(ctx sdk.Context, data []byte) (common.Hash, error)
	PrepareQC(ctx sdk.Context, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote)
}

type BlockCommiter interface {
	OnCommit(ctx sdk.Context, block *types.Block) error
}

type ElectionModule interface {
	Module
	BlockCommiter
	NewHeader(ctx sdk.Context, header *types.Header) error
	GetLastNumber(ctx sdk.Context, blockNumber uint64) uint64
	GetValidator(ctx sdk.Context, blockNumber uint64) (*cbfttypes.Validators, error)
	IsCandidateNode(ctx sdk.Context, nodeID enode.ID) bool
}

type GenesisModule interface {
	Module
	InitGenesis(ctx sdk.Context, db sdk.StateDB, genesis *core.Genesis, data json.RawMessage)
}

type BlockerModule interface {
	Module
	BeginBlock(ctx sdk.Context)
	EndBlock(ctx sdk.Context)
}

type WorkerModule interface {
	Module
	SortTxs(ctx sdk.Context, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error)
}

type TransactionModule interface {
	Module
	AddTxs(ctx sdk.Context, local, remote map[common.Address]types.Transactions) (map[common.Address]types.Transactions, map[common.Address]types.Transactions)
}
