# 应用初始化

基于AppChain-SDK开发的应用是一个类以太坊应用。对于以太坊来说，启动一个节点一般先初始化创世块，才能运行节点。同样的，使用AppChain-SDK 开发的应用也需要通过这两个步骤启动。

针对这两个步骤，AppChain-SDK提供了:

- **GenesisModule接口** 为模块提供了创世功能。
- **InitModule接口** 初始化模块。
- **ContractModule接口** 注册合约模块到底层虚拟机中。
- **RpcModule接口** 注册模块定义的RPC接口到底层RPC服务中。
- **P2PModule接口** 注册模块定义的P2P协议到底层P2P服务中。

其中`RpcModule`和`P2PModule`由其他篇幅单独描述。

## GenesisModule

模块通过实现`GenesisModule`接口来完成以下功能：

- 读取创世配置
- 读取创世配置中模块自定义的配置
- 通过创世配置初始化模块：

    * 初始化模块数据到StateDB
    * 初始化模块数据到本地数据库中（sdk数据库，不上链）

看两个例子：

- 初始化模块数据到StateDB

```go title="x/staking/module.go"
func (s *StakeModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	configParams := config.DefualtStakeNetworkParams()
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON StakeNetworkParams bytes", "error", err)
		return err
	}

	var conf config.StakeNetworkParams
	if err := json.Unmarshal(raw, &conf); nil != err {
		log.Error("Failed UnmarshalJSON StakeNetworkParams", "error", err)
		return err
	} else {
		configParams = &conf
	}

	// init staking handler account nonce
	initAccountNonce(db, s.Address())
	// store configParms
	initStakeConfigParams(db, s.Address(), configParams)

	if err := initValidators(db, s.Address(), chainConfig, configParams); nil != err {
		log.Error("Failed initialize genesis validators", "error", err)
		return err
	}

	log.Info("Succeed init genesis", "module", s.Name(), "StakeNetworkParams", configParams.String())
	return nil
}
```

- 初始化模块数据到本地数据库

```go title="x/l1/module.go"
func (l *L1Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	var g L1ConfigParams
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON l1 L1ConfigParams bytes", "error", err)
		return err
	}

	if err := json.Unmarshal(raw, &g); nil != err {
		log.Error("Failed UnmarshalJSON l1 L1ConfigParams", "error", err)
		return err
	}

	l.db.setChainID(g.ChainID)
	l.db.setStateAddress(g.State)
	l.db.setCheckpointAddress(g.Checkpoint)
	l.db.setStakeManagerAddress(g.StakeManager)
	l.db.setDepositManagerAddress(g.DepositManager)

	log.Info("Succeed init genesis", "module", l.Name(), "chainId", g.ChainID, "state", g.State.Hex(), "checkpoint", g.Checkpoint.Hex(), "stakeManager", g.StakeManager.Hex(), "depositManager", g.DepositManager.Hex())
	return nil
}
```

## InitModule

为了支持底层`account` `init` `attach`等子命令，`InitModule`接口将模块构造和模块启动逻辑分开。`InitModule`接口一般用于：

- 读取模块自定义的命令行参数
- 初始化模块运行所需的参数

```go title="x/checkpoint/module.go"
func (m *Module) Init(ctx sdk.InitContext) error {
	if m.keystoreFile == "" {
		return fmt.Errorf("checkpoint.keystore not set")
	}
	if m.passwordFile == "" {
		return fmt.Errorf("checkpoint.password not set")
	}

	key, err := sdkcom.DecryptKey(m.keystoreFile, m.passwordFile)
	if err != nil {
		return err
	}
	m.key = key
	return nil
}
```

## ContractModule

实现了`ContractModule`接口的模块称为合约模块。应用启动时，Module Manager通过断言模块是否实现`ContractModule`接口来将合约模块注册到底层EVM虚拟机中。以便后续执行合约交易能够找到对应的合约模块。

```go title="x/statesync/module.go"
func (s *StateSync) Address() common.Address {
	return constants.StateSyncAddress
}

func (s *StateSync) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	stateReceiver, _ := contracts.NewStateReceiver(evm, contract, readOnly)
	return stateReceiver.Run(input)
}
```
