package nontxpool

import (
	"gopkg.in/urfave/cli.v1"
	"time"
)

type TxQueueConfig struct {
	GlobalTxCount uint64
}
type NonTxPoolConfig struct {
	TxQueueConfig
	EnableBroadcast   bool
	PendingLimit      int
	TxsPerAccount     int
	TxsCacheSize      int
	BroadcastInterval time.Duration
}

func InitConfig(ctx *cli.Context) *NonTxPoolConfig {
	txsCacheSize := ctx.GlobalInt(TxsCacheSizeFlag.Name)
	broadcastInterval := ctx.GlobalInt(BroadcastIntervalFlag.Name)
	return &NonTxPoolConfig{
		TxQueueConfig: TxQueueConfig{
			GlobalTxCount: ctx.GlobalUint64(GlobalTxCountFlag.Name),
		},
		EnableBroadcast:   ctx.GlobalBool(BroadcastFlag.Name),
		PendingLimit:      ctx.GlobalInt(PendingLimitFlag.Name),
		TxsPerAccount:     ctx.GlobalInt(TxsPerAccountFlag.Name),
		TxsCacheSize:      txsCacheSize,
		BroadcastInterval: time.Duration(broadcastInterval) * time.Millisecond,
	}
}
