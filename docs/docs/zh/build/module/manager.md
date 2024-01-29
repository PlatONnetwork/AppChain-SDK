# Module Manger

## 应用模块接口


应用模块接口的存在是为了便于将模块组合在一起，形成一个功能完整的AppChain-SDK应用程序。

### Module

`Module`定义了每个模块都需要实现的方法。

```go title="types/module/module.go"
type Module interface {
	Name() string
}
```

- `Name() :` 以字符串形式返回模块的名称。

### InitModule

`InitModule`是一个扩展接口，用于模块初始化。

```go title="types/module/module.go"
type InitModule interface {
	Init(ctx sdk.InitContext) error
}
```

### ContractModule

`ContractModule`定义了合约模块需要实现的方法。

```go title="types/module/module.go"
type ContractModule interface {
	Module
	Address() common.Address
	Run(evm *EVM, contract *Contract, input []byte, readOnly bool) ([]byte, error)
}
```

- `Address() :` 返回合约地址。
- `Run(evm *EVM, contract *Contract, input []byte, readOnly bool) :` 调用合约入口函数。

### TxPoolModule

`TxPoolModule`是一个交易扩展接口，用于验证交易和过滤待打包的交易。

```go title="types/module/module.go"
type TxPoolModule interface {
	Module
	CheckTx(ctx sdk.Context, tx *types.Transaction) error
	FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions
}
```

- `CheckTx(ctx sdk.Context, tx *types.Transaction) :` 验证入交易池的交易，若验证失败返回`false`，否则返回`true`。
- `FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) :` 过滤待打包的交易，返回过滤后的交易。

### RpcModule

`RpcModule`RPC服务扩展接口，用于定义模块的RPC接口。

```go title="types/module/module.go"
type RpcModule interface {
	Module
	APIs() []rpc.API
}
```

- `APIs() :` 返回模块实现的RPC接口列表。

### P2PModule

`P2PModule` P2P服务扩展接口，用于模块自定义P2P协议。

```go title="types/module/module.go"
type P2PModule interface {
	Module
	Protocols() []p2p.Protocol
}
```

- `Protocols() :` 返回模块自定义的P2P协议。

### ConsensusExtendModule

`ConsensusExtendModule` 是一个共识引擎扩展接口，用于模块利用共识引擎的能力完成特定业务逻辑的场景。

```go title="types/module/module.go"
type ConsensusExtendModule interface {
	Module
	ExtendData(ctx sdk.ConsensusContext) []byte
	VerifyExtendData(ctx sdk.ConsensusContext, data []byte) (common.Hash, error)
	PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote)
}
```

- `ExtendData(ctx sdk.ConsensusContext) :` 返回模块需要共识的扩展数据。
- `VerifyExtendData(ctx sdk.ConsensusContext, data []byte) :` 验证模块的扩展数据，并返回扩展数据的Hash。
- `PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) :` 模块扩展数据共识结果，模块可通过该实现该方法记录扩展数据对应的区块信息以及投票结果。

### BlockCommitter

`BlockCommitter` 是一个共识扩展接口，用于通知模块区块`committing`。

```go title="types/module/module.go"
type BlockCommitter interface {
	Module
	OnCommit(ctx sdk.ConsensusContext, block *types.Block) error
}
```

- `OnCommit(ctx sdk.ConsensusContext, block *types.Block) :` 此方法为模块开发人员提供了实现区块提交时自动触发的逻辑。

### ElectionModule

`ElectionModule` 是一个共识扩展接口，为共识引擎提供验证人列表以及共识轮次信息。AppChain-SDK必需有一个实现该接口的模块（有且只有一个），目前系统内置模块`staking`模块实现了该接口。大多数情况下，应用使用内置的`staking`模块即可，无需额外开发。

```go title="types/module/module.go"
type ElectionModule interface {
	Module
	NewHeader(ctx sdk.ConsensusContext, header *types.Header) error
	GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64
	GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error)
	IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool
}
```

- `NewHeader(ctx sdk.ConsensusContext, header *types.Header) :` 打包区块时，共识引擎会调用该方法。模块可通过实现该方法，自定义区块头相关信息。
- `GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) :` 返回`blockNumber`对应的共识轮的最后一个区块块高。
- `GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) :` 返回`blockNumber`对应的共识轮的验证人列表。
- `IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) :` 返回`nodeID`对应的节点是否是候选验证人。

### GenesisModule

`GenesisModule`是一个允许模块实现创世功能的扩展接口。

```go title="types/module/module.go"
type GenesisModule interface {
	Module
	InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error
}
```

### BeginBlockerModule

```go title="types/module/module.go"
type BeginBlockerModule interface {
	Module
	BeginBlock(ctx sdk.WorkerContext) error
}
```

- `BeginBlock(ctx sdk.WorkerContext) :` 此方法为模块开发人员提供了实现在每个区块开始时（打包/执行）自动触发的逻辑。

### EndBlockerModule


```go title="types/module/module.go"
type EndBlockerModule interface {
	Module
	EndBlock(ctx sdk.WorkerContext) error
}
```

- `EndBlock(ctx sdk.WorkerContext) :` 此方法为模块开发人员提供了实现每个区块介绍时（打包/执行）自动触发的逻辑。

### WorkerModule

```go title="types/module/module.go"
type WorkerModule interface {
	Module
	SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error)
}
```

- `SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) :` 此方法用于在区块打包时允许模块排序待执行的交易。

### TransactionModule

```go title="types/module/module.go"
type TransactionModule interface {
	Module
	AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error)
}
```

- `AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transcations) :` 此方法用于在区块打包过程中，插入模块的交易。

## Manager

`Manager`是一个保存应用所有`Module`的结构，并定义了这些模块的执行顺序。

```go title="types/module/module.go"
type Manager struct {
	Modules map[string]interface{}

	ConsensusExtend    string
	Election           string
	Worker             string
	OrderInit          []string
	OrderTxPool        []string
	OrderBlockCommiter []string
	OrderGenesis       []string
	OrderBeginBlocker  []string
	OrderEndBlocker    []string
	OrderBlocker       []string
	OrderTransaction   []string
}
```

模块管理主要对模块集合进行集体操作，它实现了以下方法：

- `NewModule(moudles ...Module) *Manager :` 构造函数。传入应用的`Module`集合去构造新的`Manager`。它通常用于应用的主构造函数中调用。

- `SetOrderInit(moduleNames ...string) :` 设置应用启动时调用每个模块的`InitModule`接口的顺序。

- `SetOrderTxPool(moduleNames ...string) :` 设置每个模块的`TxPoolModule`接口的调用顺序。

- `SetConsensusExtend(moduleName string) :` 设置`ConsensusExtendModule`接口的实现模块。每个应用仅允许一个模块实现`ConsensusExtendModule`接口。

- `SetElection(moduleName string) :` 设置`ElectionModule`接口的实现模块。每个应用仅允许一个模块实现`ElectionModule`接口。

- `SetOrderBlockCommiter(moduleNames ...string) :` 设置在区块提交时调用每个模块的`BlockCommiterModule`接口的顺序。

- `SetOrderGenesis(moduleNames ...string) :` 设置应用初始化创世块时调用每个模块的`GenesisModule`接口的顺序。

- `SetOrderBeginBlock(moduleNames ...string) :` 设置在区块开始时调用每个模块的`BeginBlock()`方法的顺序。

- `SetOrderEndBlock(moduleNames ...string) :` 设置在区块结束时调用每个模块的`EndBlock()`方法的顺序。

- `SetOrderTransaction(moduleNames ...string) :` 设置在区块打包时调用每个模块的`AddTxs()`方法的顺序。

- `InitChain(ctx sdk.Context) :` 在应用启动时，此方法经由`BaseApp`调用，在此方法中按`OrderInit`指定的顺序调用模块的`Init()`方法。

- `Contracts() :` 在应用启动时，此方法经由`BaseApp`调用，返回合约模块列表。

- `APIs() :` 在应用启动时，此方法经由`BaseApp`调用，返回模块自定义的RPC接口列表。

- `Protocols() :` 在应用启动时，此方法经由`BaseApp`调用，返回模块自定义的P2P协议。

- `CheckTx(ctx sdk.Context, tx *types.Transaction) :` 在交易进入交易池校验时，此方法经由`BaseApp`调用，在此方法中按照`OrderTxPool`指定的顺序调用`TxPoolModule`接口的`CheckTx()`方法。

- `FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) :` 在获取`penging`交易列表后，此方法经由`BaseApp`调用，在此方法中按照`OrderTxPool`指定的顺序调用`TxPoolModule`接口的`FilterPendingTxs()`方法。

- `ExtendData(ctx sdk.ConsensusContext) :` 共识引擎提议`PrepareBlock`时，此方法经由`BaseApp`调用，在此方法中调用`ConsensusExtend`指定模块的`ExtendData()`方法。

- `VerifyExtendData(ctx sdk.ConsensusContext, data []byte) :` 共识引擎执行`PrepareBlock`之后，此方法经由`BaseApp`调用，在此方法中调用`ConsensusExtend`指定模块的`VerifyExtendData()`方法。

- `PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocol.PrepareVote) :` 当`PrepareBlock`收集到足够的投票后达成QC后，此方法经由`BaseApp`调用，在此方法中调用`ConsensusExtend`指定模块的`PrepareQC()`方法。

- `NewHeader(ctx sdk.ConsensusContext, header *types.Header) :` 在打包区块时，此方法经由`BaseApp`调用，在此方法中调用`Election`指定模块的`NewHeader()`方法。

- `GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) :` 该方法经由`BaseApp`调用，然后在此方法中调用`Election`指定模块的`GetLastNumber()`方法。

- `GetValiator(ctx sdk.ConsensusContext, blockNumber uint64) :` 该方法经由`BaseApp`调用，然后在此方法中调用`Election`指定模块的`GetValidator()`方法。返回`blockNumber`对应共识轮的验证人列表。

- `IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) :` 该方法经由`BaseApp`调用，然后在此方法中调用`Election`指定模块的`IsCandidateNode()`方法。返回`nodeID`指定的节点是否是候选验证人。

- `OnCommit(ctx sdk.ConsensusContext, block *types.Block) :` 在区块提交时，此方法经由`BaseApp`调用，然后在此方法中按照`OrderBlockCommiter`指定的顺序调用模块的`OnCommit()`方法。

- `InitGenesis(ctx sdk.ConsensusContext, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) :` 应用初始化时，此方法经由`BaseApp`调用。在此方法中，按照`OrderGenesis`的顺序调用模块的`InitGenesis()`方法。

- `BeginBlock(ctx sdk.WorkerContext) :` 在区块开始时（执行），此方法经由`BaseApp`调用，然后按照`OrderBeginBlocker`的顺序调用模块的`BeginBlock()`方法。

- `EndBlock(ctx sdk.WorkerContext) :` 在区块结束时（执行），此方法经由`BaseApp`调用，然后按照`OrderEndBlocker`的顺序调用模块的`EndBlock()`方法。

- `AddTxs(ctx sdk.WorkerContext) :` 在区块打包时，此方法经由`BaseApp`调用，然后按照`OrderTranscation`的顺序调用模块的`AddTxs()`方法。返回交易列表。

- `SortTxs(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) :` 在区块打包时交易执行前，经由`BaseApp`调用，然后调用每个实现了`WorkerModule`接口的模块。返回排序后的交易列表。

以下是`simapp`集成模块的例子：

```go title="simapp/app.go"
    datadir := node.DefaultDataDir()
	if ctx.GlobalIsSet(utils.DataDirFlag.Name) {
		datadir = ctx.GlobalString(utils.DataDirFlag.Name)
	}
	if datadir != "" {
		absdatadir, err := filepath.Abs(datadir)
		if err != nil {
			return nil, err
		}
		datadir = absdatadir
	}

	dbfile := filepath.Join(datadir, "sdk")
	store, err := storage.NewStorage(dbfile, 256, 512, "sdk")
	if err != nil {
		log.Error("failed to new storage", "err", err)
		return nil, err
	}

	l1Module := l1.NewModule(store)

	stateEvent := stateevent.NewModule(store)

	stageModule := stage.NewModule(ctx)
	vrfModule := vrf.NewModule(ctx, stageModule)
	stakeModule := staking.NewModule(ctx, l1Module, stageModule)
	stateSync, err := statesync.NewModule(ctx, l1Module, stakeModule, store, extravote.NewExtraVoteDB(store))
	if err != nil {
		return nil, err
	}
	rewardModule := reward.NewModule(ctx, stageModule)
	depositModule := deposit.NewModule(ctx, l1Module)
	l2StateSender := statesender.NewModule(ctx)

	vrfModule.SetStakeModule(stakeModule)
	stakeModule.SetRewardModule(rewardModule)
	stakeModule.SetVRFModule(vrfModule)
	rewardModule.SetStakeModule(stakeModule)

	rootchainRpc := ctx.GlobalString(x.RootchainNodeRPCFlag.Name)
	rootchainTxRelayer := txrelayer.NewModule(rootchainRpc, txrelayer.DefaultReceiptTimeout, txrelayer.DefaultNumRetries)

	manager := module.NewManager(
        stateSync,
        stateEvent,
        l1Module,
        extraVote,
        rootchainTxRelayer,
        checkpoint,
        stageModule,
        vrfModule,
        stakeModule,
        rewardModule,
        depositModule,
        l2StateSender)
	manager.SetElection(stakeModule.Name())
	manager.SetConsensusExtend(extraVote.Name()
	manager.SetOrderTransaction(stateSync.Name(), vrfModule.Name(), stakeModule.Name())
	manager.SetOrderInit(stateSync.Name(), rootchainTxRelayer.Name(), checkpoint.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderGenesis(l1Module.Name(), stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name(), depositModule.Name(), l2StateSender.Name(), stateSync.Name())
	manager.SetOrderBeginBlocker(stageModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderEndBlocker(stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderBlockCommiter(stakeModule.Name(), stateEvent.Name(), checkpoint.Name())

    app := &SimApp{}
	baseApp, err := baseapp.NewBaseApp("simapp", store, manager)
	if err != nil {
		return nil, err
	}
	app.BaseApp = baseApp
```
