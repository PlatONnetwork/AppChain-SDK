package config

import (
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
)

type StageNetworkParams struct {
	module.ModuleGenesisConfig
	// The number of blocks  between the block height of the election next round validator
	// and the block height at the end of the current round.
	RoundValidatorElectionDistance uint64 `json:"roundValidatorElectionDistance"`
	RoundSize                      uint64 `json:"roundSize"` // number of blocks in round
	EpochSize                      uint64 `json:"epochSize"` // number of blocks in epoch
}

func DefualtStageNetworkParams() *StageNetworkParams {
	return &StageNetworkParams{
		ModuleGenesisConfig: module.ModuleGenesisConfig{CreateBlock: 0},
		RoundValidatorElectionDistance: constants.ROUND_VALIDATOR_ELECTION_DISTANCE,
		RoundSize:                      constants.ROUND_SIZE,
		EpochSize:                      constants.EPOCH_SIZE,
	}
}

func (params *StageNetworkParams) String() string {
	return fmt.Sprintf(`{"createBlock: %d,roundValidatorElectionDistance": %d,"roundSize": %d, "epochSize": %d}`,
		params.CreateBlock, params.RoundValidatorElectionDistance, params.RoundSize, params.EpochSize)
}
