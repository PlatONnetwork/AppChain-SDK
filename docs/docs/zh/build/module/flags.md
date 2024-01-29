# Custom Flags

模块除了通过创世配置配置外，还可通过命令行参数配置。AppChain-SDK提供了一个统一的方法`AddModuleInitFlags`，模块可实现此方法，以便在应用启动时将模块的命令行参数增加到应用的命令行参数列表中。

## AddModuleInitFlags

!!! info "AddModuleInitFlags"
    `AddModuleInitFlags(app cli.App)`

    参数：

    - **app** App 对象。

使用示例：

- 定义命令行参数

```go title="x/checkpoint/flags.go"
package checkpoint

import "gopkg.in/urfave/cli.v1"

var (
	KeystoreFlag = cli.StringFlag{
		Name:  "checkpoint.keystore",
		Usage: "Keystore for signing checkpoint transaction",
	}
	PasswordFlag = cli.StringFlag{
		Name:  "checkpoint.password",
		Usage: "Password for keystore",
		EnvVar: "CHECKPOINT_PASSWORD",
	}
)
```

- 实现`AddModuleInitFlags`方法

```go title="x/checkpoint/module.go"
func AddModuleInitFlags(app *cli.App) {
	app.Flags = append(app.Flags, KeystoreFlag)
	app.Flags = append(app.Flags, PasswordFlag)
}
```

- 应用启动时调用`AddModuleInitFlags`

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
