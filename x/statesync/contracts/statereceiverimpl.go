package contracts

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
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

type StateReceiver struct {
	abi          *abi.ABI
	methodEntry  map[string]func([]byte) ([]byte, error)
	readOnly     bool
	contract     *vm.Contract
	evm          *vm.EVM
	burner       contracts.Burn
	stateDb      *contracts.StateDB
	fallback     func(input []byte) ([]byte, error)
	verifyQCFunc func(qc *QuorumCert) error
}

func NewStateReceiver(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*StateReceiver, error) {
	s := &StateReceiver{
		abi:      &Abi,
		evm:      evm,
		contract: contract,
		burner:   contracts.NewBurner(contract),
		stateDb:  contracts.NewStateDB(evm, contract),
		readOnly: readOnly,
	}
	s.verifyQCFunc = s.verifySignature
	s.initMethodEntry()
	return s, nil
}

func (c *StateReceiver) GetCommitmentByStateSyncId(id *big.Int) (StateSyncCommitment, error) {
	sm := c.FindCommitment(id)
	if sm == nil {
		return StateSyncCommitment{}, typesdk.NewRevertError("StateReceiver: NO_COMMITMENT_FOR_ID")
	}
	return *sm, nil
}

func (c *StateReceiver) GetRootByStateSyncId(id *big.Int) (common.Hash, error) {
	sm := c.FindCommitment(id)
	if sm == nil {
		return common.Hash{}, typesdk.NewRevertError("StateReceiver: NO_ROOT_FOR_ID")
	}
	return sm.Root, nil
}

func (c *StateReceiver) BatchExecute(proofs [][]common.Hash, objs []StateSync) error {
	if len(proofs) != len(objs) {
		return typesdk.NewRevertError("StateReceiver: UNMATCHED_LENGTH_PARAMETERS")
	}
	for i := 0; i < len(proofs); i++ {
		if err := c.Execute(proofs[i], objs[i]); err != nil {
			return err
		}
	}
	return nil
}

func (c *StateReceiver) Commit(commitment StateSyncCommitment, index uint64, voteProof []common.Hash, qc QuorumCert) error {
	end := c.GetLastCommittedId()
	if commitment.StartId.Cmp(new(big.Int).Add(end, big.NewInt(1))) != 0 {
		return typesdk.NewRevertError("StateReceiver: INVALID_START_ID")
	}
	if commitment.StartId.Cmp(commitment.EndId) > 0 {
		return typesdk.NewRevertError("StateReceiver: INVALID_END_ID")
	}
	if err := c.verifyQCFunc(&qc); err != nil {
		return typesdk.NewRevertError("StateReceiver: SIGNATURE_VERIFICATION_FAILED")
	}

	value, _ := rlp.EncodeToBytes(&commitment)

	if err := merkle.VerifyProof(index, crypto.Keccak256Hash(value).Bytes(), voteProof, qc.ExtendHash); err != nil {
		return typesdk.NewRevertError("StateReceiver: MERKLE_VERIFICATION_FAILED")
	}

	c.SetCommitment(&commitment)

	c.SetLastCommittedId(commitment.EndId)

	return nil
}

func (c *StateReceiver) verifySignature(qc *QuorumCert) error {
	return nil
}

func (c *StateReceiver) Execute(proof []common.Hash, obj StateSync) error {
	execId := c.getExecutedId()
	if obj.Id.Cmp(new(big.Int).Add(execId, big.NewInt(1))) != 0 {
		return typesdk.NewRevertError("StateReceiver: INVALID_EXEC_ID")
	}
	sm := c.FindCommitment(obj.Id)
	if sm == nil {
		return typesdk.NewRevertError("StateReceiver: NO_COMMITMENT_FOR_ID")
	}
	bytes, _ := rlp.EncodeToBytes(obj)
	hash := crypto.Keccak256Hash(bytes)
	if err := merkle.VerifyProof(new(big.Int).Sub(obj.Id, sm.StartId).Uint64(), hash[:], proof, sm.Root); err != nil {
		return typesdk.NewRevertError("StateReceiver: MERKLE_VERIFICATION_FAILED")
	}
	c.SetExecutedId(obj.Id)
	return nil
}
func (c *StateReceiver) GetExecutedId() (*big.Int, error) {
	return c.getExecutedId(), nil
}
func (c *StateReceiver) GetStateSyncId() (*big.Int, error) {
	return c.GetLastCommittedId(), nil
}
