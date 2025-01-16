# 升级

升级用例展示如何使用升级模块对现有系统进行升级，用例中展示了两种类型升级，Golang合约模块升级，功能性模块进行升级


## 构建

* Oracle 模块升级

在 oracle 用例中，在创世区块就已经构建了oracle 模块。在本用例中将采用升级的方式，在指定块高初始化 oracle 模块。

模块的升级是通过实现 RegistryModule 接口进行实现，
```go
type Module struct {
	*oracle.Module
	extraVote *extravote.ExtraVote
	logger    log.Logger
}

func NewModule(store store.Store, key *ecdsa.PrivateKey, rateClient oracle.RateMarketClient) *Module {
	return &Module{
		Module: oracle.NewModule(store, key, rateClient),
		logger: log.New("module", oracle.ModuleName),
	}
}

func (m *Module) SetExtraVote(extraVote *extravote.ExtraVote) {
	m.extraVote = extraVote
}

func (m *Module) RegistryUpgradeHandler(registrar module.UpgradeRegistrar) error {
	registrar.RegisterUpgradeHandler(oracle.ModuleName, 0, func(ctx sdk.WorkerContext) error {
        //根据升级块高构建创世文件
		og, _ := json.Marshal(oracle.GenesisConfig{Decimals: 3, BlockNumber: ctx.Header().Number.Uint64()})
        //调用oracle 模块进行模块的创世初始化
		if err := m.InitGenesis(ctx, ctx.StateDB(), ctx.Backend().ChainConfig(), og); err != nil {
			return err
		}
        //在extravote 模块中进行设置
		m.extraVote.AddEnableVerifiers(oracle.ModuleName)
		m.logger.Info("Run upgrade handler success")
		return nil
	})
	return nil
}
```

* Node 合约的升级

Node 合约分为三个版本，   

1. 仅有addNode 函数
```solidity
interface Node{
    function addNode(string memory name, string memory host, uint16 port) external;
}
```
2. addNode 函数增加事件
```solidity
interface NodeV1{
    event AddNode(string name, string host, uint16 port);
    function addNode(string memory name, string memory host, uint16 port) external;
}
```
3. 增加 getNode 查询，新增NodeInfo结构体
```solidity
struct NodeInfo{
    string name;
    string host;
    uint16 port;
}
interface NodeV2{
    function addNode(string memory name, string memory host, uint16 port) external;
    function getNode(string memory name) external view returns(NodeInfo memory);
}
```
4.增加 delNode 删除函数
```solidity
struct NodeInfo{
    string name;
    string host;
    uint16 port;
}
interface NodeV3{
    event AddNode(string name, string host, uint16 port);
    event DelNode(string name);
    function addNode(string memory name, string memory host, uint16 port) external;
    function getNode(string memory name) external view returns(NodeInfo memory);
    function delNode(string memory name) external;
}
```

生成框架代码

```shell
make upgrade
```

* 实现 Node 框架代码

node_impl.go
```go
type Info struct {
	Host string
	port uint16
}
//创建存储结构
type Storage struct {
	Nodes *container.Map[*Info]
}
func NewNode(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Node, error) {
	s := &Node{
		abi:           nil,
		abis:          make(map[uint64]*abi.ABI),
		methodEntry:   make(map[string]func([]byte) ([]byte, error)),
		methodEntries: make(map[uint64]map[string]func([]byte) ([]byte, error)),
		evm:           evm,
		contract:      contract,
		burner:        contracts.NewBurner(contract),
		stateDb:       contracts.NewStateDB(evm, contract),
		context:       contracts.NewContext(evm, contract),
		readOnly:      readOnly,
	}
	s.storage = &Storage{
		Nodes: container.NewMap[*Info](nodeKey, contract.Address(), s.stateDb),
	}
    //为了便于启动，这里将三个版本ABI全部加载
	s.initABI()
	s.initABIV1()
	s.initABIV2()
	s.initABIV3()
	s.initMethodEntry()
	s.initMethodV1Entry()
	s.initMethodV2Entry()
	s.initMethodV3Entry()
	s.loadMethodABI()
	return s, nil
}
func (c *Node) AddNode(name string, host string, port uint16) error {
	contracts.Require(len(host) != 0, "Node: invalid host")
	contracts.Require(port != 0, "Node: invalid port")
	c.storage.Nodes.MustSet(name, &Info{
		Host: host,
		port: port,
	})
	return nil
}
```
node_v1.go

```go
func (c *Node) AddNodeV1(name string, host string, port uint16) error {
	contracts.Require(len(host) != 0, "Node: invalid host")
	contracts.Require(port != 0, "Node: invalid port")
	c.AddNode(name, host, port)
	if c.GetVersion() > 0 {
		c.EmitAddNodeEvent(name, host, port)
	}
	return nil
}
```

node_v2.go
```go
func (c *Node) GetNode(name string) (NodeInfo, error) {
	var nodeInfo NodeInfo
	if node := c.storage.Nodes.MustGet(name); node.port != 0 {
		nodeInfo = NodeInfo{
			Name: name,
			Host: node.Host,
			Port: node.port,
		}
	}
	return nodeInfo, nil
}
```

node_v3.go

```go
func (c *Node) DelNode(name string) error {
	c.storage.Nodes.MustSet(name, nil)
	c.EmitDelNodeEvent(name)
	return nil
}
```

* 创建 node 模块

```go
const ModuleVersion uint64 = 0
const ModuleName = "node"

var (
	NodeAddress = common.BigToAddress(big.NewInt(2223))
)

type Module struct {
	logger log.Logger
}

func NewModule() *Module {
	return &Module{
		logger: log.New("module", ModuleName),
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	return nil
}
```

* ContractModule 接口

```go
func (m *Module) Address() common.Address {
	return NodeAddress
}

func (m *Module) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	upgrade, _ := contracts.NewNode(evm, contract, readOnly)
	return upgrade.Run(input)
}

func (m *Module) ContractCreateBlockNumber(db sdk.StateDBReader) uint64 {
	upgrade, _ := contracts.NewNode(sdkcontracts.NewEVM(types.NewStateDBWrapper(db), big.NewInt(0)), sdkcontracts.NewContract(m, m), true)
	return upgrade.GetCreateBlock()
}
```

* 实现 RegistryModule 接口

```go
func (m *Module) RegistryUpgradeHandler(registrar module.UpgradeRegistrar) error {
    //注册4个版本
	registrar.RegisterUpgradeHandler(ModuleName, 0, func(ctx sdk.WorkerContext) error {

		c, _ := contracts.NewNode(sdkcontracts.NewEVM(ctx.StateDB(), ctx.Header().Number), sdkcontracts.NewContract(m, m), false)
		c.InitGenesis(ctx.Header().Number.Uint64())
		m.logger.Info("Run upgrade handler success")
		return nil
	})
	registrar.RegisterUpgradeHandler(ModuleName, 1, func(ctx sdk.WorkerContext) error {
		m.logger.Info("Update contract version", "version", 1)
		c, _ := contracts.NewNode(sdkcontracts.NewEVM(ctx.StateDB(), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
		c.SetVersion(1)
		return nil
	})
	registrar.RegisterUpgradeHandler(ModuleName, 2, func(ctx sdk.WorkerContext) error {
		m.logger.Info("Update contract version", "version", 2)
		c, _ := contracts.NewNode(sdkcontracts.NewEVM(ctx.StateDB(), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
		c.SetVersion(2)
		return nil
	})
	registrar.RegisterUpgradeHandler(ModuleName, 3, func(ctx sdk.WorkerContext) error {
		m.logger.Info("Update contract version", "version", ModuleVersion)
		c, _ := contracts.NewNode(sdkcontracts.NewEVM(ctx.StateDB(), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
		c.SetVersion(3)
		return nil
	})
	return nil
}

```

## 构建测试程序

* Server
```go
func Server(ctx *cli.Context) error {
	store := memorydb.New()
    //创建升级版本的oracle模块
	oracleModule := oracle2.NewModule(store, testutil.DefaultAccount[0].NodePrivateKey(), oracle.NewTimerRate(time.Second*2))
    //创建extravote
	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{oracleModule})
    //添加oracle
	oracleModule.SetExtraVote(extraVote)
    //创建选举模块
	vals := election.NewModule()
    //创建升级模块
	upgradeModule := upgrade.NewModule(store)
    //创建node模块
	nodeModule := examplenode.NewModule()
	manager := module.NewManager(vals, extraVote, upgradeModule, oracleModule, nodeModule)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(vals.Name(), extraVote.Name(), upgradeModule.Name())
	manager.SetConsensusExtend(extraVote.Name())
    //设置升级模块注册
	manager.RegisterUpgradeHandler(upgradeModule)
    //设置模块检查
	manager.SetModuleValidChecker(upgradeModule.IsModuleValid)
    //设置模块是否是合约的检查函数
	upgradeModule.SetIsContractModule(manager.IsContractModule)
    //oracle注册升级 handler
	oracleModule.RegistryUpgradeHandler(upgradeModule)
    //node 模块注册升级handler
	nodeModule.RegistryUpgradeHandler(upgradeModule)
	app := testutil.NewApp(manager)

	config := election.GenesisConfig{
		InitialNodes: election.Nodes{
			election.Node{
				Name:        "node1",
				Owner:       testutil.DefaultAccount[0].NodeAddress(),
				Desc:        "node1",
				PublicKey:   crypto.FromECDSAPub(&testutil.DefaultAccount[0].NodePrivateKey().PublicKey),
				BlsPubKey:   testutil.DefaultAccount[0].BlsSecretKey().GetPublicKey().Serialize(),
				HostAddress: "127.0.0.1",
				RpcPort:     uint16(testutil.DefaultAccount[0].HTTP),
				P2pPort:     uint16(testutil.DefaultAccount[0].P2PPort),
			},
		},
		AdminAddress:     common2.UserAddrs[0],
		EpochSize:        50,
		ElectionDistance: 20,
	}
	s, _ := json.Marshal(config)
	upgradeConfig, _ := json.Marshal(&types.GenesisConfig{
		ModuleGenesisConfig: module.ModuleGenesisConfig{},
		Owner:               common2.UserAddrs[9],
	})
	var stack []*node.Node
	var backend []*eth.Ethereum
	var err error
	if stack, backend, err = testutil.CreateCluster(testutil.DefaultAccount[0:1], testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		vals.Name():          s,
		upgradeModule.Name(): upgradeConfig,
	}, common2.UserAddrs); err != nil {
		return err
	}

	stack[0].Start()
	backend[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}

```

* Client

oracle 升级

```go
func upgradeOracle(ctx *cli.Context) error {
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	height := ctx.Uint64(upgradeBlockFlag.Name)
	addr := common2.UserAddrs[9]
	key := common2.UserPrivateKeys[addr]
	nonce, err := cli.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}
	up, err := NewUpgrade(constants.UpgradeAddress, cli)
	if err != nil {
		return err
	}
    //构建交易
	tx, err := up.AddUpgradePlan(createOpt(addr, key, nonce), IUpgradePlan{
		Name: "oracle-upgrade",
		Modules: []IUpgradeModule{
			{
				ModuleName: "oracle",
				Version:    0,
			},
		},
		Info:   "init genesis oracle module",
		Height: height,
		Status: 0,
	})
	if err != nil {
		return err
	}
	if err = waitTx(cli, tx); err != nil {
		return err
	}
    //向升级合约查询升级计划
	queryPlan(up, addr, height)
	return nil
}
```

* 查询汇率

```go
func queryRate(ctx *cli.Context) error {
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	lc, err := contracts.NewRate(oracle.RateAddr, cli)

	for i := 0; i < 10; i++ {
		blockNumber, err := lc.BlockNumber(nil)
		if err != nil {
			return err
		}
		rate, err := lc.Rate(nil)
		if err != nil {
			return err
		}
		fmt.Println("blockNumber:", blockNumber, "rate:", rate)
		time.Sleep(5 * time.Second)
	}

	return nil
}
```

* node 合约升级

```go
func upgradeNode(ctx *cli.Context) error {
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
    //升级计划块高
	height := ctx.Uint64(upgradeBlockFlag.Name)
    //升级版本
	version := ctx.Uint64(upgradeVersionFlag.Name)
	addr := common2.UserAddrs[9]
	key := common2.UserPrivateKeys[addr]
	nonce, err := cli.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}
	up, err := NewUpgrade(constants.UpgradeAddress, cli)
	if err != nil {
		return err
	}
	tx, err := up.AddUpgradePlan(createOpt(addr, key, nonce), IUpgradePlan{
		Name: "node-upgrade-" + randomString(6),
		Modules: []IUpgradeModule{
			{
                //设置模块名字及版本
				ModuleName: "node",
				Version:    version,
			},
		},
		Info:   fmt.Sprintf("upgrade node module, version:%d", version),
		Height: height,
		Status: 0,
	})
	if err != nil {
		return err
	}
	if err = waitTx(cli, tx); err != nil {
		return err
	}
	queryPlan(up, addr, height)
	return nil
}

```

* 检查 node 版本的升级情况

```go
func checkNode(ctx *cli.Context) (err error) {
	addr := common2.UserAddrs[9]
	key := common2.UserPrivateKeys[addr]
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	nonce, err := cli.NonceAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	node, err := NewNode(examplenode.NodeAddress, cli)
	if err != nil {
		return err
	}
	name := ctx.String(nameFlag.Name)
	if len(name) == 0 {
		name = randomString(6)
	}

	//检查 addNode 及事件
	fmt.Println("check addNode method")
	tx, err := node.AddNode(createOpt(addr, key, nonce), name, ctx.String(hostFlag.Name), uint16(ctx.Uint64(portFlag.Name)))
	if err != nil {
		return err
	}

	err = waitTx(cli, tx)
	if err != nil {
		return err
	}
    receipt, err := cli.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil
	}
	fmt.Println("add node success", "name:", name，"event", len(receipt.Logs) != 0)


	//检查 getNode
	fmt.Println("check getNode method")
	info, err := node.GetNode(&bind.CallOpts{
		From: addr,
	}, name)
	if err != nil {
		return err
	}
	fmt.Println("get node:%s", name, "info:", info)

	//检查 delNode
	fmt.Println("check delNode method")
	tx, err = node.DelNode(createOpt(addr, key, nonce+1), name)
	if err != nil {
		return err
	}

	err = waitTx(cli, tx)
	if err != nil {
		return err
	}
	info, err = node.GetNode(&bind.CallOpts{
		From: addr,
	}, name)
	if err != nil {
		return err
	}
	fmt.Println("get node:%s", name, "info:", info)
	return nil
}

```

* 编译

```shell
make upgrade
```

* 启动 server
```shell
./upgrade server
```

* oracle 升级

```shell
./upgrade client oracle-upgrade --upgrade-block 100
try call tx success, hash: 0xdbc40270905d0664e3d69d6b30c790df548e6bea7dc378ba7de7ba220a1eb522
set plan success name: oracle-upgrade module: oracle version: 0 info: init genesis oracle module height: 100

```

* oracle 汇率查询

没有到升级区块
```shell
./upgrade client query-rate
no contract code at given address
```

```shell
./upgrade client query-rate
blockNumber: 0 rate: 0
blockNumber: 107 rate: 2
blockNumber: 115 rate: 4
blockNumber: 123 rate: 6

```

* node 节点创世升级

```shell
./upgrade client node-upgrade --upgrade-block 50 --upgrade-version 0
try call tx success, hash: 0x352237d21dc95ad0417b294d8f92339ee6ab0f19b2fb54ce8ad928f2c4f7469f
set plan success name: node-upgrade-pewQDa module: node version: 0 info: upgrade node module, version:0 height: 50

```
* node合约查询测试
```shell
./upgrade client check-node
check addNode method
try call tx success, hash: 0xe7ef250e15272ce17be39b2f29f7e9ef7d7a32078b3ac1dca86a13e34b8eb6c3
add node success name: hsPplq event false
check getNode method
methods not found

```
* node 节点V1升级

```shell
./upgrade client node-upgrade --upgrade-block 100 --upgrade-version 1
```
* node合约查询测试
```shell
./upgrade client check-node
check addNode method
try call tx success, hash: 0xc0fb1b755fc452eca379ef6cac3a4c457900be286a2dc6bfc37939dfd99a4825
add node success name: QtaOgC event true
check getNode method
methods not found

```
* node 节点V2升级

```shell
./upgrade client node-upgrade --upgrade-block 150 --upgrade-version 2
```
* node合约查询测试
```shell
./upgrade client check-node
check addNode method
try call tx success, hash: 0x7ca6109c9b4d4b0b04f27da061172a88622363b7a3622722a05a2579935c93d3
add node success name: GCuPpb event true
check getNode method
get node:%s GCuPpb info: {  0}
check delNode method
methods not found

```
* node 节点V3升级

```shell
./upgrade client node-upgrade --upgrade-block 200 --upgrade-version 3
try call tx success, hash: 0x7a3fdd54660af541e450023b90d4bad4b989dc75d47778403cdaef1c54241d7f
set plan success name: node-upgrade-SXKGUK module: node version: 3 info: upgrade node module, version:3 height: 200

```

* node合约查询测试
```shell
./upgrade client check-node
check addNode method
try call tx success, hash: 0x79fe4f7164a421b95e77702758ab502ffdcae85ea07607071889d3404aa441f3
add node success name: nroNqh event true
check getNode method
get node:%s nroNqh info: {  0}
check delNode method
try call tx success, hash: 0xc070dfc2f9466027fa6325a9d599c151bed5d7168c15571761a0413ccfd70c4a
get node:%s nroNqh info: {  0}

```