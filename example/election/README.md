# 验证人模块

验证人模块，利用 Solidity 合约实现 POA 共识合约。用例涉及到 AppChain-SDK 扩展功能有：
* GenesisModule
* EndBlockerModule
* ElectionModule
* Solidity 内置合约

## 构建

* Solidity  合约负责管理验证节点的准入以及验证节点的切换，合约分为两个合约，一个合约是对外即外部用户可以调用，这部分功能包括添加节点、更新节点、删除节点等。另一个合约则是模块内部调用，负责更新每个Epoch周期节点。

这两个合约的存储布局是一致的，保证数据的更新的一致性。

定义验证节点信息

```solidity
struct Node  {
	string name;                 
	address owner;   
	string desc;                 
	bytes publicKey; 
	bytes blsPubKey;  
	string hostAddress;    
	uint16 rpcPort;
	uint16 p2pPort;
}
```

节点类型及审核状态

```solidity
enum NodeType {
    ConsensusNode,
    ObserverNode
}

enum Status {
    UnauditedNode,
    ValidNode,
    InvalidNode,
    DisabledNode,
    DeletedNode
}

enum AuditState{
    AuditPassed,
    AuditRejected
}
```

共识周期中保存节点列表，起始结束块高

```solidity
struct RoundNodeList{
    string[] nodes;
    uint256 epoch;
    uint256 start;
    uint256 end;
}
```
合约创建 

```solidity
contract Election is Initializable, OwnableUpgradeable {

    event Apply(string  name, string  desc, bytes  publicKey, bytes  blsPubKey, string  hostAddress, uint16 rpcPort, uint16 p2pPort);
    event Audit(string name, AuditState auditStat, string auditReason);
    event Add(string  name, address owner, string desc, bytes publicKey, bytes blsPubKey, string hostAddress, uint16 rpcPort, uint16 p2pPort);
    event Update(string  name, string  hostAddress,uint16 rpcPort ,uint16 p2pPort, string desc );
    event Delete(string name);
    event ChangeEpoch(uint256 epoch, uint256 start, uint256 end);

    mapping(string => NodeInfo) nodeInfoList; //节点信息
    RoundNodeList lastRoundValidator; //上一轮节点列表
    RoundNodeList currentRoundValidator; //当前节点列表
    RoundNodeList nextRoundValidator; //下一轮节点列表
    string[] public nodeList; //节点名字，用于遍历
    uint64 public epochSize; //每个周期区块个数
    uint64 public distance; //距离周期结束多少区块选举下一轮
    uint64 public epoch; //epoch 周期值
    bool updateEpoch; //产生更新 epoch 事件

    function initialize(Node[] memory _nodes, uint64 _epochSize, uint64 _distance) external initializer {
        epochSize = _epochSize;
        distance = _distance;
        require(epochSize > distance, "Election: epoch size too small");
        for(uint i = 0; i < _nodes.length; i++) {
            _add(_nodes[i]);
            nodeInfoList[_nodes[i].name].root = true;
            nodeInfoList[_nodes[i].name].status = Status.ValidNode;
            nodeInfoList[_nodes[i].name].nodeType = NodeType.ConsensusNode;
        }

        lastRoundValidator = RoundNodeList({
            nodes:nodeList,
            epoch: 0,
            start: 0,
            end:0
        });
        
        currentRoundValidator = RoundNodeList({
            nodes:nodeList,
            epoch: 0,
            start: 1,
            end:epochSize
        });
    }

    function changeValidator() external {
        if (updateEpoch){
            emit ChangeEpoch(currentRoundValidator.epoch, currentRoundValidator.start, currentRoundValidator.end);
            updateEpoch = false;
        }
    }



    function audit(string calldata name, AuditState auditStat, string calldata auditReason) external onlyOwner {  
        NodeInfo memory node = nodeInfoList[name];
        if (auditStat == AuditState.AuditPassed) {
            node.status = Status.ValidNode;
        } else if (auditStat == AuditState.AuditRejected){
            node.status = Status.InvalidNode;
        }
        nodeInfoList[node.name] = node;
        node.updateTime = uint64(block.timestamp);
        emit Audit(name, auditStat, auditReason);
    }

    function addNode(Node memory node) external onlyOwner {
        _add(node);
        emit Add(node.name, node.owner, node.desc, node.publicKey, node.blsPubKey, node.hostAddress, node.rpcPort, node.p2pPort);
    }

    function _add(Node memory node) internal {
        nodeList.push(node.name);
        nodeInfoList[node.name] = NodeInfo({
            name:node.name,
            owner:node.owner,
            desc:node.desc,
            publicKey:node.publicKey,
            blsPubKey:node.blsPubKey,
            hostAddress:node.hostAddress,
            rpcPort:node.rpcPort,
            p2pPort:node.p2pPort,
            nodeType: NodeType.ObserverNode,
            status: Status.UnauditedNode,
            root:false,
            createTime:uint64(block.timestamp),
            updateTime:uint64(block.timestamp)
        });

    }
    
    function update(string calldata name, string calldata hostAddress,uint16 rpcPort ,uint16 p2pPort, string calldata desc ) external onlyOwner {
        NodeInfo memory node = nodeInfoList[name];
        require(bytes(node.name).length != 0, "Election: the node is not valid");
        node.hostAddress = hostAddress;
        node.rpcPort = rpcPort;
        node.p2pPort = p2pPort;
        node.desc = desc;
        emit Update(name, hostAddress, rpcPort, p2pPort, desc);
    }

    function deleteNode(string calldata name) external onlyOwner {
        NodeInfo memory node = nodeInfoList[name];
        require(bytes(node.name).length != 0, "Election: the node does not exists");
        require(node.status == Status.DeletedNode,  "Election: the node is not valid");
        require(node.root == false, "Election: the root node does not support delete");
        require(node.nodeType != NodeType.ConsensusNode, "Election: the consensus node cannot be deleted");
        require(hasNode(currentRoundValidator, name), "Election: the consensus node cannot be deleted");
        require(hasNode(nextRoundValidator, name), "Election: the consensus node cannot be deleted");
        node.status = Status.DeletedNode;
        node.updateTime = uint64(block.timestamp);
        nodeInfoList[name] = node;
        emit Delete(name);
    }

    function isCurrentRoundValidator(string calldata name) external view returns(bool) {
        return hasNode(currentRoundValidator, name);
    }

    function isNextRoundValidator(string calldata name) external view returns(bool) {
        return hasNode(nextRoundValidator, name);
    }

    function hasNode(RoundNodeList memory round, string memory name) internal pure returns(bool){
        for (uint i = 0; i < round.nodes.length; i++) {
            if (keccak256(bytes(round.nodes[i])) == keccak256(bytes(name))){
                return true;
            }
        }
        return false;
    }

    function disable(string calldata name) external onlyOwner {
        NodeInfo memory node = nodeInfoList[name];
        require(bytes(node.name).length != 0, "Election: the node does not exists" );
        require(node.status == Status.ValidNode, "Election: the node is not valid");
        require(node.root == false, "Election: the root node does not support delete");
        require(node.nodeType != NodeType.ConsensusNode, "Election: the consensus node cannot be deleted");
        require(hasNode(currentRoundValidator, name), "Election: the consensus node cannot be deleted");
        require(hasNode(nextRoundValidator, name), "Election: the consensus node cannot be deleted");
        node.status = Status.DisabledNode;
        node.updateTime = uint64(block.timestamp);
        nodeInfoList[name] = node;
    }
    function enable(string calldata name) external onlyOwner {
        NodeInfo memory node = nodeInfoList[name];
        require(bytes(node.name).length != 0, "Election: the node does not exists" );
        require(node.status == Status.DisabledNode, "Election: the node is not valid");
        node.status = Status.ValidNode;
        node.updateTime = uint64(block.timestamp);
        nodeInfoList[name] = node;
    }
    function updateType(string calldata name, NodeType nodeType) external onlyOwner {
        NodeInfo memory node = nodeInfoList[name];
        require(bytes(node.name).length != 0, "Election: the node does not exists" );
        require(node.status == Status.DisabledNode, "Election: the node is not valid");
        node.nodeType = nodeType;
        node.updateTime = uint64(block.timestamp);
        nodeInfoList[name] = node;
    }
    function getByName(string calldata name) external view returns(NodeInfo memory) {
       return nodeInfoList[name];
    }
    function getByPubKey(string calldata publicKey) external view returns(NodeInfo memory) {
        NodeInfo memory nodeInfo;
        for (uint i = 0; i <nodeList.length; i++){
            nodeInfo = nodeInfoList[nodeList[i]];
            if (keccak256(bytes(nodeInfo.publicKey)) == keccak256(bytes(publicKey))) {
                return nodeInfo;
            }
        }
        return nodeInfo;
    }
    function getLastRoundValidator() external view returns(RoundNodeList memory) {
        return lastRoundValidator;
    }
    function getCurrentRoundValidator() external view returns(RoundNodeList memory) {
        return currentRoundValidator;
    }

    function getNextRoundValidator() external view returns(RoundNodeList memory) {
        return nextRoundValidator;
    }

}
```

后端合约仅进行更新每个周期的验证人列表

```solidity
contract ElectionBackend is Initializable, OwnableUpgradeable {

    event Apply(string  name, string  desc, string  publicKey, string  blsPubKey, string  hostAddress, uint16 rpcPort, uint16 p2pPort);
    event Audit(string name, AuditState auditStat, string auditReason);
    event Add(string  name, address owner, string desc, string publicKey, string blsPubKey, string hostAddress, uint16 rpcPort, uint16 p2pPort);
    event Update(string  name, string  hostAddress,uint16 rpcPort ,uint16 p2pPort, string desc );
    event Delete(string name);
    event ChangeEpoch(uint256 epoch, uint256 start, uint256 end);

    mapping(string => NodeInfo) nodeInfoList;
    RoundNodeList lastRoundValidator;
    RoundNodeList currentRoundValidator;
    RoundNodeList nextRoundValidator;
    string[] public nodeList;
    uint64 public epochSize;
    uint64 public distance;
    uint64 public epoch;
    bool updateEpoch;

    function changeEpoch() external {
        uint256 epochOfBlocks = epochSize;
        bool endOfEpoch = (block.number % epochOfBlocks) == 0;
        if (endOfEpoch){
            epoch++;
            updateEpoch = true;
            lastRoundValidator = currentRoundValidator;
            currentRoundValidator = nextRoundValidator;
        }

        bool isElection = ((block.number + distance) % epochOfBlocks) == 0;

        if (isElection) {
            uint256 start = block.number + distance + 1;
		    uint256 end = block.number + distance + epochOfBlocks;
            _getValidNodes(nextRoundValidator);
            nextRoundValidator.epoch = epoch+1;
            nextRoundValidator.start = start;
            nextRoundValidator.end = end;
        }
    }

    function _getValidNodes(RoundNodeList storage round)internal{
        for (uint i = 0; i < nodeList.length; i++) {
            if (nodeInfoList[nodeList[i]].status == Status.ValidNode && nodeInfoList[nodeList[i]].nodeType == NodeType.ConsensusNode) {
                round.nodes.push(nodeList[i]);
            }
        }
    }
}
```

* 模块初始化

```go
const ModuleVersion uint64 = 0
const ModuleName = "election"

var ProxyAddress = common.BigToAddress(big.NewInt(101)) //代理合约地址
var ElectionAddress = common.BigToAddress(big.NewInt(102)) //选举合约实现地址
var CallerAddress = common.BigToAddress(big.NewInt(133)) //调用合约地址
type Module struct {
    chainConfig    *params.ChainConfig
    nodePrivateKey *ecdsa.PrivateKey
}

func NewModule() *Module {
    return &Module{}
}

func (m *Module) Name() string {
    return ModuleName
}

func (m *Module) Version() uint64 {
    return ModuleVersion
}
```

* 实现 GenesisModule 接口，创世配置需要指定初始节点，合约管理者，epoch 区块数量，提前多少区块进行选举

```go
type GenesisConfig struct {
	module.ModuleGenesisConfig
	InitialNodes     Nodes          `json:"initialNodes"`
	AdminAddress     common.Address `json:"adminAddress"`
	EpochSize        uint64         `json:"epochSize"`
	ElectionDistance uint64         `json:"electionDistance"`
}
func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	var genesis GenesisConfig
	if err := json.Unmarshal(data, &genesis); err != nil {
		return err
	}

    //构建代理合约
	proxyCaller, err := proxy.NewInitializableTransparentUpgradeableProxyGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
    //部署代理合约
	if err = proxyCaller.WithCaller(genesis.AdminAddress).WithTo(ProxyAddress).DeployInitializableTransparentUpgradeableProxy(); err != nil {
		return err
	}
    //创建选举合约
	electionCaller, err := contracts.NewElectionGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
    //部署选举合约
	if err = electionCaller.WithCaller(genesis.AdminAddress).WithTo(ElectionAddress).DeployElection(); err != nil {
		return err
	}
    //构建初始化参数
	input, _ := electionCaller.PackInitialize(genesis.InitialNodes.toContractNode(), genesis.EpochSize, genesis.ElectionDistance)
    //调用代理合约执行初始化
	if err = proxyCaller.WithCaller(genesis.AdminAddress).WithTo(ProxyAddress).Initialize(ElectionAddress, genesis.AdminAddress, input); err != nil {
		return err
	}
	return nil
}
```

* 实现 EndBlockerModule 接口，调用 ElectionBackend 的 changeEpoch 方法进行合约节点切换

```go
func (m *Module) EndBlock(ctx sdk.WorkerContext) error {
    //构建 ElectionBackend 
	caller, err := contracts.NewElectionBackendBackendCaller(ProxyAddress)
	if err != nil {
		return err
	}
    //执行切换
	return caller.WithCaller(CallerAddress).ChangeEpoch(ctx)
}
```

* 实现 ElectionModule 接口，在 NewHeader 中需要给区块头设置 coinbase，GetLastNumber 返回 epoch的最后一个区块的块高，GetValidator 返回当前验证节点信息，IsCandidateNode 用来判断指定NodeId是否是候选节点，由于我们是 POA 测试，没有候选节点，直接返回 false

```go
func (m *Module) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
    //判断是否是 Proposer，即自己生产的区块
	if ctx.IsProposer() {
		currentValidatorPubKey := hex.EncodeToString(crypto.FromECDSAPub(&m.nodePrivateKey.PublicKey))
		caller, _ := contracts.NewElectionGenesisCaller(ctx, types.NewStateDBWrapper(ctx.ParentStateDB()), m.chainConfig)
		nodeInfo, _ := caller.WithCaller(CallerAddress).WithTo(ProxyAddress).GetByPubKey(currentValidatorPubKey)
		if len(nodeInfo.Name) == 0 {
			return errors.New("not found validator")
		}
        //设置区块头 coinbase
		header.Coinbase = nodeInfo.Owner
	}
	return nil
}

func (m *Module) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
    //调用合约获取epoch size，计算当前 epoch 结束块高
	caller, _ := contracts.NewElectionGenesisCaller(ctx, types.NewStateDBWrapper(ctx.ParentStateDB()), m.chainConfig)
	epochOfBlocks, _ := caller.WithCaller(CallerAddress).WithTo(ProxyAddress).EpochSize()
    return (blockNumber + epochOfBlocks - 1) / epochOfBlocks * epochOfBlocks
}

func (m *Module) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	caller, _ := contracts.NewElectionGenesisCaller(ctx, types.NewStateDBWrapper(ctx.ParentStateDB()), m.chainConfig)
	caller = caller.WithCaller(CallerAddress).WithTo(ProxyAddress)
    //从合约中获取当前验证人
	rvn, _ := caller.GetCurrentRoundValidator()
    //大于当前块高，则获取下一轮验证人
	if rvn.End.Uint64() < blockNumber {
		rvn, _ = caller.GetNextRoundValidator()
	}
	vnm := make(cbfttypes.ValidateNodeMap, len(rvn.Nodes))
	for i, name := range rvn.Nodes {
		vd, _ := caller.GetByName(name)
		var blsKey bls.PublicKey
		pubKey, err := crypto.UnmarshalPubkey(vd.PublicKey)
		if err != nil {
			return nil, err
		}
		err = blsKey.Deserialize(vd.BlsPubKey)
		if err != nil {
			return nil, err
		}
		nodeId := enode.PubkeyToIDV4(pubKey)
		addr := crypto.PubkeyToNodeAddress(*pubKey)
		vn := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   addr,
			PubKey:    pubKey,
			BlsPubKey: &blsKey,
			NodeID:    nodeId,
		}
		vnm[nodeId] = vn
	}

	cvd := &cbfttypes.Validators{
		Nodes:            vnm,
		ValidBlockNumber: rvn.Start.Uint64(),
	}
	return cvd, nil
}

func (m *Module) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	return false
}

```


## 构建测试程序


* Server

```go
func Server(ctx *cli.Context) error {
	vals := election.NewModule()
	manager := module.NewManager(vals)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(vals.Name())
	app := testutil.NewApp(manager)
	nodeNumber := ctx.Int(nodeFlag.Name) - 1
	testutil.InitLog(log.LvlDebug)
    //初始Election创世节点配置
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
		AdminAddress:     common.HexToAddress(ctx.String(adminFlag.Name)),
		OwnerAddress:     common.HexToAddress(ctx.String(ownerFlag.Name)),
		EpochSize:        10,
		ElectionDistance: 5,
	}
	s, _ := json.Marshal(config)
	var stack []*node.Node
	var err error
    //根据节点编号，设置启动节点
	if stack, _, err = testutil.CreateCluster([]*testutil.Account{testutil.DefaultAccount[nodeNumber]}, testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		vals.Name(): s,
	}, common2.UserAddrs); err != nil {
		return err
	}

	stack[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}

```

* Client

```go
func Client(ctx *cli.Context) error {
	nodeNumber := ctx.Int(nodeFlag.Name) - 1
	if nodeNumber < 1 || nodeNumber > 2 {
		return errors.New("wrong node number")
	}
	adminKey, err := crypto.HexToECDSA(ctx.String(adminKeyFlag.Name))
	if err != nil {
		return err
	}
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	election, err := contracts.NewElection(election.ProxyAddress, cli)
	if err != nil {
		return err
	}
    
	opts := &bind.TransactOpts{
		From: crypto.PubkeyToAddress(adminKey.PublicKey),
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signer := types.NewLondonSigner(big.NewInt(123083))
			signature, err := crypto.Sign(signer.Hash(tx, nil).Bytes(), adminKey)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
	name := fmt.Sprintf("node%d", ctx.Int(nodeFlag.Name))
    //发送添加节点交易
	tx, err := election.AddNode(opts, contracts.Node{
		Name:        name,
		Owner:       testutil.DefaultAccount[nodeNumber].NodeAddress(),
		Desc:        fmt.Sprintf("node%d", nodeNumber),
		PublicKey:   crypto.FromECDSAPub(&testutil.DefaultAccount[nodeNumber].NodePrivateKey().PublicKey),
		BlsPubKey:   testutil.DefaultAccount[nodeNumber].BlsSecretKey().GetPublicKey().Serialize(),
		HostAddress: "127.0.0.1",
		RpcPort:     uint16(testutil.DefaultAccount[nodeNumber].HTTP),
		P2pPort:     uint16(testutil.DefaultAccount[nodeNumber].HTTP),
	})
	if err != nil {
		return err
	}
	fmt.Println("addNode tx:", tx.Hash().Hex())
	if err = common2.WaitTx(cli, tx); err != nil {
		return err
	}
    //发送审计通过交易
	tx, err = election.Audit(opts, name, 0, "pass")
	if err != nil {
		return err
	}
	fmt.Println("audit tx:", tx.Hash().Hex())
	if err = common2.WaitTx(cli, tx); err != nil {
		return err
	}
    //更新节点类型为共识节点
	tx, err = election.UpdateType(opts, name, 0)
	if err != nil {
		return err
	}
	fmt.Println("updateType tx:", tx.Hash().Hex())
	if err = common2.WaitTx(cli, tx); err != nil {
		return err
	}
	info, err := election.GetByName(nil, name)
	if err != nil {
		return err
	}
	fmt.Println("name:", name, "type", info.NodeType, "status:", info.Status)
	return nil
}

```

* 启动创世节点

```shell
./election server --node 1
```
查询节点共识状态
```shell
./election attach http://127.0.0.1:8801
> debug.consensusStatus().validator
true

```

* 节点 2 为共识节点

```shell
./election server --node 2
```

查询节点共识状态
```shell
./election attach http://127.0.0.1:8802
> debug.consensusStatus().validator
false

```

P2P 连接到创世节点

```shell
> admin.addPeer("enode://5c79bf8b836bdc85fe513a64a558291e96ac2405b6b95d5ca05bb20db9c0a1a00e12d11dd540d8685267015ce71cef43781a75157ec6f266cc88a8fb8b5c6c17@127.0.0.1:18001")
true
```

添加节点2 为共识节点

```shell
./election client --node 2
addNode tx: 0x7f167076d754373650c8a2ebbbe3013fd327c5e6cf75ec49fd2049eaf4e44692
audit tx: 0x66f0d1f5ec495f771dfc93288535665385b5556ca1b9b564656cb29f0daf578b
updateType tx: 0x8c836475cc24fd46568e505c3016509a2a930c354a63d95248a3912abb24c503
name: node2 type 0 status: 1
```
```shell
> debug.consensusStatus().validator
true
```

* 节点3 为共识节点

```shell
./election server --node 3
```

查询节点共识状态
```shell
./election attach http://127.0.0.1:8803
> debug.consensusStatus().validator
false

```

P2P 连接到创世节点

```shell
> admin.addPeer("enode://5c79bf8b836bdc85fe513a64a558291e96ac2405b6b95d5ca05bb20db9c0a1a00e12d11dd540d8685267015ce71cef43781a75157ec6f266cc88a8fb8b5c6c17@127.0.0.1:18001")
true
```

添加节点3 为共识节点

```shell
./election client --node 3
addNode tx: 0xfb81e2b1c7b64b36559a15c38a23c5c003b816b5fdaca196021741c4de247f87
audit tx: 0x7fb187c90453af83e4524d64c0bf5ecabaab5b21bfcbe798fc71793023501524
updateType tx: 0x2f1d20a849940cacf7c0716802d2fe1060ca8e6f835f022d5f37915c12c7eada
name: node3 type 0 status: 1

```
```shell
> debug.consensusStatus().validator
true
```
