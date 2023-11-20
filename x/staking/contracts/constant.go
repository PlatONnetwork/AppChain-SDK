package contracts

import "math/big"

// todo maybe be governanced
var (
	MIN_STAKE                       = new(big.Int).SetUint64(10000000000)
	MIN_DELEGATE                    = new(big.Int).SetUint64(1000000)
	STAKE_WITHDRAWAL_WAIT_PERIOD    = uint64(2)
	DELEGATE_WITHDRAWAL_WAIT_PERIOD = uint64(2)
)
