# 脚本说明

* deploy 文件夹为生成节点staking脚本工具
* tools 为节点脚本相关模板文件

## 使用说明

* 确认已经使用 simapp_client 工具生成节点相关脚本
* 进入 deploy 文件夹
* 执行 `./deploy.sh [output] [owner] [stakeAmount] [key] [user] [rootchainurl] [childchainurl] [toolsPath] [targetAddr]`
    * output 为 simapp_client 相关文件输出的文件夹
    * owner 质押参数中验证人 owner 参数，owner 通常为 user保持一致
    * stakeAmount 质押参数中 质押token数
    * key 用户发送交易的私钥，建议每个节点单独使用一个私钥
    * key 私钥对应的地址
    * rootchainurl layer1 rpc地址
    * childchainurl 通常为 http://127.0.0.1:8801
    * toolsPath tools 工具地址，需要将其拷贝到节点脚本执行目录
    * targetAddr 验证人地址，根据验证人地址找到验证节点目录来生成脚本
* 为 user 指定的地址转账，三种类型
    * layer1 coin 确保 user能发送交易
    * layer1 token 确保质押有足够的token进行质押
    * layer2 coin 确保有足够的代币发送解质押操作