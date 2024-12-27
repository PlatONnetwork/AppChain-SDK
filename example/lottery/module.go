package lottery

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	contracts2 "github.com/PlatONnetwork/AppChain-SDK/example/coupon/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/example/lottery/contracts"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/vrf"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
	"sync"
)

const ModuleVersion uint64 = 0
const ModuleName = "lottery"

var (
	CallerAddress = common.BigToAddress(big.NewInt(103))

	VRFSystemAddr  = common.BigToAddress(big.NewInt(111))
	VRFStorageAddr = common.BigToAddress(big.NewInt(222))
	LotteryAddr    = common.BigToAddress(big.NewInt(333))
)

type GenesisConfig struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Nonce  string `json:"nonce"` //0x03f3b376f00863de14440eff826835d16ffa3c8b0fc7ad1402beee7ccf076aa9282f795c3e92d2f61e45e85abe5fcfef134ae6700a50b0885942a92d92b9c88a280450a416880be13a23e449f41ef12f43
}
type Module struct {
	sync.Mutex
	chainConfig         *params.ChainConfig
	key                 *ecdsa.PrivateKey
	address             common.Address
	chainId             *big.Int
	lotteryTxBuilder    *contracts.LotteryTxBuilder
	vrfStorageTxBuilder *contracts.VRFStorageTxBuilder
}

func NewModule(key *ecdsa.PrivateKey) *Module {
	return &Module{
		key:     key,
		address: crypto.PubkeyToAddress(key.PublicKey),
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
	m.lotteryTxBuilder, err = contracts.NewLotteryTxBuilder(LotteryAddr, m.key, m.chainId)
	if err != nil {
		return err
	}
	m.vrfStorageTxBuilder, err = contracts.NewVRFStorageTxBuilder(VRFStorageAddr, m.key, m.chainId)
	if err != nil {
		return err
	}
	return nil
}

func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	db.SetNonce(VRFSystemAddr, 1)
	var genesis GenesisConfig
	if err := json.Unmarshal(data, &genesis); err != nil {
		return err
	}

	caller2, err := contracts2.NewCouponGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
	caller2.WithCaller(CallerAddress).WithTo(common.BigToAddress(big.NewInt(1212))).DeployCoupon(genesis.Name, genesis.Symbol)

	caller, err := contracts.NewVRFStorageGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
	nonceProof := hexutil.MustDecode(genesis.Nonce)
	caller.WithCaller(CallerAddress).WithTo(VRFStorageAddr).DeployVRFStorage(VRFSystemAddr, nonceProof)
	lottery, err := contracts.NewLotteryGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
	proof, _ := caller.GetNonceProof(big.NewInt(0))
	if !bytes.Equal(nonceProof, proof) {
		panic(err)
	}
	addr, _ := caller.VrfAddr()
	fmt.Println(addr.Hex())
	lottery.WithCaller(CallerAddress).WithTo(LotteryAddr).DeployLottery(VRFStorageAddr, genesis.Name, genesis.Symbol)

	return nil
}

func (m *Module) AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error) {
	nonce := ctx.StateDB().GetNonce(m.address)
	vrfTx, err := m.generateVRFProofTx(ctx, nonce)
	if err != nil {
		return nil, err
	}
	drawTx, err := m.lotteryTxBuilder.WithNonce(nonce + 1).Drawing()
	if err != nil {
		return nil, err
	}

	return map[common.Address]types.Transactions{
		m.address: {vrfTx, drawTx},
	}, nil
}
func (m *Module) generateVRFProofTx(ctx sdk.WorkerContext, nonce uint64) (*types.Transaction, error) {
	blockNumber := ctx.Header().Number
	caller, err := contracts.NewVRFStorageGenesisCaller(ctx, ctx.StateDB(), m.chainConfig)
	if err != nil {
		return nil, err
	}
	previous, err := caller.WithCaller(CallerAddress).WithTo(VRFStorageAddr).GetNonceProof(new(big.Int).Sub(blockNumber, big.NewInt(1)))
	if err != nil {
		return nil, err
	}
	nonceProof, err := vrf.Prove(m.key, vrf.ProofToHash(previous))
	if err != nil {
		return nil, err
	}

	tx, err := m.vrfStorageTxBuilder.WithNonce(nonce).AddProve(crypto.FromECDSAPub(&m.key.PublicKey), nonceProof)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (m *Module) Address() common.Address {
	return VRFSystemAddr
}

func (m *Module) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	vrfManager, _ := contracts.NewVRF(evm, contract, readOnly)
	return vrfManager.Run(input)
}

func (m *Module) ContractCreateBlockNumber(statedb sdk.StateDBReader) uint64 {
	return 0
	vrf, _ := contracts.NewVRF(sdkcontracts.NewEVM(types.NewStateDBWrapper(statedb), big.NewInt(0)), sdkcontracts.NewContract(m, m), false)
	return vrf.GetCreateBlock()
}
