package checkpoint

import (
	"fmt"
	"math/big"

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
	ctypes "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

var (
	_ module.Module     = (*Module)(nil)
	_ module.InitModule = (*Module)(nil)
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

func (m *Module) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: types.ModuleName,
			Version:   "1.0",
			Service:   NewRpcService(m),
			Public:    true,
		},
	}
}

func (m *Module) GetLogFilters() map[common.Address][]common.Hash {
	return map[common.Address][]common.Hash{
		m.l2StateSenderAddr: {contractsapi.L2StateSenderABI.Events["L2StateSynced"].ID},
	}
}

func (m *Module) ProcessLog(header *coretypes.Header, qc *ctypes.QuorumCert, log *coretypes.Log) error {
	epoch := qc.Epoch
	block := qc.BlockNumber
	var i uint64 = 0
	for ; i < uint64(types.CheckpointCommitDis); i++ {
		if m.staking.IsEndOfEpoch(block + i) {
			block = block + i + 1
			epoch = epoch + 1
			break
		}
	}

	exitEvent, err := contractsapi.DecodeExitEvent(log, epoch, block)
	if err != nil {
		m.logger.Error("Failed to decode exit event", "err", err)
		return err
	}

	return m.store.InsertExitEvent(exitEvent)
}

func (m *Module) ExtendData(ctx sdk.ConsensusContext) []byte {
	header := ctx.Header()
	view := ctx.View()
	epoch := ctx.Epoch()
	blockIndex := ctx.BlockIndex()
	logger := m.logger.New("epoch", epoch, "view", view, "index", blockIndex, "number", header.Number, "hash", header.Hash())

	logger.Info("Extend data")

	blockNumber := header.Number.Uint64()
	if m.staking.IsEndOfEpoch(blockNumber) {
		currentValidators, err := m.staking.GetValidator(ctx, blockNumber)
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

		nextValidators, err := m.staking.GetValidator(ctx, blockNumber+1)
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

		eventRoot, err := m.BuildEventRoot(epoch)
		if err != nil {
			logger.Error("Failed to build event root", "epoch", epoch, "err", err)
			return []byte{}
		}

		chainID, _ := ctx.Backend().ChainId()

		checkpoint := &types.CheckpointData{
			ChainID:               chainID.Uint64(),
			ViewNumber:            view,
			EpochNumber:           epoch,
			BlockIndex:            blockIndex,
			BlockNumber:           header.Number.Uint64(),
			BlockHash:             header.Hash(),
			CurrentValidatorsHash: currentValidatorHash,
			NextValidatorsHash:    nextValidatorHash,
			EventRoot:             eventRoot,
		}
		logger.Info("Make checkpoint data", "checkpoint data", checkpoint.String())
		return checkpoint.MarshalRLP()
	}
	return []byte{}
}

func (m *Module) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) error {
	header := ctx.Header()
	view := ctx.View()
	epoch := ctx.Epoch()
	blockIndex := ctx.BlockIndex()
	logger := m.logger.New("epoch", epoch, "view", view, "index", blockIndex, "number", header.Number, "hash", header.Hash())

	logger.Info("Verify extend data")

	blockNumber := header.Number.Uint64()
	if m.staking.IsEndOfEpoch(blockNumber) {
		var checkpoint types.CheckpointData
		if err := checkpoint.UnmarshalRLP(data); err != nil {
			logger.Error("Failed to unmarshal rlp", "data", fmt.Sprintf("%x", data), "err", err)
			return err
		}
		chainID, _ := ctx.Backend().ChainId()
		if checkpoint.ChainID != chainID.Uint64() {
			return fmt.Errorf("mismatch chainID(checkpoint:%d,actual:%d)", checkpoint.ChainID, chainID)
		}
		if checkpoint.EpochNumber != epoch {
			return fmt.Errorf("mismatch epoch(checkpoint:%d,actual:%d)", checkpoint.EpochNumber, epoch)
		}
		if checkpoint.ViewNumber != view {
			return fmt.Errorf("mismatch view(checkpoint:%d,actual:%d)", checkpoint.ViewNumber, view)
		}
		if checkpoint.BlockIndex != blockIndex {
			return fmt.Errorf("mismatch blockIndex(checkpoint:%d,actual:%d)", checkpoint.BlockIndex, blockIndex)
		}
		if checkpoint.BlockNumber != blockNumber {
			return fmt.Errorf("mismatch blockNumber(checkpoint:%d,actual:%d)", checkpoint.BlockNumber, header.Number)
		}
		if checkpoint.BlockHash != header.Hash() {
			return fmt.Errorf("mismatch blockHash(checkpoint:%s,actual:%s)", checkpoint.BlockHash.String(), header.Hash().String())
		}

		currentValidators, err := m.staking.GetValidator(ctx, blockNumber)
		if err != nil {
			logger.Error("Failed to get current round validators", "err", err)
			return err
		}
		currentAccountSet := types.NewAccountSet(currentValidators)
		currentValidatorHash, err := currentAccountSet.Hash()
		if err != nil {
			logger.Error("Failed to get current validator set hash", "err", err)
			return err
		}

		if checkpoint.CurrentValidatorsHash != currentValidatorHash {
			logger.Error("Current validators hash mismatch",
				"checkpoint.currentValidatorHash", checkpoint.CurrentValidatorsHash.TerminalString(),
				"actual", currentValidatorHash.TerminalString())
			return fmt.Errorf("mismatch currentValidatorsHash(checkpoint:%s,actual:%s)",
				checkpoint.CurrentValidatorsHash.TerminalString(),
				currentValidatorHash.TerminalString())
		}

		nextValidators, err := m.staking.GetValidator(ctx, blockNumber+1)
		if err != nil {
			logger.Error("Failed to get next round validators", "err", err)
			return err
		}
		nextAccountSet := types.NewAccountSet(nextValidators)
		nextValidatorHash, err := nextAccountSet.Hash()
		if err != nil {
			logger.Error("Failed to get next validator set hash", "err", err)
			return err
		}

		if checkpoint.NextValidatorsHash != nextValidatorHash {
			logger.Error("Next validators hash mismatch",
				"checkpoint.nextValidatorHash", checkpoint.NextValidatorsHash.TerminalString(),
				"actual", nextValidatorHash.TerminalString())
			return fmt.Errorf("mismatch nextValidatorsHash(checkpoint:%s,actual:%s)", checkpoint.NextValidatorsHash.TerminalString(), nextValidatorHash.TerminalString())
		}

		eventRoot, err := m.BuildEventRoot(epoch)
		if err != nil {
			logger.Error("Failed to build event root", "epoch", epoch, "err", err)
			return err
		}

		if eventRoot != checkpoint.EventRoot {
			logger.Error("Event root dismatch", "checkpoint.EventRoot", checkpoint.EventRoot, "actual", eventRoot)
			return fmt.Errorf("mismatch event root(checkpoint:%s,acutal:%s)", checkpoint.EventRoot.TerminalString(), eventRoot.TerminalString())
		}

		logger.Info("Verify checkpoint successfully", "checkpoint", checkpoint.String())
	}
	return nil
}

func (m *Module) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
}

func (m *Module) OnCommit(ctx sdk.ConsensusContext, block *coretypes.Block) error {
	logger := m.logger.New("epoch", ctx.Epoch(), "view", ctx.View(), "index", ctx.BlockIndex(), "number", ctx.Header().Number, "hash", ctx.Header().Hash())
	logger.Info("OnCommit")

	blockNumber := block.NumberU64()
	if m.staking.IsEndOfEpoch(blockNumber) {
		_, qc, err := ctypes.DecodeExtra(block.ExtraData())
		if err != nil {
			logger.Error("Failed to decode block extra data", "err", err)
			return err
		}

		if ctx.IsProposer() {
			go func(number uint64) {
				if err := m.submitCheckpoint(ctx, number, qc); err != nil {
					logger.Error("Failed to submit checkpoint", "err", err)
				}
			}(blockNumber)
		}
	}
	return nil
}

func (m *Module) submitCheckpoint(ctx sdk.ConsensusContext, latestNumber uint64, latestQC *ctypes.QuorumCert) error {
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
		qc := latestQC
		if blockNumber != latestNumber {
			block := ctx.Backend().GetBlockByNumber(blockNumber)
			if block == nil {
				m.logger.Error("Failed to get block", "number", blockNumber)
				return fmt.Errorf("block(%d) not found", blockNumber)
			}
			_, qc, err = ctypes.DecodeExtra(block.ExtraData())
			if err != nil {
				m.logger.Error("Failed to decode extra data", "number", block.Number(), "hash", block.Hash(), "err", err)
				return err
			}
		}

		checkpoint, err := m.rebuildCheckpoint(ctx, qc)
		if err != nil {
			m.logger.Error("Failed to get checkpoint", "err", err)
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

func (m *Module) rebuildCheckpoint(ctx sdk.ConsensusContext, qc *ctypes.QuorumCert) (*types.Checkpoint, error) {
	currentValidators, err := m.staking.GetValidator(ctx, qc.BlockNumber)
	if err != nil {
		return nil, err
	}
	currentAccountSet := types.NewAccountSet(currentValidators)
	currentValidatorsHash, err := currentAccountSet.Hash()
	if err != nil {
		return nil, err
	}

	nextValidators, err := m.staking.GetValidator(ctx, qc.BlockNumber+1)
	if err != nil {
		return nil, err
	}
	nextAccountSet := types.NewAccountSet(nextValidators)
	nextValidatorsHash, err := nextAccountSet.Hash()
	if err != nil {
		return nil, err
	}

	eventRoot, err := m.BuildEventRoot(qc.Epoch)
	if err != nil {
		return nil, err
	}

	sig := qc.Signature.Bytes()
	var blsSig bls.Sign
	blsSig.Deserialize(sig)

	chainID, _ := ctx.Backend().ChainId()

	return &types.Checkpoint{
		CheckpointData: &types.CheckpointData{
			ChainID:               chainID.Uint64(),
			EpochNumber:           qc.Epoch,
			ViewNumber:            qc.ViewNumber,
			BlockIndex:            qc.BlockIndex,
			BlockNumber:           qc.BlockNumber,
			BlockHash:             qc.BlockHash,
			CurrentValidatorsHash: currentValidatorsHash,
			NextValidatorsHash:    nextValidatorsHash,
			EventRoot:             eventRoot,
		},
		ExtendRoot: qc.ExtendHash,
		Signature:  blsSig.SerializeUncompressed(),
		Bitmap:     qc.ValidatorSet.Bytes(),
	}, nil
}

func (m *Module) encodeAndSendCheckpoint(ctx sdk.ConsensusContext, checkpoint *types.Checkpoint) error {
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
		newValidatorSet[i] = checkpoint_manager.ICheckpointManagerValidator{
			Address: account.Address,
			BlsKey:  account.BlsKey.SerializeUncompressed(),
		}
	}

	leaf := checkpoint.MarshalRLP()
	leafIndex, proof, err := m.extraVote.GetProof(checkpoint.EpochNumber, checkpoint.ViewNumber, checkpoint.BlockIndex, leaf)
	if err != nil {
		m.logger.Error("Failed to get extend data proof", "epoch", checkpoint.EpochNumber, "view", checkpoint.ViewNumber, "index", checkpoint.BlockIndex, "err", err)
		return err
	}

	data, err := contractsapi.CheckpointManagerABI.Pack("submit", checkpointMetadata, cp, checkpoint.Signature, checkpoint.Bitmap, newValidatorSet, big.NewInt(0).SetUint64(leafIndex), proof)
	if err != nil {
		return err
	}
	m.logger.Info("Submit checkpoint", "checkpoint", checkpoint.String())

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
	m.logger.Info("Send checkpoint txn success", "checkpoint", checkpoint.String(), "txHash", receipt.TxHash.Hex(), "gasUsed", receipt.GasUsed)
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
