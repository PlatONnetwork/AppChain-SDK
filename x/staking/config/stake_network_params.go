package config

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

type StakeNetworkParams struct {
	// For genesis node stake
	// ###### NOTE: ######
	// (Suggest a value greater than or equal to the "minStack" parameter set in
	// the StackManager contract in the root chain,
	// and must be used to perform the same 'GenesisValidator stack' on the root chain`)
	GenesisStakeAmount           uint64
	GenesisValidatorOwner        common.Address // The account can be used to perform 'unstack' operations on 'GenesisValidator'
	StakeWithdrawalWaitPeriod    uint64         // Asset lock up period after unstake (unit: epoch)
	DelegateWithdrawalWaitPeriod uint64         // Asset lock up period after undelegate (unit: epoch)
	SlashingPercentage           uint64         // To be read through NetworkParams later
	SlashIncentivePercentage     uint64         // Exitor reward, to be read through NetworkParams later
	MaxRoundValidatorsSize       uint64         // Maximum number of validators for each round
	MaxEpochValidatorsSize       uint64         // Maximum number of validators for each epoch
	MinRoundValidatorBlockNumber uint64         // Minimum blocks quantity threshold for validators in one round

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
		MinRoundValidatorBlockNumber: constants.MIN_ROUND_VALIDATOR_BLOCK_NUMBER,
	}
}
