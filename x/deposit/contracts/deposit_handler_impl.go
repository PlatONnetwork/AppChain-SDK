package contracts

import (
	"bytes"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	deposittypes "github.com/PlatONnetwork/AppChain-SDK/x/deposit/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = typesdk.RevertError{}
	_ = vm.EVM{}
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

type DepositHandler struct {
	abi         *abi.ABI
	methodEntry map[string]func([]byte) ([]byte, error)
	readOnly    bool
	contract    *vm.Contract
	evm         *vm.EVM
	burner      contracts.Burn
	stateDb     *contracts.StateDB
	fallback    func(input []byte) ([]byte, error)
	l1Module    deposittypes.L1Moduler
}

func NewDepositHandler(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*DepositHandler, error) {
	s := &DepositHandler{
		abi:      &Abi,
		evm:      evm,
		contract: contract,
		burner:   contracts.NewBurner(contract),
		stateDb:  contracts.NewStateDB(evm, contract),
		readOnly: readOnly,
	}
	s.initMethodEntry()
	return s, nil
}

func (c *DepositHandler) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {

	rootchainDepositManagerAddress, err := c.l1Module.GetDepositManagerAddress()
	if nil != err {
		return typesdk.NewRevertError("DepositHandler: NOT FOUND DEPOSIT MANAGER ADDR")
	}

	if c.contract.Caller() != constants.StateSyncAddress || sender != rootchainDepositManagerAddress {
		return typesdk.NewRevertError("DepositHandler: INVALID_SENDER")
	}
	if bytes.Compare(data[:METHODID_SIZE], DEPOSIT_SIG.Bytes()) == 0 {
		return c.onDeposit(data[METHODID_SIZE:])
	} else {
		return typesdk.NewRevertError("DepositHandler: INVALID_METHOD_SIGN")
	}
}

func (c *DepositHandler) Withdraw(recipient common.Address, amount *big.Int) error {

	if c.evm.StateDB.GetBalance(c.contract.Caller()).Cmp(amount) < 0 {
		return typesdk.NewRevertError("DepositHandler: insufficient balance")
	}

	c.evm.StateDB.SubBalance(c.contract.Caller(), amount)

	if err := c.syncStateWithdraw(c.contract.Caller(), recipient, amount); nil != err {
		return err
	}

	if err := c.addLogL2CoinWithdrawEvent(recipient, c.contract.Caller(), amount); nil != err {
		return err
	}

	log.Info("Withdraw for", "withdrawer", c.contract.Caller().Hex(), "recipient", recipient.Hex(), "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}
