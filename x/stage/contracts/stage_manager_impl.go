package contracts

import (
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

type StageManager struct {
	abi           *abi.ABI
	abis          map[uint64]*abi.ABI
	methodEntry   map[string]func([]byte) ([]byte, error)
	methodEntries map[uint64]map[string]func([]byte) ([]byte, error)
	readOnly      bool
	contract      *vm.Contract
	evm           *vm.EVM
	burner        contracts.Burn
	stateDb       *contracts.StateDB
	context       *contracts.Context
	fallback      func(input []byte) ([]byte, error)
}

func NewStageManager(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*StageManager, error) {
	s := &StageManager{
		abi:           nil,
		abis:          make(map[uint64]*abi.ABI),
		methodEntry:   make(map[string]func([]byte) ([]byte, error)),
		methodEntries: make(map[uint64]map[string]func([]byte) ([]byte, error)),
		evm:           evm,
		contract:      contract,
		burner:        contracts.NewBurner(contract),
		stateDb:       contracts.NewStateDB(evm, contract),
		context:       contracts.NewContext(evm, contract),
		readOnly:      readOnly,
	}
	s.initABI()
	s.initMethodEntry()
	s.loadMethodABI()
	return s, nil
}

// @notice Query the period number where the block height is located
// @dev For the convenience of expanding the expansion of serial number with multiple period properties
// @param periodType represents a period of a certain type
// @param blockNumber represents the block sequence number
// @return the period number for query
//
// ## NOTE ##
// periodType options:
// 0: unknown
// 1: round
// 2: epoch
// ...
func (c *StageManager) GetPeriodByBlockNumber(periodType uint8, blockNumber *big.Int) (*big.Int, error) {
	switch periodType {
	case 1: // round
		round, err := c.getRoundByBlockNumber(blockNumber.Uint64())
		return new(big.Int).SetUint64(round), err
	case 2: // epoch
		epoch, err := c.getEpochByBlockNumber(blockNumber.Uint64())
		return new(big.Int).SetUint64(epoch), err
	default:
		return nil, typesdk.NewRevertError("StageManager: UNKNOWN PERIOD TYPE")
	}
}

// ## NOTE ##
// periodType options:
// 0: unknown
// 1: round
// 2: epoch
// ...
func (c *StageManager) GetPeriodEdge(periodType uint8, period *big.Int) (PeriodEdge, error) {

	switch periodType {
	case 1: // round
		edge, err := c.getPeriodEdgeForRound(period.Uint64())
		if nil != err {
			return PeriodEdge{}, err
		}
		return *edge, nil
	case 2: // epoch
		edge, err := c.getPeriodEdgeForEpoch(period.Uint64())
		if nil != err {
			return PeriodEdge{}, err
		}
		return *edge, nil
	default:
		return PeriodEdge{}, typesdk.NewRevertError("StageManager: UNKNOWN PERIOD TYPE")
	}
}

// ## NOTE ##
// periodType options:
// 0: unknown
// 1: round
// 2: epoch
// ...
func (c *StageManager) GetPeriodEdges(periodType uint8, start *big.Int, size *big.Int) (*big.Int, []*big.Int, []PeriodEdge, error) {
	switch periodType {
	case 1: // round
		return c.getPeriodEdgesForRound(start.Uint64(), size.Uint64())
	case 2: // epoch
		return c.getPeriodEdgesForEpoch(start.Uint64(), size.Uint64())
	default:
		return common.Big0, nil, nil, typesdk.NewRevertError("StageManager: UNKNOWN PERIOD TYPE")
	}
}
