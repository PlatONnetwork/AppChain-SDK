# 脚本说明

* deploy 文件夹为生成节点staking脚本工具，
  * deploytestnet.sh 生成新节点环境配置，加入到已知测试网络
  * deployprivatenetwork.sh 生成新的 Layer2 集群网络
* tools 为节点脚本相关模板文件

## deploytestnet 使用说明

* 确认已经使用 simapp_client 工具生成节点相关脚本，如果没有请联系测试网络管理者
* 进入 deploy 文件夹
* 创建输出文件夹 `mkdir testnestnode` 
* 执行 `./deploytestnet.sh empty [output] [rootchainurl] [childchainurl] [toolsPath] [benchmarkclient] [newNodeDir]` e.g. `./deploytestnet.sh empty output https://devnet2openapi.platon.network/rpc http://127.0.0.1:8801 ../../../../build/bin/tools  ../../../../x/benchmark/build/bin/benchmark_client testnestnode`
    * output 为 simapp_client 相关文件输出的文件夹
    * rootchainurl layer1 rpc地址
    * childchainurl 本地节点rpc地址，通常为 http://127.0.0.1:8801
    * toolsPath tools 工具路径，需要将其拷贝到节点脚本执行目录， 默认在`build/bin/tools`
    * benchmarkclient 压测客户端路径，生成压测交易、执行压测， 默认在`x/benchmark/build/bin/benchmark_client`
    * newNodeDir 生成的环境配置输出路径
  
## deployprivatenetwork 使用说明

* 进入 deploy 文件夹
* 创建输出文件夹，`mkdir privatenetwork`
* 执行 `./deployprivatenetwork.sh all [OUTPUTDIR] [SIMAPP] [SIMAPPCLIENT] [TOOLS] [BENCHMARKCLIENT]` e.g. `./deployprivatenetwork.sh all privatenetwork ../../../../build/bin/simapp ../../../../build/bin/simapp_client ../../../../build/bin/tools ../../../../x/benchmark/build/bin/benchmark_client`
  * OUTPUTDIR 为生成的环境配置输出路径
  * SIMAPP simapp 二进制路径，默认在 `build/bin/simapp`
  * SIMAPPCLIENT simapp 客户端工具，用于部署集群环境。默认在 `build/bin/simapp_client`
  * TOOLS tools 工具路径，执行系统合约相关交易， 默认在 `build/bin/tools`
  * benchmarkclient 压测客户端路径，生成压测交易、执行压测,默认在 `x/benchmark/build/bin/benchmark_client`