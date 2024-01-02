# 交易测试工具

交易测试工具主要针对 AppChain-SDK 中内置模块合约以及 L1中重要合约的交易调用，事件过滤

工具支持的合约有：

**L1**

* ERC20，操作L1代币合约
* DepositManager，代币质押合约，L1->L2跨链转账
* ExitHelper，用于提交L2->L1动作，如L2->L1提取，L2->L1 解质押等
* StakeManager，用于进行Staking相关动作，如质押，委托，释放等

**L2**

* Deposit，接收L1代币质押
* Reward，处理Staking中奖励
* Staking，处理L1->L2质押，委托，释放等，以及L2的奖励，处罚等
* StateSender，生成L2->L1事件，用于checkpoint提交


## 使用

* 查看合约ABI

```shell
tool cast erc20 --type abi

methods:
   transfer(address,uint256)
   transferFrom(address,address,uint256)
   allowance(address,address)
   approve(address,uint256)
   balanceOf(address)
   mint(address,uint256)
   totalSupply()
events:
   Approval(address indexed,address indexed,uint256)
   Transfer(address indexed,address indexed,uint256)

```


* 发送交易

```shell
tool cast erc20 --rpc {ip} --key {sender_key} --address {contract address} --type send --method mint 0x628f6017E774F8Fc0fBf59C1ea18CB4e398c9168 200000000000000
```
**--rpc:** 节点RPC地址
**--key:** 发送者私钥
**--address:** 合约地址
**--type:** 发送类型，
**--method:** 调用合约函数名，后面为具体函数

* 查询交易

```shell
tool cast erc20 --rpc {ip} --address {contract address} --type call --method balanceOf 0x628f6017E774F8Fc0fBf59C1ea18CB4e398c9168
```
返回 

```shell
[
  20000000000
 ]
```

* 查询事件

```shell
tool cast l2.statesender --rpc http://node1.appchain.network:8801 --address 0x1000000000000000000000000000000000000001 --type logs --method L2StateSynced  "[]" "[]" "[]"
```
**--method:** 事件名称，后面的参数为事件中具有索引的字段，为数组类型，`L2StateSynced(uint256 indexed Id,address indexed Sender,address indexed Receiver,bytes CallData)`如过滤L2StateSynced 中的Id，则参数可以写为`"[1,2,3]"` 

返回l

```shell
[
  {
   "Id": 1,
   "Sender": "0x1000000000000000000000000000000000000007",
   "Receiver": "0x628f6017E774F8Fc0fBf59C1ea18CB4e398c9168",
   "CallData": "eo3CZ5ah5Q5uGQtwJZ9Y9qTt1bIigM7syCtoe46YKGkAAAAAAAAAAAAAAABilT+SE/iZ8qUWgML7tCgqJZG/yAAAAAAAAAAAAAAAAGKVP5IT+JnypRaAwvu0KColkb/IAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA+g=",
   "Raw": {
    "address": "0x1000000000000000000000000000000000000001",
    "topics": [
     "0xedaf3c471ebd67d60c29efe34b639ede7d6a1d92eaeb3f503e784971e67118a5",
     "0x0000000000000000000000000000000000000000000000000000000000000001",
     "0x0000000000000000000000001000000000000000000000000000000000000007",
     "0x000000000000000000000000628f6017e774f8fc0fbf59c1ea18cb4e398c9168"
    ],
    "data": "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000807a8dc26796a1e50e6e190b70259f58f6a4edd5b22280ceecc82b687b8e98286900000000000000000000000062953f9213f899f2a51680c2fbb4282a2591bfc800000000000000000000000062953f9213f899f2a51680c2fbb4282a2591bfc800000000000000000000000000000000000000000000000000000000000003e8",
    "blockNumber": "0x73b",
    "transactionHash": "0x9908e67b784e4c94cbb89fdf8c7ef35a81f978a2aceadbb917c123b207f58943",
    "transactionIndex": "0x1",
    "blockHash": "0x88faa442eb62cc3f5a45312e0f17384a36ed0482690e0a85428753880105e8d4",
    "logIndex": "0x1",
    "removed": false
   }
  }
 ]
```

Id，Sender，Receiver，CallData，为事件各个字段的值，Raw 为log的原始信息