package blacklist

import (
	"errors"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"sync"
)

const ModuleVersion uint64 = 0
const ModuleName = "blacklist"

type Module struct {
	sync.Mutex
	signer    types.Signer
	blacklist map[common.Address]struct{}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}

func (m *Module) Init(ctx sdk.InitContext) error {
	chainId, err := ctx.Backend().ChainId()
	if err != nil {
		return err
	}
	m.signer = types.NewLondonSigner(chainId)
	m.blacklist = make(map[common.Address]struct{})
	return nil
}
func (m *Module) CheckTx(ctx sdk.Context, tx *types.Transaction) error {
	addr := tx.FromAddr(m.signer)
	if _, ok := m.blacklist[addr]; ok {
		return errors.New("blacklist address")
	}
	return nil
}
func (m *Module) FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	return txs
}
func (m *Module) Set(addrs []common.Address) {
	m.Lock()
	defer m.Unlock()
	for _, addr := range addrs {
		m.blacklist[addr] = struct{}{}
	}
}
func (m *Module) Get() []common.Address {
	m.Lock()
	defer m.Unlock()
	var list []common.Address
	for addr := range m.blacklist {
		list = append(list, addr)
	}
	return list
}
func (m *Module) APIs() []rpc.API {
	return []rpc.API{
		rpc.API{
			Namespace: "blacklist",
			Version:   "1",
			Service:   NewRPC(m),
			Public:    false,
		},
	}
}
