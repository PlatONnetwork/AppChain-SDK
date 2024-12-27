package coupon

import (
	"encoding/json"
	"fmt"
	contracts2 "github.com/PlatONnetwork/AppChain-SDK/example/coupon/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
	"sync"
)

const ModuleVersion uint64 = 0
const ModuleName = "coupon"

var CouponAddress = common.BigToAddress(big.NewInt(102))
var CallerAddress = common.BigToAddress(big.NewInt(103))

type GenesisConfig struct {
	module.ModuleGenesisConfig
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}
type Module struct {
	sync.Mutex
	priority []common.Address
	cache    map[common.Address]int
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	var genesis GenesisConfig
	if err := json.Unmarshal(data, &genesis); err != nil {
		return err
	}
	caller, err := contracts2.NewCouponGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
	caller.WithCaller(CallerAddress).WithTo(CouponAddress).DeployCoupon(genesis.Name, genesis.Symbol)

	return nil
}
func (m *Module) SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	target := make([]types.Transactions, len(m.priority), len(m.priority))
	other := make(map[common.Address]types.Transactions)

	findFunc := func(txs map[common.Address]types.Transactions) {
		for k, v := range txs {
			pos, ok := m.cache[k]
			if ok {
				target[pos] = v
			} else {
				other[k] = v
			}
		}
	}
	findFunc(local)
	findFunc(remote)

	var txs types.Transactions
	for _, ts := range target {
		txs = append(txs, ts...)
	}
	for _, v := range other {
		txs = append(txs, v...)
	}
	for _, tx := range txs {
		fmt.Println(tx.Hash().Hex())
	}
	return txs, nil
}

func (m *Module) Set(priority []common.Address) {
	m.Lock()
	defer m.Unlock()
	m.cache = make(map[common.Address]int)
	m.priority = priority
	for i, p := range priority {
		m.cache[p] = i
	}
}
func (m *Module) Get() []common.Address {
	m.Lock()
	defer m.Unlock()
	return m.priority
}
func (m *Module) APIs() []rpc.API {
	return []rpc.API{
		rpc.API{
			Namespace: "coupon",
			Version:   "1",
			Service:   NewRPC(m),
			Public:    false,
		},
	}
}
