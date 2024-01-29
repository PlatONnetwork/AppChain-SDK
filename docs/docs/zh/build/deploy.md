# 部署链

本章节将讲解如何部署应用链，本章节采取的部署 4 个应用链节点，其中 3 个创世节点，1 个候选节点方案。为了简化创建过程，我们使用 `AppChain-SDK/simapp/docker/quickstart/cluster` 作为 L2 节点的部署环境，`AppChain-Contract/docker/quickstart/` 作为 L1 合约的部署环境

## 前期准备

### 账号准备

#### L1 

* 创建合约账号（本文档使用 `AppChain-Contract/62d63dd4-bfb3-451e-b4d7-b9c75b02a894）
* 合约初始化账号 （本文档使用 `AppChain-Contract/75a8f9b1-73cc-4d57-952d-673efe1efee2`)
* 创世验证人账号（包括 验证人的账户地址、BLS 账户地址、P2P 地址公钥，本文档私钥放在 `AppChain-SDK/simapp/docker/quickstart/cluster/node1~3`）

```
| file                                 | address                                    | bench32                                     |
|--------------------------------------|--------------------------------------------|---------------------------------------------|
| 75a8f9b1-73cc-4d57-952d-673efe1efee2 | 0x62953f9213f899f2A51680C2fBB4282a2591bfc8 | lat1v22nlysnlzvl9fgksrp0hdpg9gjer07gkkhvu3 |
| 62d63dd4-bfb3-451e-b4d7-b9c75b02a894 | 0x4779472533Ed1B8Fe6d6E0d212F2aA02Ffd84b9F | lat1gau5wffna5dclekkurfp9u42qtlasjulxwwhyu  |
```
### 链信息
* 链 ID
* 最小质押金额
* 最小委托金额
* L2 Stake 合约地址（默认 0x1000000000000000000000000000000000000005）
* L2 Deposit 合约地址 （默认 0x1000000000000000000000000000000000000007）
* 代币发行选项，目前只支持 L1 ERC20 发行（1：L1 发行，2:L2 发行）
* Giskard 验证节点出块数，出块间隔
### 安装包

#### L2 应用程序

编译 AppChain-SDK 中的 simapp

```shell
go build
```

####  交易发送工具

编译 AppChain-SDK/tools

```shell
go build
```

#### 合约

使用 AppChain-Contracts 仓库

```shell
forge build
```

#### L1 链

我们使用 `AppChain-SDK/simapp/docker-composer.yml``

```shell
docker-composer build
```

## 部署 L1

```shell
docker-composer start rootchain
```

RPC 地址 `http://127.0.0.1:8800`

## 部署合约

* 编辑 配置 文件

我们使用 AppChain-Contracts 部署脚本（仅在测试环境使用）

将 `AppChain-Contracts/docker/中的` `.env` `testnetQuickStartGenesisValidatorKeys.json` 拷贝到根目录

```shell
cd AppChain-Contracts
cp docker/.env .
cp docker/testnetQuickStartGenesisValidatorKeys.json .
```

修改 `.env` 文件
```yaml
BLOCK_CHAIN_NET_RPC_URL=http://127.0.0.1:8800
```

* 部署合约模板

```shell
./deploy_templete.sh 62d63dd4-bfb3-451e-b4d7-b9c75b02a894 password
```
部署结果 .deploy_templete_result 文件

```shell
RegistryFactory=0x434cd2795FF145D9f1F8D8429347CdE507CD6226
RegistryManager=0xE10a672d155A3A40CB6066f7adb2F2f182953F2F
ExitHelper=0x7395a22a078db2CFe7C1115efD4f31A703d2516e
CustomChildChainManager=0x9fC1B347Ef69f840830cE5A7Aa3EB5dAEF9ba282
StakeManager=0x8889f92b036851051D4eA3C45bB0F3cFaBeCF08a
DepositManager=0xC2b280bA48485448b709b7e0b5e88bc40e8D7e9d
WithdrawHandler=0xC8bE854caf35b72503D732f00A61DA4996ab3970
Deployer=0x4779472533Ed1B8Fe6d6E0d212F2aA02Ffd84b9F
```

* 编辑 .env 文件

```shell
BLOCK_CHAIN_NET_RPC_URL=http://127.0.0.1:8800

## create testnet child chian options:
## 写入上一步部署模板合约地址
REGISTRY_FACTORY=0x434cd2795FF145D9f1F8D8429347CdE507CD6226
CHILDCHAIN_ID=123083
MIN_STAKE=1000
MIN_DELEGATE=10
CHILD_STAKE_HANDLER=0x1000000000000000000000000000000000000005
CHILD_DEPOSIT_HANDLER=0x1000000000000000000000000000000000000007
CHILD_WITHDRAW_MANAGER=0x1000000000000000000000000000000000000008
ISSUANCE_OPTION=1

## genesis validator key options (for checkpoint manager and stake manager):
GENESIS_VALIDATORS_KEY_JSON_FILE="/testnetQuickStartGenesisValidatorKeys.json"

## quick stake validator options:
STAKE_OWNER=0xf22c7e0702483b876d22ebC6Ac0542C7ffF9A4Eb
STAKE_AMOUNT=1000000000
COMMISSION_RATE=0
```

* 部署 L2 链合约

```shell
./testnet_quickstart.sh 62d63dd4-bfb3-451e-b4d7-b9c75b02a894 password  75a8f9b1-73cc-4d57-952d-673efe1efee2 password
```

部署结果保存在 .quickstart_testnet_childchain_result 文件中

```shell
RegistryFactory=0x434cd2795FF145D9f1F8D8429347CdE507CD6226
ChildchainId=123083
StateSender=0x92b68E65Ef272afA1B8F001501D17F13B836E735
CheckpointManager=0x0bb7AC917fA06100cDfF38641E6c2C6Bae4F4890
ExitHelper=0x7B66bf9dAba52733231e32D3e2fe1008205c1458
CustomChildChainManager=0x8BbfE720aEbC36EE7B2F54866fbED59f94360fDD
StakeManager=0xB04B468EB388C613cb453A60f6d610A009E1Ff55
DepositManager=0x628f6017E774F8Fc0fBf59C1ea18CB4e398c9168
WithdrawHandler=0x0000000000000000000000000000000000000000
Token=0x88edBfC0d17BFa9ffC064244Bc90Bf0F264D0880
Creater=0x4779472533Ed1B8Fe6d6E0d212F2aA02Ffd84b9F
Owner=0x62953f9213f899f2A51680C2fBB4282a2591bfc8

```

## 部署 L2
我们使用 AppChain-SDK/simapp/docker/quickstart/cluster, node1, node2, node3, node4 作为后面测试的 4 个节点

* 编辑创世文件
可以直接使用 AppChain-SDK/simapp/docker/quickstart/cluster/genesis.json，下面对参数简要说明，我们需要更改 initialNodes 中 IP 地址

```shell
{
    "config":{
        "chainId":123083,
	    "pauliBlock":1,
        "cbft":{
            //创世节点
            "initialNodes":[
                {
                "node": "enode://5c79bf8b836bdc85fe513a64a558291e96ac2405b6b95d5ca05bb20db9c0a1a00e12d11dd540d8685267015ce71cef43781a75157ec6f266cc88a8fb8b5c6c17@127.0.0.1:18001",
                "blsPubKey": "b6f85c577ff890f9737595a9e326d5539cc1cc859c879017cf25accef8a96518ec9747ef621625063dc17ae95300a74f"
                },
                {
                "node": "enode://8d84e41f83e833f622c45766e7e425cf03a225867facb05baf90eaf29c1bf53680988dab4e6f058759871c91b3e6ff888aabc41daec4b8e0877fc8fe8feed27f@127.0.0.1:18002",
                "blsPubKey": "828b858eab99526f901ce610ae3d2d9b08f2302e56a29698776d2135e94fc074c575a53e27b8a68e9a7692f3be65965a"
                },
                {
                "node": "enode://c77d41c1feebf5ad53b4cd3cd52d55659a850485b68825012e0573025799a6b818787ef3a4951861a0aae7345222f6d6b3af73974fd019187d5014c12c3ac9f7@127.0.0.1:18003",
                "blsPubKey": "a33240d0c07e9a1abf747dfe4e87e54d25b2314bc33b8a700357e54a1fd4ced34f18ffe6ddc97e67a4f23655ecc38915"
                }
            ],
            "amount":10, //验证人出块数
            "period":10000, // 出块间隔
        },
        "genesisVersion":65792,
        "addressHRP":"lat"
    },
    "nonce":"0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23",
    "timestamp":"0x17538ac5720",
    "extraData":"",
    "gasLimit":"201600000",
    "alloc":{ // 内置账户
        "lat1v22nlysnlzvl9fgksrp0hdpg9gjer07gkkhvu3":{
            "balance":"2000000000000000000000000000000000000"
        },
        "lat1wrfq0sfj9n9eq6wn0yxkw6yxdk4l7yp4hpmt8p":{
            "balance":"2000000000000000000000000000000000000"
        },
        "lat1xsuh90mr6yrzwcd242y36f6s7q7tfvhh5844xy":{
            "balance":"2000000000000000000000000000000000000"
        },
        "lat1j74n6nml2pglzfasa8udzpmjzfwef4jm7y5c5a":{
            "balance":"2000000000000000000000000000000000000"
        },
        "lat15apf4czd3wufel2h9936m65vh7ppjcy6tcts33":{
            "balance":"2000000000000000000000000000000000000"
        },
        "lat1jpklm0ndy0pw76vhkg9uqnn5kcl0s9398rsfp4":{
            "balance":"2000000000000000000000000000000000000"
        },
        "lat1t9xn2zjdxjm3rcvsmaxku5mtvz7xtzg3uwzueu":{
            "balance":"2000000000000000000000000000000000000"
        },
        "lat165pu38wg9cx5x6ygvr5mmul08mlq0khzeamh5v":{
            "balance":"2000000000000000000000000000000000000"
        }
    },
    "modules":{
	"l1": {//L1 部署合约相关信息
	  "chainId": 2203181,
          "stateSender": "lat1j2mgue00yu405xu0qq2sr5tlzwurdee4l8gkff",
          "checkpoint": "lat1pwm6eytl5psspn0l8pjpumpvdwhy7jysg335uv",
          "stakeManager": "lat13wl7wg9whsmwu7e02jrxl0k4n72rvr7ag9c40n",
          "depositManager": "lat1v28kq9l8wnu0cralt8q75xxtfcuceytg3z68qs"
	},
	"stage":{//BFT 选举相关参数
	  "roundValidatorElectionDistance": 20, //共识
	  "roundSize":250,
	  "epochSize": 500
	},
	"vrf": {//VRF 创世种子
	  "genesisVRFNonce":"0x000000000000000000000000000000000067656e657369735652464e6f6e6365"
	}, 
	"staking":{//staking 相关参数
	    "genesisStakeAmount": 1000000000, // 创世每个验证人质押金额，与 L1 minStake 保持一致，
	    "genesisValidatorOwner": "lat1wrfq0sfj9n9eq6wn0yxkw6yxdk4l7yp4hpmt8p", // L1 创世的 ValidatorStake Owner 一致
	    "genesisCommissionRate": 100, // 与 ValidatorStake commissionRate 一致
	    "stakeWithdrawalWaitPeriod": 6,// 解质押提款等待期， stage.epochSize * period.
	    "delegateWithdrawalWaitPeriod": 6, // 委托提款等待期， stage.epochSize * period.
	    "slashingPercentage": 50, // 惩罚质押金额的百分比
	    "slashIncentivePercentage": 30, // 举报人奖励， 惩罚金额的百分比作为举报人奖励
	    "maxRoundValidatorsSize": 4, // 每轮共识最大验证人数
	    "maxEpochValidatorsSize": 201, // 每个结算周期最大验证人数
	    "minRoundValidatorBlockNumber": 1 // 每个共识周期最小出块数
	},
	"reward": {//奖励
	    "rewardPerBlock": 4000000000000000000, //每出块奖励
	    "rewardPerEpoch": 60000000000000000000000 //每个结算周期奖励
	}
    },
    "number":"0x0",
    "gasUsed":"0x0",
    "parentHash":"0x0000000000000000000000000000000000000000000000000000000000000000"
}

```

* 初始化

```shell
./init.sh 1
./init.sh 2
./init.sh 3
```

* 启动

更改 start.sh 脚本中 ROOTCHAIN_RPC

```shell
ROOTCHAIN_RPC=http://127.0.0.1:8800
```
启动节点

```shell
./start.sh 1 &
./start.sh 2 &
./start.sh 3 &
```

## 质押验证人

铸币

```shell
./tool cast erc20 --rpc http://127.0.0.1:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x88edBfC0d17BFa9ffC064244Bc90Bf0F264D0880 --type send --method mint 0x62953f9213f899f2A51680C2fBB4282a2591bfc8 20000000000
```


代币授权


```shell
./tool cast erc20 --rpc http://127.0.0.1:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x88edBfC0d17BFa9ffC064244Bc90Bf0F264D0880 --type send --method approve 0xB04B468EB388C613cb453A60f6d610A009E1Ff55 200000000000000
```

我们将 cluster 中 node4 节点在 L1 合约进行质押

```shell
./tool cast l1.staking $l1rpc --key $from_private_key --address $stakemanager --type send --method stakeFor $method_args
./tool cast l1.staking --rpc http://127.0.0.1:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0xB04B468EB388C613cb453A60f6d610A009E1Ff55 --type send --method stakeFor '{"owner":"0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b","stakeAmount":1000000000,"commissionRate":80,"pubKey":"0x7c278b7e4320528bdbf5c02db611282b431dbaec2509b1fbebe3de4be7442a4114f1c9970ffd0c589afa8ef6262692efd4edae5d77c87641f21f44b1d082c348","blsKey":"0xb969678ef2cf458b49b8c568d95e63221efe0f30383a9b0c5eb683bf2e23d118664631ce992d81ec4b6127ec0a760f86"}'
```

在 L2 通过 tool 查看添加情况，先查看当前l2 验证人round 和 epoch


```
./tool cast l2.validator --rpc http://127.0.0.1:8801 --round 250 --epoch 500
round: 4 epoch 2
```

* 根据round，epoch 进行查询，看是否在epoch+2个验证人列表里面

```shell
./tool cast l2.staking $l2rpc --address $l2stake --type call --method getValidatorAddrs $method_args
./tool cast l2.staking --rpc http://127.0.0.1:8801 --address 0x1000000000000000000000000000000000000005 --type call --method getValidatorAddrs 2 4

[
  [
   "0x1dD26DfB60B996FD5D5152aF723949971D9119eE",
   "0x70d207C1322CCB9069D3790D6768866DaBFf1035",
   "0x343972bF63D1062761aaaA891D2750f03cB4b2f7",
   "0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b" // 新加入的验证人
  ]
 ]

```

* 启动 node4 节点

```shell
./start.sh 4
```

* 连入到创世节点

```shell
./simapp attach http://127.0.0.1:8804
> admin.addPeer("enode://5c79bf8b836bdc85fe513a64a558291e96ac2405b6b95d5ca05bb20db9c0a1a00e12d11dd540d8685267015ce71cef43781a75157ec6f266cc88a8fb8b5c6c17@127.0.0.1:18001")

```

在下一个 epoch node4 可能会被选中

## 解质押验证人

* 发送解质押交易
```shell
./tool cast l2.staking $l2rpc --address $l2stake --key $from_private_key --type send --method getValidatorAddrs $method_args
./tool cast l2.staking --rpc http://127.0.0.1:8801 --key 8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd --address 0x1000000000000000000000000000000000000005 --type send --method unstake 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b 10000
```

* 查询解质押交易
  
```shell
./tool cast l2.staking --rpc http://127.0.0.1:8801 --key 8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd --address 0x1000000000000000000000000000000000000005 --type call --method pendingWithdrawalsOfStake 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b
[
  10000
 ]
```

* 发送质押提取交易

```shell
./tool cast l2.staking --rpc http://127.0.0.1:8801 --key 8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd --address 0x1000000000000000000000000000000000000005 --type send --method withdrawUnstake 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b
```

* 查询事件 ID

```shell
./tool cast l2.statesender --rpc http://127.0.0.1:8801 --address 0x1000000000000000000000000000000000000001 --type logs --method L2StateSynced  "[]" "[]" "[]"
[
  {
   "Id": 1,
   "Sender": "0x1000000000000000000000000000000000000005",
   "Receiver": "0x8BbfE720aEbC36EE7B2F54866fbED59f94360fDD",
   "CallData": "jKmpXkG17s4lPJP1sx7tElOu1rFF2KbhTZE/345zIpMAAAAAAAAAAAAAAACXqz1Pf1BR8Sew6fjRB3ISXZTWWwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPo",
   "Raw": {
    "address": "0x1000000000000000000000000000000000000001",
    "topics": [
     "0xedaf3c471ebd67d60c29efe34b639ede7d6a1d92eaeb3f503e784971e67118a5",
     "0x0000000000000000000000000000000000000000000000000000000000000001",
     "0x0000000000000000000000001000000000000000000000000000000000000005",
     "0x0000000000000000000000008bbfe720aebc36ee7b2f54866fbed59f94360fdd"
    ],
    "data": "0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000608ca9a95e41b5eece253c93f5b31eed1253aed6b145d8a6e14d913fdf8e73229300000000000000000000000097ab3d4f7f5051f127b0e9f8d10772125d94d65b00000000000000000000000000000000000000000000000000000000000003e8",
    "blockNumber": "0x1da3",
    "transactionHash": "0xf18f3f088a65bb8b4b2137f7094cd4c78b5e3a77bc667adb72c3714df5e95fc0",
    "transactionIndex": "0x1",
    "blockHash": "0x5b09ce32a2ee32ffaf3007b8a919359a03f6b7d97cbbf235ad0719c7a3483047",
    "logIndex": "0x2",
    "removed": false
   }
  }
 ]
```

* 查询 checkpoint

```shell
curl http://127.0.0.1:8801 -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"checkpoint_generateExitProof","params":[1],"id":0}'
{"jsonrpc":"2.0","id":0,"result":{"Data":["0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"],"Metadata":{"CheckpointBlock":7750,"ExitEvent":"000000000000000000000000000000000000000000000000000000000000000100000000000000000000000010000000000000000000000000000000000000050000000000000000000000008bbfe720aebc36ee7b2f54866fbed59f94360fdd000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000608ca9a95e41b5eece253c93f5b31eed1253aed6b145d8a6e14d913fdf8e73229300000000000000000000000097ab3d4f7f5051f127b0e9f8d10772125d94d65b00000000000000000000000000000000000000000000000000000000000003e8","LeafIndex":1}}}

```

* 发送 L1 解质押交易

```shell
./tool cast l1.exithelper --appchain-rpc http://127.0.0.1:8801 --rpc http://127.0.0.1:8800 --key 8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd --exit-id 1 --address 0x7B66bf9dAba52733231e32D3e2fe1008205c1458
```

* 发送提取交易

```shell
./tool cast l1.staking --rpc http://127.0.0.1:8800 --key 8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd --address 0xB04B468EB388C613cb453A60f6d610A009E1Ff55 --type send --method withdrawStake 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b 1000
```

## 委托

* 发送L1委托交易

```shell
./tool cast l1.staking --rpc http://127.0.0.1:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0xB04B468EB388C613cb453A60f6d610A009E1Ff55 --type send --method delegateFor 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b 0x62953f9213f899f2A51680C2fBB4282a2591bfc8 10000
```

* L2 查询委托

```shell
./tool cast l2.staking --rpc http://127.0.0.1:8801  --address 0x1000000000000000000000000000000000000005 --type call --method getDelegationsWithValidator '["0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b"]' 0x62953f9213f899f2A51680C2fBB4282a2591bfc8
[
  [
   {
    "ValidatorAddr": "0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b",
    "DelegatorAddr": "0x62953f9213f899f2A51680C2fBB4282a2591bfc8",
    "Amount": 10000,
    "StakeEpoch": 5,
    "DelegateEpoch": 17
   }
  ]
 ]

```

## 解委托

* 发送解质押交易
```shell
./tool cast l2.staking $l2rpc --address $l2stake --key $from_private_key --type send --method getValidatorAddrs $method_args
./tool cast l2.staking --rpc http://127.0.0.1:8801 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x1000000000000000000000000000000000000005 --type send --method undelegate 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b 10000
```

* 查询可提取的解委托

```shell
./tool cast l2.staking --rpc http://127.0.0.1:8801 --address 0x1000000000000000000000000000000000000005 --type call --method pendingWithdrawalsOfDelegate 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b 0x62953f9213f899f2A51680C2fBB4282a2591bfc8
[
  10000
 ]

```

* 提取委托

```shell
./tool cast l2.staking --rpc http://127.0.0.1:8801 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0x1000000000000000000000000000000000000005 --type send --method withdrawUndelegate 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b
```


* 查询事件 ID

```shell
./tool cast l2.statesender --rpc http://127.0.0.1:8801 --address 0x1000000000000000000000000000000000000001 --type logs --method L2StateSynced  "[]" "[]" "[]"
[
  {
   "Id": 3,
   "Sender": "0x1000000000000000000000000000000000000005",
   "Receiver": "0x8BbfE720aEbC36EE7B2F54866fbED59f94360fDD",
   "CallData": "WOWAyhzb5RjyfYV4c7YV6Aejw5VYSpPXXoDJIcmR5Q8AAAAAAAAAAAAAAACXqz1Pf1BR8Sew6fjRB3ISXZTWWwAAAAAAAAAAAAAAAGKVP5IT+JnypRaAwvu0KColkb/IAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAJxA=",
   "Raw": {
    "address": "0x1000000000000000000000000000000000000001",
    "topics": [
     "0xedaf3c471ebd67d60c29efe34b639ede7d6a1d92eaeb3f503e784971e67118a5",
     "0x0000000000000000000000000000000000000000000000000000000000000003",
     "0x0000000000000000000000001000000000000000000000000000000000000005",
     "0x0000000000000000000000008bbfe720aebc36ee7b2f54866fbed59f94360fdd"
    ],
    "data": "0x0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000008058e580ca1cdbe518f27d857873b615e807a3c395584a93d75e80c921c991e50f00000000000000000000000097ab3d4f7f5051f127b0e9f8d10772125d94d65b00000000000000000000000062953f9213f899f2a51680c2fbb4282a2591bfc80000000000000000000000000000000000000000000000000000000000002710",
    "blockNumber": "0x3001",
    "transactionHash": "0x601b26b6ea5ada0abb9ded05201e3f24a1a76152935628259b7ac1bb2f196c38",
    "transactionIndex": "0x1",
    "blockHash": "0x3842c99de17fb7dd64f73004db51f190b2803aad10d9b0812ebc8755a7ee1a34",
    "logIndex": "0x2",
    "removed": false
   }
  }
 ]

```

* 查询 checkpoint

```shell
curl http://127.0.0.1:8801 -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"checkpoint_generateExitProof","params":[3],"id":0}'
{"jsonrpc":"2.0","id":0,"result":{"Data":["0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"],"Metadata":{"CheckpointBlock":12500,"ExitEvent":"000000000000000000000000000000000000000000000000000000000000000300000000000000000000000010000000000000000000000000000000000000050000000000000000000000008bbfe720aebc36ee7b2f54866fbed59f94360fdd0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000008058e580ca1cdbe518f27d857873b615e807a3c395584a93d75e80c921c991e50f00000000000000000000000097ab3d4f7f5051f127b0e9f8d10772125d94d65b00000000000000000000000062953f9213f899f2a51680c2fbb4282a2591bfc80000000000000000000000000000000000000000000000000000000000002710","LeafIndex":1}}}

```

* 发送 L1 解质押交易

```shell
./tool cast l1.exithelper --appchain-rpc http://127.0.0.1:8801 --rpc http://127.0.0.1:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --exit-id 3 --address 0x7B66bf9dAba52733231e32D3e2fe1008205c1458
```

* 发送提取交易

```shell
./tool cast l1.staking --rpc http://127.0.0.1:8800 --key 99c18ee459838a2935b014657c00c68dc35ae9516019e67d2746f8d9ff43e1be --address 0xB04B468EB388C613cb453A60f6d610A009E1Ff55 --type send --method withdrawDelegation 0x97AB3d4f7F5051F127b0e9f8d10772125D94D65b 0x62953f9213f899f2A51680C2fBB4282a2591bfc8 1000
```