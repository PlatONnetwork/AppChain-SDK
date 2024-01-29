# Worker 扩展

## 简介

Worker扩展主要是对区块打包流程进行扩展。为开发者提供两个接口`WorkerModule`和`TransactionModule`，分别用在以下场景：

- 排序交易，改变交易的执行顺序。

- 增加交易，且增加的交易优先打包，保证打包成功。

## WorkerModule

!!! info "SortTxs"
    `SortTxs(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) (types.Transactions, error)`

    参数：

    - **ctx** AppChain SDK worker context.
    - **local** 本地的交易列表。
    - **remote** 远程的交易列表（通过P2P广播接收的交易）。

    返回值：

    - 排序成功，返回排序后的交易列表（按先local后remote排序）。

    - 失败返回具体的错误。

## TransactionModule

!!! info "AddTxs"

    `AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error)`

    参数：

    - **ctx** AppChain SDK worker context.
    - **local** 本地的交易列表。

    返回值：

    - 成功，模块增加的交易包含在返回的交易列表中。

    - 失败返回具体的错误。


```go title="x/statesync/module.go"
func (s *StateSync) AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error) {
	block := ctx.Backend().GetBlock(ctx.Header().ParentHash, ctx.Header().Number.Uint64()-1)
	if block == nil {
		s.logger.Warn("Get block failed", "number", ctx.Header().Number.Uint64()-1)
		return local, nil
	}
	receiver, err := s.newStateSyncCallContract(ctx, block.Header())
	if err != nil {
		s.logger.Warn("New state sync caller failed", "err", err)
		return local, nil
	}
	from := crypto.PubkeyToAddress(s.privateKey.PublicKey)

	nonce := common2.EnableNonce(local[from], func() uint64 {
		return ctx.StateDB().GetNonce(from)
	})
	if err != nil {
		s.logger.Warn("Get pool nonce failed", "nonce", nonce, "err", err)
		return local, nil
	}
	cmtx, err := s.addCommitTx(ctx, receiver, nonce)
	if err == nil {
		if local[from] == nil {
			local[from] = types.Transactions{}
		}
		local[from] = append(local[from], cmtx)
		nonce += 1
	}
	exTxs, err := s.addExecutedTx(ctx, receiver, nonce)
	if err != nil {
		s.logger.Warn("Get executed txs failed", "err", err)
		return local, nil
	}
	if local[from] == nil {
		local[from] = types.Transactions{}
	}
	local[from] = append(local[from], exTxs...)
	return local, nil
}
```
