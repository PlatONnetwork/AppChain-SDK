package checkpoint

import (
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/contractapi"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/storage"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/types"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
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

	checkpointAbi, _ = abi.JSON(strings.NewReader(contractapi.ContractapiABI))
)

type Module struct {
	checkpointManagerAddr common.Address

	logger log.Logger
	store  *storage.Storage

	staking   types.Staking
	signer    types.Signer
	txRealyer types.TxRelayer
	extraVote types.ExtraVote
}

func NewModule(
	checkpointManagerAddr common.Address,
	store *storage.Storage,
	staking types.Staking,
	signer types.Signer,
	txRealyer types.TxRelayer,
	extraVote types.ExtraVote) *Module {
	return &Module{
		checkpointManagerAddr: checkpointManagerAddr,
		logger:                log.New("module", types.ModuleName),
		store:                 store,
		staking:               staking,
		signer:                signer,
		txRealyer:             txRealyer,
		extraVote:             extraVote,
	}
}

func (m *Module) Name() string {
	return types.ModuleName
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

		checkpoint := &types.CheckpointData{
			ViewNumber:            view,
			EpochNumber:           epoch,
			BlockIndex:            blockIndex,
			BlockNumber:           header.Number.Uint64(),
			BlockHash:             header.Hash(),
			CurrentValidatorsHash: currentValidatorHash,
			NextValidatorsHash:    nextValidatorHash,
			EventRoot:             common.ZeroHash, // TODO: get event root
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

		// TODO: check event root
		if err := m.store.SetCheckpoint(header.Number.Uint64(),
			&types.StorageCheckpointData{
				CheckpointData: &checkpoint,
			}); err != nil {
			m.logger.Error("Failed to set checkpoint",
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
		m.store.SetCheckpoint(block.BlockNum(), checkpoint)
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
	submitAbiType := checkpointAbi.Methods["submit"]

	checkpointMetadata := contractapi.ICheckpointManagerCheckpointMetadata{
		BlockHash:               checkpoint.BlockHash,
		BlockIndex:              checkpoint.BlockIndex,
		CurrentValidatorSetHash: checkpoint.CurrentValidatorsHash,
	}

	cp := contractapi.ICheckpointManagerCheckpoint{
		Epoch:       checkpoint.EpochNumber,
		ViewNumber:  checkpoint.ViewNumber,
		BlockNumber: checkpoint.BlockNumber,
		EventRoot:   checkpoint.EventRoot,
		ExtendRoot:  checkpoint.ExtendRoot,
	}

	nextValidators := m.staking.GetValidator(checkpoint.BlockNumber + 1)
	accountSet := types.NewAccountSet(nextValidators)
	newValidatorSet := make([]contractapi.ICheckpointManagerValidator, len(accountSet))
	for i, account := range accountSet {
		var blsKey [2]*big.Int
		b := account.BlsKey.Serialize()
		blsKey[0] = big.NewInt(0).SetBytes(b[:16])
		blsKey[1] = big.NewInt(0).SetBytes(b[16:])
		newValidatorSet[i] = contractapi.ICheckpointManagerValidator{
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

	receipt, err := m.txRealyer.SendTransaction(coretypes.NewTransaction(0, m.checkpointManagerAddr, nil, 0, nil, data), m.signer)
	if  err != nil {
		return err
	}

	if receipt.Status == coretypes.ReceiptStatusFailed {
		return fmt.Errorf("checkpoint submission transaction failed for block %d", checkpoint.BlockNumber)
	}
	m.logger.Debug("send checkpoint txn success", "checkpoint number", checkpoint.BlockNumber, "gasUsed", receipt.GasUsed)
	return nil
}

func getCurrentCheckpointBlock(relayer types.TxRelayer, checkpointManagerAddr common.Address) (uint64, error) {
	input := checkpointAbi.Methods["currentCheckpointBlockNumber"].ID
	currentCheckpointBlockRaw, err := relayer.Call(common.ZeroAddr, checkpointManagerAddr, input)
	if err != nil {
		return 0, fmt.Errorf("invoke currentCheckpointBlockNumber on rootchain: %w", err)
	}

	currentCheckpointBlock, err := strconv.ParseUint(currentCheckpointBlockRaw, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("convert current checkpoint block number '%s' to number: %w",
			currentCheckpointBlockRaw, err)
	}

	return currentCheckpointBlock, nil
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
