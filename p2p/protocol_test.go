package p2p

import (
	"crypto/ecdsa"
	"crypto/rand"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/stretchr/testify/require"
	"testing"
)

func genSk(num int) []*ecdsa.PrivateKey {
	var sks []*ecdsa.PrivateKey
	for i := 0; i < num; i++ {
		sk, _ := crypto.GenerateKey()
		sks = append(sks, sk)
	}
	return sks
}

type MessageA struct {
	MessageCode
	Info string
}

type MessageB struct {
	MessageCode
	Info uint64
}

func TestProtocolEcho(t *testing.T) {
	protocol := NewProtocol("test", 1, 20)
	protocol.RegistryMessageType([]Message{&MessageA{}, &MessageB{}})
	protocol.SetUserHandleMsg(func(peer Peer, msg Message) error {
		switch m := msg.(type) {
		case *MessageA:
			t.Log("receive message a", m.Info)
			protocol.Send(peer, msg)
		case *MessageB:
			t.Log("receive message b", m.Info)
			for i := 0; i < 5; i++ {
				var id enode.ID
				rand.Read(id[:])
				protocol.peerSet.Register(NewDefaultPeer(p2p.NewPeer(id, "fake", nil), nil))
			}
			count := 0
			filter := func(p Peer) bool {
				if peer.Id() != p.Id() {
					return false
				}
				count++
				return true
			}
			selectFunc := func(s []Peer) []Peer {
				return s
			}
			protocol.Broadcast(filter, selectFunc, msg)
			require.Equal(t, 1, count)
		}
		return nil
	})
	sks := genSk(2)
	_, rw1, peer2, rw2 := p2p.NewPeerByNodeID(&sks[0].PublicKey, &sks[1].PublicKey, protocol.Protocol())
	go protocol.Protocol()[0].Run(peer2, rw2)
	p2p.Send(rw1, 0, &MessageA{Info: "messageA"})
	msg, err := rw1.ReadMsg()
	require.Nil(t, err)
	var a MessageA
	msg.Decode(&a)
	require.Equal(t, "messageA", a.Info)

	p2p.Send(rw1, 1, &MessageB{Info: 888})
	msg, err = rw1.ReadMsg()
	require.Nil(t, err)
	var b MessageB
	msg.Decode(&b)
	require.Equal(t, uint64(888), b.Info)
}
