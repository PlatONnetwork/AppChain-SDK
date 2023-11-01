package p2p

type Message interface {
	SetCode(code uint64)
	Code() uint64
}

type MessageCode struct {
	code uint64
}

func (c *MessageCode) SetCode(code uint64) {
	c.code = code
}

func (c *MessageCode) Code() uint64 {
	return c.code
}
