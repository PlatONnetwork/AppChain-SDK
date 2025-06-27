package asyncblock

import (
	"errors"
	sdkp2p "github.com/PlatONnetwork/AppChain-SDK/p2p"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"math/rand"
	"sync"
)

type Entry struct {
	Epoch          uint64
	View           uint64
	BlockNumber    uint64
	ParentSealHash common.Hash
	Header         *types.Header `rlp:"nil"`
	EntryNumber    uint32
	Transactions   []*types.Transaction
	Ending         uint8
	Signature      []byte
}

func (e *Entry) CannibalizeBytes() ([]byte, error) {
	type SignEntry struct {
		Epoch        uint64
		View         uint64
		BlockNumber  uint64
		Header       *types.Header `rlp:"nil"`
		EntryNumber  uint32
		Transactions []*types.Transaction
		Ending       uint8
	}
	return rlp.EncodeToBytes(&SignEntry{
		Epoch:        e.Epoch,
		View:         e.View,
		BlockNumber:  e.BlockNumber,
		Header:       e.Header,
		EntryNumber:  e.EntryNumber,
		Transactions: e.Transactions,
		Ending:       e.Ending,
	})
}

func (e *Entry) Sign(key *bls.SecretKey) error {
	buf, err := e.CannibalizeBytes()
	if err != nil {
		return err
	}
	e.Signature = key.Sign(string(buf)).Serialize()
	if len(e.Signature) == 0 {
		return nil
	}
	return nil
}
func (e *Entry) VerifySign(pubKey *bls.PublicKey) error {
	var sig bls.Sign
	err := sig.Deserialize(e.Signature)
	if err != nil {
		return err
	}
	buf, err := e.CannibalizeBytes()
	if err != nil {
		return err
	}
	if !sig.Verify(pubKey, string(buf)) {
		return errors.New("verify signature failed")
	}
	return nil
}

type EntryMsg struct {
	sdkp2p.MessageCode
	Entry
}

type GetEntryMsg struct {
	sdkp2p.MessageCode
	Epoch       uint64
	View        uint64
	BlockNumber uint64
	EntryNumber uint32
}
type AsyncBlockP2P interface {
	Update(validators []*cbfttypes.ValidateNode)
	Protocols() []p2p.Protocol
	BroadcastEntry(m sdkp2p.Message)
	Send(peer sdkp2p.Peer, msg sdkp2p.Message)
	SendValidator(m *GetEntryMsg)
}

type asyncBlockP2P struct {
	sync.Mutex
	validators     []*cbfttypes.ValidateNode
	ValidatorPeers []sdkp2p.Peer
	p2p            *sdkp2p.Protocol
}

func NewAsyncBlockP2P(userHandleMsg func(peer sdkp2p.Peer, msg sdkp2p.Message) error) AsyncBlockP2P {
	ab := &asyncBlockP2P{
		p2p: sdkp2p.NewProtocol(ModuleName, 1, 16),
	}
	ab.p2p.RegistryMessageType([]sdkp2p.Message{&EntryMsg{}, &GetEntryMsg{}})
	ab.p2p.SetNewPeer(ab.newPeer)
	ab.p2p.SetUserHandleMsg(userHandleMsg)
	return ab
}
func (a *asyncBlockP2P) newPeer(p *p2p.Peer, rw p2p.MsgReadWriter) sdkp2p.Peer {
	a.Lock()
	defer a.Unlock()
	peer := sdkp2p.NewDefaultPeer(p, rw)
	id := p.ID().String()
	for _, v := range a.validators {
		if v.NodeID.String() == id {
			a.ValidatorPeers = append(a.ValidatorPeers, peer)
		}
	}
	return peer
}

func (a *asyncBlockP2P) Update(validators []*cbfttypes.ValidateNode) {
	nodes := make(map[string]struct{})
	for _, v := range validators {
		nodes[v.NodeID.String()] = struct{}{}
	}
	a.Lock()
	defer a.Unlock()
	var vp []sdkp2p.Peer
	peers := a.p2p.Peers()
	for _, p := range peers {
		id := p.Id()
		if _, ok := nodes[id]; ok {
			vp = append(vp, p)
		}
	}
	a.ValidatorPeers = vp
	a.validators = validators
}

func (a *asyncBlockP2P) Protocols() []p2p.Protocol {
	return a.p2p.Protocol()
}

func (a *asyncBlockP2P) BroadcastEntry(m sdkp2p.Message) {
	a.Lock()
	defer a.Unlock()
	for _, peer := range a.ValidatorPeers {
		a.p2p.Send(peer, m)
	}
}

func (a *asyncBlockP2P) Send(peer sdkp2p.Peer, msg sdkp2p.Message) {
	a.p2p.Send(peer, msg)
}

func (a *asyncBlockP2P) SendValidator(m *GetEntryMsg) {
	a.Lock()
	defer a.Unlock()
	if len(a.ValidatorPeers) > 0 {
		a.p2p.Send(a.ValidatorPeers[rand.Intn(len(a.ValidatorPeers))], m)
	}
}
