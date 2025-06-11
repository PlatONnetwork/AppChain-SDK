package p2p

import "github.com/PlatONnetwork/PlatON-Go/p2p"

type Peer interface {
	Id() string
	Disconnect(reason p2p.DiscReason)
	ReadWriter() p2p.MsgReadWriter
	Info() interface{}
}

type DefaultPeer struct {
	id string // Unique ID for the peer, cached

	peer *p2p.Peer         // The embedded P2P package peer
	rw   p2p.MsgReadWriter // Input/output streams for snap
}

func NewDefaultPeer(peer *p2p.Peer, rw p2p.MsgReadWriter) *DefaultPeer {
	return &DefaultPeer{
		id:   peer.ID().String(),
		peer: peer,
		rw:   rw,
	}
}

func (d DefaultPeer) Disconnect(reason p2p.DiscReason) {
	d.peer.Disconnect(reason)
}

func (d DefaultPeer) Id() string {
	return d.id
}

func (d DefaultPeer) ReadWriter() p2p.MsgReadWriter {
	return d.rw
}
func (d DefaultPeer) Info() interface{} {
	return d.id
}
