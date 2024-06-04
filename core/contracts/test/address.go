package test

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

var (
	OneAddr   = common.BigToAddress(big.NewInt(1))
	TwoAddr   = common.BigToAddress(big.NewInt(1))
	ThreeAddr = common.BigToAddress(big.NewInt(1))
	From      = common.BigToAddress(big.NewInt(101))
	To        = common.BigToAddress(big.NewInt(102))
)

func ToAddress(num int64) common.Address {
	return common.BigToAddress(big.NewInt(num))
}
