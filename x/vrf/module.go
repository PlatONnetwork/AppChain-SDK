package vrf

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"

	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"

	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf/config"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf/contracts"
	vrftypes "github.com/PlatONnetwork/AppChain-SDK/x/vrf/types"
	vrfwrap "github.com/PlatONnetwork/AppChain-SDK/x/vrf/wrap"
)

const (
	ModuleName           = "vrf"
	ModuleVersion uint64 = 1
)

var (
	_ module.ContractModule = (*VRFModule)(nil)

	NonceStorageKey = []byte("nonceStorageKey")
)

type VRFModule struct {
	logger         log.Logger
	nodePrivateKey *ecdsa.PrivateKey
	stageModule    vrftypes.StageModuler
	stakeModule    vrftypes.StakeModuler
}

func NewModule(ctx *cli.Context, stage vrftypes.StageModuler) *VRFModule {
	return &VRFModule{
		logger:      log.New("module", ModuleName),
		stageModule: stage,
	}
}

func (v *VRFModule) SetStakeModule(stake vrftypes.StakeModuler) {
	v.stakeModule = stake
}

func (v *VRFModule) Name() string {
	return ModuleName
}

func (v *VRFModule) Version() uint64 {
	return ModuleVersion
}

func (v *VRFModule) Init(ctx sdk.InitContext) error {
	v.nodePrivateKey = ctx.NodeKey()
	return nil
}

func (v *VRFModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {

	configParams := config.DefualtVRFNetworkParams()
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON VRFNetworkParams bytes", "error", err)
		return err
	}

	var conf config.VRFNetworkParams
	if err := json.Unmarshal(raw, &conf); nil != err {
		log.Error("Failed UnmarshalJSON VRFNetworkParams", "error", err)
		return err
	} else {
		configParams = &conf
	}
	// init vrf manager  account nonce
	initAccountNonce(db, v.Address())
	initGenesisVRFNonce(db, v.Address(), chainConfig, configParams)
	log.Info("Succeed init genesis", "module", v.Name(), "VRFNetworkParams", configParams.String())
	return nil
}

func (v *VRFModule) Address() basecommon.Address {
	return constants.VRFManagerAddress
}

func (v *VRFModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	vrfManager, _ := contracts.NewVRFManager(evm, contract, readOnly)
	vrfManager.SetStageModule(v.stageModule)
	vrfManager.SetStakeModule(v.stakeModule)
	return vrfManager.Run(input)
}

func (v *VRFModule) ContractCreateBlockNumber(statedb sdk.StateDBReader) uint64 {
	// TODO: implement me
	return 0
}

func (v *VRFModule) AddTxs(ctx sdk.WorkerContext, local map[basecommon.Address]types.Transactions) (map[basecommon.Address]types.Transactions, error) {

	start := time.Now()

	blockNumber := ctx.Header().Number.Uint64()
	if blockNumber == 0 {
		return local, nil
	}
	from := crypto.PubkeyToAddress(v.nodePrivateKey.PublicKey)

	// generate nonceAndProof by validator pubKey and previousNonce
	// nonAndProof: 81 byte
	// flag |nonce |proof
	// 1byte|32byte|48byte
	nonceAndProof, err := v.GenerateNonceAndProof(ctx, blockNumber)
	if nil != err {
		return local, err
	}

	txNonce := common.EnableNonce(local[from], func() uint64 {
		return ctx.StateDB().GetNonce(from)
	})

	pushNonceAndProofTx, err := v.createPushNonceAndProofTx(ctx, nonceAndProof, txNonce)
	if nil != err {
		v.logger.Error("Failed to create pushNonceAndProof tx", "blockNumber", blockNumber, "error", err)
		return local, err
	}
	if nil == local[from] {
		local[from] = make(types.Transactions, 0)
	}
	local[from] = append(local[from], pushNonceAndProofTx)
	end := time.Now()
	v.logger.Debug("create pushNonceAndProof tx", "blockNumber", blockNumber, "start", basecommon.Millis(start), "end", basecommon.Millis(end), "duration", end.Sub(start), "txHash", pushNonceAndProofTx.Hash().Hex(), "from", from.Hex())
	return local, nil
}

func (v *VRFModule) EndBlock(ctx sdk.WorkerContext) error {
	header := ctx.Header()

	// not worker validator
	if !ctx.IsWorker() {
		currentBlock := header.Number.Uint64()

		// get nonceAndProof by block (After the `pushNonceAndProof` transaction was executed)
		nonceAndProof, err := vrfwrap.GetCurrentNonceAndProof(ctx.StateDB(), v.Address(), currentBlock)
		if nil != err {
			return fmt.Errorf("can not get current nonceAndProof, %s", err)
		}

		// Extract the validator public key of the build block based on the signature in the block header
		sign := header.Signature()
		sealhash := header.SealHash().Bytes()
		pk, err := crypto.SigToPub(sealhash, sign)
		if err != nil {
			return fmt.Errorf("can not handle sigToPub, %s", err)
		}

		// verify nonce and
		if err := v.VerifyVrf(ctx, currentBlock, nonceAndProof, pk); nil != err {
			return fmt.Errorf("can not verify vrf nonce and proof, %s", err)
		}
	}
	return nil
}

func (v *VRFModule) GenerateNonceAndProof(ctx sdk.WorkerContext, blockNumber uint64) ([]byte, error) {
	start := time.Now()
	nonceAndProof, err := vrfwrap.GenerateNonceAndProof(ctx.StateDB(), v.Address(), blockNumber, v.nodePrivateKey)
	if nil != err {
		v.logger.Error("Failed to generate vrf nonceAndProof", "blockNumber", blockNumber, "error", err)
		return nil, err
	}
	end := time.Now()
	v.logger.Info("Succeed to generate vrf nonceAndProof", "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof),
		"nodeId", enode.PublicKeyToIDv0(&(v.nodePrivateKey.PublicKey)).String(), "start", basecommon.Millis(start), "end", basecommon.Millis(end), "duration", end.Sub(start))
	return nonceAndProof, nil
}

func (v *VRFModule) VerifyVrf(ctx sdk.WorkerContext, blockNumber uint64, nonceAndProof []byte, key *ecdsa.PublicKey) error {

	previousNonce, err := vrfwrap.GetPreviousNonce(ctx.StateDB(), v.Address(), blockNumber)
	if nil != err {
		v.logger.Error("Failed to get previous vrf nonce", "blockNumber", blockNumber, "error", err)
		return err
	}

	if err := vrfwrap.VerifyVrf(nonceAndProof, previousNonce, key); nil != err {
		v.logger.Error("Failed to verify vrf", "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof), "data", previousNonce.Hex(), "error", err)
		return err
	}
	v.logger.Info("Succeed to verify vrf nonceAndProof", "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof), "data", previousNonce.Hex())
	return nil
}

func (v *VRFModule) createPushNonceAndProofTx(ctx sdk.WorkerContext, nonceAndProof []byte, txNonce uint64) (*types.Transaction, error) {

	method := contracts.Abi.Methods["pushNonceAndProof"]

	input, err := method.Inputs.Pack(nonceAndProof)
	if nil != err {
		return nil, err
	}
	input = append(method.ID, input...)

	tx := types.NewTransaction(txNonce, v.Address(), nil, 100000, big.NewInt(0), input)
	chainId, _ := ctx.Backend().ChainId()
	signer := types.NewEIP155Signer(chainId)
	tx, err = types.SignTx(tx, signer, v.nodePrivateKey)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// extern

func (v *VRFModule) GetNonceQueueFromTail(stateDB sdk.StateDBReader, blockNumber, size uint64) ([]basecommon.Hash, error) {
	return vrfwrap.GetNonceQueueFromTail(stateDB, v.Address(), blockNumber, size)
}

func (v *VRFModule) GetCurrentNonce(stateDB sdk.StateDBReader, blockNumber uint64) (basecommon.Hash, error) {
	return vrfwrap.GetCurrentNonce(stateDB, v.Address(), blockNumber)
}
