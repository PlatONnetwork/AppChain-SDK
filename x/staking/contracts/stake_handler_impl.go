package contracts

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common/math"

	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
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
	stage       staketypes.Stage
	reward      staketypes.Reward
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

// internal
func (c *StakeHandler) SetStageModule(stage staketypes.Stage) {
	c.stage = stage
}

func (c *StakeHandler) SetRewardModule(reward staketypes.Reward) {
	c.reward = reward
}

// external
func (c *StakeHandler) PendingWithdrawalsOfDelegate(validator common.Address, delegater common.Address) (*big.Int, error) {
	return c.getDelegateWithdrawalPending(delegater, validator, c.getCurrentEpoch()), nil
}

func (c *StakeHandler) PendingWithdrawalsOfStake(validator common.Address) (*big.Int, error) {
	return c.getStakeWithdrawalPending(validator, c.getCurrentEpoch()), nil
}

func (c *StakeHandler) WithdrawableOfDelegate(validator common.Address, delegater common.Address) (*big.Int, error) {
	return c.getDelegateWithdrawable(delegater, validator, c.getCurrentEpoch()), nil
}

func (c *StakeHandler) WithdrawableOfStake(validator common.Address) (*big.Int, error) {
	return c.getStakeWithdrawable(validator, c.getCurrentEpoch()), nil
}

func (c *StakeHandler) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	// todo need to change the inner contract address file path
	if c.contract.Caller() != constants.StateReceiverAddress || sender != constants.RootchainStakeManagerAddress {
		return typesdk.NewRevertError("StakeHandler: INVALID_SENDER")
	}
	if bytes.Compare(data[:METHODID_SIZE], STAKE_SIG.Bytes()) == 0 {
		return c.onStake(data[METHODID_SIZE:])
	} else if bytes.Compare(data[:METHODID_SIZE], ADDSTAKE_SIG.Bytes()) == 0 {
		return c.onAddStake(data[METHODID_SIZE:])
	} else if bytes.Compare(data[:METHODID_SIZE], SLASH_SIG.Bytes()) == 0 {
		return c.onSlash(data) // don't be data[METHODID_SIZE:], it must be data
	} else if bytes.Compare(data[:METHODID_SIZE], DELEGATE_SIG.Bytes()) == 0 {
		return c.onDelegate(data[METHODID_SIZE:])
	} else {
		return typesdk.NewRevertError("StakeHandler: INVALID_METHOD_SIGN")
	}
}

func (c *StakeHandler) Slash() error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	validators := db.CheckLowBlocksValidatorForPreviousRound(c.evm.StateDB, c.contract.Address())
	// ###### NOTE: ######
	// remove validator from epoch validators
	cache := make(map[common.Address]struct{}, 0)
	for _, validatorAddr := range validators {
		cache[validatorAddr] = struct{}{}
	}
	currentEpoch := c.getCurrentEpoch()
	// NTOE: update epoch validator snapshot queue (after remove low blocks validators)
	epochValidatorAddrQueue := db.GetEpochValidatorSharesSnapshotQueue(c.evm.StateDB, c.contract.Address(), currentEpoch)
	for i := 0; i < len(epochValidatorAddrQueue); i++ {
		validator := epochValidatorAddrQueue[i]
		if _, ok := cache[validator.ValidatorAddr]; !ok {
			// remove the validatorAddr from epoch validatorAddrQueue
			epochValidatorAddrQueue = append(epochValidatorAddrQueue[:i], epochValidatorAddrQueue[i+1:]...)
			i--
		}
	}
	if err := db.SetEpochValidatorSharesSnapshotQueue(c.evm.StateDB, c.contract.Address(), currentEpoch, epochValidatorAddrQueue); nil != err {
		log.Error("Failed to update epochValidators", "epoch", currentEpoch, "error", err)
		return typesdk.NewRevertError("StakeHandler: UPDATE EPOCH VALIDATORS FAILED")
	}

	if err := c.syncStateSlash(validators); nil != err {
		return err
	}

	log.Info("Slash for", "validators", fmt.Sprintf("%+v", validators), "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) Undelegate(validatorAddr common.Address, amount *big.Int) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}

	validator := c.getValidator(validatorAddr)

	delegaterAddr := c.contract.Caller()

	epochs, _ := c.getValidatorDelegationRcPendingAndEpoch(validatorAddr, math.MaxUint64)

	paid := common.Big0

	for _, stakeEpoch := range epochs {

		if amount.Cmp(common.Big0) == 0 {
			break
		}

		delegation := c.getDelegation(delegaterAddr, validatorAddr, stakeEpoch)
		if delegation.IsEmpty() {
			continue
		}

		// NOTE:
		// Priority must be given to settling commission rewards before proceeding with the `withdraw` operation.
		if err := c.reward.UpdateDelegationRewardsByStakeEpoch(c.evm.StateDB, delegaterAddr, validatorAddr, stakeEpoch); nil != err {
			log.Error("Failed to update delegation rewards by stakeEpoch", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
				"stakeEpoch", stakeEpoch, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber, "error", err)
			return typesdk.NewRevertError("StakeHandler: UPDATE DELEGATION REWARDS BY STAKE EPOCH FAILED")
		}

		use := common.Big0
		if delegation.Amount.Cmp(amount) <= 0 {
			// remove the delegation by stakeEpoch
			c.removeDelegation(delegaterAddr, validatorAddr, stakeEpoch)

			// decrement validator-delegater-rc
			if err := c.releaseValidatorDelegationRcItem(validatorAddr, stakeEpoch, 1); nil != err {
				log.Error("Failed to release validatorDelegation rc", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
					"stakeEpoch", stakeEpoch, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber, "error", err)
				return typesdk.NewRevertError("StakeHandler: RELEASE VALIDATOR DELEGATION RC FAILED")
			}
			use = delegation.Amount
		} else {
			// update delegation with new epoch and new amount
			delegation.UpdateEpoch(c.getCurrentEpoch())
			delegation.DecrementAmount(amount)
			if err := c.setDelegation(delegaterAddr, validatorAddr, stakeEpoch, delegation); nil != err {
				log.Error("Failed to set delegation", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
					"stakeEpoch", stakeEpoch, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber, "error", err)
				return typesdk.NewRevertError("StakeHandler: SET DELEGATION FAILED")
			}
			use = amount
		}

		amount = new(big.Int).Sub(amount, use)
		paid = new(big.Int).Add(paid, use)

		// update validator priority
		if validator.IsValid() && validator.Epoch == stakeEpoch {

			validator.SubDelegateAmount(amount)
			if err := c.updateValidatorByPriority(validatorAddr, validator); nil != err {
				log.Error("Failed to update validator priority", "validatorAddr", validatorAddr.Hex(), "error", err)
				return typesdk.NewRevertError("StakeHandler: SUB DELEGATE AMOUNT OF VALIDATOR FAILED")
			}
		}
	}
	if err := c.registerDelegateWithdrawal(delegaterAddr, validatorAddr, paid, true); nil != err {
		return err
	}
	if err := c.addLogUnDelegatedEvent(delegaterAddr, validatorAddr, paid); nil != err {
		return err
	}
	log.Info("Undelegate for", "delegater", delegaterAddr.Hex(), "validator", validatorAddr.Hex(), "amount", amount, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) Unstake(validatorAddr common.Address, amount *big.Int) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}
	if err := c.unStake(validatorAddr, amount); nil != err {
		return err
	}

	if err := c.registerStakeWithdrawal(validatorAddr, amount, true); nil != err {
		return err
	}

	if err := c.addLogUnStakedEvent(validatorAddr, amount); nil != err {
		return err
	}
	log.Info("Unstake for", "validator", validatorAddr.Hex(), "amount", amount, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) WithdrawUndelegate(validator common.Address) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}

	currentEpoch := c.getCurrentEpoch()
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

	log.Info("Withdraw undelegate for", "delegater", delegater, "validator", validator.Hex(), "amount", amount, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) WithdrawUnstake(validator common.Address) error {
	if err := upgradecontracts.OnlyInitialized(c.evm.StateDB, c.contract.Address()); err != nil {
		return err
	}

	currentEpoch := c.getCurrentEpoch()
	amount, err := c.applyStakeWithdrawable(validator, currentEpoch)
	if nil != err {
		log.Error("Failed to withdraw unstake", "validatorAddr", validator.Hex(),
			"currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: UPDATE STAKE WITHDRAW PENDDING HEAD FAILED")
	}

	// remove unstake validator
	validatorInfo := c.getValidator(validator)
	if validatorInfo.IsInvalidUnstaked() {
		c.removeValidator(validator)
	}

	if err := c.addLogStakeWithdrawalEvent(validator, amount); nil != err {
		return err
	}

	if err := c.syncStateUnStake(validator, amount); nil != err {
		return err
	}

	log.Info("Withdraw unstake for", "validator", validator.Hex(), "amount", amount, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}
