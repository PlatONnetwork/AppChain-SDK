package types

import "github.com/PlatONnetwork/PlatON-Go/common"

type L1Moduler interface {
	GetDepositManagerAddress() (common.Address, error)
}
