package types

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
)

type EventSubscriber interface {
	GetLogFilters() map[common.Address][]common.Hash
	ProcessLog(header *coretypes.Header, log *coretypes.Log) error
}

type StateEvent interface {
	Subscribe(subscriber EventSubscriber)
}
