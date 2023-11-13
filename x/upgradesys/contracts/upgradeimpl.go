package contracts

import (
	"errors"
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
	//TODO 调用implement合约
	panic("")
}
func (c *Upgrade) Implement() (common.Address, error) {
	return c.getImplement(), nil
}

func (c *Upgrade) CommitUpgrade(commitment UpgradeCommitment, index uint64, proof [][32]byte, qc QuorumCert) error {
	//TODO 验证签名，默克尔树
	origin := c.getImplement()
	if commitment.Origin != origin {
		return newRevertError("UPGRADE: ORIGIN_INVALID")
	}
	c.setImplement(commitment.Upgrade)
	if len(commitment.Data) > 0 {
		c.evm.Call(c.contract, commitment.Upgrade, commitment.Data, c.contract.Gas, c.contract.Value())
	}
	return nil
}

func (c *Upgrade) Initialized(implement common.Address) error {
	if c.getInit() {
		return newRevertError("UPGRADE: HAD_Initialized")
	}
	c.setInit()
	c.setImplement(implement)
	//TODO 调用实现合约
	return nil
}
