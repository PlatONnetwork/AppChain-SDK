# 优惠券

优惠券利用模块交易的自定义排序扩展，实现对指定账户获取到最大金额券。用例涉及到 AppChain-SDK 扩展功能有：
* GenesisModule
* WorkerModule
* RpcModule
* Solidity 内置合约


## 构建

* 实现优惠券合约，利用ERC20 合约，每个区块内向调用用户发送代币，根据调用顺序，每一位获得余额的 10%

```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
contract Coupon is ERC20 {
    using SafeMath for uint256;
    uint256 public blockNumber;
    uint256 public roundBalance;
    event Apply(address to, uint256 value, uint256 balance);
    constructor(string memory name, string memory symbol) ERC20(name, symbol){
    }
    function applyCoupon() external {
        if (blockNumber != block.number){
            roundBalance = 1000000000;
            blockNumber = block.number;
        }
        (bool ok, uint256 result) = roundBalance.tryDiv(10);
        require(ok, "Coupon: balance is insufficient");
        _mint(msg.sender, result);
        roundBalance = roundBalance - result;
        emit Apply(msg.sender, result, roundBalance);
    }
}
```

* 生成 genesis 框架

```shell
tools contract genesis --bin out/Coupon.sol/Coupon.bin --abi Coupon.sol/Coupon.abi.json --output . --type Coupon --pkg contracts
```

* 编译
```shell
make coupon
```
在 build/coupon/contracts 将会生成 coupon.go 、coupon.sol.go 两个文件

* 模块构建

```go
const ModuleVersion uint64 = 0
const ModuleName = "coupon"

type Module struct {
	sync.Mutex
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
```

* 初始化创世合约

合约需要定义地址

```go
var CouponAddress = common.BigToAddress(big.NewInt(102))
var CallerAddress = common.BigToAddress(big.NewInt(103))
```

coupon 合约的构造函数需要两个参数(name,symbol)，这需要通过创世文件的coupon 模块的创世参数传递，定义创世结构体

```go
type GenesisConfig struct {
	module.ModuleGenesisConfig
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
```
coupon 合约的初始化需要实现 GenesisModule

```go
func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	var genesis GenesisConfig
	if err := json.Unmarshal(data, &genesis); err != nil {
		return err
	}
	coupon, err := contracts2.NewCouponGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
	return coupon.WithCaller(CallerAddress).WithTo(CouponAddress).DeployCoupon(genesis.Name, genesis.Symbol)
}
```

* 交易排序

交易排序需要实现 WorkerModule 接口，通过设置 priority 来通过发送者地址判断交易的优先级

```go
type Module struct {
    sync.Mutex
    priority []common.Address
    cache    map[common.Address]int
}
func (m *Module) SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	target := make([]types.Transactions, len(m.priority), len(m.priority))
	other := make(map[common.Address]types.Transactions)

	findFunc := func(txs map[common.Address]types.Transactions) {
		for k, v := range txs {
			pos, ok := m.cache[k]
			if ok {
				target[pos] = v
			} else {
				other[k] = v
			}
		}
	}
	findFunc(local)
	findFunc(remote)

	var txs types.Transactions
	for _, ts := range target {
		txs = append(txs, ts...)
	}
	for _, v := range other {
		txs = append(txs, v...)
	}
	return txs, nil
}
func (m *Module) Set(priority []common.Address) {
    m.Lock()
    defer m.Unlock()
    m.cache = make(map[common.Address]int)
    m.priority = priority
    for i, p := range priority {
    m.cache[p] = i
}
}
func (m *Module) Get() []common.Address {
    m.Lock()
    defer m.Unlock()
    return m.priority
}
```

* 实现 RPC，RPC需要实现 RpcModule
```go
func (m *Module) APIs() []rpc.API {
	return []rpc.API{
		rpc.API{
			Namespace: "coupon",
			Version:   "1",
			Service:   NewRPC(m),
			Public:    false,
		},
	}
}

```
```go
type Priority interface {
	Set([]common.Address)
	Get() []common.Address
}

type RPC struct {
	p Priority
}

func NewRPC(priority Priority) *RPC {
	return &RPC{p: priority}
}

func (r *RPC) SetPriority(address []common.Address) error {
	r.p.Set(address)
	return nil
}

func (r *RPC) GetPriority() ([]common.Address, error) {
	return r.p.Get(), nil
}

```

实现  coupon_setPriority、coupon_getPriority rpc 客户端

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

func (c *Client) SetPriority(ctx context.Context, address []common.Address) error {
	return c.rpc.CallContext(ctx, nil, "coupon_setPriority", address)
}

func (c *Client) GetPriority(ctx context.Context) ([]common.Address, error) {
	var address []common.Address
	err := c.rpc.CallContext(ctx, &address, "coupon_getPriority")
	return address, err
}
```

## 测试程序

* Server

```go
func Server(ctx *cli.Context) error {
	couponModule := coupon.Module{}

	vals, _ := testutil.NewValidator(testutil.DefaultAccount[0:1])

	manager := module.NewManager(vals, &couponModule)
	manager.SetElection(vals.Name())
	//设置创世启动模块
	manager.SetOrderGenesis(couponModule.Name())
	//设置workModule模块
	manager.SetWorker(couponModule.Name())
	app := testutil.NewApp(manager)
	//构建coupon 模块创世参数
	config := coupon.GenesisConfig{
		Name:   "Coupon",
		Symbol: "CPN",
	}
	s, _ := json.Marshal(config)
	var stack []*node.Node
	var backend []*eth.Ethereum
	var err error
	if stack, backend, err = testutil.CreateCluster(testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		couponModule.Name(): s,
	}, common2.UserAddrs); err != nil {
		return err
	}
    //增加coupon namespace rpc
	stack[0].Config().HTTPModules = append(stack[0].Config().HTTPModules, "coupon")

	stack[0].Start()
	backend[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return nil
}
```

* Client

```go
func Client(ctx *cli.Context) error {

	cli, err := coupon.NewClient("http://127.0.0.1:8801")
	if err != nil {
		return err
	}
	//设置地址优先级
	addrs := decodeAddr(ctx)
	err = cli.SetPriority(context.Background(), addrs)
	if err != nil {
		return err
	}
	addrs, err = cli.GetPriority(context.Background())
	if err != nil {
		return err
	}
    for i, addr := range addrs {
        fmt.Printf("No.%d %s\n", i, addr.Hex())
    }

	coupon, err := contracts.NewContracts(coupon.CouponAddress, cli)
	var nonces []uint64
	for _, addr := range common2.UserAddrs {
		nonce, err := cli.NonceAt(context.Background(), addr, nil)
		if err != nil {
			return err
		}
		nonces = append(nonces, nonce)
	}
	//向 coupon 合约发送交易
	var txs types.Transactions
	for i, addr := range common2.UserAddrs {
		tx, err := coupon.ApplyCoupon(&bind.TransactOpts{
			From:  addr,
			Nonce: new(big.Int).SetUint64(nonces[i]),
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				signer := types.NewLondonSigner(big.NewInt(123083))
				signature, err := crypto.Sign(signer.Hash(tx, nil).Bytes(), common2.UserPrivateKeys[addr])
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, signature)
			},
			Value:     nil,
			GasPrice:  big.NewInt(11000000000),
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  0,
			Context:   nil,
			NoSend:    false,
		})
		if err != nil {
			return err
		}
		fmt.Println("tx:", tx.Hash().String())
		txs = append(txs, tx)
	}
	time.Sleep(time.Second * 5)
	//查询交易事件信息
	for _, tx := range txs {
		receipt, _ := cli.TransactionReceipt(context.Background(), tx.Hash())
		for _, log := range receipt.Logs {
			event, err := coupon.ParseApply(*log)
			if err == nil {
				fmt.Println("block:", receipt.BlockNumber, "index:", receipt.TransactionIndex, "address:", event.To.Hex(), "value:", event.Value, "balance:", event.Balance)
				break
			}
		}
	}
	return nil
}
```

## 执行

* 编译

```shell
make coupon
```
* 启动 Server

```shell
./coupon server
```

* 启动 Client

```shell
./coupon client --priority "0x1FD13cA28ccc423e0565fD3C9b64187ad91Ae7D6,0xd1391ba0Be0eECA4031e952402b92e26417D5F38"
```

参考设置 No.0、No.1 为交易发送者账户优先级，可以看到区块为6，交易索引为0，1的为设置的账户的发送者交易
```shell
No.0 0x1FD13cA28ccc423e0565fD3C9b64187ad91Ae7D6
No.1 0xd1391ba0Be0eECA4031e952402b92e26417D5F38
tx: 0x66b16a281b86c4387cbed4e044d23bf1d0dedfa46d128009a49db181a36194b4
tx: 0x1b4adfbc7fff649ad2a5cd4789e588250fbd8c3f7146bf5ac82c8139977c185b
tx: 0x37c8da5744af6d55950e056656956a714c522622b98b83d655f83c1657f9bf28
tx: 0x8b9124582d53e1808a77abf973829096d3a4902ae5afbe52c38b2334bbbb6ec3
tx: 0x200aa9cc34059433f5a0720ad54c2e4ca954f780a1734c60942d25bf4c1273d7
tx: 0x3e19503d5158c4eeedc8f177593a1731a606a7af23cdc39e20b21474f8d8b4b5
tx: 0x80f2d113972ee1d04c3765305e9ba8a59666c58ad5e13e27317da67ff333afe2
tx: 0x227475ddc75866ba6670b942075ef2bd4f5c210b0405e7a42f1835c1f37cb59f
tx: 0x707138440b8d2d0d0b50886a2cea5583522d959ad664d94d4c21b59059028126
tx: 0xe8fe9e36da263c67998e106d4aa79b2ba83cb25ced7fa9b1ff3987f882d7514d
block: 6 index: 6 address: 0x8a0f3F8389F79a05Dd03027dE0Bcbb06cADB3F9c value: 53144100 balance: 478296900
block: 6 index: 0 address: 0x1FD13cA28ccc423e0565fD3C9b64187ad91Ae7D6 value: 100000000 balance: 900000000
block: 6 index: 2 address: 0x27E1a2e05186c9e4C4A3A957F8748fA35EA19076 value: 81000000 balance: 729000000
block: 6 index: 7 address: 0x6151aE08d3acD4c8194bB029633Ac7059066C75F value: 47829690 balance: 430467210
block: 6 index: 3 address: 0xE9565da74A5149e43475e62646dC4f47c2FcEa6f value: 72900000 balance: 656100000
block: 6 index: 1 address: 0xd1391ba0Be0eECA4031e952402b92e26417D5F38 value: 90000000 balance: 810000000
block: 6 index: 4 address: 0xEdA947963D73F7Ca96c94843E2bd091d31Ec9F90 value: 65610000 balance: 590490000
block: 6 index: 8 address: 0xD03ae6Da0708D073f3f920Ea1AAEbfdc847971b3 value: 43046721 balance: 387420489
block: 6 index: 9 address: 0xD9bE8f83736b508514183363a1fec4c0Bd1E80Db value: 38742048 balance: 348678441
block: 6 index: 5 address: 0x0504d04808022aC69bBb37206FD97Edaef9268e9 value: 59049000 balance: 531441000


```

