package main

import (
	"context"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/customchildchainmanager"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/stakemanager"
	deploytools "github.com/PlatONnetwork/AppChain-SDK/tools/deploy"

	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/registryfactory"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/testnettoken"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
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
	DeployChildChainCommand = cli.Command{
		Name:   "deploychildchain",
		Action: deployChildChain,
		Flags: []cli.Flag{
			ChildChainIdFlag,
			StakeAmountFlag,
			StakeOwnerFlag,
			CommissionRateFlag,
			MinStakeFlag,
			MinDelegateFlag,
			ChildchainOwnerFlag,
			ChildchainTokenFlag,
			ChildchainStakeFlag,
			ChildchainDepositFlag,
			ChildchainWithdrawFlag,
			ChildchainGovernanceFlag,
			ChildchainIssuanceOptionFlag,
			DeployKeyFlag,
			GasPriceFlag,
			GasLimitFlag,
			RootChianUrlFlag,
			OutputFlag,
			TemplateResultFileFlag,
			deploytools.NodeConfigFileFlag,
			ChildchainContractFileFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func deployChildChain(ctx *cli.Context) error {
	client, err := ethclient.Dial(ctx.String(RootChianUrlFlag.Name))
	if err != nil {
		return err
	}
	chainid, err := client.ChainID(context.Background())
	if err != nil {
		return err
	}
	deployerKey, err := crypto.HexToECDSA(ctx.String(DeployKeyFlag.Name))
	if err != nil {
		return err
	}
	authOpts := createTransactionOpts(deployerKey, chainid, ctx.Uint64(GasLimitFlag.Name), big.NewInt(int64(ctx.Uint64(GasPriceFlag.Name))))
	childChainOwnerKey, err := crypto.HexToECDSA(ctx.String(ChildchainOwnerFlag.Name))
	if err != nil {
		return err
	}
	childChainOwnerAddress := crypto.PubkeyToAddress(childChainOwnerKey.PublicKey)
	childChainOwnerOpts := createTransactionOpts(childChainOwnerKey, chainid, ctx.Uint64(GasLimitFlag.Name), big.NewInt(int64(ctx.Uint64(GasPriceFlag.Name))))
	outputDir := ctx.String(OutputFlag.Name)
	templateConfig, err := readTemplateConfig(filepath.Join(outputDir, ctx.String(TemplateResultFileFlag.Name)))
	if err != nil {
		return err
	}
	childNodeConfig, err := deploytools.ReadNodeConfig(filepath.Join(outputDir, ctx.String(deploytools.NodeConfigFileFlag.Name)))
	if err != nil {
		return err
	}

	childChainId := big.NewInt(int64(ctx.Uint64(ChildChainIdFlag.Name)))
	stakeOwner := deploytools.MustStringToAddress(ctx.String(StakeOwnerFlag.Name))
	stakeAmount, _ := new(big.Int).SetString(ctx.String(StakeAmountFlag.Name), 10)
	commissionRate, _ := new(big.Int).SetString(ctx.String(CommissionRateFlag.Name), 10)
	minStake, _ := new(big.Int).SetString(ctx.String(MinStakeFlag.Name), 10)
	minDelegate, _ := new(big.Int).SetString(ctx.String(MinDelegateFlag.Name), 10)

	token := common.Address{}

	if len(ctx.String(ChildchainTokenFlag.Name)) == 0 {
		token, err = deploy(func() (common.Address, *types.Transaction, error) {
			addr, tx, _, err := testnettoken.DeployTestnetToken(authOpts, client, "simapp", "SIMAPP")
			return addr, tx, err
		}, client, "Testnettoken")
		if err != nil {
			return err
		}
	} else {
		token = deploytools.MustStringToAddress(ctx.String(ChildchainTokenFlag.Name))
	}
	childStakeHandler := deploytools.MustStringToAddress(ctx.String(ChildchainStakeFlag.Name))
	childDepositHandler := deploytools.MustStringToAddress(ctx.String(ChildchainDepositFlag.Name))
	childWithdrawManager := deploytools.MustStringToAddress(ctx.String(ChildchainWithdrawFlag.Name))
	childGovernanceManager := deploytools.MustStringToAddress(ctx.String(ChildchainGovernanceFlag.Name))
	issuanceOp := uint8(ctx.Int(ChildchainIssuanceOptionFlag.Name))

	registryFactory, err := registryfactory.NewRegistryFactory(deploytools.MustStringToAddress(templateConfig.RegistryFactoryAddr), client)

	receipt, err := send(func() (*types.Transaction, error) {
		tx, err := registryFactory.CreateChildchain(authOpts, childChainId, registryfactory.RegistryFactoryCreateOption{
			MinStake:               minStake,
			MinDelegate:            minDelegate,
			Owner:                  childChainOwnerAddress,
			Token:                  token,
			ChildStakeHandler:      childStakeHandler,
			ChildDepositHandler:    childDepositHandler,
			ChildWithdrawManager:   childWithdrawManager,
			ChildGovernanceManager: childGovernanceManager,
			IssuanceOp:             issuanceOp,
		}, genValidator(childNodeConfig.Nodes))
		return tx, err
	}, client, "CreateChildchain")
	if err != nil {
		return err
	}
	res, err := decodeCreateChildchainReceipt(receipt)
	if err != nil {
		log.Error("Decode ChildChainCreated event failed", "err", err)
	}
	genesisNodes := genesisNodes(childNodeConfig.Nodes)
	totalStakeAmount := new(big.Int).Mul(stakeAmount, big.NewInt(int64(len(genesisNodes))))
	testnetToken, _ := testnettoken.NewTestnetToken(token, client)

	if _, err := send(func() (*types.Transaction, error) {
		return testnetToken.Mint(childChainOwnerOpts, childChainOwnerAddress, totalStakeAmount)
	}, client, "Token mint"); err != nil {
		return err
	}
	if _, err := send(func() (*types.Transaction, error) {
		return testnetToken.Approve(childChainOwnerOpts, res.StakeManager, totalStakeAmount)
	}, client, "Token Approve"); err != nil {
		return err
	}
	stakeManager, _ := stakemanager.NewStakeManager(res.StakeManager, client)
	for _, n := range genesisNodes {
		if _, err := send(func() (*types.Transaction, error) {
			return stakeManager.StakeFor(childChainOwnerOpts, stakemanager.ValidatorStake{
				Owner:          stakeOwner,
				StakeAmount:    stakeAmount,
				CommissionRate: commissionRate,
				PubKey:         deploytools.MustDecodeString(n.NodePublicKey),
				BlsKey:         deploytools.MustDecodeString(n.BlsPublicKey),
			})
		}, client, "stakeFor"); err != nil {
			return err
		}
	}

	customChildChainManager, _ := customchildchainmanager.NewCustomChildChainManager(res.ChildChainManager, client)
	if _, err := send(func() (*types.Transaction, error) {
		return customChildChainManager.FinalizeGenesis(childChainOwnerOpts)
	}, client, "FinalizeGenesis"); err != nil {
		return err
	}
	if _, err := send(func() (*types.Transaction, error) {
		return customChildChainManager.EnableStaking(childChainOwnerOpts)
	}, client, "EnableStaking"); err != nil {
		return err
	}
	data, err := toml.Marshal(&ChildChainContractConfig{
		GenesisValidatorOwner: ctx.String(StakeOwnerFlag.Name),
		GenesisStakeAmount:    ctx.String(StakeAmountFlag.Name),
		GenesisCommissionRate: ctx.String(CommissionRateFlag.Name),
		DeployNumber:          receipt.BlockNumber.Uint64(),
		ChainId:               res.ChainId.Uint64(),
		StateSenderAddr:       res.StateSender.Hex(),
		CheckpointManagerAddr: res.CheckpointManager.Hex(),
		ExitHelperAddr:        res.ExitHelper.Hex(),
		ChildChainManagerAddr: res.ChildChainManager.Hex(),
		StakeManagerAddr:      res.StakeManager.Hex(),
		DepositManagerAddr:    res.DepositManager.Hex(),
		WithdrawHandlerAddr:   res.WithdrawHandler.Hex(),
		GovernanceHandlerAddr: res.GovernanceHandler.Hex(),
	})
	if err != nil {
		return err
	}
	path := filepath.Join(ctx.String(OutputFlag.Name), ctx.String(ChildchainContractFileFlag.Name))
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		log.Error("Write config failed", "path", path)
	}
	return nil
}

func genValidator(conf []*deploytools.Node) []registryfactory.ICheckpointManagerValidator {
	var vals []registryfactory.ICheckpointManagerValidator
	for _, v := range conf {
		if v.Genesis {
			pub := deploytools.MustDecodePubKey(v.NodePublicKey)
			vals = append(vals, registryfactory.ICheckpointManagerValidator{
				Address: crypto.PubkeyToAddress(*pub),
				BlsKey:  deploytools.MustDecodeString(v.BlsUncompressedPublicKey),
			})
		}
	}
	return vals
}

func decodeCreateChildchainReceipt(receipt *types.Receipt) (*registryfactory.RegistryFactoryChildChainCreated, error) {
	event := new(registryfactory.RegistryFactoryChildChainCreated)
	factoryAbi, _ := registryfactory.RegistryFactoryMetaData.GetAbi()
	for _, log := range receipt.Logs {
		if log.Topics[0] == factoryAbi.Events["ChildChainCreated"].ID {
			if len(log.Data) > 0 {
				if err := factoryAbi.UnpackIntoInterface(event, "ChildChainCreated", log.Data); err != nil {
					return nil, err
				}
			}
			var indexed abi.Arguments
			for _, arg := range factoryAbi.Events["ChildChainCreated"].Inputs {
				if arg.Indexed {
					indexed = append(indexed, arg)
				}
			}
			if err := abi.ParseTopics(event, indexed, log.Topics[1:]); err != nil {
				return nil, err
			}
		}
	}
	return event, nil
}
func genesisNodes(nodes []*deploytools.Node) []*deploytools.Node {
	var genesis []*deploytools.Node
	for _, n := range nodes {
		if n.Genesis {
			genesis = append(genesis, n)
		}
	}
	return genesis
}
