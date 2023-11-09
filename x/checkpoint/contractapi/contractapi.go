package contractapi

import (
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/umbracle/ethgo/abi"
)

type Validator struct {
	Address common.Address `abi:"_address"`
	BlsKey  [2]*big.Int    `abi:"blsKey"`
}

var ValidatorABIType = abi.MustNewType("tuple(address _address,uint256[2] blsKey)")

func (v *Validator) EncodeAbi() ([]byte, error) {
	return ValidatorABIType.Encode(v)
}
