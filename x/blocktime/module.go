package blocktime

import (
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"time"
)

const (
	ModuleName    = "blocktime"
	ModuleVersion = 0
)

type Module struct {
	nextBlockTime          time.Duration
	blockProductionTimeout time.Duration
}

func NewModule(ctx *cli.Context) *Module {
	return &Module{
		nextBlockTime:          time.Duration(ctx.GlobalUint64(NextBlockTimeFlag.Name)) * time.Millisecond,
		blockProductionTimeout: time.Duration(ctx.GlobalUint64(BlockProductionTimeoutFlag.Name)) * time.Millisecond,
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}

func (m *Module) CalcBlockDeadline(ctx sdk.ConsensusBlockTimeContext, timePoint time.Time) time.Time {
	blockTimeout := m.blockProductionTimeout
	if blockTimeout > ctx.ProduceInterval() {
		blockTimeout = ctx.ProduceInterval()
	}
	if ctx.Deadline().Sub(timePoint) > blockTimeout {
		return timePoint.Add(blockTimeout)
	}
	return ctx.Deadline()
}
func (m *Module) CalcNextBlockTime(ctx sdk.ConsensusBlockTimeContext, blockTime time.Time) time.Time {
	if m.nextBlockTime < ctx.ProduceInterval() {
		return blockTime.Add(m.nextBlockTime)
	}
	return blockTime.Add(ctx.ProduceInterval())
}
