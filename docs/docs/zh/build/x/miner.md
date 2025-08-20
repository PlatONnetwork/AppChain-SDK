# Miner

miner 模块实现 PEVM 虚拟机执行交易，主要实现 FillTransactionsModule、TxExecutorModule 

## 启动参数

* miner.concurrency_level PEVM 虚拟机并行度
* miner.force_sequential PEVM 是否强制串行执行
* miner.txs_batch 每批多少交易进行并行执行

## 构建

```go
//加载参数
miner.AddModuleInitFlags(cliApp)
//创建模块
miner := miner.NewModule(ctx)
manger := module.NewManager(...,miner)

manager.SetOrderGenesis(...,asyncBlock)
//设置FillTransaction
manager.SetTxFiller(miner.Name())
//设置ExecuteTxs
manager.SetTxExecutor(miner.Name())
```