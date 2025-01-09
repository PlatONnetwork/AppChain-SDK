# 交易黑名单

交易黑名单模块展示如何用 AppChain-SDK 对客户端交易进行过滤。

## 构建

* 实现 Module 接口


```go
const ModuleVersion uint64 = 0
const ModuleName = "blacklist"

type Module struct {
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
```

* 实现黑名单存储

黑名单地址通过 RPC 的方式进行调用，存在竞态调用


```go
type Module struct {
    sync.Mutex
    signer    types.Signer
    blacklist map[common.Address]struct{}
}
func (m *Module) Set(addrs []common.Address) {
    m.Lock()
    defer m.Unlock()
    for _, addr := range addrs {
        m.blacklist[addr] = struct{}{}
    }
}
func (m *Module) Get() []common.Address {
    m.Lock()
    defer m.Unlock()
    var list []common.Address
    for addr := range m.blacklist {
        list = append(list, addr)
    }
    return list
}
```

* 实现启动初始化 InitModule，解析交易签名获得 From 地址，所以需要通过 Backend 获取到 chainId来构建 signer

```go
func (m *Module) Init(ctx sdk.InitContext) error {
	chainId, err := ctx.Backend().ChainId()
	if err != nil {
		return err
	}
	m.signer = types.NewLondonSigner(chainId)
	m.blacklist = make(map[common.Address]struct{})
	return nil
}
```

* 实现 TxPoolModule，在 CheckTx 函数中获取交易 From 地址，再根据 blacklist 判断是否是合法地址。FilterPendingTxs 在此用例中没有特殊处理，直接返回入参 txs.

```go
func (m *Module) CheckTx(ctx sdk.Context, tx *types.Transaction) error {
	addr := tx.FromAddr(m.signer)
	if _, ok := m.blacklist[addr]; ok {
		return errors.New("blacklist address")
	}
	return nil
}
func (m *Module) FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	return txs
}
```

* 实现 RpcModule 

```go
func (m *Module) APIs() []rpc.API {
	return []rpc.API{
		rpc.API{
			Namespace: "blacklist",
			Version:   "1",
			Service:   NewRPC(m),
			Public:    false,
		},
	}
}

```
```go
type Blacklist interface {
	Set([]common.Address)
	Get() []common.Address
}

type RPC struct {
	p Blacklist
}

func NewRPC(blacklist Blacklist) *RPC {
	return &RPC{p: blacklist}
}

func (r *RPC) SetBlacklist(address []common.Address) error {
	r.p.Set(address)
	return nil
}

func (r *RPC) GetBlacklist() ([]common.Address, error) {
	return r.p.Get(), nil
}
```

* 实现 RPC 客户端，增加 blacklist_setBlacklist、blacklist_getBlacklist

```go
func NewClient(rawurl string) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: ethclient.NewClient(c),
		rpc:    c,
	}, nil
}

func (c *Client) SetBlacklist(ctx context.Context, address []common.Address) error {
	return c.rpc.CallContext(ctx, nil, "blacklist_setBlacklist", address)
}

func (c *Client) GetBlacklist(ctx context.Context) ([]common.Address, error) {
	var address []common.Address
	err := c.rpc.CallContext(ctx, &address, "blacklist_getBlacklist")
	return address, err
}

```


## 测试程序

* Server

server 端使用 testutil 来在内存中建立单节点测试网络
```go
func Server(ctx *cli.Context) error {
	blacklistModule := blacklist.Module{}
    
	vals, _ := testutil.NewValidator(testutil.DefaultAccount[0:1])

	manager := module.NewManager(vals, &blacklistModule)
	//设置选举模块
	manager.SetElection(vals.Name())
	//设置交易模块
	manager.SetOrderTxPool(blacklistModule.Name())
	app := testutil.NewApp(manager)

	var stack []*node.Node
	var backend []*eth.Ethereum
	var err error
	if stack, backend, err = testutil.CreateCluster(
		//选取用例账户
        testutil.DefaultAccount[0:1],
        []sdk.App{app},
        map[string]json.RawMessage{},
		//设置预置的测试账户
        common2.UserAddrs); err != nil {
        return err
	}
    //将namespace为 blacklist 的 rpc 加入到 rpcserver 中
	stack[0].Config().HTTPModules = append(stack[0].Config().HTTPModules, "blacklist")

	stack[0].Start()
	backend[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}
```

* Client 

在server启动预置了测试账户，测试客户端，1. 设置黑名单地址，2. 账户尝试发交易，判断是否能发送成功


```go
func Client(ctx *cli.Context) error {

	cli, err := blacklist.NewClient("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	
	addrs := decodeAddr(ctx)
	fmt.Println("addrs:", len(addrs))

	//设置黑名单
	err = cli.SetBlacklist(context.Background(), addrs)
	if err != nil {
		return err
	}
	addrs, err = cli.GetBlacklist(context.Background())
	if err != nil {
		return err
	}

	var nonces []uint64
	for _, addr := range common2.UserAddrs {
		nonce, err := cli.NonceAt(context.Background(), addr, nil)
		if err != nil {
			return err
		}
		nonces = append(nonces, nonce)
	}
	chainId, err := cli.ChainID(context.Background())
	if err != nil {
		return err
	}
	//预置地址尝试发送转账交易
	for i, addr := range common2.UserAddrs {
		tx := types.NewTransaction(nonces[i], common.BigToAddress(big.NewInt(1111)), big.NewInt(1000), 100000, big.NewInt(11000000000), nil)
		signer := types.NewLondonSigner(chainId)
		signature, err := crypto.Sign(signer.Hash(tx, nil).Bytes(), common2.UserPrivateKeys[addr])
		if err != nil {
			return err
		}
		tx, _ = tx.WithSignature(signer, signature)
		fmt.Println("tx:", tx.Hash().String())
		err = cli.SendTransaction(context.Background(), tx)
		if err != nil {
			fmt.Println("send failed, from", addr.Hex(), "err:", err.Error())
		}
		time.Sleep(time.Second * 2)
	}

	return nil
}
```

### 启动

```shell
make blacklist
```