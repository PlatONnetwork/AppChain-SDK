package contracts

import (
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

func (c *DepositHandler) addLogL2CoinDepositEvent(recipient, depositor basecommon.Address, amount *big.Int) error {
	c.EmitL2CoinDepositEvent(recipient, depositor, amount)
	return nil
}

func (c *DepositHandler) addLogL2CoinWithdrawEvent(recipient, withdrawer basecommon.Address, amount *big.Int) error {
	c.EmitL2CoinWithdrawEvent(recipient, withdrawer, amount)
	return nil
}
