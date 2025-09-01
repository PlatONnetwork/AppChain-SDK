package stateevent

import (
	"errors"
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/stateevent/storage"
	"github.com/PlatONnetwork/PlatON-Go/common"
	ctypes "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const ModuleVersion uint64 = 0
const ModuleName = "eventState"

var (
	_ module.Module               = (*Module)(nil)
	_ module.BlockCommitterModule = (*Module)(nil)
)

// EventSubscriber specifies functions needed for a component to subscribe to stateEvent
type EventSubscriber interface {
	// GetLogFilters returns a map of log filters for getting desired events,
	// where the key is the address of contract that emits desired events,
	// and the value is a slice of signatures of events we want to get.
	GetLogFilters() map[common.Address][]common.Hash

	// ProcessLog is used to handle a log defined in GetLogFilters, provid
	ProcessLog(ctx sdk.ConsensusContext, header *coretypes.Header, qc *ctypes.QuorumCert, log *coretypes.Log) error
}

type Module struct {
	logger              log.Logger
	store               *storage.Storage
	subscriberIDCounter uint64

	subscribers map[uint64]EventSubscriber
	allFilters  map[common.Address]map[common.Hash][]uint64
}

func NewModule(kvStore store.Store) *Module {
	return &Module{
		logger:      log.New("module", ModuleName),
		store:       storage.NewStorage(kvStore),
		subscribers: make(map[uint64]EventSubscriber),
		allFilters:  make(map[common.Address]map[common.Hash][]uint64, 0),
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}

func (m *Module) OnCommit(ctx sdk.ConsensusContext, block *coretypes.Block) error {
	m.logger.Debug("OnCommit", "number", block.Number(), "hash", block.Hash())
	if ctx.Backend() == nil {
		m.logger.Error("Empty backend")
		return errors.New("empty backend")
	}
	lastProcessedBlock, err := m.store.GetLastProcessedEventsBlock()
	if err != nil {
		m.logger.Error("Failed to get last proccessed events block", "number", block.Number(), "hash", block.Hash(), "err", err)
		return err
	}

	if err := m.getEventsFromBlocks(ctx, lastProcessedBlock, block); err != nil {
		m.logger.Error("Failed to get events from blocks", "lastProcessedBlock", lastProcessedBlock, "latestBlock", block.Number(), "err", err)
		return err
	}
	if err := m.store.InsertLastProcessedEventBlock(block.NumberU64()); err != nil {
		m.logger.Error("Failed to insert last processed event block", "number", block.Number(), "err", err)
		return err
	}
	return nil
}

func (m *Module) Subscribe(subscriber EventSubscriber) {
	m.subscriberIDCounter++
	subscriberID := m.subscriberIDCounter
	m.subscribers[subscriberID] = subscriber
	m.logger.Debug("Subscribe state event", "id", subscriberID)

	for address, filters := range subscriber.GetLogFilters() {
		existingAddressFilters, exist := m.allFilters[address]
		if !exist {
			existingAddressFilters = make(map[common.Hash][]uint64, 0)
			m.allFilters[address] = existingAddressFilters
			m.logger.Debug("Subscribe success", "address", address.Hex())
		}

		for _, f := range filters {
			m.logger.Debug("Subscribe success", "address", address.Hex(), "topic", f.Hex())
			existingAddressFilters[f] = append(existingAddressFilters[f], subscriberID)
		}
	}
}

func (m *Module) getEventsFromBlocks(ctx sdk.ConsensusContext, lastProcessedBlock uint64, latestBlock *coretypes.Block) error {
	if err := m.getEventsFromBlocksRange(ctx, lastProcessedBlock+1, latestBlock.NumberU64()-1); err != nil {
		return err
	}
	_, qc, err := ctypes.DecodeExtra(latestBlock.ExtraData())
	if err != nil {
		return err
	}
	return m.getEventsFromReceipts(ctx, latestBlock.Header(), qc, ctx.Backend().ReadReceipts(latestBlock.Header().SealHash()))
}

func (m *Module) getEventsFromBlocksRange(ctx sdk.ConsensusContext, from, to uint64) error {
	for i := from; i <= to; i++ {
		blockHeader := ctx.Backend().GetHeaderByNumber(i)
		if blockHeader == nil {
			return fmt.Errorf("block header not found(number: %d)", i)
		}

		block := ctx.Backend().GetBlock(blockHeader.Hash(), blockHeader.Number.Uint64())
		if block == nil {
			return fmt.Errorf("block not found(number: %d,hash: %s)", blockHeader.Number, blockHeader.Hash().String())
		}

		_, qc, err := ctypes.DecodeExtra(block.ExtraData())
		if err != nil {
			return err
		}

		receipts := ctx.Backend().GetReceiptsByHash(blockHeader.Hash())
		if err := m.getEventsFromReceipts(ctx, blockHeader, qc, receipts); err != nil {
			return err
		}
	}
	return nil
}

func (m *Module) getEventsFromReceipts(ctx sdk.ConsensusContext, blockHeader *coretypes.Header, qc *ctypes.QuorumCert, receipts coretypes.Receipts) error {
	m.logger.Debug("Get event", "number", qc.BlockNumber, "receipts", receipts.Len(), "filters", len(m.allFilters))
	for _, receipt := range receipts {
		if receipt.Status != coretypes.ReceiptStatusSuccessful {
			continue
		}

		for _, log := range receipt.Logs {
			logFilters, isRelevantLog := m.allFilters[log.Address]
			if !isRelevantLog {
				continue
			}
			m.logger.Debug("Find relevant log", "address", log.Address.Hex(), "event", log.Topics[0].Hex())
			for logFilter, subscribers := range logFilters {
				if log.Topics[0] == logFilter {
					for _, subscriber := range subscribers {
						if err := m.subscribers[subscriber].ProcessLog(ctx, blockHeader, qc, log); err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
