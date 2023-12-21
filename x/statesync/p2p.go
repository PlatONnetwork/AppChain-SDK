package statesync

import (
	"context"
	"errors"
	"math/big"
	"slices"
	"sync"
	"time"

	sdkp2p "github.com/PlatONnetwork/AppChain-SDK/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
)

const (
	loopInterval = time.Second * 2
)

type Heartbeat struct {
	sdkp2p.MessageCode
	SyncStatus
}

type SyncStatus struct {
	Id          *big.Int
	BlockNumber *big.Int
}

type Peer struct {
	sync.Mutex
	*sdkp2p.DefaultPeer
	syncStatus *SyncStatus
}

func (p *Peer) SetSyncStatus(status *SyncStatus) {
	p.Lock()
	defer p.Unlock()
	p.syncStatus = status
}
func (p *Peer) SyncId() *big.Int {
	p.Lock()
	defer p.Unlock()
	if p.syncStatus == nil {
		return nil
	}
	return p.syncStatus.Id
}

type SyncP2P struct {
	mutex      sync.Mutex
	syncStatus *SyncStatus
	p2p        *sdkp2p.Protocol
}

func NewSyncP2P() *SyncP2P {
	syncP2P := &SyncP2P{}
	protocol := sdkp2p.NewProtocol("l1sync", 1, 10)
	protocol.RegistryMessageType([]sdkp2p.Message{&Heartbeat{}})
	protocol.SetNewPeer(func(p *p2p.Peer, rw p2p.MsgReadWriter) sdkp2p.Peer {
		return &Peer{
			DefaultPeer: sdkp2p.NewDefaultPeer(p, rw),
		}
	})
	protocol.SetUserHandleMsg(func(peer sdkp2p.Peer, msg sdkp2p.Message) error {
		return syncP2P.handleMsg(peer, msg)
	})
	syncP2P.p2p = protocol
	return syncP2P
}
func (s *SyncP2P) Protocols() []p2p.Protocol {
	return s.p2p.Protocol()
}
func (s *SyncP2P) sendHeartbeat() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if s.syncStatus != nil {
		s.p2p.Broadcast(nil, nil, &Heartbeat{
			SyncStatus: *s.syncStatus,
		})
	}
}
func (s *SyncP2P) Run(ctx context.Context) {
	ticker := time.NewTicker(loopInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.sendHeartbeat()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (s *SyncP2P) handleMsg(peer sdkp2p.Peer, msg sdkp2p.Message) error {
	switch m := msg.(type) {
	case *Heartbeat:
		peer.(*Peer).SetSyncStatus(&m.SyncStatus)
	default:
		return errors.New("unknown message type")
	}
	return nil
}

func (s *SyncP2P) SetSyncStatus(status *SyncStatus) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.syncStatus = status
}

func (s *SyncP2P) GetQuorumSyncId(validPeers map[string]struct{}) *big.Int {
	var status []*big.Int
	for _, peer := range s.p2p.Peers() {
		if _, ok := validPeers[peer.Id()]; ok {
			syncId := peer.(*Peer).SyncId()
			if syncId != nil {
				status = append(status, syncId)
			}
		}
	}
	if len(status) == 0 || len(status) < (len(validPeers))/3+1 {
		return nil
	}
	slices.SortFunc(status, func(a, b *big.Int) int {
		return a.Cmp(b)
	})
	return status[0]
}
