package sync

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
)

type SimL1Sync struct {
	L1Sync *L1Sync
}

func MakeL1Sync(stateSenderAddr common.Address, cli PlatonClient, start *big.Int, db store.Store, updateCh chan struct{}) (*SimL1Sync, error) {
	syncdb := NewL1SyncDB(db)
	log := log.New("l1sync")

	number, err := syncdb.LastBlockNumber()
	if err != nil {
		return nil, err
	}

	if start == nil && number == nil {
		return nil, errors.New("need set start block number")
	}
	return &SimL1Sync{
		L1Sync: &L1Sync{
			log:             log,
			stateSenderAddr: stateSenderAddr,
			db:              syncdb,
			cli:             cli,
			updateCh:        updateCh,
		},
	}, nil
}

func (s *SimL1Sync) Scan(ctx context.Context, start, end uint64) error {
	return s.L1Sync.scanLogs(ctx, start, end)
}

type SimClient struct {
	id     *big.Int
	Number uint64
}

func NewSimClient() *SimClient {
	return &SimClient{
		id:     big.NewInt(1),
		Number: 1,
	}
}

func (s *SimClient) BlockNumber(ctx context.Context) (uint64, error) {
	return s.Number, nil
}

func (s *SimClient) FilterLogs(ctx context.Context, q platon.FilterQuery) ([]types.Log, error) {
	var logs []types.Log
	topics, err := abi.PackTopics(SyncAbi.Events["StateSynced"].Inputs, s.id, common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2)))
	topics = append([]common.Hash{SyncAbi.Events["StateSynced"].ID}, topics...)
	var input abi.Arguments
	for _, p := range SyncAbi.Events["StateSynced"].Inputs {
		if !p.Indexed {
			input = append(input, p)
		}
	}
	//data, err := input.Pack([]byte{1, 2, 3, 4})
	data, err := PackLogData(s.id, common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2)), []byte{1, 2, 3, 4})
	if err != nil {
		return nil, err
	}
	fmt.Println(hex.EncodeToString(data))
	logs = append(logs, types.Log{
		Address:     common.Address{},
		Topics:      topics,
		Data:        data,
		BlockNumber: q.ToBlock.Uint64(),
		TxHash:      common.BigToHash(q.ToBlock),
		TxIndex:     0,
		BlockHash:   common.BigToHash(q.ToBlock),
		Index:       0,
		Removed:     false,
	})
	s.Number = q.ToBlock.Uint64()
	return logs, nil
}

func PackLogData(args ...interface{}) ([]byte, error) {
	var input abi.Arguments
	var needArgs []interface{}
	for i, p := range SyncAbi.Events["StateSynced"].Inputs {
		if !p.Indexed {
			input = append(input, p)
			needArgs = append(needArgs, args[i])
		}
	}
	return input.Pack(needArgs...)
}
