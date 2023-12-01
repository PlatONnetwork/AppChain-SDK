package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	baselog "github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
)

func (c *WithdrawManager) addLogL2MintableCoinDepositEvent(recipient, depositor basecommon.Address, amount *big.Int) error {
	log, err := c.EmitL2MintableCoinDepositEvent(recipient, depositor, amount)
	if nil != err {
		baselog.Error("Failed to emit L2MintableCoinDepositEvent", "recipient", recipient.Hex(), "depositor", depositor.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("WithdrawManager: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *WithdrawManager) addLogL2MintableCoinWithdrawEvent(recipient, withdrawer basecommon.Address, amount *big.Int) error {
	log, err := c.EmitL2MintableCoinWithdrawEvent(recipient, withdrawer, amount)
	if nil != err {
		baselog.Error("Failed to emit L2MintableCoinWithdrawEvent", "recipient", recipient.Hex(), "withdrawer", withdrawer.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("WithdrawManager: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}
