package contracts

import (
	"encoding/binary"
	"math/big"
)

var (
	key = []byte("counter")
)

func (c *Counter) count() uint64 {
	val := c.stateDb.GetState(c.contract.Address(), key)
	if len(val) != 0 {
		return binary.BigEndian.Uint64(val)
	}
	return 0
}

func (c *Counter) incr() {
	n := c.count()
	n = n + 1

	c.set(n)
}

func (c *Counter) dec() {
	old := c.count()
	if old > 0 {
		old = old - 1
	}

	c.set(old)
}

func (c *Counter) add(n *big.Int) {
	i := c.count()
	i = i + n.Uint64()

	c.set(i)
}

func (c *Counter) minus(n *big.Int) {
	i := c.count()
	if i < n.Uint64() {
		i = 0
	} else {
		i = i - n.Uint64()
	}

	c.set(i)
}

func (c *Counter) set(n uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], n)
	c.stateDb.SetState(c.contract.Address(), key, data[:])
}
