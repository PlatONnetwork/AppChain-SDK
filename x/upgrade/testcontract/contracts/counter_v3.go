package contracts

import (
	"errors"
	"math/big"
	"strings"

	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
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

var (
	ABIV3    = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"minus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	AbiV3, _ = abi.JSON(strings.NewReader(ABIV3))
)

func (c *Counter) initABIV3() {
	V3 := uint64(3)
	c.abis[V3] = &AbiV3
}

func (c *Counter) initMethodV3Entry() {
	methodEntry := map[string]func([]byte) ([]byte, error){

		"06661abd": c.CountEntry,
		"06fdde03": c.NameEntry,
		"1003e2d2": c.AddEntry,
		"119fbbd4": c.IncrEntry,
		"b3bcfa82": c.DecEntry,
		"d23aa460": c.MinusEntry,
	}
	V3 := uint64(3)
	c.methodEntries[V3] = methodEntry
}

func (c *Counter) MinusEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["minus"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Minus(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}
func (c *Counter) Minus(n *big.Int) error {
	c.minus(n)
	return nil
}
