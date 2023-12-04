package common

import (
	"math/big"
)

// todo maybe be governanced
var (
	REWARD_PER_BLOCK = new(big.Int).SetUint64(4)
	REWARD_PER_EPOCH = new(big.Int).SetUint64(60000) // 240 * 25
)
