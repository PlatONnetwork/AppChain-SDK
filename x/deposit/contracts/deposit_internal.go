package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	statesenderC "github.com/PlatONnetwork/AppChain-SDK/x/statesender/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"math/big"
)

const (
	METHODID_SIZE = 32
)

var (
	DEPOSIT_SIG  = crypto.Keccak256Hash([]byte("DEPOSIT"))
	WITHDRAW_SIG = crypto.Keccak256Hash([]byte("WITHDRAW"))
)

var (
	DEPOSIT_PARAMS_TYPE = abi.MustNewType("tuple(address depositor, address recipient, uint256 amount)")

	WITHDRAW_PARAMS_TYPE = abi.MustNewType("tuple(bytes32 sig, address withdrawer, address recipient, uint256 amount)")
)

func (c *DepositHandler) onDeposit(input []byte) error {
	decoded, err := abi.Decode(DEPOSIT_PARAMS_TYPE, input)
	if nil != err {
		return typesdk.NewRevertError("DepositHandler: DECODE_DEPOSIT_DATA_FAILED")
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		return typesdk.NewRevertError("DepositHandler: INVALID_DEPOSIT_DATA")
	}

	depositor, ok := res["depositor"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_DEPOSITOR")
	}

	recipient, ok := res["recipient"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_RECIPIENT")
	}

	amount, ok := res["amount"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_AMOUNT")
	}
	return c.deposit(basecommon.Address(depositor), basecommon.Address(recipient), amount)
}

func (c *DepositHandler) deposit(depositor, recipient basecommon.Address, amount *big.Int) error {

	c.evm.StateDB.AddBalance(recipient, amount)
	if err := c.addLogL2CoinDepositEvent(recipient, depositor, amount); nil != err {
		return err
	}
	log.Info("Deposit for", "depositor", depositor.Hex(), "recipient", recipient.Hex(), "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *DepositHandler) syncStateWithdraw(withdrawer, recipient basecommon.Address, amount *big.Int) error {

	data, err := abi.Encode([]interface{}{WITHDRAW_SIG, withdrawer, recipient, amount}, WITHDRAW_PARAMS_TYPE)
	if nil != err {
		return typesdk.NewRevertError(fmt.Sprintf("encode L2StateSender withdraw data %s", err))
	}

	l2statesender, err := statesenderC.NewL2StateSenderCaller(c.evm, c.contract, address.StakeSenderAddress)
	if nil != err {
		return typesdk.NewRevertError(fmt.Sprintf("call withdraw by L2StateSender %s", err))
	}

	if err := l2statesender.SyncState(address.RootchainDepositManagerAddress, data); nil != err {
		return typesdk.NewRevertError(fmt.Sprintf("call withdraw by L2StateSender %s", err))
	}
	return nil
}
