package contracts

import "encoding/binary"

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
	old := c.count()
	old = old + 1

	var data [8]byte
	binary.BigEndian.PutUint64(data[:], old)
	c.stateDb.SetState(c.contract.Address(), key, data[:])
}
