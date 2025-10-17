## PlatON应用链测试环境

PlatON实现了一个基于PlatON的L2应用链，采用PlatON相同的共识算法、虚拟机、经济模型，该项目的源码位于xxxxx。

### 测试应用链

PlatON已经部署了一套用于测试的PlatON L2应用链，测试环境有4个节点组成，IP地址分别为8.219.210.248、8.222.253.224、8.222.207.89、8.219.242.242，浏览器地址为[http://8.219.177.143:8000/](http://8.219.177.143:8000/ "http://8.219.177.143:8000/")。

用户可在测试环境进行合约和dAPP开发测试，也可自行搭建节点加入测试链，并进行性能验证，方法如下：


#### 下载测试环境配置
操作系统为 Ubuntu 24.04 LTS。下载链接为：
https://app-chain.oss-ap-southeast-1.aliyuncs.com/testnetnode.tar.gz
#### 生成配置

* `cd scripts/bin` 进入到脚本目录
* `./generate.sh  --stakeamount 1000000000`

生成节点所需的私钥，1000000000 表示该节点将要质押的代币数量。

#### 测试Token

在 `generate.sh` 脚本执行后会有如下信息，请将`user address`、`checkpoint address`对应的地址提供给我们来获取测试 Token。
```
generate keys success please contact the official team to ensure sufficient balance for both the user address:0x6f9f59826097D25f0D6F4d719805D1cF8ba7a07d and checkpoint address:0x1Df85FA030d2E00bA28EfA7115eF5dB88B14CcDD
```

* **user address:** 验证人账户地址，需要获取PlatON测试网LAT与应用链测试Token
* **checkpoint address:** 发送 checkpoint交易账户地址，需要获取PlatON测试网LAT

#### 启动节点

* 切换到项目根目录
* `./init.sh`
* `./start.sh`

#### 节点质押

在执行质押操作前，请确认您的节点已完全同步到最新区块。否则在执行质押操作时可能会受到处罚。
##### 执行流程

* 检查 `config/env` 参数是否正确。默认下发的配置已与节点保持一致。
* 进入 `bin` 目录。
* 执行 `./staking.sh stake` 进行质押。如果余额不足，请联系官方团队。
* 执行 `./staking.sh validators` 检查是否已成为验证节点。这通常需要一些时间。
* 执行 `./staking.sh unstake` 进行取消质押。节点将退出验证人列表。
* 执行 `./staking.sh validators` 检查节点是否已退出验证人列表。这通常需要一些时间。

#### 性能测试

* `cd scripts/bin`
* `./benchmark.sh gencontracttx 200000`  生成200000 ERC20 合约转账交易
* `./benchmark.sh status` 查看是否生成完 cache:200000
* `./benchmark.sh startcontracttx` 启动压测
* `./benchmark.sh status` 查看是否发送完 sent:200000

通过浏览器可查询到区块交易情况




### 私有应用链部署

用户可以基于该源码搭建自己的PlatON L2应用链，下面提供基于编译后的私有链配置环境进行搭建：

#### 下载版本
操作系统为 Ubuntu 24.04 LTS（安装ansible）。下载链接为：
https://app-chain.oss-ap-southeast-1.aliyuncs.com/privatenetwork.tar.gz

#### 生成配置参数

- `cd scripts/bin` 进入到脚本目录
- `./gensimapp.sh [IPS] [USERNAME] [PASSWORD]`，如机器为10.1.1.33,10.1.1.34,10.1.1.35,服务器账户密码为simapp, 123456则命令为`./gensimapp.sh gen 10.1.1.33,10.1.1.34,10.1.1.35 simapp 123456`，scripts/config 将生成如下文件
  - **benchmarkenv** 压测的环境变量配置，主要是RPC地址、压测性能参数配置
  - **l1checkpointsender.json**  向 Layer1 发送checkpoint交易的keystore 账户文件，需要保证有足够余额
  - **l1checkpointsender_password** `l1checkpointsender.json` 的keystore密码
  - **l2txsender.json**  Layer2 发送系统交易的账户
  - **l2txsender_password** `l2txsender.json` 的keystore密码
  - **simappenv** 应用链部署的环境变量配置

#### 测试代币

在 `gensimapp.sh` 脚本执行后会有如下信息，请确保`deploy address`、`initialize address`、`checkpoint address`对应的地址有充足代币。
```
generate keys success please contact the official team to ensure sufficient balance for both the deploy address:0x311F5F59B1170e095709eE8da4f6880d912946F3, initialize address:0x8CFF458f3d7cFF30943701979b336adFf42D2776 and checkpoint address:0xabEe1cCfb84853d954f8a005a1Becc18d37A86ec
```
**deploy address:** 部署合约账户地址，需要PlatON测试网LAT
**initialize address:** 合约初始化账户地址，需要PlatON测试网LAT
**checkpoint address:** 发送 checkpoint 交易账户地址，需要PlatON测试网LAT
#### 创建链


- `source scripts/config/simappenv`应用环境变量
- `cd bin`切换到应用程序目录
- `./simappclient deploytemplate`Layer1 部署合约模板
- `./simappclient createnode`生成节点配置
- `./simappclient deploychildchain`部署Layer2 在Layer1的合约
- `./simappclient creategenesis`创建创世文件
- `./simappclient createansible`创建ansible 目录
- `./simappclient createansiblenode`创建ansible 节点
- `cd scripts/output/ansible`切换到ansible目录
- `ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml`部署到目标机器
- `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init"`初始化节点
- `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start"`启动节点

#### 性能测试

- `cd scripts/bin`
- `./benchmarkcluster.sh gencontracttx 200000`生成200000 ERC20 合约转账交易
- `./benchmarkcluster.sh status`查看是否生成完 cache:200000
- `./benchmarkcluster.sh startcontracttx`启动压测
- `./benchmarkcluster.sh report [startblock] [endblock] report.csv`压测结束使用 report 命令 生成压测区块数据
