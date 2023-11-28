package p2p

import (
	sdkp2p "github.com/PlatONnetwork/AppChain-SDK/p2p"
	basep2p "github.com/PlatONnetwork/PlatON-Go/p2p"
)

type StakingP2P struct {
	p2p      *sdkp2p.Protocol
	iterator *sdkp2p.ListenNode
}

func NewStakingP2P() *StakingP2P {
	protocol := sdkp2p.NewProtocol("validator", 1, 10)
	iter := sdkp2p.NewListenNode()
	protocol.SetDialvalidators(iter)
	return &StakingP2P{
		p2p:      protocol,
		iterator: iter,
	}
}

func (s *StakingP2P) Protocols() []basep2p.Protocol {
	return s.p2p.Protocol()
}

func (s *StakingP2P) Addnode(rawUrl string) {
	s.iterator.AddNode(rawUrl)
}
