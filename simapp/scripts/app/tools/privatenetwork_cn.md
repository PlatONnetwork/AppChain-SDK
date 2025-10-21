# 私有网络部署

## 生成配置参数

```
cd scripts/bin
./gensimapp.sh gen [IPS] [USERNAME] [PASSWORD] 
```

假设安装四个节点：10.1.1.33,10.1.1.34,10.1.1.35,10.1.1.36 服务器账户密码统一为simapp/123456，则命令为`./gensimapp gen 10.1.1.33,10.1.1.34,10.1.1.35,10.1.1.36 simapp 123456`，scripts/config目录下将生成以下文件：

- benchmarkenv: 压测的环境变量配置，包括RPC地址、压测性能参数配置。
- l1checkpointsender.json：用以向Layer1发送checkpoint交易的账户keystore，需要保证有足够余额。
- l1checkpointsender_password：l1checkpointsender.json中账户keystore的密码。
- l2txsender.json：Layer2发送系统交易的账户keystore。
- l2txsender_password：l2txsender.json中账户keystore的密码。
- simappenv：应用链部署的环境变量配置。

## 创建链

```
source scripts/config/simappenv #应用环境变量
cd bin #切换到应用程序目录
./simappclient deploytemplate #Layer1部署合约模板
./simappclient createnode #生成节点配置
./simappclient deploychildchain #部署Layer2在Layer1的合约
./simappclient creategenesis #创建创世文件
./simappclient createansible #创建ansible目录
./simappclient createansiblenode #创建ansible节点
cd scripts/output/ansible #切换到ansible目录
ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml #部署到目标机器
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init" #初始化节点
ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start" #启动节点
```
## 压测

```
cd scripts/bin
./benchmarkcluster.sh gencontracttx 200000 #生成200000(ERC20合约)转账交易
./benchmarkcluster.sh status #查看是否生成完cache:200000
./benchmarkcluster.sh startcontracttx #启动压测
./benchmarkcluster.sh report [startblock] [endblock] report.csv #压测结束使用report命令生成压测区块数据
```