package sync

import (
	"context"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
	"strings"
	"time"
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
}

func NewL1Sync(stateSenderAddr common.Address, url string, start *big.Int, db store.Store) (*L1Sync, error) {
	syncdb := NewL1SyncDB(db)
	log := log.New("l1sync")
	if url == "" {
		return &L1Sync{
			log: log,
			db:  syncdb,
		}, nil
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
	}, nil
}

func (l *L1Sync) Run(ctx context.Context) {
	if l.cli == nil {
		l.log.Info("l1 sync is not start")
		return
	}
	timer := time.AfterFunc(scanBlockInterval, func() {
		l.scanRootChain(ctx)
	})
	go func() {
		select {
		case <-ctx.Done():
			timer.Stop()
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
		event := SyncAbi.Events["StateSynced"]
		args, err := event.Inputs.Unpack(log.Data)
		if err != nil {
			return err
		}
		events = append(events, &StateSender{
			Id:       args[0].(*big.Int),
			Sender:   args[1].(common.Address),
			Receiver: args[2].(common.Address),
			Data:     args[3].([]byte),
		})
	}
	if len(events) != 0 {
		l.db.SetMaxSyncId(events[len(events)-1].Id)
	}
	if err := l.db.WriteStateSenderEvent(events); err != nil {
		return err
	}
	return nil
}
