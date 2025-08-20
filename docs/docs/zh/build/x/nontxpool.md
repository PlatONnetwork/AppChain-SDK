# Nontxpool

Nontxpool 模块是为了应对传统交易池性能问题做出优化，主要以下几点：

* 取消交易P2P gossip 广播，采取转发Leader形式或关闭P2P转发，减少高压力下网络压力
* 取消交易池Reset，TxPool的reset操作，再区块交易数过多时，reset 性能显著下降


## 启动参数

* nontxpool.txscachesize 多少笔交易进行广播
* nontxpool.broadcastinterval 广播交易间隔时间
* nontxpool.validatetx 是否验证接收到的交易
* nontxpool.pendinglimit 每次从交易池拿多少笔交易
* nontxpool.txsperaccount 每个账户每次最多拿多少笔交易
* nontxpool.broadcast 是否开启交易广播
* nontxpool.globaltxcount 交易池最大缓存交易数量

## 模块接口

Nontxpool 通过实现 WorkerModule 接口将交易给到worker，如果有其它模块实现WorkerModule 想调用Nontxpool 获取交易，则通过 Pending 获取交易
```go
Pending(getNonce func(addr common.Address) uint64, limit int, txsPerAccount int) types.Transactions
```

通过实现 ViewChangeModule 将本地的Remote 交易清除掉，非当前Leader 则对交易进行转发到Leader节点(需要开启广播交易)

通过实现 BlockCommitterModule 清除掉已经上链交易


## RPC

Nontxpool 实现pendingNonce用于获取指定地址的nonce值

```go
func (m *Module) APIs() []rpc.API {
	return []rpc.API{
		rpc.API{
			Namespace: "nontxpool",
			Version:   "1",
			Service:   NewRPC(m),
			Public:    true,
		},
	}
}

func PendingNonce(addr common.Address) (uint64, error)
```

## 构建

```go
nontxpool.AddNonTxPoolFlags(cliApp)
nonTxPool = nontxpool.NewModule(ctx)
manager := module.NewManager(...,nonTxPool)
manager.SetOrderBlockCommitter(nonTxPool.Name())
manager.SetOrderInit(nonTxPool.Name())
manager.SetOrderGenesis(nonTxPool.Name())
```