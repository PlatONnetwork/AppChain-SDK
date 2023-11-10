package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
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

func (c *StateReceiver) GetCommitmentByStateSyncId(id *big.Int) (StateSyncCommitment, error) {
	sm := c.FindCommitment(id)
	if sm == nil {
		return StateSyncCommitment{}, newRevertError("StateReceiver: NO_COMMITMENT_FOR_ID")
	}
	return *sm, nil
}

func (c *StateReceiver) GetRootByStateSyncId(id *big.Int) ([32]byte, error) {
	sm := c.FindCommitment(id)
	if sm == nil {
		return common.Hash{}, newRevertError("StateReceiver: NO_ROOT_FOR_ID")
	}
	return sm.Root, nil
}

func (c *StateReceiver) BatchExecute(proofs [][][32]byte, objs []StateSync) error {
	if len(proofs) != len(objs) {
		return newRevertError("StateReceiver: UNMATCHED_LENGTH_PARAMETERS")
	}
	for i := 0; i < len(proofs); i++ {
		if err := c.Execute(proofs[i], objs[i]); err != nil {
			return err
		}
	}
	return nil
}

func (c *StateReceiver) Commit(commitment StateSyncCommitment, index uint64, voteProof [][32]byte, qc QuorumCert) error {
	end := c.GetLastCommittedId()
	if commitment.StartId.Cmp(new(big.Int).Add(end, big.NewInt(1))) != 0 {
		return newRevertError("StateReceiver: INVALID_START_ID")
	}
	if commitment.StartId.Cmp(commitment.EndId) > 0 {
		return newRevertError("StateReceiver: INVALID_END_ID")
	}
	if err := c.VerifySignature(&qc); err != nil {
		return newRevertError("StateReceiver: SIGNATURE_VERIFICATION_FAILED")
	}

	value, _ := rlp.EncodeToBytes(&commitment)

	if err := merkle.VerifyProof(index, crypto.Keccak256Hash(value).Bytes(), c.toProof(voteProof), qc.ExtendHash); err != nil {
		return newRevertError("StateReceiver: MERKLE_VERIFICATION_FAILED")
	}

	c.SetCommitment(&commitment)

	c.SetLastCommittedId(commitment.EndId)

	return nil
}
func (c *StateReceiver) toProof(proof [][32]byte) []common.Hash {
	path := make([]common.Hash, len(proof), len(proof))
	for i := 0; i < len(proof); i++ {
		path[i] = proof[i]
	}
	return path
}
func (c *StateReceiver) VerifySignature(qc *QuorumCert) error {
	return nil
}

func (c *StateReceiver) Execute(proof [][32]byte, obj StateSync) error {
	execId := c.GetExecutedId()
	if execId.Cmp(new(big.Int).Add(obj.Id, big.NewInt(1))) != 0 {
		return newRevertError("StateReceiver: INVALID_EXEC_ID")
	}
	sm := c.FindCommitment(obj.Id)
	if sm == nil {
		return newRevertError("StateReceiver: NO_COMMITMENT_FOR_ID")
	}
	bytes, _ := rlp.EncodeToBytes(obj)
	hash := crypto.Keccak256Hash(bytes)
	path := c.toProof(proof)
	if err := merkle.VerifyProof(new(big.Int).Sub(obj.Id, sm.StartId).Uint64(), hash[:], path, sm.Root); err != nil {
		newRevertError("StateReceiver: MERKLE_VERIFICATION_FAILED")
	}
	c.SetExecutedId(obj.Id)
	return nil
}

func (c *StateReceiver) GetStateSyncId() (*big.Int, error) {
	return c.GetLastCommittedId(), nil
}
