package pevm

import (
	"math"
	"sync"
	"sync/atomic"

	"github.com/PlatONnetwork/AppChain-SDK/x/pevm/types"
)

type LockedTxStatus struct {
	sync.RWMutex
	*types.TxStatus
}

func initializeLockedTxStatus(blockSize uint32) []*LockedTxStatus {
	ts := make([]*LockedTxStatus, blockSize)
	for i := range ts {
		ts[i] = &LockedTxStatus{
			TxStatus: &types.TxStatus{
				Incarnation: 0,
				Status:      types.ReadyToExecute,
			},
		}
	}
	return ts
}

type LockedTxDependent struct {
	sync.RWMutex
	deps []uint32
}

func initializeTxDependents(blockSize uint32) []*LockedTxDependent {
	td := make([]*LockedTxDependent, blockSize)
	for i := range td {
		td[i] = &LockedTxDependent{
			deps: make([]uint32, 0),
		}
	}
	return td
}

type Scheduler struct {
	blockSize        uint32
	txsStatus        []*LockedTxStatus
	txsDependents    []*LockedTxDependent
	executionIdx     atomic.Uint32
	validationIdx    atomic.Uint32
	minValidationIdx atomic.Uint32
	numValidated     atomic.Uint32
	aborted          atomic.Bool
}

func NewScheduler(blockSize uint32) *Scheduler {
	s := &Scheduler{
		blockSize:     blockSize,
		txsStatus:     initializeLockedTxStatus(blockSize),
		txsDependents: initializeTxDependents(blockSize),
	}
	s.executionIdx.Store(0)
	s.validationIdx.Store(blockSize)
	s.minValidationIdx.Store(0)
	s.numValidated.Store(0)
	s.aborted.Store(false)
	return s
}

func (s *Scheduler) Abort() {
	s.aborted.Store(true)
}

func (s *Scheduler) NextTask() types.Task {
	for !s.aborted.Load() {
		var (
			executionIdx  = s.executionIdx.Load()
			validationIdx = s.validationIdx.Load()
		)
		if executionIdx >= s.blockSize && validationIdx >= s.blockSize {
			if s.numValidated.Load() >= s.blockSize-s.minValidationIdx.Load() {
				break
			}
			continue
		}

		// Prioritize a validation task to minimize re-execution
		if validationIdx < executionIdx {
			txIdx := s.validationIdx.Add(1)
			if txIdx < s.blockSize {
				tx := s.txsStatus[txIdx]
				tx.Lock()

				if tx.Status == types.ReadyToExecute {
					tx.Status = types.Executing
					tx.Unlock()
					return types.NewExection(types.TxVersion{
						TxIdx:         txIdx,
						TxIncarnation: tx.Incarnation,
					})
				}

				// Start a typical validation task
				if tx.Status == types.Executed || tx.Status == types.Validated {
					tx.Unlock()
					return types.NewValidation(types.TxVersion{
						TxIdx:         txIdx,
						TxIncarnation: tx.Incarnation,
					})
				}

				// Validation index is still catching up so continue a
				// new loop iteration to refetch the latest indices
				// before deciding again.
				if tx.Status == types.Aborting {
					tx.Unlock()
					continue
				}
				tx.Unlock()
			}
		}

		// Prioritize execution task
		if txVer := s.tryExecute(s.executionIdx.Add(1)); txVer != nil {
			return types.NewExection(*txVer)
		}
	}
	return nil
}

func (s *Scheduler) tryExecute(txIdx uint32) *types.TxVersion {
	if txIdx < s.blockSize {
		tx := s.txsStatus[txIdx]
		tx.Lock()
		defer tx.Unlock()
		return &types.TxVersion{
			TxIdx:         txIdx,
			TxIncarnation: tx.Incarnation,
		}
	}
	return nil
}

func (s *Scheduler) AddDependency(txIdx, blockingIdx uint32) bool {
	// This is an important lock to prevent a race condition where the blocking
	// transaction completes re-execution before this dependency can be added.
	blockingTx := s.txsStatus[blockingIdx]
	blockingTx.RLock()
	if blockingTx.Status == types.Executed || blockingTx.Status == types.Validated {
		blockingTx.RUnlock()
		return false
	}
	blockingTx.RUnlock()

	tx := s.txsStatus[txIdx]
	tx.Lock()
	tx.Status = types.Aborting
	tx.Unlock()

	blockingDeps := s.txsDependents[blockingIdx]
	blockingDeps.Lock()
	blockingDeps.deps = append(blockingDeps.deps, txIdx)
	blockingDeps.Unlock()

	return true
}

func (s *Scheduler) SetReadyStatus(txIdx uint32) {
	tx := s.txsStatus[txIdx]
	tx.Lock()
	defer tx.Unlock()
	tx.Status = types.ReadyToExecute
	tx.Incarnation += 1
}

func (s *Scheduler) FinishExecution(txVersion types.TxVersion, flags types.FinishExecFlags) types.Task {
	tx := s.txsStatus[txVersion.TxIdx]
	tx.Lock()
	defer tx.Unlock()

	deps := s.txsDependents[txVersion.TxIdx]
	deps.Lock()
	for _, txIdx := range deps.deps {
		s.SetReadyStatus(txIdx)
		fetchMinU32(&s.executionIdx, txIdx)
	}
	deps.deps = deps.deps[:0]
	deps.Unlock()

	minValidationIdx := s.minValidationIdx.Load()
	if flags.Has(types.NeedValidation) {
		minValidationIdx = uint32(math.Min(float64(fetchMinU32(&s.minValidationIdx, txVersion.TxIdx)),
			float64(txVersion.TxIdx)))
	}
	// Have found a min validation index to even bother
	if minValidationIdx < s.blockSize {
		// Must re-validate from min as this transaction is lower
		if txVersion.TxIdx < minValidationIdx {
			if flags.Has(types.WroteNewLocation) {
				fetchMinU32(&s.validationIdx, minValidationIdx)
			}
		} else if txVersion.TxIdx < s.validationIdx.Load() {
			// Validate from this transaction as it's in between min and the current
			// validation index.
			if flags.Has(types.WroteNewLocation) {
				fetchMinU32(&s.validationIdx, txVersion.TxIdx+1)
			}
			if flags.Has(types.NeedValidation) {
				tx.Status = types.Executed
				return types.NewExection(txVersion)
			}
			tx.Status = types.Validated
			s.numValidated.Add(1)
		}
		// Don't need to validate anything if the current validation index is
		// lower or equal -- it will catch up later.
	}

	if flags.Has(types.NeedValidation) {
		tx.Status = types.Executed
	} else {
		tx.Status = types.Validated
		s.numValidated.Add(1)
	}

	return nil
}

func (s *Scheduler) TryValidationAbort(txVersion types.TxVersion) bool {
	tx := s.txsStatus[txVersion.TxIdx]
	tx.Lock()
	defer tx.Unlock()

	if tx.Status == types.Validated {
		s.numValidated.Add(1)
	}

	aborting := tx.Status == types.Executed || tx.Status == types.Validated
	if aborting {
		tx.Status = types.Aborting
	}
	return aborting
}

func (s *Scheduler) FinishValidation(txVersion types.TxVersion, aborted bool) types.Task {
	if aborted {
		s.SetReadyStatus(txVersion.TxIdx)
		fetchMinU32(&s.validationIdx, txVersion.TxIdx+1)
		if s.executionIdx.Load() > txVersion.TxIdx {
			return types.NewExection(*s.tryExecute(txVersion.TxIdx))
		}
	} else {
		tx := s.txsStatus[txVersion.TxIdx]
		tx.Lock()
		defer tx.Unlock()
		if tx.Status == types.Executed {
			tx.Status = types.Validated
			s.numValidated.Add(1)
		}
	}
	return nil
}

func fetchMinU32(a *atomic.Uint32, newVal uint32) uint32 {
	current := a.Load()
	for newVal < current {
		if a.CompareAndSwap(current, newVal) {
			return current
		}
		current = a.Load()
	}
	return current
}
