package config

import "github.com/PlatONnetwork/AppChain-SDK/x/constants"

type StageNetworkParams struct {
	// The number of blocks  between the block height of the election next round validator
	// and the block height at the end of the current round.
	RoundValidatorElectionDistance uint64
	RoundSize                      uint64 // number of blocks in round
	EpochSize                      uint64 // number of blocks in epoch
}

func DefualtStageNetworkParams() *StageNetworkParams {
	return &StageNetworkParams{
		RoundValidatorElectionDistance: constants.ROUND_VALIDATOR_ELECTION_DISTANCE,
		RoundSize:                      constants.ROUND_SIZE,
		EpochSize:                      constants.EPOCH_SIZE,
	}
}
