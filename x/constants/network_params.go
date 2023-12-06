package constants

import (
	"math/big"
)

// todo maybe be governanced
var (
	// for stage
	ROUND_VALIDATOR_ELECTION_DISTANCE = uint64(20)
	ROUND_SIZE                        = uint64(250) // number of block in round
	EPOCH_SIZE                        = uint64(250) // number of block in epoch
	// for stake
	STAKE_WITHDRAWAL_WAIT_PERIOD     = uint64(2)
	DELEGATE_WITHDRAWAL_WAIT_PERIOD  = uint64(2)
	SLASHING_PERCENTAGE              = uint64(50) // to be read through NetworkParams later
	SLASH_INCENTIVE_PERCENTAGE       = uint64(30) // exitor reward, to be read through NetworkParams later
	MAX_ROUND_VALIDATORS_SIZE        = uint64(25)
	MAX_EPOCH_VALIDATORS_SIZE        = uint64(201)
	MIN_ROUND_VALIDATOR_BLOCK_NUMBER = uint64(1) // Minimum blocks quantity threshold for validators in one round
	// for reward
	REWARD_PER_BLOCK = new(big.Int).SetUint64(4)
	REWARD_PER_EPOCH = new(big.Int).SetUint64(60000) // 240 * 25
)
