package types

import (
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
)

type L1 interface {
	GetChainID() (*big.Int, error)
	GetStateAddress() (common.Address, error)
	GetCheckpointAddress() (common.Address, error)
}
