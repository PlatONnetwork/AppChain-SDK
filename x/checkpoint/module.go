package checkpoint

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"

	sdkcom "github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi/checkpoint_manager"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/storage"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/types"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/utils"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

var (
	_ module.Module                = (*Module)(nil)
	_ module.InitModule            = (*Module)(nil)
	_ module.ConsensusExtendModule = (*Module)(nil)
)

func AddModuleInitFlags(app *cli.App) {
	app.Flags = append(app.Flags, KeystoreFlag)
	app.Flags = append(app.Flags, PasswordFlag)
}

type Module struct {
	checkpointManagerAddr common.Address
	l2StateSenderAddr     common.Address

	keystoreFile string
	passwordFile string
	key          *keystore.Key

	logger log.Logger
	store  *storage.Storage

	staking   types.Staking
	txRelayer types.TxRelayer
	extraVote types.ExtraVote
	l1        types.L1
}

func NewModule(
	cliCtx *cli.Context,
	store store.Store,
	staking types.Staking,
	txRelayer types.TxRelayer,
	extraVote types.ExtraVote,
	stateEvent types.StateEvent,
	l1 types.L1) (*Module, error) {
	m := &Module{
		keystoreFile: cliCtx.GlobalString(KeystoreFlag.Name),
		passwordFile: cliCtx.GlobalString(PasswordFlag.Name),
		logger:       log.New("module", types.ModuleName),
		store:        storage.NewStorage(store),
		staking:      staking,
		txRelayer:    txRelayer,
		extraVote:    extraVote,
		l1:           l1,
	}

	checkpointAddr, err := l1.GetCheckpointAddress()
	if err != nil {
		return nil, err
	}

	// FIXME: set L2StateSender contract address
	m.l2StateSenderAddr = common.ZeroAddr
	m.checkpointManagerAddr = checkpointAddr

	stateEvent.Subscribe(m)
	return m, nil
}

func (m *Module) Name() string {
	return types.ModuleName
}

func (m *Module) Init() error {
	if err := m.txRelayer.Init(); err != nil {
		return err
	}

	if m.keystoreFile == "" {
		return fmt.Errorf("checkpoint.keystore not set")
	}
	if m.passwordFile == "" {
		return fmt.Errorf("checkpoint.password not set")
	}

	key, err := sdkcom.DecryptKey(m.keystoreFile, m.passwordFile)
	if err != nil {
		return err
	}
	m.key = key
	return nil
}

func (m *Module) GetLogFilters() map[common.Address][]common.Hash {
	return map[common.Address][]common.Hash{
		m.l2StateSenderAddr: {contractsapi.L2StateSenderABI.Events["L2StateSynced"].ID},
	}
}

func (m *Module) ProcessLog(header *coretypes.Header, log *coretypes.Log) error {
	exitEvent, err := contractsapi.DecodeExitEvent(log, header.Number.Uint64())
	if err != nil {
		m.logger.Error("Failed to decode exit event", "err", err)
		return err
	}

	return m.store.InsertExitEvent(exitEvent)
}

func (m *Module) ExtendData(ctx sdk.Context) []byte {
	sdkCtx, ok := ctx.(sdk.ConsensusContext)
	if !ok {
		m.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return []byte{}
	}

	header := sdkCtx.Header()
	view := sdkCtx.View()
	epoch := sdkCtx.Epoch()
	blockIndex := sdkCtx.BlockIndex()
	logger := m.logger.New("epoch", epoch, "view", view, "index", blockIndex, "number", header.Number, "hash", header.Hash())

	logger.Debug("Extend data")

	if m.staking.IsEndOfEpoch(header.Number.Uint64()) {
		currentValidators, err := m.staking.GetValidator(ctx, header.Number.Uint64())
		if err != nil {
			logger.Error("Failed to get current round valdiators", "err", err)
			return []byte{}
		}
		currentAccountSet := types.NewAccountSet(currentValidators)
		currentValidatorHash, err := currentAccountSet.Hash()
		if err != nil {
			logger.Error("Failed to get current validator set hash", "err", err)
			return []byte{}
		}

		nextValidators, err := m.staking.GetValidator(ctx, header.Number.Uint64()+1)
		if err != nil {
			logger.Error("Failed to get next round valdiators", "err", err)
			return []byte{}
		}
		nextAccountSet := types.NewAccountSet(nextValidators)
		nextValidatorHash, err := nextAccountSet.Hash()
		if err != nil {
			logger.Error("Failed to get next validator set hash", "err", err)
			return []byte{}
		}

		lastCheckpointBlockNumber, err := getCurrentCheckpointBlock(m.txRelayer, m.checkpointManagerAddr)
		if err != nil {
			logger.Error("Failed to get current checkpoint block", "err", err)
			return []byte{}
		}

		// ExitEvent insert store when block committing.
		// Block consensus sequence: qc -> locked -> committed
		// The checkpoint number is qcblock,
		// so the range is [commitblock, qcblock-2]
		if lastCheckpointBlockNumber > types.CheckpointCommitDis {
			lastCheckpointBlockNumber = lastCheckpointBlockNumber - types.CheckpointCommitDis
		}
		end := header.Number.Uint64() - types.CheckpointCommitDis

		eventRoot, err := m.BuildEventRoot(lastCheckpointBlockNumber, end)
		if err != nil {
			logger.Error("Failed to build event root", "epoch", epoch, "err", err)
			return []byte{}
		}

		checkpoint := &types.CheckpointData{
			ViewNumber:            view,
			EpochNumber:           epoch,
			BlockIndex:            blockIndex,
			BlockNumber:           header.Number.Uint64(),
			BlockHash:             header.Hash(),
			CurrentValidatorsHash: currentValidatorHash,
			NextValidatorsHash:    nextValidatorHash,
			EventRoot:             eventRoot,
		}
		logger.Info("Make checkpoint data",
			"currentValidatorsHash", checkpoint.CurrentValidatorsHash,
			"nextValidatorsHash", checkpoint.NextValidatorsHash,
			"eventRoot", checkpoint.EventRoot)
		return checkpoint.MarshalRLP()
	}
	return []byte{}
}

func (m *Module) VerifyExtendData(ctx sdk.Context, data []byte) (common.Hash, error) {
	sdkCtx, ok := ctx.(sdk.ConsensusContext)
	if !ok {
		m.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return common.ZeroHash, errors.New("unexpected sdk context")
	}

	header := sdkCtx.Header()
	view := sdkCtx.View()
	epoch := sdkCtx.Epoch()
	logger := m.logger.New("epoch", epoch, "view", view, "index", sdkCtx.BlockIndex(), "number", header.Number, "hash", header.Hash())

	logger.Debug("Verify extend data")

	if m.staking.IsEndOfEpoch(header.Number.Uint64()) {
		var checkpoint types.CheckpointData
		if err := checkpoint.UnmarshalRLP(data); err != nil {
			logger.Error("Failed to unmarshal rlp", "data", fmt.Sprintf("%x", data), "err", err)
			return common.ZeroHash, err
		}
		if checkpoint.EpochNumber != epoch {
			return common.ZeroHash, fmt.Errorf("mismatch epoch(checkpoint:%d,actual:%d)", checkpoint.EpochNumber, epoch)
		}
		if checkpoint.ViewNumber != view {
			return common.ZeroHash, fmt.Errorf("mismatch view(checkpoint:%d,actual:%d)", checkpoint.ViewNumber, view)
		}

		currentValidators, err := m.staking.GetValidator(ctx, header.Number.Uint64())
		if err != nil {
			logger.Error("Failed to get current round validators", "err", err)
			return common.ZeroHash, err
		}
		currentAccountSet := types.NewAccountSet(currentValidators)
		currentValidatorHash, err := currentAccountSet.Hash()
		if err != nil {
			logger.Error("Failed to get current validator set hash", "err", err)
			return common.ZeroHash, err
		}

		if checkpoint.CurrentValidatorsHash != currentValidatorHash {
			logger.Error("Current validators hash mismatch",
				"checkpoint.currentValidatorHash", checkpoint.CurrentValidatorsHash.TerminalString(),
				"actual", currentValidatorHash.TerminalString())
			return common.ZeroHash, fmt.Errorf("mismatch currentValidatorsHash(checkpoint:%s,actual:%s)",
				checkpoint.CurrentValidatorsHash.TerminalString(),
				currentValidatorHash.TerminalString())
		}

		nextValidators, err := m.staking.GetValidator(ctx, header.Number.Uint64()+1)
		if err != nil {
			logger.Error("Failed to get next round validators", "err", err)
			return common.Hash{}, err
		}
		nextAccountSet := types.NewAccountSet(nextValidators)
		nextValidatorHash, err := nextAccountSet.Hash()
		if err != nil {
			logger.Error("Failed to get next validator set hash", "err", err)
			return common.ZeroHash, err
		}

		if checkpoint.NextValidatorsHash != nextValidatorHash {
			logger.Error("Next validators hash mismatch",
				"checkpoint.nextValidatorHash", checkpoint.NextValidatorsHash.TerminalString(),
				"actual", nextValidatorHash.TerminalString())
			return common.ZeroHash, fmt.Errorf("mismatch nextValidatorsHash(checkpoint:%s,actual:%s)", checkpoint.NextValidatorsHash.TerminalString(), nextValidatorHash.TerminalString())
		}

		lastCheckpointBlockNumber, err := getCurrentCheckpointBlock(m.txRelayer, m.checkpointManagerAddr)
		if err != nil {
			logger.Error("Failed to get current checkpoint block", "err", err)
			return common.ZeroHash, err
		}

		// ExitEvent insert store when block committing.
		// Block consensus sequence: qc -> locked -> committed
		// The checkpoint number is qcblock,
		// so the range is [commitblock, qcblock-2]
		if lastCheckpointBlockNumber > types.CheckpointCommitDis {
			lastCheckpointBlockNumber = lastCheckpointBlockNumber - types.CheckpointCommitDis
		}
		end := header.Number.Uint64() - types.CheckpointCommitDis

		eventRoot, err := m.BuildEventRoot(lastCheckpointBlockNumber, end)
		if err != nil {
			logger.Error("Failed to build event root", "epoch", sdkCtx.Epoch(), "err", err)
			return common.ZeroHash, err
		}

		if eventRoot != checkpoint.EventRoot {
			logger.Error("Event root dismatch", "checkpoint.EventRoot", checkpoint.EventRoot, "actual", eventRoot)
			return common.ZeroHash, fmt.Errorf("mismatch event root(checkpoint:%s,acutal:%s)", checkpoint.EventRoot.TerminalString(), eventRoot.TerminalString())
		}

		if err := m.store.InsertCheckpoint(header.Number.Uint64(),
			&types.StorageCheckpointData{
				CheckpointData: &checkpoint,
			}); err != nil {
			logger.Error("Failed to insert checkpoint", "err", err)
			return common.ZeroHash, err
		}
		return checkpoint.Hash()
	}
	return common.ZeroHash, nil
}

func (m *Module) PrepareQC(ctx sdk.Context, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	sdkCtx, ok := ctx.(sdk.ConsensusContext)
	if !ok {
		m.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return
	}

	logger := m.logger.New("epoch", sdkCtx.Epoch(), "view", sdkCtx.View(), "index", sdkCtx.BlockIndex(), "number", sdkCtx.Header().Number, "hash", sdkCtx.Header().Hash())
	logger.Debug("Prepare QC")

	if m.staking.IsEndOfEpoch(block.BlockNum()) {
		checkpoint, err := m.store.GetCheckpoint(block.BlockNum())
		if err != nil {
			m.logger.Error("Failed to get checkpoint from store", "err", err)
			return
		}
		signature, bitmap, extendRoot, err := aggSignatures(votes)
		if err != nil {
			logger.Error("Failed to aggregate votes signature", "err", err)
			return
		}
		checkpoint.ExtendRoot = extendRoot
		checkpoint.Signature = signature
		checkpoint.Bitmap = bitmap
		m.store.InsertCheckpoint(block.BlockNum(), checkpoint)
		return
	}

	latestNumber := block.Block.NumberU64() - types.CheckpointCommitDis
	if m.staking.IsEndOfEpoch(latestNumber) && sdkCtx.IsProposer() {
		go func(number uint64, epoch uint64) {
			if err := m.submitCheckpoint(ctx, number); err != nil {
				logger.Error("Failed to submit checkpoint", "checkpoint number", number, "err", err)
			}
		}(latestNumber, sdkCtx.Epoch())
	}
}

func (m *Module) submitCheckpoint(ctx sdk.Context, latestNumber uint64) error {
	lastCheckpointBlockNumber, err := getCurrentCheckpointBlock(m.txRelayer, m.checkpointManagerAddr)
	if err != nil {
		return err
	}

	if lastCheckpointBlockNumber > latestNumber {
		// node is out of sync (haven't reached the tip of the chain), so even though it is a proposer,
		// it would checkpoint block that is already checkpointed and transaction would fail anyway
		return nil
	}

	m.logger.Debug("submitCheckpoint invoked...",
		"latest checkpoint block", lastCheckpointBlockNumber,
		"checkpoint block", latestNumber)

	blocksOfEpoch := m.staking.BlocksOfEpoch()
	initialBlockNumber := lastCheckpointBlockNumber + blocksOfEpoch

	for blockNumber := initialBlockNumber; blockNumber <= latestNumber; {
		checkpoint, err := m.store.GetCheckpoint(blockNumber)
		if err != nil {
			return err
		}

		if err := m.encodeAndSendCheckpoint(ctx, checkpoint); err != nil {
			m.logger.Error("Failed to encode and send checkpoint", "checkpoint", checkpoint.String(), "err", err)
			return err
		}

		blockNumber = initialBlockNumber + blocksOfEpoch
	}
	return nil
}

func (m *Module) encodeAndSendCheckpoint(ctx sdk.Context, checkpoint *types.StorageCheckpointData) error {
	submitAbiType := contractsapi.CheckpointManagerABI.Methods["submit"]

	checkpointMetadata := checkpoint_manager.ICheckpointManagerCheckpointMetadata{
		BlockHash:               checkpoint.BlockHash,
		BlockIndex:              checkpoint.BlockIndex,
		CurrentValidatorSetHash: checkpoint.CurrentValidatorsHash,
	}

	cp := checkpoint_manager.ICheckpointManagerCheckpoint{
		Epoch:       checkpoint.EpochNumber,
		ViewNumber:  checkpoint.ViewNumber,
		BlockNumber: checkpoint.BlockNumber,
		EventRoot:   checkpoint.EventRoot,
		ExtendRoot:  checkpoint.ExtendRoot,
	}

	nextValidators, err := m.staking.GetValidator(ctx, checkpoint.BlockNumber+1)
	accountSet := types.NewAccountSet(nextValidators)
	newValidatorSet := make([]checkpoint_manager.ICheckpointManagerValidator, len(accountSet))
	for i, account := range accountSet {
		var blsKey [2]*big.Int
		b := account.BlsKey.Serialize()
		blsKey[0] = big.NewInt(0).SetBytes(b[:16])
		blsKey[1] = big.NewInt(0).SetBytes(b[16:])
		newValidatorSet[i] = checkpoint_manager.ICheckpointManagerValidator{
			Address: account.Address,
			BlsKey:  blsKey,
		}
	}

	leaf, _ := checkpoint.CheckpointData.Hash()
	leafIndex, proof, err := m.extraVote.GetProof(checkpoint.EpochNumber, checkpoint.ViewNumber, checkpoint.BlockIndex, leaf.Bytes())
	if err != nil {
		return err
	}

	data, err := submitAbiType.Inputs.Pack(checkpointMetadata, cp, checkpoint.Signature, checkpoint.Bitmap, newValidatorSet, big.NewInt(0).SetUint64(leafIndex), proof)
	if err != nil {
		return err
	}

	receipt, err := m.txRelayer.SendTransaction(coretypes.NewTx(&coretypes.LegacyTx{
		To:   &m.checkpointManagerAddr,
		Data: data,
	}), m.key)
	if err != nil {
		return err
	}

	if receipt.Status == coretypes.ReceiptStatusFailed {
		return fmt.Errorf("checkpoint submission transaction failed for block %d", checkpoint.BlockNumber)
	}
	m.logger.Debug("send checkpoint txn success", "checkpoint", checkpoint.String(), "gasUsed", receipt.GasUsed)
	return nil
}

func (m *Module) BuildEventRoot(start, end uint64) (common.Hash, error) {
	exitEvents, err := m.store.GetExitEventsByNumberRange(start, end)
	if err != nil {
		return common.ZeroHash, err
	}

	if len(exitEvents) == 0 {
		return common.ZeroHash, nil
	}

	tree, err := createExitTree(exitEvents)
	if err != nil {
		return common.ZeroHash, err
	}
	return tree.Hash(), nil
}

func getCurrentCheckpointBlock(relayer types.TxRelayer, checkpointManagerAddr common.Address) (uint64, error) {
	input := contractsapi.CheckpointManagerABI.Methods["currentCheckpointBlockNumber"].ID
	currentCheckpointBlockRaw, err := relayer.Call(common.ZeroAddr, checkpointManagerAddr, input)
	if err != nil {
		return 0, fmt.Errorf("invoke currentCheckpointBlockNumber on rootchain: %w", err)
	}

	currentCheckpointBlock := big.NewInt(0).SetBytes(currentCheckpointBlockRaw)
	if err != nil {
		return 0, fmt.Errorf("convert current checkpoint block number '%s' to number: %w",
			currentCheckpointBlockRaw, err)
	}

	return currentCheckpointBlock.Uint64(), nil
}

func aggSignatures(votes map[uint32]*protocols.PrepareVote) (signature []byte, bitmap []byte, extendHash common.Hash, err error) {
	var aggSig bls.Sign
	var sets utils.BitArray
	aggSigInit := false
	for _, vote := range votes {
		if extendHash == common.ZeroHash {
			extendHash = vote.ExtendHash
		}

		if !aggSigInit {
			if err = aggSig.Deserialize(vote.Signature[:]); err != nil {
				return
			}
			aggSigInit = true
			sets.SetIndex(vote.NodeIndex(), true)
			continue
		}

		var sig bls.Sign
		if err = sig.Deserialize(vote.Signature[:]); err != nil {
			return
		}
		aggSig.Add(&sig)
		sets.SetIndex(vote.NodeIndex(), true)
	}
	signature = aggSig.Serialize()
	bitmap = sets.Bytes()
	return
}

func createExitTree(exitEvents []*contractsapi.ExitEvent) (*merkle.MerkleTree, error) {
	numOfEvents := len(exitEvents)
	data := make([][]byte, numOfEvents)

	for i := 0; i < numOfEvents; i++ {
		b, err := exitEvents[i].Encode()
		if err != nil {
			return nil, err
		}
		data = append(data, b)
	}
	return merkle.NewMerkleTree(data)
}
