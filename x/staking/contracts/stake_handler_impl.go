package contracts

import (
	"bytes"
	"errors"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgradesys/contracts"
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

type StakeHandler struct {
	abi         *abi.ABI
	methodEntry map[string]func([]byte) ([]byte, error)
	readOnly    bool
	contract    *vm.Contract
	evm         *vm.EVM
	fallback    func(input []byte) ([]byte, error)
}

func NewStakeHandler(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*StakeHandler, error) {
	s := &StakeHandler{
		abi:      &Abi,
		evm:      evm,
		contract: contract,
		readOnly: readOnly,
	}
	s.initMethodEntry()
	return s, nil
}

func (c *StakeHandler) PendingWithdrawalsOfDelegate(account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *StakeHandler) PendingWithdrawalsOfStake(account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *StakeHandler) WithdrawableOfDelegate(account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *StakeHandler) WithdrawableOfStake(account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *StakeHandler) CommitEpoch(id *big.Int, epoch Epoch, epochSize *big.Int) error {
	panic("implement")
}

func (c *StakeHandler) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {
	if err := contracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	// todo change the address
	if c.contract.Caller() != address.StateReceiverAddress || sender != address.RootchainStakeManagerAddress {
		return typesdk.NewRevertError("StakeHandler: INVALID_SENDER")
	}
	if bytes.Compare(data[:METHODID_SIZE], _STAKE_SIG.Bytes()) == 0 {

	} else if bytes.Compare(data[:METHODID_SIZE], _ADDSTAKE_SIG.Bytes()) == 0 {

	} else if bytes.Compare(data[:METHODID_SIZE], _SLASH_SIG.Bytes()) == 0 {

	} else if bytes.Compare(data[:METHODID_SIZE], _DELEGATE_SIG.Bytes()) == 0 {

	}

	panic("implement")
}

func (c *StakeHandler) Slash(validators []common.Address) error {
	if err := contracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	panic("implement")
}

func (c *StakeHandler) Undelegate(validator common.Address, amount *big.Int) error {
	if err := contracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}

	panic("implement")
}

func (c *StakeHandler) Unstake(amount *big.Int) error {
	if err := contracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}

	panic("implement")
}

func (c *StakeHandler) WithdrawUndelegate() error {
	panic("implement")
}

func (c *StakeHandler) WithdrawUnstake() error {
	panic("implement")
}
