package pevm

import (
	"math"
	"sync"
	"sync/atomic"
)

type LockedTxStatus struct {
	sync.RWMutex
	*TxStatus
}

func initializeLockedTxStatus(blockSize int32) []*LockedTxStatus {
	ts := make([]*LockedTxStatus, blockSize)
	for i := range ts {
		ts[i] = &LockedTxStatus{
			TxStatus: &TxStatus{
				Incarnation: 0,
				Status:      ReadyToExecute,
			},
		}
	}
	return ts
}

type LockedTxDependent struct {
	sync.RWMutex
	deps []int32
}

func initializeTxDependents(blockSize int32) []*LockedTxDependent {
	td := make([]*LockedTxDependent, blockSize)
	for i := range td {
		td[i] = &LockedTxDependent{
			deps: make([]int32, 0),
		}
	}
	return td
}

type Scheduler struct {
	blockSize        int32
	txsStatus        []*LockedTxStatus
	txsDependents    []*LockedTxDependent
	executionIdx     atomic.Int32
	validationIdx    atomic.Int32
	minValidationIdx atomic.Int32
	numValidated     atomic.Int32
	aborted          atomic.Bool
}

func NewScheduler(blockSize int32) *Scheduler {
	s := &Scheduler{
		blockSize:     blockSize,
		txsStatus:     initializeLockedTxStatus(blockSize),
		txsDependents: initializeTxDependents(blockSize),
	}
	s.executionIdx.Store(0)
	s.validationIdx.Store(int32(blockSize))
	s.minValidationIdx.Store(blockSize)
	s.numValidated.Store(0)
	s.aborted.Store(false)
	return s
}

func (s *Scheduler) Abort() {
	s.aborted.Store(true)
}

func (s *Scheduler) NextTask() Task {
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
			txIdx := fetchAddI32(&s.validationIdx, 1)
			if txIdx < s.blockSize {
				tx := s.txsStatus[txIdx]
				tx.Lock()

				if tx.Status == ReadyToExecute {
					tx.Status = Executing
					tx.Unlock()
					return NewExection(TxVersion{
						TxIdx:         txIdx,
						TxIncarnation: tx.Incarnation,
					})
				}

				// Start a typical validation task
				if tx.Status == Executed || tx.Status == Validated {
					tx.Unlock()
					return NewValidation(TxVersion{
						TxIdx:         txIdx,
						TxIncarnation: tx.Incarnation,
					})
				}

				// Validation index is still catching up so continue a
				// new loop iteration to refetch the latest indices
				// before deciding again.
				if tx.Status == Aborting {
					tx.Unlock()
					continue
				}
				tx.Unlock()
			}
		}

		// Prioritize execution task
		if txVer := s.tryExecute(fetchAddI32(&s.executionIdx, 1)); txVer != nil {
			return NewExection(*txVer)
		}
	}
	return nil
}

func (s *Scheduler) tryExecute(txIdx int32) *TxVersion {
	if txIdx < s.blockSize {
		tx := s.txsStatus[txIdx]
		tx.Lock()
		defer tx.Unlock()
		return &TxVersion{
			TxIdx:         txIdx,
			TxIncarnation: tx.Incarnation,
		}
	}
	return nil
}

func (s *Scheduler) AddDependency(txIdx, blockingIdx int32) bool {
	// This is an important lock to prevent a race condition where the blocking
	// transaction completes re-execution before this dependency can be added.
	blockingTx := s.txsStatus[blockingIdx]
	blockingTx.RLock()
	if blockingTx.Status == Executed || blockingTx.Status == Validated {
		blockingTx.RUnlock()
		return false
	}
	blockingTx.RUnlock()

	tx := s.txsStatus[txIdx]
	tx.Lock()
	tx.Status = Aborting
	tx.Unlock()

	blockingDeps := s.txsDependents[blockingIdx]
	blockingDeps.Lock()
	blockingDeps.deps = append(blockingDeps.deps, txIdx)
	blockingDeps.Unlock()

	return true
}

func (s *Scheduler) SetReadyStatus(txIdx int32) {
	tx := s.txsStatus[txIdx]
	tx.Lock()
	defer tx.Unlock()
	tx.Status = ReadyToExecute
	tx.Incarnation += 1
}

func (s *Scheduler) FinishExecution(txVersion TxVersion, flags FinishExecFlags) Task {
	tx := s.txsStatus[txVersion.TxIdx]
	tx.Lock()
	defer tx.Unlock()

	deps := s.txsDependents[txVersion.TxIdx]
	deps.Lock()
	for _, txIdx := range deps.deps {
		s.SetReadyStatus(txIdx)
		fetchMinI32(&s.executionIdx, txIdx)
	}
	deps.deps = deps.deps[:0]
	deps.Unlock()

	minValidationIdx := s.minValidationIdx.Load()
	if flags.Has(NeedValidation) {
		minValidationIdx = int32(math.Min(float64(fetchMinI32(&s.minValidationIdx, txVersion.TxIdx)),
			float64(txVersion.TxIdx)))
	}
	// Have found a min validation index to even bother
	if minValidationIdx < s.blockSize {
		// Must re-validate from min as this transaction is lower
		if txVersion.TxIdx < minValidationIdx {
			if flags.Has(WroteNewLocation) {
				fetchMinI32(&s.validationIdx, minValidationIdx)
			}
		} else if txVersion.TxIdx < s.validationIdx.Load() {
			// Validate from this transaction as it's in between min and the current
			// validation index.
			if flags.Has(WroteNewLocation) {
				fetchMinI32(&s.validationIdx, txVersion.TxIdx+1)
			}
			if flags.Has(NeedValidation) {
				tx.Status = Executed
				return NewValidation(txVersion)
			}
			tx.Status = Validated
			s.numValidated.Add(1)
		}
		// Don't need to validate anything if the current validation index is
		// lower or equal -- it will catch up later.
	}

	if flags.Has(NeedValidation) {
		tx.Status = Executed
	} else {
		tx.Status = Validated
		s.numValidated.Add(1)
	}

	return nil
}

func (s *Scheduler) TryValidationAbort(txVersion TxVersion) bool {
	tx := s.txsStatus[txVersion.TxIdx]
	tx.Lock()
	defer tx.Unlock()

	if tx.Status == Validated {
		s.numValidated.Add(-1)
	}

	aborting := tx.Status == Executed || tx.Status == Validated
	if aborting {
		tx.Status = Aborting
	}
	return aborting
}

func (s *Scheduler) FinishValidation(txVersion TxVersion, aborted bool) Task {
	if aborted {
		s.SetReadyStatus(txVersion.TxIdx)
		fetchMinI32(&s.validationIdx, txVersion.TxIdx+1)
		if s.executionIdx.Load() > int32(txVersion.TxIdx) {
			return NewExection(*s.tryExecute(int32(txVersion.TxIdx)))
		}
	} else {
		tx := s.txsStatus[txVersion.TxIdx]
		tx.Lock()
		defer tx.Unlock()
		if tx.Status == Executed {
			tx.Status = Validated
			s.numValidated.Add(1)
		}
	}
	return nil
}

func fetchMinI32(a *atomic.Int32, newVal int32) int32 {
	current := a.Load()
	for newVal < current {
		if a.CompareAndSwap(current, newVal) {
			return current
		}
		current = a.Load()
	}
	return current
}

func fetchAddI32(a *atomic.Int32, delta int32) int32 {
	old := a.Load()
	for {
		newVal := old + delta
		if a.CompareAndSwap(old, newVal) {
			return old
		}
		old = a.Load()
	}
}
