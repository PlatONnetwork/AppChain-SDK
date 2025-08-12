package main

import (
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type TemplateConfig struct {
	RegistryManagerImplementAddr string `toml:"registry_manager_implement"                     comment:"RegistryManager implement contract address"`
	RegistryManagerProxyAddr     string `toml:"registry_manager_proxy"                     comment:"RegistryManager proxy contract address"`
	RegistryFactoryAddr          string `toml:"registry_factory"               comment:"RegistryFactory contract address"`
}

type ChildChainContractConfig struct {
	GenesisValidatorOwner string `toml:"genesis_validator_owner" comment:"Child chain genesis validator owner"`
	GenesisStakeAmount    string `toml:"genesis_stake_amount" comment:"Child chain stake amount"`
	GenesisCommissionRate string `toml:"genesis_commission_rate" comment:"Child chain commission rate"`
	DeployNumber          uint64 `toml:"deploy_number" comment:"Child chain contracts deploy number"`
	ChainId               uint64 `toml:"chainid" comment:"Child chain chainid"`
	StateSenderAddr       string `toml:"state_sender" comment:"The contract address for Layer2 listening to rootchain events"`
	CheckpointManagerAddr string `toml:"checkpoint_manager" comment:"The contract address for Layer2 to submit checkpoints to the rootchain"`
	ExitHelperAddr        string `toml:"exit_helper" comment:"Exit contract"`
	ChildChainManagerAddr string `toml:"childchain_manager" comment:"Childchain management contract"`
	StakeManagerAddr      string `toml:"stake_manager" comment:"The contract for Layer2 to submit stake-related events to the rootchain"`
	DepositManagerAddr    string `toml:"deposit_manager" comment:"Layer2 token operation contract"`
	WithdrawHandlerAddr   string `toml:"withdraw_handler" comment:"Layer2 token operation contract"`
	GovernanceHandlerAddr string `toml:"governance" comment:"Rootchain governance management contract"`
}

func readTemplateConfig(file string) (*TemplateConfig, error) {
	var config TemplateConfig
	data, err := os.ReadFile(file)
	if err != nil {
		log.Error("Read template config failed", "err", err)
		return nil, err
	}
	if err := toml.Unmarshal(data, &config); err != nil {
		log.Error("Decode template config failed", "err", err)
		return nil, err
	}
	return &config, nil
}

func readChildChainContractConfig(file string) (*ChildChainContractConfig, error) {
	var config ChildChainContractConfig
	data, err := os.ReadFile(file)
	if err != nil {
		log.Error("Read childchain contract config failed", "err", err)
		return nil, err
	}

	if err := toml.Unmarshal(data, &config); err != nil {
		log.Error("Decode childchain contract config failed", "err", err)
		return nil, err
	}

	return &config, nil
}
