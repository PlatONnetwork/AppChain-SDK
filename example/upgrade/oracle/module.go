package oracle

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/PlatONnetwork/AppChain-SDK/example/oracle"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type Module struct {
	*oracle.Module
	extraVote *extravote.ExtraVote
	logger    log.Logger
}

func NewModule(store store.Store, key *ecdsa.PrivateKey, rateClient oracle.RateClient) *Module {
	return &Module{
		Module: oracle.NewModule(store, key, rateClient),
		logger: log.New("module", oracle.ModuleName),
	}
}
func (m *Module) SetExtraVote(extraVote *extravote.ExtraVote) {
	m.extraVote = extraVote
}
func (m *Module) RegistryUpgradeHandler(registrar module.UpgradeRegistrar) error {
	registrar.RegisterUpgradeHandler(oracle.ModuleName, 0, func(ctx sdk.WorkerContext) error {
		og, _ := json.Marshal(oracle.GenesisConfig{Decimals: 3, BlockNumber: ctx.Header().Number.Uint64()})

		if err := m.InitGenesis(ctx, ctx.StateDB(), ctx.Backend().ChainConfig(), og); err != nil {
			return err
		}
		m.extraVote.AddEnableVerifiers(oracle.ModuleName)
		m.logger.Info("Run upgrade handler success")
		return nil
	})
	return nil
}
