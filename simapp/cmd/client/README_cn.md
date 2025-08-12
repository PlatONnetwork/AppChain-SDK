# simapp 部署工具

## 创建输出文件夹

```shell
mkdir  output
```

## 部署模板合约

设置参数，采用环境变量方式

```
//与创建的文件夹路径一致
export OUTPUT=./output
//部署合约的私钥
export DEPLOY_KEY=8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd
//初始化模板合约私钥
export INITIALIZE_KEY=5856fb978afcb48fd023a0e6ba76eb90ed35de98ca4a6873d28d1c5b9a47f39e
//rootchain rpc地址
export ROOTCHAIN_URL=http://127.0.0.1:8801
//注册管理合约的Owner
export REGISTRY_MANAGER_OWNER=0x8a0f3F8389F79a05Dd03027dE0Bcbb06cADB3F9c
```

执行部署
```shell
./client deploytemplate
```
输出文件为
```shell
output/template.toml
```
## 创建Layer2 节点配置文件
```shell
//创世节点host
export GENESIS_NODE=10.1.1.33,10.1.1.34,10.1.1.35
//普通节点host
export NORMAL_NODE=10.1.1.36
```
执行
```shell
./client createnode
```
输出文件为
```shell
output/nodes.toml
```

## 部署子链合约

```shell
//创世节点的奖励地址
export STAKE_OWNER=0x8a0f3F8389F79a05Dd03027dE0Bcbb06cADB3F9c
//子链合约Owner地址
export CHILDCHAIN_OWNER=5856fb978afcb48fd023a0e6ba76eb90ed35de98ca4a6873d28d1c5b9a47f39e
```
输出文件为
```shell
output/childchaincontract.toml
```

## 创建 genesis.json

```shell
export ROOTCHAINID=101
```
执行
```shell
./client creategenesis
```

输出文件为
```shell
output/genesis.json
```

## 创建ansible 文件夹

```shell
//远程部署节点路径
export REMOTE_ANSIBLE_DIR=~/opt
```
执行
```shell
./client createansible
```

输出文件为
```shell
output/ansible
```

## 创建 ansible 节点配置

```shell
//simapp 执行程序路径
export BIN=./node
//远程服务器账户名
export USERNAME=simapp
export PASSWORD=123456
```

## 节点部署
切换到ansible 工程文件夹
```shell
cd output/ansible
```
执行部署
```shell
ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml
```

初始化

```shell
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init"
```

启动
```shell
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start"
```

停止
```shell
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=stop"
```
清理
```shell
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=clean"
```