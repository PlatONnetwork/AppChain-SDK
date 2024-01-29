# 快速开始

## 环境准备

* docker, 没有安装参考[官方文档](https://docs.docker.com/engine/install)
* docker-compose. 没有参考[官方文档](https://docs.docker.com/compose/install/linux/#install-the-plugin-manually)

## Docker 镜像说明

**rootchain:** L1 链节点，对外暴露 8800 http rpc 端口

**appchain-cluster:** L2 节点，docker-compose 部署3节点集群，HTTP RPC 端口分别为 node1: 8801, node2:8802, node3:8803

**appchain-contract:** 系统合约，包括staking,checkpoint,statesync 等

## 启动

* 启动L1 节点
```shell
docker-compose up -d rootchain
```

* 部署L1合约
```shell
docker-compose run l1init -i bash
./start.sh
```

* 启动 L2 节点

```shell
docker-compose up -d node1 node2 node3
```

## L1 合约地址

```shell
ChildChainId=123083
StateSender=0x92b68E65Ef272afA1B8F001501D17F13B836E735
CheckpointManager=0x0bb7AC917fA06100cDfF38641E6c2C6Bae4F4890
ExitHelper=0x7B66bf9dAba52733231e32D3e2fe1008205c1458
CustomChildChainManager=0x8BbfE720aEbC36EE7B2F54866fbED59f94360fDD
StakeManager=0xB04B468EB388C613cb453A60f6d610A009E1Ff55
DepositManager=0x628f6017E774F8Fc0fBf59C1ea18CB4e398c9168
WithdrawHandler=0x0000000000000000000000000000000000000000
Token=0x88edBfC0d17BFa9ffC064244Bc90Bf0F264D0880
Creator=0x4779472533Ed1B8Fe6d6E0d212F2aA02Ffd84b9F
Owner=0x62953f9213f899f2A51680C2fBB4282a2591bfc8
```
L2 合约地址

```shell
L2Stake=0x1000000000000000000000000000000000000005
L2Deposit=0x1000000000000000000000000000000000000007
```

## 启动工具 Docker

```shell
docker-compose run tool bash
```

## 查看 Token 余额

* 铸币

```shell
tool cast erc20 --rpc http://rpc.rootchain.network:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x88edBfC0d17BFa9ffC064244Bc90Bf0F264D0880 --type send --method mint 0x62953f9213f899f2A51680C2fBB4282a2591bfc8 20000000000
```

* 查看 L1 余额
```shell
tool cast erc20 --rpc http://rpc.rootchain.network:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x88edBfC0d17BFa9ffC064244Bc90Bf0F264D0880 --type call --method balanceOf 0x62953f9213f899f2A51680C2fBB4282a2591bfc8
```


* 查看 L2 余额

```shell
simapp attach http://node1.appchain.network:8801
> platon.getBalance("0x62953f9213f899f2A51680C2fBB4282a2591bfc8")
```

## L1->L2 转账

* Approve DepositManager 合约

```shell
tool cast erc20 --rpc http://rpc.rootchain.network:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x88edBfC0d17BFa9ffC064244Bc90Bf0F264D0880 --type send --method approve 0x628f6017E774F8Fc0fBf59C1ea18CB4e398c9168 200000000000000
```

* Deposit
```shell
tool cast l1.deposit --rpc http://rpc.rootchain.network:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x628f6017E774F8Fc0fBf59C1ea18CB4e398c9168  --type send --method deposit 0x62953f9213f899f2A51680C2fBB4282a2591bfc8 200000000
```

* StateSync 同步

```shell
tool cast l2.statesync --rpc http://node1.appchain.network:8801 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x1000000000000000000000000000000000000002  --type call --method getExecutedId
```
id 为1 表明事件已经从L1->同步到L2

* 查看余额变化

```shell
simapp attach http://node1.appchain.network:8801
> platon.getBalance("0x62953f9213f899f2A51680C2fBB4282a2591bfc8")
```

## L2 提取

* 发送提取交易

```shell
tool cast l2.deposit --rpc http://node1.appchain.network:8801 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x1000000000000000000000000000000000000007 --type send --method withdraw 0x62953f9213f899f2A51680C2fBB4282a2591bfc8 1000
```

* 查看事件

```shell
tool cast l2.statesender --rpc http://node1.appchain.network:8801 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x1000000000000000000000000000000000000001 --type logs --method L2StateSynced  "[]" "[]" "[]"
```


* 查看余额变化

```shell
simapp attach http://node1.appchain.network:8801
> platon.getBalance("0x62953f9213f899f2A51680C2fBB4282a2591bfc8")
```

* L2 查看checkpoint 是否生成提取交易证明
  params中的1 为l2.statesender事件中Id字段值
```shell
curl http://node1.appchain.network:8801 -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"checkpoint_generateExitProof","params":[1],"id":0}'
```

* L1 发送提取交易

```shell
tool cast l1.exithelper --appchain-rpc http://node1.appchain.network:8801 --rpc http://rpc.rootchain.network:8800 --key 5cd44b68d91f4b42dc240c658fee14922a0e26245ed6f406c0015bb50cf66431 --exit-id 1 --address 0x7B66bf9dAba52733231e32D3e2fe1008205c1458
```

* 查看 L1 余额
```shell
tool cast erc20 --rpc http://192.168.213.2:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x88edBfC0d17BFa9ffC064244Bc90Bf0F264D0880 --type call --method balanceOf 0x62953f9213f899f2A51680C2fBB4282a2591bfc8
```