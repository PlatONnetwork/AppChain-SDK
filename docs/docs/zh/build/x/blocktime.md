# BlockTime

BlockTime 模块是控制出块时间、出块间隔。BlockTime 实现了 ConsensusBlockTimeModule 接口

由两个参数控制

* blocktime.nextblocktime 区块的最小区块间隔
* blocktime.blockproductiontimeout 生产区块最大的生产时间，超时时间为预留网络传播时间。如共识的出块间隔为1s，网络平均延时200ms，那么最大出块时间可设置为800ms
