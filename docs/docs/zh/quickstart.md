# 快速开始

## 编译

```shell
make simapp
make example
```
使用到的应用程序

```
build/bin/simapp
build/bin/simapp_client
example/build/bin/emptynode
```

## 启动

* 启动L1 节点
```shell
emptynode server
```

* 初始化环境变量

```shell
cd simapp/quick_start
source ./env_simapp_client
```

* 部署L1合约

模板合约
```shell
simapp_client deploytemplate
```
子链合约

```shell
simapp_client deploychildchain
```

* 创建创世文件

```shell
simapp_client creategenesis 
```

* 创建部署脚本

```shell
simapp_client createansible
simapp_client createansiblenode
```

* 部署及初始化 L2节点

节点默认部署在`~/opt/node` 路径下
```shell
cd output/ansible
ansible-playbook playbooks/deploy.yml
ansible-playbook playbooks/command.yml --extra-vars "cmd=init"
ansible-playbook playbooks/command.yml --extra-vars "cmd=start"
```

## 初始化工具环境

```shell
source env_tools
```

## 查看 Token 余额

* 铸币

```shell
tools cast --module erc20 --address $token --rpc $rootchainurl --type send --method mint $user 100000000000000000
```

* 查看 L1 余额
```shell
tools cast --module erc20 --address $token --rpc $rootchainurl --type call --method balanceOf $user
```


* 查看 L2 余额

```shell
simapp attach $rootchainurl
> platon.getBalance("0xE9565da74A5149e43475e62646dC4f47c2FcEa6f")
```

## L1->L2 转账

* Approve DepositManager 合约

```shell
tools cast --module erc20 --address $token --rpc $rootchainurl --type send --method approve $l1depositmanager 2000000000000000000
* Deposit
```shell
tools cast --module l1.deposit --address $l1depositmanager --rpc $rootchainurl --type send --method deposit $user 100000000000000000
```

* StateSync 同步

```shell
tools cast --module l2.statesync --rpc $childchainurl  --address $l2statesync  --type call --method getExecutedId
```
id 为1 表明事件已经从L1->同步到L2

* 查看余额变化

```shell
simapp attach http://node1.appchain.network:8801
> platon.getBalance("0xE9565da74A5149e43475e62646dC4f47c2FcEa6f")
```

## L2 提取

* 发送提取交易

```shell
tools cast --module l2.deposit --rpc $childchainurl  --address $l2deposit --type send --method withdraw $user 1000
```

* 查看事件

```shell
tools cast --module l2.statesender --rpc $childchainurl  --address $l2statesender --type logs --method L2StateSynced
```


* 查看余额变化

```shell
simapp attach $childchainurl
> platon.getBalance("0xE9565da74A5149e43475e62646dC4f47c2FcEa6f")
```

* L2 查看checkpoint 是否生成提取交易证明
  params中的1 为l2.statesender事件中Id字段值
```shell
curl $childchainurl -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"checkpoint_generateExitProof","params":[1],"id":0}'
```

* L1 发送提取交易

```shell
tools cast l1.exithelper --address $l1exithelper --exit-id 1 
```

* 查看 L1 余额
```shell
tools cast --module erc20 --address $token --rpc $rootchainurl --type call --method balanceOf $user
```