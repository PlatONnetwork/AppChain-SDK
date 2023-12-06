package db

import (
	"errors"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto/vrf"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	ErrNotFound = errors.New("not found")
)
var (
	nonceAndProofKey = []byte("nonceAndProof") // "nonce":blockNumber => nonceAndProof
)

func EncodeNonceAndProofKey(block uint64) []byte {
	return append(nonceAndProofKey, basecommon.Uint64ToBytes(block)...)
}

func GetNonce(db sdk.StateDBReader, addr basecommon.Address, block uint64) basecommon.Hash {
	nonceAndProof := GetNonceAndProof(db, addr, block)
	if len(nonceAndProof) == 0 {
		return basecommon.ZeroHash
	}
	if block == 0 { // genesis block, only nonce 32 byte
		return basecommon.BytesToHash(nonceAndProof)
	}
	nonceBytes := vrf.ProofToHash(nonceAndProof)
	return basecommon.BytesToHash(nonceBytes)
}

func GetNonceAndProof(db sdk.StateDBReader, addr basecommon.Address, block uint64) []byte {
	return db.GetState(addr, EncodeNonceAndProofKey(block))
}

func SetNonceAndProof(db sdk.StateDB, addr basecommon.Address, block uint64, nonceAndProof []byte) {
	db.SetState(addr, EncodeNonceAndProofKey(block), nonceAndProof)
}
