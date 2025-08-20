# ConsensusNetwork

ConsensusNetwork 模块实现共识网络的扩展，通过实现 ConsensusNetworkModule 模块接口，实现网络一系列操作

* ConsensusNetworkModule 实现共识网络消息发送
* ViewChangeModule 更新验证人
* P2PModule P2P 协议
```go
type ConsensusNetworkEngine interface {
    //启动网络引擎
	StartNetworkEngine()
	//发布消息
	Publish(msg ctypes.Message)
    //广播
	Broadcast(msg ctypes.Message)
    //部分广播
	PartBroadcast(msg ctypes.Message)
    //转发消息
	Forwarding(nodeID string, msg ctypes.Message) error
    //向指定peer发送消息
	Send(peerID string, msg ctypes.Message)
    //网络平均延时
	AvgLatency() time.Duration
	// 设置peer 状态，bType 有QC、Lock、Commit三种区块共识状态
	PeerSetting(peerID string, bType uint64, blockNumber uint64) error
    //移除peer
	RemovePeer(id string)
}
```

## 消息转发逻辑

# 消息类型-转发规则

| 消息类型             | 消息说明                     | 发送规则                                     | 转发规则                           | 说明                                        | 备注                                             |
| :------------------- | :--------------------------- | :------------------------------------------- | :--------------------------------- | ------------------------------------------- | ------------------------------------------------ |
| PrepareBlock         | 提议区块消息                 | Leader节点发送给所有共识节点和部分非共识节点        | 转发给所有共识节点和部分非共识节点            | 全网所有共识节点都需要接收  | 消息尽可能扩散到所有节点（和之前版本保持一致） |
| PrepareVote          | 区块的投票签名消息           | 投票节点只发送给Leader节点                   | 不转发                             | Leader节点负责收集投票并生成BlockQuorumCert | 消息只向Leader节点扩散 |
| BlockQuorumCert      | 对区块投票的完整聚合签名消息 | Leader节点发送给所有共识节点和部分非共识节点（优先发送给下一任Leader） | 转发给部分非共识节点 | 全网所有节点都需要接收                      | 消息尽可能扩散到所有节点（之前版本不转发）                |
| ViewChange           | 视图切换的签名消息           | 发起节点发送给所有共识节点                   | 转发给所有共识节点                 |                                             | 消息只在共识节点间扩散          |
| ViewChangeQuorumCert | 对视图切换的完整聚合签名信息 | 发起节点发送给所有共识节点和部分非共识节点   | 转发给所有共识节点和部分非共识节点 | 全网所有节点都需要接收                      | 消息尽可能扩散所有节点（之前版本不发送也不转发） |
| GetPrepareBlock | 区块同步消息 | 发起节点发送给部分共识节点 | 不转发 |                       |  |
| GetPrepareVote | 区块投票同步消息 | 无 | 无 |                       |  |
| GetViewChange | 视图投票同步消息 | 发起节点发送给部分共识节点 | 不转发 |                       |  |
| GetLatestStatus | 共识状态同步消息 | 发起节点发送给部分共识节点 | 不转发 |                       |  |


## 构建

```go
consensusNetwork = consensusnetwork.NewModule(ctx)
manager := module.NewManager(...,consensusNetwork)
manager.SetOrderBlockCommitter(...,consensusNetwork.Name())
manager.SetOrderInit(...,consensusNetwork.Name())
manager.SetConsensusNetwork(consensusNetwork.Name())
```
