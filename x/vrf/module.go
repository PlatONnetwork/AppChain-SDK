package vrf

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	sdkcommon "github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/message"
	"github.com/PlatONnetwork/AppChain-SDK/x/vrf/contracts"
	vrfdb "github.com/PlatONnetwork/AppChain-SDK/x/vrf/db"
	vrfInternal "github.com/PlatONnetwork/AppChain-SDK/x/vrf/internal"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"reflect"

	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

var (
	NonceStorageKey = []byte("nonceStorageKey")
)

type VRFModule struct {
	keystoreFile string
	passwordFile string
	logger       log.Logger
	privateKey   *ecdsa.PrivateKey
}

func NewVRFModule() *VRFModule {
	return &VRFModule{
		logger: log.New("module", "vrf"),
	}
}

func (v *VRFModule) Name() string {
	return "staking"
}

func (v *VRFModule) Init() error {

	key, err := decodePrivateKey(v.keystoreFile, v.passwordFile)
	if err != nil {
		return err
	}
	v.privateKey = key.PrivateKey
	return nil
}

func (v *VRFModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) {
	// TODO 初始化 vrf nonce
	vrfdb.SetNonceAndProof(db, address.VRFHandlerAddress, 0, []byte("genesisVRFNonce"))
}

func (v *VRFModule) AddTxs(ctx sdk.Context, local, remote map[basecommon.Address]types.Transactions) (map[basecommon.Address]types.Transactions, map[basecommon.Address]types.Transactions) {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		v.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return local, remote
	}

	blockNumber := ctx.Backend().CurrentHeader().Number.Uint64()
	from := crypto.PubkeyToAddress(v.privateKey.PublicKey)

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

func (v *VRFModule) BeginBlock(ctx sdk.Context) {

	//isWorker := func(extra []byte) bool {
	//	return len(extra) > 32 && len(extra[32:]) >= common.ExtraSeal && bytes.Equal(extra[32:97], make([]byte, common.ExtraSeal))
	//}
	//header := ctx.Backend().CurrentHeader()
	//if isWorker(ctx.Backend().CurrentHeader().Extra) {
	//	// Generate vrf proof
	//	if value, err := v.GenerateNonce(header.Number, header.ParentHash); nil != err {
	//		return err
	//	} else {
	//		header.Nonce = types.EncodeNonce(value)
	//	}
	//} else {
	//	blockHash = header.CacheHash()
	//	// Verify vrf proof
	//	pk := header.CachePublicKey()
	//	if pk == nil {
	//		return errors.New("failed to get the public key of the block producer")
	//	}
	//	if err := v.VerifyVrf(pk, header.Number, header.ParentHash, blockHash, header.Nonce.Bytes()); nil != err {
	//		return err
	//	}
	//}

}
func (v *VRFModule) EndBlock(ctx sdk.Context) {

	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		v.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return
	}

	header := ctx.Backend().CurrentHeader()

	blockNumber := header.Number.Uint64()
	// get nonceAndProof by block (After the `pushNonceAndProof` transaction was executed)
	nonceAndProof, err := v.getCurrentNonceAndProof(wctx, blockNumber)
	if nil != err {
		v.logger.Error("Failed to get current nonceAndProof", "blockNumber", blockNumber, "error", err)
		return
	}

	// Extract the validator public key of the build block based on the signature in the block header
	sign := header.Signature()
	sealhash := header.SealHash().Bytes()
	pk, err := crypto.SigToPub(sealhash, sign)
	if err != nil {
		log.Error("can not sigToPub", "blockNumber", blockNumber, "err", err)
		return
	}

	// verify nonce and
	if err := v.VerifyVrf(wctx, blockNumber, nonceAndProof, pk); nil != err {
		panic(err)
	}
}

func (v *VRFModule) GenerateNonceAndProof(ctx sdk.WorkerContext, blockNumber uint64) ([]byte, error) {
	nonceAndProof, err := vrfInternal.GenerateNonceAndProof(ctx.StateDB(), address.VRFHandlerAddress, blockNumber, v.privateKey)
	if nil != err {
		v.logger.Error(err.Error(), "blockNumber", blockNumber)
		return nil, err
	}
	v.logger.Info("Succeed to generate vrf nonce and proof", "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof),
		"nodeId", enode.PublicKeyToIDv0(&(v.privateKey.PublicKey)).String())
	return nonceAndProof, nil
}

func (v *VRFModule) VerifyVrf(ctx sdk.WorkerContext, blockNumber uint64, nonceAndProof []byte, key *ecdsa.PublicKey) error {

	previousNonce, err := v.getPreviousNonce(ctx, blockNumber)
	if nil != err {
		v.logger.Error("Failed to get previous vrf nonce", "blockNumber", blockNumber, "error", err)
		return err
	}

	if err := vrfInternal.VerifyVrf(nonceAndProof, previousNonce, key); nil != err {
		v.logger.Error(err.Error(), "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof), "data", previousNonce.Hex())
		return err
	}
	v.logger.Info("Succeed to verify vrf nonceAndProof", "blockNumber", blockNumber, "nonceAndProof", hex.EncodeToString(nonceAndProof), "data", previousNonce.Hex())
	return nil
}

func (v *VRFModule) getPreviousNonce(ctx sdk.WorkerContext, blockNumber uint64) (basecommon.Hash, error) {
	return vrfInternal.GetPreviousNonce(ctx.StateDB(), address.VRFHandlerAddress, blockNumber)
}

func (v *VRFModule) getPreviousNonceAndProof(ctx sdk.WorkerContext, blockNumber uint64) ([]byte, error) {
	return vrfInternal.GetPreviousNonceAndProof(ctx.StateDB(), address.VRFHandlerAddress, blockNumber)
}

func (v *VRFModule) getCurrentNonce(ctx sdk.WorkerContext, blockNumber uint64) (basecommon.Hash, error) {
	return vrfInternal.GetCurrentNonce(ctx.StateDB(), address.VRFHandlerAddress, blockNumber)
}

func (v *VRFModule) getCurrentNonceAndProof(ctx sdk.WorkerContext, blockNumber uint64) ([]byte, error) {
	return vrfInternal.GetCurrentNonceAndProof(ctx.StateDB(), address.VRFHandlerAddress, blockNumber)
}

func (v *VRFModule) createPushNonceAndProofTx(ctx sdk.Context, nonceAndProof []byte) (*types.Transaction, error) {

	input, err := contracts.Abi.Methods["pushNonceAndProof"].Inputs.Pack(nonceAndProof)
	if nil != err {
		return nil, err
	}
	from := crypto.PubkeyToAddress(v.privateKey.PublicKey)
	txNonce, err := ctx.Backend().GetPoolNonce(from)
	if nil != err {
		return nil, err
	}

	tx := types.NewTransaction(txNonce, address.VRFHandlerAddress, nil, 100000, big.NewInt(0), input)
	chainId, _ := ctx.Backend().ChainId()
	signer := types.NewEIP155Signer(chainId)
	tx, err = types.SignTx(tx, signer, v.privateKey)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (v *VRFModule) newVRFHandlerCallContract(ctx sdk.Context, header *types.Header) (*contracts.VRFHandler, error) {
	from := crypto.PubkeyToAddress(v.privateKey.PublicKey)
	evm, _, err := ctx.Backend().GetEVM(message.NewOnlyCallMessage(from), header)
	if err != nil {
		return nil, err
	}
	return contracts.NewVRFHandler(evm, vm.NewContract(vm.AccountRef(from), vm.AccountRef(address.VRFHandlerAddress), big.NewInt(0), 1000000), true)
}

func decodePrivateKey(keystoreFile, passwordFile string) (*keystore.Key, error) {
	if keystoreFile == "" {
		return nil, fmt.Errorf("statesync.keystore not set")
	}
	if passwordFile == "" {
		return nil, fmt.Errorf("statesync.password not set")

	}

	key, err := sdkcommon.DecryptKey(keystoreFile, passwordFile)
	if err != nil {
		return nil, err
	}
	return key, nil
}
