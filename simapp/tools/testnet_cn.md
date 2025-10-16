# 生成配置

* `cd scripts/bin` 进入到脚本目录
* `./generate.sh  --stakeamount 1000000000`

生成节点所需的私钥，1000000000 表示该节点将要质押的代币数量。

# 启动节点

* 切换到项目根目录
* `./init.sh`
* `./start.sh`

# 质押脚本
简单的质押与取消质押脚本

!!!~~~~ 在使用此脚本前，请确认您的节点已完全同步到最新区块。否则在执行质押操作时可能会受到处罚。
## 执行流程

* 检查 `config/env` 参数是否正确。默认下发的配置已与节点保持一致。
* 进入 `bin` 目录。
* 执行 `./staking.sh stake` 进行质押。如果余额不足，请联系官方团队。
* 执行 `./staking.sh validators` 检查是否已成为验证节点。这通常需要一些时间。
* 执行 `./staking.sh unstake` 进行取消质押。节点将退出验证人列表。
* 执行 `./staking.sh validators` 检查节点是否已退出验证人列表。这通常需要一些时间。

# 压测

* `cd scripts/bin`
* `./benchmark.sh gencontracttx 200000`  生成200000 ERC20 合约转账交易
* `./benchmark.sh status` 查看是否生成完 cache:200000
* `./benchmark.sh startcontracttx` 启动压测
* `./benchmark.sh status` 查看是否发送完 sent:200000

通过浏览器可查询到区块交易情况