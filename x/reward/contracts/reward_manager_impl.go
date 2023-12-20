package contracts

import (
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	rewarddb "github.com/PlatONnetwork/AppChain-SDK/x/reward/db"
	rewardtypes "github.com/PlatONnetwork/AppChain-SDK/x/reward/types"
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

type RewardManager struct {
	abi          *abi.ABI
	methodEntry  map[string]func([]byte) ([]byte, error)
	readOnly     bool
	contract     *vm.Contract
	evm          *vm.EVM
	burner       contracts.Burn
	stateDb      *contracts.StateDB
	fallback     func(input []byte) ([]byte, error)
	stageModule  rewardtypes.StageModuler
	stakeModule  rewardtypes.StakeModuler
	rewardModule rewardtypes.RewardModuler
}

func NewRewardManager(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*RewardManager, error) {
	s := &RewardManager{
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

// external
func (c *RewardManager) PaidRewardPerEpoch(epochId *big.Int) (*big.Int, error) {
	return rewarddb.GetPaidRewardPerEpoch(c.evm.StateDB, c.contract.Address(), epochId.Uint64()), nil
}

func (c *RewardManager) PendingDelegatorRewards(validator common.Address) (*big.Int, error) {
	return rewarddb.GetPendingDelegatorReward(c.evm.StateDB, c.contract.Address(), c.contract.Caller(), validator), nil
}

func (c *RewardManager) PendingValidatorRewards(validator common.Address) (*big.Int, error) {
	return rewarddb.GetPendingValidatorReward(c.evm.StateDB, c.contract.Address(), validator), nil
}

func (c *RewardManager) WithdrawDelegatorReward(validator common.Address) error {

	if err := c.updateDelegationRewards(c.contract.Caller(), validator); nil != err {
		return typesdk.NewRevertError(fmt.Sprintf("RewardManager: can not update delegation rewards, %s", err))
	}

	rewards, err := c.withdrawDelegationRewards(c.contract.Caller(), validator)
	if nil != err {
		return err
	}

	if err := c.addLogEmitDelegatorRewardWithdrawalEvent(validator, rewards, c.contract.Caller()); nil != err {
		return err
	}
	log.Info("WithdrawDelegatorReward for", "validator", validator.Hex(), "rewards", rewards, "delegator", c.contract.Caller().Hex(),
		"currentEpoch", c.stageModule.GetCurrentEpoch(c.evm.StateDB), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *RewardManager) WithdrawValidatorReward(validator common.Address) error {

	owner := c.stakeModule.GetValidatorOwner(c.evm.StateDB, validator)

	if owner != c.contract.Caller() {
		return typesdk.NewRevertError("RewardManager: invalid caller")
	}

	rewards, err := c.withdrawValidatorRewrads(c.contract.Caller(), validator)
	if nil != err {
		return err
	}

	if err := c.addLogEmitValidatorRewardWithdrawalEvent(validator, rewards, c.contract.Caller()); nil != err {
		return err
	}
	log.Info("WithdrawValidatorReward for", "validator", validator.Hex(), "rewards", rewards, "caller", c.contract.Caller().Hex(),
		"currentEpoch", c.stageModule.GetCurrentEpoch(c.evm.StateDB), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}
