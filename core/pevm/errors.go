package pevm

import "fmt"

type ErrBlocking struct {
	TxIdx uint32
	err   error
}

func NewErrBlocking(txIdx uint32, err error) ErrBlocking {
	return ErrBlocking{
		TxIdx: txIdx,
		err:   err,
	}
}

func (e ErrBlocking) Error() string {
	return fmt.Sprintf("transaction %d blocking: %v", e.TxIdx, e.err)
}

type ErrExecution struct {
	TxIdx uint32
	err   error
}

func NewErrExecution(txIdx uint32, err error) ErrExecution {
	return ErrExecution{
		TxIdx: txIdx,
		err:   err,
	}
}

func (e ErrExecution) Error() string {
	return fmt.Sprintf("transaction %d execution error: %v", e.TxIdx, e.err)
}
