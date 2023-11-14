package contracts

import (
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
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

var (
	initialized, _ = hex.DecodeString("8129fc1c")
)

type Upgrade struct {
	abi         *abi.ABI
	methodEntry map[string]func([]byte) ([]byte, error)
	readOnly    bool
	contract    *vm.Contract
	evm         *vm.EVM
	fallback    func(input []byte) ([]byte, error)
}

func NewUpgrade(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Upgrade, error) {
	s := &Upgrade{
		abi:      &Abi,
		evm:      evm,
		contract: contract,
		readOnly: readOnly,
	}
	s.fallback = s.Fallback
	s.initMethodEntry()
	return s, nil
}

func (c *Upgrade) Fallback(input []byte) ([]byte, error) {
	addr := c.getImplement()
	ret, err := contracts.DelegateCall(c.evm, c.contract, addr, input, c.contract.Gas)
	return ret, err
}

func (c *Upgrade) Implement() (common.Address, error) {
	return c.getImplement(), nil
}

func (c *Upgrade) CommitUpgrade(commitment UpgradeCommitment, index uint64, proof [][32]byte, qc QuorumCert) error {
	//TODO 验证签名，默克尔树
	origin := c.getImplement()
	if commitment.Origin != origin {
		return typesdk.NewRevertError("UPGRADE: ORIGIN_INVALID")
	}
	c.setImplement(commitment.Upgrade)
	if len(commitment.Data) > 0 {
		c.evm.Call(c.contract, commitment.Upgrade, commitment.Data, c.contract.Gas, c.contract.Value())
	}
	return nil
}

func (c *Upgrade) Initialize(implement common.Address) error {
	if c.getInitialized() {
		return typesdk.NewRevertError("UPGRADE: HAD_INITIALIZED")
	}
	c.setInitializing()
	c.setImplement(implement)
	ret, err := contracts.DelegateCall(c.evm, c.contract, implement, initialized, c.contract.Gas)
	if err != nil {
		return typesdk.NewRevertError(string(ret))
	}
	c.setInitialized()
	return nil
}
