package sync

import (
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestDB(t *testing.T) {
	store := memorydb.New()
	db := NewL1SyncDB(store)
	var events []*StateSender
	for i := 0; i < 10; i++ {
		events = append(events, &StateSender{
			Id:       big.NewInt(int64(i)),
			Sender:   common.BigToAddress(big.NewInt(int64(i))),
			Receiver: common.BigToAddress(big.NewInt(int64(i))),
			Data:     []byte{byte(i)},
		})
	}
	err := db.WriteStateSenderEvent(events)
	require.Nil(t, err)
	res, err := db.FindStateSenderEvent(big.NewInt(1), big.NewInt(5))
	require.Nil(t, err)
	require.Equal(t, 5, len(res))

	for i := 0; i < 10; i++ {
		event, err := db.GetStateSenderEvent(big.NewInt(int64(i)))
		require.Nil(t, err)
		require.Equal(t, big.NewInt(int64(i)), event.Id)
		require.Equal(t, []byte{byte(i)}, event.Data)
	}
	db.SetLastBlockNumber(big.NewInt(10))
	n, err := db.LastBlockNumber()
	require.Nil(t, err)
	require.Equal(t, big.NewInt(10), n)
}
