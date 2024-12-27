package oracle

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/example/oracle/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	types2 "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
	"sync"
)

const ModuleVersion uint64 = 0
const ModuleName = "oracle"

var (
	CallerAddress = common.BigToAddress(big.NewInt(103))
	ElectionAddr  = common.BigToAddress(big.NewInt(101))
	RateAddr      = common.BigToAddress(big.NewInt(111))
)

type GenesisConfig struct {
	Decimals uint8 `json:"decimals,omitempty"`
}

type Module struct {
	sync.Mutex
	chainConfig   *params.ChainConfig
	key           *ecdsa.PrivateKey
	address       common.Address
	chainId       *big.Int
	db            *QCRateDB
	rateTxBuilder *contracts.RateTxBuilder
	extraVoteDb   *extravote.ExtraVoteDB
	rateClient    RateClient
}

func NewModule(store store.Store, key *ecdsa.PrivateKey, rateClient RateClient) *Module {
	return &Module{
		key:        key,
		address:    crypto.PubkeyToAddress(key.PublicKey),
		db:         NewQCRateDB(store),
		rateClient: rateClient,
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
func (m *Module) Init(ctx sdk.InitContext) error {
	m.chainConfig = ctx.Backend().ChainConfig()
	var err error
	m.chainId, err = ctx.Backend().ChainId()
	if err != nil {
		return err
	}
	m.rateTxBuilder, err = contracts.NewRateTxBuilder(RateAddr, m.key, m.chainId)
	if err != nil {
		return err
	}
	return nil
}

func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	var genesis GenesisConfig
	if err := json.Unmarshal(data, &genesis); err != nil {
		return err
	}
	caller2, _ := contracts.NewRateGenesisCaller(ctx, db, chainConfig)
	caller2.WithCaller(CallerAddress).WithTo(RateAddr).DeployRate(ElectionAddr, genesis.Decimals)

	return nil
}

func (m *Module) ExtendData(ctx sdk.ConsensusContext) []byte {
	rate, err := m.rateClient.Rate()
	if err != nil {
		return nil
	}
	data, _ := rlp.EncodeToBytes(rate)
	return data
}

func (m *Module) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	expectRate, err := m.rateClient.Rate()
	if err != nil {
		return errors.New("request rate failed")
	}
	var rate uint64
	rlp.DecodeBytes(data, &rate)
	if rate != expectRate {
		return errors.New(fmt.Sprintf("invalid rate, expect:%d, actual:%d", expectRate, rate))
	}
	return m.db.InsertRate(ctx.Header().Hash(), rate)
}

func (m *Module) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	m.db.InsertQCRate(block.Block.Hash())
}

func (m *Module) AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error) {
	nonce := ctx.StateDB().GetNonce(m.address)
	blockHash, rate, err := m.db.QCRate()
	if err != nil {
		return nil, err
	}
	block := ctx.Backend().GetBlockByHash(blockHash)
	if block == nil {
		return nil, errors.New(fmt.Sprintf("get block failed:%s", blockHash.Hex()))
	}
	_, qc, err := types2.DecodeExtra(block.ExtraData())
	if err != nil {
		return nil, err
	}
	leaf, _ := rlp.EncodeToBytes(rate)
	index, proof, err := m.extraVoteDb.GetProof(qc.Epoch, qc.ViewNumber, qc.BlockIndex, leaf)
	if err != nil {
		return nil, err
	}
	tx, err := m.rateTxBuilder.WithNonce(nonce).Update(big.NewInt(int64(rate)), contracts.QuorumCert{
		Epoch:       qc.Epoch,
		ViewNumber:  qc.ViewNumber,
		BlockHash:   qc.BlockHash,
		BlockNumber: qc.BlockNumber,
		BlockIndex:  qc.BlockIndex,
		ExtendHash:  qc.ExtendHash,
	}, qc.ValidatorSet.Bytes(), qc.Signature.Bytes(), big.NewInt(int64(index)), proof)
	return map[common.Address]types.Transactions{
		m.address: {tx},
	}, nil
}
