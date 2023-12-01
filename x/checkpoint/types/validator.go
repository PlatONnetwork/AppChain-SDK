package types

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/umbracle/ethgo/abi"
)

var accountSetABIType = abi.MustNewType(`tuple(tuple(address _address, bytes blskey)[])`)

type ValidatorMetadata struct {
	Address common.Address
	BlsKey  *bls.PublicKey
}

type AccountSet []*ValidatorMetadata

func NewAccountSet(validators *cbfttypes.Validators) AccountSet {
	as := make(AccountSet, validators.Len())
	for _, node := range validators.Nodes {
		n := node
		as[node.Index] = &ValidatorMetadata{
			Address: common.Address(n.Address),
			BlsKey:  n.BlsPubKey,
		}
	}
	return as
}

func (as AccountSet) Hash() (common.Hash, error) {
	abiEncoded, err := accountSetABIType.Encode([]interface{}{as.ToAPIBinding()})
	if err != nil {
		return common.ZeroHash, err
	}

	return crypto.Keccak256Hash(abiEncoded), nil
}

func (as AccountSet) ToAPIBinding() []*contractsapi.Validator {
	apiBinding := make([]*contractsapi.Validator, len(as))
	for i, v := range as {
		apiBinding[i] = &contractsapi.Validator{
			Address: v.Address,
			BlsKey:  v.BlsKey.SerializeUncompressed(),
		}
	}

	return apiBinding
}
