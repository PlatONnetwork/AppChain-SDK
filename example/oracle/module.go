package oracle

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/example/election"
	contracts2 "github.com/PlatONnetwork/AppChain-SDK/example/election/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/example/oracle/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	types2 "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
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
	CallerAddress = common.BigToAddress(big.NewInt(201))
	RateAddr      = common.BigToAddress(big.NewInt(203))
	BlsVerifyAddr = common.BigToAddress(big.NewInt(204))
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
		key:         key,
		address:     crypto.PubkeyToAddress(key.PublicKey),
		db:          NewQCRateDB(store),
		extraVoteDb: extravote.NewExtraVoteDB(store),
		rateClient:  rateClient,
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
func (m *Module) Address() common.Address {
	return BlsVerifyAddr
}

func (m *Module) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	verify, _ := contracts.NewBlsVerify(evm, contract, readOnly)
	return verify.Run(input)
}
func (m *Module) ContractCreateBlockNumber(statedb vm.StateDBReader) uint64 {
	return 0
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
	fmt.Println(BlsVerifyAddr.Hex())
	err := caller2.WithCaller(CallerAddress).WithTo(RateAddr).DeployRate(BlsVerifyAddr, genesis.Decimals)
	if err != nil {
		return err
	}
	fmt.Println("election:", election.ProxyAddress.Hex())
	decimal, _ := caller2.Decimals()
	fmt.Println("decimal:", decimal)
	xx, err := caller2.GetLastRoundValidator()
	fmt.Println("xx:", xx)
	rns, _ := caller2.GetCurrentRoundValidator()
	fmt.Println("rns:", rns.Nodes)
	names, _ := caller2.FindNodes(big.NewInt(0))
	fmt.Println("names:", names)

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

}

func (m *Module) EndBlock(ctx sdk.WorkerContext) error {
	return nil
}

func (m *Module) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	m.db.InsertQCRate(block.Hash())
	return nil
}

func (m *Module) AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error) {
	nonce := ctx.StateDB().GetNonce(m.address)
	blockHash, rate, err := m.db.QCRate()
	if err != nil {
		return nil, nil
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
	caller, _ := contracts.NewRateGenesisCaller(ctx, ctx.StateDB(), m.chainConfig)
	num, _ := caller.WithCaller(CallerAddress).WithTo(RateAddr).BlockNumber()
	rns, err := caller.GetCurrentRoundValidator()
	fmt.Println(rns)
	rns2, err := caller.GetLastRoundValidator()
	fmt.Println(rns2)
	elec, _ := contracts2.NewElectionGenesisCaller(ctx, ctx.StateDB(), m.chainConfig)
	rnss, _ := elec.WithCaller(CallerAddress).WithTo(election.ProxyAddress).GetCurrentRoundValidator()
	fmt.Println(rnss)
	rnss2, _ := elec.WithCaller(CallerAddress).WithTo(election.ProxyAddress).GetLastRoundValidator()
	fmt.Println(rnss2)
	//if err := caller.WithCaller(CallerAddress).WithTo(RateAddr).Update(big.NewInt(int64(rate)), contracts.QuorumCert{
	//	Epoch:       qc.Epoch,
	//	ViewNumber:  qc.ViewNumber,
	//	BlockHash:   qc.BlockHash,
	//	BlockNumber: qc.BlockNumber,
	//	BlockIndex:  qc.BlockIndex,
	//	ExtendHash:  qc.ExtendHash,
	//}, qc.ValidatorSet.Bytes(), qc.Signature.Bytes(), big.NewInt(int64(index)), proof); err != nil {
	//	panic(err)
	//}

	tx, err := m.rateTxBuilder.WithNonce(nonce).Update(big.NewInt(int64(rate)), contracts.QuorumCert{
		Epoch:       qc.Epoch,
		ViewNumber:  qc.ViewNumber,
		BlockHash:   qc.BlockHash,
		BlockNumber: qc.BlockNumber,
		BlockIndex:  qc.BlockIndex,
		ExtendHash:  qc.ExtendHash,
	}, qc.ValidatorSet.Bytes(), qc.Signature.Bytes(), big.NewInt(int64(index)), proof)
	fmt.Println("---tx:", tx.Hash().Hex(), "num:", num)
	return map[common.Address]types.Transactions{
		m.address: {tx},
	}, nil
}
