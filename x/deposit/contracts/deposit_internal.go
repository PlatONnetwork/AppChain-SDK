package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	deposittypes "github.com/PlatONnetwork/AppChain-SDK/x/deposit/types"
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
	DEPOSIT_PARAMS_TYPE  = abi.MustNewType("tuple(address depositor, address recipient, uint256 amount)")
	WITHDRAW_PARAMS_TYPE = abi.MustNewType("tuple(bytes32 sig, address withdrawer, address recipient, uint256 amount)")
)

func (c *DepositHandler) SetL1Module(l1Module deposittypes.L1Moduler) {
	c.l1Module = l1Module
}

func (c *DepositHandler) onDeposit(input []byte) error {
	decoded, err := abi.Decode(DEPOSIT_PARAMS_TYPE, input)
	if nil != err {
		log.Error("Failed to decode deposit data", "error", err)
		return typesdk.NewRevertError("DepositHandler: DECODE_DEPOSIT_DATA_FAILED")
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		return typesdk.NewRevertError("DepositHandler: INVALID_DEPOSIT_DATA")
	}

	depositor, ok := res["depositor"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("DepositHandler: INVALID_DEPOSITOR")
	}

	recipient, ok := res["recipient"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("DepositHandler: INVALID_RECIPIENT")
	}

	amount, ok := res["amount"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("DepositHandler: INVALID_AMOUNT")
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
		log.Error("Failed to encode withdraw syncState data", "withdrawer", withdrawer.Hex(), "recipient", recipient, "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("DepositHandler: encode L2StateSender withdraw data %s", err))
	}

	l2statesender, err := statesenderC.NewL2StateSenderCaller(c.evm, c.contract, constants.StateSenderAddress)
	if nil != err {
		log.Error("Failed to call NewL2StateSenderCaller", "withdrawer", withdrawer.Hex(), "recipient", recipient, "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("DepositHandler: call withdraw by L2StateSender %s", err))
	}

	rootchainDepositManagerAddress, err := c.l1Module.GetDepositManagerAddress()
	if nil != err {
		return typesdk.NewRevertError("DepositHandler: NOT FOUND DEPOSIT MANAGER ADDR")
	}

	if err := l2statesender.SyncState(rootchainDepositManagerAddress, data); nil != err {
		log.Error("Failed to call SyncState", "withdrawer", withdrawer.Hex(), "recipient", recipient, "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("DepositHandler: call withdraw by L2StateSender %s", err))
	}
	return nil
}
