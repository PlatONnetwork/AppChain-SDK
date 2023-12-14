package wrap

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	vrfdb "github.com/PlatONnetwork/AppChain-SDK/x/vrf/db"

	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto/vrf"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	ErrInvalidVrfProve = errors.New("invalid vrf prove")
)

func GenerateNonceAndProof(db sdk.StateDBReader, addr basecommon.Address, blockNumber uint64, key *ecdsa.PrivateKey) ([]byte, error) {

	parentNonce, err := GetPreviousNonce(db, addr, blockNumber)
	if nil != err {
		//return nil, fmt.Errorf("can not get previous vrf nonce, blockNumber: %d, %v", blockNumber, err)
		panic(fmt.Errorf("can not get previous vrf nonce, blockNumber: %d, %v", blockNumber, err))
	}
	nonceAndProof, err := vrf.Prove(key, parentNonce.Bytes())
	if nil != err {
		return nil, fmt.Errorf("can not generate vrf nonce and proof, %v", err)
	}

	if len(nonceAndProof) == 0 {
		return nil, fmt.Errorf("can not generate proof, seed:%x", parentNonce)
	}

	return nonceAndProof, nil
}

func VerifyVrf(nonceAndProof []byte, data basecommon.Hash, key *ecdsa.PublicKey) error {

	if succeed, err := vrf.Verify(key, nonceAndProof, data.Bytes()); nil != err {
		return fmt.Errorf("can not verify vrf nonceAndProof, %v", err)
	} else if !succeed {
		return ErrInvalidVrfProve
	}
	return nil
}

func StorageNonceAndProof(db sdk.StateDB, addr basecommon.Address, blockNumber uint64, nonceAndProof []byte) {
	vrfdb.SetNonceAndProof(db, addr, blockNumber, nonceAndProof)
}
