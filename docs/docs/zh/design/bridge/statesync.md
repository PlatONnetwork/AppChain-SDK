# 状态同步

状态同步（StateSync）是一种用于根据L1上发生的事件更新L2上合约状态的机制。它是区块链技术的关键组成部分，因为它可以实现两条链之间安全高效的通信。状态同步允许以更高效、更安全的方式更新L2的链状态，而无需处理来自创世块的所有交易。

## 状态同步的实现

!!! info "关键点"

    - `状态同步`实现L2和L1之间高效、安全的数据传输。
    - `状态同步`通过 StateSender 合约在L2上启动，并通过 StateReceiver 合约在L1上执行。
    - `状态同步`机制用于根据L2上发生的事件更新L1上合约的状态

### StateSender

`StateSender`合约部署在L1上，并由关联的L1合约触发。它的主要职责是根据提供的数据和目标合约地址生成状态同步事件。任何人都可以调用`syncState`接口去发出状态同步事件。数据和事件一起发送，代表需要在L2上执行的状态变化。

### StateReceiver

`StateReceiver`合约部署在L2上，负责执行和转发L1发送的状态数据。它从L1合约接收状态变更数据，这些数据以`commitment`的形式捆绑在一起，并生成Merkle根一起发送出去。这棵Merkle树通过捆绑`StateSender`接收到的`StateSync`事件来创建。`Commitments`作为一个系统交易由区块提议人提交到`StateReceiver`合约。它们用于验证从L1到L2的状态数据的执行情况，例如从L1到L2的资金转移。`Commitments`与`checkpoints`类似，用于将数据从L1传输到L2的过程，而checkpoint则用于将数据从L2传输到L1的过程。


## L2StateSender 和 ExitHelper

为了实现从L2到L1的通信，`L2StateSender`部署在L2上负责发出`L2StateSync`事件（也称为退出事件）。这些事件由验证人排序并作为`checkpoint`提交到L1，从而允许延迟执行。与`StateSender`不同，`L2StateSender`不会向L1发送交易。

在L1上，`ExitHelper`负责验证和执行退出事件，从而使用户能够将L2资产提取到L1中。它类似于L2上的`StateReceiver`合约，两个合约协同工作，以实现L1和L2之间的双向通信。

!!! info "同步和承诺"

    `状态同步`主要分为两步：同步和承诺。

    在同步阶段，部署在L1的`StateSender`合约基于接收人和数据生成状态同步事件。任何人都可以通过合约`syncState`接口发出状态同步事件。数据随事件一起发送，代表需要在L2上执行的状态变化。

    在承诺阶段，L2上的`StateReceiver`合约接收状态改变数据以及来自`StateSender`合约的Merkle证明， 并验证该证明以确保数据的完整性。如果证明有效，则将在L2上执行状态变更。

    为了确保状态变更的有效性，`StateSender`合约为每个状态同步事件生成一个唯一的ID。`StateReceiver`使用该ID来防止重放攻击，因为重放可能导致执行重复的状态变更。

    `StateReceiver`合约还实现了BLS签名方案，以验证验证人提交的签名。验证人的签名聚合后，合约检查是否满足接受状态变更所需的投票权阈值。
