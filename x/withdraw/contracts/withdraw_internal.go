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
	DEPOSIT_PARAMS_TYPE  = abi.MustNewType("tuple(bytes32 sig, address depositor, address recipient, uint256 amount)")
	WITHDRAW_PARAMS_TYPE = abi.MustNewType("tuple(address withdrawer, address recipient, uint256 amount)")
)

func (c *WithdrawManager) onWithdraw(input []byte) error {
	decoded, err := abi.Decode(WITHDRAW_PARAMS_TYPE, input)
	if nil != err {
		log.Error("Failed to decode withdraw data", "error", err)
		return typesdk.NewRevertError("WithdrawManager: DECODE_WITHDRAW_DATA_FAILED")
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		return typesdk.NewRevertError("WithdrawManager: INVALID_WITHDRAW_DATA")
	}

	withdrawer, ok := res["withdrawer"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("WithdrawManager: INVALID_WITHDRAWER")
	}

	recipient, ok := res["recipient"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("WithdrawManager: INVALID_RECIPIENT")
	}

	amount, ok := res["amount"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("WithdrawManager: INVALID_AMOUNT")
	}
	return c.withdraw(basecommon.Address(withdrawer), basecommon.Address(recipient), amount)
}

func (c *WithdrawManager) withdraw(withdrawer, recipient basecommon.Address, amount *big.Int) error {

	// ###3 NOTE ####
	// defferent to DepositHandler
	// withdraw coin from `WithdrawManager` contract to recipient
	if c.evm.StateDB.GetBalance(c.contract.Address()).Cmp(amount) < 0 {
		return typesdk.NewRevertError("WithdrawManager: insufficient balance")
	}
	c.evm.Context.Transfer(c.evm.StateDB, c.contract.Address(), recipient, amount)
	if err := c.addLogL2MintableCoinWithdrawEvent(recipient, withdrawer, amount); nil != err {
		return err
	}
	log.Info("Withdraw for", "withdrawer", withdrawer.Hex(), "recipient", recipient.Hex(), "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *WithdrawManager) syncStateDeposit(depositor, recipient basecommon.Address, amount *big.Int) error {

	data, err := abi.Encode([]interface{}{DEPOSIT_SIG, depositor, recipient, amount}, DEPOSIT_PARAMS_TYPE)
	if nil != err {
		log.Error("Failed to encode deposit syncState data", "depositor", depositor.Hex(), "recipient", recipient, "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("encode L2StateSender deposit data %s", err))
	}

	l2statesender, err := statesenderC.NewL2StateSenderCaller(c.evm, c.contract, address.StateSenderAddress)
	if nil != err {
		log.Error("Failed to call NewL2StateSenderCaller", "depositor", depositor.Hex(), "recipient", recipient, "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("call deposit by L2StateSender %s", err))
	}

	if err := l2statesender.SyncState(address.RootchainWithdrawHandlerAddress, data); nil != err {
		log.Error("Failed to call SyncState", "depositor", depositor.Hex(), "recipient", recipient, "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("call deposit by L2StateSender %s", err))
	}
	return nil
}
