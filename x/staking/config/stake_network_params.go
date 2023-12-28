package config

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

type StakeNetworkParams struct {
	// For genesis node stake
	// ###### NOTE: ######
	// (Suggest a value greater than or equal to the "minStack" parameter set in
	// the StackManager contract in the root chain,
	// and must be used to perform the same 'GenesisValidator stack' on the root chain`)
	GenesisValidatorOwner        common.Address `json:"genesisValidatorOwner"`        // The account can be used to perform 'unstack' operations on 'GenesisValidator'
	GenesisStakeAmount           uint64         `json:"genesisStakeAmount"`           // The stakeAmount for genesis validator
	GenesisCommissionRate        uint64         `json:"genesisCommissionRate"`        // The epoch reward commissionRate for genesis validator
	StakeWithdrawalWaitPeriod    uint64         `json:"stakeWithdrawalWaitPeriod"`    // Asset lock up period after unstake (unit: epoch)
	DelegateWithdrawalWaitPeriod uint64         `json:"delegateWithdrawalWaitPeriod"` // Asset lock up period after undelegate (unit: epoch)
	SlashingPercentage           uint64         `json:"slashingPercentage"`           // To be read through NetworkParams later
	SlashIncentivePercentage     uint64         `json:"slashIncentivePercentage"`     // Exitor reward, to be read through NetworkParams later
	MaxRoundValidatorsSize       uint64         `json:"maxRoundValidatorsSize"`       // Maximum number of validators for each round
	MaxEpochValidatorsSize       uint64         `json:"maxEpochValidatorsSize"`       // Maximum number of validators for each epoch
	MinBlocksOfRoundValidator    uint64         `json:"minBlocksOfRoundValidator"`    // Minimum blocks quantity threshold for validators in one round

}

func DefualtStakeNetworkParams() *StakeNetworkParams {
	return &StakeNetworkParams{
		GenesisStakeAmount:           0,
		GenesisValidatorOwner:        common.ZeroAddr,
		StakeWithdrawalWaitPeriod:    constants.STAKE_WITHDRAWAL_WAIT_PERIOD,
		DelegateWithdrawalWaitPeriod: constants.DELEGATE_WITHDRAWAL_WAIT_PERIOD,
		SlashingPercentage:           constants.SLASHING_PERCENTAGE,
		SlashIncentivePercentage:     constants.SLASH_INCENTIVE_PERCENTAGE,
		MaxRoundValidatorsSize:       constants.MAX_ROUND_VALIDATORS_SIZE,
		MaxEpochValidatorsSize:       constants.MAX_EPOCH_VALIDATORS_SIZE,
		MinBlocksOfRoundValidator:    constants.MIN_BLOCKS_OF_ROUND_VALIDATOR,
	}
}

func (params *StakeNetworkParams) String() string {
	return fmt.Sprintf(`{"genesisValidatorOwner": "%s", "genesisStakeAmount": %d, "genesisCommissionRate": %d,  "stakeWithdrawalWaitPeriod": %d, "delegateWithdrawalWaitPeriod": %d, "slashingPercentage": %d, "slashIncentivePercentage": %d, "maxRoundValidatorsSize": %d, "maxEpochValidatorsSize": %d, "minBlocksOfRoundValidator": %d}`,
		params.GenesisValidatorOwner.Hex(), params.GenesisStakeAmount, params.GenesisCommissionRate, params.StakeWithdrawalWaitPeriod, params.DelegateWithdrawalWaitPeriod, params.SlashingPercentage, params.SlashIncentivePercentage, params.MaxRoundValidatorsSize, params.MaxEpochValidatorsSize, params.MinBlocksOfRoundValidator)
}
