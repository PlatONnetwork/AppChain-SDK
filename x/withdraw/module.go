package withdraw

import (
	"encoding/json"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/withdraw/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	MODULE_NAME_WITHDRAW = "withdraw"
)

type WithdrawModule struct {
	logger log.Logger
}

func NewWithdrawModule(ctx *cli.Context) *WithdrawModule {
	return &WithdrawModule{
		logger: log.New("module", MODULE_NAME_WITHDRAW),
	}
}

func (w *WithdrawModule) Name() string {
	return MODULE_NAME_WITHDRAW
}

func (w *WithdrawModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	// init withdraw manager  account nonce
	initAccountNonce(db, w.Address())
	return nil
}

func (w *WithdrawModule) Address() basecommon.Address {
	return constants.WithdrawManagerAddress
}

func (w *WithdrawModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	withdrawManaher, _ := contracts.NewWithdrawManager(evm, contract, readOnly)
	return withdrawManaher.Run(input)
}
