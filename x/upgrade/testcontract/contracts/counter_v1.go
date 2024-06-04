package contracts

import (
	"errors"
	"math/big"
	"strings"

	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"

	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
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
	ABIV1    = "[{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dec\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
	AbiV1, _ = abi.JSON(strings.NewReader(ABIV1))
)

func (c *Counter) initABIV1() {
	V1 := uint64(1)
	c.abis[V1] = &AbiV1
}

func (c *Counter) initMethodV1Entry() {
	methodEntry := map[string]func([]byte) ([]byte, error){

		"06661abd": c.CountEntry,
		"06fdde03": c.NameEntry,
		"119fbbd4": c.IncrEntry,
		"b3bcfa82": c.DecEntry,
	}
	V1 := uint64(1)
	c.methodEntries[V1] = methodEntry
}

func (c *Counter) DecEntry(input []byte) ([]byte, error) {

	var err error

	err = c.Dec()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}
func (c *Counter) Dec() error {
	c.dec()
	return nil
}
