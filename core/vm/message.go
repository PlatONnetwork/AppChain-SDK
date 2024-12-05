package vm

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"math/big"
)

type Message struct {
	to         *common.Address
	from       common.Address
	nonce      uint64
	amount     *big.Int
	gasLimit   uint64
	gasPrice   *big.Int
	gasFeeCap  *big.Int
	gasTipCap  *big.Int
	data       []byte
	accessList types.AccessList
	isFake     bool
}

func NewOnlyCallMessage(from common.Address) *Message {
	return &Message{
		from:     from,
		gasPrice: big.NewInt(0),
		gasLimit: 100000000000,
	}
}
func (m Message) From() common.Address {
	return m.from
}

func (m Message) To() *common.Address {
	return m.to
}

func (m Message) GasPrice() *big.Int {
	return m.gasPrice
}

func (m Message) GasFeeCap() *big.Int {
	return m.gasFeeCap
}

func (m Message) GasTipCap() *big.Int {
	return m.gasTipCap
}

func (m Message) Gas() uint64 {
	return m.gasLimit
}

func (m Message) Value() *big.Int {
	return m.amount
}

func (m Message) Nonce() uint64 {
	return m.nonce
}

func (m Message) IsFake() bool {
	return m.isFake
}

func (m Message) Data() []byte {
	return m.data
}

func (m Message) AccessList() types.AccessList {
	return m.accessList
}
