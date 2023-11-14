package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"math/big"
)

var (
	lastCommittedIdKey = []byte("lastCommittedId")
	executedIdKey      = []byte("executedId")
	commitmentKey      = []byte("commitment")
)

func encodeCommitmentKey(id *big.Int) []byte {
	return append(commitmentKey, id.Bytes()...)
}

func decodeCommitmentKey(value []byte) *big.Int {
	if len(value) >= len(commitmentKey) {
		return new(big.Int).SetBytes(value[len(commitmentKey):])
	}
	return nil
}

func (c *StateReceiver) SetLastCommittedId(id *big.Int) {
	c.evm.StateDB.SetState(c.contract.Address(), lastCommittedIdKey, id.Bytes())
}

func (c *StateReceiver) GetLastCommittedId() *big.Int {
	id := big.NewInt(0)
	value := c.evm.StateDB.GetState(c.contract.Address(), lastCommittedIdKey)
	if len(value) != 0 {
		id.SetBytes(value)
	}
	return id
}

func (c *StateReceiver) SetExecutedId(id *big.Int) {
	c.evm.StateDB.SetState(c.contract.Address(), executedIdKey, id.Bytes())
}

func (c *StateReceiver) GetExecutedId() *big.Int {
	id := big.NewInt(0)
	value := c.evm.StateDB.GetState(c.contract.Address(), executedIdKey)
	if len(value) != 0 {
		id.SetBytes(value)
	}
	return id
}

func (c *StateReceiver) SetCommitment(cm *StateSyncCommitment) {
	value, _ := rlp.EncodeToBytes(cm)
	c.evm.StateDB.SetState(c.contract.Address(), encodeCommitmentKey(cm.EndId), value)
}

func (c *StateReceiver) GetCommitment(end *big.Int) *StateSyncCommitment {
	id := big.NewInt(0)
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeCommitmentKey(end))
	if len(value) != 0 {
		id.SetBytes(value)
	}
	var sc StateSyncCommitment
	rlp.DecodeBytes(value, &sc)

	return &sc
}

func (c *StateReceiver) FindCommitment(id *big.Int) *StateSyncCommitment {
	end := c.GetLastCommittedId()
	if id.Cmp(end) > 0 {
		return nil
	}
	for {
		cm := c.GetCommitment(end)
		if cm != nil {
			if cm.StartId.Cmp(id) >= 0 {
				return cm
			}
			end = new(big.Int).Sub(cm.StartId, big.NewInt(1))
		} else {
			break
		}
	}
	return nil
}
