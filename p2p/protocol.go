package p2p

import (
	"errors"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"reflect"
)

type Protocol struct {
	Name                string
	Version             uint
	Length              uint64
	peerSet             *PeerSet
	registryMessageType []Message
	messageMap          map[reflect.Type]uint64
	dialvalidators      enode.Iterator
	newPeer             func(p *p2p.Peer, rw p2p.MsgReadWriter) Peer
	handshake           func(p Peer, rw p2p.MsgReadWriter) error

	userHandleMsg func(peer Peer, msg Message) error
	nodeInfo      func() interface{}
	peerInfo      func(id enode.ID) interface{}
}

func NewProtocol(name string, version uint, length uint64) *Protocol {
	return &Protocol{
		Name:       name,
		Version:    version,
		Length:     length,
		peerSet:    NewPeerSet(),
		messageMap: make(map[reflect.Type]uint64),
		newPeer: func(p *p2p.Peer, rw p2p.MsgReadWriter) Peer {
			return NewDefaultPeer(p, rw)
		},
		userHandleMsg: func(peer Peer, msg Message) error {
			return nil
		},
	}
}
func (p *Protocol) Protocol() []p2p.Protocol {
	return []p2p.Protocol{
		{
			Name:           p.Name,
			Version:        p.Version,
			Length:         p.Length,
			Run:            p.handleMsg,
			NodeInfo:       p.nodeInfo,
			PeerInfo:       p.peerInfo,
			DialCandidates: p.dialvalidators,
		},
	}
}

func (p *Protocol) SetDialvalidators(iterator enode.Iterator) {
	p.dialvalidators = iterator
}

func (p *Protocol) handleMsg(peer *p2p.Peer, rw p2p.MsgReadWriter) error {
	newPeer := p.newPeer(peer, rw)
	if p.handshake != nil {
		err := p.handshake(newPeer, rw)
		if err != nil {
			return err
		}
	}
	p.peerSet.Register(newPeer)
	defer func() {
		p.peerSet.Unregister(newPeer.Id())
		peer.Disconnect(p2p.DiscUselessPeer)
	}()
	for {
		if err := p.handleUserMsg(newPeer); err != nil {
			return err
		}
	}
	return nil
}

func (p *Protocol) SetNewPeer(newPeer func(p *p2p.Peer, rw p2p.MsgReadWriter) Peer) {
	p.newPeer = newPeer
}

func (p *Protocol) SetHandshake(handshake func(p Peer, rw p2p.MsgReadWriter) error) {
	p.handshake = handshake
}

func (p *Protocol) SetUserHandleMsg(userHandleMsg func(peer Peer, msg Message) error) {
	p.userHandleMsg = userHandleMsg
}

func (p *Protocol) SetNodeInfo(nodeInfo func() interface{}) {
	p.nodeInfo = nodeInfo
}

func (p *Protocol) SetPeerInfo(peerInfo func(id enode.ID) interface{}) {
	p.peerInfo = peerInfo
}

func (p *Protocol) RegistryMessageType(msgs []Message) {
	for i, msg := range msgs {
		msg.SetCode(uint64(i))
		p.registryMessageType = append(p.registryMessageType, msg)
		p.messageMap[getType(msg)] = uint64(i)
	}
}

func (p *Protocol) handleUserMsg(peer Peer) error {
	msg, err := peer.ReadWriter().ReadMsg()
	if err != nil {
		return err
	}
	msgType := p.registryMessageType[msg.Code]
	t := reflect.TypeOf(msgType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	s := reflect.New(t)
	if err = msg.Decode(s.Interface()); err != nil {
		return err
	}
	return p.userHandleMsg(peer, s.Interface().(Message))
}

func (p *Protocol) Peers() []Peer {
	return p.peerSet.Peers()
}

func (p *Protocol) setMessageCode(msg Message) error {
	mtype := getType(msg)
	code, ok := p.messageMap[mtype]
	if !ok {
		return errors.New("invalid message type")
	}
	msg.SetCode(code)
	return nil
}

func (p *Protocol) Send(peer Peer, message Message) {
	p.setMessageCode(message)
	p2p.Send(peer.ReadWriter(), message.Code(), message)
}

func (p *Protocol) Broadcast(filter func(peer Peer) bool, selectFunc func([]Peer) []Peer, message Message) {
	var peers []Peer
	for _, peer := range p.peerSet.Peers() {
		if filter == nil || filter != nil && filter(peer) {
			peers = append(peers, peer)
		}
	}
	if selectFunc != nil {
		peers = selectFunc(peers)
	}
	for _, peer := range peers {
		p.Send(peer, message)
	}
}

func getType(message Message) reflect.Type {
	mtype := reflect.TypeOf(message)
	if mtype.Kind() == reflect.Ptr {
		mtype = mtype.Elem()
	}
	return mtype
}
