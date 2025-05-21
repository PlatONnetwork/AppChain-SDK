package types

import (
	"encoding/json"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type InitChaner func(ctx sdk.InitContext) error

type Contractser func(statedb sdk.StateDBReader, blockNumber uint64) []sdk.SDKContract

type Starter func() error
type Stoper func() error

type CheckTxer func(ctx sdk.Context, tx *types.Transaction) error
type FilterPendingTxser func(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions

type APIser func() []rpc.API

type Protocolser func() []p2p.Protocol

type ExtendDataer func(ctx sdk.ConsensusContext) []byte
type VerifyExtendDataer func(ctx sdk.ConsensusContext, data []byte) (common.Hash, error)
type PrepareQCer func(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote)

type NewHeaderer func(ctx sdk.ConsensusContext, header *types.Header) error
type GetLastNumberer func(ctx sdk.ConsensusContext, blockNumber uint64) uint64
type GetValidatorer func(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error)
type IsCandidateNoder func(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool
type OnCommiter func(ctx sdk.ConsensusContext, block *types.Block) error

type InitGenesiser func(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) error

type BeginBlocker func(ctx sdk.WorkerContext) error
type EndBlocker func(ctx sdk.WorkerContext) error

type AddTxser func(ctx sdk.WorkerContext) (types.Transactions, error)
type SortTxser func(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) (types.Transactions, error)
type TxFiller func(sdk.WorkerContext, sdk.TxApplyCallbackApp) (types.Transactions, types.Receipts, error)

type TxExecutor func(sdk.WorkerContext, sdk.ContractsApp, types.Transactions) (types.Receipts, uint64, error)
