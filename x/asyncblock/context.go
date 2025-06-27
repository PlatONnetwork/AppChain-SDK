package asyncblock

import (
	"context"
	coresdk "github.com/PlatONnetwork/PlatON-Go/core/sdk"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"time"
)

type asyncWorkerContext struct {
	parent      *types.Block
	statedb     coresdk.StateDB
	backend     coresdk.Backend
	header      *types.Header
	chainConfig *params.ChainConfig
	vmConfig    *vm.Config
}

func NewAsyncWorkerContext(parent *types.Block,
	statedb coresdk.StateDB,
	backend coresdk.Backend,
	header *types.Header,
	chainConfig *params.ChainConfig,
	vmConfig *vm.Config) *asyncWorkerContext {
	return &asyncWorkerContext{
		parent:      parent,
		statedb:     statedb,
		backend:     backend,
		header:      header,
		chainConfig: chainConfig,
		vmConfig:    vmConfig,
	}
}

func (a *asyncWorkerContext) Context() context.Context {
	return context.Background()
}

func (a *asyncWorkerContext) Backend() coresdk.Backend {
	return a.backend
}

func (a *asyncWorkerContext) StateDB() coresdk.StateDB {
	return a.statedb
}

func (a *asyncWorkerContext) Header() *types.Header {
	return a.header
}

func (a *asyncWorkerContext) IsWorker() bool {
	return false
}

func (a *asyncWorkerContext) ParentBlock() *types.Block {
	return a.parent
}

func (a *asyncWorkerContext) ChainConfig() *params.ChainConfig {
	return a.chainConfig
}

func (a *asyncWorkerContext) VMConfig() *vm.Config {
	return a.vmConfig
}

func (a asyncWorkerContext) BlockDeadline() time.Time {
	return time.Time{}
}
