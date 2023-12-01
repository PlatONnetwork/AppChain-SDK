package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	baselog "github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
)

func (c *DepositHandler) addLogL2CoinDepositEvent(recipient, depositor basecommon.Address, amount *big.Int) error {
	log, err := c.EmitL2CoinDepositEvent(recipient, depositor, amount)
	if nil != err {
		baselog.Error("Failed to emit L2CoinDepositEvent", "recipient", recipient.Hex(), "depositor", depositor.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("DepositHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *DepositHandler) addLogL2CoinWithdrawEvent(recipient, withdrawer basecommon.Address, amount *big.Int) error {
	log, err := c.EmitL2CoinWithdrawEvent(recipient, withdrawer, amount)
	if nil != err {
		baselog.Error("Failed to emit LogL2CoinWithdrawEvent", "recipient", recipient.Hex(), "withdrawer", withdrawer.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("DepositHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}
