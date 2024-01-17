package constants

import (
	"math/big"
)

// todo maybe be governanced
var (
	// NOTE: for stage
	//
	// The number of blocks  between the block height of the election next round validator
	// and the block height at the end of the current round.
	ROUND_VALIDATOR_ELECTION_DISTANCE = uint64(20)
	ROUND_SIZE                        = uint64(250)   // number of blocks in round
	EPOCH_SIZE                        = uint64(25000) // number of blocks in epoch
	// NOTE: for stake
	STAKE_WITHDRAWAL_WAIT_PERIOD    = uint64(6)   // Asset lock up period after unstake (unit: epoch)
	DELEGATE_WITHDRAWAL_WAIT_PERIOD = uint64(6)   // Asset lock up period after undelegate (unit: epoch)
	SLASHING_PERCENTAGE             = uint64(50)  // To be read through NetworkParams later
	SLASH_INCENTIVE_PERCENTAGE      = uint64(30)  // Exitor reward, to be read through NetworkParams later
	MAX_ROUND_VALIDATORS_SIZE       = uint64(25)  // Maximum number of validators for each round
	MAX_EPOCH_VALIDATORS_SIZE       = uint64(201) // Maximum number of validators for each epoch
	MIN_BLOCKS_OF_ROUND_VALIDATOR   = uint64(1)   // Minimum blocks quantity threshold for validators in one round
	// NOTE: for reward
	REWARD_PER_BLOCK = new(big.Int).SetUint64(4)     // The validator receives rewards for each block builded
	REWARD_PER_EPOCH = new(big.Int).SetUint64(60000) // The validator receives rewards based on 'stack shares' for each epoch (6000 = 240 * 25)
)
