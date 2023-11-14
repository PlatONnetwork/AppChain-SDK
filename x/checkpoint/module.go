package checkpoint

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"

	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractsapi/checkpoint_manager"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/storage"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/utils"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	_ module.Module                = (*Module)(nil)
	_ module.ConsensusExtendModule = (*Module)(nil)
	_ types.EventSubscriber        = (*Module)(nil)
)

type Module struct {
	checkpointManagerAddr common.Address
	l2StateSenderAddr     common.Address

	logger log.Logger
	store  *storage.Storage

	staking   types.Staking
	signer    types.Signer
	txRealyer types.TxRelayer
	extraVote types.ExtraVote
}

func NewModule(
	checkpointManagerAddr common.Address,
	l2StateSenderAddr common.Address,
	store *storage.Storage,
	staking types.Staking,
	signer types.Signer,
	txRealyer types.TxRelayer,
	extraVote types.ExtraVote,
	stateEvent types.StateEvent) *Module {
	m := &Module{
		checkpointManagerAddr: checkpointManagerAddr,
		l2StateSenderAddr:     l2StateSenderAddr,
		logger:                log.New("module", types.ModuleName),
		store:                 store,
		staking:               staking,
		signer:                signer,
		txRealyer:             txRealyer,
		extraVote:             extraVote,
	}

	stateEvent.Subscribe(m)
	return m
}

func (m *Module) Name() string {
	return types.ModuleName
}

func (m *Module) GetLogFilters() map[common.Address][]common.Hash {
	return map[common.Address][]common.Hash{
		m.l2StateSenderAddr: {contractsapi.L2StateSenderABI.Events["L2StateSynced"].ID},
	}
}

func (m *Module) ProcessLog(header *coretypes.Header, log *coretypes.Log) error {
	checkpoint, err := m.store.GetCheckpoint(header.Number.Uint64())
	if err != nil {
		m.logger.Error("Failed to get checkpoint", "number", header.Number, "err", err)
		return err
	}

	//  exit events that happened in epoch ending blocks,
	// should be added to the tree of the next epoch
	epoch := checkpoint.EpochNumber + 1
	number := checkpoint.BlockNumber + 1

	exitEvent, err := contractsapi.DecodeExitEvent(log, epoch, number)
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

	if m.staking.IsEndOfEpoch(header.Number.Uint64()) {
		currentValidators := m.staking.GetValidator(header.Number.Uint64())
		currentAccountSet := types.NewAccountSet(currentValidators)
		currentValidatorHash, err := currentAccountSet.Hash()
		if err != nil {
			m.logger.Error("Failed to current validator set hash", "err", err)
			return []byte{}
		}

		nextValidators := m.staking.GetValidator(header.Number.Uint64() + 1)
		nextAccountSet := types.NewAccountSet(nextValidators)
		nextValidatorHash, err := nextAccountSet.Hash()
		if err != nil {
			m.logger.Error("Failed to get next validator set hash", "err", err)
			return []byte{}
		}

		eventRoot, err := m.BuildEventRoot(epoch)
		if err != nil {
			m.logger.Error("Failed to build event root", "epoch", epoch, "err", err)
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
		m.logger.Info("Make checkpoint data",
			"blockNumber", header.Number,
			"viewNumber", view,
			"epochNumber", epoch,
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

	m.logger.Debug("Verify extend data", "epoch", epoch, "view", view,
		"number", "index", sdkCtx.BlockIndex(), header.Number, "hash", header.Hash())

	if m.staking.IsEndOfEpoch(header.Number.Uint64()) {
		var checkpoint types.CheckpointData
		if err := checkpoint.UnmarshalRLP(data); err != nil {
			return common.ZeroHash, err
		}
		if checkpoint.EpochNumber != epoch {
			return common.ZeroHash, fmt.Errorf("mismatch epoch(checkpoint:%d,actual:%d)", checkpoint.EpochNumber, epoch)
		}
		if checkpoint.ViewNumber != view {
			return common.ZeroHash, fmt.Errorf("mismatch view(checkpoint:%d,actual:%d)", checkpoint.ViewNumber, view)
		}

		currentValidators := m.staking.GetValidator(header.Number.Uint64())
		currentAccountSet := types.NewAccountSet(currentValidators)
		currentValidatorHash, err := currentAccountSet.Hash()
		if err != nil {
			m.logger.Error("Failed to current validator set hash", "err", err)
			return common.ZeroHash, err
		}

		if checkpoint.CurrentValidatorsHash != currentValidatorHash {
			m.logger.Error("Current validators hash mismatch",
				"checkpoint.currentValidatorHash", checkpoint.CurrentValidatorsHash.TerminalString(),
				"actual", currentValidatorHash.TerminalString())
			return common.ZeroHash, fmt.Errorf("mismatch currentValidatorsHash(checkpoint:%s,actual:%s)",
				checkpoint.CurrentValidatorsHash.TerminalString(),
				currentValidatorHash.TerminalString())
		}

		nextValidators := m.staking.GetValidator(header.Number.Uint64() + 1)
		nextAccountSet := types.NewAccountSet(nextValidators)
		nextValidatorHash, err := nextAccountSet.Hash()
		if err != nil {
			m.logger.Error("Failed to get next validator set hash", "err", err)
			return common.ZeroHash, err
		}

		if checkpoint.NextValidatorsHash != nextValidatorHash {
			m.logger.Error("Next validators hash mismatch",
				"checkpoint.nextValidatorHash", checkpoint.NextValidatorsHash.TerminalString(),
				"actual", nextValidatorHash.TerminalString())
			return common.ZeroHash, fmt.Errorf("mismatch nextValidatorsHash(checkpoint:%s,actual:%s)", checkpoint.NextValidatorsHash.TerminalString(), nextValidatorHash.TerminalString())
		}

		eventRoot, err := m.BuildEventRoot(sdkCtx.Epoch())
		if err != nil {
			m.logger.Error("Failed to build event root", "epoch", sdkCtx.Epoch(), "err", err)
			return common.ZeroHash, err
		}

		if eventRoot != checkpoint.EventRoot {
			m.logger.Error("Event root dismatch", "checkpoint.EventRoot", checkpoint.EventRoot, "actual", eventRoot)
			return common.ZeroHash, fmt.Errorf("mismatch event root(checkpoint:%s,acutal:%s)", checkpoint.EventRoot.TerminalString(), eventRoot.TerminalString())
		}

		if err := m.store.InsertCheckpoint(header.Number.Uint64(),
			&types.StorageCheckpointData{
				CheckpointData: &checkpoint,
			}); err != nil {
			m.logger.Error("Failed to insert checkpoint",
				"epoch", epoch,
				"view", view,
				"blockIndex", checkpoint.BlockIndex,
				"blockNumber", checkpoint.BlockNumber,
				"blockHash", checkpoint.BlockHash,
				"err", err,
			)
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

	if m.staking.IsEndOfEpoch(block.BlockNum()) {
		checkpoint, err := m.store.GetCheckpoint(block.BlockNum())
		if err != nil {
			m.logger.Error("Failed to get checkpoint from store",
				"epoch", block.Epoch,
				"view", block.ViewNumber,
				"blockNumber", block.BlockNum(),
				"blockHash", block.Block.Hash(),
				"err", err,
			)
			return
		}
		signature, bitmap, extendRoot, err := aggSignatures(votes)
		if err != nil {
			m.logger.Error("Failed to aggregate votes signature",
				"epoch", block.Epoch,
				"view", block.ViewNumber,
				"blockNumber", block.BlockNum(),
				"blockHash", block.Block.Hash(),
				"err", err,
			)
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
		latestHeader := sdkCtx.Backend().GetHeaderByNumber(latestNumber)
		if latestHeader == nil {
			m.logger.Error("Failed to get block header", "checkpoint number", latestNumber)
			return
		}
		go func(header *coretypes.Header, epoch uint64) {
			if err := m.submitCheckpoint(header); err != nil {
				m.logger.Error("Failed to submit checkpoint",
					"checkpoint number", header.Number,
					"epoch", epoch,
					"err", err)
			}
		}(latestHeader, sdkCtx.Epoch())
	}
}

func (m *Module) submitCheckpoint(latestHeader *coretypes.Header) error {
	lastCheckpointBlockNumber, err := getCurrentCheckpointBlock(m.txRealyer, m.checkpointManagerAddr)
	if err != nil {
		return err
	}

	if lastCheckpointBlockNumber > latestHeader.Number.Uint64() {
		// node is out of sync (haven't reached the tip of the chain), so even though it is a proposer,
		// it would checkpoint block that is already checkpointed and transaction would fail anyway
		return nil
	}

	m.logger.Debug("submitCheckpoint invoked...",
		"latest checkpoint block", lastCheckpointBlockNumber,
		"checkpoint block", latestHeader.Number)

	blocksOfEpoch := m.staking.BlocksOfEpoch()
	initialBlockNumber := lastCheckpointBlockNumber + blocksOfEpoch

	for blockNumber := initialBlockNumber; blockNumber <= latestHeader.Number.Uint64(); {
		checkpoint, err := m.store.GetCheckpoint(blockNumber)
		if err != nil {
			return err
		}

		if err := m.encodeAndSendCheckpoint(checkpoint); err != nil {
			return err
		}

		blockNumber = initialBlockNumber + blocksOfEpoch
	}
	return nil
}

func (m *Module) encodeAndSendCheckpoint(checkpoint *types.StorageCheckpointData) error {
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

	nextValidators := m.staking.GetValidator(checkpoint.BlockNumber + 1)
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
	leafIndex, proof, err := m.extraVote.GetProof(checkpoint.EpochNumber, checkpoint.ViewNumber, leaf)
	if err != nil {
		return err
	}

	data, err := submitAbiType.Inputs.Pack(checkpointMetadata, cp, checkpoint.Signature, checkpoint.Bitmap, newValidatorSet, leafIndex, proof)
	if err != nil {
		return err
	}

	receipt, err := m.txRealyer.SendTransaction(coretypes.NewTx(&coretypes.LegacyTx{
		To:   &m.checkpointManagerAddr,
		Data: data,
	}), m.signer)
	if err != nil {
		return err
	}

	if receipt.Status == coretypes.ReceiptStatusFailed {
		return fmt.Errorf("checkpoint submission transaction failed for block %d", checkpoint.BlockNumber)
	}
	m.logger.Debug("send checkpoint txn success", "checkpoint number", checkpoint.BlockNumber, "gasUsed", receipt.GasUsed)
	return nil
}

func (m *Module) BuildEventRoot(epoch uint64) (common.Hash, error) {
	exitEvents, err := m.store.GetExitEventsByEpoch(epoch)
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
	for _, vote := range votes {
		if extendHash == common.ZeroHash {
			extendHash = vote.ExtendHash
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
