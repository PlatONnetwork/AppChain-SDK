package vrf

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	"github.com/PlatONnetwork/AppChain-SDK/x/util"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf/config"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf/contracts"
	vrfdb "github.com/PlatONnetwork/AppChain-SDK/x/vrf/db"
	vrftypes "github.com/PlatONnetwork/AppChain-SDK/x/vrf/types"
	vrfwrap "github.com/PlatONnetwork/AppChain-SDK/x/vrf/wrap"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"reflect"

	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
)

var (
	NonceStorageKey = []byte("nonceStorageKey")
)

type VRFModule struct {
	logger         log.Logger
	configParams   *config.VRFNetworkParams
	nodePrivateKey *ecdsa.PrivateKey
	stageModule    vrftypes.StageModuler
	stakeModule    vrftypes.StakeModuler
}

func NewVRFModule(ctx *cli.Context, stage vrftypes.StageModuler) *VRFModule {
	return &VRFModule{
		logger:         log.New("module", "vrf"),
		nodePrivateKey: l1.DecodeNodePrivateKey(ctx),
		configParams:   config.DefualtVRFNetworkParams(),
		stageModule:    stage,
	}
}

func (v *VRFModule) SetStakeModule(stake vrftypes.StakeModuler) {
	v.stakeModule = stake
}

func (v *VRFModule) Name() string {
	return "staking"
}
func (v *VRFModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) {
	var conf config.VRFNetworkParams
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON VRFNetworkParams bytes", "error", err)
	}
	if err := json.Unmarshal(raw, &conf); nil != err {
		log.Error("Failed UnmarshalJSON VRFNetworkParams", "error", err)
	} else {
		v.configParams = &conf
	}

	// set genesis vrf nonce (32 byte)
	vrfdb.SetNonceAndProof(db, v.Address(), 0, v.configParams.GenesisVRFNonce.Bytes())

	log.Info("Succeed init genesis", "module", v.Name(), "VRFNetworkParams", string(raw))
}

func (v *VRFModule) Address() basecommon.Address {
	return constants.VRFHandlerAddress
}

func (v *VRFModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	vrfHandler, _ := contracts.NewVRFHandler(evm, contract, readOnly)
	vrfHandler.SetStageModule(v.stageModule)
	vrfHandler.SetStakeModule(v.stakeModule)
	return vrfHandler.Run(input)
}

func (v *VRFModule) AddTxs(ctx sdk.WorkerContext, local, remote map[basecommon.Address]types.Transactions) (map[basecommon.Address]types.Transactions, map[basecommon.Address]types.Transactions) {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		v.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return local, remote
	}

	blockNumber := ctx.Backend().CurrentHeader().Number.Uint64()
	from := crypto.PubkeyToAddress(v.nodePrivateKey.PublicKey)

	// generate nonceAndProof by validator pubKey and previousNonce
	// nonAndProof: 81 byte
	// flag |nonce |proof
	// 1byte|32byte|48byte
	nonceAndProof, err := v.GenerateNonceAndProof(wctx, blockNumber)
	if nil != err {
		return local, remote
	}

	pushNonceAndProofTx, err := v.createPushNonceAndProofTx(ctx, nonceAndProof)
	if nil != err {
		v.logger.Error("Failed to create pushNonceAndProof tx", "blockNumber", blockNumber, "error", err)
		return local, remote
	}
	if nil == local[from] {
		local[from] = make(types.Transactions, 0)
	}
	local[from] = append(local[from], pushNonceAndProofTx)
	return local, remote
}

func (v *VRFModule) BeginBlock(ctx sdk.WorkerContext) {

}
func (v *VRFModule) EndBlock(ctx sdk.WorkerContext) {

	header := ctx.Backend().CurrentHeader()

	// not worker validator
	if util.IsNotWorker(header) {
		currentBlock := header.Number.Uint64()

		// get nonceAndProof by block (After the `pushNonceAndProof` transaction was executed)
		nonceAndProof, err := vrfwrap.GetCurrentNonceAndProof(ctx.StateDB(), v.Address(), currentBlock)
		if nil != err {
			panic(fmt.Sprintf("Failed to get current nonceAndProof, blockNumber: %d, error: %s", currentBlock, err))
		}

		// Extract the validator public key of the build block based on the signature in the block header
		sign := header.Signature()
		sealhash := header.SealHash().Bytes()
		pk, err := crypto.SigToPub(sealhash, sign)
		if err != nil {
			panic(fmt.Sprintf("Failed to handle sigToPub, blockNumber: %d, error: %s", currentBlock, err))
		}

		// verify nonce and
		if err := v.VerifyVrf(ctx, currentBlock, nonceAndProof, pk); nil != err {
			panic(fmt.Sprintf("Failed to verify vrf nonce and proof, blockNumber: %d, error: %s", currentBlock, err))
		}
	}
}

func (v *VRFModule) GenerateNonceAndProof(ctx sdk.WorkerContext, blockNumber uint64) ([]byte, error) {
	nonceAndProof, err := vrfwrap.GenerateNonceAndProof(ctx.StateDB(), v.Address(), blockNumber, v.nodePrivateKey)
	if nil != err {
		v.logger.Error(err.Error(), "blockNumber", blockNumber)
		return nil, err
	}
	v.logger.Info("Succeed to generate vrf nonce and proof", "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof),
		"nodeId", enode.PublicKeyToIDv0(&(v.nodePrivateKey.PublicKey)).String())
	return nonceAndProof, nil
}

func (v *VRFModule) VerifyVrf(ctx sdk.WorkerContext, blockNumber uint64, nonceAndProof []byte, key *ecdsa.PublicKey) error {

	previousNonce, err := vrfwrap.GetPreviousNonce(ctx.StateDB(), v.Address(), blockNumber)
	if nil != err {
		v.logger.Error("Failed to get previous vrf nonce", "blockNumber", blockNumber, "error", err)
		return err
	}

	if err := vrfwrap.VerifyVrf(nonceAndProof, previousNonce, key); nil != err {
		v.logger.Error(err.Error(), "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof), "data", previousNonce.Hex())
		return err
	}
	v.logger.Info("Succeed to verify vrf nonceAndProof", "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof), "data", previousNonce.Hex())
	return nil
}

func (v *VRFModule) createPushNonceAndProofTx(ctx sdk.WorkerContext, nonceAndProof []byte) (*types.Transaction, error) {

	input, err := contracts.Abi.Methods["pushNonceAndProof"].Inputs.Pack(nonceAndProof)
	if nil != err {
		return nil, err
	}
	from := crypto.PubkeyToAddress(v.nodePrivateKey.PublicKey)
	txNonce, err := ctx.Backend().GetPoolNonce(from)
	if nil != err {
		return nil, err
	}

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

func (v *VRFModule) GetNonceQueueUtil(stateDB sdk.StateDBReader, blockNumber, size uint64) ([]basecommon.Hash, error) {
	return vrfwrap.GetNonceQueueUtil(stateDB, v.Address(), blockNumber, size)
}

func (v *VRFModule) GetCurrentNonce(stateDB sdk.StateDBReader, blockNumber uint64) (basecommon.Hash, error) {
	return vrfwrap.GetCurrentNonce(stateDB, v.Address(), blockNumber)
}
