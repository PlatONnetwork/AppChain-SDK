package nontxpool

import (
	"sync"
	"time"

	sdkp2p "github.com/PlatONnetwork/AppChain-SDK/p2p"
)

const (
	// This is the target size for the packs of transactions or announcements. A
	// pack can get larger than this if a single transactions exceeds this size.
	maxTxPacketSize = 100 * 1024

	// maxQueuedTxs is the maximum number of transactions to queue up before dropping
	// older broadcasts.
	maxQueuedTxs = 4096
)

type Rate struct {
	count  uint64
	update time.Time
}
type Peer struct {
	sync.Mutex
	*sdkp2p.DefaultPeer
	rate Rate
}
