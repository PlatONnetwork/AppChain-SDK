package staking

import (
	"encoding/json"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type StakeModule struct {
}

func NewStakeModule() *StakeModule {
	return &StakeModule{}
}

func (s *StakeModule) Name() string {
	return "StakeModule"
}

func (s *StakeModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) {
	if err := initStakeHandler(db); nil != err {
		log.Error("Failed initialize StakeHandler", "error", err)
		panic(err)
	}

}

func (s *StakeModule) Address() common.Address {
	return address.StakeHandlerAddres
}

func (s *StakeModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	stateReceiver, _ := contracts.NewStakeHandler(evm, contract, readOnly)
	return stateReceiver.Run(input)
}

// TODO 还未实现 ...
func (s *StakeModule) IsEndOfRound(blockNumber uint64) bool {
	return false
}
func (s *StakeModule) GetValidator(ctx sdk.Context, blockNumber uint64) (*cbfttypes.Validators, error) {
	return nil, nil
}
func (s *StakeModule) BlocksOfRound() uint64 {
	return 0
}
