package sync

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/log"
)

var (
	AbiJson    = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"id","type":"uint256"},{"indexed":true,"internalType":"address","name":"sender","type":"address"},{"indexed":true,"internalType":"address","name":"receiver","type":"address"},{"indexed":false,"internalType":"bytes","name":"data","type":"bytes"}],"name":"StateSynced","type":"event"}]`
	SyncAbi, _ = abi.JSON(strings.NewReader(AbiJson))
)

const (
	maxScanBlock      = 1000
	scanBlockInterval = time.Second * 3
)

type StateSender struct {
	Id       *big.Int
	Sender   common.Address
	Receiver common.Address
	Data     []byte
}

type PlatonClient interface {
	BlockNumber(ctx context.Context) (uint64, error)
	FilterLogs(ctx context.Context, q platon.FilterQuery) ([]types.Log, error)
}

type L1Sync struct {
	log             log.Logger
	stateSenderAddr common.Address
	cli             PlatonClient
	db              *L1SyncDB
	updateCh        chan struct{}
}

func NewL1Sync(stateSenderAddr common.Address, url string, start *big.Int, db store.Store, updateCh chan struct{}) (*L1Sync, error) {
	syncdb := NewL1SyncDB(db)
	log := log.New("module", "l1sync")
	if last, _ := syncdb.LastBlockNumber(); (last == nil || last.Uint64() == 0) && start != nil {
		syncdb.SetLastBlockNumber(start)
	}
	cli, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	number, err := syncdb.LastBlockNumber()
	if err != nil {
		return nil, err
	}

	if start == nil && number == nil {
		return nil, errors.New("need set start block number")
	}
	return &L1Sync{
		log:             log,
		stateSenderAddr: stateSenderAddr,
		cli:             cli,
		db:              syncdb,
		updateCh:        updateCh,
	}, nil
}

func (l *L1Sync) Run(ctx context.Context) {
	if l.cli == nil {
		l.log.Info("l1 sync is not start")
		return
	}
	ticker := time.NewTicker(scanBlockInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				l.scanRootChain(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (l *L1Sync) SyncDB() *L1SyncDB {
	return l.db
}

func (l *L1Sync) scanRootChain(ctx context.Context) error {
	start, err := l.db.LastBlockNumber()
	if err != nil {
		return err
	}
	end, err := l.cli.BlockNumber(ctx)
	if err != nil {
		return err
	}
	l.log.Debug("try scan root chain", "start", start, "end", end)
	return l.scanLogs(ctx, start.Uint64(), end)
}

func (l *L1Sync) scanLogs(ctx context.Context, start, end uint64) error {
	for from := start; from < end; from += maxScanBlock {
		to := from + maxScanBlock
		if from+maxScanBlock > end {
			to = end
		}
		l.log.Debug("try scan logs", "from", from, "to", to)
		logs, err := l.cli.FilterLogs(ctx, platon.FilterQuery{
			FromBlock: new(big.Int).SetUint64(from),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{l.stateSenderAddr},
			Topics:    [][]common.Hash{[]common.Hash{SyncAbi.Events["StateSynced"].ID}},
		})
		if err != nil {
			return err
		}
		l.log.Debug("try scan logs success", "total", len(logs))
		if err := l.handleLogs(logs); err != nil {
			return err
		}

		if err := l.db.SetLastBlockNumber(new(big.Int).SetUint64(to)); err != nil {
			return err
		}
	}
	return nil
}

func (l *L1Sync) handleLogs(logs []types.Log) error {
	var events []*StateSender
	for _, log := range logs {
		event, err := UnpackLog(log)
		if err != nil {
			return err
		}
		events = append(events, event)
	}
	if len(events) != 0 {
		l.db.SetMaxSyncId(events[len(events)-1].Id)
	}
	if err := l.db.WriteStateSenderEvent(events); err != nil {
		return err
	}
	if l.updateCh != nil {
		l.updateCh <- struct{}{}
	}
	return nil
}

func UnpackLog(log types.Log) (*StateSender, error) {
	out := new(StateSender)
	event := SyncAbi.Events["StateSynced"]
	// Anonymous events are not supported.
	if len(log.Topics) == 0 {
		return nil, errors.New("no event signature")
	}
	if log.Topics[0] != event.ID {
		return nil, errors.New("event signature mismatch")
	}
	if len(log.Data) > 0 {
		v, err := event.Inputs.Unpack(log.Data)
		if err != nil {
			return nil, err
		}
		event.Inputs.Copy(out, v)
		if err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range event.Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	err := abi.ParseTopics(out, indexed, log.Topics[1:])
	if err != nil {
		return nil, err
	}
	return out, err
}
