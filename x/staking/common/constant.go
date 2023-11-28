package common

// todo maybe be governanced
var (
	STAKE_WITHDRAWAL_WAIT_PERIOD      = uint64(2)
	DELEGATE_WITHDRAWAL_WAIT_PERIOD   = uint64(2)
	SLASHING_PERCENTAGE               = uint64(50) // to be read through NetworkParams later
	SLASH_INCENTIVE_PERCENTAGE        = uint64(30) // exitor reward, to be read through NetworkParams later
	ROUND_VALIDATOR_ELECTION_DISTANCE = uint64(20)
	ROUND_SIZE                        = uint64(250) // number of block in round
	EPOCH_SIZE                        = uint64(250) // number of block in epoch
	MAX_ROUND_VALIDATORS_SIZE         = uint64(25)
	MAX_EPOCH_VALIDATORS_SIZE         = uint64(201)
)
