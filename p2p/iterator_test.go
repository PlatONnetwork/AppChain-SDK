package p2p

import (
	"github.com/stretchr/testify/require"
	"sync"
	"sync/atomic"
	"testing"
)

func TestIterator(t *testing.T) {
	listen := NewListenNode()
	var counter atomic.Int32
	var group sync.WaitGroup
	group.Add(2)
	go func() {
		for listen.Next() {
			t.Log("next")
			if listen.Node() != nil {
				counter.Add(1)
			}
		}
		t.Log("done")
		group.Done()
	}()
	go func() {
		for i := 0; i < 5; i++ {
			t.Log("add node")
			listen.AddNode("enode://4fcc251cf6bf3ea53a748971a223f5676225ee4380b65c7889a2b491e1551d45fe9fcc19c6af54dcf0d5323b5aa8ee1d919791695082bae1f86dd282dba4150f@0.0.0.0:16789")
		}
		listen.Close()
		group.Done()
	}()
	group.Wait()
	require.Equal(t, int32(5), counter.Load())
}
