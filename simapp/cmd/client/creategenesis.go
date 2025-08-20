package main

import (
	"encoding/json"
	deploytools "github.com/PlatONnetwork/AppChain-SDK/tools/deploy"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/asyncblock"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	"github.com/PlatONnetwork/AppChain-SDK/x/blocktime"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint"
	"github.com/PlatONnetwork/AppChain-SDK/x/consensus"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/deposit"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/gov"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	"github.com/PlatONnetwork/AppChain-SDK/x/nontxpool"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward"
	config4 "github.com/PlatONnetwork/AppChain-SDK/x/reward/config"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/config"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking"
	config3 "github.com/PlatONnetwork/AppChain-SDK/x/staking/config"
	"github.com/PlatONnetwork/AppChain-SDK/x/stateevent"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesender"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync"
	"github.com/PlatONnetwork/AppChain-SDK/x/txrelayer"
	"github.com/PlatONnetwork/AppChain-SDK/x/votetoken"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf"
	config2 "github.com/PlatONnetwork/AppChain-SDK/x/vrf/config"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"net"
	"os"
	"path/filepath"
)

var (
	CreateGenesisCommand = cli.Command{
		Name:   "creategenesis",
		Action: createGenesis,
		Flags: []cli.Flag{
			OutputFlag,
			RootChainIdFlag,
			deploytools.NodeConfigFileFlag,
			ChildchainContractFileFlag,
			deploytools.GenesisFileFlag,
			GenesisCbftPeriodFlag,
			GenesisCbftAmountFlag,
			GenesisStageRoundValidatorElectionDistanceFlag,
			GenesisStageRoundSizeFlag,
			GenesisStageEpochSizeFlag,
			GenesisVrfNonceFlag,
			GenesisStakeWithdrawalWaitPeriodFlag,
			GenesisStakeDelegateWithdrawalWaitPeriodFlag,
			GenesisStakeSlashingPercentageFlag,
			GenesisStakeSlashIncentivePercentageFlag,
			GenesisStakeMaxRoundValidatorsSizeFlag,
			GenesisStakeMaxEpochValidatorsSizeFlag,
			GenesisStakeMinBlocksOfRoundValidatorFlag,
			GenesisRewardPerBlockFlag,
			GenesisRewardPerEpochFlag,
			GenesisVoteTokenNameFlag,
			GenesisVoteTokenSymbolFlag,
			GenesisVoteTokenVersionFlag,
			GenesisVoteTokenOwnerFlag,
			GenesisGovNameFlag,
			GenesisGovVersionFlag,
			GenesisGovVoteDelayFlag,
			GenesisGovVotePeriodFlag,
			GenesisGovQuorumNumeratorFlag,
			GenesisGovProposalThresholdFlag,
			GenesisGovOwnerFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	DefaultGenesisTemplate = core.Genesis{
		Config: &params.ChainConfig{
			AddressHRP: "lat",
			EmptyBlock: "",
			Cbft: &params.CbftConfig{
				Period:        10000,
				Amount:        10,
				InitialNodes:  nil,
				ValidatorMode: "",
			},
			GenesisVersion: 65792,
			Modules:        nil,
		},
		Nonce:      hexutil.MustDecode("0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23"),
		Timestamp:  1602973620000,
		ExtraData:  hexutil.MustDecode("0x22cf84e1bf86cf8220e1bc84cebdcf8920e1bd81ceb4cebfe1bfa620e1bc80ceb5e1bdb620e1bc91cebecf8ccebcceb5ceb8ceb120cebaceb1e1bdb620ceb4ceb9cebaceb1ceb9cebfcf83cf8dcebdceb7cebd20cebcceb5cf84e1bdb020cf86cf81cebfcebdceaecf83ceb5cf89cf8220cf80ceb1cebdcf84e1bdb620cf84cf81cf8ccf80e1bfb320e1bc90cf80ceb9cf84ceb7ceb4ceb5cf8dcf83cebfcebcceb5cebd2e220a225765207368616c6c20616c77617973206b65657020746f2074686520757070657220726f616420616e64207072616374696365206a75737469636520776974682070727564656e636520696e206576657279207761792e220a22e8aea9e68891e4bbace6b0b8e8bf9ce59d9ae68c81e8b5b0e59091e4b88ae79a84e8b7afefbc8ce585a8e58a9be4bba5e5aea1e6858ee8b7b5e8a18ce6ada3e4b989e3808222"),
		GasLimit:   91201600000,
		Coinbase:   common.Address{},
		Alloc:      nil,
		Number:     0,
		GasUsed:    0,
		ParentHash: common.Hash{},
	}
)

func createGenesis(ctx *cli.Context) error {
	outputDir := ctx.String(OutputFlag.Name)

	childNodeConfig, err := deploytools.ReadNodeConfig(filepath.Join(outputDir, ctx.String(deploytools.NodeConfigFileFlag.Name)))
	if err != nil {
		return err
	}
	childContractConfig, err := readChildChainContractConfig(filepath.Join(outputDir, ctx.String(ChildchainContractFileFlag.Name)))
	if err != nil {
		return err
	}
	nodes, err := createCbftNode(childNodeConfig.Nodes)
	if err != nil {
		return err
	}
	allocs, err := createAlloc(childNodeConfig.Nodes)
	if err != nil {
		return err
	}
	genesis := DefaultGenesisTemplate
	genesis.Config.Cbft.Period = ctx.Uint64(GenesisCbftPeriodFlag.Name)
	genesis.Config.Cbft.Amount = uint32(ctx.Uint64(GenesisCbftAmountFlag.Name))
	genesis.Alloc = allocs
	genesis.Config.ChainID = big.NewInt(int64(childContractConfig.ChainId))
	genesis.Config.Cbft.InitialNodes = nodes
	genesis.Config.Modules = make(map[string]json.RawMessage)
	createEmptyModuleGenesis(&genesis)
	genesis.Config.Modules[l1.ModuleName] = marshal(l1.L1ConfigParams{
		ChainID:        big.NewInt(ctx.Int64(RootChainIdFlag.Name)),
		StateSender:    deploytools.MustStringToAddress(childContractConfig.StateSenderAddr),
		Checkpoint:     deploytools.MustStringToAddress(childContractConfig.CheckpointManagerAddr),
		StakeManager:   deploytools.MustStringToAddress(childContractConfig.StakeManagerAddr),
		DepositManager: deploytools.MustStringToAddress(childContractConfig.DepositManagerAddr),
	})
	genesis.Config.Modules[stage.ModuleName] = marshal(config.StageNetworkParams{
		RoundValidatorElectionDistance: ctx.Uint64(GenesisStageRoundValidatorElectionDistanceFlag.Name),
		RoundSize:                      ctx.Uint64(GenesisStageRoundSizeFlag.Name),
		EpochSize:                      ctx.Uint64(GenesisStageEpochSizeFlag.Name),
	})
	genesis.Config.Modules[vrf.ModuleName] = marshal(config2.VRFNetworkParams{
		GenesisVRFNonce: common.HexToHash(ctx.String(GenesisVrfNonceFlag.Name)),
	})
	genesis.Config.Modules[staking.ModuleName] = marshal(config3.StakeNetworkParams{
		GenesisValidatorOwner:        deploytools.MustStringToAddress(childContractConfig.GenesisValidatorOwner),
		GenesisStakeAmount:           deploytools.MustStringToDecimal(childContractConfig.GenesisStakeAmount).Uint64(),
		GenesisCommissionRate:        deploytools.MustStringToDecimal(childContractConfig.GenesisCommissionRate).Uint64(),
		StakeWithdrawalWaitPeriod:    ctx.Uint64(GenesisStakeWithdrawalWaitPeriodFlag.Name),
		DelegateWithdrawalWaitPeriod: ctx.Uint64(GenesisStakeDelegateWithdrawalWaitPeriodFlag.Name),
		SlashingPercentage:           ctx.Uint64(GenesisStakeSlashingPercentageFlag.Name),
		SlashIncentivePercentage:     ctx.Uint64(GenesisStakeSlashIncentivePercentageFlag.Name),
		MaxRoundValidatorsSize:       ctx.Uint64(GenesisStakeMaxRoundValidatorsSizeFlag.Name),
		MaxEpochValidatorsSize:       ctx.Uint64(GenesisStakeMaxEpochValidatorsSizeFlag.Name),
		MinBlocksOfRoundValidator:    ctx.Uint64(GenesisStakeMinBlocksOfRoundValidatorFlag.Name),
	})
	genesis.Config.Modules[reward.ModuleName] = marshal(config4.RewardNetworkParams{
		ModuleGenesisConfig: module.ModuleGenesisConfig{},
		RewardPerBlock:      deploytools.MustStringToDecimal(ctx.String(GenesisRewardPerBlockFlag.Name)),
		RewardPerEpoch:      deploytools.MustStringToDecimal(ctx.String(GenesisRewardPerEpochFlag.Name)),
	})
	genesis.Config.Modules[votetoken.ModuleName] = marshal(votetoken.GenesisParams{
		ModuleGenesisConfig: module.ModuleGenesisConfig{},
		Name:                ctx.String(GenesisVoteTokenNameFlag.Name),
		Symbol:              ctx.String(GenesisVoteTokenSymbolFlag.Name),
		Version:             ctx.String(GenesisVoteTokenVersionFlag.Name),
		Owner:               deploytools.MustStringToAddress(ctx.String(GenesisVoteTokenOwnerFlag.Name)),
	})

	genesis.Config.Modules[gov.ModuleName] = marshal(gov.GenesisParams{
		ModuleGenesisConfig: module.ModuleGenesisConfig{},
		Name:                ctx.String(GenesisGovNameFlag.Name),
		Version:             ctx.String(GenesisGovVersionFlag.Name),
		VoteDelay:           deploytools.MustStringToDecimal(ctx.String(GenesisGovVoteDelayFlag.Name)),
		VotePeriod:          deploytools.MustStringToDecimal(ctx.String(GenesisGovVotePeriodFlag.Name)),
		QuorumNumerator:     deploytools.MustStringToDecimal(ctx.String(GenesisGovQuorumNumeratorFlag.Name)),
		ProposalThreshold:   deploytools.MustStringToDecimal(ctx.String(GenesisGovProposalThresholdFlag.Name)),
		Owner:               deploytools.MustStringToAddress(ctx.String(GenesisGovOwnerFlag.Name)),
		VoteToken:           constants.VoteTokenAddress,
	})
	data, err := json.MarshalIndent(genesis, " ", " ")
	if err != nil {
		return err
	}
	path := filepath.Join(ctx.String(OutputFlag.Name), ctx.String(deploytools.GenesisFileFlag.Name))
	err = os.WriteFile(path, data, 0666)
	if err != nil {
		log.Error("Write config failed", "path", path)
	}
	return nil
}

func marshal(m interface{}) json.RawMessage {
	v, _ := json.MarshalIndent(m, " ", " ")
	return v
}

func createEmptyModuleGenesis(genesis *core.Genesis) {
	for _, name := range []string{
		asyncblock.ModuleName,
		benchmark.ModuleName,
		blocktime.ModuleName,
		checkpoint.ModuleName,
		consensusnetwork.ModuleName,
		deposit.ModuleName,
		stateevent.ModuleName,
		extravote.ModuleName,
		statesender.ModuleName,
		nontxpool.ModuleName,
		statesync.ModuleName,
		txrelayer.ModuleName,
	} {
		genesis.Config.Modules[name] = json.RawMessage("{}")
	}
}

func createCbftNode(nodes []*deploytools.Node) ([]params.CbftNode, error) {
	var cbftNode []params.CbftNode
	for _, acc := range nodes {
		if !acc.Genesis {
			continue
		}
		var blsKey bls.SecretKey
		keyBuf := deploytools.MustDecodeString(acc.BlsPrivateKey)
		blsKey.Deserialize(keyBuf)
		nodePrivate, err := crypto.HexToECDSA(acc.NodePrivateKey)
		if err != nil {
			return nil, err
		}
		host := net.ParseIP(acc.IP)
		node := enode.NewV4(&nodePrivate.PublicKey, host, acc.P2pPort, acc.P2pPort)

		cbftNode = append(cbftNode, params.CbftNode{
			Node:      node,
			BlsPubKey: *blsKey.GetPublicKey(),
		})
	}
	return cbftNode, nil
}
func createAlloc(nodes []*deploytools.Node) (core.GenesisAlloc, error) {
	alloc := make(core.GenesisAlloc)
	balance, _ := new(big.Int).SetString("2000000000000000000000000000000000000", 16)
	for _, acc := range nodes {
		alloc[deploytools.MustStringToAddress(acc.NodeAddress)] = core.GenesisAccount{
			Balance: balance,
		}
		if acc.Extra != nil {
			addr := deploytools.MustStringToAddress(acc.Extra.(map[string]interface{})["l2_keystore_address"].(string))
			alloc[addr] = core.GenesisAccount{
				Balance: balance,
			}
		}
	}
	return alloc, nil
}
