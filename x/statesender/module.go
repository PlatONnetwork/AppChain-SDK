package statesender

import (
	"encoding/json"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesender/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	ModuleName = "l2StateSender"
)

type StateSenderModule struct {
	logger log.Logger
}

func NewModule(ctx *cli.Context) *StateSenderModule {
	return &StateSenderModule{
		logger: log.New("module", ModuleName),
	}
}

func (s *StateSenderModule) Name() string {
	return ModuleName
}

func (s *StateSenderModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	// init l2 state sender  account nonce
	initAccountNonce(db, s.Address())
	return nil
}

func (s *StateSenderModule) Address() basecommon.Address {
	return constants.StateSenderAddress
}

func (s *StateSenderModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	l2StateSender, _ := contracts.NewL2StateSender(evm, contract, readOnly)
	return l2StateSender.Run(input)
}
