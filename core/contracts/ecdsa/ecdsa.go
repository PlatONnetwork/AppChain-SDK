package ecdsa

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"math/big"
)

func Recover(hash common.Hash, v uint8, r common.Hash, s common.Hash) common.Address {
	var sig [65]byte
	copy(sig[:], r.Bytes())
	copy(sig[32:], s.Bytes())
	sig[64] = v
	pubKey, err := crypto.Ecrecover(hash.Bytes(), sig[:])
	contracts.Require(err == nil, "ECDSA: invalid signature")
	signer := common.BytesToAddress(crypto.Keccak256(pubKey[1:])[12:])
	return signer
}

func ToTypedDataHash(domainSeparator, structHash common.Hash) common.Hash {
	data, _ := AbiEncodePacked([]byte{0x19, 0x01}, domainSeparator, structHash)
	return crypto.Keccak256Hash(data)
}

func AbiEncodePacked(args ...interface{}) ([]byte, error) {
	bytes := make([]byte, 0)
	for _, arg := range args {
		switch val := arg.(type) {
		case *big.Int:
			bytes = append(bytes, common.LeftPadBytes(val.Bytes(), 32)...)
		case bool:
			if val {
				bytes = append(bytes, []byte{0x0, 0x1}...)
			}
		case common.Hash:
			bytes = append(bytes, val[:]...)
		case []byte:
			bytes = append(bytes, val...)
		case common.Address:
			bytes = append(bytes, val[:]...)
		default:
			return nil, fmt.Errorf("unsupport type %T", arg)
		}
	}
	return bytes, nil
}
