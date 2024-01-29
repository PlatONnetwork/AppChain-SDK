# Context

AppChain-SDK 为模块开发人员提供的接口的参数都携带了`context`，根据接口的使用场景，`context`又分为`Context`、`ConsensusContext`和`WorkerContext`。

## Context

`Context`是其他context的基础。

```go title="AppChain-Base/sdk/context.go"
type Context interface {
	Context() context.Context
	Backend() Backend
}
```

方法说明：

- `Context()` 返回context.Context对象。
- `Backend()` 返回实现了Backend接口的对象，用于与底层数据打交道。

## ConsensusContext

`ConsensusContext`用于共识扩展和验证人选举接口，代表当前共识状态。

```go title="AppChain-Base/sdk/context.go"
type ConsensusContext interface {
	Context

	Epoch() uint64
	View() uint64
	BlockIndex() uint32
	Header() *types.Header
	StateDB() StateDBReader
	ParentStateDB() StateDBReader
	IsConsensusNode() bool
	IsProposer() bool
	NumberValidators(epoch uint64) int
}
```

方法说明：

- `Epoch()` 返回当前共识轮的epoch。
- `View()` 返回当前共识轮的view number。
- `BlockIndex()` 返回当前共识区块的序号（表示验证人打包的第几个区块）。
- `Header()` 返回当前共识的区块头。
- `ParentStateDB()` 返回当前共识区块的父区块的链状态。
- `IsConsensusNode()` 返回当前节点是否是共识节点。
- `IsProposer()` 返回当前节点是否是区块提议人。
- `NumberValidators(epoch uint64)` 返回`epoch`对应共识轮的验证人数量。

## WorkerConext

`WorkerContext`用于worker接口，代表当前正在打包或执行的区块的上下文信息。

```go title="AppChain-Base/sdk/context.go"
type WorkerContext interface {
	Context

	StateDB() StateDB
	Header() *types.Header
	IsWorker() bool
	ParentBlock() *types.Block
}
```

方法说明：

- `StateDB()` 返回当前区块的链状态。
- `Header()` 返回当前区块的区块头。
- `IsWorker()` 返回当前调用放是否是`worker`。
- `ParentBlock()` 返回当前区块的父区块。

## Backend

```go title="AppChain-Base/sdk/backend.go"
type Backend interface {
	ChainId() (*big.Int, error)
	GetAddressHrp() string

	GetPoolNonce(addr common.Address) (uint64, error)

	CurrentHeader() *types.Header
	GetHeader(hash common.Hash, number uint64) *types.Header
	GetHeaderByHash(hash common.Hash) *types.Header
	GetHeaderByNumber(number uint64) *types.Header

	GetBlock(hash common.Hash, number uint64) *types.Block
	GetBlockByHash(hash common.Hash) *types.Block
	GetBlockByNumber(number uint64) *types.Block

	GetReceiptsByHash(hash common.Hash) types.Receipts
	ReadReceipts(sealHash common.Hash) types.Receipts

	ContractCode(hash common.Hash) ([]byte, error)

	State() (StateDBReader, error)
	StateAt(root common.Hash) (StateDBReader, error)

	GetEVMCommit(msg Message, blockHash common.Hash) (*vm.EVM, func() error, error)
	GetEVM(msg Message, header *types.Header) (*vm.EVM, func() error, error)

	GetTransaction(txHash common.Hash) (*types.Transaction, common.Hash, uint64, uint64, error)
}
```

方法说明：

- `ChainId()` 返回链ID。
- `GetAddressHrp()` 返回HRP地址前缀。
- `GetPoolNonce(addr common.Address)` 获取指定地址的nonce。
- `CurrentHeader()` 返回最新上链的区块头。
- `GetHeader(hash common.Hash, number uint64)` 从链上获取指定的区块头。
- `GetHeaderByHash(hash common.Hash)` 通过区块hash从链上获取区块头。
- `GetHeaderByNumber(number uint64)` 通过块高从链上获取区块头。
- `GetBlock(hash common.Hash, number uint64)` 从缓存获取指定的区块。
- `GetBlockByHash(hash common.Hash)` 通过区块hash从链上获取区块。
- `GetBlockByNumber(number uint64)` 通过块高从链上获取区块。
- `GetReceiptsByHash(hash common.Hash)` 通过交易hash从链上获取回执。
- `ReadReceipts(sealHash common.Hash)` 通过区块seal hash从缓存获取交易回执列表。
- `ContractCode(hash common.Hash)` 通过合约代码hash从链上获取合约代码。
- `State()` 获取当前链状态。
- `StateAt(root common.Hash)` 从链上获取指定root的状态。
- `GetEVMCommit(msg Message, blockHash common.Hash)` 从链上获取blockHash对应链状态的EVM虚拟机。
- `GetEVM(msg Message, blockHash common.Hash)` 从缓存获取blockHash对应链状态的EVM虚拟机。
