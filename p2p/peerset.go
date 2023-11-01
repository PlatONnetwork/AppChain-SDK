package p2p

import (
	"errors"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"sync"
)

type PeerSet struct {
	peers  map[string]Peer
	lock   sync.RWMutex
	closed bool
}

func NewPeerSet() *PeerSet {
	return &PeerSet{
		peers:  make(map[string]Peer),
		lock:   sync.RWMutex{},
		closed: false,
	}
}

func (ps *PeerSet) Register(peer Peer) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	if ps.closed {
		return errors.New("peerset closed")
	}
	ps.peers[peer.Id()] = peer
	return nil
}

func (ps *PeerSet) Peers() []Peer {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	var peers []Peer
	for _, peer := range ps.peers {
		peers = append(peers, peer)
	}
	return peers
}

func (ps *PeerSet) Unregister(id string) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()
	delete(ps.peers, id)
	return nil
}

func (ps *PeerSet) Peer(id string) Peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	return ps.peers[id]
}

func (ps *PeerSet) Len() int {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	return len(ps.peers)
}

func (ps *PeerSet) Close() {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	for _, p := range ps.peers {
		p.Disconnect(p2p.DiscQuitting)
	}
	ps.closed = true
}
