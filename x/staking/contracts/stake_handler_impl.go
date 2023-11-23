package contracts

import (
	"bytes"
	"errors"

	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	upgradecontracts "github.com/PlatONnetwork/AppChain-SDK/x/upgradesys/contracts"
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

func (c *StakeHandler) PendingWithdrawalsOfDelegate(validator common.Address, account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *StakeHandler) PendingWithdrawalsOfStake(validator common.Address, account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *StakeHandler) WithdrawableOfDelegate(validator common.Address, account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *StakeHandler) WithdrawableOfStake(validator common.Address, account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *StakeHandler) CommitEpoch(id *big.Int, epoch Epoch, epochSize *big.Int) error {
	panic("implement")
}

func (c *StakeHandler) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	// todo need to change the inner contract address file path
	if c.contract.Caller() != address.StateReceiverAddress || sender != address.RootchainStakeManagerAddress {
		return typesdk.NewRevertError("StakeHandler: INVALID_SENDER")
	}
	if bytes.Compare(data[:METHODID_SIZE], STAKE_SIG.Bytes()) == 0 {
		return c.onStake(data[METHODID_SIZE:])
	} else if bytes.Compare(data[:METHODID_SIZE], ADDSTAKE_SIG.Bytes()) == 0 {
		return c.onAddStake(data[METHODID_SIZE:])
	} else if bytes.Compare(data[:METHODID_SIZE], SLASH_SIG.Bytes()) == 0 {
		return c.onSlash(data[METHODID_SIZE:])
	} else if bytes.Compare(data[:METHODID_SIZE], DELEGATE_SIG.Bytes()) == 0 {
		return c.onDelegate(data[METHODID_SIZE:])
	} else {
		return typesdk.NewRevertError("StakeHandler: INVALID_METHOD_SIGN")
	}
}

func (c *StakeHandler) Slash(validators []common.Address) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	panic("implement")
}

func (c *StakeHandler) Undelegate(validatorAddr common.Address, amount *big.Int) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	// todo ...
	// c.contract.Caller(): msg.sender
	return c.registerDelegateWithdrawal(c.contract.Caller(), validatorAddr, amount)
}

func (c *StakeHandler) Unstake(validatorAddr common.Address, amount *big.Int) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	if err := c.unStake(validatorAddr, amount); nil != err {
		return err
	}
	return c.registerStakeWithdrawal(validatorAddr, amount)
}

func (c *StakeHandler) WithdrawUndelegate(validator common.Address) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}

	currentEpoch := c.GetCurrentEpoch()
	delegater := c.contract.Caller()
	amount, err := c.applyDelegateWithdrawable(delegater, validator, currentEpoch)
	if nil != err {
		log.Error("Failed to withdraw undelegate", "validatorAddr", validator.Hex(),
			"currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: UPDATE STAKE WITHDRAW PENDDING HEAD FAILED")
	}

	if err := c.addLogDelegateWithdrawalEvent(delegater, validator, amount); nil != err {
		return err
	}

	if err := c.syncStateUnDelegate(validator, delegater, amount); nil != err {
		return err
	}

	log.Info("Withdraw undelegate for", "delegater", delegater, "validator", validator.Hex(), "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) WithdrawUnstake(validator common.Address) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}

	currentEpoch := c.GetCurrentEpoch()
	amount, err := c.applyStakeWithdrawable(validator, currentEpoch)
	if nil != err {
		log.Error("Failed to withdraw unstake", "validatorAddr", validator.Hex(),
			"currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: UPDATE STAKE WITHDRAW PENDDING HEAD FAILED")
	}

	if err := c.addLogStakeWithdrawalEvent(validator, amount); nil != err {
		return err
	}

	if err := c.syncStateUnStake(validator, amount); nil != err {
		return err
	}

	log.Info("Withdraw unstake for", "validator", validator.Hex(), "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}
