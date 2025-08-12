package main

import (
	"context"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/customchildchainmanager"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/depositmanager"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/exithelper"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/governancehandler"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/logicproxy"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/registryfactory"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/registrymanager"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/stakemanager"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/withdrawhandler"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"os"
	"path/filepath"
)

var (
	DeployTemplateCommand = cli.Command{
		Name:   "deploytemplate",
		Action: deployTemplate,
		Flags: []cli.Flag{
			DeployKeyFlag,
			InitializeKeyFlag,
			RegistryManagerOwnerFlag,
			RootChianUrlFlag,
			GasLimitFlag,
			GasPriceFlag,
			TemplateResultFileFlag,
			OutputFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func deployTemplate(ctx *cli.Context) error {
	client, err := ethclient.Dial(ctx.String(RootChianUrlFlag.Name))
	if err != nil {
		return err
	}
	deployerKey, err := crypto.HexToECDSA(ctx.String(DeployKeyFlag.Name))
	if err != nil {
		return err
	}
	initializeKey, err := crypto.HexToECDSA(ctx.String(InitializeKeyFlag.Name))
	if err != nil {
		return err
	}
	chainid, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	owner, err := common.StringToAddress(ctx.String(RegistryManagerOwnerFlag.Name))
	if err != nil {
		log.Error("Decode owner failed", "err", err)
		return err
	}
	authOpts := createTransactionOpts(deployerKey, chainid, ctx.Uint64(GasLimitFlag.Name), big.NewInt(int64(ctx.Uint64(GasPriceFlag.Name))))
	initializeOpts := createTransactionOpts(initializeKey, chainid, ctx.Uint64(GasLimitFlag.Name), big.NewInt(int64(ctx.Uint64(GasPriceFlag.Name))))
	var registryManagerAddr common.Address
	var registryFactoryAddr common.Address
	var logicProxyAddr common.Address
	if registryManagerAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := registrymanager.DeployRegistryManager(authOpts, client)
		return addr, tx, err
	}, client, "registry manager"); err != nil {
		return err
	}

	if logicProxyAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := logicproxy.DeployLogicProxy(authOpts, client, registryManagerAddr)
		return addr, tx, err
	}, client, "LogicProxy"); err != nil {
		return err
	}
	if registryFactoryAddr, err = deployRegistryFactoryComponent(authOpts, client, logicProxyAddr); err != nil {
		return err
	}

	registryManager, _ := registrymanager.NewRegistryManager(logicProxyAddr, client)
	if _, err = send(func() (*types.Transaction, error) {
		return registryManager.Initialize(initializeOpts, owner, registryFactoryAddr)
	}, client, "Initialize"); err != nil {
		return err
	}
	data, err := toml.Marshal(&TemplateConfig{
		RegistryManagerImplementAddr: registryManagerAddr.Hex(),
		RegistryManagerProxyAddr:     logicProxyAddr.Hex(),
		RegistryFactoryAddr:          registryFactoryAddr.Hex(),
	})
	if err != nil {
		return err
	}
	path := filepath.Join(ctx.String(OutputFlag.Name), ctx.String(TemplateResultFileFlag.Name))
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		log.Error("Write config failed", "path", path)
	}
	return err
}

func deployRegistryFactoryComponent(authOpts *bind.TransactOpts, client *ethclient.Client, registryManagerAddr common.Address) (common.Address, error) {
	var registryFactoryAddr common.Address
	var exitHelperAddr common.Address
	var customChildChainManagerAddr common.Address
	var stakeManagerAddr common.Address
	var depositManagerAddr common.Address
	var withdrawHandlerAddr common.Address
	var governanceHandlerAddr common.Address
	var err error

	if exitHelperAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := exithelper.DeployExitHelper(authOpts, client)
		return addr, tx, err
	}, client, "ExitHelper"); err != nil {
		return common.Address{}, err
	}
	if customChildChainManagerAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := customchildchainmanager.DeployCustomChildChainManager(authOpts, client)
		return addr, tx, err
	}, client, "CustomChildChainManager"); err != nil {
		return common.Address{}, err
	}
	if stakeManagerAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := stakemanager.DeployStakeManager(authOpts, client)
		return addr, tx, err
	}, client, "StakeManager"); err != nil {
		return common.Address{}, err
	}
	if depositManagerAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := depositmanager.DeployDepositManager(authOpts, client)
		return addr, tx, err
	}, client, "DepositManager"); err != nil {
		return common.Address{}, err
	}
	if withdrawHandlerAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := withdrawhandler.DeployWithdrawHandler(authOpts, client)
		return addr, tx, err
	}, client, "WithdrawHandler"); err != nil {
		return common.Address{}, err
	}

	if governanceHandlerAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := governancehandler.DeployGovernanceHandler(authOpts, client)
		return addr, tx, err
	}, client, "GovernanceHandler"); err != nil {
		return common.Address{}, err
	}
	if registryFactoryAddr, err = deploy(func() (common.Address, *types.Transaction, error) {
		addr, tx, _, err := registryfactory.DeployRegistryFactory(authOpts, client, registryfactory.RegistryFactoryInitializeParams{
			RegistryManager:   registryManagerAddr,
			ExitHelper:        exitHelperAddr,
			ChildChainManager: customChildChainManagerAddr,
			StakeManager:      stakeManagerAddr,
			DepositManager:    depositManagerAddr,
			WithdrawHandler:   withdrawHandlerAddr,
			GovernanceHandler: governanceHandlerAddr,
		})
		return addr, tx, err
	}, client, "RegistryFactory"); err != nil {
		return common.Address{}, err
	}
	return registryFactoryAddr, nil
}
