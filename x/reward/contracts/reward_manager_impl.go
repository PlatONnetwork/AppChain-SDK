package contracts

import (
	"errors"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
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

func (c *RewardManager) PendingDelegaterRewards(account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *RewardManager) PendingValidatorRewards(account common.Address) (*big.Int, error) {
	panic("implement")
}

func (c *RewardManager) WithdrawDelegaterReward(validator common.Address) error {
	panic("implement")
}

func (c *RewardManager) WithdrawValidatorReward(validator common.Address) error {
	panic("implement")
}
