# BENCHMARK


性能测试工具模块

## 编译

```shell
cd x/benchmark && make
```

在 x/benchmark/build/bin 生成执行程序，benchmark 为server运行，client用于控制压测模块，生成压测配置

## 部署流程

### ansible 工程


```shell
mkdir  output
```

```shell
./client createansible
```

### 生成节点文件


```shell
./client --hosts "127.0.0.1,127.0.0.2" --ansible_dir ~/tmp/ansible --user sdk --password 123123 --bin benchmark
```
分别生成 output/ansible/inventories/hosts.yml ，output/ansible/playbooks/vars/env.yml配置。在output/ansible/playbooks/files 生成分发到各个节点的数据


### ansible 部署节点
切换到 output/ansible

```shell
ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml
```

### ansible 启动节点
```shell
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init"
```

### ansible 启动节点
```shell
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start"
```


### 压测准备

配置参数
```shell
//配置rpc
export BENCHMARK_URLS=http://192.168.21.17:8801
//压测地址数
export BENCHMARK_ADDRS=10
//每个账户每次区块最多发送笔数
export BENCHMARK_TXSPERACCOUNT=100
//转账交易占生成交易的比例，0-100
export BENCHMARK_RAWTX=0
//生成压测交易数量
export BENCHMARK_COUNT=10000
//每秒钟发送交易数
export BENCHMARK_TPS=100
```

### 生成交易

```shell
./client gentx
```

### 启动

```shell
./client start
```

### 查看状态

```shell
./client status
```

### 停止压测

```shell
./client stop
```

### ansible 停止节点
```shell
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=stop"
```