package contracts

import (
	"encoding/hex"
	"errors"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	vrftypes "github.com/PlatONnetwork/AppChain-SDK/x/vrf/types"
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

type VRFHandler struct {
	abi         *abi.ABI
	methodEntry map[string]func([]byte) ([]byte, error)
	readOnly    bool
	contract    *vm.Contract
	evm         *vm.EVM
	fallback    func(input []byte) ([]byte, error)
	stage       vrftypes.Stage
	stake       vrftypes.Stake
}

func NewVRFHandler(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*VRFHandler, error) {
	s := &VRFHandler{
		abi:      &Abi,
		evm:      evm,
		contract: contract,
		readOnly: readOnly,
	}
	s.initMethodEntry()
	return s, nil
}

// for vrf module

func (c *VRFHandler) SetStageModule(stage vrftypes.Stage) {
	c.stage = stage
}

func (c *VRFHandler) SetStakeModule(stake vrftypes.Stake) {
	c.stake = stake
}

func (c *VRFHandler) PushNonceAndProof(nonceAndProof []byte) error {

	if len(nonceAndProof) != 81 { // 81 byte, nonce and proof, flag |nonce |proof, 1byte|32byte|48byte
		return typesdk.NewRevertError("VRFHandler: INVALID PARAM")
	}

	validatorAddr := c.contract.Caller()

	if c.stake.IsInvalidValidator(c.evm.StateDB, validatorAddr) {
		return typesdk.NewRevertError("VRFHandler: INVALID CALLER")
	}
	currentBlock := c.evm.Context.BlockNumber.Uint64()
	if err := c.verifyNonceAndProof(validatorAddr, currentBlock, nonceAndProof); nil != err {
		return err
	}

	c.setNonceAndProof(currentBlock, nonceAndProof)

	log.Info("PushNonceAndProof for", "validatorAddr", validatorAddr, "nonceAndProof", hex.EncodeToString(nonceAndProof),
		"currentEpoch", c.stage.GetCurrentEpoch(c.evm.StateDB), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}
