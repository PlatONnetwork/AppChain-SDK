
# 创建App
在本章节中，我们将介绍怎么使用AppChain-SDK去构建一个简单的应用。这个应用将展示如何使用系统模块以及如何启动应用。

## 定义SimApp

SimApp必需继承`BaseApp`的能力。

```go title="simapp/app.go"
type SimApp struct {
    *baseapp.BaseApp
    // ...
}
```

## 创建数据库

除区块链底层数据库外，模块在节点本地也需要存储数据。

```go title="simapp/main.go"
	datadir := node.DefaultDataDir()
	if ctx.GlobalIsSet(utils.DataDirFlag.Name) {
		datadir = ctx.GlobalString(utils.DataDirFlag.Name)
	}
	if datadir != "" {
		absdatadir, err := filepath.Abs(datadir)
		if err != nil {
			return nil, err
		}
		datadir = absdatadir
	}

	dbfile := filepath.Join(datadir, "sdk")
	store, err := storage.NewStorage(dbfile, 256, 512, "sdk")
	if err != nil {
		log.Error("failed to new storage", "err", err)
		return nil, err
	}
```

## 加载系统模块

目前以下这些系统模块都是应用必需的，除deposit和withdraw模块外。根据应用的代币发行模式选择deposit或withdraw模块，对于在L1上发行代币的应用使用deposit模块，在L2上发行的使用withdraw模块。`SimApp`使用deposit模块。

!!! note "模块依赖关系"

    - `l1`模块存储所有模块的创世块配置信息，其他模块通过`l1`提供的接口去获取这些配置。
    - `staking`模块通过`stage`模块获取共识轮和结算轮数据，通过`vrf`模块产生随机性选举共识节点，通过`reward`模块发放区块奖励。
    - `stage` `vrf` `reward` `staking`模块之间相互依赖。
    - `statesync`模块通过`staking`模块获取验证人列表，通过`extravote`生成Merkle证明。
    - `checkpoint`模块通过`staking`模块获取验证人列表，共识轮数据，通过`txrealyer`模块向根链提交checkpoint，通过`extravote`模块生成Merkle树的根以及证明，通过`stateevent`模块生成退出事件的根和生成对应事件的Merkle证明。

```go title="simapp/app.go"
	l1Module := l1.NewModule(store)

	stateEvent := stateevent.NewModule(store)

	stageModule := stage.NewModule(ctx)
	vrfModule := vrf.NewModule(ctx, stageModule)
	stakeModule := staking.NewModule(ctx, l1Module, stageModule)
	stateSync, err := statesync.NewModule(ctx, l1Module, stakeModule, store, extravote.NewExtraVoteDB(store))
	if err != nil {
		return nil, err
	}
	rewardModule := reward.NewModule(ctx, stageModule)
	depositModule := deposit.NewModule(ctx, l1Module)
	l2StateSender := statesender.NewModule(ctx)

	vrfModule.SetStakeModule(stakeModule)
	stakeModule.SetRewardModule(rewardModule)
	stakeModule.SetVRFModule(vrfModule)
	rewardModule.SetStakeModule(stakeModule)

	rootchainRpc := ctx.GlobalString(x.RootchainNodeRPCFlag.Name)
	rootchainTxRelayer := txrelayer.NewModule(rootchainRpc, txrelayer.DefaultReceiptTimeout, txrelayer.DefaultNumRetries)

	checkpoint, err := checkpoint.NewModule(
		ctx,
		store,
		stakeModule,
		rootchainTxRelayer,
		extravote.NewExtraVoteDB(store),
		stateEvent,
		l1Module)

	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{stateSync, checkpoint})

    // 创建ModuleManager并将模块注册
	manager := module.NewManager(stateSync, stateEvent, l1Module, extraVote, rootchainTxRelayer, checkpoint, stageModule, vrfModule, stakeModule, rewardModule, depositModule, l2StateSender)
    // 设置ElectionModule接口回调的模块
	manager.SetElection(stakeModule.Name())
    // 设置ConsensusExtendModule接口回调模块
	manager.SetConsensusExtend(extraVote.Name())
    // 设置TransactionModule接口回调模块的顺序
	manager.SetOrderTransaction(stateSync.Name(), vrfModule.Name(), stakeModule.Name())
    // 设置InitModule接口回调模块的顺序
	manager.SetOrderInit(stateSync.Name(), rootchainTxRelayer.Name(), checkpoint.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
    // 设置GenesisModule接口回调模块的顺序
	manager.SetOrderGenesis(l1Module.Name(), stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name(), depositModule.Name(), l2StateSender.Name(), stateSync.Name())
    // 设置BeginBlockerModule接口回调模块的顺序
	manager.SetOrderBeginBlocker(stageModule.Name(), stakeModule.Name(), rewardModule.Name())
    // 设置EndBlockerModule接口回调模块的顺序
	manager.SetOrderEndBlocker(stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
    // 设置BlockCommiter接口回调模块的顺序
	manager.SetOrderBlockCommiter(stakeModule.Name(), stateEvent.Name(), checkpoint.Name())
```
## 启动App

有部分模块自定义了命令行参数，在启动App之前需要先调用模块的`AddModuleInitFlags`函数将自定义参数加入App的命令行参数中。

```go title="simapp/main.go"
func main() {
	cliApp := cli.NewApp()
	x.AddModuleInitFlags(cliApp)
	checkpoint.AddModuleInitFlags(cliApp)
	statesync.AddModuleInitFlags(cliApp)
	utils.AddModuleInitFlags(cliApp)
	app.InitApp(cliApp, func(ctx *cli.Context) sdk.App {
		simApp, err := NewSimApp(ctx)
		if err != nil {
			panic(fmt.Sprintf("Create simple app error: %v", err))
		}
		return simApp
	}, nil, nil)

	if err := cliApp.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

## 完整代码

```go title="simapp/main.go"
package main

import (
    "fmt"
    "github.com/PlatONnetwork/AppChain-SDK/utils"
    "github.com/PlatONnetwork/AppChain-SDK/x/statesync"
    "os"

    "github.com/PlatONnetwork/AppChain-SDK/x"
    "github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
    "github.com/PlatONnetwork/PlatON-Go/sdk"
    "github.com/PlatONnetwork/PlatON-Go/sdk/app"
    "gopkg.in/urfave/cli.v1"
)

func main() {
    cliApp := cli.NewApp()
    x.AddModuleInitFlags(cliApp)
    checkpoint.AddModuleInitFlags(cliApp)
    statesync.AddModuleInitFlags(cliApp)
    utils.AddModuleInitFlags(cliApp)
    app.InitApp(cliApp, func(ctx *cli.Context) sdk.App {
        simApp, err := NewSimApp(ctx)
        if err != nil {
            panic(fmt.Sprintf("Create simple app error: %v", err))
        }
        return simApp
    }, nil, nil)

    if err := cliApp.Run(os.Args); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

```go title="simapp/app.go"
package main

import (
	"path/filepath"

	"github.com/PlatONnetwork/AppChain-SDK/x/deposit"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesender"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf"

	"github.com/PlatONnetwork/AppChain-SDK/baseapp"
	"github.com/PlatONnetwork/AppChain-SDK/store/storage"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking"
	"github.com/PlatONnetwork/AppChain-SDK/x/stateevent"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync"
	"github.com/PlatONnetwork/AppChain-SDK/x/txrelayer"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"gopkg.in/urfave/cli.v1"
)

type SimApp struct {
	*baseapp.BaseApp
}

func NewSimApp(ctx *cli.Context) (*SimApp, error) {
	datadir := node.DefaultDataDir()
	if ctx.GlobalIsSet(utils.DataDirFlag.Name) {
		datadir = ctx.GlobalString(utils.DataDirFlag.Name)
	}
	if datadir != "" {
		absdatadir, err := filepath.Abs(datadir)
		if err != nil {
			return nil, err
		}
		datadir = absdatadir
	}

	dbfile := filepath.Join(datadir, "sdk")
	store, err := storage.NewStorage(dbfile, 256, 512, "sdk")
	if err != nil {
		log.Error("failed to new storage", "err", err)
		return nil, err
	}

	l1Module := l1.NewModule(store)

	stateEvent := stateevent.NewModule(store)

	stageModule := stage.NewModule(ctx)
	vrfModule := vrf.NewModule(ctx, stageModule)
	stakeModule := staking.NewModule(ctx, l1Module, stageModule)
	stateSync, err := statesync.NewModule(ctx, l1Module, stakeModule, store, extravote.NewExtraVoteDB(store))
	if err != nil {
		return nil, err
	}
	rewardModule := reward.NewModule(ctx, stageModule)
	depositModule := deposit.NewModule(ctx, l1Module)
	l2StateSender := statesender.NewModule(ctx)

	vrfModule.SetStakeModule(stakeModule)
	stakeModule.SetRewardModule(rewardModule)
	stakeModule.SetVRFModule(vrfModule)
	rewardModule.SetStakeModule(stakeModule)

	rootchainRpc := ctx.GlobalString(x.RootchainNodeRPCFlag.Name)
	rootchainTxRelayer := txrelayer.NewModule(rootchainRpc, txrelayer.DefaultReceiptTimeout, txrelayer.DefaultNumRetries)

	checkpoint, err := checkpoint.NewModule(
		ctx,
		store,
		stakeModule,
		rootchainTxRelayer,
		extravote.NewExtraVoteDB(store),
		stateEvent,
		l1Module)

	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{stateSync, checkpoint})

	manager := module.NewManager(stateSync, stateEvent, l1Module, extraVote, rootchainTxRelayer, checkpoint, stageModule, vrfModule, stakeModule, rewardModule, depositModule, l2StateSender)
	manager.SetElection(stakeModule.Name())
	manager.SetConsensusExtend(extraVote.Name())
	//manager.SetWorker(stateSync.Name())
	manager.SetOrderTransaction(stateSync.Name(), vrfModule.Name(), stakeModule.Name())
	manager.SetOrderInit(stateSync.Name(), rootchainTxRelayer.Name(), checkpoint.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderGenesis(l1Module.Name(), stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name(), depositModule.Name(), l2StateSender.Name(), stateSync.Name())
	manager.SetOrderBeginBlocker(stageModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderEndBlocker(stageModule.Name(), vrfModule.Name(), stakeModule.Name(), rewardModule.Name())
	manager.SetOrderBlockCommiter(stakeModule.Name(), stateEvent.Name(), checkpoint.Name())

	app := &SimApp{}
	baseApp, err := baseapp.NewBaseApp("simapp", store, manager)
	if err != nil {
		return nil, err
	}
	app.BaseApp = baseApp

	return app, nil
}
```
