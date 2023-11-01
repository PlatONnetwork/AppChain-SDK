package p2p

import (
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"sync"
)

type ListenNode struct {
	mutex     sync.Mutex
	listenNew chan *enode.Node
	close     chan struct{}
	closeFlag bool
}

func NewListenNode() *ListenNode {
	return &ListenNode{
		listenNew: make(chan *enode.Node),
		close:     make(chan struct{}, 1),
		closeFlag: false,
	}
}
func (l *ListenNode) AddNode(rawUrl string) {
	l.mutex.Lock()
	if l.closeFlag {
		l.mutex.Unlock()
		return
	}
	l.mutex.Unlock()
	l.listenNew <- enode.MustParse(rawUrl)

}

func (l *ListenNode) Next() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	return !l.closeFlag || len(l.listenNew) != 0
}

func (l *ListenNode) Node() *enode.Node {
	if l.closeFlag && len(l.listenNew) != 0 {
		return <-l.listenNew
	}
	select {
	case n := <-l.listenNew:
		return n
	case <-l.close:
		return nil
	}
}

func (l *ListenNode) Close() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.close <- struct{}{}
	l.closeFlag = true
}
