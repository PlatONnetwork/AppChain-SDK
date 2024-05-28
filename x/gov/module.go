package gov

import (
	"crypto/ecdsa"
	"encoding/json"
	common2 "github.com/PlatONnetwork/AppChain-SDK/common"
	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/utils"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	contracts2 "github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/message"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"sync"
)

const (
	ModuleName           = "governance"
	ModuleVersion uint64 = 0
)

type GenesisParams struct {
	Name              string         `json:"name"`
	Version           string         `json:"version"`
	VoteDelay         *big.Int       `json:"voteDelay"`
	VotePeriod        *big.Int       `json:"votePeriod"`
	QuorumNumerator   *big.Int       `json:"quorumNumerator"`
	ProposalThreshold *big.Int       `json:"proposalThreshold"`
	Owner             common.Address `json:"owner"`
	VoteToken         common.Address `json:"voteToken"`
}
type Module struct {
	sync.Mutex
	logger       log.Logger
	keystoreFile string
	passwordFile string
	privateKey   *ecdsa.PrivateKey
	proposals    map[*big.Int]*Proposal
}

func NewModule(ctx *cli.Context) (*Module, error) {
	return &Module{
		logger:       log.New("module", "governance"),
		keystoreFile: ctx.GlobalString(utils.KeystoreFlag.Name),
		passwordFile: ctx.GlobalString(utils.PasswordFlag.Name),
	}, nil
}

func (g *Module) Name() string {
	return ModuleName
}

func (g *Module) Version() uint64 {
	return ModuleVersion
}

func (g *Module) Address() common.Address {
	return constants.GovAddress
}

func (g *Module) Init(ctx sdk.InitContext) error {
	key, err := utils.DecodePrivateKey(g.keystoreFile, g.passwordFile)
	if err != nil {
		return err
	}
	g.privateKey = key.PrivateKey
	return nil
}

func (g *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, raw json.RawMessage) error {
	var params GenesisParams
	if err := json.Unmarshal(raw, &params); nil != err {
		log.Error("Failed UnmarshalJSON RewardNetworkParams", "error", err)
		return err
	}
	gov, _ := contracts2.NewGovernance(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(g, g), false)
	gov.Init(params.Name, params.Version, params.VoteDelay, params.VotePeriod, params.QuorumNumerator, params.ProposalThreshold, params.Owner, params.VoteToken)
	return nil
}

func (g *Module) AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error) {
	from := crypto.PubkeyToAddress(g.privateKey.PublicKey)

	nonce := common2.EnableNonce(local[from], func() uint64 {
		return ctx.StateDB().GetNonce(from)
	})
	caller, err := g.newGovCallContract(ctx, ctx.Header())
	if err != nil {
		return nil, err
	}
	g.Lock()
	defer g.Unlock()
	for _, proposal := range g.proposals {
		g.logger.Debug("handle proposal", "proposal", proposal)
		delete(g.proposals, proposal.ProposalId)

		state, err := caller.State(proposal.ProposalId)
		if err != nil {
			g.logger.Error("get proposal state failed", "proposal", proposal)
			continue
		}
		if state != contracts2.Succeeded {
			g.logger.Info("proposal state is error", "proposal", proposal, "state", contracts2.StateToString(state), "expect", "Succeeded")
			continue
		}
		tx, err := g.createExecuteTx(ctx, nonce, proposal)
		if err != nil {
			g.logger.Error("create execute proposal tx failed", "proposal", proposal)
			continue
		}
		if local[from] == nil {
			local[from] = types.Transactions{}
		}
		local[from] = append(local[from], tx)
		nonce += 1
	}
	return local, nil
}

func (g *Module) APIs() []rpc.API {
	return []rpc.API{
		rpc.API{
			Namespace: "gov",
			Version:   "1",
			Service:   NewRpcService(g),
			Public:    false,
		},
	}
}

func (g *Module) getProposals() []*Proposal {
	g.Lock()
	defer g.Unlock()
	var proposals []*Proposal
	for _, v := range g.proposals {
		proposals = append(proposals, v)
	}
	return proposals
}

func (g *Module) addProposal(proposal *Proposal) {
	g.Lock()
	defer g.Unlock()
	g.proposals[proposal.ProposalId] = proposal
}

func (g *Module) newGovCallContract(ctx sdk.Context, header *types.Header) (*contracts2.GovernanceCaller, error) {
	from := crypto.PubkeyToAddress(g.privateKey.PublicKey)
	evm, _, err := ctx.Backend().GetEVM(message.NewOnlyCallMessage(from), header)
	if err != nil {
		return nil, err
	}
	return contracts2.NewGovernanceCaller(evm, vm.NewContract(vm.AccountRef(from), vm.AccountRef(constants.StateSyncAddress), big.NewInt(0), 1000000), constants.GovAddress)
}

func (g *Module) createExecuteTx(ctx sdk.Context, nonce uint64, proposal *Proposal) (*types.Transaction, error) {
	method := contracts2.Abi.Methods["execute"]
	input, err := method.Inputs.Pack(proposal.Targets, proposal.Values, proposal.Calldatas, proposal.DescriptionHash)
	if err != nil {
		return nil, err
	}

	input = append(method.ID, input...)
	tx := types.NewTransaction(nonce, constants.StateSyncAddress, nil, 3000000, big.NewInt(0), input)
	chainId, _ := ctx.Backend().ChainId()
	signer := types.NewEIP155Signer(chainId)
	tx, err = types.SignTx(tx, signer, g.privateKey)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
