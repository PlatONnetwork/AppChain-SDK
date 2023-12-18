package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesender/db"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
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

type L2StateSender struct {
	abi         *abi.ABI
	methodEntry map[string]func([]byte) ([]byte, error)
	readOnly    bool
	contract    *vm.Contract
	evm         *vm.EVM
	burner      contracts.Burn
	stateDb     *contracts.StateDB
	fallback    func(input []byte) ([]byte, error)
	maxLength   uint64
}

func NewL2StateSender(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*L2StateSender, error) {
	s := &L2StateSender{
		abi:       &Abi,
		evm:       evm,
		contract:  contract,
		burner:    contracts.NewBurner(contract),
		stateDb:   contracts.NewStateDB(evm, contract),
		readOnly:  readOnly,
		maxLength: 2048,
	}
	s.initMethodEntry()
	return s, nil
}

func (c *L2StateSender) MAXLENGTH() (*big.Int, error) {
	return new(big.Int).SetUint64(c.maxLength), nil
}

func (c *L2StateSender) Counter() (*big.Int, error) {
	return db.GetCounter(c.evm.StateDB, c.contract.Address()), nil
}

func (c *L2StateSender) SyncState(receiver common.Address, data []byte) error {

	// check receiver
	if receiver == common.ZeroAddr {
		return typesdk.NewRevertError("L2StateSender: INVALID_RECEIVER")
	}

	// check data length
	if uint64(len(data)) > c.maxLength {
		return typesdk.NewRevertError("L2StateSender: EXCEEDS_MAX_LENGTH")
	}
	// State sync id will start with 1
	counter := db.IncrementCounter(c.evm.StateDB, c.contract.Address())
	if err := c.addLogL2StateSyncedEvent(counter, c.contract.Caller(), receiver, data); nil != err {
		return err
	}
	return nil
}
