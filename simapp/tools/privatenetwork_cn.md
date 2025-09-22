# 私有网络部署

## 生成配置参数

* `cd scripts/bin` 进入到脚本目录
* `./gensimapp.sh [IPS] [USERNAME] [PASSWORD]`，如机器为10.1.1.33,10.1.1.34,10.1.1.35,服务器账户密码为simapp,123456则命令为`./gensimapp 10.1.1.33,10.1.1.34,10.1.1.35 simapp 123456，scripts/config 将生成`benchmarkenv  l1checkpointsender.json  l1checkpointsender_password  l2txsender.json  l2txsender_password  simappenv`文件
    * benchmarkenv 压测的环境变量配置，主要是RPC地址、压测性能参数配置
    * l1checkpointsender.json l1checkpointsender_password 向Layer1 发送checkpoint交易的账户，需要保证有足够余额
    * l2txsender.json  l2txsender_password Layer2 发送系统交易的账户
    * simappenv app部署的环境变量配置

## 创建链

* `cd bin` 切换到应用程序目录
* `source scripts/config/simappenv` 应用环境变量
* `./simappclient deploytemplate` Layer1 部署合约模板
* `./simappclient createnode` 生成节点配置
* `./simappclient deploychildchain` 部署Layer2 在Layer1的合约
* `./simappclient creategenesis` 创建创世文件
* `./simappclient createansible` 创建ansible 目录
* `./simappclient createansiblenode` 创建ansible 节点
* `cd scripts/output/ansible` 切换到ansible目录
* `ansible-playbook -i inventories/hosts.yml playbooks/deploy.yml` 部署到目标机器
* `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=init"` 初始化节点
* `ansible-playbook -i inventories/hosts.yml playbooks/command.yml --extra-vars "cmd=start"` 启动节点

## 压测

* `cd scripts/bin`
* `./benchmarkcluster.sh gencontracttx 200000`  生成200000 ERC20 合约转账交易
* `./benchmarkcluster.sh status` 查看是否生成完 cache:200000
* `./benchmarkcluster.sh startcontracttx` 启动压测
* `./benchmarkcluster.sh report [startblock] [endblock] report.csv` 压测结束使用 report 命令 生成压测区块数据