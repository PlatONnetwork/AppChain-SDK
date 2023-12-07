package wrap

import (
	"errors"
	vrfdb "github.com/PlatONnetwork/AppChain-SDK/x/vrf/db"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func GetPreviousNonce(db sdk.StateDBReader, addr basecommon.Address, blockNumber uint64) (basecommon.Hash, error) {
	if blockNumber == 0 {
		return basecommon.ZeroHash, errors.New("can not found vrf nonce of parent of genesis block")
	}
	nonce := vrfdb.GetNonce(db, addr, blockNumber-1)
	if nonce == basecommon.ZeroHash {
		return basecommon.ZeroHash, vrfdb.ErrNotFound
	}
	return nonce, nil
}

func GetPreviousNonceAndProof(db sdk.StateDBReader, addr basecommon.Address, blockNumber uint64) ([]byte, error) {
	if blockNumber == 0 {
		return nil, errors.New("can not found vrf nonce of parent of genesis block")
	}
	nonceAndProof := vrfdb.GetNonceAndProof(db, addr, blockNumber-1)
	if len(nonceAndProof) == 0 {
		return nil, vrfdb.ErrNotFound
	}
	return nonceAndProof, nil
}

func GetCurrentNonce(db sdk.StateDBReader, addr basecommon.Address, blockNumber uint64) (basecommon.Hash, error) {

	nonce := vrfdb.GetNonce(db, addr, blockNumber)
	if nonce == basecommon.ZeroHash {
		return basecommon.ZeroHash, vrfdb.ErrNotFound
	}
	return nonce, nil
}

func GetCurrentNonceAndProof(db sdk.StateDBReader, addr basecommon.Address, blockNumber uint64) ([]byte, error) {

	nonceAndProof := vrfdb.GetNonceAndProof(db, addr, blockNumber)
	if len(nonceAndProof) == 0 {
		return nil, vrfdb.ErrNotFound
	}
	return nonceAndProof, nil
}

// from tail to head
func GetNonceQueueFromTail(db sdk.StateDBReader, addr basecommon.Address, blockNumber, size uint64) ([]basecommon.Hash, error) {

	index := blockNumber
	count := uint64(0)

	queue := make([]basecommon.Hash, size)

	for index != 0 && count < size {
		nonce, err := GetCurrentNonce(db, addr, index)
		if nil != err {
			return nil, err
		}
		queue[count] = nonce
		index--
		count++
	}
	return queue[:count], nil

}
