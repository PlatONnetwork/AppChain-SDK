# Checkpoints

Checkpoint是一个在特定时间点记录和提交系统快照的过程。在区块链中，checkpoint用于通过允许更快地验证区块链状态来提供安全性，而无需处理从创世块开始的所有交易。Checkpoint可用于加快新节点的同步时间，并帮助防止某些类型的攻击，例如51%攻击。

## Checkpoint实现

!!! note "关键点"

    - Checkpoint由L2的验证人生成并提交到L1。
    - Checkpoint是L2的快照。该快照存储为 Merkle 根，代表L2在创建时的状态。
    - 提交checkpoint对于确保L2网络安全非常重要，因为它使L1能够检测并防止针对子链任何潜在的欺诈或恶意活动。

### CheckpointManager

`CheckpointManager`负责管理网络中的checkpoint。

Checkpoint代表L2链状态的快照，验证人定期将checkpoint提交到L1。Checkpoint用作L2的参考点，以验证L2上数据的完整性和准确性。

该合约有多个函数用于管理checkpoint，例如提交checkpoint，验证签名以及使用块高或epoch获取退出事件的根。同时该合约存储了checkpoint和当前验证人列表，以及一个用于索引checkpoint块高的数组。

该合约使用Merkle树来证明L2的退出事件是否包含在checkpoint中。Merkle树使用L2StateSender生成的L2StateSync事件来构建，L2客户端在执行交易后将这些事件保存到本地存储中。成员证明由用户提供的Merkle证明进行验证。

该合约使用BLS算法来验证验证人提交的签名。验证人的签名被聚合，合约检查是否满足接受checkpoint所需的投票权阈值。

!!! note "Checkpoint详情"

    具体来说，Merkle树的根是一个哈希值，代表特定时间点L2链状态的特定子集。这个状态仅包含L2StateSender合约发出的退出事件。当用户想要退出L2（即将资产从L2转移到L1），他们的退出交易将包含在这棵Merkle树中。

    当checkpoint创建时，Merkle树的跟和其他元数据一起作为checkpoint的一部分包含在内。

    稍后，当用户想验证L2上的特定退出事件时，他们可以提供Merkle树中包含特定退出事件的证明，这是一种加密证明，证明Merkle树中包含特定退出事件。L1可以使用checkpoint中包含的Merkle树的根来验证Merkle证明。

    简而言之，Merkle树的根是L2特定时间点退出事件的紧凑表示，它包含在checkpoint中并用于验证是否包含某些退出事件。
