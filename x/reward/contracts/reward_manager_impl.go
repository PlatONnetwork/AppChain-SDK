package contracts

import (
	"errors"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	rewarddb "github.com/PlatONnetwork/AppChain-SDK/x/reward/db"
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
	abi         *abi.ABI
	methodEntry map[string]func([]byte) ([]byte, error)
	readOnly    bool
	contract    *vm.Contract
	evm         *vm.EVM
	fallback    func(input []byte) ([]byte, error)
}

func NewRewardManager(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*RewardManager, error) {
	s := &RewardManager{
		abi:      &Abi,
		evm:      evm,
		contract: contract,
		readOnly: readOnly,
	}
	s.initMethodEntry()
	return s, nil
}

func (c *RewardManager) PaidRewardPerEpoch(epochId *big.Int) (*big.Int, error) {
	panic("implement")
}

func (c *RewardManager) PendingDelegaterRewards(validator common.Address) (*big.Int, error) {
	return rewarddb.GetPendingDelegaterReward(c.evm.StateDB, c.contract.Address(), c.contract.Caller(), validator), nil
}

func (c *RewardManager) PendingValidatorRewards(validator common.Address) (*big.Int, error) {
	return rewarddb.GetPendingValidatorReward(c.evm.StateDB, c.contract.Address(), validator), nil
}

func (c *RewardManager) WithdrawDelegaterReward(validator common.Address) error {

	// todo 这个还要做下合并，再 withdraw

	//stakeHandler, err := stakingC.NewStakeHandlerCaller(c.evm, c.contract, constants.StakeHandlerAddress)
	//if nil != err {
	//	log.Error("Failed to call NewStakeHandlerCaller", "delegaterAddr", c.contract.Caller(), "validatorAddr", validator.Hex(), "error", err)
	//	return typesdk.NewRevertError("RewardManager: invalid stakeHandler")
	//}
	//epochIndex := rewarddb.GetDelegaterRewardPendingIndex(c.evm.StateDB, c.contract.Address(), c.contract.Caller(), validator)
	//stakeHandler.

	return nil
}

func (c *RewardManager) WithdrawValidatorReward(validator common.Address) error {

	owner := rewarddb.GetValidatorRewardOwner(c.evm.StateDB, c.contract.Address(), validator)

	if owner != c.contract.Caller() {
		return typesdk.NewRevertError("RewardManager: invalid caller")
	}

	validatorRewards := rewarddb.GetPendingValidatorReward(c.evm.StateDB, c.contract.Address(), validator)

	rewardPoolBalance := c.evm.StateDB.GetBalance(c.contract.Address())
	if rewardPoolBalance.Cmp(validatorRewards) < 0 {
		log.Error("insufficient balance on reward pool", "will withdraw validator reward", validatorRewards, "reward pool balance", rewardPoolBalance)
		return typesdk.NewRevertError("RewardManager: insufficient balance on reward pool")
	}

	c.evm.Context.Transfer(c.evm.StateDB, c.contract.Address(), c.contract.Caller(), validatorRewards)

	if err := c.addLogEmitValidatorRewardWithdrawalEvent(validator, validatorRewards, c.contract.Caller()); nil != err {
		return err
	}
	log.Info("WithdrawValidatorReward for", "validator", validator.Hex(), "amount", validatorRewards, "caller", c.contract.Caller(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}
