package contracts

import (
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

func (c *WithdrawManager) addLogL2MintableCoinDepositEvent(recipient, depositor basecommon.Address, amount *big.Int) error {
	c.EmitL2MintableCoinDepositEvent(recipient, depositor, amount)
	return nil
}

func (c *WithdrawManager) addLogL2MintableCoinWithdrawEvent(recipient, withdrawer basecommon.Address, amount *big.Int) error {
	c.EmitL2MintableCoinWithdrawEvent(recipient, withdrawer, amount)
	return nil
}
