# 交易池扩展

## 简介

为了扩展交易池的功能，模块需要实现`TxPoolModule`，主要从两方面进行扩展：

1. 交易进入交易池模块可通过`CheckTx()`方法验证交易是否有效

2. 交易打包前模块可通过`FilterPendingTxs()`方法过滤交易

## CheckTx

!!! info "CheckTx"
    `CheckTx(ctx sdk.Context, tx *types.Transaction) error`

    参数：

    - **ctx** AppChain SDK base context
    - **tx** 待验证的交易

    返回值：验证通过返回`nil`，否则返回具体的错误。

## FilterPendingTxs

!!! info "FilterPendingTxs"

    `FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions`

    参数：

    - **ctx** AppChain SDK base context
    - **txs** 待过滤的交易列表

    返回值：返回过滤后的交易列表，如果没有过滤交易，则返回入参的交易列表。
